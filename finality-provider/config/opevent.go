package config

import (
	"github.com/ethereum/go-ethereum/common"
	"time"
)

var (
	defaultScanSize           = uint32(500)
	defaultEthRpc             = "127.0.0.1:8545"
	defaultL2OutputOracleAddr = "0x0"
	defaultPollInterval       = 5 * time.Second
	defaultStartHeight        = uint64(1)
)

type OpEventConfig struct {
	ChainId            uint
	BlockStep          uint64
	BufferSize         uint32 `long:"buffersize" description:"The maximum number of ethereum blocks that can be stored in the buffer"`
	EthRpc             string `long:"ethrpc" description:"The rpc uri of ethereum"`
	Contracts          []common.Address
	L2OutputOracleAddr string        `long:"l2outputoracleaddr" description:"The contract address of L2OutputOracle address"`
	PollInterval       time.Duration `long:"pollinterval" description:"The interval between each polling of blocks; the value should be set depending on the block production time but could be set smaller for quick catching up"`
	ScanStartHeight    uint64        `long:"scantartheight" description:"The height from which we start polling the chain"`

	OPFinalityGadgetAddress  string `long:"op-finality-gadget" description:"the contract address of the op-finality-gadget"`
	BabylonFinalityGadgetRpc string `long:"babylon-finality-gadget-rpc" description:"the rpc address of babylon op finality gadget"`
}

func DefaultOpEventConfig() OpEventConfig {
	return OpEventConfig{
		BufferSize:         defaultScanSize,
		EthRpc:             defaultEthRpc,
		L2OutputOracleAddr: defaultL2OutputOracleAddr,
		PollInterval:       defaultPollInterval,
		ScanStartHeight:    defaultStartHeight,
	}
}
