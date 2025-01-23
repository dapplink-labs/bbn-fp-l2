package hsm

import (
	"context"
	"encoding/asn1"
	"fmt"
	"math/big"

	kms "cloud.google.com/go/kms/apiv1"
	btcecdsa "github.com/btcsuite/btcd/btcec/v2/ecdsa"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type ManagedKey struct {
	KeyName      string
	EthereumAddr common.Address
	Gclient      *kms.KeyManagementClient
}

func NewManagedKey(ctx context.Context, client *kms.KeyManagementClient, address string, keyName string) (*ManagedKey, error) {
	addr := common.HexToAddress(address)

	return &ManagedKey{
		KeyName:      keyName,
		EthereumAddr: addr,
	}, nil
}

func (mk *ManagedKey) NewEthereumTransactorrWithChainID(ctx context.Context, chainID *big.Int) (*bind.TransactOpts, error) {
	if chainID == nil {
		return nil, bind.ErrNoChainID
	}
	signer := types.LatestSignerForChainID(chainID)
	return &bind.TransactOpts{
		Context: ctx,
		From:    mk.EthereumAddr,
		Signer:  mk.NewEthereumSigner(ctx, signer),
	}, nil
}

func (mk *ManagedKey) NewEthereumTransactor(ctx context.Context, txIdentification types.Signer) *bind.TransactOpts {
	return &bind.TransactOpts{
		Context: ctx,
		From:    mk.EthereumAddr,
		Signer:  mk.NewEthereumSigner(ctx, txIdentification),
	}
}

func (mk *ManagedKey) NewEthereumSigner(ctx context.Context, txIdentification types.Signer) bind.SignerFn {
	return func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if addr != mk.EthereumAddr {
			return nil, bind.ErrNotAuthorized
		}
		sig, err := mk.SignHash(ctx, txIdentification.Hash(tx))
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(txIdentification, sig)
	}
}

func (mk *ManagedKey) SignHash(ctx context.Context, hash common.Hash) ([]byte, error) {
	req := kmspb.AsymmetricSignRequest{
		Name: mk.KeyName,
		Digest: &kmspb.Digest{
			Digest: &kmspb.Digest_Sha256{
				Sha256: hash[:],
			},
		},
	}
	resp, err := mk.Gclient.AsymmetricSign(ctx, &req)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("Google KMS asymmetric sign operation: %w", err)
	}

	var params struct{ R, S *big.Int }
	_, err = asn1.Unmarshal(resp.Signature, &params)
	if err != nil {
		return nil, fmt.Errorf("Google KMS asymmetric signature encoding: %w", err)
	}
	var rLen, sLen int // byte size
	if params.R != nil {
		rLen = (params.R.BitLen() + 7) / 8
	}
	if params.S != nil {
		sLen = (params.S.BitLen() + 7) / 8
	}
	if rLen == 0 || rLen > 32 || sLen == 0 || sLen > 32 {
		return nil, fmt.Errorf("Google KMS asymmetric signature with %d-byte r and %d-byte s denied on size", rLen, sLen)
	}

	var sig [66]byte // + 1-byte header + 1-byte tailer
	params.R.FillBytes(sig[33-rLen : 33])
	params.S.FillBytes(sig[65-sLen : 65])

	// brute force try includes KMS verification
	var recoverErr error
	for recoveryID := byte(0); recoveryID < 2; recoveryID++ {
		sig[0] = recoveryID + 27 // BitCoin header
		btcsig := sig[:65]       // exclude Ethereum 'v' parameter
		pubKey, _, err := btcecdsa.RecoverCompact(btcsig, hash[:])
		if err != nil {
			recoverErr = err
			continue
		}

		if pubKeyAddr(pubKey.SerializeUncompressed()) == mk.EthereumAddr {
			sig[65] = recoveryID // Ethereum 'v' parameter
			return sig[1:], nil  // exclude BitCoin header
		}
	}
	return nil, fmt.Errorf("Google KMS asymmetric signature address recovery mis: %w", recoverErr)
}

func pubKeyAddr(bytes []byte) common.Address {
	digest := crypto.Keccak256(bytes[1:])
	var addr common.Address
	copy(addr[:], digest[12:])
	return addr
}
