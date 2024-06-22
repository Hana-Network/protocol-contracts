// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanainterfaces

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

// HanaCommonErrorsMetaData contains all meta data concerning the HanaCommonErrors contract.
var HanaCommonErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"}]",
}

// HanaCommonErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaCommonErrorsMetaData.ABI instead.
var HanaCommonErrorsABI = HanaCommonErrorsMetaData.ABI

// HanaCommonErrors is an auto generated Go binding around an Ethereum contract.
type HanaCommonErrors struct {
	HanaCommonErrorsCaller     // Read-only binding to the contract
	HanaCommonErrorsTransactor // Write-only binding to the contract
	HanaCommonErrorsFilterer   // Log filterer for contract events
}

// HanaCommonErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaCommonErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaCommonErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaCommonErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaCommonErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaCommonErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaCommonErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaCommonErrorsSession struct {
	Contract     *HanaCommonErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HanaCommonErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaCommonErrorsCallerSession struct {
	Contract *HanaCommonErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// HanaCommonErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaCommonErrorsTransactorSession struct {
	Contract     *HanaCommonErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// HanaCommonErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaCommonErrorsRaw struct {
	Contract *HanaCommonErrors // Generic contract binding to access the raw methods on
}

// HanaCommonErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaCommonErrorsCallerRaw struct {
	Contract *HanaCommonErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// HanaCommonErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaCommonErrorsTransactorRaw struct {
	Contract *HanaCommonErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaCommonErrors creates a new instance of HanaCommonErrors, bound to a specific deployed contract.
func NewHanaCommonErrors(address common.Address, backend bind.ContractBackend) (*HanaCommonErrors, error) {
	contract, err := bindHanaCommonErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaCommonErrors{HanaCommonErrorsCaller: HanaCommonErrorsCaller{contract: contract}, HanaCommonErrorsTransactor: HanaCommonErrorsTransactor{contract: contract}, HanaCommonErrorsFilterer: HanaCommonErrorsFilterer{contract: contract}}, nil
}

// NewHanaCommonErrorsCaller creates a new read-only instance of HanaCommonErrors, bound to a specific deployed contract.
func NewHanaCommonErrorsCaller(address common.Address, caller bind.ContractCaller) (*HanaCommonErrorsCaller, error) {
	contract, err := bindHanaCommonErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaCommonErrorsCaller{contract: contract}, nil
}

// NewHanaCommonErrorsTransactor creates a new write-only instance of HanaCommonErrors, bound to a specific deployed contract.
func NewHanaCommonErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaCommonErrorsTransactor, error) {
	contract, err := bindHanaCommonErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaCommonErrorsTransactor{contract: contract}, nil
}

// NewHanaCommonErrorsFilterer creates a new log filterer instance of HanaCommonErrors, bound to a specific deployed contract.
func NewHanaCommonErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaCommonErrorsFilterer, error) {
	contract, err := bindHanaCommonErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaCommonErrorsFilterer{contract: contract}, nil
}

// bindHanaCommonErrors binds a generic wrapper to an already deployed contract.
func bindHanaCommonErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaCommonErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaCommonErrors *HanaCommonErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaCommonErrors.Contract.HanaCommonErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaCommonErrors *HanaCommonErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaCommonErrors.Contract.HanaCommonErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaCommonErrors *HanaCommonErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaCommonErrors.Contract.HanaCommonErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaCommonErrors *HanaCommonErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaCommonErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaCommonErrors *HanaCommonErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaCommonErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaCommonErrors *HanaCommonErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaCommonErrors.Contract.contract.Transact(opts, method, params...)
}
