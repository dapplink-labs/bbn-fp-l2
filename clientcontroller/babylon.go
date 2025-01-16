package clientcontroller

import (
	"context"
	"encoding/json"
	"fmt"
	wasmdparams "github.com/CosmWasm/wasmd/app/params"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"math/big"
	"strings"
	"time"

	"github.com/Manta-Network/manta-fp/finality-provider/proto"

	sdkErr "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	cwclient "github.com/Manta-Network/manta-fp/cosmwasmclient/client"
	cwconfig "github.com/Manta-Network/manta-fp/cosmwasmclient/config"
	bbnapp "github.com/babylonlabs-io/babylon/app"
	bbnclient "github.com/babylonlabs-io/babylon/client/client"
	bbntypes "github.com/babylonlabs-io/babylon/types"
	btcctypes "github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
	btclctypes "github.com/babylonlabs-io/babylon/x/btclightclient/types"
	btcstakingtypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	finalitytypes "github.com/babylonlabs-io/babylon/x/finality/types"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	cmtcrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	sttypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/relayer/v2/relayer/provider"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"go.uber.org/zap"
	protobuf "google.golang.org/protobuf/proto"

	opclient "github.com/Manta-Network/manta-fp/ethereum/node"
	fpcfg "github.com/Manta-Network/manta-fp/finality-provider/config"
	"github.com/Manta-Network/manta-fp/types"
)

const (
	BabylonChainName = "Babylon"
)

var _ ClientController = &BabylonController{}

var emptyErrs = []*sdkErr.Error{}

type BabylonController struct {
	bbnClient  *bbnclient.Client
	opl2Client opclient.EthClient
	CwClient   *cwclient.Client
	cfg        *fpcfg.BBNConfig
	btcParams  *chaincfg.Params
	logger     *zap.Logger

	OPFinalityGadgetAddress  string
	BabylonFinalityGadgetRpc string
}

func NewBabylonController(
	cfg *fpcfg.BBNConfig,
	opCfg *fpcfg.OpEventConfig,
	btcParams *chaincfg.Params,
	logger *zap.Logger,
) (*BabylonController, error) {
	bbnConfig := fpcfg.BBNConfigToBabylonConfig(cfg)

	bc, err := bbnclient.New(
		&bbnConfig,
		logger,
	)

	cwConfig := fpcfg.BBNConfigToCosmwasmConfig(cfg)

	cwClient, err := NewCwClient(&cwConfig, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create CW client: %w", err)
	}

	opClient, err := opclient.DialEthClient(context.Background(), opCfg.EthRpc)
	if err != nil {
		return nil, fmt.Errorf("failed to create op client: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create Babylon client: %w", err)
	}

	// makes sure that the key in config really exists and is a valid bech32 addr
	// to allow using mustGetTxSigner
	if _, err := bc.GetAddr(); err != nil {
		return nil, err
	}

	return &BabylonController{
		bc,
		opClient,
		cwClient,
		cfg,
		btcParams,
		logger,
		opCfg.OPFinalityGadgetAddress,
		opCfg.BabylonFinalityGadgetRpc,
	}, nil
}

func (bc *BabylonController) mustGetTxSigner() string {
	signer := bc.GetKeyAddress()
	prefix := bc.cfg.AccountPrefix
	return sdk.MustBech32ifyAddressBytes(prefix, signer)
}

func (bc *BabylonController) GetKeyAddress() sdk.AccAddress {
	// get key address, retrieves address based on the key name which is configured in
	// cfg *stakercfg.BBNConfig. If this fails, it means we have a misconfiguration problem
	// and we should panic.
	// This is checked at the start of BabylonController, so if it fails something is really wrong

	keyRec, err := bc.bbnClient.GetKeyring().Key(bc.cfg.Key)
	if err != nil {
		panic(fmt.Sprintf("Failed to get key address: %s", err))
	}

	addr, err := keyRec.GetAddress()
	if err != nil {
		panic(fmt.Sprintf("Failed to get key address: %s", err))
	}

	return addr
}

func (bc *BabylonController) reliablySendMsg(msg sdk.Msg, expectedErrs []*sdkErr.Error, unrecoverableErrs []*sdkErr.Error) (*provider.RelayerTxResponse, error) {
	return bc.reliablySendMsgs([]sdk.Msg{msg}, expectedErrs, unrecoverableErrs)
}

func (bc *BabylonController) reliablySendMsgs(msgs []sdk.Msg, expectedErrs []*sdkErr.Error, unrecoverableErrs []*sdkErr.Error) (*provider.RelayerTxResponse, error) {
	return bc.CwClient.ReliablySendMsgs(
		context.Background(),
		msgs,
		expectedErrs,
		unrecoverableErrs,
	)
}

// RegisterFinalityProvider registers a finality provider via a MsgCreateFinalityProvider to Babylon
// it returns tx hash and error
func (bc *BabylonController) RegisterFinalityProvider(
	fpPk *btcec.PublicKey,
	pop []byte,
	commission *sdkmath.LegacyDec,
	description []byte,
) (*types.TxResponse, error) {
	var bbnPop btcstakingtypes.ProofOfPossessionBTC
	if err := bbnPop.Unmarshal(pop); err != nil {
		return nil, fmt.Errorf("invalid proof-of-possession: %w", err)
	}

	var sdkDescription sttypes.Description
	if err := sdkDescription.Unmarshal(description); err != nil {
		return nil, fmt.Errorf("invalid description: %w", err)
	}

	fpAddr := bc.mustGetTxSigner()
	msg := &btcstakingtypes.MsgCreateFinalityProvider{
		Addr:        fpAddr,
		BtcPk:       bbntypes.NewBIP340PubKeyFromBTCPK(fpPk),
		Pop:         &bbnPop,
		Commission:  commission,
		Description: &sdkDescription,
	}

	res, err := bc.reliablySendMsg(msg, emptyErrs, emptyErrs)
	if err != nil {
		return nil, err
	}

	return &types.TxResponse{TxHash: res.TxHash, Events: res.Events}, nil
}

// CommitPubRandList commits a list of Schnorr public randomness via a MsgCommitPubRand to Babylon
// it returns tx hash and error
func (bc *BabylonController) CommitPubRandList(
	fpPk *btcec.PublicKey,
	startHeight uint64,
	numPubRand uint64,
	commitment []byte,
	sig *schnorr.Signature,
) (*types.TxResponse, error) {
	msg := CommitPublicRandomnessMsg{
		CommitPublicRandomness: CommitPublicRandomnessMsgParams{
			FpPubkeyHex: bbntypes.NewBIP340PubKeyFromBTCPK(fpPk).MarshalHex(),
			StartHeight: startHeight,
			NumPubRand:  numPubRand,
			Commitment:  commitment,
			Signature:   sig.Serialize(),
		},
	}
	payload, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	execMsg := &wasmtypes.MsgExecuteContract{
		Sender:   bc.CwClient.MustGetAddr(),
		Contract: bc.OPFinalityGadgetAddress,
		Msg:      payload,
	}

	res, err := bc.reliablySendMsg(execMsg, emptyErrs, nil)
	if err != nil {
		return nil, err
	}

	return &types.TxResponse{TxHash: res.TxHash, Events: res.Events}, nil
}

// SubmitFinalitySig submits the finality signature via a MsgAddVote to Babylon
func (bc *BabylonController) SubmitFinalitySig(
	fpPk *btcec.PublicKey,
	block *types.BlockInfo,
	pubRand *btcec.FieldVal,
	proof []byte, // TODO: have a type for proof
	sig *btcec.ModNScalar,
) (*types.TxResponse, error) {
	return bc.SubmitBatchFinalitySigs(
		fpPk, block, pubRand,
		proof, sig,
	)
}

// SubmitBatchFinalitySigs submits a batch of finality signatures to Babylon
func (bc *BabylonController) SubmitBatchFinalitySigs(
	fpPk *btcec.PublicKey,
	block *types.BlockInfo,
	pubRand *btcec.FieldVal,
	proof []byte,
	sig *btcec.ModNScalar,
) (*types.TxResponse, error) {

	unrecoverableErrs := []*sdkErr.Error{
		finalitytypes.ErrInvalidFinalitySig,
		finalitytypes.ErrPubRandNotFound,
		btcstakingtypes.ErrFpAlreadySlashed,
	}

	cmtProof := cmtcrypto.Proof{}
	if err := cmtProof.Unmarshal(proof); err != nil {
		return nil, err
	}

	msg := SubmitFinalitySignatureMsg{
		SubmitFinalitySignature: SubmitFinalitySignatureMsgParams{
			FpPubkeyHex: bbntypes.NewBIP340PubKeyFromBTCPK(fpPk).MarshalHex(),
			Height:      block.Height,
			PubRand:     bbntypes.NewSchnorrPubRandFromFieldVal(pubRand).MustMarshal(),
			Proof:       ConvertProof(cmtProof),
			BlockHash:   block.Hash,
			Signature:   bbntypes.NewSchnorrEOTSSigFromModNScalar(sig).MustMarshal(),
		},
	}
	if msg.SubmitFinalitySignature.Proof.Aunts == nil {
		msg.SubmitFinalitySignature.Proof.Aunts = [][]byte{}
	}
	payload, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	execMsg := &wasmtypes.MsgExecuteContract{
		Sender:   bc.CwClient.MustGetAddr(),
		Contract: bc.OPFinalityGadgetAddress,
		Msg:      payload,
	}

	res, err := bc.reliablySendMsg(execMsg, emptyErrs, unrecoverableErrs)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return &types.TxResponse{}, nil
	}

	return &types.TxResponse{TxHash: res.TxHash, Events: res.Events}, nil
}

func ConvertProof(cmtProof cmtcrypto.Proof) Proof {
	return Proof{
		Total:    uint64(cmtProof.Total),
		Index:    uint64(cmtProof.Index),
		LeafHash: cmtProof.LeafHash,
		Aunts:    cmtProof.Aunts,
	}
}

// UnjailFinalityProvider sends an unjail transaction to the consumer chain
func (bc *BabylonController) UnjailFinalityProvider(fpPk *btcec.PublicKey) (*types.TxResponse, error) {
	msg := &finalitytypes.MsgUnjailFinalityProvider{
		Signer:  bc.mustGetTxSigner(),
		FpBtcPk: bbntypes.NewBIP340PubKeyFromBTCPK(fpPk),
	}

	unrecoverableErrs := []*sdkErr.Error{
		btcstakingtypes.ErrFpNotFound,
		btcstakingtypes.ErrFpNotJailed,
		btcstakingtypes.ErrFpAlreadySlashed,
	}

	res, err := bc.reliablySendMsg(msg, emptyErrs, unrecoverableErrs)
	if err != nil {
		return nil, err
	}

	return &types.TxResponse{TxHash: res.TxHash, Events: res.Events}, nil
}

// QueryFinalityProviderSlashedOrJailed - returns if the fp has been slashed, jailed, err
func (bc *BabylonController) QueryFinalityProviderSlashedOrJailed(fpPk *btcec.PublicKey) (bool, bool, error) {
	fpPubKey := bbntypes.NewBIP340PubKeyFromBTCPK(fpPk)
	res, err := bc.bbnClient.QueryClient.FinalityProvider(fpPubKey.MarshalHex())
	if err != nil {
		return false, false, fmt.Errorf("failed to query the finality provider %s: %w", fpPubKey.MarshalHex(), err)
	}

	return res.FinalityProvider.SlashedBtcHeight > 0, res.FinalityProvider.Jailed, nil
}

// QueryFinalityProviderVotingPower queries the voting power of the finality provider at a given height
func (bc *BabylonController) QueryFinalityProviderVotingPower(fpPk *btcec.PublicKey, blockHeight uint64) (bool, error) {
	fpBtcPkHex := bbntypes.NewBIP340PubKeyFromBTCPK(fpPk).MarshalHex()
	var nextKey []byte

	btcStakingParams, err := bc.bbnClient.QueryClient.BTCStakingParams()
	if err != nil {
		return false, err
	}
	for {
		resp, err := bc.bbnClient.QueryClient.FinalityProviderDelegations(fpBtcPkHex, &sdkquery.PageRequest{Key: nextKey, Limit: 100})
		if err != nil {
			return false, err
		}

		for _, btcDels := range resp.BtcDelegatorDelegations {
			for _, btcDel := range btcDels.Dels {
				active, err := bc.isDelegationActive(btcStakingParams, btcDel)
				if err != nil {
					continue
				}
				if active {
					return true, nil
				}
			}
		}

		if resp.Pagination == nil || resp.Pagination.NextKey == nil {
			break
		}
		nextKey = resp.Pagination.NextKey
	}

	return false, nil

}

func (bc *BabylonController) isDelegationActive(
	btcStakingParams *btcstakingtypes.QueryParamsResponse,
	btcDel *btcstakingtypes.BTCDelegationResponse,
) (bool, error) {

	covQuorum := btcStakingParams.GetParams().CovenantQuorum
	ud := btcDel.UndelegationResponse

	if uint32(len(btcDel.CovenantSigs)) < covQuorum {
		return false, nil
	}
	if len(ud.CovenantUnbondingSigList) < int(covQuorum) {
		return false, nil
	}
	if len(ud.CovenantSlashingSigs) < int(covQuorum) {
		return false, nil
	}

	return true, nil
}

func (bc *BabylonController) QueryLatestFinalizedBlocks() (uint64, error) {
	l2LatestBlock, err := bc.opl2Client.BlockHeaderByNumber(big.NewInt(ethrpc.LatestBlockNumber.Int64()))
	if err != nil {
		return 0, err
	}

	return l2LatestBlock.Number.Uint64(), nil
}

// QueryLastCommittedPublicRand returns the last public randomness commitments
func (bc *BabylonController) QueryLastCommittedPublicRand(fpPk *btcec.PublicKey) (*types.PubRandCommit, error) {

	fpPubKey := bbntypes.NewBIP340PubKeyFromBTCPK(fpPk)
	queryMsg := &QueryMsg{
		LastPubRandCommit: &PubRandCommit{
			BtcPkHex: fpPubKey.MarshalHex(),
		},
	}

	jsonData, err := json.Marshal(queryMsg)
	if err != nil {
		return nil, fmt.Errorf("failed marshaling to JSON: %w", err)
	}

	stateResp, err := bc.CwClient.QuerySmartContractState(bc.OPFinalityGadgetAddress, string(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to query smart contract state: %w", err)
	}
	if len(stateResp.Data) == 0 {
		return nil, nil
	}

	var resp *types.PubRandCommit
	err = json.Unmarshal(stateResp.Data, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if resp == nil {
		return nil, nil
	}
	if err := resp.Validate(); err != nil {
		return nil, err
	}

	return resp, nil
}

func (bc *BabylonController) QueryBlocks(startHeight, endHeight uint64, limit uint32) ([]*types.BlockInfo, error) {
	if endHeight < startHeight {
		return nil, fmt.Errorf("the startHeight %v should not be higher than the endHeight %v", startHeight, endHeight)
	}
	count := endHeight - startHeight + 1
	if count > uint64(limit) {
		count = uint64(limit)
	}
	return bc.queryLatestBlocks(sdk.Uint64ToBigEndian(startHeight), count, finalitytypes.QueriedBlockStatus_ANY, false)
}

func (bc *BabylonController) queryLatestBlocks(startKey []byte, count uint64, status finalitytypes.QueriedBlockStatus, reverse bool) ([]*types.BlockInfo, error) {
	var blocks []*types.BlockInfo
	pagination := &sdkquery.PageRequest{
		Limit:   count,
		Reverse: reverse,
		Key:     startKey,
	}

	res, err := bc.bbnClient.QueryClient.ListBlocks(status, pagination)
	if err != nil {
		return nil, fmt.Errorf("failed to query finalized blocks: %w", err)
	}

	for _, b := range res.Blocks {
		ib := &types.BlockInfo{
			Height: b.Height,
			Hash:   b.AppHash,
		}
		blocks = append(blocks, ib)
	}

	return blocks, nil
}

func getContextWithCancel(timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return ctx, cancel
}

func (bc *BabylonController) QueryBlock(height uint64) (bool, error) {
	//Todo change to finalized block
	l2Block, err := bc.opl2Client.BlockNumber()
	if err != nil {
		return false, err
	}

	//if l2Block == nil {
	//	return false, nil
	//}
	if height > l2Block {
		return false, nil
	}
	return true, nil

}

func (bc *BabylonController) QueryActivatedHeight() (uint64, error) {
	res, err := bc.bbnClient.QueryClient.ActivatedHeight()
	if err != nil {
		return 0, fmt.Errorf("failed to query activated height: %w", err)
	}

	return res.Height, nil
}

func (bc *BabylonController) QueryBestBlock() (*types.BlockInfo, error) {
	l2Block, err := bc.opl2Client.BlockNumber()
	if err != nil {
		return nil, err
	}

	return &types.BlockInfo{Height: l2Block}, nil
}

func (bc *BabylonController) queryCometBestBlock() (*types.BlockInfo, error) {
	ctx, cancel := getContextWithCancel(bc.cfg.Timeout)
	// this will return 20 items at max in the descending order (highest first)
	chainInfo, err := bc.bbnClient.RPCClient.BlockchainInfo(ctx, 0, 0)
	defer cancel()

	if err != nil {
		return nil, err
	}

	headerHeightInt64 := chainInfo.BlockMetas[0].Header.Height
	if headerHeightInt64 < 0 {
		return nil, fmt.Errorf("block height %v should be positive", headerHeightInt64)
	}
	// Returning response directly, if header with specified number did not exist
	// at request will contain nil header
	return &types.BlockInfo{
		Height: uint64(headerHeightInt64),
		Hash:   chainInfo.BlockMetas[0].Header.AppHash,
	}, nil
}

func (bc *BabylonController) Close() error {
	if !bc.bbnClient.IsRunning() {
		return nil
	}

	return bc.bbnClient.Stop()
}

/*
	Implementations for e2e tests only
*/

func (bc *BabylonController) CreateBTCDelegation(
	delBtcPk *bbntypes.BIP340PubKey,
	fpPks []*btcec.PublicKey,
	pop *btcstakingtypes.ProofOfPossessionBTC,
	stakingTime uint32,
	stakingValue int64,
	stakingTxInfo *btcctypes.TransactionInfo,
	slashingTx *btcstakingtypes.BTCSlashingTx,
	delSlashingSig *bbntypes.BIP340Signature,
	unbondingTx []byte,
	unbondingTime uint32,
	unbondingValue int64,
	unbondingSlashingTx *btcstakingtypes.BTCSlashingTx,
	delUnbondingSlashingSig *bbntypes.BIP340Signature,
) (*types.TxResponse, error) {
	fpBtcPks := make([]bbntypes.BIP340PubKey, 0, len(fpPks))
	for _, v := range fpPks {
		fpBtcPks = append(fpBtcPks, *bbntypes.NewBIP340PubKeyFromBTCPK(v))
	}
	msg := &btcstakingtypes.MsgCreateBTCDelegation{
		StakerAddr:                    bc.mustGetTxSigner(),
		Pop:                           pop,
		BtcPk:                         delBtcPk,
		FpBtcPkList:                   fpBtcPks,
		StakingTime:                   stakingTime,
		StakingValue:                  stakingValue,
		SlashingTx:                    slashingTx,
		DelegatorSlashingSig:          delSlashingSig,
		UnbondingTx:                   unbondingTx,
		UnbondingTime:                 unbondingTime,
		UnbondingValue:                unbondingValue,
		UnbondingSlashingTx:           unbondingSlashingTx,
		DelegatorUnbondingSlashingSig: delUnbondingSlashingSig,
	}

	res, err := bc.reliablySendMsg(msg, emptyErrs, emptyErrs)
	if err != nil {
		return nil, err
	}

	return &types.TxResponse{TxHash: res.TxHash}, nil
}

func (bc *BabylonController) InsertBtcBlockHeaders(headers []bbntypes.BTCHeaderBytes) (*provider.RelayerTxResponse, error) {
	msg := &btclctypes.MsgInsertHeaders{
		Signer:  bc.mustGetTxSigner(),
		Headers: headers,
	}

	res, err := bc.reliablySendMsg(msg, emptyErrs, emptyErrs)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (bc *BabylonController) QueryFinalityProviders() ([]*btcstakingtypes.FinalityProviderResponse, error) {
	var fps []*btcstakingtypes.FinalityProviderResponse
	pagination := &sdkquery.PageRequest{
		Limit: 100,
	}

	for {
		res, err := bc.bbnClient.QueryClient.FinalityProviders(pagination)
		if err != nil {
			return nil, fmt.Errorf("failed to query finality providers: %w", err)
		}
		fps = append(fps, res.FinalityProviders...)
		if res.Pagination == nil || res.Pagination.NextKey == nil {
			break
		}

		pagination.Key = res.Pagination.NextKey
	}

	return fps, nil
}

func (bc *BabylonController) QueryFinalityProvider(fpPk *btcec.PublicKey) (*btcstakingtypes.QueryFinalityProviderResponse, error) {
	fpPubKey := bbntypes.NewBIP340PubKeyFromBTCPK(fpPk)
	res, err := bc.bbnClient.QueryClient.FinalityProvider(fpPubKey.MarshalHex())
	if err != nil {
		return nil, fmt.Errorf("failed to query the finality provider %s: %w", fpPubKey.MarshalHex(), err)
	}

	return res, nil
}

func (bc *BabylonController) EditFinalityProvider(fpPk *btcec.PublicKey,
	rate *sdkmath.LegacyDec, description []byte) (*btcstakingtypes.MsgEditFinalityProvider, error) {
	var reqDesc proto.Description
	if err := protobuf.Unmarshal(description, &reqDesc); err != nil {
		return nil, err
	}
	fpPubKey := bbntypes.NewBIP340PubKeyFromBTCPK(fpPk)

	fpRes, err := bc.QueryFinalityProvider(fpPk)
	if err != nil {
		return nil, err
	}

	if !strings.EqualFold(fpRes.FinalityProvider.Addr, bc.mustGetTxSigner()) {
		return nil, fmt.Errorf("the signer does not correspond to the finality provider's "+
			"Babylon address, expected %s got %s", bc.mustGetTxSigner(), fpRes.FinalityProvider.Addr)
	}

	getValueOrDefault := func(reqValue, defaultValue string) string {
		if reqValue != "" {
			return reqValue
		}
		return defaultValue
	}

	resDesc := fpRes.FinalityProvider.Description

	desc := &sttypes.Description{
		Moniker:         getValueOrDefault(reqDesc.Moniker, resDesc.Moniker),
		Identity:        getValueOrDefault(reqDesc.Identity, resDesc.Identity),
		Website:         getValueOrDefault(reqDesc.Website, resDesc.Website),
		SecurityContact: getValueOrDefault(reqDesc.SecurityContact, resDesc.SecurityContact),
		Details:         getValueOrDefault(reqDesc.Details, resDesc.Details),
	}

	msg := &btcstakingtypes.MsgEditFinalityProvider{
		Addr:        bc.mustGetTxSigner(),
		BtcPk:       fpPubKey.MustMarshal(),
		Description: desc,
		Commission:  fpRes.FinalityProvider.Commission,
	}

	if rate != nil {
		msg.Commission = rate
	}

	_, err = bc.reliablySendMsg(msg, emptyErrs, emptyErrs)
	if err != nil {
		return nil, fmt.Errorf("failed to query the finality provider %s: %w", fpPk.SerializeCompressed(), err)
	}

	return msg, nil
}

func (bc *BabylonController) QueryBtcLightClientTip() (*btclctypes.BTCHeaderInfoResponse, error) {
	res, err := bc.bbnClient.QueryClient.BTCHeaderChainTip()
	if err != nil {
		return nil, fmt.Errorf("failed to query BTC tip: %w", err)
	}

	return res.Header, nil
}

func (bc *BabylonController) QueryCurrentEpoch() (uint64, error) {
	res, err := bc.bbnClient.QueryClient.CurrentEpoch()
	if err != nil {
		return 0, fmt.Errorf("failed to query BTC tip: %w", err)
	}

	return res.CurrentEpoch, nil
}

func (bc *BabylonController) QueryVotesAtHeight(height uint64) ([]bbntypes.BIP340PubKey, error) {
	res, err := bc.bbnClient.QueryClient.VotesAtHeight(height)
	if err != nil {
		return nil, fmt.Errorf("failed to query BTC delegations: %w", err)
	}

	return res.BtcPks, nil
}

func (bc *BabylonController) QueryPendingDelegations(limit uint64) ([]*btcstakingtypes.BTCDelegationResponse, error) {
	return bc.queryDelegationsWithStatus(btcstakingtypes.BTCDelegationStatus_PENDING, limit)
}

func (bc *BabylonController) QueryActiveDelegations(limit uint64) ([]*btcstakingtypes.BTCDelegationResponse, error) {
	return bc.queryDelegationsWithStatus(btcstakingtypes.BTCDelegationStatus_ACTIVE, limit)
}

// queryDelegationsWithStatus queries BTC delegations
// with the given status (either pending or unbonding)
// it is only used when the program is running in Covenant mode
func (bc *BabylonController) queryDelegationsWithStatus(status btcstakingtypes.BTCDelegationStatus, limit uint64) ([]*btcstakingtypes.BTCDelegationResponse, error) {
	pagination := &sdkquery.PageRequest{
		Limit: limit,
	}

	res, err := bc.bbnClient.QueryClient.BTCDelegations(status, pagination)
	if err != nil {
		return nil, fmt.Errorf("failed to query BTC delegations: %w", err)
	}

	return res.BtcDelegations, nil
}

func (bc *BabylonController) QueryStakingParams() (*types.StakingParams, error) {
	// query btc checkpoint params
	ckptParamRes, err := bc.bbnClient.QueryClient.BTCCheckpointParams()
	if err != nil {
		return nil, fmt.Errorf("failed to query params of the btccheckpoint module: %w", err)
	}

	// query btc staking params
	stakingParamRes, err := bc.bbnClient.QueryClient.BTCStakingParams()
	if err != nil {
		return nil, fmt.Errorf("failed to query staking params: %w", err)
	}

	covenantPks := make([]*btcec.PublicKey, 0, len(stakingParamRes.Params.CovenantPks))
	for _, pk := range stakingParamRes.Params.CovenantPks {
		covPk, err := pk.ToBTCPK()
		if err != nil {
			return nil, fmt.Errorf("invalid covenant public key")
		}
		covenantPks = append(covenantPks, covPk)
	}

	return &types.StakingParams{
		ComfirmationTimeBlocks:    uint32(ckptParamRes.Params.BtcConfirmationDepth),
		FinalizationTimeoutBlocks: uint32(ckptParamRes.Params.CheckpointFinalizationTimeout),
		MinSlashingTxFeeSat:       btcutil.Amount(stakingParamRes.Params.MinSlashingTxFeeSat),
		CovenantPks:               covenantPks,
		SlashingPkScript:          stakingParamRes.Params.SlashingPkScript,
		CovenantQuorum:            stakingParamRes.Params.CovenantQuorum,
		SlashingRate:              stakingParamRes.Params.SlashingRate,
		UnbondingTime:             stakingParamRes.Params.MinUnbondingTimeBlocks,
	}, nil
}

func (bc *BabylonController) SubmitCovenantSigs(
	covPk *btcec.PublicKey,
	stakingTxHash string,
	slashingSigs [][]byte,
	unbondingSig *schnorr.Signature,
	unbondingSlashingSigs [][]byte,
) (*types.TxResponse, error) {
	bip340UnbondingSig := bbntypes.NewBIP340SignatureFromBTCSig(unbondingSig)

	msg := &btcstakingtypes.MsgAddCovenantSigs{
		Signer:                  bc.mustGetTxSigner(),
		Pk:                      bbntypes.NewBIP340PubKeyFromBTCPK(covPk),
		StakingTxHash:           stakingTxHash,
		SlashingTxSigs:          slashingSigs,
		UnbondingTxSig:          bip340UnbondingSig,
		SlashingUnbondingTxSigs: unbondingSlashingSigs,
	}

	res, err := bc.reliablySendMsg(msg, emptyErrs, emptyErrs)
	if err != nil {
		return nil, err
	}

	return &types.TxResponse{TxHash: res.TxHash, Events: res.Events}, nil
}

func (bc *BabylonController) GetBBNClient() *bbnclient.Client {
	return bc.bbnClient
}

func (bc *BabylonController) InsertSpvProofs(submitter string, proofs []*btcctypes.BTCSpvProof) (*provider.RelayerTxResponse, error) {
	msg := &btcctypes.MsgInsertBTCSpvProof{
		Submitter: submitter,
		Proofs:    proofs,
	}

	res, err := bc.reliablySendMsg(msg, emptyErrs, emptyErrs)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewCwClient(cwConfig *cwconfig.CosmwasmConfig, logger *zap.Logger) (*cwclient.Client, error) {
	if err := cwConfig.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config for OP consumer controller: %w", err)
	}

	bbnEncodingCfg := bbnapp.GetEncodingConfig()
	cwEncodingCfg := wasmdparams.EncodingConfig{
		InterfaceRegistry: bbnEncodingCfg.InterfaceRegistry,
		Codec:             bbnEncodingCfg.Codec,
		TxConfig:          bbnEncodingCfg.TxConfig,
		Amino:             bbnEncodingCfg.Amino,
	}

	cwClient, err := cwclient.New(
		cwConfig,
		BabylonChainName,
		cwEncodingCfg,
		logger,
	)

	return cwClient, err
}
