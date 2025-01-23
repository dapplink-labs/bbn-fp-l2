package manta_fp

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp/bindings"
	common2 "github.com/Manta-Network/manta-fp/fp/manta-fp/common"
	"github.com/Manta-Network/manta-fp/fp/manta-fp/txmgr"
)

type SignerFn func(context.Context, common.Address, *types.Transaction) (*types.Transaction, error)

type MantaStakingMiddlewareConfig struct {
	EthClient                  *ethclient.Client
	ChainID                    *big.Int
	MantaStakingMiddlewareAddr common.Address
	PrivateKey                 *ecdsa.PrivateKey
	SignerFn                   SignerFn
	NumConfirmations           uint64
	SafeAbortNonceTooLowCount  uint64
	EnableHsm                  bool
	HsmApiName                 string
	HsmCreden                  string
	HsmAddress                 string
}

type MantaStakingMiddleware struct {
	Ctx                               context.Context
	Cfg                               *MantaStakingMiddlewareConfig
	MantaStakingMiddlewareContract    *bindings.MantaStakingMiddleware
	RawMantaStakingMiddlewareContract *bind.BoundContract
	WalletAddr                        common.Address
	MantaStakingMiddlewareABI         *abi.ABI
	txMgr                             txmgr.TxManager
}

func NewMantaStakingMiddleware(cfg *MantaStakingMiddlewareConfig) (*MantaStakingMiddleware, error) {
	mantaStakingMiddlewareContract, err := bindings.NewMantaStakingMiddleware(
		cfg.MantaStakingMiddlewareAddr, cfg.EthClient,
	)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(
		bindings.MantaStakingMiddlewareMetaData.ABI,
	))
	if err != nil {
		return nil, err
	}
	mantaStakingMiddlewareABI, err := bindings.MantaStakingMiddlewareMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	rawTreasureManagerContract := bind.NewBoundContract(
		cfg.MantaStakingMiddlewareAddr, parsed, cfg.EthClient, cfg.EthClient,
		cfg.EthClient,
	)

	txManagerConfig := txmgr.Config{
		ResubmissionTimeout:       time.Second * 5,
		ReceiptQueryInterval:      time.Second,
		NumConfirmations:          cfg.NumConfirmations,
		SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount,
	}

	txMgr := txmgr.NewSimpleTxManager(txManagerConfig, cfg.EthClient)
	var walletAddr common.Address
	if cfg.EnableHsm {
		walletAddr = common.HexToAddress(cfg.HsmAddress)
	} else {
		walletAddr = crypto.PubkeyToAddress(cfg.PrivateKey.PublicKey)
	}
	return &MantaStakingMiddleware{
		Cfg:                               cfg,
		MantaStakingMiddlewareContract:    mantaStakingMiddlewareContract,
		RawMantaStakingMiddlewareContract: rawTreasureManagerContract,
		WalletAddr:                        walletAddr,
		MantaStakingMiddlewareABI:         mantaStakingMiddlewareABI,
		txMgr:                             txMgr,
	}, nil
}

func (msm *MantaStakingMiddleware) UpdateGasPrice(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	var opts *bind.TransactOpts
	var err error
	if !msm.Cfg.EnableHsm {
		opts, err = bind.NewKeyedTransactorWithChainID(
			msm.Cfg.PrivateKey, msm.Cfg.ChainID,
		)
	} else {
		opts, err = common2.NewHSMTransactOpts(ctx, msm.Cfg.HsmApiName,
			msm.Cfg.HsmAddress, msm.Cfg.ChainID, msm.Cfg.HsmCreden)
	}
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
	opts.NoSend = true
	finalTx, err := msm.RawMantaStakingMiddlewareContract.RawTransact(opts, tx.Data())
	if err != nil {
		return nil, err
	}
	return finalTx, nil
}

func (msm *MantaStakingMiddleware) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return msm.Cfg.EthClient.SendTransaction(ctx, tx)
}

func (msm *MantaStakingMiddleware) IsMaxPriorityFeePerGasNotFoundError(err error) bool {
	return strings.Contains(
		err.Error(), common2.ErrMaxPriorityFeePerGasNotFound.Error(),
	)
}

func (msm *MantaStakingMiddleware) registerOperator(ctx context.Context) (*types.Transaction, error) {
	balance, err := msm.Cfg.EthClient.BalanceAt(
		msm.Ctx, msm.WalletAddr, nil,
	)
	if err != nil {
		return nil, err
	}
	log.Info("manta wallet address balance", "balance", balance)

	nonce64, err := msm.Cfg.EthClient.NonceAt(
		msm.Ctx, msm.WalletAddr, nil,
	)
	if err != nil {
		return nil, err
	}
	nonce := new(big.Int).SetUint64(nonce64)
	var opts *bind.TransactOpts
	if !msm.Cfg.EnableHsm {
		opts, err = bind.NewKeyedTransactorWithChainID(
			msm.Cfg.PrivateKey, msm.Cfg.ChainID,
		)
	} else {
		opts, err = common2.NewHSMTransactOpts(ctx, msm.Cfg.HsmApiName,
			msm.Cfg.HsmAddress, msm.Cfg.ChainID, msm.Cfg.HsmCreden)
	}
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.NoSend = true
	tx, err := msm.MantaStakingMiddlewareContract.RegisterOperator(opts)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (msm *MantaStakingMiddleware) RegisterOperator() (*types.Receipt, error) {
	tx, err := msm.registerOperator(msm.Ctx)
	if err != nil {
		return nil, err
	}
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		return msm.UpdateGasPrice(ctx, tx)
	}
	receipt, err := msm.txMgr.Send(
		msm.Ctx, updateGasPrice, msm.SendTransaction,
	)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}
