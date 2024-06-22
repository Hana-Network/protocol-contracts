// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interfaces

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

// HContractMetaData contains all meta data concerning the HContract contract.
var HContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"onCrossChainCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HContractABI is the input ABI used to generate the binding from.
// Deprecated: Use HContractMetaData.ABI instead.
var HContractABI = HContractMetaData.ABI

// HContract is an auto generated Go binding around an Ethereum contract.
type HContract struct {
	HContractCaller     // Read-only binding to the contract
	HContractTransactor // Write-only binding to the contract
	HContractFilterer   // Log filterer for contract events
}

// HContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type HContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HContractSession struct {
	Contract     *HContract        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HContractCallerSession struct {
	Contract *HContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HContractTransactorSession struct {
	Contract     *HContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type HContractRaw struct {
	Contract *HContract // Generic contract binding to access the raw methods on
}

// HContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HContractCallerRaw struct {
	Contract *HContractCaller // Generic read-only contract binding to access the raw methods on
}

// HContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HContractTransactorRaw struct {
	Contract *HContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHContract creates a new instance of HContract, bound to a specific deployed contract.
func NewHContract(address common.Address, backend bind.ContractBackend) (*HContract, error) {
	contract, err := bindHContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HContract{HContractCaller: HContractCaller{contract: contract}, HContractTransactor: HContractTransactor{contract: contract}, HContractFilterer: HContractFilterer{contract: contract}}, nil
}

// NewHContractCaller creates a new read-only instance of HContract, bound to a specific deployed contract.
func NewHContractCaller(address common.Address, caller bind.ContractCaller) (*HContractCaller, error) {
	contract, err := bindHContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HContractCaller{contract: contract}, nil
}

// NewHContractTransactor creates a new write-only instance of HContract, bound to a specific deployed contract.
func NewHContractTransactor(address common.Address, transactor bind.ContractTransactor) (*HContractTransactor, error) {
	contract, err := bindHContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HContractTransactor{contract: contract}, nil
}

// NewHContractFilterer creates a new log filterer instance of HContract, bound to a specific deployed contract.
func NewHContractFilterer(address common.Address, filterer bind.ContractFilterer) (*HContractFilterer, error) {
	contract, err := bindHContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HContractFilterer{contract: contract}, nil
}

// bindHContract binds a generic wrapper to an already deployed contract.
func bindHContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HContract *HContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HContract.Contract.HContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HContract *HContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HContract.Contract.HContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HContract *HContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HContract.Contract.HContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HContract *HContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HContract *HContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HContract *HContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HContract.Contract.contract.Transact(opts, method, params...)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xc8522691.
//
// Solidity: function onCrossChainCall(address hrc20, uint256 amount, bytes message) returns()
func (_HContract *HContractTransactor) OnCrossChainCall(opts *bind.TransactOpts, hrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _HContract.contract.Transact(opts, "onCrossChainCall", hrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xc8522691.
//
// Solidity: function onCrossChainCall(address hrc20, uint256 amount, bytes message) returns()
func (_HContract *HContractSession) OnCrossChainCall(hrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _HContract.Contract.OnCrossChainCall(&_HContract.TransactOpts, hrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xc8522691.
//
// Solidity: function onCrossChainCall(address hrc20, uint256 amount, bytes message) returns()
func (_HContract *HContractTransactorSession) OnCrossChainCall(hrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _HContract.Contract.OnCrossChainCall(&_HContract.TransactOpts, hrc20, amount, message)
}
