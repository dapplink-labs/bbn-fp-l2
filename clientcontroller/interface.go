package clientcontroller

import (
	"fmt"

	"cosmossdk.io/math"
	btcstakingtypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/chaincfg"
	"go.uber.org/zap"

	fpcfg "github.com/Manta-Network/manta-fp/fp/bbn-fp/config"
	"github.com/Manta-Network/manta-fp/types"
)

const (
	babylonConsumerChainType = "babylon"
)

type ClientController interface {
	// RegisterFinalityProvider registers a finality provider to the consumer chain
	// it returns tx hash and error. The address of the finality provider will be
	// the signer of the msg.
	RegisterFinalityProvider(
		fpPk *btcec.PublicKey,
		pop []byte,
		commission *math.LegacyDec,
		description []byte,
	) (*types.TxResponse, error)

	// EditFinalityProvider edits description and commission of a finality provider
	EditFinalityProvider(fpPk *btcec.PublicKey, commission *math.LegacyDec, description []byte) (*btcstakingtypes.MsgEditFinalityProvider, error)

	// CommitPubRandList commits a list of EOTS public randomness the consumer chain
	// it returns tx hash and error
	CommitPubRandList(fpPk *btcec.PublicKey, startHeight uint64, numPubRand uint64, commitment []byte, sig *schnorr.Signature) (*types.TxResponse, error)

	// SubmitFinalitySig submits the finality signature to the consumer chain
	SubmitFinalitySig(fpPk *btcec.PublicKey, block *types.BlockInfo, pubRand *btcec.FieldVal, proof []byte, sig *btcec.ModNScalar) (*types.TxResponse, error)

	// SubmitBatchFinalitySigs submits a batch of finality signatures to the consumer chain
	SubmitBatchFinalitySigs(fpPk *btcec.PublicKey, blocks *types.BlockInfo, pubRandList *btcec.FieldVal, proofList []byte, sigs *btcec.ModNScalar) (*types.TxResponse, error)

	// UnjailFinalityProvider sends an unjail transaction to the consumer chain
	UnjailFinalityProvider(fpPk *btcec.PublicKey) (*types.TxResponse, error)

	/*
		The following methods are queries to the consumer chain
	*/

	// QueryFinalityProviderVotingPower queries the voting power of the finality provider at a given height
	QueryFinalityProviderVotingPower(fpPk *btcec.PublicKey, blockHeight uint64) (bool, error)

	// QueryFinalityProviderSlashedOrJailed queries if the finality provider is slashed or jailed
	QueryFinalityProviderSlashedOrJailed(fpPk *btcec.PublicKey) (slashed bool, jailed bool, err error)

	// QueryLatestFinalizedBlocks returns the latest finalized blocks
	QueryLatestFinalizedBlocks() (uint64, error)

	// QueryLastCommittedPublicRand returns the last committed public randomness
	QueryLastCommittedPublicRand(fpPk *btcec.PublicKey) (*types.PubRandCommit, error)

	// QueryBlock queries the block at the given height
	QueryBlock(height uint64) (bool, error)

	// QueryBlocks returns a list of blocks from startHeight to endHeight
	QueryBlocks(startHeight, endHeight uint64, limit uint32) ([]*types.BlockInfo, error)

	// QueryBestBlock queries the tip block of the consumer chain
	QueryBestBlock() (*types.BlockInfo, error)

	// QueryActivatedHeight returns the activated height of the consumer chain
	// error will be returned if the consumer chain has not been activated
	QueryActivatedHeight() (uint64, error)

	Close() error
}

func NewClientController(chainType string, bbnConfig *fpcfg.BBNConfig, opConfig *fpcfg.OpEventConfig, netParams *chaincfg.Params, logger *zap.Logger) (ClientController, error) {
	var (
		cc  ClientController
		err error
	)

	switch chainType {
	case babylonConsumerChainType:
		cc, err = NewBabylonController(bbnConfig, opConfig, netParams, logger)
		if err != nil {
			return nil, fmt.Errorf("failed to create Babylon rpc client: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported consumer chain")
	}

	return cc, err
}
