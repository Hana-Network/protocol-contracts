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

// HanaInterfacesHanaMessage is an auto generated low-level Go binding around an user-defined struct.
type HanaInterfacesHanaMessage struct {
	HanaTxSenderAddress []byte
	SourceChainId       *big.Int
	DestinationAddress  common.Address
	HanaValue           *big.Int
	Message             []byte
}

// HanaInterfacesHanaRevert is an auto generated low-level Go binding around an user-defined struct.
type HanaInterfacesHanaRevert struct {
	HanaTxSenderAddress common.Address
	SourceChainId       *big.Int
	DestinationAddress  []byte
	DestinationChainId  *big.Int
	RemainingHanaValue  *big.Int
	Message             []byte
}

// HanaReceiverMetaData contains all meta data concerning the HanaReceiver contract.
var HanaReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"hanaTxSenderAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hanaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structHanaInterfaces.HanaMessage\",\"name\":\"hanaMessage\",\"type\":\"tuple\"}],\"name\":\"onHanaMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingHanaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structHanaInterfaces.HanaRevert\",\"name\":\"hanaRevert\",\"type\":\"tuple\"}],\"name\":\"onHanaRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HanaReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaReceiverMetaData.ABI instead.
var HanaReceiverABI = HanaReceiverMetaData.ABI

// HanaReceiver is an auto generated Go binding around an Ethereum contract.
type HanaReceiver struct {
	HanaReceiverCaller     // Read-only binding to the contract
	HanaReceiverTransactor // Write-only binding to the contract
	HanaReceiverFilterer   // Log filterer for contract events
}

// HanaReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaReceiverSession struct {
	Contract     *HanaReceiver     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HanaReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaReceiverCallerSession struct {
	Contract *HanaReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HanaReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaReceiverTransactorSession struct {
	Contract     *HanaReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HanaReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaReceiverRaw struct {
	Contract *HanaReceiver // Generic contract binding to access the raw methods on
}

// HanaReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaReceiverCallerRaw struct {
	Contract *HanaReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// HanaReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaReceiverTransactorRaw struct {
	Contract *HanaReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaReceiver creates a new instance of HanaReceiver, bound to a specific deployed contract.
func NewHanaReceiver(address common.Address, backend bind.ContractBackend) (*HanaReceiver, error) {
	contract, err := bindHanaReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaReceiver{HanaReceiverCaller: HanaReceiverCaller{contract: contract}, HanaReceiverTransactor: HanaReceiverTransactor{contract: contract}, HanaReceiverFilterer: HanaReceiverFilterer{contract: contract}}, nil
}

// NewHanaReceiverCaller creates a new read-only instance of HanaReceiver, bound to a specific deployed contract.
func NewHanaReceiverCaller(address common.Address, caller bind.ContractCaller) (*HanaReceiverCaller, error) {
	contract, err := bindHanaReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaReceiverCaller{contract: contract}, nil
}

// NewHanaReceiverTransactor creates a new write-only instance of HanaReceiver, bound to a specific deployed contract.
func NewHanaReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaReceiverTransactor, error) {
	contract, err := bindHanaReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaReceiverTransactor{contract: contract}, nil
}

// NewHanaReceiverFilterer creates a new log filterer instance of HanaReceiver, bound to a specific deployed contract.
func NewHanaReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaReceiverFilterer, error) {
	contract, err := bindHanaReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaReceiverFilterer{contract: contract}, nil
}

// bindHanaReceiver binds a generic wrapper to an already deployed contract.
func bindHanaReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaReceiver *HanaReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaReceiver.Contract.HanaReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaReceiver *HanaReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaReceiver.Contract.HanaReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaReceiver *HanaReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaReceiver.Contract.HanaReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaReceiver *HanaReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaReceiver *HanaReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaReceiver *HanaReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnHanaMessage is a paid mutator transaction binding the contract method 0xb204be93.
//
// Solidity: function onHanaMessage((bytes,uint256,address,uint256,bytes) hanaMessage) returns()
func (_HanaReceiver *HanaReceiverTransactor) OnHanaMessage(opts *bind.TransactOpts, hanaMessage HanaInterfacesHanaMessage) (*types.Transaction, error) {
	return _HanaReceiver.contract.Transact(opts, "onHanaMessage", hanaMessage)
}

// OnHanaMessage is a paid mutator transaction binding the contract method 0xb204be93.
//
// Solidity: function onHanaMessage((bytes,uint256,address,uint256,bytes) hanaMessage) returns()
func (_HanaReceiver *HanaReceiverSession) OnHanaMessage(hanaMessage HanaInterfacesHanaMessage) (*types.Transaction, error) {
	return _HanaReceiver.Contract.OnHanaMessage(&_HanaReceiver.TransactOpts, hanaMessage)
}

// OnHanaMessage is a paid mutator transaction binding the contract method 0xb204be93.
//
// Solidity: function onHanaMessage((bytes,uint256,address,uint256,bytes) hanaMessage) returns()
func (_HanaReceiver *HanaReceiverTransactorSession) OnHanaMessage(hanaMessage HanaInterfacesHanaMessage) (*types.Transaction, error) {
	return _HanaReceiver.Contract.OnHanaMessage(&_HanaReceiver.TransactOpts, hanaMessage)
}

// OnHanaRevert is a paid mutator transaction binding the contract method 0x4204cf9b.
//
// Solidity: function onHanaRevert((address,uint256,bytes,uint256,uint256,bytes) hanaRevert) returns()
func (_HanaReceiver *HanaReceiverTransactor) OnHanaRevert(opts *bind.TransactOpts, hanaRevert HanaInterfacesHanaRevert) (*types.Transaction, error) {
	return _HanaReceiver.contract.Transact(opts, "onHanaRevert", hanaRevert)
}

// OnHanaRevert is a paid mutator transaction binding the contract method 0x4204cf9b.
//
// Solidity: function onHanaRevert((address,uint256,bytes,uint256,uint256,bytes) hanaRevert) returns()
func (_HanaReceiver *HanaReceiverSession) OnHanaRevert(hanaRevert HanaInterfacesHanaRevert) (*types.Transaction, error) {
	return _HanaReceiver.Contract.OnHanaRevert(&_HanaReceiver.TransactOpts, hanaRevert)
}

// OnHanaRevert is a paid mutator transaction binding the contract method 0x4204cf9b.
//
// Solidity: function onHanaRevert((address,uint256,bytes,uint256,uint256,bytes) hanaRevert) returns()
func (_HanaReceiver *HanaReceiverTransactorSession) OnHanaRevert(hanaRevert HanaInterfacesHanaRevert) (*types.Transaction, error) {
	return _HanaReceiver.Contract.OnHanaRevert(&_HanaReceiver.TransactOpts, hanaRevert)
}
