// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanatokenconsumertrident

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

// HanaTokenConsumerTridentErrorsMetaData contains all meta data concerning the HanaTokenConsumerTridentErrors contract.
var HanaTokenConsumerTridentErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ErrorSendingETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"}]",
}

// HanaTokenConsumerTridentErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaTokenConsumerTridentErrorsMetaData.ABI instead.
var HanaTokenConsumerTridentErrorsABI = HanaTokenConsumerTridentErrorsMetaData.ABI

// HanaTokenConsumerTridentErrors is an auto generated Go binding around an Ethereum contract.
type HanaTokenConsumerTridentErrors struct {
	HanaTokenConsumerTridentErrorsCaller     // Read-only binding to the contract
	HanaTokenConsumerTridentErrorsTransactor // Write-only binding to the contract
	HanaTokenConsumerTridentErrorsFilterer   // Log filterer for contract events
}

// HanaTokenConsumerTridentErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaTokenConsumerTridentErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerTridentErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaTokenConsumerTridentErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerTridentErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaTokenConsumerTridentErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerTridentErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaTokenConsumerTridentErrorsSession struct {
	Contract     *HanaTokenConsumerTridentErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// HanaTokenConsumerTridentErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaTokenConsumerTridentErrorsCallerSession struct {
	Contract *HanaTokenConsumerTridentErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// HanaTokenConsumerTridentErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaTokenConsumerTridentErrorsTransactorSession struct {
	Contract     *HanaTokenConsumerTridentErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// HanaTokenConsumerTridentErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaTokenConsumerTridentErrorsRaw struct {
	Contract *HanaTokenConsumerTridentErrors // Generic contract binding to access the raw methods on
}

// HanaTokenConsumerTridentErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaTokenConsumerTridentErrorsCallerRaw struct {
	Contract *HanaTokenConsumerTridentErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// HanaTokenConsumerTridentErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaTokenConsumerTridentErrorsTransactorRaw struct {
	Contract *HanaTokenConsumerTridentErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaTokenConsumerTridentErrors creates a new instance of HanaTokenConsumerTridentErrors, bound to a specific deployed contract.
func NewHanaTokenConsumerTridentErrors(address common.Address, backend bind.ContractBackend) (*HanaTokenConsumerTridentErrors, error) {
	contract, err := bindHanaTokenConsumerTridentErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerTridentErrors{HanaTokenConsumerTridentErrorsCaller: HanaTokenConsumerTridentErrorsCaller{contract: contract}, HanaTokenConsumerTridentErrorsTransactor: HanaTokenConsumerTridentErrorsTransactor{contract: contract}, HanaTokenConsumerTridentErrorsFilterer: HanaTokenConsumerTridentErrorsFilterer{contract: contract}}, nil
}

// NewHanaTokenConsumerTridentErrorsCaller creates a new read-only instance of HanaTokenConsumerTridentErrors, bound to a specific deployed contract.
func NewHanaTokenConsumerTridentErrorsCaller(address common.Address, caller bind.ContractCaller) (*HanaTokenConsumerTridentErrorsCaller, error) {
	contract, err := bindHanaTokenConsumerTridentErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerTridentErrorsCaller{contract: contract}, nil
}

// NewHanaTokenConsumerTridentErrorsTransactor creates a new write-only instance of HanaTokenConsumerTridentErrors, bound to a specific deployed contract.
func NewHanaTokenConsumerTridentErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaTokenConsumerTridentErrorsTransactor, error) {
	contract, err := bindHanaTokenConsumerTridentErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerTridentErrorsTransactor{contract: contract}, nil
}

// NewHanaTokenConsumerTridentErrorsFilterer creates a new log filterer instance of HanaTokenConsumerTridentErrors, bound to a specific deployed contract.
func NewHanaTokenConsumerTridentErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaTokenConsumerTridentErrorsFilterer, error) {
	contract, err := bindHanaTokenConsumerTridentErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerTridentErrorsFilterer{contract: contract}, nil
}

// bindHanaTokenConsumerTridentErrors binds a generic wrapper to an already deployed contract.
func bindHanaTokenConsumerTridentErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaTokenConsumerTridentErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaTokenConsumerTridentErrors *HanaTokenConsumerTridentErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaTokenConsumerTridentErrors.Contract.HanaTokenConsumerTridentErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaTokenConsumerTridentErrors *HanaTokenConsumerTridentErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumerTridentErrors.Contract.HanaTokenConsumerTridentErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaTokenConsumerTridentErrors *HanaTokenConsumerTridentErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaTokenConsumerTridentErrors.Contract.HanaTokenConsumerTridentErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaTokenConsumerTridentErrors *HanaTokenConsumerTridentErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaTokenConsumerTridentErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaTokenConsumerTridentErrors *HanaTokenConsumerTridentErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumerTridentErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaTokenConsumerTridentErrors *HanaTokenConsumerTridentErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaTokenConsumerTridentErrors.Contract.contract.Transact(opts, method, params...)
}
