// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package withdrawals

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

// WithdrawalsMetaData contains all meta data concerning the Withdrawals contract.
var WithdrawalsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogMintWithdrawalPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"LogNftWithdrawalPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"LogWithdrawalPerformed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"getAssetInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"}],\"name\":\"getEthKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"presumedAssetType\",\"type\":\"uint256\"}],\"name\":\"getQuantum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"mintingBlob\",\"type\":\"bytes\"}],\"name\":\"withdrawAndMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WithdrawalsABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawalsMetaData.ABI instead.
var WithdrawalsABI = WithdrawalsMetaData.ABI

// Withdrawals is an auto generated Go binding around an Ethereum contract.
type Withdrawals struct {
	WithdrawalsCaller     // Read-only binding to the contract
	WithdrawalsTransactor // Write-only binding to the contract
	WithdrawalsFilterer   // Log filterer for contract events
}

// WithdrawalsCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithdrawalsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithdrawalsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithdrawalsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithdrawalsSession struct {
	Contract     *Withdrawals      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WithdrawalsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithdrawalsCallerSession struct {
	Contract *WithdrawalsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WithdrawalsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithdrawalsTransactorSession struct {
	Contract     *WithdrawalsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WithdrawalsRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithdrawalsRaw struct {
	Contract *Withdrawals // Generic contract binding to access the raw methods on
}

// WithdrawalsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithdrawalsCallerRaw struct {
	Contract *WithdrawalsCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawalsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithdrawalsTransactorRaw struct {
	Contract *WithdrawalsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawals creates a new instance of Withdrawals, bound to a specific deployed contract.
func NewWithdrawals(address common.Address, backend bind.ContractBackend) (*Withdrawals, error) {
	contract, err := bindWithdrawals(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Withdrawals{WithdrawalsCaller: WithdrawalsCaller{contract: contract}, WithdrawalsTransactor: WithdrawalsTransactor{contract: contract}, WithdrawalsFilterer: WithdrawalsFilterer{contract: contract}}, nil
}

// NewWithdrawalsCaller creates a new read-only instance of Withdrawals, bound to a specific deployed contract.
func NewWithdrawalsCaller(address common.Address, caller bind.ContractCaller) (*WithdrawalsCaller, error) {
	contract, err := bindWithdrawals(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawalsCaller{contract: contract}, nil
}

// NewWithdrawalsTransactor creates a new write-only instance of Withdrawals, bound to a specific deployed contract.
func NewWithdrawalsTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawalsTransactor, error) {
	contract, err := bindWithdrawals(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawalsTransactor{contract: contract}, nil
}

// NewWithdrawalsFilterer creates a new log filterer instance of Withdrawals, bound to a specific deployed contract.
func NewWithdrawalsFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawalsFilterer, error) {
	contract, err := bindWithdrawals(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawalsFilterer{contract: contract}, nil
}

// bindWithdrawals binds a generic wrapper to an already deployed contract.
func bindWithdrawals(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WithdrawalsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdrawals *WithdrawalsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Withdrawals.Contract.WithdrawalsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdrawals *WithdrawalsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawals.Contract.WithdrawalsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdrawals *WithdrawalsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdrawals.Contract.WithdrawalsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdrawals *WithdrawalsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Withdrawals.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdrawals *WithdrawalsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawals.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdrawals *WithdrawalsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdrawals.Contract.contract.Transact(opts, method, params...)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_Withdrawals *WithdrawalsCaller) GetAssetInfo(opts *bind.CallOpts, assetType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Withdrawals.contract.Call(opts, &out, "getAssetInfo", assetType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_Withdrawals *WithdrawalsSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _Withdrawals.Contract.GetAssetInfo(&_Withdrawals.CallOpts, assetType)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_Withdrawals *WithdrawalsCallerSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _Withdrawals.Contract.GetAssetInfo(&_Withdrawals.CallOpts, assetType)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_Withdrawals *WithdrawalsCaller) GetEthKey(opts *bind.CallOpts, ownerKey *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Withdrawals.contract.Call(opts, &out, "getEthKey", ownerKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_Withdrawals *WithdrawalsSession) GetEthKey(ownerKey *big.Int) (common.Address, error) {
	return _Withdrawals.Contract.GetEthKey(&_Withdrawals.CallOpts, ownerKey)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_Withdrawals *WithdrawalsCallerSession) GetEthKey(ownerKey *big.Int) (common.Address, error) {
	return _Withdrawals.Contract.GetEthKey(&_Withdrawals.CallOpts, ownerKey)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_Withdrawals *WithdrawalsCaller) GetQuantum(opts *bind.CallOpts, presumedAssetType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Withdrawals.contract.Call(opts, &out, "getQuantum", presumedAssetType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_Withdrawals *WithdrawalsSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _Withdrawals.Contract.GetQuantum(&_Withdrawals.CallOpts, presumedAssetType)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_Withdrawals *WithdrawalsCallerSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _Withdrawals.Contract.GetQuantum(&_Withdrawals.CallOpts, presumedAssetType)
}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 ownerKey, uint256 assetId) view returns(uint256 balance)
func (_Withdrawals *WithdrawalsCaller) GetWithdrawalBalance(opts *bind.CallOpts, ownerKey *big.Int, assetId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Withdrawals.contract.Call(opts, &out, "getWithdrawalBalance", ownerKey, assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 ownerKey, uint256 assetId) view returns(uint256 balance)
func (_Withdrawals *WithdrawalsSession) GetWithdrawalBalance(ownerKey *big.Int, assetId *big.Int) (*big.Int, error) {
	return _Withdrawals.Contract.GetWithdrawalBalance(&_Withdrawals.CallOpts, ownerKey, assetId)
}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 ownerKey, uint256 assetId) view returns(uint256 balance)
func (_Withdrawals *WithdrawalsCallerSession) GetWithdrawalBalance(ownerKey *big.Int, assetId *big.Int) (*big.Int, error) {
	return _Withdrawals.Contract.GetWithdrawalBalance(&_Withdrawals.CallOpts, ownerKey, assetId)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_Withdrawals *WithdrawalsCaller) IsFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Withdrawals.contract.Call(opts, &out, "isFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_Withdrawals *WithdrawalsSession) IsFrozen() (bool, error) {
	return _Withdrawals.Contract.IsFrozen(&_Withdrawals.CallOpts)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_Withdrawals *WithdrawalsCallerSession) IsFrozen() (bool, error) {
	return _Withdrawals.Contract.IsFrozen(&_Withdrawals.CallOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 ownerKey, uint256 assetType) returns()
func (_Withdrawals *WithdrawalsTransactor) Withdraw(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _Withdrawals.contract.Transact(opts, "withdraw", ownerKey, assetType)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 ownerKey, uint256 assetType) returns()
func (_Withdrawals *WithdrawalsSession) Withdraw(ownerKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _Withdrawals.Contract.Withdraw(&_Withdrawals.TransactOpts, ownerKey, assetType)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 ownerKey, uint256 assetType) returns()
func (_Withdrawals *WithdrawalsTransactorSession) Withdraw(ownerKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _Withdrawals.Contract.Withdraw(&_Withdrawals.TransactOpts, ownerKey, assetType)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 ownerKey, uint256 assetType, bytes mintingBlob) returns()
func (_Withdrawals *WithdrawalsTransactor) WithdrawAndMint(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _Withdrawals.contract.Transact(opts, "withdrawAndMint", ownerKey, assetType, mintingBlob)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 ownerKey, uint256 assetType, bytes mintingBlob) returns()
func (_Withdrawals *WithdrawalsSession) WithdrawAndMint(ownerKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _Withdrawals.Contract.WithdrawAndMint(&_Withdrawals.TransactOpts, ownerKey, assetType, mintingBlob)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 ownerKey, uint256 assetType, bytes mintingBlob) returns()
func (_Withdrawals *WithdrawalsTransactorSession) WithdrawAndMint(ownerKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _Withdrawals.Contract.WithdrawAndMint(&_Withdrawals.TransactOpts, ownerKey, assetType, mintingBlob)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_Withdrawals *WithdrawalsTransactor) WithdrawNft(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Withdrawals.contract.Transact(opts, "withdrawNft", ownerKey, assetType, tokenId)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_Withdrawals *WithdrawalsSession) WithdrawNft(ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Withdrawals.Contract.WithdrawNft(&_Withdrawals.TransactOpts, ownerKey, assetType, tokenId)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_Withdrawals *WithdrawalsTransactorSession) WithdrawNft(ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Withdrawals.Contract.WithdrawNft(&_Withdrawals.TransactOpts, ownerKey, assetType, tokenId)
}

// WithdrawalsLogMintWithdrawalPerformedIterator is returned from FilterLogMintWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogMintWithdrawalPerformed events raised by the Withdrawals contract.
type WithdrawalsLogMintWithdrawalPerformedIterator struct {
	Event *WithdrawalsLogMintWithdrawalPerformed // Event containing the contract specifics and raw log

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
func (it *WithdrawalsLogMintWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalsLogMintWithdrawalPerformed)
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
		it.Event = new(WithdrawalsLogMintWithdrawalPerformed)
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
func (it *WithdrawalsLogMintWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalsLogMintWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalsLogMintWithdrawalPerformed represents a LogMintWithdrawalPerformed event raised by the Withdrawals contract.
type WithdrawalsLogMintWithdrawalPerformed struct {
	OwnerKey           *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	AssetId            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogMintWithdrawalPerformed is a free log retrieval operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_Withdrawals *WithdrawalsFilterer) FilterLogMintWithdrawalPerformed(opts *bind.FilterOpts) (*WithdrawalsLogMintWithdrawalPerformedIterator, error) {

	logs, sub, err := _Withdrawals.contract.FilterLogs(opts, "LogMintWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &WithdrawalsLogMintWithdrawalPerformedIterator{contract: _Withdrawals.contract, event: "LogMintWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogMintWithdrawalPerformed is a free log subscription operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_Withdrawals *WithdrawalsFilterer) WatchLogMintWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *WithdrawalsLogMintWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _Withdrawals.contract.WatchLogs(opts, "LogMintWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalsLogMintWithdrawalPerformed)
				if err := _Withdrawals.contract.UnpackLog(event, "LogMintWithdrawalPerformed", log); err != nil {
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

// ParseLogMintWithdrawalPerformed is a log parse operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_Withdrawals *WithdrawalsFilterer) ParseLogMintWithdrawalPerformed(log types.Log) (*WithdrawalsLogMintWithdrawalPerformed, error) {
	event := new(WithdrawalsLogMintWithdrawalPerformed)
	if err := _Withdrawals.contract.UnpackLog(event, "LogMintWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalsLogNftWithdrawalPerformedIterator is returned from FilterLogNftWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogNftWithdrawalPerformed events raised by the Withdrawals contract.
type WithdrawalsLogNftWithdrawalPerformedIterator struct {
	Event *WithdrawalsLogNftWithdrawalPerformed // Event containing the contract specifics and raw log

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
func (it *WithdrawalsLogNftWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalsLogNftWithdrawalPerformed)
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
		it.Event = new(WithdrawalsLogNftWithdrawalPerformed)
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
func (it *WithdrawalsLogNftWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalsLogNftWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalsLogNftWithdrawalPerformed represents a LogNftWithdrawalPerformed event raised by the Withdrawals contract.
type WithdrawalsLogNftWithdrawalPerformed struct {
	OwnerKey  *big.Int
	AssetType *big.Int
	TokenId   *big.Int
	AssetId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogNftWithdrawalPerformed is a free log retrieval operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_Withdrawals *WithdrawalsFilterer) FilterLogNftWithdrawalPerformed(opts *bind.FilterOpts) (*WithdrawalsLogNftWithdrawalPerformedIterator, error) {

	logs, sub, err := _Withdrawals.contract.FilterLogs(opts, "LogNftWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &WithdrawalsLogNftWithdrawalPerformedIterator{contract: _Withdrawals.contract, event: "LogNftWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogNftWithdrawalPerformed is a free log subscription operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_Withdrawals *WithdrawalsFilterer) WatchLogNftWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *WithdrawalsLogNftWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _Withdrawals.contract.WatchLogs(opts, "LogNftWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalsLogNftWithdrawalPerformed)
				if err := _Withdrawals.contract.UnpackLog(event, "LogNftWithdrawalPerformed", log); err != nil {
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

// ParseLogNftWithdrawalPerformed is a log parse operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_Withdrawals *WithdrawalsFilterer) ParseLogNftWithdrawalPerformed(log types.Log) (*WithdrawalsLogNftWithdrawalPerformed, error) {
	event := new(WithdrawalsLogNftWithdrawalPerformed)
	if err := _Withdrawals.contract.UnpackLog(event, "LogNftWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalsLogWithdrawalPerformedIterator is returned from FilterLogWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogWithdrawalPerformed events raised by the Withdrawals contract.
type WithdrawalsLogWithdrawalPerformedIterator struct {
	Event *WithdrawalsLogWithdrawalPerformed // Event containing the contract specifics and raw log

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
func (it *WithdrawalsLogWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalsLogWithdrawalPerformed)
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
		it.Event = new(WithdrawalsLogWithdrawalPerformed)
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
func (it *WithdrawalsLogWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalsLogWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalsLogWithdrawalPerformed represents a LogWithdrawalPerformed event raised by the Withdrawals contract.
type WithdrawalsLogWithdrawalPerformed struct {
	OwnerKey           *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Recipient          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalPerformed is a free log retrieval operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_Withdrawals *WithdrawalsFilterer) FilterLogWithdrawalPerformed(opts *bind.FilterOpts) (*WithdrawalsLogWithdrawalPerformedIterator, error) {

	logs, sub, err := _Withdrawals.contract.FilterLogs(opts, "LogWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &WithdrawalsLogWithdrawalPerformedIterator{contract: _Withdrawals.contract, event: "LogWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalPerformed is a free log subscription operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_Withdrawals *WithdrawalsFilterer) WatchLogWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *WithdrawalsLogWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _Withdrawals.contract.WatchLogs(opts, "LogWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalsLogWithdrawalPerformed)
				if err := _Withdrawals.contract.UnpackLog(event, "LogWithdrawalPerformed", log); err != nil {
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

// ParseLogWithdrawalPerformed is a log parse operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_Withdrawals *WithdrawalsFilterer) ParseLogWithdrawalPerformed(log types.Log) (*WithdrawalsLogWithdrawalPerformed, error) {
	event := new(WithdrawalsLogWithdrawalPerformed)
	if err := _Withdrawals.contract.UnpackLog(event, "LogWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
