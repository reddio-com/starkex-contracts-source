// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package acceptModifications

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

// AcceptModificationsMetaData contains all meta data concerning the AcceptModifications contract.
var AcceptModificationsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogAssetWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogMintableWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogNftWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawalAllowed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FREEZE_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FORCED_ACTIONS_REQS_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_VERIFIER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNFREEZE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_REMOVAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"presumedAssetType\",\"type\":\"uint256\"}],\"name\":\"getQuantum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AcceptModificationsABI is the input ABI used to generate the binding from.
// Deprecated: Use AcceptModificationsMetaData.ABI instead.
var AcceptModificationsABI = AcceptModificationsMetaData.ABI

// AcceptModifications is an auto generated Go binding around an Ethereum contract.
type AcceptModifications struct {
	AcceptModificationsCaller     // Read-only binding to the contract
	AcceptModificationsTransactor // Write-only binding to the contract
	AcceptModificationsFilterer   // Log filterer for contract events
}

// AcceptModificationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AcceptModificationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AcceptModificationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AcceptModificationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AcceptModificationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AcceptModificationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AcceptModificationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AcceptModificationsSession struct {
	Contract     *AcceptModifications // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AcceptModificationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AcceptModificationsCallerSession struct {
	Contract *AcceptModificationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AcceptModificationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AcceptModificationsTransactorSession struct {
	Contract     *AcceptModificationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AcceptModificationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AcceptModificationsRaw struct {
	Contract *AcceptModifications // Generic contract binding to access the raw methods on
}

// AcceptModificationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AcceptModificationsCallerRaw struct {
	Contract *AcceptModificationsCaller // Generic read-only contract binding to access the raw methods on
}

// AcceptModificationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AcceptModificationsTransactorRaw struct {
	Contract *AcceptModificationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAcceptModifications creates a new instance of AcceptModifications, bound to a specific deployed contract.
func NewAcceptModifications(address common.Address, backend bind.ContractBackend) (*AcceptModifications, error) {
	contract, err := bindAcceptModifications(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AcceptModifications{AcceptModificationsCaller: AcceptModificationsCaller{contract: contract}, AcceptModificationsTransactor: AcceptModificationsTransactor{contract: contract}, AcceptModificationsFilterer: AcceptModificationsFilterer{contract: contract}}, nil
}

// NewAcceptModificationsCaller creates a new read-only instance of AcceptModifications, bound to a specific deployed contract.
func NewAcceptModificationsCaller(address common.Address, caller bind.ContractCaller) (*AcceptModificationsCaller, error) {
	contract, err := bindAcceptModifications(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AcceptModificationsCaller{contract: contract}, nil
}

// NewAcceptModificationsTransactor creates a new write-only instance of AcceptModifications, bound to a specific deployed contract.
func NewAcceptModificationsTransactor(address common.Address, transactor bind.ContractTransactor) (*AcceptModificationsTransactor, error) {
	contract, err := bindAcceptModifications(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AcceptModificationsTransactor{contract: contract}, nil
}

// NewAcceptModificationsFilterer creates a new log filterer instance of AcceptModifications, bound to a specific deployed contract.
func NewAcceptModificationsFilterer(address common.Address, filterer bind.ContractFilterer) (*AcceptModificationsFilterer, error) {
	contract, err := bindAcceptModifications(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AcceptModificationsFilterer{contract: contract}, nil
}

// bindAcceptModifications binds a generic wrapper to an already deployed contract.
func bindAcceptModifications(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AcceptModificationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AcceptModifications *AcceptModificationsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AcceptModifications.Contract.AcceptModificationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AcceptModifications *AcceptModificationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AcceptModifications.Contract.AcceptModificationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AcceptModifications *AcceptModificationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AcceptModifications.Contract.AcceptModificationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AcceptModifications *AcceptModificationsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AcceptModifications.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AcceptModifications *AcceptModificationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AcceptModifications.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AcceptModifications *AcceptModificationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AcceptModifications.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_AcceptModifications *AcceptModificationsCaller) DEPOSITCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AcceptModifications.contract.Call(opts, &out, "DEPOSIT_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_AcceptModifications *AcceptModificationsSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _AcceptModifications.Contract.DEPOSITCANCELDELAY(&_AcceptModifications.CallOpts)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_AcceptModifications *AcceptModificationsCallerSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _AcceptModifications.Contract.DEPOSITCANCELDELAY(&_AcceptModifications.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_AcceptModifications *AcceptModificationsCaller) FREEZEGRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AcceptModifications.contract.Call(opts, &out, "FREEZE_GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_AcceptModifications *AcceptModificationsSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _AcceptModifications.Contract.FREEZEGRACEPERIOD(&_AcceptModifications.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_AcceptModifications *AcceptModificationsCallerSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _AcceptModifications.Contract.FREEZEGRACEPERIOD(&_AcceptModifications.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_AcceptModifications *AcceptModificationsCaller) MAXFORCEDACTIONSREQSPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AcceptModifications.contract.Call(opts, &out, "MAX_FORCED_ACTIONS_REQS_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_AcceptModifications *AcceptModificationsSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _AcceptModifications.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_AcceptModifications.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_AcceptModifications *AcceptModificationsCallerSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _AcceptModifications.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_AcceptModifications.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_AcceptModifications *AcceptModificationsCaller) MAXVERIFIERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AcceptModifications.contract.Call(opts, &out, "MAX_VERIFIER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_AcceptModifications *AcceptModificationsSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _AcceptModifications.Contract.MAXVERIFIERCOUNT(&_AcceptModifications.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_AcceptModifications *AcceptModificationsCallerSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _AcceptModifications.Contract.MAXVERIFIERCOUNT(&_AcceptModifications.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_AcceptModifications *AcceptModificationsCaller) UNFREEZEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AcceptModifications.contract.Call(opts, &out, "UNFREEZE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_AcceptModifications *AcceptModificationsSession) UNFREEZEDELAY() (*big.Int, error) {
	return _AcceptModifications.Contract.UNFREEZEDELAY(&_AcceptModifications.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_AcceptModifications *AcceptModificationsCallerSession) UNFREEZEDELAY() (*big.Int, error) {
	return _AcceptModifications.Contract.UNFREEZEDELAY(&_AcceptModifications.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_AcceptModifications *AcceptModificationsCaller) VERIFIERREMOVALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AcceptModifications.contract.Call(opts, &out, "VERIFIER_REMOVAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_AcceptModifications *AcceptModificationsSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _AcceptModifications.Contract.VERIFIERREMOVALDELAY(&_AcceptModifications.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_AcceptModifications *AcceptModificationsCallerSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _AcceptModifications.Contract.VERIFIERREMOVALDELAY(&_AcceptModifications.CallOpts)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_AcceptModifications *AcceptModificationsCaller) GetQuantum(opts *bind.CallOpts, presumedAssetType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AcceptModifications.contract.Call(opts, &out, "getQuantum", presumedAssetType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_AcceptModifications *AcceptModificationsSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _AcceptModifications.Contract.GetQuantum(&_AcceptModifications.CallOpts, presumedAssetType)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_AcceptModifications *AcceptModificationsCallerSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _AcceptModifications.Contract.GetQuantum(&_AcceptModifications.CallOpts, presumedAssetType)
}

// AcceptModificationsLogAssetWithdrawalAllowedIterator is returned from FilterLogAssetWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogAssetWithdrawalAllowed events raised by the AcceptModifications contract.
type AcceptModificationsLogAssetWithdrawalAllowedIterator struct {
	Event *AcceptModificationsLogAssetWithdrawalAllowed // Event containing the contract specifics and raw log

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
func (it *AcceptModificationsLogAssetWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcceptModificationsLogAssetWithdrawalAllowed)
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
		it.Event = new(AcceptModificationsLogAssetWithdrawalAllowed)
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
func (it *AcceptModificationsLogAssetWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcceptModificationsLogAssetWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcceptModificationsLogAssetWithdrawalAllowed represents a LogAssetWithdrawalAllowed event raised by the AcceptModifications contract.
type AcceptModificationsLogAssetWithdrawalAllowed struct {
	OwnerKey        *big.Int
	AssetId         *big.Int
	QuantizedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogAssetWithdrawalAllowed is a free log retrieval operation binding the contract event 0xd31f59c968eb53320f624721416cf88605da9aadbaa723405e52affe9de4b07f.
//
// Solidity: event LogAssetWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_AcceptModifications *AcceptModificationsFilterer) FilterLogAssetWithdrawalAllowed(opts *bind.FilterOpts) (*AcceptModificationsLogAssetWithdrawalAllowedIterator, error) {

	logs, sub, err := _AcceptModifications.contract.FilterLogs(opts, "LogAssetWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &AcceptModificationsLogAssetWithdrawalAllowedIterator{contract: _AcceptModifications.contract, event: "LogAssetWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogAssetWithdrawalAllowed is a free log subscription operation binding the contract event 0xd31f59c968eb53320f624721416cf88605da9aadbaa723405e52affe9de4b07f.
//
// Solidity: event LogAssetWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_AcceptModifications *AcceptModificationsFilterer) WatchLogAssetWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *AcceptModificationsLogAssetWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _AcceptModifications.contract.WatchLogs(opts, "LogAssetWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcceptModificationsLogAssetWithdrawalAllowed)
				if err := _AcceptModifications.contract.UnpackLog(event, "LogAssetWithdrawalAllowed", log); err != nil {
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

// ParseLogAssetWithdrawalAllowed is a log parse operation binding the contract event 0xd31f59c968eb53320f624721416cf88605da9aadbaa723405e52affe9de4b07f.
//
// Solidity: event LogAssetWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_AcceptModifications *AcceptModificationsFilterer) ParseLogAssetWithdrawalAllowed(log types.Log) (*AcceptModificationsLogAssetWithdrawalAllowed, error) {
	event := new(AcceptModificationsLogAssetWithdrawalAllowed)
	if err := _AcceptModifications.contract.UnpackLog(event, "LogAssetWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AcceptModificationsLogMintableWithdrawalAllowedIterator is returned from FilterLogMintableWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogMintableWithdrawalAllowed events raised by the AcceptModifications contract.
type AcceptModificationsLogMintableWithdrawalAllowedIterator struct {
	Event *AcceptModificationsLogMintableWithdrawalAllowed // Event containing the contract specifics and raw log

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
func (it *AcceptModificationsLogMintableWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcceptModificationsLogMintableWithdrawalAllowed)
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
		it.Event = new(AcceptModificationsLogMintableWithdrawalAllowed)
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
func (it *AcceptModificationsLogMintableWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcceptModificationsLogMintableWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcceptModificationsLogMintableWithdrawalAllowed represents a LogMintableWithdrawalAllowed event raised by the AcceptModifications contract.
type AcceptModificationsLogMintableWithdrawalAllowed struct {
	OwnerKey        *big.Int
	AssetId         *big.Int
	QuantizedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogMintableWithdrawalAllowed is a free log retrieval operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_AcceptModifications *AcceptModificationsFilterer) FilterLogMintableWithdrawalAllowed(opts *bind.FilterOpts) (*AcceptModificationsLogMintableWithdrawalAllowedIterator, error) {

	logs, sub, err := _AcceptModifications.contract.FilterLogs(opts, "LogMintableWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &AcceptModificationsLogMintableWithdrawalAllowedIterator{contract: _AcceptModifications.contract, event: "LogMintableWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogMintableWithdrawalAllowed is a free log subscription operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_AcceptModifications *AcceptModificationsFilterer) WatchLogMintableWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *AcceptModificationsLogMintableWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _AcceptModifications.contract.WatchLogs(opts, "LogMintableWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcceptModificationsLogMintableWithdrawalAllowed)
				if err := _AcceptModifications.contract.UnpackLog(event, "LogMintableWithdrawalAllowed", log); err != nil {
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

// ParseLogMintableWithdrawalAllowed is a log parse operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_AcceptModifications *AcceptModificationsFilterer) ParseLogMintableWithdrawalAllowed(log types.Log) (*AcceptModificationsLogMintableWithdrawalAllowed, error) {
	event := new(AcceptModificationsLogMintableWithdrawalAllowed)
	if err := _AcceptModifications.contract.UnpackLog(event, "LogMintableWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AcceptModificationsLogNftWithdrawalAllowedIterator is returned from FilterLogNftWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogNftWithdrawalAllowed events raised by the AcceptModifications contract.
type AcceptModificationsLogNftWithdrawalAllowedIterator struct {
	Event *AcceptModificationsLogNftWithdrawalAllowed // Event containing the contract specifics and raw log

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
func (it *AcceptModificationsLogNftWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcceptModificationsLogNftWithdrawalAllowed)
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
		it.Event = new(AcceptModificationsLogNftWithdrawalAllowed)
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
func (it *AcceptModificationsLogNftWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcceptModificationsLogNftWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcceptModificationsLogNftWithdrawalAllowed represents a LogNftWithdrawalAllowed event raised by the AcceptModifications contract.
type AcceptModificationsLogNftWithdrawalAllowed struct {
	OwnerKey *big.Int
	AssetId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNftWithdrawalAllowed is a free log retrieval operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 ownerKey, uint256 assetId)
func (_AcceptModifications *AcceptModificationsFilterer) FilterLogNftWithdrawalAllowed(opts *bind.FilterOpts) (*AcceptModificationsLogNftWithdrawalAllowedIterator, error) {

	logs, sub, err := _AcceptModifications.contract.FilterLogs(opts, "LogNftWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &AcceptModificationsLogNftWithdrawalAllowedIterator{contract: _AcceptModifications.contract, event: "LogNftWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogNftWithdrawalAllowed is a free log subscription operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 ownerKey, uint256 assetId)
func (_AcceptModifications *AcceptModificationsFilterer) WatchLogNftWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *AcceptModificationsLogNftWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _AcceptModifications.contract.WatchLogs(opts, "LogNftWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcceptModificationsLogNftWithdrawalAllowed)
				if err := _AcceptModifications.contract.UnpackLog(event, "LogNftWithdrawalAllowed", log); err != nil {
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

// ParseLogNftWithdrawalAllowed is a log parse operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 ownerKey, uint256 assetId)
func (_AcceptModifications *AcceptModificationsFilterer) ParseLogNftWithdrawalAllowed(log types.Log) (*AcceptModificationsLogNftWithdrawalAllowed, error) {
	event := new(AcceptModificationsLogNftWithdrawalAllowed)
	if err := _AcceptModifications.contract.UnpackLog(event, "LogNftWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AcceptModificationsLogWithdrawalAllowedIterator is returned from FilterLogWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogWithdrawalAllowed events raised by the AcceptModifications contract.
type AcceptModificationsLogWithdrawalAllowedIterator struct {
	Event *AcceptModificationsLogWithdrawalAllowed // Event containing the contract specifics and raw log

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
func (it *AcceptModificationsLogWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcceptModificationsLogWithdrawalAllowed)
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
		it.Event = new(AcceptModificationsLogWithdrawalAllowed)
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
func (it *AcceptModificationsLogWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcceptModificationsLogWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcceptModificationsLogWithdrawalAllowed represents a LogWithdrawalAllowed event raised by the AcceptModifications contract.
type AcceptModificationsLogWithdrawalAllowed struct {
	OwnerKey           *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalAllowed is a free log retrieval operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_AcceptModifications *AcceptModificationsFilterer) FilterLogWithdrawalAllowed(opts *bind.FilterOpts) (*AcceptModificationsLogWithdrawalAllowedIterator, error) {

	logs, sub, err := _AcceptModifications.contract.FilterLogs(opts, "LogWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &AcceptModificationsLogWithdrawalAllowedIterator{contract: _AcceptModifications.contract, event: "LogWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalAllowed is a free log subscription operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_AcceptModifications *AcceptModificationsFilterer) WatchLogWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *AcceptModificationsLogWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _AcceptModifications.contract.WatchLogs(opts, "LogWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcceptModificationsLogWithdrawalAllowed)
				if err := _AcceptModifications.contract.UnpackLog(event, "LogWithdrawalAllowed", log); err != nil {
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

// ParseLogWithdrawalAllowed is a log parse operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_AcceptModifications *AcceptModificationsFilterer) ParseLogWithdrawalAllowed(log types.Log) (*AcceptModificationsLogWithdrawalAllowed, error) {
	event := new(AcceptModificationsLogWithdrawalAllowed)
	if err := _AcceptModifications.contract.UnpackLog(event, "LogWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
