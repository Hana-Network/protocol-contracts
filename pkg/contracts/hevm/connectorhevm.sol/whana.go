// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package connectorhevm

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

// WHANAMetaData contains all meta data concerning the WHANA contract.
var WHANAMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WHANAABI is the input ABI used to generate the binding from.
// Deprecated: Use WHANAMetaData.ABI instead.
var WHANAABI = WHANAMetaData.ABI

// WHANA is an auto generated Go binding around an Ethereum contract.
type WHANA struct {
	WHANACaller     // Read-only binding to the contract
	WHANATransactor // Write-only binding to the contract
	WHANAFilterer   // Log filterer for contract events
}

// WHANACaller is an auto generated read-only Go binding around an Ethereum contract.
type WHANACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WHANATransactor is an auto generated write-only Go binding around an Ethereum contract.
type WHANATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WHANAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WHANAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WHANASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WHANASession struct {
	Contract     *WHANA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WHANACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WHANACallerSession struct {
	Contract *WHANACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WHANATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WHANATransactorSession struct {
	Contract     *WHANATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WHANARaw is an auto generated low-level Go binding around an Ethereum contract.
type WHANARaw struct {
	Contract *WHANA // Generic contract binding to access the raw methods on
}

// WHANACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WHANACallerRaw struct {
	Contract *WHANACaller // Generic read-only contract binding to access the raw methods on
}

// WHANATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WHANATransactorRaw struct {
	Contract *WHANATransactor // Generic write-only contract binding to access the raw methods on
}

// NewWHANA creates a new instance of WHANA, bound to a specific deployed contract.
func NewWHANA(address common.Address, backend bind.ContractBackend) (*WHANA, error) {
	contract, err := bindWHANA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WHANA{WHANACaller: WHANACaller{contract: contract}, WHANATransactor: WHANATransactor{contract: contract}, WHANAFilterer: WHANAFilterer{contract: contract}}, nil
}

// NewWHANACaller creates a new read-only instance of WHANA, bound to a specific deployed contract.
func NewWHANACaller(address common.Address, caller bind.ContractCaller) (*WHANACaller, error) {
	contract, err := bindWHANA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WHANACaller{contract: contract}, nil
}

// NewWHANATransactor creates a new write-only instance of WHANA, bound to a specific deployed contract.
func NewWHANATransactor(address common.Address, transactor bind.ContractTransactor) (*WHANATransactor, error) {
	contract, err := bindWHANA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WHANATransactor{contract: contract}, nil
}

// NewWHANAFilterer creates a new log filterer instance of WHANA, bound to a specific deployed contract.
func NewWHANAFilterer(address common.Address, filterer bind.ContractFilterer) (*WHANAFilterer, error) {
	contract, err := bindWHANA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WHANAFilterer{contract: contract}, nil
}

// bindWHANA binds a generic wrapper to an already deployed contract.
func bindWHANA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WHANAMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WHANA *WHANARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WHANA.Contract.WHANACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WHANA *WHANARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WHANA.Contract.WHANATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WHANA *WHANARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WHANA.Contract.WHANATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WHANA *WHANACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WHANA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WHANA *WHANATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WHANA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WHANA *WHANATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WHANA.Contract.contract.Transact(opts, method, params...)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WHANA *WHANATransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WHANA.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WHANA *WHANASession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WHANA.Contract.TransferFrom(&_WHANA.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WHANA *WHANATransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WHANA.Contract.TransferFrom(&_WHANA.TransactOpts, src, dst, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WHANA *WHANATransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _WHANA.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WHANA *WHANASession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WHANA.Contract.Withdraw(&_WHANA.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WHANA *WHANATransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WHANA.Contract.Withdraw(&_WHANA.TransactOpts, wad)
}
