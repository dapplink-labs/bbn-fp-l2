// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MantaStakingMiddlewareDefaultVaultInitParams is an auto generated low-level Go binding around an user-defined struct.
type MantaStakingMiddlewareDefaultVaultInitParams struct {
	Version              uint64
	DelegatorIndex       uint64
	SlasherIndex         uint64
	VaultSlashBurner     common.Address
	VaultDefaultAdmin    common.Address
	VaultBeforeSlashHook common.Address
}

// MantaStakingMiddlewareMetaData contains all meta data concerning the MantaStakingMiddleware contract.
var MantaStakingMiddlewareMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_VAULT_PARAMS\",\"inputs\":[],\"outputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delegatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slasherIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"vaultSlashBurner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultDefaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultBeforeSlashHook\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REQUIRED_OPERATOR_STAKE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKE_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNREGISTER_TOKEN_UNLOCK_WINDOW\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VAULT_CONFIGURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimUnlockedToken\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"operatorRegistry_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultConfiguration_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"epochDuration_\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"requiredOperatorStake_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stakeToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"unregisterTokenUnlockWindow_\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"defaultVaultParams_\",\"type\":\"tuple\",\"internalType\":\"structMantaStakingMiddleware.DefaultVaultInitParams\",\"components\":[{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"delegatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slasherIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"vaultSlashBurner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultDefaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultBeforeSlashHook\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorTokenUnlockTimestamps\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operators\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paused\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slash\",\"inputs\":[{\"name\":\"subnetwork\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"captureTimestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[{\"name\":\"slashedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpauseOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterOperator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorPaused\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"slasher\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUnpaused\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUnregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"CannotClaimBeforeUnlock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotUnregisterPausedOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedToCreateVault\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorIsNotPendingUnlock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorIsPendingUnlock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegisteredToNetwork\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611be5806100d96000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c806372f9adab116100c3578063a314db161161007c578063a314db1614610305578063a70b9f0c14610318578063a876b89a14610331578063c93c6db414610339578063d547741f146103ca578063f3f2c422146103dd57600080fd5b806372f9adab146102a957806381de864a146102bc57806383ce0322146102cf57806391d14854146102e2578063a0f98db6146102f5578063a217fddf146102fd57600080fd5b80632e5aaf33116101155780632e5aaf33146102245780632f2ff15d1461023757806336568abe1461024a57806337148a241461025d5780635d91c2a6146102705780636d942667146102a057600080fd5b806301ffc9a71461015257806313e7c9d81461017a5780631c39b672146101ce578063248a9ca3146101f95780632acde0981461021a575b600080fd5b610165610160366004611546565b610405565b60405190151581526020015b60405180910390f35b6101af61018836600461159c565b6008602052600090815260409020546001600160a01b03811690600160a01b900460ff1682565b604080516001600160a01b039093168352901515602083015201610171565b6007546101e1906001600160a01b031681565b6040516001600160a01b039091168152602001610171565b61020c6102073660046115b9565b61043c565b604051908152602001610171565b61022261045e565b005b61022261023236600461159c565b6105c5565b6102226102453660046115d2565b610690565b6102226102583660046115d2565b6106b2565b61022261026b36600461162f565b6106e5565b60015461028990600160d01b900465ffffffffffff1681565b60405165ffffffffffff9091168152602001610171565b61020c60065481565b6102226102b736600461159c565b6109e5565b6001546101e1906001600160a01b031681565b6000546101e1906001600160a01b031681565b6101656102f03660046115d2565b610aab565b610222610ae3565b61020c600081565b61020c610313366004611759565b610b9c565b60015461028990600160a01b900465ffffffffffff1681565b610222610cef565b60025460035460045460055461037e936001600160401b0380821694600160401b8304821694600160801b909304909116926001600160a01b03918216928216911686565b604080516001600160401b039788168152958716602087015293909516928401929092526001600160a01b039081166060840152908116608083015290911660a082015260c001610171565b6102226103d83660046115d2565b610e1e565b6102896103eb36600461159c565b60096020526000908152604090205465ffffffffffff1681565b60006001600160e01b03198216637965db0b60e01b148061043657506301ffc9a760e01b6001600160e01b03198316145b92915050565b6000908152600080516020611b70833981519152602052604090206001015490565b610466610e3a565b3360008181526009602052604090205465ffffffffffff161561049c5760405163574a3d2d60e11b815260040160405180910390fd5b336000908152600860205260409020546001600160a01b0316156104d3576040516342ee68b560e01b815260040160405180910390fd5b6104dc33610e72565b6006546007546104fb916001600160a01b039091169033903090610efd565b600080600061050933610f64565b6040805180820182526001600160a01b03858116808352600060208085018281523380845260088352928790209551865491511515600160a01b026001600160a81b0319909216908616171790945584519081529283015284811692820192909252908216606082015292955090935091507f46646bab0dd157a3684acf6b4684ec8f475dfa314fe44fb2db2c67ab3f865ee19060800160405180910390a1505050506105c36001600080516020611b9083398151915255565b565b6001600160a01b038082166000908152600860205260409020548291166105ff576040516325ec6c1f60e01b815260040160405180910390fd5b600061060a81611280565b6001600160a01b038316600090815260086020526040902054600160a01b900460ff161561068b576001600160a01b038316600081815260086020908152604091829020805460ff60a01b1916905590519182527fae02c1bd695006b6d891af37fdeefea45a10ebcc17071e3471787db4f177288591015b60405180910390a15b505050565b6106998261043c565b6106a281611280565b6106ac838361128a565b50505050565b6001600160a01b03811633146106db5760405163334bd91960e11b815260040160405180910390fd5b61068b8282611336565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b031660008115801561072a5750825b90506000826001600160401b031660011480156107465750303b155b905081158015610754575080155b156107725760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561079c57845460ff60401b1916600160401b1785555b6107a46113b2565b6107af60003361128a565b506107b86113ba565b8b6000806101000a8154816001600160a01b0302191690836001600160a01b031602179055508a600160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555089600160146101000a81548165ffffffffffff021916908365ffffffffffff1602179055508860068190555087600760006101000a8154816001600160a01b0302191690836001600160a01b0316021790555085600260008201518160000160006101000a8154816001600160401b0302191690836001600160401b0316021790555060208201518160000160086101000a8154816001600160401b0302191690836001600160401b0316021790555060408201518160000160106101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160010160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060808201518160020160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060a08201518160030160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550905050866001601a6101000a81548165ffffffffffff021916908365ffffffffffff16021790555083156109d757845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b6001600160a01b03808216600090815260086020526040902054829116610a1f576040516325ec6c1f60e01b815260040160405180910390fd5b6000610a2a81611280565b6001600160a01b038316600090815260086020526040902054600160a01b900460ff1661068b576001600160a01b038316600081815260086020908152604091829020805460ff60a01b1916600160a01b17905590519182527fc5437eb8dd091f69800961953f2bb0bc16ae1ff2d3e52caa96796db65f8271da9101610682565b6000918252600080516020611b70833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3360009081526009602052604081205465ffffffffffff169003610b1a57604051635e0580eb60e01b815260040160405180910390fd5b3360009081526009602052604090205465ffffffffffff16610b3a6113ca565b65ffffffffffff161015610b61576040516352d1770f60e01b815260040160405180910390fd5b600654600754610b7e916001600160a01b039091169033906113da565b336000908152600960205260409020805465ffffffffffff19169055565b6001600160a01b038084166000908152600860205260408120549091859116610bd8576040516325ec6c1f60e01b815260040160405180910390fd5b6000610be381611280565b6001600160a01b03808716600090815260086020908152604080832054815163b134427160e01b81529151941693849263b134427192600480820193918290030181865afa158015610c39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c5d91906117a1565b6040805160208101825260008152905163010d40ab60e11b81529192506001600160a01b0383169163021a815691610c9f918d918d918d918d91600401611804565b6020604051808303816000875af1158015610cbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce2919061184a565b9998505050505050505050565b336000818152600860205260409020546001600160a01b0316610d25576040516325ec6c1f60e01b815260040160405180910390fd5b610d2d610e3a565b33600090815260086020526040902054600160a01b900460ff1615610d645760405162417c2d60e11b815260040160405180910390fd5b33600090815260086020526040902080546001600160a81b0319169055600154600160d01b900465ffffffffffff16610d9b6113ca565b610da59190611863565b33600081815260096020908152604091829020805465ffffffffffff191665ffffffffffff9590951694909417909355519081527f6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f910160405180910390a1610e1b6001600080516020611b9083398151915255565b50565b610e278261043c565b610e3081611280565b6106ac8383611336565b600080516020611b90833981519152805460011901610e6c57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6000546040516302910f8b60e31b81526001600160a01b038381166004830152909116906314887c5890602401602060405180830381865afa158015610ebc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ee09190611890565b610e1b57604051632c4b254960e01b815260040160405180910390fd5b6040516001600160a01b0384811660248301528381166044830152606482018390526106ac9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b03838183161783525050505061140b565b60408051610160810182526007546001600160a01b0390811682526003548116602083015260015465ffffffffffff600160a01b9091041682840152600060608084018290526080840182905260a0840182905260045490921660c084015260e0830181905261010083018190526101208301819052610140830181905283516002808252928101909452928392839290918391816020016020820280368337505060045482519293506001600160a01b03169183915060009061102a5761102a6118b2565b60200260200101906001600160a01b031690816001600160a01b031681525050308160018151811061105e5761105e6118b2565b6001600160a01b039283166020918202929092018101919091526040805160c0810182526004548416606082018181526005548616608084015260a083018290528252818401869052938a168183015281518083018352600181850190815281528251610100810184526002546001600160401b03168152808501959095528251919490936000939192918301916110f8918991016118c8565b6040516020818303038152906040528152602001600260000160089054906101000a90046001600160401b03166001600160401b031681526020018460405160200161114491906119b0565b60408051808303601f1901815291815290825260016020808401829052600254600160801b90046001600160401b031684840152825187515115158183015283518082039092018252830183526060909301929092529054905163312249f960e21b81529192506001600160a01b03169063c48927e4906111c9908490600401611a44565b6060604051808303816000875af11580156111e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061120c9190611b22565b919950975095506001600160a01b038816158061123057506001600160a01b038716155b8061124257506001600160a01b038616155b15611260576040516307cbdadb60e11b815260040160405180910390fd5b50505050509193909250565b6001600080516020611b9083398151915255565b610e1b8133611481565b6000600080516020611b708339815191526112a58484610aab565b611325576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556112db3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610436565b6000915050610436565b5092915050565b6000600080516020611b708339815191526113518484610aab565b15611325576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610436565b6105c36114be565b6113c26114be565b6105c3611507565b60006113d54261150f565b905090565b6040516001600160a01b0383811660248301526044820183905261068b91859182169063a9059cbb90606401610f32565b600080602060008451602086016000885af18061142e576040513d6000823e3d81fd5b50506000513d91508115611446578060011415611453565b6001600160a01b0384163b155b156106ac57604051635274afe760e01b81526001600160a01b03851660048201526024015b60405180910390fd5b61148b8282610aab565b6114ba5760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401611478565b5050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166105c357604051631afcd79f60e31b815260040160405180910390fd5b61126c6114be565b600065ffffffffffff821115611542576040516306dfcc6560e41b81526030600482015260248101839052604401611478565b5090565b60006020828403121561155857600080fd5b81356001600160e01b03198116811461157057600080fd5b9392505050565b6001600160a01b0381168114610e1b57600080fd5b803561159781611577565b919050565b6000602082840312156115ae57600080fd5b813561157081611577565b6000602082840312156115cb57600080fd5b5035919050565b600080604083850312156115e557600080fd5b8235915060208301356115f781611577565b809150509250929050565b803565ffffffffffff8116811461159757600080fd5b80356001600160401b038116811461159757600080fd5b600080600080600080600087890361018081121561164c57600080fd5b883561165781611577565b9750602089013561166781611577565b965061167560408a01611602565b955060608901359450608089013561168c81611577565b935061169a60a08a01611602565b925060c060bf19820112156116ae57600080fd5b5060405160c081018181106001600160401b03821117156116df57634e487b7160e01b600052604160045260246000fd5b6040526116ee60c08a01611618565b81526116fc60e08a01611618565b602082015261170e6101008a01611618565b60408201526117206101208a0161158c565b60608201526117326101408a0161158c565b60808201526117446101608a0161158c565b60a08201528091505092959891949750929550565b6000806000806080858703121561176f57600080fd5b84359350602085013561178181611577565b92506040850135915061179660608601611602565b905092959194509250565b6000602082840312156117b357600080fd5b815161157081611577565b6000815180845260005b818110156117e4576020818501810151868301820152016117c8565b506000602082860101526020601f19601f83011685010191505092915050565b85815260018060a01b038516602082015283604082015265ffffffffffff8316606082015260a06080820152600061183f60a08301846117be565b979650505050505050565b60006020828403121561185c57600080fd5b5051919050565b65ffffffffffff81811683821601908082111561132f57634e487b7160e01b600052601160045260246000fd5b6000602082840312156118a257600080fd5b8151801515811461157057600080fd5b634e487b7160e01b600052603260045260246000fd5b81516001600160a01b03168152610160810160208301516118f460208401826001600160a01b03169052565b50604083015161190e604084018265ffffffffffff169052565b506060830151611922606084018215159052565b506080830151611936608084018215159052565b5060a083015160a083015260c083015161195b60c08401826001600160a01b03169052565b5060e083015161197660e08401826001600160a01b03169052565b50610100838101516001600160a01b0390811691840191909152610120808501518216908401526101409384015116929091019190915290565b6020808252825180516001600160a01b039081168484015281830151811660408086019190915290910151811660608401528382015160a06080850152805160c085018190526000939291830191849160e08701905b80841015611a2857845183168252938501936001939093019290850190611a06565b5060408801516001600160a01b03811660a0890152945061183f565b60208152611a5e6020820183516001600160401b03169052565b60006020830151611a7a60408401826001600160a01b03169052565b506040830151610100806060850152611a976101208501836117be565b91506060850151611ab360808601826001600160401b03169052565b506080850151601f19808685030160a0870152611ad084836117be565b935060a08701519150611ae760c087018315159052565b60c08701516001600160401b03811660e0880152915060e0870151915080868503018387015250611b1883826117be565b9695505050505050565b600080600060608486031215611b3757600080fd5b8351611b4281611577565b6020850151909350611b5381611577565b6040850151909250611b6481611577565b80915050925092509256fe02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220c6078f12275d958856c1e5f7d5280a3557f9547065dcc95db834df588547079e64736f6c63430008190033",
}

// MantaStakingMiddlewareABI is the input ABI used to generate the binding from.
// Deprecated: Use MantaStakingMiddlewareMetaData.ABI instead.
var MantaStakingMiddlewareABI = MantaStakingMiddlewareMetaData.ABI

// MantaStakingMiddlewareBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MantaStakingMiddlewareMetaData.Bin instead.
var MantaStakingMiddlewareBin = MantaStakingMiddlewareMetaData.Bin

// DeployMantaStakingMiddleware deploys a new Ethereum contract, binding an instance of MantaStakingMiddleware to it.
func DeployMantaStakingMiddleware(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MantaStakingMiddleware, error) {
	parsed, err := MantaStakingMiddlewareMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MantaStakingMiddlewareBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MantaStakingMiddleware{MantaStakingMiddlewareCaller: MantaStakingMiddlewareCaller{contract: contract}, MantaStakingMiddlewareTransactor: MantaStakingMiddlewareTransactor{contract: contract}, MantaStakingMiddlewareFilterer: MantaStakingMiddlewareFilterer{contract: contract}}, nil
}

// MantaStakingMiddleware is an auto generated Go binding around an Ethereum contract.
type MantaStakingMiddleware struct {
	MantaStakingMiddlewareCaller     // Read-only binding to the contract
	MantaStakingMiddlewareTransactor // Write-only binding to the contract
	MantaStakingMiddlewareFilterer   // Log filterer for contract events
}

// MantaStakingMiddlewareCaller is an auto generated read-only Go binding around an Ethereum contract.
type MantaStakingMiddlewareCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MantaStakingMiddlewareTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MantaStakingMiddlewareTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MantaStakingMiddlewareFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MantaStakingMiddlewareFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MantaStakingMiddlewareSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MantaStakingMiddlewareSession struct {
	Contract     *MantaStakingMiddleware // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MantaStakingMiddlewareCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MantaStakingMiddlewareCallerSession struct {
	Contract *MantaStakingMiddlewareCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MantaStakingMiddlewareTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MantaStakingMiddlewareTransactorSession struct {
	Contract     *MantaStakingMiddlewareTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MantaStakingMiddlewareRaw is an auto generated low-level Go binding around an Ethereum contract.
type MantaStakingMiddlewareRaw struct {
	Contract *MantaStakingMiddleware // Generic contract binding to access the raw methods on
}

// MantaStakingMiddlewareCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MantaStakingMiddlewareCallerRaw struct {
	Contract *MantaStakingMiddlewareCaller // Generic read-only contract binding to access the raw methods on
}

// MantaStakingMiddlewareTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MantaStakingMiddlewareTransactorRaw struct {
	Contract *MantaStakingMiddlewareTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMantaStakingMiddleware creates a new instance of MantaStakingMiddleware, bound to a specific deployed contract.
func NewMantaStakingMiddleware(address common.Address, backend bind.ContractBackend) (*MantaStakingMiddleware, error) {
	contract, err := bindMantaStakingMiddleware(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddleware{MantaStakingMiddlewareCaller: MantaStakingMiddlewareCaller{contract: contract}, MantaStakingMiddlewareTransactor: MantaStakingMiddlewareTransactor{contract: contract}, MantaStakingMiddlewareFilterer: MantaStakingMiddlewareFilterer{contract: contract}}, nil
}

// NewMantaStakingMiddlewareCaller creates a new read-only instance of MantaStakingMiddleware, bound to a specific deployed contract.
func NewMantaStakingMiddlewareCaller(address common.Address, caller bind.ContractCaller) (*MantaStakingMiddlewareCaller, error) {
	contract, err := bindMantaStakingMiddleware(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareCaller{contract: contract}, nil
}

// NewMantaStakingMiddlewareTransactor creates a new write-only instance of MantaStakingMiddleware, bound to a specific deployed contract.
func NewMantaStakingMiddlewareTransactor(address common.Address, transactor bind.ContractTransactor) (*MantaStakingMiddlewareTransactor, error) {
	contract, err := bindMantaStakingMiddleware(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareTransactor{contract: contract}, nil
}

// NewMantaStakingMiddlewareFilterer creates a new log filterer instance of MantaStakingMiddleware, bound to a specific deployed contract.
func NewMantaStakingMiddlewareFilterer(address common.Address, filterer bind.ContractFilterer) (*MantaStakingMiddlewareFilterer, error) {
	contract, err := bindMantaStakingMiddleware(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareFilterer{contract: contract}, nil
}

// bindMantaStakingMiddleware binds a generic wrapper to an already deployed contract.
func bindMantaStakingMiddleware(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MantaStakingMiddlewareMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MantaStakingMiddleware *MantaStakingMiddlewareRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MantaStakingMiddleware.Contract.MantaStakingMiddlewareCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MantaStakingMiddleware *MantaStakingMiddlewareRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.MantaStakingMiddlewareTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MantaStakingMiddleware *MantaStakingMiddlewareRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.MantaStakingMiddlewareTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MantaStakingMiddleware.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MantaStakingMiddleware.Contract.DEFAULTADMINROLE(&_MantaStakingMiddleware.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MantaStakingMiddleware.Contract.DEFAULTADMINROLE(&_MantaStakingMiddleware.CallOpts)
}

// DEFAULTVAULTPARAMS is a free data retrieval call binding the contract method 0xc93c6db4.
//
// Solidity: function DEFAULT_VAULT_PARAMS() view returns(uint64 version, uint64 delegatorIndex, uint64 slasherIndex, address vaultSlashBurner, address vaultDefaultAdmin, address vaultBeforeSlashHook)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) DEFAULTVAULTPARAMS(opts *bind.CallOpts) (struct {
	Version              uint64
	DelegatorIndex       uint64
	SlasherIndex         uint64
	VaultSlashBurner     common.Address
	VaultDefaultAdmin    common.Address
	VaultBeforeSlashHook common.Address
}, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "DEFAULT_VAULT_PARAMS")

	outstruct := new(struct {
		Version              uint64
		DelegatorIndex       uint64
		SlasherIndex         uint64
		VaultSlashBurner     common.Address
		VaultDefaultAdmin    common.Address
		VaultBeforeSlashHook common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Version = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.DelegatorIndex = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.SlasherIndex = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.VaultSlashBurner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.VaultDefaultAdmin = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.VaultBeforeSlashHook = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DEFAULTVAULTPARAMS is a free data retrieval call binding the contract method 0xc93c6db4.
//
// Solidity: function DEFAULT_VAULT_PARAMS() view returns(uint64 version, uint64 delegatorIndex, uint64 slasherIndex, address vaultSlashBurner, address vaultDefaultAdmin, address vaultBeforeSlashHook)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) DEFAULTVAULTPARAMS() (struct {
	Version              uint64
	DelegatorIndex       uint64
	SlasherIndex         uint64
	VaultSlashBurner     common.Address
	VaultDefaultAdmin    common.Address
	VaultBeforeSlashHook common.Address
}, error) {
	return _MantaStakingMiddleware.Contract.DEFAULTVAULTPARAMS(&_MantaStakingMiddleware.CallOpts)
}

// DEFAULTVAULTPARAMS is a free data retrieval call binding the contract method 0xc93c6db4.
//
// Solidity: function DEFAULT_VAULT_PARAMS() view returns(uint64 version, uint64 delegatorIndex, uint64 slasherIndex, address vaultSlashBurner, address vaultDefaultAdmin, address vaultBeforeSlashHook)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) DEFAULTVAULTPARAMS() (struct {
	Version              uint64
	DelegatorIndex       uint64
	SlasherIndex         uint64
	VaultSlashBurner     common.Address
	VaultDefaultAdmin    common.Address
	VaultBeforeSlashHook common.Address
}, error) {
	return _MantaStakingMiddleware.Contract.DEFAULTVAULTPARAMS(&_MantaStakingMiddleware.CallOpts)
}

// EPOCHDURATION is a free data retrieval call binding the contract method 0xa70b9f0c.
//
// Solidity: function EPOCH_DURATION() view returns(uint48)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) EPOCHDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "EPOCH_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EPOCHDURATION is a free data retrieval call binding the contract method 0xa70b9f0c.
//
// Solidity: function EPOCH_DURATION() view returns(uint48)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) EPOCHDURATION() (*big.Int, error) {
	return _MantaStakingMiddleware.Contract.EPOCHDURATION(&_MantaStakingMiddleware.CallOpts)
}

// EPOCHDURATION is a free data retrieval call binding the contract method 0xa70b9f0c.
//
// Solidity: function EPOCH_DURATION() view returns(uint48)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) EPOCHDURATION() (*big.Int, error) {
	return _MantaStakingMiddleware.Contract.EPOCHDURATION(&_MantaStakingMiddleware.CallOpts)
}

// OPERATORREGISTRY is a free data retrieval call binding the contract method 0x83ce0322.
//
// Solidity: function OPERATOR_REGISTRY() view returns(address)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) OPERATORREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "OPERATOR_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OPERATORREGISTRY is a free data retrieval call binding the contract method 0x83ce0322.
//
// Solidity: function OPERATOR_REGISTRY() view returns(address)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) OPERATORREGISTRY() (common.Address, error) {
	return _MantaStakingMiddleware.Contract.OPERATORREGISTRY(&_MantaStakingMiddleware.CallOpts)
}

// OPERATORREGISTRY is a free data retrieval call binding the contract method 0x83ce0322.
//
// Solidity: function OPERATOR_REGISTRY() view returns(address)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) OPERATORREGISTRY() (common.Address, error) {
	return _MantaStakingMiddleware.Contract.OPERATORREGISTRY(&_MantaStakingMiddleware.CallOpts)
}

// REQUIREDOPERATORSTAKE is a free data retrieval call binding the contract method 0x6d942667.
//
// Solidity: function REQUIRED_OPERATOR_STAKE() view returns(uint256)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) REQUIREDOPERATORSTAKE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "REQUIRED_OPERATOR_STAKE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REQUIREDOPERATORSTAKE is a free data retrieval call binding the contract method 0x6d942667.
//
// Solidity: function REQUIRED_OPERATOR_STAKE() view returns(uint256)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) REQUIREDOPERATORSTAKE() (*big.Int, error) {
	return _MantaStakingMiddleware.Contract.REQUIREDOPERATORSTAKE(&_MantaStakingMiddleware.CallOpts)
}

// REQUIREDOPERATORSTAKE is a free data retrieval call binding the contract method 0x6d942667.
//
// Solidity: function REQUIRED_OPERATOR_STAKE() view returns(uint256)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) REQUIREDOPERATORSTAKE() (*big.Int, error) {
	return _MantaStakingMiddleware.Contract.REQUIREDOPERATORSTAKE(&_MantaStakingMiddleware.CallOpts)
}

// STAKETOKEN is a free data retrieval call binding the contract method 0x1c39b672.
//
// Solidity: function STAKE_TOKEN() view returns(address)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) STAKETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "STAKE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKETOKEN is a free data retrieval call binding the contract method 0x1c39b672.
//
// Solidity: function STAKE_TOKEN() view returns(address)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) STAKETOKEN() (common.Address, error) {
	return _MantaStakingMiddleware.Contract.STAKETOKEN(&_MantaStakingMiddleware.CallOpts)
}

// STAKETOKEN is a free data retrieval call binding the contract method 0x1c39b672.
//
// Solidity: function STAKE_TOKEN() view returns(address)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) STAKETOKEN() (common.Address, error) {
	return _MantaStakingMiddleware.Contract.STAKETOKEN(&_MantaStakingMiddleware.CallOpts)
}

// UNREGISTERTOKENUNLOCKWINDOW is a free data retrieval call binding the contract method 0x5d91c2a6.
//
// Solidity: function UNREGISTER_TOKEN_UNLOCK_WINDOW() view returns(uint48)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) UNREGISTERTOKENUNLOCKWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "UNREGISTER_TOKEN_UNLOCK_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNREGISTERTOKENUNLOCKWINDOW is a free data retrieval call binding the contract method 0x5d91c2a6.
//
// Solidity: function UNREGISTER_TOKEN_UNLOCK_WINDOW() view returns(uint48)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) UNREGISTERTOKENUNLOCKWINDOW() (*big.Int, error) {
	return _MantaStakingMiddleware.Contract.UNREGISTERTOKENUNLOCKWINDOW(&_MantaStakingMiddleware.CallOpts)
}

// UNREGISTERTOKENUNLOCKWINDOW is a free data retrieval call binding the contract method 0x5d91c2a6.
//
// Solidity: function UNREGISTER_TOKEN_UNLOCK_WINDOW() view returns(uint48)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) UNREGISTERTOKENUNLOCKWINDOW() (*big.Int, error) {
	return _MantaStakingMiddleware.Contract.UNREGISTERTOKENUNLOCKWINDOW(&_MantaStakingMiddleware.CallOpts)
}

// VAULTCONFIGURATION is a free data retrieval call binding the contract method 0x81de864a.
//
// Solidity: function VAULT_CONFIGURATION() view returns(address)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) VAULTCONFIGURATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "VAULT_CONFIGURATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VAULTCONFIGURATION is a free data retrieval call binding the contract method 0x81de864a.
//
// Solidity: function VAULT_CONFIGURATION() view returns(address)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) VAULTCONFIGURATION() (common.Address, error) {
	return _MantaStakingMiddleware.Contract.VAULTCONFIGURATION(&_MantaStakingMiddleware.CallOpts)
}

// VAULTCONFIGURATION is a free data retrieval call binding the contract method 0x81de864a.
//
// Solidity: function VAULT_CONFIGURATION() view returns(address)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) VAULTCONFIGURATION() (common.Address, error) {
	return _MantaStakingMiddleware.Contract.VAULTCONFIGURATION(&_MantaStakingMiddleware.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MantaStakingMiddleware.Contract.GetRoleAdmin(&_MantaStakingMiddleware.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MantaStakingMiddleware.Contract.GetRoleAdmin(&_MantaStakingMiddleware.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MantaStakingMiddleware.Contract.HasRole(&_MantaStakingMiddleware.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MantaStakingMiddleware.Contract.HasRole(&_MantaStakingMiddleware.CallOpts, role, account)
}

// OperatorTokenUnlockTimestamps is a free data retrieval call binding the contract method 0xf3f2c422.
//
// Solidity: function operatorTokenUnlockTimestamps(address ) view returns(uint48)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) OperatorTokenUnlockTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "operatorTokenUnlockTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorTokenUnlockTimestamps is a free data retrieval call binding the contract method 0xf3f2c422.
//
// Solidity: function operatorTokenUnlockTimestamps(address ) view returns(uint48)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) OperatorTokenUnlockTimestamps(arg0 common.Address) (*big.Int, error) {
	return _MantaStakingMiddleware.Contract.OperatorTokenUnlockTimestamps(&_MantaStakingMiddleware.CallOpts, arg0)
}

// OperatorTokenUnlockTimestamps is a free data retrieval call binding the contract method 0xf3f2c422.
//
// Solidity: function operatorTokenUnlockTimestamps(address ) view returns(uint48)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) OperatorTokenUnlockTimestamps(arg0 common.Address) (*big.Int, error) {
	return _MantaStakingMiddleware.Contract.OperatorTokenUnlockTimestamps(&_MantaStakingMiddleware.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(address vault, bool paused)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Vault  common.Address
	Paused bool
}, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "operators", arg0)

	outstruct := new(struct {
		Vault  common.Address
		Paused bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Vault = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Paused = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(address vault, bool paused)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) Operators(arg0 common.Address) (struct {
	Vault  common.Address
	Paused bool
}, error) {
	return _MantaStakingMiddleware.Contract.Operators(&_MantaStakingMiddleware.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(address vault, bool paused)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) Operators(arg0 common.Address) (struct {
	Vault  common.Address
	Paused bool
}, error) {
	return _MantaStakingMiddleware.Contract.Operators(&_MantaStakingMiddleware.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MantaStakingMiddleware.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MantaStakingMiddleware.Contract.SupportsInterface(&_MantaStakingMiddleware.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MantaStakingMiddleware *MantaStakingMiddlewareCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MantaStakingMiddleware.Contract.SupportsInterface(&_MantaStakingMiddleware.CallOpts, interfaceId)
}

// ClaimUnlockedToken is a paid mutator transaction binding the contract method 0xa0f98db6.
//
// Solidity: function claimUnlockedToken() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) ClaimUnlockedToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "claimUnlockedToken")
}

// ClaimUnlockedToken is a paid mutator transaction binding the contract method 0xa0f98db6.
//
// Solidity: function claimUnlockedToken() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) ClaimUnlockedToken() (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.ClaimUnlockedToken(&_MantaStakingMiddleware.TransactOpts)
}

// ClaimUnlockedToken is a paid mutator transaction binding the contract method 0xa0f98db6.
//
// Solidity: function claimUnlockedToken() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) ClaimUnlockedToken() (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.ClaimUnlockedToken(&_MantaStakingMiddleware.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.GrantRole(&_MantaStakingMiddleware.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.GrantRole(&_MantaStakingMiddleware.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x37148a24.
//
// Solidity: function initialize(address operatorRegistry_, address vaultConfiguration_, uint48 epochDuration_, uint256 requiredOperatorStake_, address stakeToken_, uint48 unregisterTokenUnlockWindow_, (uint64,uint64,uint64,address,address,address) defaultVaultParams_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) Initialize(opts *bind.TransactOpts, operatorRegistry_ common.Address, vaultConfiguration_ common.Address, epochDuration_ *big.Int, requiredOperatorStake_ *big.Int, stakeToken_ common.Address, unregisterTokenUnlockWindow_ *big.Int, defaultVaultParams_ MantaStakingMiddlewareDefaultVaultInitParams) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "initialize", operatorRegistry_, vaultConfiguration_, epochDuration_, requiredOperatorStake_, stakeToken_, unregisterTokenUnlockWindow_, defaultVaultParams_)
}

// Initialize is a paid mutator transaction binding the contract method 0x37148a24.
//
// Solidity: function initialize(address operatorRegistry_, address vaultConfiguration_, uint48 epochDuration_, uint256 requiredOperatorStake_, address stakeToken_, uint48 unregisterTokenUnlockWindow_, (uint64,uint64,uint64,address,address,address) defaultVaultParams_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) Initialize(operatorRegistry_ common.Address, vaultConfiguration_ common.Address, epochDuration_ *big.Int, requiredOperatorStake_ *big.Int, stakeToken_ common.Address, unregisterTokenUnlockWindow_ *big.Int, defaultVaultParams_ MantaStakingMiddlewareDefaultVaultInitParams) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.Initialize(&_MantaStakingMiddleware.TransactOpts, operatorRegistry_, vaultConfiguration_, epochDuration_, requiredOperatorStake_, stakeToken_, unregisterTokenUnlockWindow_, defaultVaultParams_)
}

// Initialize is a paid mutator transaction binding the contract method 0x37148a24.
//
// Solidity: function initialize(address operatorRegistry_, address vaultConfiguration_, uint48 epochDuration_, uint256 requiredOperatorStake_, address stakeToken_, uint48 unregisterTokenUnlockWindow_, (uint64,uint64,uint64,address,address,address) defaultVaultParams_) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) Initialize(operatorRegistry_ common.Address, vaultConfiguration_ common.Address, epochDuration_ *big.Int, requiredOperatorStake_ *big.Int, stakeToken_ common.Address, unregisterTokenUnlockWindow_ *big.Int, defaultVaultParams_ MantaStakingMiddlewareDefaultVaultInitParams) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.Initialize(&_MantaStakingMiddleware.TransactOpts, operatorRegistry_, vaultConfiguration_, epochDuration_, requiredOperatorStake_, stakeToken_, unregisterTokenUnlockWindow_, defaultVaultParams_)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) PauseOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "pauseOperator", operator)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) PauseOperator(operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.PauseOperator(&_MantaStakingMiddleware.TransactOpts, operator)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) PauseOperator(operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.PauseOperator(&_MantaStakingMiddleware.TransactOpts, operator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x2acde098.
//
// Solidity: function registerOperator() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) RegisterOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "registerOperator")
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x2acde098.
//
// Solidity: function registerOperator() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) RegisterOperator() (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RegisterOperator(&_MantaStakingMiddleware.TransactOpts)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x2acde098.
//
// Solidity: function registerOperator() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) RegisterOperator() (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RegisterOperator(&_MantaStakingMiddleware.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RenounceRole(&_MantaStakingMiddleware.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RenounceRole(&_MantaStakingMiddleware.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RevokeRole(&_MantaStakingMiddleware.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.RevokeRole(&_MantaStakingMiddleware.TransactOpts, role, account)
}

// Slash is a paid mutator transaction binding the contract method 0xa314db16.
//
// Solidity: function slash(bytes32 subnetwork, address operator, uint256 amount, uint48 captureTimestamp) returns(uint256 slashedAmount)
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) Slash(opts *bind.TransactOpts, subnetwork [32]byte, operator common.Address, amount *big.Int, captureTimestamp *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "slash", subnetwork, operator, amount, captureTimestamp)
}

// Slash is a paid mutator transaction binding the contract method 0xa314db16.
//
// Solidity: function slash(bytes32 subnetwork, address operator, uint256 amount, uint48 captureTimestamp) returns(uint256 slashedAmount)
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) Slash(subnetwork [32]byte, operator common.Address, amount *big.Int, captureTimestamp *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.Slash(&_MantaStakingMiddleware.TransactOpts, subnetwork, operator, amount, captureTimestamp)
}

// Slash is a paid mutator transaction binding the contract method 0xa314db16.
//
// Solidity: function slash(bytes32 subnetwork, address operator, uint256 amount, uint48 captureTimestamp) returns(uint256 slashedAmount)
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) Slash(subnetwork [32]byte, operator common.Address, amount *big.Int, captureTimestamp *big.Int) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.Slash(&_MantaStakingMiddleware.TransactOpts, subnetwork, operator, amount, captureTimestamp)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) UnpauseOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "unpauseOperator", operator)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) UnpauseOperator(operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UnpauseOperator(&_MantaStakingMiddleware.TransactOpts, operator)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address operator) returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) UnpauseOperator(operator common.Address) (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UnpauseOperator(&_MantaStakingMiddleware.TransactOpts, operator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0xa876b89a.
//
// Solidity: function unregisterOperator() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactor) UnregisterOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MantaStakingMiddleware.contract.Transact(opts, "unregisterOperator")
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0xa876b89a.
//
// Solidity: function unregisterOperator() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareSession) UnregisterOperator() (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UnregisterOperator(&_MantaStakingMiddleware.TransactOpts)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0xa876b89a.
//
// Solidity: function unregisterOperator() returns()
func (_MantaStakingMiddleware *MantaStakingMiddlewareTransactorSession) UnregisterOperator() (*types.Transaction, error) {
	return _MantaStakingMiddleware.Contract.UnregisterOperator(&_MantaStakingMiddleware.TransactOpts)
}

// MantaStakingMiddlewareInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareInitializedIterator struct {
	Event *MantaStakingMiddlewareInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareInitialized represents a Initialized event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterInitialized(opts *bind.FilterOpts) (*MantaStakingMiddlewareInitializedIterator, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareInitializedIterator{contract: _MantaStakingMiddleware.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareInitialized) (event.Subscription, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareInitialized)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseInitialized(log types.Log) (*MantaStakingMiddlewareInitialized, error) {
	event := new(MantaStakingMiddlewareInitialized)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareOperatorPausedIterator is returned from FilterOperatorPaused and is used to iterate over the raw logs and unpacked data for OperatorPaused events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorPausedIterator struct {
	Event *MantaStakingMiddlewareOperatorPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareOperatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareOperatorPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareOperatorPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareOperatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareOperatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareOperatorPaused represents a OperatorPaused event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorPaused struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorPaused is a free log retrieval operation binding the contract event 0xc5437eb8dd091f69800961953f2bb0bc16ae1ff2d3e52caa96796db65f8271da.
//
// Solidity: event OperatorPaused(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterOperatorPaused(opts *bind.FilterOpts) (*MantaStakingMiddlewareOperatorPausedIterator, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "OperatorPaused")
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareOperatorPausedIterator{contract: _MantaStakingMiddleware.contract, event: "OperatorPaused", logs: logs, sub: sub}, nil
}

// WatchOperatorPaused is a free log subscription operation binding the contract event 0xc5437eb8dd091f69800961953f2bb0bc16ae1ff2d3e52caa96796db65f8271da.
//
// Solidity: event OperatorPaused(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchOperatorPaused(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareOperatorPaused) (event.Subscription, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "OperatorPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareOperatorPaused)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorPaused is a log parse operation binding the contract event 0xc5437eb8dd091f69800961953f2bb0bc16ae1ff2d3e52caa96796db65f8271da.
//
// Solidity: event OperatorPaused(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseOperatorPaused(log types.Log) (*MantaStakingMiddlewareOperatorPaused, error) {
	event := new(MantaStakingMiddlewareOperatorPaused)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorRegisteredIterator struct {
	Event *MantaStakingMiddlewareOperatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareOperatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareOperatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareOperatorRegistered represents a OperatorRegistered event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorRegistered struct {
	Operator  common.Address
	Vault     common.Address
	Delegator common.Address
	Slasher   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x46646bab0dd157a3684acf6b4684ec8f475dfa314fe44fb2db2c67ab3f865ee1.
//
// Solidity: event OperatorRegistered(address operator, address vault, address delegator, address slasher)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterOperatorRegistered(opts *bind.FilterOpts) (*MantaStakingMiddlewareOperatorRegisteredIterator, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "OperatorRegistered")
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareOperatorRegisteredIterator{contract: _MantaStakingMiddleware.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x46646bab0dd157a3684acf6b4684ec8f475dfa314fe44fb2db2c67ab3f865ee1.
//
// Solidity: event OperatorRegistered(address operator, address vault, address delegator, address slasher)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareOperatorRegistered) (event.Subscription, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "OperatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareOperatorRegistered)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRegistered is a log parse operation binding the contract event 0x46646bab0dd157a3684acf6b4684ec8f475dfa314fe44fb2db2c67ab3f865ee1.
//
// Solidity: event OperatorRegistered(address operator, address vault, address delegator, address slasher)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseOperatorRegistered(log types.Log) (*MantaStakingMiddlewareOperatorRegistered, error) {
	event := new(MantaStakingMiddlewareOperatorRegistered)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareOperatorUnpausedIterator is returned from FilterOperatorUnpaused and is used to iterate over the raw logs and unpacked data for OperatorUnpaused events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorUnpausedIterator struct {
	Event *MantaStakingMiddlewareOperatorUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareOperatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareOperatorUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareOperatorUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareOperatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareOperatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareOperatorUnpaused represents a OperatorUnpaused event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorUnpaused struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorUnpaused is a free log retrieval operation binding the contract event 0xae02c1bd695006b6d891af37fdeefea45a10ebcc17071e3471787db4f1772885.
//
// Solidity: event OperatorUnpaused(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterOperatorUnpaused(opts *bind.FilterOpts) (*MantaStakingMiddlewareOperatorUnpausedIterator, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "OperatorUnpaused")
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareOperatorUnpausedIterator{contract: _MantaStakingMiddleware.contract, event: "OperatorUnpaused", logs: logs, sub: sub}, nil
}

// WatchOperatorUnpaused is a free log subscription operation binding the contract event 0xae02c1bd695006b6d891af37fdeefea45a10ebcc17071e3471787db4f1772885.
//
// Solidity: event OperatorUnpaused(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchOperatorUnpaused(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareOperatorUnpaused) (event.Subscription, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "OperatorUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareOperatorUnpaused)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorUnpaused is a log parse operation binding the contract event 0xae02c1bd695006b6d891af37fdeefea45a10ebcc17071e3471787db4f1772885.
//
// Solidity: event OperatorUnpaused(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseOperatorUnpaused(log types.Log) (*MantaStakingMiddlewareOperatorUnpaused, error) {
	event := new(MantaStakingMiddlewareOperatorUnpaused)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareOperatorUnregisteredIterator is returned from FilterOperatorUnregistered and is used to iterate over the raw logs and unpacked data for OperatorUnregistered events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorUnregisteredIterator struct {
	Event *MantaStakingMiddlewareOperatorUnregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareOperatorUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareOperatorUnregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareOperatorUnregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareOperatorUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareOperatorUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareOperatorUnregistered represents a OperatorUnregistered event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareOperatorUnregistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorUnregistered is a free log retrieval operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterOperatorUnregistered(opts *bind.FilterOpts) (*MantaStakingMiddlewareOperatorUnregisteredIterator, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "OperatorUnregistered")
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareOperatorUnregisteredIterator{contract: _MantaStakingMiddleware.contract, event: "OperatorUnregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorUnregistered is a free log subscription operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchOperatorUnregistered(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareOperatorUnregistered) (event.Subscription, error) {

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "OperatorUnregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareOperatorUnregistered)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorUnregistered is a log parse operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address operator)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseOperatorUnregistered(log types.Log) (*MantaStakingMiddlewareOperatorUnregistered, error) {
	event := new(MantaStakingMiddlewareOperatorUnregistered)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareRoleAdminChangedIterator struct {
	Event *MantaStakingMiddlewareRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareRoleAdminChanged represents a RoleAdminChanged event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MantaStakingMiddlewareRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareRoleAdminChangedIterator{contract: _MantaStakingMiddleware.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareRoleAdminChanged)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseRoleAdminChanged(log types.Log) (*MantaStakingMiddlewareRoleAdminChanged, error) {
	event := new(MantaStakingMiddlewareRoleAdminChanged)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareRoleGrantedIterator struct {
	Event *MantaStakingMiddlewareRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareRoleGranted represents a RoleGranted event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MantaStakingMiddlewareRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareRoleGrantedIterator{contract: _MantaStakingMiddleware.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareRoleGranted)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseRoleGranted(log types.Log) (*MantaStakingMiddlewareRoleGranted, error) {
	event := new(MantaStakingMiddlewareRoleGranted)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaStakingMiddlewareRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareRoleRevokedIterator struct {
	Event *MantaStakingMiddlewareRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MantaStakingMiddlewareRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaStakingMiddlewareRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MantaStakingMiddlewareRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MantaStakingMiddlewareRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaStakingMiddlewareRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaStakingMiddlewareRoleRevoked represents a RoleRevoked event raised by the MantaStakingMiddleware contract.
type MantaStakingMiddlewareRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MantaStakingMiddlewareRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MantaStakingMiddleware.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MantaStakingMiddlewareRoleRevokedIterator{contract: _MantaStakingMiddleware.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MantaStakingMiddlewareRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MantaStakingMiddleware.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaStakingMiddlewareRoleRevoked)
				if err := _MantaStakingMiddleware.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaStakingMiddleware *MantaStakingMiddlewareFilterer) ParseRoleRevoked(log types.Log) (*MantaStakingMiddlewareRoleRevoked, error) {
	event := new(MantaStakingMiddlewareRoleRevoked)
	if err := _MantaStakingMiddleware.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
