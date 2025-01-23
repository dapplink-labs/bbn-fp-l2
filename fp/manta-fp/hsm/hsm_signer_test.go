package hsm

import (
	"context"
	"fmt"

	"log"
	"math/big"
	"testing"

	kms "cloud.google.com/go/kms/apiv1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/api/option"
)

// cloud hsm public key address:
var (
	addr    = ""
	nodeurl = ""
	apiName = ""
)

// This test will fail because we lack the mantle-666-keystore.json
func TestManagedKey_NewEthereumTransactor(t *testing.T) {
	fmt.Println("this is a simple gcp-cloudhsm demo")
	ethClient, err := ethclient.Dial(nodeurl)
	if err != nil {
		log.Fatal(err)
	}

	tx, chainID := constructTx(addr, ethClient)

	testSigner := types.NewEIP155Signer(chainID) // Mumbai
	ctx := context.Background()
	apikey := option.WithCredentialsFile("mantle-666-keystore.json")
	client, err := kms.NewKeyManagementClient(ctx, apikey)

	mk := &ManagedKey{
		KeyName:      apiName,
		EthereumAddr: common.HexToAddress(addr),
		Gclient:      client,
	}

	signedTx, err := mk.NewEthereumTransactor(context.Background(), testSigner).Signer(mk.EthereumAddr, tx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("hash")
	fmt.Println("hash:" + signedTx.Hash().String())

	v, r, s := signedTx.RawSignatureValues()
	fmt.Printf("signatures, r:%s, s:%s, v:%d\n", r.String(), s.String(), v.Int64())

	err = ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func constructTx(addr string, client *ethclient.Client) (*types.Transaction, *big.Int) {
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(addr))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nonce)

	value := big.NewInt(1000000000000000000)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return tx, chainID
}
