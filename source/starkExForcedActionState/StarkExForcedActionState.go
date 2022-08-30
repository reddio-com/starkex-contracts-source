// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package starkExForcedActionState

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
)

// StarkExForcedActionStateMetaData contains all meta data concerning the StarkExForcedActionState contract.
var StarkExForcedActionStateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DEPOSIT_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FREEZE_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FORCED_ACTIONS_REQS_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_VERIFIER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNFREEZE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_REMOVAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultVaultWithdrawalLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actionIndex\",\"type\":\"uint256\"}],\"name\":\"getActionHashByIndex\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getFullWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StarkExForcedActionStateABI is the input ABI used to generate the binding from.
// Deprecated: Use StarkExForcedActionStateMetaData.ABI instead.
var StarkExForcedActionStateABI = StarkExForcedActionStateMetaData.ABI

// StarkExForcedActionState is an auto generated Go binding around an Ethereum contract.
type StarkExForcedActionState struct {
	StarkExForcedActionStateCaller     // Read-only binding to the contract
	StarkExForcedActionStateTransactor // Write-only binding to the contract
	StarkExForcedActionStateFilterer   // Log filterer for contract events
}

// StarkExForcedActionStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type StarkExForcedActionStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarkExForcedActionStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StarkExForcedActionStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarkExForcedActionStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StarkExForcedActionStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarkExForcedActionStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StarkExForcedActionStateSession struct {
	Contract     *StarkExForcedActionState // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StarkExForcedActionStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StarkExForcedActionStateCallerSession struct {
	Contract *StarkExForcedActionStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// StarkExForcedActionStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StarkExForcedActionStateTransactorSession struct {
	Contract     *StarkExForcedActionStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// StarkExForcedActionStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type StarkExForcedActionStateRaw struct {
	Contract *StarkExForcedActionState // Generic contract binding to access the raw methods on
}

// StarkExForcedActionStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StarkExForcedActionStateCallerRaw struct {
	Contract *StarkExForcedActionStateCaller // Generic read-only contract binding to access the raw methods on
}

// StarkExForcedActionStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StarkExForcedActionStateTransactorRaw struct {
	Contract *StarkExForcedActionStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStarkExForcedActionState creates a new instance of StarkExForcedActionState, bound to a specific deployed contract.
func NewStarkExForcedActionState(address common.Address, backend bind.ContractBackend) (*StarkExForcedActionState, error) {
	contract, err := bindStarkExForcedActionState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StarkExForcedActionState{StarkExForcedActionStateCaller: StarkExForcedActionStateCaller{contract: contract}, StarkExForcedActionStateTransactor: StarkExForcedActionStateTransactor{contract: contract}, StarkExForcedActionStateFilterer: StarkExForcedActionStateFilterer{contract: contract}}, nil
}

// NewStarkExForcedActionStateCaller creates a new read-only instance of StarkExForcedActionState, bound to a specific deployed contract.
func NewStarkExForcedActionStateCaller(address common.Address, caller bind.ContractCaller) (*StarkExForcedActionStateCaller, error) {
	contract, err := bindStarkExForcedActionState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StarkExForcedActionStateCaller{contract: contract}, nil
}

// NewStarkExForcedActionStateTransactor creates a new write-only instance of StarkExForcedActionState, bound to a specific deployed contract.
func NewStarkExForcedActionStateTransactor(address common.Address, transactor bind.ContractTransactor) (*StarkExForcedActionStateTransactor, error) {
	contract, err := bindStarkExForcedActionState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StarkExForcedActionStateTransactor{contract: contract}, nil
}

// NewStarkExForcedActionStateFilterer creates a new log filterer instance of StarkExForcedActionState, bound to a specific deployed contract.
func NewStarkExForcedActionStateFilterer(address common.Address, filterer bind.ContractFilterer) (*StarkExForcedActionStateFilterer, error) {
	contract, err := bindStarkExForcedActionState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StarkExForcedActionStateFilterer{contract: contract}, nil
}

// bindStarkExForcedActionState binds a generic wrapper to an already deployed contract.
func bindStarkExForcedActionState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StarkExForcedActionStateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StarkExForcedActionState *StarkExForcedActionStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StarkExForcedActionState.Contract.StarkExForcedActionStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StarkExForcedActionState *StarkExForcedActionStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StarkExForcedActionState.Contract.StarkExForcedActionStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StarkExForcedActionState *StarkExForcedActionStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StarkExForcedActionState.Contract.StarkExForcedActionStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StarkExForcedActionState *StarkExForcedActionStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StarkExForcedActionState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StarkExForcedActionState *StarkExForcedActionStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StarkExForcedActionState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StarkExForcedActionState *StarkExForcedActionStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StarkExForcedActionState.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCaller) DEPOSITCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkExForcedActionState.contract.Call(opts, &out, "DEPOSIT_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.DEPOSITCANCELDELAY(&_StarkExForcedActionState.CallOpts)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCallerSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.DEPOSITCANCELDELAY(&_StarkExForcedActionState.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCaller) FREEZEGRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkExForcedActionState.contract.Call(opts, &out, "FREEZE_GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.FREEZEGRACEPERIOD(&_StarkExForcedActionState.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCallerSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.FREEZEGRACEPERIOD(&_StarkExForcedActionState.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCaller) MAXFORCEDACTIONSREQSPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkExForcedActionState.contract.Call(opts, &out, "MAX_FORCED_ACTIONS_REQS_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_StarkExForcedActionState.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCallerSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_StarkExForcedActionState.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCaller) MAXVERIFIERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkExForcedActionState.contract.Call(opts, &out, "MAX_VERIFIER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.MAXVERIFIERCOUNT(&_StarkExForcedActionState.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCallerSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.MAXVERIFIERCOUNT(&_StarkExForcedActionState.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCaller) UNFREEZEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkExForcedActionState.contract.Call(opts, &out, "UNFREEZE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateSession) UNFREEZEDELAY() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.UNFREEZEDELAY(&_StarkExForcedActionState.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCallerSession) UNFREEZEDELAY() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.UNFREEZEDELAY(&_StarkExForcedActionState.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCaller) VERIFIERREMOVALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkExForcedActionState.contract.Call(opts, &out, "VERIFIER_REMOVAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.VERIFIERREMOVALDELAY(&_StarkExForcedActionState.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCallerSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.VERIFIERREMOVALDELAY(&_StarkExForcedActionState.CallOpts)
}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCaller) DefaultVaultWithdrawalLock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkExForcedActionState.contract.Call(opts, &out, "defaultVaultWithdrawalLock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateSession) DefaultVaultWithdrawalLock() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.DefaultVaultWithdrawalLock(&_StarkExForcedActionState.CallOpts)
}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCallerSession) DefaultVaultWithdrawalLock() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.DefaultVaultWithdrawalLock(&_StarkExForcedActionState.CallOpts)
}

// GetActionCount is a free data retrieval call binding the contract method 0x5eecd218.
//
// Solidity: function getActionCount() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCaller) GetActionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarkExForcedActionState.contract.Call(opts, &out, "getActionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActionCount is a free data retrieval call binding the contract method 0x5eecd218.
//
// Solidity: function getActionCount() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateSession) GetActionCount() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.GetActionCount(&_StarkExForcedActionState.CallOpts)
}

// GetActionCount is a free data retrieval call binding the contract method 0x5eecd218.
//
// Solidity: function getActionCount() view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCallerSession) GetActionCount() (*big.Int, error) {
	return _StarkExForcedActionState.Contract.GetActionCount(&_StarkExForcedActionState.CallOpts)
}

// GetActionHashByIndex is a free data retrieval call binding the contract method 0x5e586cd1.
//
// Solidity: function getActionHashByIndex(uint256 actionIndex) view returns(bytes32)
func (_StarkExForcedActionState *StarkExForcedActionStateCaller) GetActionHashByIndex(opts *bind.CallOpts, actionIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _StarkExForcedActionState.contract.Call(opts, &out, "getActionHashByIndex", actionIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionHashByIndex is a free data retrieval call binding the contract method 0x5e586cd1.
//
// Solidity: function getActionHashByIndex(uint256 actionIndex) view returns(bytes32)
func (_StarkExForcedActionState *StarkExForcedActionStateSession) GetActionHashByIndex(actionIndex *big.Int) ([32]byte, error) {
	return _StarkExForcedActionState.Contract.GetActionHashByIndex(&_StarkExForcedActionState.CallOpts, actionIndex)
}

// GetActionHashByIndex is a free data retrieval call binding the contract method 0x5e586cd1.
//
// Solidity: function getActionHashByIndex(uint256 actionIndex) view returns(bytes32)
func (_StarkExForcedActionState *StarkExForcedActionStateCallerSession) GetActionHashByIndex(actionIndex *big.Int) ([32]byte, error) {
	return _StarkExForcedActionState.Contract.GetActionHashByIndex(&_StarkExForcedActionState.CallOpts, actionIndex)
}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCaller) GetFullWithdrawalRequest(opts *bind.CallOpts, ownerKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StarkExForcedActionState.contract.Call(opts, &out, "getFullWithdrawalRequest", ownerKey, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateSession) GetFullWithdrawalRequest(ownerKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _StarkExForcedActionState.Contract.GetFullWithdrawalRequest(&_StarkExForcedActionState.CallOpts, ownerKey, vaultId)
}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) view returns(uint256)
func (_StarkExForcedActionState *StarkExForcedActionStateCallerSession) GetFullWithdrawalRequest(ownerKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _StarkExForcedActionState.Contract.GetFullWithdrawalRequest(&_StarkExForcedActionState.CallOpts, ownerKey, vaultId)
}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_StarkExForcedActionState *StarkExForcedActionStateCaller) OrderRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StarkExForcedActionState.contract.Call(opts, &out, "orderRegistryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_StarkExForcedActionState *StarkExForcedActionStateSession) OrderRegistryAddress() (common.Address, error) {
	return _StarkExForcedActionState.Contract.OrderRegistryAddress(&_StarkExForcedActionState.CallOpts)
}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_StarkExForcedActionState *StarkExForcedActionStateCallerSession) OrderRegistryAddress() (common.Address, error) {
	return _StarkExForcedActionState.Contract.OrderRegistryAddress(&_StarkExForcedActionState.CallOpts)
}
