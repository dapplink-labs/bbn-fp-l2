package opstack

import (
	"context"
	"go.uber.org/zap"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum-optimism/optimism/op-proposer/bindings"

	"github.com/Manta-Network/manta-fp/types"
)

type EventProvider struct {
	Log          *zap.Logger
	L2ooFilterer *bindings.L2OutputOracleFilterer
	L2ooABI      *abi.ABI
	EpCtx        context.Context
}

func NewEventProvider(ctx context.Context, logger *zap.Logger) (*EventProvider, error) {

	l2OutputOracleAbi, err := bindings.L2OutputOracleMetaData.GetAbi()
	if err != nil {
		logger.Error("get l2 output oracle abi fail", zap.String("err", err.Error()))
		return nil, err
	}

	l2OutputOracleUnpack, err := bindings.NewL2OutputOracleFilterer(common.Address{}, nil)
	if err != nil {
		logger.Error("new l2output oracle fail", zap.String("err", err.Error()))
		return nil, err
	}

	return &EventProvider{
		Log:          logger,
		L2ooFilterer: l2OutputOracleUnpack,
		L2ooABI:      l2OutputOracleAbi,
		EpCtx:        ctx,
	}, nil
}

func (epd *EventProvider) ProcessStateRootEvent(log types2.Log) (*types.StateRoot, error) {
	outputProposed, err := epd.L2ooFilterer.ParseOutputProposed(log)
	if err != nil {
		epd.Log.Error("parse output proposed fail", zap.String("err", err.Error()))
		return nil, err
	}
	epd.Log.Info("state root event parse success", zap.String("OutputRoot", common.Bytes2Hex(outputProposed.OutputRoot[:])))
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
