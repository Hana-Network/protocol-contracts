// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanaconnectorhevm

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

// HanaInterfacesMetaData contains all meta data concerning the HanaInterfaces contract.
var HanaInterfacesMetaData = &bind.MetaData{
	ABI: "[]",
}

// HanaInterfacesABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaInterfacesMetaData.ABI instead.
var HanaInterfacesABI = HanaInterfacesMetaData.ABI

// HanaInterfaces is an auto generated Go binding around an Ethereum contract.
type HanaInterfaces struct {
	HanaInterfacesCaller     // Read-only binding to the contract
	HanaInterfacesTransactor // Write-only binding to the contract
	HanaInterfacesFilterer   // Log filterer for contract events
}

// HanaInterfacesCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaInterfacesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaInterfacesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaInterfacesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaInterfacesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaInterfacesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaInterfacesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaInterfacesSession struct {
	Contract     *HanaInterfaces   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HanaInterfacesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaInterfacesCallerSession struct {
	Contract *HanaInterfacesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HanaInterfacesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaInterfacesTransactorSession struct {
	Contract     *HanaInterfacesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HanaInterfacesRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaInterfacesRaw struct {
	Contract *HanaInterfaces // Generic contract binding to access the raw methods on
}

// HanaInterfacesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaInterfacesCallerRaw struct {
	Contract *HanaInterfacesCaller // Generic read-only contract binding to access the raw methods on
}

// HanaInterfacesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaInterfacesTransactorRaw struct {
	Contract *HanaInterfacesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaInterfaces creates a new instance of HanaInterfaces, bound to a specific deployed contract.
func NewHanaInterfaces(address common.Address, backend bind.ContractBackend) (*HanaInterfaces, error) {
	contract, err := bindHanaInterfaces(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaInterfaces{HanaInterfacesCaller: HanaInterfacesCaller{contract: contract}, HanaInterfacesTransactor: HanaInterfacesTransactor{contract: contract}, HanaInterfacesFilterer: HanaInterfacesFilterer{contract: contract}}, nil
}

// NewHanaInterfacesCaller creates a new read-only instance of HanaInterfaces, bound to a specific deployed contract.
func NewHanaInterfacesCaller(address common.Address, caller bind.ContractCaller) (*HanaInterfacesCaller, error) {
	contract, err := bindHanaInterfaces(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaInterfacesCaller{contract: contract}, nil
}

// NewHanaInterfacesTransactor creates a new write-only instance of HanaInterfaces, bound to a specific deployed contract.
func NewHanaInterfacesTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaInterfacesTransactor, error) {
	contract, err := bindHanaInterfaces(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaInterfacesTransactor{contract: contract}, nil
}

// NewHanaInterfacesFilterer creates a new log filterer instance of HanaInterfaces, bound to a specific deployed contract.
func NewHanaInterfacesFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaInterfacesFilterer, error) {
	contract, err := bindHanaInterfaces(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaInterfacesFilterer{contract: contract}, nil
}

// bindHanaInterfaces binds a generic wrapper to an already deployed contract.
func bindHanaInterfaces(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaInterfacesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaInterfaces *HanaInterfacesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaInterfaces.Contract.HanaInterfacesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaInterfaces *HanaInterfacesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaInterfaces.Contract.HanaInterfacesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaInterfaces *HanaInterfacesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaInterfaces.Contract.HanaInterfacesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaInterfaces *HanaInterfacesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaInterfaces.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaInterfaces *HanaInterfacesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaInterfaces.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaInterfaces *HanaInterfacesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaInterfaces.Contract.contract.Transact(opts, method, params...)
}
