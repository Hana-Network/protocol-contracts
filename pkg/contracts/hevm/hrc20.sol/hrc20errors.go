// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hrc20

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

// HRC20ErrorsMetaData contains all meta data concerning the HRC20Errors contract.
var HRC20ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CallerIsNotFungibleModule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasFeeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LowAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LowBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroGasCoin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroGasPrice\",\"type\":\"error\"}]",
}

// HRC20ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use HRC20ErrorsMetaData.ABI instead.
var HRC20ErrorsABI = HRC20ErrorsMetaData.ABI

// HRC20Errors is an auto generated Go binding around an Ethereum contract.
type HRC20Errors struct {
	HRC20ErrorsCaller     // Read-only binding to the contract
	HRC20ErrorsTransactor // Write-only binding to the contract
	HRC20ErrorsFilterer   // Log filterer for contract events
}

// HRC20ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type HRC20ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HRC20ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HRC20ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HRC20ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HRC20ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HRC20ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HRC20ErrorsSession struct {
	Contract     *HRC20Errors      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HRC20ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HRC20ErrorsCallerSession struct {
	Contract *HRC20ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HRC20ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HRC20ErrorsTransactorSession struct {
	Contract     *HRC20ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HRC20ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type HRC20ErrorsRaw struct {
	Contract *HRC20Errors // Generic contract binding to access the raw methods on
}

// HRC20ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HRC20ErrorsCallerRaw struct {
	Contract *HRC20ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// HRC20ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HRC20ErrorsTransactorRaw struct {
	Contract *HRC20ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHRC20Errors creates a new instance of HRC20Errors, bound to a specific deployed contract.
func NewHRC20Errors(address common.Address, backend bind.ContractBackend) (*HRC20Errors, error) {
	contract, err := bindHRC20Errors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HRC20Errors{HRC20ErrorsCaller: HRC20ErrorsCaller{contract: contract}, HRC20ErrorsTransactor: HRC20ErrorsTransactor{contract: contract}, HRC20ErrorsFilterer: HRC20ErrorsFilterer{contract: contract}}, nil
}

// NewHRC20ErrorsCaller creates a new read-only instance of HRC20Errors, bound to a specific deployed contract.
func NewHRC20ErrorsCaller(address common.Address, caller bind.ContractCaller) (*HRC20ErrorsCaller, error) {
	contract, err := bindHRC20Errors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HRC20ErrorsCaller{contract: contract}, nil
}

// NewHRC20ErrorsTransactor creates a new write-only instance of HRC20Errors, bound to a specific deployed contract.
func NewHRC20ErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*HRC20ErrorsTransactor, error) {
	contract, err := bindHRC20Errors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HRC20ErrorsTransactor{contract: contract}, nil
}

// NewHRC20ErrorsFilterer creates a new log filterer instance of HRC20Errors, bound to a specific deployed contract.
func NewHRC20ErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*HRC20ErrorsFilterer, error) {
	contract, err := bindHRC20Errors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HRC20ErrorsFilterer{contract: contract}, nil
}

// bindHRC20Errors binds a generic wrapper to an already deployed contract.
func bindHRC20Errors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HRC20ErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HRC20Errors *HRC20ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HRC20Errors.Contract.HRC20ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HRC20Errors *HRC20ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HRC20Errors.Contract.HRC20ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HRC20Errors *HRC20ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HRC20Errors.Contract.HRC20ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HRC20Errors *HRC20ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HRC20Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HRC20Errors *HRC20ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HRC20Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HRC20Errors *HRC20ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HRC20Errors.Contract.contract.Transact(opts, method, params...)
}
