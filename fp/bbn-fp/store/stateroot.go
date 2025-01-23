package store

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/btcsuite/btcwallet/walletdb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lightningnetwork/lnd/kvdb"

	"github.com/Manta-Network/manta-fp/types"
)

var (
	ErrCorruptedLatestBlockrDb = errors.New("latest block manager db is corrupted")

	ErrCorruptedBlockHeaderDb = errors.New("block header manager db is corrupted")

	ErrDuplicateBlock = errors.New("block header key name already exists")

	ErrBlockNotFound = errors.New("block header key name not found")

	ErrStateRootNotFound = errors.New("state root record not found")

	ErrDuplicateStateRoot = errors.New("state root for given height already exists")
)

var (
	LatestBlock         = []byte("latestBlock")
	LatestBlockKey      = []byte("latestBlockKey")
	BlockHeaderName     = []byte("blockHeader")
	StateRootBucketName = []byte("opStateRoot19")
)

type OpStateRootStore struct {
	db kvdb.Backend
}

func NewOpStateRootStore(db kvdb.Backend) (*OpStateRootStore, error) {
	store := &OpStateRootStore{db}
	if err := store.initBuckets(); err != nil {
		return nil, err
	}
	return store, nil
}

func (s *OpStateRootStore) initBuckets() error {
	return kvdb.Batch(s.db, func(tx kvdb.RwTx) error {
		_, err := tx.CreateTopLevelBucket(LatestBlock)
		if err != nil {
			return err
		}

		_, err = tx.CreateTopLevelBucket(BlockHeaderName)
		if err != nil {
			return err
		}

		_, err = tx.CreateTopLevelBucket(StateRootBucketName)
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *OpStateRootStore) AddLatestBlock(blockNumber *big.Int) error {
	return kvdb.Batch(s.db, func(tx kvdb.RwTx) error {
		latestBlockBucket := tx.ReadWriteBucket(LatestBlock)
		if latestBlockBucket == nil {
			return ErrCorruptedLatestBlockrDb
		}

		return latestBlockBucket.Put(LatestBlockKey, blockNumber.Bytes())
	})
}

func (s *OpStateRootStore) GetLatestBlock() (*big.Int, error) {
	var blockNumber *big.Int
	err := s.db.View(func(tx kvdb.RTx) error {
		blockNumberBucket := tx.ReadBucket(LatestBlock)
		if blockNumberBucket == nil {
			return ErrCorruptedLatestBlockrDb
		}

		blockByte := blockNumberBucket.Get(LatestBlockKey)
		if blockByte == nil {
			return nil
		}
		blockNumber = new(big.Int).SetBytes(blockByte)
		return nil
	}, func() {})

	if err != nil {
		return nil, err
	}
	return blockNumber, nil
}

func (s *OpStateRootStore) AddBlock(
	blockNumber *big.Int,
	Hash common.Hash,
	ParentHash common.Hash,
	Timestamp uint64,
) error {
	return kvdb.Batch(s.db, func(tx kvdb.RwTx) error {
		blockHeaderBucket := tx.ReadWriteBucket(BlockHeaderName)
		if blockHeaderBucket == nil {
			return ErrCorruptedBlockHeaderDb
		}

		if blockHeaderBucket.Get(blockNumber.Bytes()) != nil {
			return ErrDuplicateBlock
		}

		return saveBlock(blockHeaderBucket, blockNumber, Hash, ParentHash, Timestamp)
	})
}

func saveBlock(
	blockHeaderBucket walletdb.ReadWriteBucket,
	blockNumber *big.Int,
	Hash common.Hash,
	ParentHash common.Hash,
	Timestamp uint64,
) error {
	block := &types.Block{
		Hash:       Hash,
		ParentHash: ParentHash,
		Number:     blockNumber,
		Timestamp:  Timestamp,
	}

	blockMarshalled, err := json.Marshal(block)
	if err != nil {
		return err
	}

	return blockHeaderBucket.Put(blockNumber.Bytes(), blockMarshalled)
}

func (s *OpStateRootStore) GetBlock(blockNumber *big.Int) (*types.Block, error) {
	var blockInfo *types.Block
	err := s.db.View(func(tx kvdb.RTx) error {
		blockNumberBucket := tx.ReadBucket(BlockHeaderName)
		if blockNumberBucket == nil {
			return ErrCorruptedBlockHeaderDb
		}

		blockByte := blockNumberBucket.Get(blockNumber.Bytes())
		if blockByte == nil {
			return ErrBlockNotFound
		}

		err := json.Unmarshal(blockByte, blockInfo)
		if err != nil {
			return err
		}
		return nil
	}, func() {})

	if err != nil {
		return nil, err
	}
	return blockInfo, nil
}

func (s *OpStateRootStore) SaveStateRoot(
	l1BlockNumber *big.Int,
	stateRoot [32]byte,
	l2BlockNumber *big.Int,
	l1BlockHash common.Hash,
	L2OutputIndex *big.Int,
	disputeGameType uint64,
) error {
	return kvdb.Batch(s.db, func(tx kvdb.RwTx) error {
		bucket := tx.ReadWriteBucket(StateRootBucketName)
		if bucket == nil {
			return ErrCorruptedBlockHeaderDb
		}

		if bucket.Get(l1BlockNumber.Bytes()) != nil {
			return ErrDuplicateStateRoot
		}

		sttRoot := &types.StateRoot{
			StateRoot:       stateRoot,
			L2BlockNumber:   l2BlockNumber,
			L1BlockHash:     l1BlockHash,
			L2OutputIndex:   L2OutputIndex,
			DisputeGameType: disputeGameType,
		}

		stateRootMarshalled, err := json.Marshal(sttRoot)
		if err != nil {
			return err
		}

		return bucket.Put(l1BlockNumber.Bytes(), stateRootMarshalled)
	})
}

func (s *OpStateRootStore) GetStateRoot(l1BlockNumber *big.Int) (*types.StateRoot, error) {
	stateRootRes := &types.StateRoot{}
	err := s.db.View(func(tx kvdb.RTx) error {
		bucket := tx.ReadBucket(StateRootBucketName)
		if bucket == nil {
			return ErrCorruptedBlockHeaderDb
		}

		StateRootBytes := bucket.Get(l1BlockNumber.Bytes())
		if StateRootBytes == nil {
			return ErrStateRootNotFound
		}
		err := json.Unmarshal(StateRootBytes, stateRootRes)
		if err != nil {
			return err
		}
		return nil
	}, func() {})

	if err != nil {
		if errors.Is(err, ErrStateRootNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return stateRootRes, nil
}
