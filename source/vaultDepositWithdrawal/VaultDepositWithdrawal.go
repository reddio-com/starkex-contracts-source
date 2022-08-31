// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaultDepositWithdrawal

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

// VaultDepositWithdrawalMetaData contains all meta data concerning the VaultDepositWithdrawal contract.
var VaultDepositWithdrawalMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDepositToVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawalFromVault\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"calculateAssetIdWithTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"mintingBlob\",\"type\":\"bytes\"}],\"name\":\"calculateMintableAssetId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultVaultWithdrawalLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"depositERC1155ToVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"depositERC20ToVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositEthToVault\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"getAssetInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getQuantizedErc1155VaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getQuantizedVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"presumedAssetType\",\"type\":\"uint256\"}],\"name\":\"getQuantum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"isVaultLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawErc1155FromVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VaultDepositWithdrawalABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultDepositWithdrawalMetaData.ABI instead.
var VaultDepositWithdrawalABI = VaultDepositWithdrawalMetaData.ABI

// VaultDepositWithdrawal is an auto generated Go binding around an Ethereum contract.
type VaultDepositWithdrawal struct {
	VaultDepositWithdrawalCaller     // Read-only binding to the contract
	VaultDepositWithdrawalTransactor // Write-only binding to the contract
	VaultDepositWithdrawalFilterer   // Log filterer for contract events
}

// VaultDepositWithdrawalCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultDepositWithdrawalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultDepositWithdrawalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultDepositWithdrawalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultDepositWithdrawalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultDepositWithdrawalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultDepositWithdrawalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultDepositWithdrawalSession struct {
	Contract     *VaultDepositWithdrawal // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VaultDepositWithdrawalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultDepositWithdrawalCallerSession struct {
	Contract *VaultDepositWithdrawalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// VaultDepositWithdrawalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultDepositWithdrawalTransactorSession struct {
	Contract     *VaultDepositWithdrawalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// VaultDepositWithdrawalRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultDepositWithdrawalRaw struct {
	Contract *VaultDepositWithdrawal // Generic contract binding to access the raw methods on
}

// VaultDepositWithdrawalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultDepositWithdrawalCallerRaw struct {
	Contract *VaultDepositWithdrawalCaller // Generic read-only contract binding to access the raw methods on
}

// VaultDepositWithdrawalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultDepositWithdrawalTransactorRaw struct {
	Contract *VaultDepositWithdrawalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultDepositWithdrawal creates a new instance of VaultDepositWithdrawal, bound to a specific deployed contract.
func NewVaultDepositWithdrawal(address common.Address, backend bind.ContractBackend) (*VaultDepositWithdrawal, error) {
	contract, err := bindVaultDepositWithdrawal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultDepositWithdrawal{VaultDepositWithdrawalCaller: VaultDepositWithdrawalCaller{contract: contract}, VaultDepositWithdrawalTransactor: VaultDepositWithdrawalTransactor{contract: contract}, VaultDepositWithdrawalFilterer: VaultDepositWithdrawalFilterer{contract: contract}}, nil
}

// NewVaultDepositWithdrawalCaller creates a new read-only instance of VaultDepositWithdrawal, bound to a specific deployed contract.
func NewVaultDepositWithdrawalCaller(address common.Address, caller bind.ContractCaller) (*VaultDepositWithdrawalCaller, error) {
	contract, err := bindVaultDepositWithdrawal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultDepositWithdrawalCaller{contract: contract}, nil
}

// NewVaultDepositWithdrawalTransactor creates a new write-only instance of VaultDepositWithdrawal, bound to a specific deployed contract.
func NewVaultDepositWithdrawalTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultDepositWithdrawalTransactor, error) {
	contract, err := bindVaultDepositWithdrawal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultDepositWithdrawalTransactor{contract: contract}, nil
}

// NewVaultDepositWithdrawalFilterer creates a new log filterer instance of VaultDepositWithdrawal, bound to a specific deployed contract.
func NewVaultDepositWithdrawalFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultDepositWithdrawalFilterer, error) {
	contract, err := bindVaultDepositWithdrawal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultDepositWithdrawalFilterer{contract: contract}, nil
}

// bindVaultDepositWithdrawal binds a generic wrapper to an already deployed contract.
func bindVaultDepositWithdrawal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultDepositWithdrawalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultDepositWithdrawal *VaultDepositWithdrawalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultDepositWithdrawal.Contract.VaultDepositWithdrawalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultDepositWithdrawal *VaultDepositWithdrawalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.Contract.VaultDepositWithdrawalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultDepositWithdrawal *VaultDepositWithdrawalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.Contract.VaultDepositWithdrawalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultDepositWithdrawal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultDepositWithdrawal *VaultDepositWithdrawalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultDepositWithdrawal *VaultDepositWithdrawalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.Contract.contract.Transact(opts, method, params...)
}

// CalculateAssetIdWithTokenId is a free data retrieval call binding the contract method 0xa1cc5e13.
//
// Solidity: function calculateAssetIdWithTokenId(uint256 assetType, uint256 tokenId) view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCaller) CalculateAssetIdWithTokenId(opts *bind.CallOpts, assetType *big.Int, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultDepositWithdrawal.contract.Call(opts, &out, "calculateAssetIdWithTokenId", assetType, tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateAssetIdWithTokenId is a free data retrieval call binding the contract method 0xa1cc5e13.
//
// Solidity: function calculateAssetIdWithTokenId(uint256 assetType, uint256 tokenId) view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) CalculateAssetIdWithTokenId(assetType *big.Int, tokenId *big.Int) (*big.Int, error) {
	return _VaultDepositWithdrawal.Contract.CalculateAssetIdWithTokenId(&_VaultDepositWithdrawal.CallOpts, assetType, tokenId)
}

// CalculateAssetIdWithTokenId is a free data retrieval call binding the contract method 0xa1cc5e13.
//
// Solidity: function calculateAssetIdWithTokenId(uint256 assetType, uint256 tokenId) view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCallerSession) CalculateAssetIdWithTokenId(assetType *big.Int, tokenId *big.Int) (*big.Int, error) {
	return _VaultDepositWithdrawal.Contract.CalculateAssetIdWithTokenId(&_VaultDepositWithdrawal.CallOpts, assetType, tokenId)
}

// CalculateMintableAssetId is a free data retrieval call binding the contract method 0xb12773fb.
//
// Solidity: function calculateMintableAssetId(uint256 assetType, bytes mintingBlob) pure returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCaller) CalculateMintableAssetId(opts *bind.CallOpts, assetType *big.Int, mintingBlob []byte) (*big.Int, error) {
	var out []interface{}
	err := _VaultDepositWithdrawal.contract.Call(opts, &out, "calculateMintableAssetId", assetType, mintingBlob)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateMintableAssetId is a free data retrieval call binding the contract method 0xb12773fb.
//
// Solidity: function calculateMintableAssetId(uint256 assetType, bytes mintingBlob) pure returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) CalculateMintableAssetId(assetType *big.Int, mintingBlob []byte) (*big.Int, error) {
	return _VaultDepositWithdrawal.Contract.CalculateMintableAssetId(&_VaultDepositWithdrawal.CallOpts, assetType, mintingBlob)
}

// CalculateMintableAssetId is a free data retrieval call binding the contract method 0xb12773fb.
//
// Solidity: function calculateMintableAssetId(uint256 assetType, bytes mintingBlob) pure returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCallerSession) CalculateMintableAssetId(assetType *big.Int, mintingBlob []byte) (*big.Int, error) {
	return _VaultDepositWithdrawal.Contract.CalculateMintableAssetId(&_VaultDepositWithdrawal.CallOpts, assetType, mintingBlob)
}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCaller) DefaultVaultWithdrawalLock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultDepositWithdrawal.contract.Call(opts, &out, "defaultVaultWithdrawalLock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) DefaultVaultWithdrawalLock() (*big.Int, error) {
	return _VaultDepositWithdrawal.Contract.DefaultVaultWithdrawalLock(&_VaultDepositWithdrawal.CallOpts)
}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCallerSession) DefaultVaultWithdrawalLock() (*big.Int, error) {
	return _VaultDepositWithdrawal.Contract.DefaultVaultWithdrawalLock(&_VaultDepositWithdrawal.CallOpts)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCaller) GetAssetInfo(opts *bind.CallOpts, assetType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _VaultDepositWithdrawal.contract.Call(opts, &out, "getAssetInfo", assetType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _VaultDepositWithdrawal.Contract.GetAssetInfo(&_VaultDepositWithdrawal.CallOpts, assetType)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCallerSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _VaultDepositWithdrawal.Contract.GetAssetInfo(&_VaultDepositWithdrawal.CallOpts, assetType)
}

// GetQuantizedErc1155VaultBalance is a free data retrieval call binding the contract method 0xf7ac8587.
//
// Solidity: function getQuantizedErc1155VaultBalance(address ethKey, uint256 assetType, uint256 tokenId, uint256 vaultId) view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCaller) GetQuantizedErc1155VaultBalance(opts *bind.CallOpts, ethKey common.Address, assetType *big.Int, tokenId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultDepositWithdrawal.contract.Call(opts, &out, "getQuantizedErc1155VaultBalance", ethKey, assetType, tokenId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantizedErc1155VaultBalance is a free data retrieval call binding the contract method 0xf7ac8587.
//
// Solidity: function getQuantizedErc1155VaultBalance(address ethKey, uint256 assetType, uint256 tokenId, uint256 vaultId) view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) GetQuantizedErc1155VaultBalance(ethKey common.Address, assetType *big.Int, tokenId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _VaultDepositWithdrawal.Contract.GetQuantizedErc1155VaultBalance(&_VaultDepositWithdrawal.CallOpts, ethKey, assetType, tokenId, vaultId)
}

// GetQuantizedErc1155VaultBalance is a free data retrieval call binding the contract method 0xf7ac8587.
//
// Solidity: function getQuantizedErc1155VaultBalance(address ethKey, uint256 assetType, uint256 tokenId, uint256 vaultId) view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCallerSession) GetQuantizedErc1155VaultBalance(ethKey common.Address, assetType *big.Int, tokenId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _VaultDepositWithdrawal.Contract.GetQuantizedErc1155VaultBalance(&_VaultDepositWithdrawal.CallOpts, ethKey, assetType, tokenId, vaultId)
}

// GetQuantizedVaultBalance is a free data retrieval call binding the contract method 0xe69a9de7.
//
// Solidity: function getQuantizedVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCaller) GetQuantizedVaultBalance(opts *bind.CallOpts, ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultDepositWithdrawal.contract.Call(opts, &out, "getQuantizedVaultBalance", ethKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantizedVaultBalance is a free data retrieval call binding the contract method 0xe69a9de7.
//
// Solidity: function getQuantizedVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) GetQuantizedVaultBalance(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _VaultDepositWithdrawal.Contract.GetQuantizedVaultBalance(&_VaultDepositWithdrawal.CallOpts, ethKey, assetId, vaultId)
}

// GetQuantizedVaultBalance is a free data retrieval call binding the contract method 0xe69a9de7.
//
// Solidity: function getQuantizedVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCallerSession) GetQuantizedVaultBalance(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _VaultDepositWithdrawal.Contract.GetQuantizedVaultBalance(&_VaultDepositWithdrawal.CallOpts, ethKey, assetId, vaultId)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCaller) GetQuantum(opts *bind.CallOpts, presumedAssetType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultDepositWithdrawal.contract.Call(opts, &out, "getQuantum", presumedAssetType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _VaultDepositWithdrawal.Contract.GetQuantum(&_VaultDepositWithdrawal.CallOpts, presumedAssetType)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCallerSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _VaultDepositWithdrawal.Contract.GetQuantum(&_VaultDepositWithdrawal.CallOpts, presumedAssetType)
}

// GetVaultBalance is a free data retrieval call binding the contract method 0x1c6bd544.
//
// Solidity: function getVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCaller) GetVaultBalance(opts *bind.CallOpts, ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultDepositWithdrawal.contract.Call(opts, &out, "getVaultBalance", ethKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultBalance is a free data retrieval call binding the contract method 0x1c6bd544.
//
// Solidity: function getVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) GetVaultBalance(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _VaultDepositWithdrawal.Contract.GetVaultBalance(&_VaultDepositWithdrawal.CallOpts, ethKey, assetId, vaultId)
}

// GetVaultBalance is a free data retrieval call binding the contract method 0x1c6bd544.
//
// Solidity: function getVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCallerSession) GetVaultBalance(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _VaultDepositWithdrawal.Contract.GetVaultBalance(&_VaultDepositWithdrawal.CallOpts, ethKey, assetId, vaultId)
}

// IsVaultLocked is a free data retrieval call binding the contract method 0x27c2b36d.
//
// Solidity: function isVaultLocked(address ethKey, uint256 assetId, uint256 vaultId) view returns(bool)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCaller) IsVaultLocked(opts *bind.CallOpts, ethKey common.Address, assetId *big.Int, vaultId *big.Int) (bool, error) {
	var out []interface{}
	err := _VaultDepositWithdrawal.contract.Call(opts, &out, "isVaultLocked", ethKey, assetId, vaultId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVaultLocked is a free data retrieval call binding the contract method 0x27c2b36d.
//
// Solidity: function isVaultLocked(address ethKey, uint256 assetId, uint256 vaultId) view returns(bool)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) IsVaultLocked(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (bool, error) {
	return _VaultDepositWithdrawal.Contract.IsVaultLocked(&_VaultDepositWithdrawal.CallOpts, ethKey, assetId, vaultId)
}

// IsVaultLocked is a free data retrieval call binding the contract method 0x27c2b36d.
//
// Solidity: function isVaultLocked(address ethKey, uint256 assetId, uint256 vaultId) view returns(bool)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCallerSession) IsVaultLocked(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (bool, error) {
	return _VaultDepositWithdrawal.Contract.IsVaultLocked(&_VaultDepositWithdrawal.CallOpts, ethKey, assetId, vaultId)
}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCaller) OrderRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultDepositWithdrawal.contract.Call(opts, &out, "orderRegistryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) OrderRegistryAddress() (common.Address, error) {
	return _VaultDepositWithdrawal.Contract.OrderRegistryAddress(&_VaultDepositWithdrawal.CallOpts)
}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalCallerSession) OrderRegistryAddress() (common.Address, error) {
	return _VaultDepositWithdrawal.Contract.OrderRegistryAddress(&_VaultDepositWithdrawal.CallOpts)
}

// DepositERC1155ToVault is a paid mutator transaction binding the contract method 0xa70c4b3b.
//
// Solidity: function depositERC1155ToVault(uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalTransactor) DepositERC1155ToVault(opts *bind.TransactOpts, assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.contract.Transact(opts, "depositERC1155ToVault", assetType, tokenId, vaultId, quantizedAmount)
}

// DepositERC1155ToVault is a paid mutator transaction binding the contract method 0xa70c4b3b.
//
// Solidity: function depositERC1155ToVault(uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) DepositERC1155ToVault(assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.Contract.DepositERC1155ToVault(&_VaultDepositWithdrawal.TransactOpts, assetType, tokenId, vaultId, quantizedAmount)
}

// DepositERC1155ToVault is a paid mutator transaction binding the contract method 0xa70c4b3b.
//
// Solidity: function depositERC1155ToVault(uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalTransactorSession) DepositERC1155ToVault(assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.Contract.DepositERC1155ToVault(&_VaultDepositWithdrawal.TransactOpts, assetType, tokenId, vaultId, quantizedAmount)
}

// DepositERC20ToVault is a paid mutator transaction binding the contract method 0x1d133002.
//
// Solidity: function depositERC20ToVault(uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalTransactor) DepositERC20ToVault(opts *bind.TransactOpts, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.contract.Transact(opts, "depositERC20ToVault", assetType, vaultId, quantizedAmount)
}

// DepositERC20ToVault is a paid mutator transaction binding the contract method 0x1d133002.
//
// Solidity: function depositERC20ToVault(uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) DepositERC20ToVault(assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.Contract.DepositERC20ToVault(&_VaultDepositWithdrawal.TransactOpts, assetType, vaultId, quantizedAmount)
}

// DepositERC20ToVault is a paid mutator transaction binding the contract method 0x1d133002.
//
// Solidity: function depositERC20ToVault(uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalTransactorSession) DepositERC20ToVault(assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.Contract.DepositERC20ToVault(&_VaultDepositWithdrawal.TransactOpts, assetType, vaultId, quantizedAmount)
}

// DepositEthToVault is a paid mutator transaction binding the contract method 0xa3c71fde.
//
// Solidity: function depositEthToVault(uint256 assetType, uint256 vaultId) payable returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalTransactor) DepositEthToVault(opts *bind.TransactOpts, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.contract.Transact(opts, "depositEthToVault", assetType, vaultId)
}

// DepositEthToVault is a paid mutator transaction binding the contract method 0xa3c71fde.
//
// Solidity: function depositEthToVault(uint256 assetType, uint256 vaultId) payable returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) DepositEthToVault(assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.Contract.DepositEthToVault(&_VaultDepositWithdrawal.TransactOpts, assetType, vaultId)
}

// DepositEthToVault is a paid mutator transaction binding the contract method 0xa3c71fde.
//
// Solidity: function depositEthToVault(uint256 assetType, uint256 vaultId) payable returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalTransactorSession) DepositEthToVault(assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.Contract.DepositEthToVault(&_VaultDepositWithdrawal.TransactOpts, assetType, vaultId)
}

// WithdrawErc1155FromVault is a paid mutator transaction binding the contract method 0xa90f34a3.
//
// Solidity: function withdrawErc1155FromVault(uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalTransactor) WithdrawErc1155FromVault(opts *bind.TransactOpts, assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.contract.Transact(opts, "withdrawErc1155FromVault", assetType, tokenId, vaultId, quantizedAmount)
}

// WithdrawErc1155FromVault is a paid mutator transaction binding the contract method 0xa90f34a3.
//
// Solidity: function withdrawErc1155FromVault(uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) WithdrawErc1155FromVault(assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.Contract.WithdrawErc1155FromVault(&_VaultDepositWithdrawal.TransactOpts, assetType, tokenId, vaultId, quantizedAmount)
}

// WithdrawErc1155FromVault is a paid mutator transaction binding the contract method 0xa90f34a3.
//
// Solidity: function withdrawErc1155FromVault(uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalTransactorSession) WithdrawErc1155FromVault(assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.Contract.WithdrawErc1155FromVault(&_VaultDepositWithdrawal.TransactOpts, assetType, tokenId, vaultId, quantizedAmount)
}

// WithdrawFromVault is a paid mutator transaction binding the contract method 0xbf422b2e.
//
// Solidity: function withdrawFromVault(uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalTransactor) WithdrawFromVault(opts *bind.TransactOpts, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.contract.Transact(opts, "withdrawFromVault", assetType, vaultId, quantizedAmount)
}

// WithdrawFromVault is a paid mutator transaction binding the contract method 0xbf422b2e.
//
// Solidity: function withdrawFromVault(uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalSession) WithdrawFromVault(assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.Contract.WithdrawFromVault(&_VaultDepositWithdrawal.TransactOpts, assetType, vaultId, quantizedAmount)
}

// WithdrawFromVault is a paid mutator transaction binding the contract method 0xbf422b2e.
//
// Solidity: function withdrawFromVault(uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_VaultDepositWithdrawal *VaultDepositWithdrawalTransactorSession) WithdrawFromVault(assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _VaultDepositWithdrawal.Contract.WithdrawFromVault(&_VaultDepositWithdrawal.TransactOpts, assetType, vaultId, quantizedAmount)
}

// VaultDepositWithdrawalLogDepositToVaultIterator is returned from FilterLogDepositToVault and is used to iterate over the raw logs and unpacked data for LogDepositToVault events raised by the VaultDepositWithdrawal contract.
type VaultDepositWithdrawalLogDepositToVaultIterator struct {
	Event *VaultDepositWithdrawalLogDepositToVault // Event containing the contract specifics and raw log

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
func (it *VaultDepositWithdrawalLogDepositToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDepositWithdrawalLogDepositToVault)
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
		it.Event = new(VaultDepositWithdrawalLogDepositToVault)
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
func (it *VaultDepositWithdrawalLogDepositToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDepositWithdrawalLogDepositToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDepositWithdrawalLogDepositToVault represents a LogDepositToVault event raised by the VaultDepositWithdrawal contract.
type VaultDepositWithdrawalLogDepositToVault struct {
	EthKey             common.Address
	AssetType          *big.Int
	AssetId            *big.Int
	VaultId            *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDepositToVault is a free log retrieval operation binding the contract event 0xeab1149e7dda57c040af22580ce50b38b8b3af38433be30a24b3938166cd45e9.
//
// Solidity: event LogDepositToVault(address ethKey, uint256 assetType, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalFilterer) FilterLogDepositToVault(opts *bind.FilterOpts) (*VaultDepositWithdrawalLogDepositToVaultIterator, error) {

	logs, sub, err := _VaultDepositWithdrawal.contract.FilterLogs(opts, "LogDepositToVault")
	if err != nil {
		return nil, err
	}
	return &VaultDepositWithdrawalLogDepositToVaultIterator{contract: _VaultDepositWithdrawal.contract, event: "LogDepositToVault", logs: logs, sub: sub}, nil
}

// WatchLogDepositToVault is a free log subscription operation binding the contract event 0xeab1149e7dda57c040af22580ce50b38b8b3af38433be30a24b3938166cd45e9.
//
// Solidity: event LogDepositToVault(address ethKey, uint256 assetType, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalFilterer) WatchLogDepositToVault(opts *bind.WatchOpts, sink chan<- *VaultDepositWithdrawalLogDepositToVault) (event.Subscription, error) {

	logs, sub, err := _VaultDepositWithdrawal.contract.WatchLogs(opts, "LogDepositToVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDepositWithdrawalLogDepositToVault)
				if err := _VaultDepositWithdrawal.contract.UnpackLog(event, "LogDepositToVault", log); err != nil {
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

// ParseLogDepositToVault is a log parse operation binding the contract event 0xeab1149e7dda57c040af22580ce50b38b8b3af38433be30a24b3938166cd45e9.
//
// Solidity: event LogDepositToVault(address ethKey, uint256 assetType, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalFilterer) ParseLogDepositToVault(log types.Log) (*VaultDepositWithdrawalLogDepositToVault, error) {
	event := new(VaultDepositWithdrawalLogDepositToVault)
	if err := _VaultDepositWithdrawal.contract.UnpackLog(event, "LogDepositToVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDepositWithdrawalLogWithdrawalFromVaultIterator is returned from FilterLogWithdrawalFromVault and is used to iterate over the raw logs and unpacked data for LogWithdrawalFromVault events raised by the VaultDepositWithdrawal contract.
type VaultDepositWithdrawalLogWithdrawalFromVaultIterator struct {
	Event *VaultDepositWithdrawalLogWithdrawalFromVault // Event containing the contract specifics and raw log

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
func (it *VaultDepositWithdrawalLogWithdrawalFromVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDepositWithdrawalLogWithdrawalFromVault)
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
		it.Event = new(VaultDepositWithdrawalLogWithdrawalFromVault)
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
func (it *VaultDepositWithdrawalLogWithdrawalFromVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDepositWithdrawalLogWithdrawalFromVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDepositWithdrawalLogWithdrawalFromVault represents a LogWithdrawalFromVault event raised by the VaultDepositWithdrawal contract.
type VaultDepositWithdrawalLogWithdrawalFromVault struct {
	EthKey             common.Address
	AssetType          *big.Int
	AssetId            *big.Int
	VaultId            *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalFromVault is a free log retrieval operation binding the contract event 0x824058caef851aafcaa37b343756b98e0c9e3a86be89eaa6a326d66df04e5a43.
//
// Solidity: event LogWithdrawalFromVault(address ethKey, uint256 assetType, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalFilterer) FilterLogWithdrawalFromVault(opts *bind.FilterOpts) (*VaultDepositWithdrawalLogWithdrawalFromVaultIterator, error) {

	logs, sub, err := _VaultDepositWithdrawal.contract.FilterLogs(opts, "LogWithdrawalFromVault")
	if err != nil {
		return nil, err
	}
	return &VaultDepositWithdrawalLogWithdrawalFromVaultIterator{contract: _VaultDepositWithdrawal.contract, event: "LogWithdrawalFromVault", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalFromVault is a free log subscription operation binding the contract event 0x824058caef851aafcaa37b343756b98e0c9e3a86be89eaa6a326d66df04e5a43.
//
// Solidity: event LogWithdrawalFromVault(address ethKey, uint256 assetType, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalFilterer) WatchLogWithdrawalFromVault(opts *bind.WatchOpts, sink chan<- *VaultDepositWithdrawalLogWithdrawalFromVault) (event.Subscription, error) {

	logs, sub, err := _VaultDepositWithdrawal.contract.WatchLogs(opts, "LogWithdrawalFromVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDepositWithdrawalLogWithdrawalFromVault)
				if err := _VaultDepositWithdrawal.contract.UnpackLog(event, "LogWithdrawalFromVault", log); err != nil {
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

// ParseLogWithdrawalFromVault is a log parse operation binding the contract event 0x824058caef851aafcaa37b343756b98e0c9e3a86be89eaa6a326d66df04e5a43.
//
// Solidity: event LogWithdrawalFromVault(address ethKey, uint256 assetType, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_VaultDepositWithdrawal *VaultDepositWithdrawalFilterer) ParseLogWithdrawalFromVault(log types.Log) (*VaultDepositWithdrawalLogWithdrawalFromVault, error) {
	event := new(VaultDepositWithdrawalLogWithdrawalFromVault)
	if err := _VaultDepositWithdrawal.contract.UnpackLog(event, "LogWithdrawalFromVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
