package opstack

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/op-proposer/bindings"
	"github.com/ethereum-optimism/optimism/op-proposer/contracts"
	"github.com/ethereum-optimism/optimism/op-service/dial"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum-optimism/optimism/op-service/sources/batching"
)

var (
	supportedL2OutputVersion = eth.Bytes32{}
	ErrProposerNotRunning    = errors.New("proposer is not running")
)

type L2OOContract interface {
	Version(*bind.CallOpts) (string, error)
	NextBlockNumber(*bind.CallOpts) (*big.Int, error)
}

type DGFContract interface {
	Version(ctx context.Context) (string, error)
}

type OpStackProviderConfig struct {
	L1EthRpc               string
	RollupRpc              string
	DisputeGameFactoryAddr common.Address
	L2OutputOracleAddr     common.Address
	AllowNonFinalized      bool
	DisputeGameType        uint32
	WaitNodeSync           bool
	LogConfig              oplog.CLIConfig
	NetworkTimeout         time.Duration
}

type OpStackProvider struct {
	opStackCfg     *OpStackProviderConfig
	Log            log.Logger
	RollupProvider dial.RollupProvider

	l2ooContract L2OOContract
	l2ooABI      *abi.ABI

	dgfContract DGFContract
	Multicaller *batching.MultiCaller
}

func NewOpStackProvider(ctx context.Context, cfg *OpStackProviderConfig) (*OpStackProvider, error) {
	opLoger := oplog.NewLogger(os.Stdout, cfg.LogConfig)
	oplog.SetGlobalLogHandler(opLoger.Handler())

	l1Client, err := dial.DialEthClientWithTimeout(ctx, dial.DefaultDialTimeout, opLoger, cfg.L1EthRpc)
	if err != nil {
		return nil, fmt.Errorf("failed to dial L1 RPC: %w", err)
	}
	multicaller := batching.NewMultiCaller(l1Client.Client(), batching.DefaultBatchSize)

	rollupProvider, err := dial.NewStaticL2RollupProvider(ctx, opLoger, cfg.RollupRpc)
	if err != nil {
		return nil, fmt.Errorf("failed to build L2 endpoint provider: %w", err)
	}

	dgfCaller := contracts.NewDisputeGameFactory(cfg.DisputeGameFactoryAddr, multicaller, cfg.NetworkTimeout)

	l2ooContract, err := bindings.NewL2OutputOracleCaller(cfg.L2OutputOracleAddr, l1Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create L2OO at address %s: %w", cfg.L2OutputOracleAddr, err)
	}

	version, err := l2ooContract.Version(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return nil, err
	}
	opLoger.Info("Connected to L2OutputOracle", "address", cfg.L2OutputOracleAddr, "version", version)

	parsed, err := bindings.L2OutputOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return &OpStackProvider{
		opStackCfg:     cfg,
		Log:            opLoger,
		RollupProvider: rollupProvider,
		l2ooContract:   l2ooContract,
		l2ooABI:        parsed,
		dgfContract:    dgfCaller,
	}, nil
}

func (op *OpStackProvider) FetchL2OOOutput(ctx context.Context, fromAddress common.Address) (*eth.OutputResponse, bool, error) {
	if op.l2ooContract == nil {
		return nil, false, fmt.Errorf("L2OutputOracle contract not set, cannot fetch next output info")
	}

	cCtx, cancel := context.WithTimeout(ctx, op.opStackCfg.NetworkTimeout)
	defer cancel()
	callOpts := &bind.CallOpts{
		From:    fromAddress,
		Context: cCtx,
	}
	nextCheckpointBlockBig, err := op.l2ooContract.NextBlockNumber(callOpts)
	if err != nil {
		return nil, false, fmt.Errorf("querying next block number: %w", err)
	}
	nextCheckpointBlock := nextCheckpointBlockBig.Uint64()
	currentBlockNumber, err := op.FetchCurrentBlockNumber(ctx)
	if err != nil {
		return nil, false, err
	}

	if currentBlockNumber < nextCheckpointBlock {
		op.Log.Debug("Proposer submission interval has not elapsed", "currentBlockNumber", currentBlockNumber, "nextBlockNumber", nextCheckpointBlock)
		return nil, false, nil
	}

	output, err := op.FetchOutput(ctx, nextCheckpointBlock)
	if err != nil {
		return nil, false, fmt.Errorf("fetching output: %w", err)
	}

	if output.BlockRef.Number > output.Status.FinalizedL2.Number && (!op.opStackCfg.AllowNonFinalized || output.BlockRef.Number > output.Status.SafeL2.Number) {
		op.Log.Debug("Not proposing yet, L2 block is not ready for proposal",
			"l2_proposal", output.BlockRef,
			"l2_safe", output.Status.SafeL2,
			"l2_finalized", output.Status.FinalizedL2,
			"allow_non_finalized", op.opStackCfg.AllowNonFinalized)
		return output, false, nil
	}
	return output, true, nil
}

func (op *OpStackProvider) FetchOutput(ctx context.Context, block uint64) (*eth.OutputResponse, error) {
	rollupClient, err := op.RollupProvider.RollupClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting rollup client: %w", err)
	}

	output, err := rollupClient.OutputAtBlock(ctx, block)
	if err != nil {
		return nil, fmt.Errorf("fetching output at block %d: %w", block, err)
	}
	if output.Version != supportedL2OutputVersion {
		return nil, fmt.Errorf("unsupported l2 output version: %v, supported: %v", output.Version, supportedL2OutputVersion)
	}
	if onum := output.BlockRef.Number; onum != block {
		return nil, fmt.Errorf("output block number %d mismatches requested %d", output.BlockRef.Number, block)
	}
	return output, nil
}

func (op *OpStackProvider) FetchCurrentBlockNumber(ctx context.Context) (uint64, error) {
	rollupClient, err := op.RollupProvider.RollupClient(ctx)
	if err != nil {
		return 0, fmt.Errorf("getting rollup client: %w", err)
	}
	status, err := rollupClient.SyncStatus(ctx)
	if err != nil {
		return 0, fmt.Errorf("getting sync status: %w", err)
	}
	if op.opStackCfg.AllowNonFinalized {
		return status.SafeL2.Number, nil
	}
	return status.FinalizedL2.Number, nil
}

func (op *OpStackProvider) FetchDGFOutput(ctx context.Context) (*eth.OutputResponse, bool, error) {
	currentBlockNumber, err := op.FetchCurrentBlockNumber(ctx)
	if err != nil {
		return nil, false, fmt.Errorf("could not fetch current block number: %w", err)
	}
	if currentBlockNumber == 0 {
		op.Log.Info("Skipping proposal for genesis block")
		return nil, false, nil
	}
	output, err := op.FetchOutput(ctx, currentBlockNumber)
	if err != nil {
		return nil, false, fmt.Errorf("could not fetch output at current block number %d: %w", currentBlockNumber, err)
	}
	return output, true, nil
}
