// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanainteractorerrors

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

// HanaInteractorErrorsMetaData contains all meta data concerning the HanaInteractorErrors contract.
var HanaInteractorErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidHanaMessageCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidHanaRevertCall\",\"type\":\"error\"}]",
}

// HanaInteractorErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaInteractorErrorsMetaData.ABI instead.
var HanaInteractorErrorsABI = HanaInteractorErrorsMetaData.ABI

// HanaInteractorErrors is an auto generated Go binding around an Ethereum contract.
type HanaInteractorErrors struct {
	HanaInteractorErrorsCaller     // Read-only binding to the contract
	HanaInteractorErrorsTransactor // Write-only binding to the contract
	HanaInteractorErrorsFilterer   // Log filterer for contract events
}

// HanaInteractorErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaInteractorErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaInteractorErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaInteractorErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaInteractorErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaInteractorErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaInteractorErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaInteractorErrorsSession struct {
	Contract     *HanaInteractorErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HanaInteractorErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaInteractorErrorsCallerSession struct {
	Contract *HanaInteractorErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// HanaInteractorErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaInteractorErrorsTransactorSession struct {
	Contract     *HanaInteractorErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// HanaInteractorErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaInteractorErrorsRaw struct {
	Contract *HanaInteractorErrors // Generic contract binding to access the raw methods on
}

// HanaInteractorErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaInteractorErrorsCallerRaw struct {
	Contract *HanaInteractorErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// HanaInteractorErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaInteractorErrorsTransactorRaw struct {
	Contract *HanaInteractorErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaInteractorErrors creates a new instance of HanaInteractorErrors, bound to a specific deployed contract.
func NewHanaInteractorErrors(address common.Address, backend bind.ContractBackend) (*HanaInteractorErrors, error) {
	contract, err := bindHanaInteractorErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorErrors{HanaInteractorErrorsCaller: HanaInteractorErrorsCaller{contract: contract}, HanaInteractorErrorsTransactor: HanaInteractorErrorsTransactor{contract: contract}, HanaInteractorErrorsFilterer: HanaInteractorErrorsFilterer{contract: contract}}, nil
}

// NewHanaInteractorErrorsCaller creates a new read-only instance of HanaInteractorErrors, bound to a specific deployed contract.
func NewHanaInteractorErrorsCaller(address common.Address, caller bind.ContractCaller) (*HanaInteractorErrorsCaller, error) {
	contract, err := bindHanaInteractorErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorErrorsCaller{contract: contract}, nil
}

// NewHanaInteractorErrorsTransactor creates a new write-only instance of HanaInteractorErrors, bound to a specific deployed contract.
func NewHanaInteractorErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaInteractorErrorsTransactor, error) {
	contract, err := bindHanaInteractorErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorErrorsTransactor{contract: contract}, nil
}

// NewHanaInteractorErrorsFilterer creates a new log filterer instance of HanaInteractorErrors, bound to a specific deployed contract.
func NewHanaInteractorErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaInteractorErrorsFilterer, error) {
	contract, err := bindHanaInteractorErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorErrorsFilterer{contract: contract}, nil
}

// bindHanaInteractorErrors binds a generic wrapper to an already deployed contract.
func bindHanaInteractorErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaInteractorErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaInteractorErrors *HanaInteractorErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaInteractorErrors.Contract.HanaInteractorErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaInteractorErrors *HanaInteractorErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaInteractorErrors.Contract.HanaInteractorErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaInteractorErrors *HanaInteractorErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaInteractorErrors.Contract.HanaInteractorErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaInteractorErrors *HanaInteractorErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaInteractorErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaInteractorErrors *HanaInteractorErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaInteractorErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaInteractorErrors *HanaInteractorErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaInteractorErrors.Contract.contract.Transact(opts, method, params...)
}
