// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deposits

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

// DepositsMetaData contains all meta data concerning the Deposits contract.
var DepositsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositorEthKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogDepositCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDepositCancelReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogDepositNftCancelReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositorEthKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogNftDeposit\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FREEZE_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FORCED_ACTIONS_REQS_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_VERIFIER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNFREEZE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_REMOVAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositCancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositNftReclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositReclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"getAssetInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getCancellationRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"request\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"}],\"name\":\"getEthKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getQuantizedDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"presumedAssetType\",\"type\":\"uint256\"}],\"name\":\"getQuantum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DepositsABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositsMetaData.ABI instead.
var DepositsABI = DepositsMetaData.ABI

// Deposits is an auto generated Go binding around an Ethereum contract.
type Deposits struct {
	DepositsCaller     // Read-only binding to the contract
	DepositsTransactor // Write-only binding to the contract
	DepositsFilterer   // Log filterer for contract events
}

// DepositsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositsSession struct {
	Contract     *Deposits         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositsCallerSession struct {
	Contract *DepositsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DepositsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositsTransactorSession struct {
	Contract     *DepositsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DepositsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositsRaw struct {
	Contract *Deposits // Generic contract binding to access the raw methods on
}

// DepositsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositsCallerRaw struct {
	Contract *DepositsCaller // Generic read-only contract binding to access the raw methods on
}

// DepositsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositsTransactorRaw struct {
	Contract *DepositsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeposits creates a new instance of Deposits, bound to a specific deployed contract.
func NewDeposits(address common.Address, backend bind.ContractBackend) (*Deposits, error) {
	contract, err := bindDeposits(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deposits{DepositsCaller: DepositsCaller{contract: contract}, DepositsTransactor: DepositsTransactor{contract: contract}, DepositsFilterer: DepositsFilterer{contract: contract}}, nil
}

// NewDepositsCaller creates a new read-only instance of Deposits, bound to a specific deployed contract.
func NewDepositsCaller(address common.Address, caller bind.ContractCaller) (*DepositsCaller, error) {
	contract, err := bindDeposits(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositsCaller{contract: contract}, nil
}

// NewDepositsTransactor creates a new write-only instance of Deposits, bound to a specific deployed contract.
func NewDepositsTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositsTransactor, error) {
	contract, err := bindDeposits(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositsTransactor{contract: contract}, nil
}

// NewDepositsFilterer creates a new log filterer instance of Deposits, bound to a specific deployed contract.
func NewDepositsFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositsFilterer, error) {
	contract, err := bindDeposits(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositsFilterer{contract: contract}, nil
}

// bindDeposits binds a generic wrapper to an already deployed contract.
func bindDeposits(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deposits *DepositsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deposits.Contract.DepositsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deposits *DepositsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposits.Contract.DepositsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deposits *DepositsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deposits.Contract.DepositsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deposits *DepositsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deposits.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deposits *DepositsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposits.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deposits *DepositsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deposits.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_Deposits *DepositsCaller) DEPOSITCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deposits.contract.Call(opts, &out, "DEPOSIT_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_Deposits *DepositsSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _Deposits.Contract.DEPOSITCANCELDELAY(&_Deposits.CallOpts)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_Deposits *DepositsCallerSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _Deposits.Contract.DEPOSITCANCELDELAY(&_Deposits.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_Deposits *DepositsCaller) FREEZEGRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deposits.contract.Call(opts, &out, "FREEZE_GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_Deposits *DepositsSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _Deposits.Contract.FREEZEGRACEPERIOD(&_Deposits.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_Deposits *DepositsCallerSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _Deposits.Contract.FREEZEGRACEPERIOD(&_Deposits.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_Deposits *DepositsCaller) MAXFORCEDACTIONSREQSPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deposits.contract.Call(opts, &out, "MAX_FORCED_ACTIONS_REQS_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_Deposits *DepositsSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _Deposits.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_Deposits.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_Deposits *DepositsCallerSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _Deposits.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_Deposits.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_Deposits *DepositsCaller) MAXVERIFIERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deposits.contract.Call(opts, &out, "MAX_VERIFIER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_Deposits *DepositsSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _Deposits.Contract.MAXVERIFIERCOUNT(&_Deposits.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_Deposits *DepositsCallerSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _Deposits.Contract.MAXVERIFIERCOUNT(&_Deposits.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_Deposits *DepositsCaller) UNFREEZEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deposits.contract.Call(opts, &out, "UNFREEZE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_Deposits *DepositsSession) UNFREEZEDELAY() (*big.Int, error) {
	return _Deposits.Contract.UNFREEZEDELAY(&_Deposits.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_Deposits *DepositsCallerSession) UNFREEZEDELAY() (*big.Int, error) {
	return _Deposits.Contract.UNFREEZEDELAY(&_Deposits.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_Deposits *DepositsCaller) VERIFIERREMOVALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deposits.contract.Call(opts, &out, "VERIFIER_REMOVAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_Deposits *DepositsSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _Deposits.Contract.VERIFIERREMOVALDELAY(&_Deposits.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_Deposits *DepositsCallerSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _Deposits.Contract.VERIFIERREMOVALDELAY(&_Deposits.CallOpts)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_Deposits *DepositsCaller) GetAssetInfo(opts *bind.CallOpts, assetType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Deposits.contract.Call(opts, &out, "getAssetInfo", assetType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_Deposits *DepositsSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _Deposits.Contract.GetAssetInfo(&_Deposits.CallOpts, assetType)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_Deposits *DepositsCallerSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _Deposits.Contract.GetAssetInfo(&_Deposits.CallOpts, assetType)
}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_Deposits *DepositsCaller) GetCancellationRequest(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Deposits.contract.Call(opts, &out, "getCancellationRequest", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_Deposits *DepositsSession) GetCancellationRequest(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Deposits.Contract.GetCancellationRequest(&_Deposits.CallOpts, starkKey, assetId, vaultId)
}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_Deposits *DepositsCallerSession) GetCancellationRequest(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Deposits.Contract.GetCancellationRequest(&_Deposits.CallOpts, starkKey, assetId, vaultId)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Deposits *DepositsCaller) GetDepositBalance(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Deposits.contract.Call(opts, &out, "getDepositBalance", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Deposits *DepositsSession) GetDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Deposits.Contract.GetDepositBalance(&_Deposits.CallOpts, starkKey, assetId, vaultId)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Deposits *DepositsCallerSession) GetDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Deposits.Contract.GetDepositBalance(&_Deposits.CallOpts, starkKey, assetId, vaultId)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_Deposits *DepositsCaller) GetEthKey(opts *bind.CallOpts, ownerKey *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Deposits.contract.Call(opts, &out, "getEthKey", ownerKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_Deposits *DepositsSession) GetEthKey(ownerKey *big.Int) (common.Address, error) {
	return _Deposits.Contract.GetEthKey(&_Deposits.CallOpts, ownerKey)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_Deposits *DepositsCallerSession) GetEthKey(ownerKey *big.Int) (common.Address, error) {
	return _Deposits.Contract.GetEthKey(&_Deposits.CallOpts, ownerKey)
}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Deposits *DepositsCaller) GetQuantizedDepositBalance(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Deposits.contract.Call(opts, &out, "getQuantizedDepositBalance", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Deposits *DepositsSession) GetQuantizedDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Deposits.Contract.GetQuantizedDepositBalance(&_Deposits.CallOpts, starkKey, assetId, vaultId)
}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Deposits *DepositsCallerSession) GetQuantizedDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Deposits.Contract.GetQuantizedDepositBalance(&_Deposits.CallOpts, starkKey, assetId, vaultId)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_Deposits *DepositsCaller) GetQuantum(opts *bind.CallOpts, presumedAssetType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Deposits.contract.Call(opts, &out, "getQuantum", presumedAssetType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_Deposits *DepositsSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _Deposits.Contract.GetQuantum(&_Deposits.CallOpts, presumedAssetType)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_Deposits *DepositsCallerSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _Deposits.Contract.GetQuantum(&_Deposits.CallOpts, presumedAssetType)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_Deposits *DepositsCaller) IsFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Deposits.contract.Call(opts, &out, "isFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_Deposits *DepositsSession) IsFrozen() (bool, error) {
	return _Deposits.Contract.IsFrozen(&_Deposits.CallOpts)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_Deposits *DepositsCallerSession) IsFrozen() (bool, error) {
	return _Deposits.Contract.IsFrozen(&_Deposits.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Deposits *DepositsTransactor) Deposit(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Deposits.contract.Transact(opts, "deposit", starkKey, assetType, vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Deposits *DepositsSession) Deposit(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.Deposit(&_Deposits.TransactOpts, starkKey, assetType, vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Deposits *DepositsTransactorSession) Deposit(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.Deposit(&_Deposits.TransactOpts, starkKey, assetType, vaultId)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Deposits *DepositsTransactor) Deposit0(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Deposits.contract.Transact(opts, "deposit0", starkKey, assetType, vaultId, quantizedAmount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Deposits *DepositsSession) Deposit0(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.Deposit0(&_Deposits.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Deposits *DepositsTransactorSession) Deposit0(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.Deposit0(&_Deposits.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Deposits *DepositsTransactor) DepositCancel(opts *bind.TransactOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Deposits.contract.Transact(opts, "depositCancel", starkKey, assetId, vaultId)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Deposits *DepositsSession) DepositCancel(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.DepositCancel(&_Deposits.TransactOpts, starkKey, assetId, vaultId)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Deposits *DepositsTransactorSession) DepositCancel(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.DepositCancel(&_Deposits.TransactOpts, starkKey, assetId, vaultId)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Deposits *DepositsTransactor) DepositERC20(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Deposits.contract.Transact(opts, "depositERC20", starkKey, assetType, vaultId, quantizedAmount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Deposits *DepositsSession) DepositERC20(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.DepositERC20(&_Deposits.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Deposits *DepositsTransactorSession) DepositERC20(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.DepositERC20(&_Deposits.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Deposits *DepositsTransactor) DepositEth(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Deposits.contract.Transact(opts, "depositEth", starkKey, assetType, vaultId)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Deposits *DepositsSession) DepositEth(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.DepositEth(&_Deposits.TransactOpts, starkKey, assetType, vaultId)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Deposits *DepositsTransactorSession) DepositEth(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.DepositEth(&_Deposits.TransactOpts, starkKey, assetType, vaultId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Deposits *DepositsTransactor) DepositNft(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Deposits.contract.Transact(opts, "depositNft", starkKey, assetType, vaultId, tokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Deposits *DepositsSession) DepositNft(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.DepositNft(&_Deposits.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Deposits *DepositsTransactorSession) DepositNft(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.DepositNft(&_Deposits.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Deposits *DepositsTransactor) DepositNftReclaim(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Deposits.contract.Transact(opts, "depositNftReclaim", starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Deposits *DepositsSession) DepositNftReclaim(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.DepositNftReclaim(&_Deposits.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Deposits *DepositsTransactorSession) DepositNftReclaim(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.DepositNftReclaim(&_Deposits.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Deposits *DepositsTransactor) DepositReclaim(opts *bind.TransactOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Deposits.contract.Transact(opts, "depositReclaim", starkKey, assetId, vaultId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Deposits *DepositsSession) DepositReclaim(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.DepositReclaim(&_Deposits.TransactOpts, starkKey, assetId, vaultId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Deposits *DepositsTransactorSession) DepositReclaim(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Deposits.Contract.DepositReclaim(&_Deposits.TransactOpts, starkKey, assetId, vaultId)
}

// DepositsLogDepositIterator is returned from FilterLogDeposit and is used to iterate over the raw logs and unpacked data for LogDeposit events raised by the Deposits contract.
type DepositsLogDepositIterator struct {
	Event *DepositsLogDeposit // Event containing the contract specifics and raw log

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
func (it *DepositsLogDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositsLogDeposit)
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
		it.Event = new(DepositsLogDeposit)
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
func (it *DepositsLogDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositsLogDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositsLogDeposit represents a LogDeposit event raised by the Deposits contract.
type DepositsLogDeposit struct {
	DepositorEthKey    common.Address
	StarkKey           *big.Int
	VaultId            *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDeposit is a free log retrieval operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Deposits *DepositsFilterer) FilterLogDeposit(opts *bind.FilterOpts) (*DepositsLogDepositIterator, error) {

	logs, sub, err := _Deposits.contract.FilterLogs(opts, "LogDeposit")
	if err != nil {
		return nil, err
	}
	return &DepositsLogDepositIterator{contract: _Deposits.contract, event: "LogDeposit", logs: logs, sub: sub}, nil
}

// WatchLogDeposit is a free log subscription operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Deposits *DepositsFilterer) WatchLogDeposit(opts *bind.WatchOpts, sink chan<- *DepositsLogDeposit) (event.Subscription, error) {

	logs, sub, err := _Deposits.contract.WatchLogs(opts, "LogDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositsLogDeposit)
				if err := _Deposits.contract.UnpackLog(event, "LogDeposit", log); err != nil {
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

// ParseLogDeposit is a log parse operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Deposits *DepositsFilterer) ParseLogDeposit(log types.Log) (*DepositsLogDeposit, error) {
	event := new(DepositsLogDeposit)
	if err := _Deposits.contract.UnpackLog(event, "LogDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositsLogDepositCancelIterator is returned from FilterLogDepositCancel and is used to iterate over the raw logs and unpacked data for LogDepositCancel events raised by the Deposits contract.
type DepositsLogDepositCancelIterator struct {
	Event *DepositsLogDepositCancel // Event containing the contract specifics and raw log

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
func (it *DepositsLogDepositCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositsLogDepositCancel)
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
		it.Event = new(DepositsLogDepositCancel)
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
func (it *DepositsLogDepositCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositsLogDepositCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositsLogDepositCancel represents a LogDepositCancel event raised by the Deposits contract.
type DepositsLogDepositCancel struct {
	StarkKey *big.Int
	VaultId  *big.Int
	AssetId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogDepositCancel is a free log retrieval operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_Deposits *DepositsFilterer) FilterLogDepositCancel(opts *bind.FilterOpts) (*DepositsLogDepositCancelIterator, error) {

	logs, sub, err := _Deposits.contract.FilterLogs(opts, "LogDepositCancel")
	if err != nil {
		return nil, err
	}
	return &DepositsLogDepositCancelIterator{contract: _Deposits.contract, event: "LogDepositCancel", logs: logs, sub: sub}, nil
}

// WatchLogDepositCancel is a free log subscription operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_Deposits *DepositsFilterer) WatchLogDepositCancel(opts *bind.WatchOpts, sink chan<- *DepositsLogDepositCancel) (event.Subscription, error) {

	logs, sub, err := _Deposits.contract.WatchLogs(opts, "LogDepositCancel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositsLogDepositCancel)
				if err := _Deposits.contract.UnpackLog(event, "LogDepositCancel", log); err != nil {
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

// ParseLogDepositCancel is a log parse operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_Deposits *DepositsFilterer) ParseLogDepositCancel(log types.Log) (*DepositsLogDepositCancel, error) {
	event := new(DepositsLogDepositCancel)
	if err := _Deposits.contract.UnpackLog(event, "LogDepositCancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositsLogDepositCancelReclaimedIterator is returned from FilterLogDepositCancelReclaimed and is used to iterate over the raw logs and unpacked data for LogDepositCancelReclaimed events raised by the Deposits contract.
type DepositsLogDepositCancelReclaimedIterator struct {
	Event *DepositsLogDepositCancelReclaimed // Event containing the contract specifics and raw log

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
func (it *DepositsLogDepositCancelReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositsLogDepositCancelReclaimed)
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
		it.Event = new(DepositsLogDepositCancelReclaimed)
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
func (it *DepositsLogDepositCancelReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositsLogDepositCancelReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositsLogDepositCancelReclaimed represents a LogDepositCancelReclaimed event raised by the Deposits contract.
type DepositsLogDepositCancelReclaimed struct {
	StarkKey           *big.Int
	VaultId            *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDepositCancelReclaimed is a free log retrieval operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Deposits *DepositsFilterer) FilterLogDepositCancelReclaimed(opts *bind.FilterOpts) (*DepositsLogDepositCancelReclaimedIterator, error) {

	logs, sub, err := _Deposits.contract.FilterLogs(opts, "LogDepositCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return &DepositsLogDepositCancelReclaimedIterator{contract: _Deposits.contract, event: "LogDepositCancelReclaimed", logs: logs, sub: sub}, nil
}

// WatchLogDepositCancelReclaimed is a free log subscription operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Deposits *DepositsFilterer) WatchLogDepositCancelReclaimed(opts *bind.WatchOpts, sink chan<- *DepositsLogDepositCancelReclaimed) (event.Subscription, error) {

	logs, sub, err := _Deposits.contract.WatchLogs(opts, "LogDepositCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositsLogDepositCancelReclaimed)
				if err := _Deposits.contract.UnpackLog(event, "LogDepositCancelReclaimed", log); err != nil {
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

// ParseLogDepositCancelReclaimed is a log parse operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Deposits *DepositsFilterer) ParseLogDepositCancelReclaimed(log types.Log) (*DepositsLogDepositCancelReclaimed, error) {
	event := new(DepositsLogDepositCancelReclaimed)
	if err := _Deposits.contract.UnpackLog(event, "LogDepositCancelReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositsLogDepositNftCancelReclaimedIterator is returned from FilterLogDepositNftCancelReclaimed and is used to iterate over the raw logs and unpacked data for LogDepositNftCancelReclaimed events raised by the Deposits contract.
type DepositsLogDepositNftCancelReclaimedIterator struct {
	Event *DepositsLogDepositNftCancelReclaimed // Event containing the contract specifics and raw log

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
func (it *DepositsLogDepositNftCancelReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositsLogDepositNftCancelReclaimed)
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
		it.Event = new(DepositsLogDepositNftCancelReclaimed)
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
func (it *DepositsLogDepositNftCancelReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositsLogDepositNftCancelReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositsLogDepositNftCancelReclaimed represents a LogDepositNftCancelReclaimed event raised by the Deposits contract.
type DepositsLogDepositNftCancelReclaimed struct {
	StarkKey  *big.Int
	VaultId   *big.Int
	AssetType *big.Int
	TokenId   *big.Int
	AssetId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogDepositNftCancelReclaimed is a free log retrieval operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Deposits *DepositsFilterer) FilterLogDepositNftCancelReclaimed(opts *bind.FilterOpts) (*DepositsLogDepositNftCancelReclaimedIterator, error) {

	logs, sub, err := _Deposits.contract.FilterLogs(opts, "LogDepositNftCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return &DepositsLogDepositNftCancelReclaimedIterator{contract: _Deposits.contract, event: "LogDepositNftCancelReclaimed", logs: logs, sub: sub}, nil
}

// WatchLogDepositNftCancelReclaimed is a free log subscription operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Deposits *DepositsFilterer) WatchLogDepositNftCancelReclaimed(opts *bind.WatchOpts, sink chan<- *DepositsLogDepositNftCancelReclaimed) (event.Subscription, error) {

	logs, sub, err := _Deposits.contract.WatchLogs(opts, "LogDepositNftCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositsLogDepositNftCancelReclaimed)
				if err := _Deposits.contract.UnpackLog(event, "LogDepositNftCancelReclaimed", log); err != nil {
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

// ParseLogDepositNftCancelReclaimed is a log parse operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Deposits *DepositsFilterer) ParseLogDepositNftCancelReclaimed(log types.Log) (*DepositsLogDepositNftCancelReclaimed, error) {
	event := new(DepositsLogDepositNftCancelReclaimed)
	if err := _Deposits.contract.UnpackLog(event, "LogDepositNftCancelReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositsLogNftDepositIterator is returned from FilterLogNftDeposit and is used to iterate over the raw logs and unpacked data for LogNftDeposit events raised by the Deposits contract.
type DepositsLogNftDepositIterator struct {
	Event *DepositsLogNftDeposit // Event containing the contract specifics and raw log

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
func (it *DepositsLogNftDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositsLogNftDeposit)
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
		it.Event = new(DepositsLogNftDeposit)
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
func (it *DepositsLogNftDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositsLogNftDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositsLogNftDeposit represents a LogNftDeposit event raised by the Deposits contract.
type DepositsLogNftDeposit struct {
	DepositorEthKey common.Address
	StarkKey        *big.Int
	VaultId         *big.Int
	AssetType       *big.Int
	TokenId         *big.Int
	AssetId         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogNftDeposit is a free log retrieval operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Deposits *DepositsFilterer) FilterLogNftDeposit(opts *bind.FilterOpts) (*DepositsLogNftDepositIterator, error) {

	logs, sub, err := _Deposits.contract.FilterLogs(opts, "LogNftDeposit")
	if err != nil {
		return nil, err
	}
	return &DepositsLogNftDepositIterator{contract: _Deposits.contract, event: "LogNftDeposit", logs: logs, sub: sub}, nil
}

// WatchLogNftDeposit is a free log subscription operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Deposits *DepositsFilterer) WatchLogNftDeposit(opts *bind.WatchOpts, sink chan<- *DepositsLogNftDeposit) (event.Subscription, error) {

	logs, sub, err := _Deposits.contract.WatchLogs(opts, "LogNftDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositsLogNftDeposit)
				if err := _Deposits.contract.UnpackLog(event, "LogNftDeposit", log); err != nil {
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

// ParseLogNftDeposit is a log parse operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Deposits *DepositsFilterer) ParseLogNftDeposit(log types.Log) (*DepositsLogNftDeposit, error) {
	event := new(DepositsLogNftDeposit)
	if err := _Deposits.contract.UnpackLog(event, "LogNftDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
