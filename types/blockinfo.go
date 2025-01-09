package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Block struct {
	Hash       common.Hash `json:"hash"`
	ParentHash common.Hash `json:"parent_hash"`
	Number     *big.Int    `json:"number"`
	Timestamp  uint64      `json:"timestamp"`
}

type BlockInfo struct {
	Height    uint64
	Hash      []byte
	Finalized bool
	StateRoot
}

type StateRoot struct {
	StateRoot       [32]byte    `json:"state_root"`
	L2BlockNumber   *big.Int    `json:"l2_block_number"`
	L2OutputIndex   *big.Int    `json:"l2_output_index"`
	L1BlockHash     common.Hash `json:"l1_block_hash"`
	L1BlockNumber   uint64      `json:"l1_block_number"`
	DisputeGameType uint64      `json:"dispute_game_type"`
}
