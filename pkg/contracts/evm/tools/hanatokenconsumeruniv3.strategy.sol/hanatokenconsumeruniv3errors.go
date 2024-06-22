// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanatokenconsumeruniv3

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

// HanaTokenConsumerUniV3ErrorsMetaData contains all meta data concerning the HanaTokenConsumerUniV3Errors contract.
var HanaTokenConsumerUniV3ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ErrorSendingETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"}]",
}

// HanaTokenConsumerUniV3ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaTokenConsumerUniV3ErrorsMetaData.ABI instead.
var HanaTokenConsumerUniV3ErrorsABI = HanaTokenConsumerUniV3ErrorsMetaData.ABI

// HanaTokenConsumerUniV3Errors is an auto generated Go binding around an Ethereum contract.
type HanaTokenConsumerUniV3Errors struct {
	HanaTokenConsumerUniV3ErrorsCaller     // Read-only binding to the contract
	HanaTokenConsumerUniV3ErrorsTransactor // Write-only binding to the contract
	HanaTokenConsumerUniV3ErrorsFilterer   // Log filterer for contract events
}

// HanaTokenConsumerUniV3ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV3ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerUniV3ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV3ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerUniV3ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaTokenConsumerUniV3ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerUniV3ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaTokenConsumerUniV3ErrorsSession struct {
	Contract     *HanaTokenConsumerUniV3Errors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// HanaTokenConsumerUniV3ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaTokenConsumerUniV3ErrorsCallerSession struct {
	Contract *HanaTokenConsumerUniV3ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// HanaTokenConsumerUniV3ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaTokenConsumerUniV3ErrorsTransactorSession struct {
	Contract     *HanaTokenConsumerUniV3ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// HanaTokenConsumerUniV3ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaTokenConsumerUniV3ErrorsRaw struct {
	Contract *HanaTokenConsumerUniV3Errors // Generic contract binding to access the raw methods on
}

// HanaTokenConsumerUniV3ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV3ErrorsCallerRaw struct {
	Contract *HanaTokenConsumerUniV3ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// HanaTokenConsumerUniV3ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV3ErrorsTransactorRaw struct {
	Contract *HanaTokenConsumerUniV3ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaTokenConsumerUniV3Errors creates a new instance of HanaTokenConsumerUniV3Errors, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV3Errors(address common.Address, backend bind.ContractBackend) (*HanaTokenConsumerUniV3Errors, error) {
	contract, err := bindHanaTokenConsumerUniV3Errors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV3Errors{HanaTokenConsumerUniV3ErrorsCaller: HanaTokenConsumerUniV3ErrorsCaller{contract: contract}, HanaTokenConsumerUniV3ErrorsTransactor: HanaTokenConsumerUniV3ErrorsTransactor{contract: contract}, HanaTokenConsumerUniV3ErrorsFilterer: HanaTokenConsumerUniV3ErrorsFilterer{contract: contract}}, nil
}

// NewHanaTokenConsumerUniV3ErrorsCaller creates a new read-only instance of HanaTokenConsumerUniV3Errors, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV3ErrorsCaller(address common.Address, caller bind.ContractCaller) (*HanaTokenConsumerUniV3ErrorsCaller, error) {
	contract, err := bindHanaTokenConsumerUniV3Errors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV3ErrorsCaller{contract: contract}, nil
}

// NewHanaTokenConsumerUniV3ErrorsTransactor creates a new write-only instance of HanaTokenConsumerUniV3Errors, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV3ErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaTokenConsumerUniV3ErrorsTransactor, error) {
	contract, err := bindHanaTokenConsumerUniV3Errors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV3ErrorsTransactor{contract: contract}, nil
}

// NewHanaTokenConsumerUniV3ErrorsFilterer creates a new log filterer instance of HanaTokenConsumerUniV3Errors, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV3ErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaTokenConsumerUniV3ErrorsFilterer, error) {
	contract, err := bindHanaTokenConsumerUniV3Errors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV3ErrorsFilterer{contract: contract}, nil
}

// bindHanaTokenConsumerUniV3Errors binds a generic wrapper to an already deployed contract.
func bindHanaTokenConsumerUniV3Errors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaTokenConsumerUniV3ErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaTokenConsumerUniV3Errors *HanaTokenConsumerUniV3ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaTokenConsumerUniV3Errors.Contract.HanaTokenConsumerUniV3ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaTokenConsumerUniV3Errors *HanaTokenConsumerUniV3ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3Errors.Contract.HanaTokenConsumerUniV3ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaTokenConsumerUniV3Errors *HanaTokenConsumerUniV3ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3Errors.Contract.HanaTokenConsumerUniV3ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaTokenConsumerUniV3Errors *HanaTokenConsumerUniV3ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaTokenConsumerUniV3Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaTokenConsumerUniV3Errors *HanaTokenConsumerUniV3ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaTokenConsumerUniV3Errors *HanaTokenConsumerUniV3ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3Errors.Contract.contract.Transact(opts, method, params...)
}
