// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fullWithdrawals

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

// FullWithdrawalsMetaData contains all meta data concerning the FullWithdrawals contract.
var FullWithdrawalsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"LogFullWithdrawalRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FREEZE_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FORCED_ACTIONS_REQS_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_VERIFIER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNFREEZE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_REMOVAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"freezeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"fullWithdrawalRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"}],\"name\":\"getEthKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getFullWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalConfigCode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBatchId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrderRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrderTreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupTreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupVaultRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidiumTreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidiumVaultRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FullWithdrawalsABI is the input ABI used to generate the binding from.
// Deprecated: Use FullWithdrawalsMetaData.ABI instead.
var FullWithdrawalsABI = FullWithdrawalsMetaData.ABI

// FullWithdrawals is an auto generated Go binding around an Ethereum contract.
type FullWithdrawals struct {
	FullWithdrawalsCaller     // Read-only binding to the contract
	FullWithdrawalsTransactor // Write-only binding to the contract
	FullWithdrawalsFilterer   // Log filterer for contract events
}

// FullWithdrawalsCaller is an auto generated read-only Go binding around an Ethereum contract.
type FullWithdrawalsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FullWithdrawalsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FullWithdrawalsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FullWithdrawalsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FullWithdrawalsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FullWithdrawalsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FullWithdrawalsSession struct {
	Contract     *FullWithdrawals  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FullWithdrawalsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FullWithdrawalsCallerSession struct {
	Contract *FullWithdrawalsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FullWithdrawalsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FullWithdrawalsTransactorSession struct {
	Contract     *FullWithdrawalsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FullWithdrawalsRaw is an auto generated low-level Go binding around an Ethereum contract.
type FullWithdrawalsRaw struct {
	Contract *FullWithdrawals // Generic contract binding to access the raw methods on
}

// FullWithdrawalsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FullWithdrawalsCallerRaw struct {
	Contract *FullWithdrawalsCaller // Generic read-only contract binding to access the raw methods on
}

// FullWithdrawalsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FullWithdrawalsTransactorRaw struct {
	Contract *FullWithdrawalsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFullWithdrawals creates a new instance of FullWithdrawals, bound to a specific deployed contract.
func NewFullWithdrawals(address common.Address, backend bind.ContractBackend) (*FullWithdrawals, error) {
	contract, err := bindFullWithdrawals(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FullWithdrawals{FullWithdrawalsCaller: FullWithdrawalsCaller{contract: contract}, FullWithdrawalsTransactor: FullWithdrawalsTransactor{contract: contract}, FullWithdrawalsFilterer: FullWithdrawalsFilterer{contract: contract}}, nil
}

// NewFullWithdrawalsCaller creates a new read-only instance of FullWithdrawals, bound to a specific deployed contract.
func NewFullWithdrawalsCaller(address common.Address, caller bind.ContractCaller) (*FullWithdrawalsCaller, error) {
	contract, err := bindFullWithdrawals(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FullWithdrawalsCaller{contract: contract}, nil
}

// NewFullWithdrawalsTransactor creates a new write-only instance of FullWithdrawals, bound to a specific deployed contract.
func NewFullWithdrawalsTransactor(address common.Address, transactor bind.ContractTransactor) (*FullWithdrawalsTransactor, error) {
	contract, err := bindFullWithdrawals(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FullWithdrawalsTransactor{contract: contract}, nil
}

// NewFullWithdrawalsFilterer creates a new log filterer instance of FullWithdrawals, bound to a specific deployed contract.
func NewFullWithdrawalsFilterer(address common.Address, filterer bind.ContractFilterer) (*FullWithdrawalsFilterer, error) {
	contract, err := bindFullWithdrawals(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FullWithdrawalsFilterer{contract: contract}, nil
}

// bindFullWithdrawals binds a generic wrapper to an already deployed contract.
func bindFullWithdrawals(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FullWithdrawalsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FullWithdrawals *FullWithdrawalsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FullWithdrawals.Contract.FullWithdrawalsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FullWithdrawals *FullWithdrawalsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FullWithdrawals.Contract.FullWithdrawalsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FullWithdrawals *FullWithdrawalsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FullWithdrawals.Contract.FullWithdrawalsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FullWithdrawals *FullWithdrawalsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FullWithdrawals.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FullWithdrawals *FullWithdrawalsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FullWithdrawals.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FullWithdrawals *FullWithdrawalsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FullWithdrawals.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) DEPOSITCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "DEPOSIT_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _FullWithdrawals.Contract.DEPOSITCANCELDELAY(&_FullWithdrawals.CallOpts)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _FullWithdrawals.Contract.DEPOSITCANCELDELAY(&_FullWithdrawals.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) FREEZEGRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "FREEZE_GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _FullWithdrawals.Contract.FREEZEGRACEPERIOD(&_FullWithdrawals.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _FullWithdrawals.Contract.FREEZEGRACEPERIOD(&_FullWithdrawals.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) MAXFORCEDACTIONSREQSPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "MAX_FORCED_ACTIONS_REQS_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _FullWithdrawals.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_FullWithdrawals.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _FullWithdrawals.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_FullWithdrawals.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) MAXVERIFIERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "MAX_VERIFIER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _FullWithdrawals.Contract.MAXVERIFIERCOUNT(&_FullWithdrawals.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _FullWithdrawals.Contract.MAXVERIFIERCOUNT(&_FullWithdrawals.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) UNFREEZEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "UNFREEZE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) UNFREEZEDELAY() (*big.Int, error) {
	return _FullWithdrawals.Contract.UNFREEZEDELAY(&_FullWithdrawals.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) UNFREEZEDELAY() (*big.Int, error) {
	return _FullWithdrawals.Contract.UNFREEZEDELAY(&_FullWithdrawals.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) VERIFIERREMOVALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "VERIFIER_REMOVAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _FullWithdrawals.Contract.VERIFIERREMOVALDELAY(&_FullWithdrawals.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _FullWithdrawals.Contract.VERIFIERREMOVALDELAY(&_FullWithdrawals.CallOpts)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_FullWithdrawals *FullWithdrawalsCaller) GetEthKey(opts *bind.CallOpts, ownerKey *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "getEthKey", ownerKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_FullWithdrawals *FullWithdrawalsSession) GetEthKey(ownerKey *big.Int) (common.Address, error) {
	return _FullWithdrawals.Contract.GetEthKey(&_FullWithdrawals.CallOpts, ownerKey)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_FullWithdrawals *FullWithdrawalsCallerSession) GetEthKey(ownerKey *big.Int) (common.Address, error) {
	return _FullWithdrawals.Contract.GetEthKey(&_FullWithdrawals.CallOpts, ownerKey)
}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) GetFullWithdrawalRequest(opts *bind.CallOpts, ownerKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "getFullWithdrawalRequest", ownerKey, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) GetFullWithdrawalRequest(ownerKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _FullWithdrawals.Contract.GetFullWithdrawalRequest(&_FullWithdrawals.CallOpts, ownerKey, vaultId)
}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) GetFullWithdrawalRequest(ownerKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _FullWithdrawals.Contract.GetFullWithdrawalRequest(&_FullWithdrawals.CallOpts, ownerKey, vaultId)
}

// GetGlobalConfigCode is a free data retrieval call binding the contract method 0xe9aa2d6b.
//
// Solidity: function getGlobalConfigCode() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) GetGlobalConfigCode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "getGlobalConfigCode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGlobalConfigCode is a free data retrieval call binding the contract method 0xe9aa2d6b.
//
// Solidity: function getGlobalConfigCode() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) GetGlobalConfigCode() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetGlobalConfigCode(&_FullWithdrawals.CallOpts)
}

// GetGlobalConfigCode is a free data retrieval call binding the contract method 0xe9aa2d6b.
//
// Solidity: function getGlobalConfigCode() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) GetGlobalConfigCode() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetGlobalConfigCode(&_FullWithdrawals.CallOpts)
}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) GetLastBatchId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "getLastBatchId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) GetLastBatchId() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetLastBatchId(&_FullWithdrawals.CallOpts)
}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) GetLastBatchId() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetLastBatchId(&_FullWithdrawals.CallOpts)
}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) GetOrderRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "getOrderRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) GetOrderRoot() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetOrderRoot(&_FullWithdrawals.CallOpts)
}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) GetOrderRoot() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetOrderRoot(&_FullWithdrawals.CallOpts)
}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) GetOrderTreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "getOrderTreeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) GetOrderTreeHeight() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetOrderTreeHeight(&_FullWithdrawals.CallOpts)
}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) GetOrderTreeHeight() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetOrderTreeHeight(&_FullWithdrawals.CallOpts)
}

// GetRollupTreeHeight is a free data retrieval call binding the contract method 0x81b47796.
//
// Solidity: function getRollupTreeHeight() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) GetRollupTreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "getRollupTreeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRollupTreeHeight is a free data retrieval call binding the contract method 0x81b47796.
//
// Solidity: function getRollupTreeHeight() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) GetRollupTreeHeight() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetRollupTreeHeight(&_FullWithdrawals.CallOpts)
}

// GetRollupTreeHeight is a free data retrieval call binding the contract method 0x81b47796.
//
// Solidity: function getRollupTreeHeight() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) GetRollupTreeHeight() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetRollupTreeHeight(&_FullWithdrawals.CallOpts)
}

// GetRollupVaultRoot is a free data retrieval call binding the contract method 0x8ed31439.
//
// Solidity: function getRollupVaultRoot() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) GetRollupVaultRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "getRollupVaultRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRollupVaultRoot is a free data retrieval call binding the contract method 0x8ed31439.
//
// Solidity: function getRollupVaultRoot() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) GetRollupVaultRoot() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetRollupVaultRoot(&_FullWithdrawals.CallOpts)
}

// GetRollupVaultRoot is a free data retrieval call binding the contract method 0x8ed31439.
//
// Solidity: function getRollupVaultRoot() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) GetRollupVaultRoot() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetRollupVaultRoot(&_FullWithdrawals.CallOpts)
}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) GetSequenceNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "getSequenceNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) GetSequenceNumber() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetSequenceNumber(&_FullWithdrawals.CallOpts)
}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) GetSequenceNumber() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetSequenceNumber(&_FullWithdrawals.CallOpts)
}

// GetValidiumTreeHeight is a free data retrieval call binding the contract method 0xbb7c3d32.
//
// Solidity: function getValidiumTreeHeight() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) GetValidiumTreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "getValidiumTreeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidiumTreeHeight is a free data retrieval call binding the contract method 0xbb7c3d32.
//
// Solidity: function getValidiumTreeHeight() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) GetValidiumTreeHeight() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetValidiumTreeHeight(&_FullWithdrawals.CallOpts)
}

// GetValidiumTreeHeight is a free data retrieval call binding the contract method 0xbb7c3d32.
//
// Solidity: function getValidiumTreeHeight() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) GetValidiumTreeHeight() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetValidiumTreeHeight(&_FullWithdrawals.CallOpts)
}

// GetValidiumVaultRoot is a free data retrieval call binding the contract method 0x3a8cc4fc.
//
// Solidity: function getValidiumVaultRoot() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCaller) GetValidiumVaultRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "getValidiumVaultRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidiumVaultRoot is a free data retrieval call binding the contract method 0x3a8cc4fc.
//
// Solidity: function getValidiumVaultRoot() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsSession) GetValidiumVaultRoot() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetValidiumVaultRoot(&_FullWithdrawals.CallOpts)
}

// GetValidiumVaultRoot is a free data retrieval call binding the contract method 0x3a8cc4fc.
//
// Solidity: function getValidiumVaultRoot() view returns(uint256)
func (_FullWithdrawals *FullWithdrawalsCallerSession) GetValidiumVaultRoot() (*big.Int, error) {
	return _FullWithdrawals.Contract.GetValidiumVaultRoot(&_FullWithdrawals.CallOpts)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_FullWithdrawals *FullWithdrawalsCaller) IsFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FullWithdrawals.contract.Call(opts, &out, "isFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_FullWithdrawals *FullWithdrawalsSession) IsFrozen() (bool, error) {
	return _FullWithdrawals.Contract.IsFrozen(&_FullWithdrawals.CallOpts)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_FullWithdrawals *FullWithdrawalsCallerSession) IsFrozen() (bool, error) {
	return _FullWithdrawals.Contract.IsFrozen(&_FullWithdrawals.CallOpts)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x93c1e466.
//
// Solidity: function freezeRequest(uint256 ownerKey, uint256 vaultId) returns()
func (_FullWithdrawals *FullWithdrawalsTransactor) FreezeRequest(opts *bind.TransactOpts, ownerKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _FullWithdrawals.contract.Transact(opts, "freezeRequest", ownerKey, vaultId)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x93c1e466.
//
// Solidity: function freezeRequest(uint256 ownerKey, uint256 vaultId) returns()
func (_FullWithdrawals *FullWithdrawalsSession) FreezeRequest(ownerKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _FullWithdrawals.Contract.FreezeRequest(&_FullWithdrawals.TransactOpts, ownerKey, vaultId)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x93c1e466.
//
// Solidity: function freezeRequest(uint256 ownerKey, uint256 vaultId) returns()
func (_FullWithdrawals *FullWithdrawalsTransactorSession) FreezeRequest(ownerKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _FullWithdrawals.Contract.FreezeRequest(&_FullWithdrawals.TransactOpts, ownerKey, vaultId)
}

// FullWithdrawalRequest is a paid mutator transaction binding the contract method 0xa93310c4.
//
// Solidity: function fullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) returns()
func (_FullWithdrawals *FullWithdrawalsTransactor) FullWithdrawalRequest(opts *bind.TransactOpts, ownerKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _FullWithdrawals.contract.Transact(opts, "fullWithdrawalRequest", ownerKey, vaultId)
}

// FullWithdrawalRequest is a paid mutator transaction binding the contract method 0xa93310c4.
//
// Solidity: function fullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) returns()
func (_FullWithdrawals *FullWithdrawalsSession) FullWithdrawalRequest(ownerKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _FullWithdrawals.Contract.FullWithdrawalRequest(&_FullWithdrawals.TransactOpts, ownerKey, vaultId)
}

// FullWithdrawalRequest is a paid mutator transaction binding the contract method 0xa93310c4.
//
// Solidity: function fullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) returns()
func (_FullWithdrawals *FullWithdrawalsTransactorSession) FullWithdrawalRequest(ownerKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _FullWithdrawals.Contract.FullWithdrawalRequest(&_FullWithdrawals.TransactOpts, ownerKey, vaultId)
}

// FullWithdrawalsLogFullWithdrawalRequestIterator is returned from FilterLogFullWithdrawalRequest and is used to iterate over the raw logs and unpacked data for LogFullWithdrawalRequest events raised by the FullWithdrawals contract.
type FullWithdrawalsLogFullWithdrawalRequestIterator struct {
	Event *FullWithdrawalsLogFullWithdrawalRequest // Event containing the contract specifics and raw log

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
func (it *FullWithdrawalsLogFullWithdrawalRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FullWithdrawalsLogFullWithdrawalRequest)
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
		it.Event = new(FullWithdrawalsLogFullWithdrawalRequest)
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
func (it *FullWithdrawalsLogFullWithdrawalRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FullWithdrawalsLogFullWithdrawalRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FullWithdrawalsLogFullWithdrawalRequest represents a LogFullWithdrawalRequest event raised by the FullWithdrawals contract.
type FullWithdrawalsLogFullWithdrawalRequest struct {
	OwnerKey *big.Int
	VaultId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogFullWithdrawalRequest is a free log retrieval operation binding the contract event 0x08eb46dbb87dcfe92d4846e5766802051525fba08a9b48318f5e0fe41186d298.
//
// Solidity: event LogFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId)
func (_FullWithdrawals *FullWithdrawalsFilterer) FilterLogFullWithdrawalRequest(opts *bind.FilterOpts) (*FullWithdrawalsLogFullWithdrawalRequestIterator, error) {

	logs, sub, err := _FullWithdrawals.contract.FilterLogs(opts, "LogFullWithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return &FullWithdrawalsLogFullWithdrawalRequestIterator{contract: _FullWithdrawals.contract, event: "LogFullWithdrawalRequest", logs: logs, sub: sub}, nil
}

// WatchLogFullWithdrawalRequest is a free log subscription operation binding the contract event 0x08eb46dbb87dcfe92d4846e5766802051525fba08a9b48318f5e0fe41186d298.
//
// Solidity: event LogFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId)
func (_FullWithdrawals *FullWithdrawalsFilterer) WatchLogFullWithdrawalRequest(opts *bind.WatchOpts, sink chan<- *FullWithdrawalsLogFullWithdrawalRequest) (event.Subscription, error) {

	logs, sub, err := _FullWithdrawals.contract.WatchLogs(opts, "LogFullWithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FullWithdrawalsLogFullWithdrawalRequest)
				if err := _FullWithdrawals.contract.UnpackLog(event, "LogFullWithdrawalRequest", log); err != nil {
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

// ParseLogFullWithdrawalRequest is a log parse operation binding the contract event 0x08eb46dbb87dcfe92d4846e5766802051525fba08a9b48318f5e0fe41186d298.
//
// Solidity: event LogFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId)
func (_FullWithdrawals *FullWithdrawalsFilterer) ParseLogFullWithdrawalRequest(log types.Log) (*FullWithdrawalsLogFullWithdrawalRequest, error) {
	event := new(FullWithdrawalsLogFullWithdrawalRequest)
	if err := _FullWithdrawals.contract.UnpackLog(event, "LogFullWithdrawalRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
