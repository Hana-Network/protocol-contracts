// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanatokenconsumeruniv2

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

// HanaTokenConsumerUniV2ErrorsMetaData contains all meta data concerning the HanaTokenConsumerUniV2Errors contract.
var HanaTokenConsumerUniV2ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"}]",
}

// HanaTokenConsumerUniV2ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaTokenConsumerUniV2ErrorsMetaData.ABI instead.
var HanaTokenConsumerUniV2ErrorsABI = HanaTokenConsumerUniV2ErrorsMetaData.ABI

// HanaTokenConsumerUniV2Errors is an auto generated Go binding around an Ethereum contract.
type HanaTokenConsumerUniV2Errors struct {
	HanaTokenConsumerUniV2ErrorsCaller     // Read-only binding to the contract
	HanaTokenConsumerUniV2ErrorsTransactor // Write-only binding to the contract
	HanaTokenConsumerUniV2ErrorsFilterer   // Log filterer for contract events
}

// HanaTokenConsumerUniV2ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV2ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerUniV2ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV2ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerUniV2ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaTokenConsumerUniV2ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerUniV2ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaTokenConsumerUniV2ErrorsSession struct {
	Contract     *HanaTokenConsumerUniV2Errors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// HanaTokenConsumerUniV2ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaTokenConsumerUniV2ErrorsCallerSession struct {
	Contract *HanaTokenConsumerUniV2ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// HanaTokenConsumerUniV2ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaTokenConsumerUniV2ErrorsTransactorSession struct {
	Contract     *HanaTokenConsumerUniV2ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// HanaTokenConsumerUniV2ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaTokenConsumerUniV2ErrorsRaw struct {
	Contract *HanaTokenConsumerUniV2Errors // Generic contract binding to access the raw methods on
}

// HanaTokenConsumerUniV2ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV2ErrorsCallerRaw struct {
	Contract *HanaTokenConsumerUniV2ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// HanaTokenConsumerUniV2ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV2ErrorsTransactorRaw struct {
	Contract *HanaTokenConsumerUniV2ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaTokenConsumerUniV2Errors creates a new instance of HanaTokenConsumerUniV2Errors, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV2Errors(address common.Address, backend bind.ContractBackend) (*HanaTokenConsumerUniV2Errors, error) {
	contract, err := bindHanaTokenConsumerUniV2Errors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV2Errors{HanaTokenConsumerUniV2ErrorsCaller: HanaTokenConsumerUniV2ErrorsCaller{contract: contract}, HanaTokenConsumerUniV2ErrorsTransactor: HanaTokenConsumerUniV2ErrorsTransactor{contract: contract}, HanaTokenConsumerUniV2ErrorsFilterer: HanaTokenConsumerUniV2ErrorsFilterer{contract: contract}}, nil
}

// NewHanaTokenConsumerUniV2ErrorsCaller creates a new read-only instance of HanaTokenConsumerUniV2Errors, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV2ErrorsCaller(address common.Address, caller bind.ContractCaller) (*HanaTokenConsumerUniV2ErrorsCaller, error) {
	contract, err := bindHanaTokenConsumerUniV2Errors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV2ErrorsCaller{contract: contract}, nil
}

// NewHanaTokenConsumerUniV2ErrorsTransactor creates a new write-only instance of HanaTokenConsumerUniV2Errors, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV2ErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaTokenConsumerUniV2ErrorsTransactor, error) {
	contract, err := bindHanaTokenConsumerUniV2Errors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV2ErrorsTransactor{contract: contract}, nil
}

// NewHanaTokenConsumerUniV2ErrorsFilterer creates a new log filterer instance of HanaTokenConsumerUniV2Errors, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV2ErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaTokenConsumerUniV2ErrorsFilterer, error) {
	contract, err := bindHanaTokenConsumerUniV2Errors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV2ErrorsFilterer{contract: contract}, nil
}

// bindHanaTokenConsumerUniV2Errors binds a generic wrapper to an already deployed contract.
func bindHanaTokenConsumerUniV2Errors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaTokenConsumerUniV2ErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaTokenConsumerUniV2Errors *HanaTokenConsumerUniV2ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaTokenConsumerUniV2Errors.Contract.HanaTokenConsumerUniV2ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaTokenConsumerUniV2Errors *HanaTokenConsumerUniV2ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2Errors.Contract.HanaTokenConsumerUniV2ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaTokenConsumerUniV2Errors *HanaTokenConsumerUniV2ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2Errors.Contract.HanaTokenConsumerUniV2ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaTokenConsumerUniV2Errors *HanaTokenConsumerUniV2ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaTokenConsumerUniV2Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaTokenConsumerUniV2Errors *HanaTokenConsumerUniV2ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaTokenConsumerUniV2Errors *HanaTokenConsumerUniV2ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2Errors.Contract.contract.Transact(opts, method, params...)
}
