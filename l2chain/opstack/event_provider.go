package opstack

import (
	"context"
	"os"

	"github.com/ethereum-optimism/optimism/op-proposer/bindings"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/dapplink-labs/bbn-fp-l2/types"
)

type EventProvider struct {
	Log          log.Logger
	L2ooFilterer *bindings.L2OutputOracleFilterer
	L2ooABI      *abi.ABI
	EpCtx        context.Context
}

func NewEventProvider(ctx context.Context, logConfig oplog.CLIConfig) (*EventProvider, error) {
	opLoger := oplog.NewLogger(os.Stdout, logConfig)
	oplog.SetGlobalLogHandler(opLoger.Handler())

	l2OutputOracleAbi, err := bindings.L2OutputOracleMetaData.GetAbi()
	if err != nil {
		log.Error("get l2 output oracle abi fail", "err", err)
		return nil, err
	}

	l2OutputOracleUnpack, err := bindings.NewL2OutputOracleFilterer(common.Address{}, nil)
	if err != nil {
		log.Error("new l2output oracle fail", "err", err)
		return nil, err
	}

	return &EventProvider{
		Log:          opLoger,
		L2ooFilterer: l2OutputOracleUnpack,
		L2ooABI:      l2OutputOracleAbi,
		EpCtx:        ctx,
	}, nil
}

func (epd *EventProvider) ProcessStateRootEvent(log types2.Log, L1BlockHash common.Hash) (*types.StateRoot, error) {
	outputProposed, err := epd.L2ooFilterer.ParseOutputProposed(log)
	if err != nil {
		epd.Log.Error("parse output proposed fail", "err", err)
		return nil, err
	}
	epd.Log.Info("state root event parse success", "OutputRoot", outputProposed.OutputRoot)
	stateRoot := &types.StateRoot{
		StateRoot:       outputProposed.OutputRoot,
		L2BlockNumber:   outputProposed.L2BlockNumber,
		L2OutputIndex:   outputProposed.L2OutputIndex,
		L1BlockHash:     outputProposed.Raw.BlockHash,
		L1BlockNumber:   outputProposed.Raw.BlockNumber,
		DisputeGameType: 0,
	}
	return stateRoot, nil
}
