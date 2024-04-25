// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanaerrors

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

// HanaErrorsMetaData contains all meta data concerning the HanaErrors contract.
var HanaErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotConnector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTss\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssOrUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HanaTransferError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"}]",
}

// HanaErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaErrorsMetaData.ABI instead.
var HanaErrorsABI = HanaErrorsMetaData.ABI

// HanaErrors is an auto generated Go binding around an Ethereum contract.
type HanaErrors struct {
	HanaErrorsCaller     // Read-only binding to the contract
	HanaErrorsTransactor // Write-only binding to the contract
	HanaErrorsFilterer   // Log filterer for contract events
}

// HanaErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaErrorsSession struct {
	Contract     *HanaErrors       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HanaErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaErrorsCallerSession struct {
	Contract *HanaErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HanaErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaErrorsTransactorSession struct {
	Contract     *HanaErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HanaErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaErrorsRaw struct {
	Contract *HanaErrors // Generic contract binding to access the raw methods on
}

// HanaErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaErrorsCallerRaw struct {
	Contract *HanaErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// HanaErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaErrorsTransactorRaw struct {
	Contract *HanaErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaErrors creates a new instance of HanaErrors, bound to a specific deployed contract.
func NewHanaErrors(address common.Address, backend bind.ContractBackend) (*HanaErrors, error) {
	contract, err := bindHanaErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaErrors{HanaErrorsCaller: HanaErrorsCaller{contract: contract}, HanaErrorsTransactor: HanaErrorsTransactor{contract: contract}, HanaErrorsFilterer: HanaErrorsFilterer{contract: contract}}, nil
}

// NewHanaErrorsCaller creates a new read-only instance of HanaErrors, bound to a specific deployed contract.
func NewHanaErrorsCaller(address common.Address, caller bind.ContractCaller) (*HanaErrorsCaller, error) {
	contract, err := bindHanaErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaErrorsCaller{contract: contract}, nil
}

// NewHanaErrorsTransactor creates a new write-only instance of HanaErrors, bound to a specific deployed contract.
func NewHanaErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaErrorsTransactor, error) {
	contract, err := bindHanaErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaErrorsTransactor{contract: contract}, nil
}

// NewHanaErrorsFilterer creates a new log filterer instance of HanaErrors, bound to a specific deployed contract.
func NewHanaErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaErrorsFilterer, error) {
	contract, err := bindHanaErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaErrorsFilterer{contract: contract}, nil
}

// bindHanaErrors binds a generic wrapper to an already deployed contract.
func bindHanaErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaErrors *HanaErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaErrors.Contract.HanaErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaErrors *HanaErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaErrors.Contract.HanaErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaErrors *HanaErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaErrors.Contract.HanaErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaErrors *HanaErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaErrors *HanaErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaErrors *HanaErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaErrors.Contract.contract.Transact(opts, method, params...)
}
