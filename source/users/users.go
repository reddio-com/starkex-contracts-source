// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package users

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

// UsersMetaData contains all meta data concerning the Users contract.
var UsersMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LogUserRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FREEZE_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FORCED_ACTIONS_REQS_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_VERIFIER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNFREEZE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_REMOVAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"starkSignature\",\"type\":\"bytes\"}],\"name\":\"registerEthAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"starkSignature\",\"type\":\"bytes\"}],\"name\":\"registerSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UsersABI is the input ABI used to generate the binding from.
// Deprecated: Use UsersMetaData.ABI instead.
var UsersABI = UsersMetaData.ABI

// Users is an auto generated Go binding around an Ethereum contract.
type Users struct {
	UsersCaller     // Read-only binding to the contract
	UsersTransactor // Write-only binding to the contract
	UsersFilterer   // Log filterer for contract events
}

// UsersCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsersSession struct {
	Contract     *Users            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsersCallerSession struct {
	Contract *UsersCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UsersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsersTransactorSession struct {
	Contract     *UsersTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsersRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsersRaw struct {
	Contract *Users // Generic contract binding to access the raw methods on
}

// UsersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsersCallerRaw struct {
	Contract *UsersCaller // Generic read-only contract binding to access the raw methods on
}

// UsersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsersTransactorRaw struct {
	Contract *UsersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsers creates a new instance of Users, bound to a specific deployed contract.
func NewUsers(address common.Address, backend bind.ContractBackend) (*Users, error) {
	contract, err := bindUsers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Users{UsersCaller: UsersCaller{contract: contract}, UsersTransactor: UsersTransactor{contract: contract}, UsersFilterer: UsersFilterer{contract: contract}}, nil
}

// NewUsersCaller creates a new read-only instance of Users, bound to a specific deployed contract.
func NewUsersCaller(address common.Address, caller bind.ContractCaller) (*UsersCaller, error) {
	contract, err := bindUsers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsersCaller{contract: contract}, nil
}

// NewUsersTransactor creates a new write-only instance of Users, bound to a specific deployed contract.
func NewUsersTransactor(address common.Address, transactor bind.ContractTransactor) (*UsersTransactor, error) {
	contract, err := bindUsers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsersTransactor{contract: contract}, nil
}

// NewUsersFilterer creates a new log filterer instance of Users, bound to a specific deployed contract.
func NewUsersFilterer(address common.Address, filterer bind.ContractFilterer) (*UsersFilterer, error) {
	contract, err := bindUsers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsersFilterer{contract: contract}, nil
}

// bindUsers binds a generic wrapper to an already deployed contract.
func bindUsers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Users *UsersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Users.Contract.UsersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Users *UsersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Users.Contract.UsersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Users *UsersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Users.Contract.UsersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Users *UsersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Users.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Users *UsersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Users.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Users *UsersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Users.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_Users *UsersCaller) DEPOSITCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Users.contract.Call(opts, &out, "DEPOSIT_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_Users *UsersSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _Users.Contract.DEPOSITCANCELDELAY(&_Users.CallOpts)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_Users *UsersCallerSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _Users.Contract.DEPOSITCANCELDELAY(&_Users.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_Users *UsersCaller) FREEZEGRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Users.contract.Call(opts, &out, "FREEZE_GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_Users *UsersSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _Users.Contract.FREEZEGRACEPERIOD(&_Users.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_Users *UsersCallerSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _Users.Contract.FREEZEGRACEPERIOD(&_Users.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_Users *UsersCaller) MAXFORCEDACTIONSREQSPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Users.contract.Call(opts, &out, "MAX_FORCED_ACTIONS_REQS_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_Users *UsersSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _Users.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_Users.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_Users *UsersCallerSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _Users.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_Users.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_Users *UsersCaller) MAXVERIFIERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Users.contract.Call(opts, &out, "MAX_VERIFIER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_Users *UsersSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _Users.Contract.MAXVERIFIERCOUNT(&_Users.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_Users *UsersCallerSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _Users.Contract.MAXVERIFIERCOUNT(&_Users.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_Users *UsersCaller) UNFREEZEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Users.contract.Call(opts, &out, "UNFREEZE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_Users *UsersSession) UNFREEZEDELAY() (*big.Int, error) {
	return _Users.Contract.UNFREEZEDELAY(&_Users.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_Users *UsersCallerSession) UNFREEZEDELAY() (*big.Int, error) {
	return _Users.Contract.UNFREEZEDELAY(&_Users.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_Users *UsersCaller) VERIFIERREMOVALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Users.contract.Call(opts, &out, "VERIFIER_REMOVAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_Users *UsersSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _Users.Contract.VERIFIERREMOVALDELAY(&_Users.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_Users *UsersCallerSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _Users.Contract.VERIFIERREMOVALDELAY(&_Users.CallOpts)
}

// RegisterEthAddress is a paid mutator transaction binding the contract method 0xbea84187.
//
// Solidity: function registerEthAddress(address ethKey, uint256 starkKey, bytes starkSignature) returns()
func (_Users *UsersTransactor) RegisterEthAddress(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "registerEthAddress", ethKey, starkKey, starkSignature)
}

// RegisterEthAddress is a paid mutator transaction binding the contract method 0xbea84187.
//
// Solidity: function registerEthAddress(address ethKey, uint256 starkKey, bytes starkSignature) returns()
func (_Users *UsersSession) RegisterEthAddress(ethKey common.Address, starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _Users.Contract.RegisterEthAddress(&_Users.TransactOpts, ethKey, starkKey, starkSignature)
}

// RegisterEthAddress is a paid mutator transaction binding the contract method 0xbea84187.
//
// Solidity: function registerEthAddress(address ethKey, uint256 starkKey, bytes starkSignature) returns()
func (_Users *UsersTransactorSession) RegisterEthAddress(ethKey common.Address, starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _Users.Contract.RegisterEthAddress(&_Users.TransactOpts, ethKey, starkKey, starkSignature)
}

// RegisterSender is a paid mutator transaction binding the contract method 0x86aeb445.
//
// Solidity: function registerSender(uint256 starkKey, bytes starkSignature) returns()
func (_Users *UsersTransactor) RegisterSender(opts *bind.TransactOpts, starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "registerSender", starkKey, starkSignature)
}

// RegisterSender is a paid mutator transaction binding the contract method 0x86aeb445.
//
// Solidity: function registerSender(uint256 starkKey, bytes starkSignature) returns()
func (_Users *UsersSession) RegisterSender(starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _Users.Contract.RegisterSender(&_Users.TransactOpts, starkKey, starkSignature)
}

// RegisterSender is a paid mutator transaction binding the contract method 0x86aeb445.
//
// Solidity: function registerSender(uint256 starkKey, bytes starkSignature) returns()
func (_Users *UsersTransactorSession) RegisterSender(starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _Users.Contract.RegisterSender(&_Users.TransactOpts, starkKey, starkSignature)
}

// UsersLogUserRegisteredIterator is returned from FilterLogUserRegistered and is used to iterate over the raw logs and unpacked data for LogUserRegistered events raised by the Users contract.
type UsersLogUserRegisteredIterator struct {
	Event *UsersLogUserRegistered // Event containing the contract specifics and raw log

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
func (it *UsersLogUserRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsersLogUserRegistered)
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
		it.Event = new(UsersLogUserRegistered)
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
func (it *UsersLogUserRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersLogUserRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersLogUserRegistered represents a LogUserRegistered event raised by the Users contract.
type UsersLogUserRegistered struct {
	EthKey   common.Address
	StarkKey *big.Int
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogUserRegistered is a free log retrieval operation binding the contract event 0xcab1cf17c190e4e2195a7b8f7b362023246fa774390432b4704ab6b29d56b07b.
//
// Solidity: event LogUserRegistered(address ethKey, uint256 starkKey, address sender)
func (_Users *UsersFilterer) FilterLogUserRegistered(opts *bind.FilterOpts) (*UsersLogUserRegisteredIterator, error) {

	logs, sub, err := _Users.contract.FilterLogs(opts, "LogUserRegistered")
	if err != nil {
		return nil, err
	}
	return &UsersLogUserRegisteredIterator{contract: _Users.contract, event: "LogUserRegistered", logs: logs, sub: sub}, nil
}

// WatchLogUserRegistered is a free log subscription operation binding the contract event 0xcab1cf17c190e4e2195a7b8f7b362023246fa774390432b4704ab6b29d56b07b.
//
// Solidity: event LogUserRegistered(address ethKey, uint256 starkKey, address sender)
func (_Users *UsersFilterer) WatchLogUserRegistered(opts *bind.WatchOpts, sink chan<- *UsersLogUserRegistered) (event.Subscription, error) {

	logs, sub, err := _Users.contract.WatchLogs(opts, "LogUserRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsersLogUserRegistered)
				if err := _Users.contract.UnpackLog(event, "LogUserRegistered", log); err != nil {
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

// ParseLogUserRegistered is a log parse operation binding the contract event 0xcab1cf17c190e4e2195a7b8f7b362023246fa774390432b4704ab6b29d56b07b.
//
// Solidity: event LogUserRegistered(address ethKey, uint256 starkKey, address sender)
func (_Users *UsersFilterer) ParseLogUserRegistered(log types.Log) (*UsersLogUserRegistered, error) {
	event := new(UsersLogUserRegistered)
	if err := _Users.contract.UnpackLog(event, "LogUserRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
