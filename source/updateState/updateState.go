// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package updateState

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

// UpdateStateMetaData contains all meta data concerning the UpdateState contract.
var UpdateStateMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"LogOperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"LogOperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validiumVaultRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rollupVaultRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderRoot\",\"type\":\"uint256\"}],\"name\":\"LogRootUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateTransitionFact\",\"type\":\"bytes32\"}],\"name\":\"LogStateTransitionFact\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"quantizedAmountChange\",\"type\":\"int256\"}],\"name\":\"LogVaultBalanceChangeApplied\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FREEZE_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FORCED_ACTIONS_REQS_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_VERIFIER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STARKEX_MAX_DEFAULT_VAULT_LOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNFREEZE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_REMOVAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultVaultWithdrawalLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getFullWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"testedOperator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"removedOperator\",\"type\":\"address\"}],\"name\":\"unregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"publicInput\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"applicationData\",\"type\":\"uint256[]\"}],\"name\":\"updateState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UpdateStateABI is the input ABI used to generate the binding from.
// Deprecated: Use UpdateStateMetaData.ABI instead.
var UpdateStateABI = UpdateStateMetaData.ABI

// UpdateState is an auto generated Go binding around an Ethereum contract.
type UpdateState struct {
	UpdateStateCaller     // Read-only binding to the contract
	UpdateStateTransactor // Write-only binding to the contract
	UpdateStateFilterer   // Log filterer for contract events
}

// UpdateStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpdateStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpdateStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpdateStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpdateStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpdateStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpdateStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpdateStateSession struct {
	Contract     *UpdateState      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpdateStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpdateStateCallerSession struct {
	Contract *UpdateStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UpdateStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpdateStateTransactorSession struct {
	Contract     *UpdateStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UpdateStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpdateStateRaw struct {
	Contract *UpdateState // Generic contract binding to access the raw methods on
}

// UpdateStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpdateStateCallerRaw struct {
	Contract *UpdateStateCaller // Generic read-only contract binding to access the raw methods on
}

// UpdateStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpdateStateTransactorRaw struct {
	Contract *UpdateStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpdateState creates a new instance of UpdateState, bound to a specific deployed contract.
func NewUpdateState(address common.Address, backend bind.ContractBackend) (*UpdateState, error) {
	contract, err := bindUpdateState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpdateState{UpdateStateCaller: UpdateStateCaller{contract: contract}, UpdateStateTransactor: UpdateStateTransactor{contract: contract}, UpdateStateFilterer: UpdateStateFilterer{contract: contract}}, nil
}

// NewUpdateStateCaller creates a new read-only instance of UpdateState, bound to a specific deployed contract.
func NewUpdateStateCaller(address common.Address, caller bind.ContractCaller) (*UpdateStateCaller, error) {
	contract, err := bindUpdateState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpdateStateCaller{contract: contract}, nil
}

// NewUpdateStateTransactor creates a new write-only instance of UpdateState, bound to a specific deployed contract.
func NewUpdateStateTransactor(address common.Address, transactor bind.ContractTransactor) (*UpdateStateTransactor, error) {
	contract, err := bindUpdateState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpdateStateTransactor{contract: contract}, nil
}

// NewUpdateStateFilterer creates a new log filterer instance of UpdateState, bound to a specific deployed contract.
func NewUpdateStateFilterer(address common.Address, filterer bind.ContractFilterer) (*UpdateStateFilterer, error) {
	contract, err := bindUpdateState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpdateStateFilterer{contract: contract}, nil
}

// bindUpdateState binds a generic wrapper to an already deployed contract.
func bindUpdateState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UpdateStateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpdateState *UpdateStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpdateState.Contract.UpdateStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpdateState *UpdateStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpdateState.Contract.UpdateStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpdateState *UpdateStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpdateState.Contract.UpdateStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpdateState *UpdateStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpdateState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpdateState *UpdateStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpdateState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpdateState *UpdateStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpdateState.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_UpdateState *UpdateStateCaller) DEPOSITCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpdateState.contract.Call(opts, &out, "DEPOSIT_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_UpdateState *UpdateStateSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _UpdateState.Contract.DEPOSITCANCELDELAY(&_UpdateState.CallOpts)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_UpdateState *UpdateStateCallerSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _UpdateState.Contract.DEPOSITCANCELDELAY(&_UpdateState.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_UpdateState *UpdateStateCaller) FREEZEGRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpdateState.contract.Call(opts, &out, "FREEZE_GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_UpdateState *UpdateStateSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _UpdateState.Contract.FREEZEGRACEPERIOD(&_UpdateState.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_UpdateState *UpdateStateCallerSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _UpdateState.Contract.FREEZEGRACEPERIOD(&_UpdateState.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_UpdateState *UpdateStateCaller) MAXFORCEDACTIONSREQSPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpdateState.contract.Call(opts, &out, "MAX_FORCED_ACTIONS_REQS_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_UpdateState *UpdateStateSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _UpdateState.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_UpdateState.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_UpdateState *UpdateStateCallerSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _UpdateState.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_UpdateState.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_UpdateState *UpdateStateCaller) MAXVERIFIERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpdateState.contract.Call(opts, &out, "MAX_VERIFIER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_UpdateState *UpdateStateSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _UpdateState.Contract.MAXVERIFIERCOUNT(&_UpdateState.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_UpdateState *UpdateStateCallerSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _UpdateState.Contract.MAXVERIFIERCOUNT(&_UpdateState.CallOpts)
}

// STARKEXMAXDEFAULTVAULTLOCK is a free data retrieval call binding the contract method 0x835a71c3.
//
// Solidity: function STARKEX_MAX_DEFAULT_VAULT_LOCK() view returns(uint256)
func (_UpdateState *UpdateStateCaller) STARKEXMAXDEFAULTVAULTLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpdateState.contract.Call(opts, &out, "STARKEX_MAX_DEFAULT_VAULT_LOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STARKEXMAXDEFAULTVAULTLOCK is a free data retrieval call binding the contract method 0x835a71c3.
//
// Solidity: function STARKEX_MAX_DEFAULT_VAULT_LOCK() view returns(uint256)
func (_UpdateState *UpdateStateSession) STARKEXMAXDEFAULTVAULTLOCK() (*big.Int, error) {
	return _UpdateState.Contract.STARKEXMAXDEFAULTVAULTLOCK(&_UpdateState.CallOpts)
}

// STARKEXMAXDEFAULTVAULTLOCK is a free data retrieval call binding the contract method 0x835a71c3.
//
// Solidity: function STARKEX_MAX_DEFAULT_VAULT_LOCK() view returns(uint256)
func (_UpdateState *UpdateStateCallerSession) STARKEXMAXDEFAULTVAULTLOCK() (*big.Int, error) {
	return _UpdateState.Contract.STARKEXMAXDEFAULTVAULTLOCK(&_UpdateState.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_UpdateState *UpdateStateCaller) UNFREEZEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpdateState.contract.Call(opts, &out, "UNFREEZE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_UpdateState *UpdateStateSession) UNFREEZEDELAY() (*big.Int, error) {
	return _UpdateState.Contract.UNFREEZEDELAY(&_UpdateState.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_UpdateState *UpdateStateCallerSession) UNFREEZEDELAY() (*big.Int, error) {
	return _UpdateState.Contract.UNFREEZEDELAY(&_UpdateState.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_UpdateState *UpdateStateCaller) VERIFIERREMOVALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpdateState.contract.Call(opts, &out, "VERIFIER_REMOVAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_UpdateState *UpdateStateSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _UpdateState.Contract.VERIFIERREMOVALDELAY(&_UpdateState.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_UpdateState *UpdateStateCallerSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _UpdateState.Contract.VERIFIERREMOVALDELAY(&_UpdateState.CallOpts)
}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_UpdateState *UpdateStateCaller) DefaultVaultWithdrawalLock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpdateState.contract.Call(opts, &out, "defaultVaultWithdrawalLock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_UpdateState *UpdateStateSession) DefaultVaultWithdrawalLock() (*big.Int, error) {
	return _UpdateState.Contract.DefaultVaultWithdrawalLock(&_UpdateState.CallOpts)
}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_UpdateState *UpdateStateCallerSession) DefaultVaultWithdrawalLock() (*big.Int, error) {
	return _UpdateState.Contract.DefaultVaultWithdrawalLock(&_UpdateState.CallOpts)
}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) view returns(uint256)
func (_UpdateState *UpdateStateCaller) GetFullWithdrawalRequest(opts *bind.CallOpts, ownerKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UpdateState.contract.Call(opts, &out, "getFullWithdrawalRequest", ownerKey, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) view returns(uint256)
func (_UpdateState *UpdateStateSession) GetFullWithdrawalRequest(ownerKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _UpdateState.Contract.GetFullWithdrawalRequest(&_UpdateState.CallOpts, ownerKey, vaultId)
}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) view returns(uint256)
func (_UpdateState *UpdateStateCallerSession) GetFullWithdrawalRequest(ownerKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _UpdateState.Contract.GetFullWithdrawalRequest(&_UpdateState.CallOpts, ownerKey, vaultId)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_UpdateState *UpdateStateCaller) IsFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UpdateState.contract.Call(opts, &out, "isFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_UpdateState *UpdateStateSession) IsFrozen() (bool, error) {
	return _UpdateState.Contract.IsFrozen(&_UpdateState.CallOpts)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_UpdateState *UpdateStateCallerSession) IsFrozen() (bool, error) {
	return _UpdateState.Contract.IsFrozen(&_UpdateState.CallOpts)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_UpdateState *UpdateStateCaller) IsOperator(opts *bind.CallOpts, testedOperator common.Address) (bool, error) {
	var out []interface{}
	err := _UpdateState.contract.Call(opts, &out, "isOperator", testedOperator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_UpdateState *UpdateStateSession) IsOperator(testedOperator common.Address) (bool, error) {
	return _UpdateState.Contract.IsOperator(&_UpdateState.CallOpts, testedOperator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_UpdateState *UpdateStateCallerSession) IsOperator(testedOperator common.Address) (bool, error) {
	return _UpdateState.Contract.IsOperator(&_UpdateState.CallOpts, testedOperator)
}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_UpdateState *UpdateStateCaller) OrderRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpdateState.contract.Call(opts, &out, "orderRegistryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_UpdateState *UpdateStateSession) OrderRegistryAddress() (common.Address, error) {
	return _UpdateState.Contract.OrderRegistryAddress(&_UpdateState.CallOpts)
}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_UpdateState *UpdateStateCallerSession) OrderRegistryAddress() (common.Address, error) {
	return _UpdateState.Contract.OrderRegistryAddress(&_UpdateState.CallOpts)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_UpdateState *UpdateStateTransactor) RegisterOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _UpdateState.contract.Transact(opts, "registerOperator", newOperator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_UpdateState *UpdateStateSession) RegisterOperator(newOperator common.Address) (*types.Transaction, error) {
	return _UpdateState.Contract.RegisterOperator(&_UpdateState.TransactOpts, newOperator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_UpdateState *UpdateStateTransactorSession) RegisterOperator(newOperator common.Address) (*types.Transaction, error) {
	return _UpdateState.Contract.RegisterOperator(&_UpdateState.TransactOpts, newOperator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_UpdateState *UpdateStateTransactor) UnregisterOperator(opts *bind.TransactOpts, removedOperator common.Address) (*types.Transaction, error) {
	return _UpdateState.contract.Transact(opts, "unregisterOperator", removedOperator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_UpdateState *UpdateStateSession) UnregisterOperator(removedOperator common.Address) (*types.Transaction, error) {
	return _UpdateState.Contract.UnregisterOperator(&_UpdateState.TransactOpts, removedOperator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_UpdateState *UpdateStateTransactorSession) UnregisterOperator(removedOperator common.Address) (*types.Transaction, error) {
	return _UpdateState.Contract.UnregisterOperator(&_UpdateState.TransactOpts, removedOperator)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] publicInput, uint256[] applicationData) returns()
func (_UpdateState *UpdateStateTransactor) UpdateState(opts *bind.TransactOpts, publicInput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _UpdateState.contract.Transact(opts, "updateState", publicInput, applicationData)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] publicInput, uint256[] applicationData) returns()
func (_UpdateState *UpdateStateSession) UpdateState(publicInput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _UpdateState.Contract.UpdateState(&_UpdateState.TransactOpts, publicInput, applicationData)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] publicInput, uint256[] applicationData) returns()
func (_UpdateState *UpdateStateTransactorSession) UpdateState(publicInput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _UpdateState.Contract.UpdateState(&_UpdateState.TransactOpts, publicInput, applicationData)
}

// UpdateStateLogOperatorAddedIterator is returned from FilterLogOperatorAdded and is used to iterate over the raw logs and unpacked data for LogOperatorAdded events raised by the UpdateState contract.
type UpdateStateLogOperatorAddedIterator struct {
	Event *UpdateStateLogOperatorAdded // Event containing the contract specifics and raw log

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
func (it *UpdateStateLogOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpdateStateLogOperatorAdded)
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
		it.Event = new(UpdateStateLogOperatorAdded)
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
func (it *UpdateStateLogOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpdateStateLogOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpdateStateLogOperatorAdded represents a LogOperatorAdded event raised by the UpdateState contract.
type UpdateStateLogOperatorAdded struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogOperatorAdded is a free log retrieval operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_UpdateState *UpdateStateFilterer) FilterLogOperatorAdded(opts *bind.FilterOpts) (*UpdateStateLogOperatorAddedIterator, error) {

	logs, sub, err := _UpdateState.contract.FilterLogs(opts, "LogOperatorAdded")
	if err != nil {
		return nil, err
	}
	return &UpdateStateLogOperatorAddedIterator{contract: _UpdateState.contract, event: "LogOperatorAdded", logs: logs, sub: sub}, nil
}

// WatchLogOperatorAdded is a free log subscription operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_UpdateState *UpdateStateFilterer) WatchLogOperatorAdded(opts *bind.WatchOpts, sink chan<- *UpdateStateLogOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _UpdateState.contract.WatchLogs(opts, "LogOperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpdateStateLogOperatorAdded)
				if err := _UpdateState.contract.UnpackLog(event, "LogOperatorAdded", log); err != nil {
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

// ParseLogOperatorAdded is a log parse operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_UpdateState *UpdateStateFilterer) ParseLogOperatorAdded(log types.Log) (*UpdateStateLogOperatorAdded, error) {
	event := new(UpdateStateLogOperatorAdded)
	if err := _UpdateState.contract.UnpackLog(event, "LogOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpdateStateLogOperatorRemovedIterator is returned from FilterLogOperatorRemoved and is used to iterate over the raw logs and unpacked data for LogOperatorRemoved events raised by the UpdateState contract.
type UpdateStateLogOperatorRemovedIterator struct {
	Event *UpdateStateLogOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *UpdateStateLogOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpdateStateLogOperatorRemoved)
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
		it.Event = new(UpdateStateLogOperatorRemoved)
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
func (it *UpdateStateLogOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpdateStateLogOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpdateStateLogOperatorRemoved represents a LogOperatorRemoved event raised by the UpdateState contract.
type UpdateStateLogOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogOperatorRemoved is a free log retrieval operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_UpdateState *UpdateStateFilterer) FilterLogOperatorRemoved(opts *bind.FilterOpts) (*UpdateStateLogOperatorRemovedIterator, error) {

	logs, sub, err := _UpdateState.contract.FilterLogs(opts, "LogOperatorRemoved")
	if err != nil {
		return nil, err
	}
	return &UpdateStateLogOperatorRemovedIterator{contract: _UpdateState.contract, event: "LogOperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchLogOperatorRemoved is a free log subscription operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_UpdateState *UpdateStateFilterer) WatchLogOperatorRemoved(opts *bind.WatchOpts, sink chan<- *UpdateStateLogOperatorRemoved) (event.Subscription, error) {

	logs, sub, err := _UpdateState.contract.WatchLogs(opts, "LogOperatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpdateStateLogOperatorRemoved)
				if err := _UpdateState.contract.UnpackLog(event, "LogOperatorRemoved", log); err != nil {
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

// ParseLogOperatorRemoved is a log parse operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_UpdateState *UpdateStateFilterer) ParseLogOperatorRemoved(log types.Log) (*UpdateStateLogOperatorRemoved, error) {
	event := new(UpdateStateLogOperatorRemoved)
	if err := _UpdateState.contract.UnpackLog(event, "LogOperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpdateStateLogRootUpdateIterator is returned from FilterLogRootUpdate and is used to iterate over the raw logs and unpacked data for LogRootUpdate events raised by the UpdateState contract.
type UpdateStateLogRootUpdateIterator struct {
	Event *UpdateStateLogRootUpdate // Event containing the contract specifics and raw log

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
func (it *UpdateStateLogRootUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpdateStateLogRootUpdate)
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
		it.Event = new(UpdateStateLogRootUpdate)
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
func (it *UpdateStateLogRootUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpdateStateLogRootUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpdateStateLogRootUpdate represents a LogRootUpdate event raised by the UpdateState contract.
type UpdateStateLogRootUpdate struct {
	SequenceNumber    *big.Int
	BatchId           *big.Int
	ValidiumVaultRoot *big.Int
	RollupVaultRoot   *big.Int
	OrderRoot         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogRootUpdate is a free log retrieval operation binding the contract event 0x54fe7a67f8957a96919a0d1b81eeb25ea8c06f96ad528671da17a4a840040664.
//
// Solidity: event LogRootUpdate(uint256 sequenceNumber, uint256 batchId, uint256 validiumVaultRoot, uint256 rollupVaultRoot, uint256 orderRoot)
func (_UpdateState *UpdateStateFilterer) FilterLogRootUpdate(opts *bind.FilterOpts) (*UpdateStateLogRootUpdateIterator, error) {

	logs, sub, err := _UpdateState.contract.FilterLogs(opts, "LogRootUpdate")
	if err != nil {
		return nil, err
	}
	return &UpdateStateLogRootUpdateIterator{contract: _UpdateState.contract, event: "LogRootUpdate", logs: logs, sub: sub}, nil
}

// WatchLogRootUpdate is a free log subscription operation binding the contract event 0x54fe7a67f8957a96919a0d1b81eeb25ea8c06f96ad528671da17a4a840040664.
//
// Solidity: event LogRootUpdate(uint256 sequenceNumber, uint256 batchId, uint256 validiumVaultRoot, uint256 rollupVaultRoot, uint256 orderRoot)
func (_UpdateState *UpdateStateFilterer) WatchLogRootUpdate(opts *bind.WatchOpts, sink chan<- *UpdateStateLogRootUpdate) (event.Subscription, error) {

	logs, sub, err := _UpdateState.contract.WatchLogs(opts, "LogRootUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpdateStateLogRootUpdate)
				if err := _UpdateState.contract.UnpackLog(event, "LogRootUpdate", log); err != nil {
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

// ParseLogRootUpdate is a log parse operation binding the contract event 0x54fe7a67f8957a96919a0d1b81eeb25ea8c06f96ad528671da17a4a840040664.
//
// Solidity: event LogRootUpdate(uint256 sequenceNumber, uint256 batchId, uint256 validiumVaultRoot, uint256 rollupVaultRoot, uint256 orderRoot)
func (_UpdateState *UpdateStateFilterer) ParseLogRootUpdate(log types.Log) (*UpdateStateLogRootUpdate, error) {
	event := new(UpdateStateLogRootUpdate)
	if err := _UpdateState.contract.UnpackLog(event, "LogRootUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpdateStateLogStateTransitionFactIterator is returned from FilterLogStateTransitionFact and is used to iterate over the raw logs and unpacked data for LogStateTransitionFact events raised by the UpdateState contract.
type UpdateStateLogStateTransitionFactIterator struct {
	Event *UpdateStateLogStateTransitionFact // Event containing the contract specifics and raw log

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
func (it *UpdateStateLogStateTransitionFactIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpdateStateLogStateTransitionFact)
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
		it.Event = new(UpdateStateLogStateTransitionFact)
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
func (it *UpdateStateLogStateTransitionFactIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpdateStateLogStateTransitionFactIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpdateStateLogStateTransitionFact represents a LogStateTransitionFact event raised by the UpdateState contract.
type UpdateStateLogStateTransitionFact struct {
	StateTransitionFact [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogStateTransitionFact is a free log retrieval operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_UpdateState *UpdateStateFilterer) FilterLogStateTransitionFact(opts *bind.FilterOpts) (*UpdateStateLogStateTransitionFactIterator, error) {

	logs, sub, err := _UpdateState.contract.FilterLogs(opts, "LogStateTransitionFact")
	if err != nil {
		return nil, err
	}
	return &UpdateStateLogStateTransitionFactIterator{contract: _UpdateState.contract, event: "LogStateTransitionFact", logs: logs, sub: sub}, nil
}

// WatchLogStateTransitionFact is a free log subscription operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_UpdateState *UpdateStateFilterer) WatchLogStateTransitionFact(opts *bind.WatchOpts, sink chan<- *UpdateStateLogStateTransitionFact) (event.Subscription, error) {

	logs, sub, err := _UpdateState.contract.WatchLogs(opts, "LogStateTransitionFact")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpdateStateLogStateTransitionFact)
				if err := _UpdateState.contract.UnpackLog(event, "LogStateTransitionFact", log); err != nil {
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

// ParseLogStateTransitionFact is a log parse operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_UpdateState *UpdateStateFilterer) ParseLogStateTransitionFact(log types.Log) (*UpdateStateLogStateTransitionFact, error) {
	event := new(UpdateStateLogStateTransitionFact)
	if err := _UpdateState.contract.UnpackLog(event, "LogStateTransitionFact", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpdateStateLogVaultBalanceChangeAppliedIterator is returned from FilterLogVaultBalanceChangeApplied and is used to iterate over the raw logs and unpacked data for LogVaultBalanceChangeApplied events raised by the UpdateState contract.
type UpdateStateLogVaultBalanceChangeAppliedIterator struct {
	Event *UpdateStateLogVaultBalanceChangeApplied // Event containing the contract specifics and raw log

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
func (it *UpdateStateLogVaultBalanceChangeAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpdateStateLogVaultBalanceChangeApplied)
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
		it.Event = new(UpdateStateLogVaultBalanceChangeApplied)
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
func (it *UpdateStateLogVaultBalanceChangeAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpdateStateLogVaultBalanceChangeAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpdateStateLogVaultBalanceChangeApplied represents a LogVaultBalanceChangeApplied event raised by the UpdateState contract.
type UpdateStateLogVaultBalanceChangeApplied struct {
	EthKey                common.Address
	AssetId               *big.Int
	VaultId               *big.Int
	QuantizedAmountChange *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterLogVaultBalanceChangeApplied is a free log retrieval operation binding the contract event 0x2b2b583bb5166d03b87a8e7c2785e61530ab776400fb69e1bcb13b33afef2b58.
//
// Solidity: event LogVaultBalanceChangeApplied(address ethKey, uint256 assetId, uint256 vaultId, int256 quantizedAmountChange)
func (_UpdateState *UpdateStateFilterer) FilterLogVaultBalanceChangeApplied(opts *bind.FilterOpts) (*UpdateStateLogVaultBalanceChangeAppliedIterator, error) {

	logs, sub, err := _UpdateState.contract.FilterLogs(opts, "LogVaultBalanceChangeApplied")
	if err != nil {
		return nil, err
	}
	return &UpdateStateLogVaultBalanceChangeAppliedIterator{contract: _UpdateState.contract, event: "LogVaultBalanceChangeApplied", logs: logs, sub: sub}, nil
}

// WatchLogVaultBalanceChangeApplied is a free log subscription operation binding the contract event 0x2b2b583bb5166d03b87a8e7c2785e61530ab776400fb69e1bcb13b33afef2b58.
//
// Solidity: event LogVaultBalanceChangeApplied(address ethKey, uint256 assetId, uint256 vaultId, int256 quantizedAmountChange)
func (_UpdateState *UpdateStateFilterer) WatchLogVaultBalanceChangeApplied(opts *bind.WatchOpts, sink chan<- *UpdateStateLogVaultBalanceChangeApplied) (event.Subscription, error) {

	logs, sub, err := _UpdateState.contract.WatchLogs(opts, "LogVaultBalanceChangeApplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpdateStateLogVaultBalanceChangeApplied)
				if err := _UpdateState.contract.UnpackLog(event, "LogVaultBalanceChangeApplied", log); err != nil {
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

// ParseLogVaultBalanceChangeApplied is a log parse operation binding the contract event 0x2b2b583bb5166d03b87a8e7c2785e61530ab776400fb69e1bcb13b33afef2b58.
//
// Solidity: event LogVaultBalanceChangeApplied(address ethKey, uint256 assetId, uint256 vaultId, int256 quantizedAmountChange)
func (_UpdateState *UpdateStateFilterer) ParseLogVaultBalanceChangeApplied(log types.Log) (*UpdateStateLogVaultBalanceChangeApplied, error) {
	event := new(UpdateStateLogVaultBalanceChangeApplied)
	if err := _UpdateState.contract.UnpackLog(event, "LogVaultBalanceChangeApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
