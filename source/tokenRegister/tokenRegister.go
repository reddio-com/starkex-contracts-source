// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenRegister

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

// TokenRegisterMetaData contains all meta data concerning the TokenRegister contract.
var TokenRegisterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAdmin\",\"type\":\"address\"}],\"name\":\"LogTokenAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAdmin\",\"type\":\"address\"}],\"name\":\"LogTokenAdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"name\":\"LogTokenRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FREEZE_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FORCED_ACTIONS_REQS_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_VERIFIER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNFREEZE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_REMOVAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"calculateAssetIdWithTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"mintingBlob\",\"type\":\"bytes\"}],\"name\":\"calculateMintableAssetId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"getAssetInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"isAssetRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"testedAdmin\",\"type\":\"address\"}],\"name\":\"isTokenAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"registerTokenAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"}],\"name\":\"unregisterTokenAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenRegisterABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenRegisterMetaData.ABI instead.
var TokenRegisterABI = TokenRegisterMetaData.ABI

// TokenRegister is an auto generated Go binding around an Ethereum contract.
type TokenRegister struct {
	TokenRegisterCaller     // Read-only binding to the contract
	TokenRegisterTransactor // Write-only binding to the contract
	TokenRegisterFilterer   // Log filterer for contract events
}

// TokenRegisterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenRegisterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegisterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenRegisterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegisterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenRegisterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegisterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenRegisterSession struct {
	Contract     *TokenRegister    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRegisterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenRegisterCallerSession struct {
	Contract *TokenRegisterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenRegisterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenRegisterTransactorSession struct {
	Contract     *TokenRegisterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenRegisterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRegisterRaw struct {
	Contract *TokenRegister // Generic contract binding to access the raw methods on
}

// TokenRegisterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenRegisterCallerRaw struct {
	Contract *TokenRegisterCaller // Generic read-only contract binding to access the raw methods on
}

// TokenRegisterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenRegisterTransactorRaw struct {
	Contract *TokenRegisterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenRegister creates a new instance of TokenRegister, bound to a specific deployed contract.
func NewTokenRegister(address common.Address, backend bind.ContractBackend) (*TokenRegister, error) {
	contract, err := bindTokenRegister(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenRegister{TokenRegisterCaller: TokenRegisterCaller{contract: contract}, TokenRegisterTransactor: TokenRegisterTransactor{contract: contract}, TokenRegisterFilterer: TokenRegisterFilterer{contract: contract}}, nil
}

// NewTokenRegisterCaller creates a new read-only instance of TokenRegister, bound to a specific deployed contract.
func NewTokenRegisterCaller(address common.Address, caller bind.ContractCaller) (*TokenRegisterCaller, error) {
	contract, err := bindTokenRegister(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRegisterCaller{contract: contract}, nil
}

// NewTokenRegisterTransactor creates a new write-only instance of TokenRegister, bound to a specific deployed contract.
func NewTokenRegisterTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenRegisterTransactor, error) {
	contract, err := bindTokenRegister(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRegisterTransactor{contract: contract}, nil
}

// NewTokenRegisterFilterer creates a new log filterer instance of TokenRegister, bound to a specific deployed contract.
func NewTokenRegisterFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenRegisterFilterer, error) {
	contract, err := bindTokenRegister(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenRegisterFilterer{contract: contract}, nil
}

// bindTokenRegister binds a generic wrapper to an already deployed contract.
func bindTokenRegister(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenRegisterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRegister *TokenRegisterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenRegister.Contract.TokenRegisterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRegister *TokenRegisterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRegister.Contract.TokenRegisterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRegister *TokenRegisterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRegister.Contract.TokenRegisterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRegister *TokenRegisterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenRegister.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRegister *TokenRegisterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRegister.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRegister *TokenRegisterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRegister.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_TokenRegister *TokenRegisterCaller) DEPOSITCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenRegister.contract.Call(opts, &out, "DEPOSIT_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_TokenRegister *TokenRegisterSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _TokenRegister.Contract.DEPOSITCANCELDELAY(&_TokenRegister.CallOpts)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_TokenRegister *TokenRegisterCallerSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _TokenRegister.Contract.DEPOSITCANCELDELAY(&_TokenRegister.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_TokenRegister *TokenRegisterCaller) FREEZEGRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenRegister.contract.Call(opts, &out, "FREEZE_GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_TokenRegister *TokenRegisterSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _TokenRegister.Contract.FREEZEGRACEPERIOD(&_TokenRegister.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_TokenRegister *TokenRegisterCallerSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _TokenRegister.Contract.FREEZEGRACEPERIOD(&_TokenRegister.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_TokenRegister *TokenRegisterCaller) MAXFORCEDACTIONSREQSPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenRegister.contract.Call(opts, &out, "MAX_FORCED_ACTIONS_REQS_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_TokenRegister *TokenRegisterSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _TokenRegister.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_TokenRegister.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_TokenRegister *TokenRegisterCallerSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _TokenRegister.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_TokenRegister.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_TokenRegister *TokenRegisterCaller) MAXVERIFIERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenRegister.contract.Call(opts, &out, "MAX_VERIFIER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_TokenRegister *TokenRegisterSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _TokenRegister.Contract.MAXVERIFIERCOUNT(&_TokenRegister.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_TokenRegister *TokenRegisterCallerSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _TokenRegister.Contract.MAXVERIFIERCOUNT(&_TokenRegister.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_TokenRegister *TokenRegisterCaller) UNFREEZEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenRegister.contract.Call(opts, &out, "UNFREEZE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_TokenRegister *TokenRegisterSession) UNFREEZEDELAY() (*big.Int, error) {
	return _TokenRegister.Contract.UNFREEZEDELAY(&_TokenRegister.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_TokenRegister *TokenRegisterCallerSession) UNFREEZEDELAY() (*big.Int, error) {
	return _TokenRegister.Contract.UNFREEZEDELAY(&_TokenRegister.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_TokenRegister *TokenRegisterCaller) VERIFIERREMOVALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenRegister.contract.Call(opts, &out, "VERIFIER_REMOVAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_TokenRegister *TokenRegisterSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _TokenRegister.Contract.VERIFIERREMOVALDELAY(&_TokenRegister.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_TokenRegister *TokenRegisterCallerSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _TokenRegister.Contract.VERIFIERREMOVALDELAY(&_TokenRegister.CallOpts)
}

// CalculateAssetIdWithTokenId is a free data retrieval call binding the contract method 0xa1cc5e13.
//
// Solidity: function calculateAssetIdWithTokenId(uint256 assetType, uint256 tokenId) view returns(uint256)
func (_TokenRegister *TokenRegisterCaller) CalculateAssetIdWithTokenId(opts *bind.CallOpts, assetType *big.Int, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenRegister.contract.Call(opts, &out, "calculateAssetIdWithTokenId", assetType, tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateAssetIdWithTokenId is a free data retrieval call binding the contract method 0xa1cc5e13.
//
// Solidity: function calculateAssetIdWithTokenId(uint256 assetType, uint256 tokenId) view returns(uint256)
func (_TokenRegister *TokenRegisterSession) CalculateAssetIdWithTokenId(assetType *big.Int, tokenId *big.Int) (*big.Int, error) {
	return _TokenRegister.Contract.CalculateAssetIdWithTokenId(&_TokenRegister.CallOpts, assetType, tokenId)
}

// CalculateAssetIdWithTokenId is a free data retrieval call binding the contract method 0xa1cc5e13.
//
// Solidity: function calculateAssetIdWithTokenId(uint256 assetType, uint256 tokenId) view returns(uint256)
func (_TokenRegister *TokenRegisterCallerSession) CalculateAssetIdWithTokenId(assetType *big.Int, tokenId *big.Int) (*big.Int, error) {
	return _TokenRegister.Contract.CalculateAssetIdWithTokenId(&_TokenRegister.CallOpts, assetType, tokenId)
}

// CalculateMintableAssetId is a free data retrieval call binding the contract method 0xb12773fb.
//
// Solidity: function calculateMintableAssetId(uint256 assetType, bytes mintingBlob) pure returns(uint256)
func (_TokenRegister *TokenRegisterCaller) CalculateMintableAssetId(opts *bind.CallOpts, assetType *big.Int, mintingBlob []byte) (*big.Int, error) {
	var out []interface{}
	err := _TokenRegister.contract.Call(opts, &out, "calculateMintableAssetId", assetType, mintingBlob)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateMintableAssetId is a free data retrieval call binding the contract method 0xb12773fb.
//
// Solidity: function calculateMintableAssetId(uint256 assetType, bytes mintingBlob) pure returns(uint256)
func (_TokenRegister *TokenRegisterSession) CalculateMintableAssetId(assetType *big.Int, mintingBlob []byte) (*big.Int, error) {
	return _TokenRegister.Contract.CalculateMintableAssetId(&_TokenRegister.CallOpts, assetType, mintingBlob)
}

// CalculateMintableAssetId is a free data retrieval call binding the contract method 0xb12773fb.
//
// Solidity: function calculateMintableAssetId(uint256 assetType, bytes mintingBlob) pure returns(uint256)
func (_TokenRegister *TokenRegisterCallerSession) CalculateMintableAssetId(assetType *big.Int, mintingBlob []byte) (*big.Int, error) {
	return _TokenRegister.Contract.CalculateMintableAssetId(&_TokenRegister.CallOpts, assetType, mintingBlob)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes)
func (_TokenRegister *TokenRegisterCaller) GetAssetInfo(opts *bind.CallOpts, assetType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _TokenRegister.contract.Call(opts, &out, "getAssetInfo", assetType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes)
func (_TokenRegister *TokenRegisterSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _TokenRegister.Contract.GetAssetInfo(&_TokenRegister.CallOpts, assetType)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes)
func (_TokenRegister *TokenRegisterCallerSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _TokenRegister.Contract.GetAssetInfo(&_TokenRegister.CallOpts, assetType)
}

// IsAssetRegistered is a free data retrieval call binding the contract method 0x049f5ade.
//
// Solidity: function isAssetRegistered(uint256 assetType) view returns(bool)
func (_TokenRegister *TokenRegisterCaller) IsAssetRegistered(opts *bind.CallOpts, assetType *big.Int) (bool, error) {
	var out []interface{}
	err := _TokenRegister.contract.Call(opts, &out, "isAssetRegistered", assetType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetRegistered is a free data retrieval call binding the contract method 0x049f5ade.
//
// Solidity: function isAssetRegistered(uint256 assetType) view returns(bool)
func (_TokenRegister *TokenRegisterSession) IsAssetRegistered(assetType *big.Int) (bool, error) {
	return _TokenRegister.Contract.IsAssetRegistered(&_TokenRegister.CallOpts, assetType)
}

// IsAssetRegistered is a free data retrieval call binding the contract method 0x049f5ade.
//
// Solidity: function isAssetRegistered(uint256 assetType) view returns(bool)
func (_TokenRegister *TokenRegisterCallerSession) IsAssetRegistered(assetType *big.Int) (bool, error) {
	return _TokenRegister.Contract.IsAssetRegistered(&_TokenRegister.CallOpts, assetType)
}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_TokenRegister *TokenRegisterCaller) IsTokenAdmin(opts *bind.CallOpts, testedAdmin common.Address) (bool, error) {
	var out []interface{}
	err := _TokenRegister.contract.Call(opts, &out, "isTokenAdmin", testedAdmin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_TokenRegister *TokenRegisterSession) IsTokenAdmin(testedAdmin common.Address) (bool, error) {
	return _TokenRegister.Contract.IsTokenAdmin(&_TokenRegister.CallOpts, testedAdmin)
}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_TokenRegister *TokenRegisterCallerSession) IsTokenAdmin(testedAdmin common.Address) (bool, error) {
	return _TokenRegister.Contract.IsTokenAdmin(&_TokenRegister.CallOpts, testedAdmin)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo) returns()
func (_TokenRegister *TokenRegisterTransactor) RegisterToken(opts *bind.TransactOpts, assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _TokenRegister.contract.Transact(opts, "registerToken", assetType, assetInfo)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo) returns()
func (_TokenRegister *TokenRegisterSession) RegisterToken(assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _TokenRegister.Contract.RegisterToken(&_TokenRegister.TransactOpts, assetType, assetInfo)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo) returns()
func (_TokenRegister *TokenRegisterTransactorSession) RegisterToken(assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _TokenRegister.Contract.RegisterToken(&_TokenRegister.TransactOpts, assetType, assetInfo)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo, uint256 quantum) returns()
func (_TokenRegister *TokenRegisterTransactor) RegisterToken0(opts *bind.TransactOpts, assetType *big.Int, assetInfo []byte, quantum *big.Int) (*types.Transaction, error) {
	return _TokenRegister.contract.Transact(opts, "registerToken0", assetType, assetInfo, quantum)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo, uint256 quantum) returns()
func (_TokenRegister *TokenRegisterSession) RegisterToken0(assetType *big.Int, assetInfo []byte, quantum *big.Int) (*types.Transaction, error) {
	return _TokenRegister.Contract.RegisterToken0(&_TokenRegister.TransactOpts, assetType, assetInfo, quantum)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo, uint256 quantum) returns()
func (_TokenRegister *TokenRegisterTransactorSession) RegisterToken0(assetType *big.Int, assetInfo []byte, quantum *big.Int) (*types.Transaction, error) {
	return _TokenRegister.Contract.RegisterToken0(&_TokenRegister.TransactOpts, assetType, assetInfo, quantum)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_TokenRegister *TokenRegisterTransactor) RegisterTokenAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _TokenRegister.contract.Transact(opts, "registerTokenAdmin", newAdmin)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_TokenRegister *TokenRegisterSession) RegisterTokenAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _TokenRegister.Contract.RegisterTokenAdmin(&_TokenRegister.TransactOpts, newAdmin)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_TokenRegister *TokenRegisterTransactorSession) RegisterTokenAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _TokenRegister.Contract.RegisterTokenAdmin(&_TokenRegister.TransactOpts, newAdmin)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_TokenRegister *TokenRegisterTransactor) UnregisterTokenAdmin(opts *bind.TransactOpts, oldAdmin common.Address) (*types.Transaction, error) {
	return _TokenRegister.contract.Transact(opts, "unregisterTokenAdmin", oldAdmin)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_TokenRegister *TokenRegisterSession) UnregisterTokenAdmin(oldAdmin common.Address) (*types.Transaction, error) {
	return _TokenRegister.Contract.UnregisterTokenAdmin(&_TokenRegister.TransactOpts, oldAdmin)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_TokenRegister *TokenRegisterTransactorSession) UnregisterTokenAdmin(oldAdmin common.Address) (*types.Transaction, error) {
	return _TokenRegister.Contract.UnregisterTokenAdmin(&_TokenRegister.TransactOpts, oldAdmin)
}

// TokenRegisterLogTokenAdminAddedIterator is returned from FilterLogTokenAdminAdded and is used to iterate over the raw logs and unpacked data for LogTokenAdminAdded events raised by the TokenRegister contract.
type TokenRegisterLogTokenAdminAddedIterator struct {
	Event *TokenRegisterLogTokenAdminAdded // Event containing the contract specifics and raw log

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
func (it *TokenRegisterLogTokenAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegisterLogTokenAdminAdded)
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
		it.Event = new(TokenRegisterLogTokenAdminAdded)
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
func (it *TokenRegisterLogTokenAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegisterLogTokenAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegisterLogTokenAdminAdded represents a LogTokenAdminAdded event raised by the TokenRegister contract.
type TokenRegisterLogTokenAdminAdded struct {
	TokenAdmin common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogTokenAdminAdded is a free log retrieval operation binding the contract event 0x9085a9044aeb6daeeb5b4bf84af42b1a1613d4056f503c4e992b6396c16bd52f.
//
// Solidity: event LogTokenAdminAdded(address tokenAdmin)
func (_TokenRegister *TokenRegisterFilterer) FilterLogTokenAdminAdded(opts *bind.FilterOpts) (*TokenRegisterLogTokenAdminAddedIterator, error) {

	logs, sub, err := _TokenRegister.contract.FilterLogs(opts, "LogTokenAdminAdded")
	if err != nil {
		return nil, err
	}
	return &TokenRegisterLogTokenAdminAddedIterator{contract: _TokenRegister.contract, event: "LogTokenAdminAdded", logs: logs, sub: sub}, nil
}

// WatchLogTokenAdminAdded is a free log subscription operation binding the contract event 0x9085a9044aeb6daeeb5b4bf84af42b1a1613d4056f503c4e992b6396c16bd52f.
//
// Solidity: event LogTokenAdminAdded(address tokenAdmin)
func (_TokenRegister *TokenRegisterFilterer) WatchLogTokenAdminAdded(opts *bind.WatchOpts, sink chan<- *TokenRegisterLogTokenAdminAdded) (event.Subscription, error) {

	logs, sub, err := _TokenRegister.contract.WatchLogs(opts, "LogTokenAdminAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegisterLogTokenAdminAdded)
				if err := _TokenRegister.contract.UnpackLog(event, "LogTokenAdminAdded", log); err != nil {
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

// ParseLogTokenAdminAdded is a log parse operation binding the contract event 0x9085a9044aeb6daeeb5b4bf84af42b1a1613d4056f503c4e992b6396c16bd52f.
//
// Solidity: event LogTokenAdminAdded(address tokenAdmin)
func (_TokenRegister *TokenRegisterFilterer) ParseLogTokenAdminAdded(log types.Log) (*TokenRegisterLogTokenAdminAdded, error) {
	event := new(TokenRegisterLogTokenAdminAdded)
	if err := _TokenRegister.contract.UnpackLog(event, "LogTokenAdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRegisterLogTokenAdminRemovedIterator is returned from FilterLogTokenAdminRemoved and is used to iterate over the raw logs and unpacked data for LogTokenAdminRemoved events raised by the TokenRegister contract.
type TokenRegisterLogTokenAdminRemovedIterator struct {
	Event *TokenRegisterLogTokenAdminRemoved // Event containing the contract specifics and raw log

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
func (it *TokenRegisterLogTokenAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegisterLogTokenAdminRemoved)
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
		it.Event = new(TokenRegisterLogTokenAdminRemoved)
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
func (it *TokenRegisterLogTokenAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegisterLogTokenAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegisterLogTokenAdminRemoved represents a LogTokenAdminRemoved event raised by the TokenRegister contract.
type TokenRegisterLogTokenAdminRemoved struct {
	TokenAdmin common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogTokenAdminRemoved is a free log retrieval operation binding the contract event 0xfa49aecb996ea8d99950bb051552dfcc0b5460a0bb209867a1ed8067c32c2177.
//
// Solidity: event LogTokenAdminRemoved(address tokenAdmin)
func (_TokenRegister *TokenRegisterFilterer) FilterLogTokenAdminRemoved(opts *bind.FilterOpts) (*TokenRegisterLogTokenAdminRemovedIterator, error) {

	logs, sub, err := _TokenRegister.contract.FilterLogs(opts, "LogTokenAdminRemoved")
	if err != nil {
		return nil, err
	}
	return &TokenRegisterLogTokenAdminRemovedIterator{contract: _TokenRegister.contract, event: "LogTokenAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchLogTokenAdminRemoved is a free log subscription operation binding the contract event 0xfa49aecb996ea8d99950bb051552dfcc0b5460a0bb209867a1ed8067c32c2177.
//
// Solidity: event LogTokenAdminRemoved(address tokenAdmin)
func (_TokenRegister *TokenRegisterFilterer) WatchLogTokenAdminRemoved(opts *bind.WatchOpts, sink chan<- *TokenRegisterLogTokenAdminRemoved) (event.Subscription, error) {

	logs, sub, err := _TokenRegister.contract.WatchLogs(opts, "LogTokenAdminRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegisterLogTokenAdminRemoved)
				if err := _TokenRegister.contract.UnpackLog(event, "LogTokenAdminRemoved", log); err != nil {
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

// ParseLogTokenAdminRemoved is a log parse operation binding the contract event 0xfa49aecb996ea8d99950bb051552dfcc0b5460a0bb209867a1ed8067c32c2177.
//
// Solidity: event LogTokenAdminRemoved(address tokenAdmin)
func (_TokenRegister *TokenRegisterFilterer) ParseLogTokenAdminRemoved(log types.Log) (*TokenRegisterLogTokenAdminRemoved, error) {
	event := new(TokenRegisterLogTokenAdminRemoved)
	if err := _TokenRegister.contract.UnpackLog(event, "LogTokenAdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRegisterLogTokenRegisteredIterator is returned from FilterLogTokenRegistered and is used to iterate over the raw logs and unpacked data for LogTokenRegistered events raised by the TokenRegister contract.
type TokenRegisterLogTokenRegisteredIterator struct {
	Event *TokenRegisterLogTokenRegistered // Event containing the contract specifics and raw log

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
func (it *TokenRegisterLogTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegisterLogTokenRegistered)
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
		it.Event = new(TokenRegisterLogTokenRegistered)
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
func (it *TokenRegisterLogTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegisterLogTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegisterLogTokenRegistered represents a LogTokenRegistered event raised by the TokenRegister contract.
type TokenRegisterLogTokenRegistered struct {
	AssetType *big.Int
	AssetInfo []byte
	Quantum   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogTokenRegistered is a free log retrieval operation binding the contract event 0x7a0efbc885500f3b4a895231945be4520e4c0ba5ef7274a225a0272c81ccbcb7.
//
// Solidity: event LogTokenRegistered(uint256 assetType, bytes assetInfo, uint256 quantum)
func (_TokenRegister *TokenRegisterFilterer) FilterLogTokenRegistered(opts *bind.FilterOpts) (*TokenRegisterLogTokenRegisteredIterator, error) {

	logs, sub, err := _TokenRegister.contract.FilterLogs(opts, "LogTokenRegistered")
	if err != nil {
		return nil, err
	}
	return &TokenRegisterLogTokenRegisteredIterator{contract: _TokenRegister.contract, event: "LogTokenRegistered", logs: logs, sub: sub}, nil
}

// WatchLogTokenRegistered is a free log subscription operation binding the contract event 0x7a0efbc885500f3b4a895231945be4520e4c0ba5ef7274a225a0272c81ccbcb7.
//
// Solidity: event LogTokenRegistered(uint256 assetType, bytes assetInfo, uint256 quantum)
func (_TokenRegister *TokenRegisterFilterer) WatchLogTokenRegistered(opts *bind.WatchOpts, sink chan<- *TokenRegisterLogTokenRegistered) (event.Subscription, error) {

	logs, sub, err := _TokenRegister.contract.WatchLogs(opts, "LogTokenRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegisterLogTokenRegistered)
				if err := _TokenRegister.contract.UnpackLog(event, "LogTokenRegistered", log); err != nil {
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

// ParseLogTokenRegistered is a log parse operation binding the contract event 0x7a0efbc885500f3b4a895231945be4520e4c0ba5ef7274a225a0272c81ccbcb7.
//
// Solidity: event LogTokenRegistered(uint256 assetType, bytes assetInfo, uint256 quantum)
func (_TokenRegister *TokenRegisterFilterer) ParseLogTokenRegistered(log types.Log) (*TokenRegisterLogTokenRegistered, error) {
	event := new(TokenRegisterLogTokenRegistered)
	if err := _TokenRegister.contract.UnpackLog(event, "LogTokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
