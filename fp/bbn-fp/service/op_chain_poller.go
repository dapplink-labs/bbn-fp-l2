package service

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"
	"time"

	"go.uber.org/atomic"
	"go.uber.org/zap"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ctypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/Manta-Network/manta-fp/ethereum/node"
	cfg "github.com/Manta-Network/manta-fp/fp/bbn-fp/config"
	"github.com/Manta-Network/manta-fp/fp/bbn-fp/store"
	"github.com/Manta-Network/manta-fp/l2chain/opstack"
	"github.com/Manta-Network/manta-fp/metrics"
	"github.com/Manta-Network/manta-fp/types"
)

type OpChainPoller struct {
	isStarted *atomic.Bool
	wg        sync.WaitGroup
	logger    *zap.Logger
	sRStore   *store.OpStateRootStore

	opClient       node.EthClient
	cfg            *cfg.OpEventConfig
	headers        []ctypes.Header
	latestBlock    *big.Int
	blockTraversal *node.BlockTraversal
	eventProvider  *opstack.EventProvider
	blockInfoChan  chan *types.BlockInfo

	metrics *metrics.FpMetrics
	quit    chan struct{}
}

func NewOpChainPoller(logger *zap.Logger, opClient node.EthClient, startHeight uint64, cfg *cfg.OpEventConfig, sRStore *store.OpStateRootStore, eventProvider *opstack.EventProvider, metrics *metrics.FpMetrics) (*OpChainPoller, error) {
	dbLatestBlock, err := sRStore.GetLatestBlock()
	if err != nil {
		return nil, err
	}
	var fromBlock *big.Int
	if dbLatestBlock != nil {
		logger.Info("sync detected last indexed block", zap.String("blockNumber", dbLatestBlock.String()))
		fromBlock = dbLatestBlock
	} else if startHeight > 0 {
		logger.Info("no sync indexed state starting from supplied ethereum height", zap.Uint64("height", startHeight))
		header, err := opClient.BlockHeaderByNumber(big.NewInt(int64(startHeight)))
		if err != nil {
			return nil, fmt.Errorf("could not fetch starting block header: %w", err)
		}
		fromBlock = header.Number
	} else {
		logger.Info("no ethereum block indexed state")
	}

	blockTraversal := node.NewBlockTraversal(opClient, fromBlock, big.NewInt(0), cfg.ChainId, logger)

	return &OpChainPoller{
		isStarted:      atomic.NewBool(false),
		logger:         logger,
		opClient:       opClient,
		sRStore:        sRStore,
		cfg:            cfg,
		latestBlock:    fromBlock,
		blockTraversal: blockTraversal,
		eventProvider:  eventProvider,
		blockInfoChan:  make(chan *types.BlockInfo, cfg.BufferSize),
		metrics:        metrics,
		quit:           make(chan struct{}),
	}, nil
}

func (ocp *OpChainPoller) Start(startHeight uint64) error {
	if ocp.isStarted.Swap(true) {
		return fmt.Errorf("the op chain poller is already started")
	}

	ocp.logger.Info("starting the op chain poller")

	ocp.wg.Add(1)
	go ocp.opPollChain()

	ocp.metrics.RecordPollerStartingHeight(startHeight) // todo: change to op stack
	ocp.logger.Info("the chain poller is successfully started")

	return nil
}

func (ocp *OpChainPoller) Stop() error {
	if !ocp.isStarted.Swap(false) {
		return fmt.Errorf("the op chain poller has already stopped")
	}

	ocp.logger.Info("stopping the op chain poller")

	close(ocp.quit)
	ocp.wg.Wait()

	ocp.logger.Info("the op chain poller is successfully stopped")

	return nil
}

func (ocp *OpChainPoller) GetBlockInfoChan() <-chan *types.BlockInfo {
	return ocp.blockInfoChan
}

func (ocp *OpChainPoller) opPollChain() {
	defer ocp.wg.Done()
	for {
		select {
		case <-time.After(ocp.cfg.PollInterval):
			if len(ocp.headers) > 0 {
				ocp.logger.Info("retrying previous batch")
			} else {
				newHeaders, err := ocp.blockTraversal.NextHeaders(uint64(ocp.cfg.BlockStep))
				if err != nil {
					ocp.logger.Error("error querying for headers", zap.String("err", err.Error()))
					continue
				} else if len(newHeaders) == 0 {
					ocp.logger.Warn("no new headers. syncer at head?")
				} else {
					ocp.headers = newHeaders
				}
				latestBlock := ocp.blockTraversal.LatestBlock()
				if latestBlock != nil {
					ocp.logger.Info("Latest header", zap.String("latestHeader Number", latestBlock.String()))
				}
				err = ocp.sRStore.AddLatestBlock(latestBlock)
				if err != nil {
					ocp.logger.Error("Add latest block fail", zap.String("err", err.Error()))
					return
				}
			}
			err := ocp.processBatch(ocp.headers, ocp.cfg)
			if err == nil {
				ocp.headers = nil
			}
		case <-ocp.quit:
			return
		}
	}
}

func (ocp *OpChainPoller) processBatch(headers []ctypes.Header, chainCfg *cfg.OpEventConfig) error {
	if len(headers) == 0 {
		return nil
	}
	firstHeader, lastHeader := headers[0], headers[len(headers)-1]
	ocp.logger.Info("extracting batch", zap.Int("size", len(headers)), zap.String("startBlock", firstHeader.Number.String()), zap.String("endBlock", lastHeader.Number.String()))
	headerMap := make(map[common.Hash]*ctypes.Header, len(headers))
	for i := range headers {
		header := headers[i]
		headerMap[header.Hash()] = &header
	}

	filterQuery := ethereum.FilterQuery{FromBlock: firstHeader.Number, ToBlock: lastHeader.Number, Addresses: ocp.cfg.Contracts}
	logs, err := ocp.opClient.FilterLogs(filterQuery)
	if err != nil {
		ocp.logger.Error("failed to extract logs", zap.String("err", err.Error()))
		return err
	}

	if logs.ToBlockHeader.Number.Cmp(lastHeader.Number) != 0 {
		return fmt.Errorf("mismatch in FilterLog#ToBlock number")
	} else if logs.ToBlockHeader.Hash() != lastHeader.Hash() {
		return fmt.Errorf("mismatch in FitlerLog#ToBlock block hash")
	}

	if len(logs.Logs) > 0 {
		ocp.logger.Info("detected logs", zap.Int("size", len(logs.Logs)))
	}

	blockList := make([]types.Block, 0, len(headers))
	for i := range headers {
		if headers[i].Number == nil {
			continue
		}
		blockItem := types.Block{
			Hash:       headers[i].Hash(),
			ParentHash: headers[i].ParentHash,
			Number:     headers[i].Number,
			Timestamp:  headers[i].Time,
		}
		blockList = append(blockList, blockItem)
	}

	// contracts parse
	for i := range logs.Logs {
		stateRootEvent, err := ocp.eventProvider.ProcessStateRootEvent(logs.Logs[i])
		if err != nil {
			return err
		}
		ocp.logger.Info("event list", zap.String("stateroot", hex.EncodeToString(stateRootEvent.StateRoot[:])))

		err = ocp.sRStore.SaveStateRoot(big.NewInt(int64(stateRootEvent.L1BlockNumber)), stateRootEvent.StateRoot,
			stateRootEvent.L2BlockNumber, stateRootEvent.L1BlockHash, stateRootEvent.L2OutputIndex, stateRootEvent.DisputeGameType)
		if err != nil {
			ocp.logger.Error("failed to store state root", zap.String("err", err.Error()))
			return err
		}
		ocp.blockInfoChan <- &types.BlockInfo{
			Height:    logs.Logs[i].BlockNumber,
			Hash:      logs.Logs[i].BlockHash.Bytes(),
			Finalized: false,
			StateRoot: *stateRootEvent,
		}
	}
	return nil
}
