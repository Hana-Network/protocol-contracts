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

// HanaInterfacesSendInput is an auto generated low-level Go binding around an user-defined struct.
type HanaInterfacesSendInput struct {
	DestinationChainId  *big.Int
	DestinationAddress  []byte
	DestinationGasLimit *big.Int
	Message             []byte
	HanaValueAndGas     *big.Int
	HanaParams          []byte
}

// HanaConnectorMetaData contains all meta data concerning the HanaConnector contract.
var HanaConnectorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"hanaValueAndGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"hanaParams\",\"type\":\"bytes\"}],\"internalType\":\"structHanaInterfaces.SendInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HanaConnectorABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaConnectorMetaData.ABI instead.
var HanaConnectorABI = HanaConnectorMetaData.ABI

// HanaConnector is an auto generated Go binding around an Ethereum contract.
type HanaConnector struct {
	HanaConnectorCaller     // Read-only binding to the contract
	HanaConnectorTransactor // Write-only binding to the contract
	HanaConnectorFilterer   // Log filterer for contract events
}

// HanaConnectorCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaConnectorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaConnectorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaConnectorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaConnectorSession struct {
	Contract     *HanaConnector    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HanaConnectorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaConnectorCallerSession struct {
	Contract *HanaConnectorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// HanaConnectorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaConnectorTransactorSession struct {
	Contract     *HanaConnectorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// HanaConnectorRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaConnectorRaw struct {
	Contract *HanaConnector // Generic contract binding to access the raw methods on
}

// HanaConnectorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaConnectorCallerRaw struct {
	Contract *HanaConnectorCaller // Generic read-only contract binding to access the raw methods on
}

// HanaConnectorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaConnectorTransactorRaw struct {
	Contract *HanaConnectorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaConnector creates a new instance of HanaConnector, bound to a specific deployed contract.
func NewHanaConnector(address common.Address, backend bind.ContractBackend) (*HanaConnector, error) {
	contract, err := bindHanaConnector(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaConnector{HanaConnectorCaller: HanaConnectorCaller{contract: contract}, HanaConnectorTransactor: HanaConnectorTransactor{contract: contract}, HanaConnectorFilterer: HanaConnectorFilterer{contract: contract}}, nil
}

// NewHanaConnectorCaller creates a new read-only instance of HanaConnector, bound to a specific deployed contract.
func NewHanaConnectorCaller(address common.Address, caller bind.ContractCaller) (*HanaConnectorCaller, error) {
	contract, err := bindHanaConnector(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorCaller{contract: contract}, nil
}

// NewHanaConnectorTransactor creates a new write-only instance of HanaConnector, bound to a specific deployed contract.
func NewHanaConnectorTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaConnectorTransactor, error) {
	contract, err := bindHanaConnector(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorTransactor{contract: contract}, nil
}

// NewHanaConnectorFilterer creates a new log filterer instance of HanaConnector, bound to a specific deployed contract.
func NewHanaConnectorFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaConnectorFilterer, error) {
	contract, err := bindHanaConnector(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorFilterer{contract: contract}, nil
}

// bindHanaConnector binds a generic wrapper to an already deployed contract.
func bindHanaConnector(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaConnectorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaConnector *HanaConnectorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaConnector.Contract.HanaConnectorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaConnector *HanaConnectorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnector.Contract.HanaConnectorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaConnector *HanaConnectorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaConnector.Contract.HanaConnectorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaConnector *HanaConnectorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaConnector.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaConnector *HanaConnectorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnector.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaConnector *HanaConnectorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaConnector.Contract.contract.Transact(opts, method, params...)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnector *HanaConnectorTransactor) Send(opts *bind.TransactOpts, input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnector.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnector *HanaConnectorSession) Send(input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnector.Contract.Send(&_HanaConnector.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnector *HanaConnectorTransactorSession) Send(input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnector.Contract.Send(&_HanaConnector.TransactOpts, input)
}
