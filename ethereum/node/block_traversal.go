package node

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/dapplink-labs/bbn-fp-l2/ethereum/bigint"
)

var (
	ErrBlockTraversalAheadOfProvider            = errors.New("the BlockTraversal's internal state is ahead of the provider")
	ErrBlockTraversalAndProviderMismatchedState = errors.New("the BlockTraversal and provider have diverged in state")
	ErrBlockTraversalCheckBlockFail             = errors.New("the BlockTraversal check block height fail")
)

type BlockTraversal struct {
	ethClient EthClient
	chainId   uint

	latestBlock        *big.Int
	lastTraversedBlock *big.Int

	blockConfirmationDepth *big.Int
}

func NewBlockTraversal(ethClient EthClient, fromBlock *big.Int, confDepth *big.Int, chainId uint) *BlockTraversal {
	return &BlockTraversal{
		ethClient:              ethClient,
		lastTraversedBlock:     fromBlock,
		blockConfirmationDepth: confDepth,
		chainId:                chainId,
	}
}

func (f *BlockTraversal) LatestBlock() *big.Int {
	return f.latestBlock
}

func (f *BlockTraversal) LastTraversedHeader() *big.Int {
	return f.lastTraversedBlock
}

func (f *BlockTraversal) NextHeaders(maxSize uint64) ([]types.Header, error) {
	latestHeader, err := f.ethClient.BlockHeaderByNumber(nil)
	if err != nil {
		return nil, fmt.Errorf("unable to query latest block: %w", err)
	} else if latestHeader == nil {
		return nil, fmt.Errorf("latest header unreported")
	} else {
		f.latestBlock = latestHeader.Number
	}

	log.Info("header traversal db latest header: ", "latestBlock", latestHeader.Number)

	endHeight := new(big.Int).Sub(latestHeader.Number, f.blockConfirmationDepth)
	if endHeight.Sign() < 0 {
		return nil, nil
	}

	log.Info("header traversal last traversed deader to json: ", "lastTraversedBlock", f.lastTraversedBlock)

	if f.lastTraversedBlock != nil {
		cmp := f.lastTraversedBlock.Cmp(endHeight)
		if cmp == 0 {
			return nil, nil
		} else if cmp > 0 {
			return nil, ErrBlockTraversalAheadOfProvider
		}
	}

	nextHeight := bigint.Zero
	if f.lastTraversedBlock != nil {
		nextHeight = new(big.Int).Add(f.lastTraversedBlock, bigint.One)
	}

	endHeight = bigint.Clamp(nextHeight, endHeight, maxSize)
	headers, err := f.ethClient.BlockHeadersByRange(nextHeight, endHeight, f.chainId)
	if err != nil {
		return nil, fmt.Errorf("error querying blocks by range: %w", err)
	}
	if len(headers) == 0 {
		return nil, nil
	}
	err = f.checkHeaderListByHash(f.lastTraversedBlock, headers)
	if err != nil {
		log.Error("next headers check blockList by hash", "error", err)
		return nil, err
	}
	numHeaders := len(headers)
	if numHeaders == 0 {
		return nil, nil
	} else if f.lastTraversedBlock != nil && headers[0].Number != new(big.Int).Add(f.lastTraversedBlock, big.NewInt(1)) {
		log.Error("Err header traversal and provider mismatched state", "parentHash = ", headers[0].ParentHash.String(), "hash")
		return nil, ErrBlockTraversalAndProviderMismatchedState
	}
	f.lastTraversedBlock = headers[numHeaders-1].Number
	return headers, nil
}

func (f *BlockTraversal) checkHeaderListByHash(dbLatestBlock *big.Int, headerList []types.Header) error {
	if len(headerList) == 0 {
		return nil
	}
	if len(headerList) == 1 {
		return nil
	}

	if dbLatestBlock != nil && headerList[0].Number != new(big.Int).Add(dbLatestBlock, big.NewInt(1)) {
		log.Error("check header list by hash", "parentHash = ", headerList[0].ParentHash.String())
		return ErrBlockTraversalCheckBlockFail
	}

	for i := 1; i < len(headerList); i++ {
		if headerList[i].ParentHash != headerList[i-1].Hash() {
			return fmt.Errorf("check header list by hash: block parent hash not equal parent block hash")
		}
	}
	return nil
}

func (f *BlockTraversal) ChangeLastTraversedHeaderByDelAfter(dbLatestBlock *big.Int) {
	f.lastTraversedBlock = dbLatestBlock
}
