// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanainteractor

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

// HanaInteractorMetaData contains all meta data concerning the HanaInteractor contract.
var HanaInteractorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidHanaMessageCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidHanaRevertCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connector\",\"outputs\":[{\"internalType\":\"contractHanaConnector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"interactorsByChainId\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"contractAddress\",\"type\":\"bytes\"}],\"name\":\"setInteractorByChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HanaInteractorABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaInteractorMetaData.ABI instead.
var HanaInteractorABI = HanaInteractorMetaData.ABI

// HanaInteractor is an auto generated Go binding around an Ethereum contract.
type HanaInteractor struct {
	HanaInteractorCaller     // Read-only binding to the contract
	HanaInteractorTransactor // Write-only binding to the contract
	HanaInteractorFilterer   // Log filterer for contract events
}

// HanaInteractorCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaInteractorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaInteractorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaInteractorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaInteractorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaInteractorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaInteractorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaInteractorSession struct {
	Contract     *HanaInteractor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HanaInteractorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaInteractorCallerSession struct {
	Contract *HanaInteractorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HanaInteractorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaInteractorTransactorSession struct {
	Contract     *HanaInteractorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HanaInteractorRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaInteractorRaw struct {
	Contract *HanaInteractor // Generic contract binding to access the raw methods on
}

// HanaInteractorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaInteractorCallerRaw struct {
	Contract *HanaInteractorCaller // Generic read-only contract binding to access the raw methods on
}

// HanaInteractorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaInteractorTransactorRaw struct {
	Contract *HanaInteractorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaInteractor creates a new instance of HanaInteractor, bound to a specific deployed contract.
func NewHanaInteractor(address common.Address, backend bind.ContractBackend) (*HanaInteractor, error) {
	contract, err := bindHanaInteractor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaInteractor{HanaInteractorCaller: HanaInteractorCaller{contract: contract}, HanaInteractorTransactor: HanaInteractorTransactor{contract: contract}, HanaInteractorFilterer: HanaInteractorFilterer{contract: contract}}, nil
}

// NewHanaInteractorCaller creates a new read-only instance of HanaInteractor, bound to a specific deployed contract.
func NewHanaInteractorCaller(address common.Address, caller bind.ContractCaller) (*HanaInteractorCaller, error) {
	contract, err := bindHanaInteractor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorCaller{contract: contract}, nil
}

// NewHanaInteractorTransactor creates a new write-only instance of HanaInteractor, bound to a specific deployed contract.
func NewHanaInteractorTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaInteractorTransactor, error) {
	contract, err := bindHanaInteractor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorTransactor{contract: contract}, nil
}

// NewHanaInteractorFilterer creates a new log filterer instance of HanaInteractor, bound to a specific deployed contract.
func NewHanaInteractorFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaInteractorFilterer, error) {
	contract, err := bindHanaInteractor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorFilterer{contract: contract}, nil
}

// bindHanaInteractor binds a generic wrapper to an already deployed contract.
func bindHanaInteractor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaInteractorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaInteractor *HanaInteractorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaInteractor.Contract.HanaInteractorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaInteractor *HanaInteractorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaInteractor.Contract.HanaInteractorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaInteractor *HanaInteractorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaInteractor.Contract.HanaInteractorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaInteractor *HanaInteractorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaInteractor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaInteractor *HanaInteractorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaInteractor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaInteractor *HanaInteractorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaInteractor.Contract.contract.Transact(opts, method, params...)
}

// Connector is a free data retrieval call binding the contract method 0x83f3084f.
//
// Solidity: function connector() view returns(address)
func (_HanaInteractor *HanaInteractorCaller) Connector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaInteractor.contract.Call(opts, &out, "connector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Connector is a free data retrieval call binding the contract method 0x83f3084f.
//
// Solidity: function connector() view returns(address)
func (_HanaInteractor *HanaInteractorSession) Connector() (common.Address, error) {
	return _HanaInteractor.Contract.Connector(&_HanaInteractor.CallOpts)
}

// Connector is a free data retrieval call binding the contract method 0x83f3084f.
//
// Solidity: function connector() view returns(address)
func (_HanaInteractor *HanaInteractorCallerSession) Connector() (common.Address, error) {
	return _HanaInteractor.Contract.Connector(&_HanaInteractor.CallOpts)
}

// InteractorsByChainId is a free data retrieval call binding the contract method 0x2618143f.
//
// Solidity: function interactorsByChainId(uint256 ) view returns(bytes)
func (_HanaInteractor *HanaInteractorCaller) InteractorsByChainId(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _HanaInteractor.contract.Call(opts, &out, "interactorsByChainId", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// InteractorsByChainId is a free data retrieval call binding the contract method 0x2618143f.
//
// Solidity: function interactorsByChainId(uint256 ) view returns(bytes)
func (_HanaInteractor *HanaInteractorSession) InteractorsByChainId(arg0 *big.Int) ([]byte, error) {
	return _HanaInteractor.Contract.InteractorsByChainId(&_HanaInteractor.CallOpts, arg0)
}

// InteractorsByChainId is a free data retrieval call binding the contract method 0x2618143f.
//
// Solidity: function interactorsByChainId(uint256 ) view returns(bytes)
func (_HanaInteractor *HanaInteractorCallerSession) InteractorsByChainId(arg0 *big.Int) ([]byte, error) {
	return _HanaInteractor.Contract.InteractorsByChainId(&_HanaInteractor.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HanaInteractor *HanaInteractorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaInteractor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HanaInteractor *HanaInteractorSession) Owner() (common.Address, error) {
	return _HanaInteractor.Contract.Owner(&_HanaInteractor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HanaInteractor *HanaInteractorCallerSession) Owner() (common.Address, error) {
	return _HanaInteractor.Contract.Owner(&_HanaInteractor.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_HanaInteractor *HanaInteractorCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaInteractor.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_HanaInteractor *HanaInteractorSession) PendingOwner() (common.Address, error) {
	return _HanaInteractor.Contract.PendingOwner(&_HanaInteractor.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_HanaInteractor *HanaInteractorCallerSession) PendingOwner() (common.Address, error) {
	return _HanaInteractor.Contract.PendingOwner(&_HanaInteractor.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_HanaInteractor *HanaInteractorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaInteractor.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_HanaInteractor *HanaInteractorSession) AcceptOwnership() (*types.Transaction, error) {
	return _HanaInteractor.Contract.AcceptOwnership(&_HanaInteractor.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_HanaInteractor *HanaInteractorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _HanaInteractor.Contract.AcceptOwnership(&_HanaInteractor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HanaInteractor *HanaInteractorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaInteractor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HanaInteractor *HanaInteractorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HanaInteractor.Contract.RenounceOwnership(&_HanaInteractor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HanaInteractor *HanaInteractorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HanaInteractor.Contract.RenounceOwnership(&_HanaInteractor.TransactOpts)
}

// SetInteractorByChainId is a paid mutator transaction binding the contract method 0x4fd3f7d7.
//
// Solidity: function setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) returns()
func (_HanaInteractor *HanaInteractorTransactor) SetInteractorByChainId(opts *bind.TransactOpts, destinationChainId *big.Int, contractAddress []byte) (*types.Transaction, error) {
	return _HanaInteractor.contract.Transact(opts, "setInteractorByChainId", destinationChainId, contractAddress)
}

// SetInteractorByChainId is a paid mutator transaction binding the contract method 0x4fd3f7d7.
//
// Solidity: function setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) returns()
func (_HanaInteractor *HanaInteractorSession) SetInteractorByChainId(destinationChainId *big.Int, contractAddress []byte) (*types.Transaction, error) {
	return _HanaInteractor.Contract.SetInteractorByChainId(&_HanaInteractor.TransactOpts, destinationChainId, contractAddress)
}

// SetInteractorByChainId is a paid mutator transaction binding the contract method 0x4fd3f7d7.
//
// Solidity: function setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) returns()
func (_HanaInteractor *HanaInteractorTransactorSession) SetInteractorByChainId(destinationChainId *big.Int, contractAddress []byte) (*types.Transaction, error) {
	return _HanaInteractor.Contract.SetInteractorByChainId(&_HanaInteractor.TransactOpts, destinationChainId, contractAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HanaInteractor *HanaInteractorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HanaInteractor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HanaInteractor *HanaInteractorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HanaInteractor.Contract.TransferOwnership(&_HanaInteractor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HanaInteractor *HanaInteractorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HanaInteractor.Contract.TransferOwnership(&_HanaInteractor.TransactOpts, newOwner)
}

// HanaInteractorOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the HanaInteractor contract.
type HanaInteractorOwnershipTransferStartedIterator struct {
	Event *HanaInteractorOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HanaInteractorOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaInteractorOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HanaInteractorOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HanaInteractorOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaInteractorOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaInteractorOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the HanaInteractor contract.
type HanaInteractorOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_HanaInteractor *HanaInteractorFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HanaInteractorOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HanaInteractor.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorOwnershipTransferStartedIterator{contract: _HanaInteractor.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_HanaInteractor *HanaInteractorFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *HanaInteractorOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HanaInteractor.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaInteractorOwnershipTransferStarted)
				if err := _HanaInteractor.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_HanaInteractor *HanaInteractorFilterer) ParseOwnershipTransferStarted(log types.Log) (*HanaInteractorOwnershipTransferStarted, error) {
	event := new(HanaInteractorOwnershipTransferStarted)
	if err := _HanaInteractor.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaInteractorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HanaInteractor contract.
type HanaInteractorOwnershipTransferredIterator struct {
	Event *HanaInteractorOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HanaInteractorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaInteractorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HanaInteractorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HanaInteractorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaInteractorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaInteractorOwnershipTransferred represents a OwnershipTransferred event raised by the HanaInteractor contract.
type HanaInteractorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HanaInteractor *HanaInteractorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HanaInteractorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HanaInteractor.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorOwnershipTransferredIterator{contract: _HanaInteractor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HanaInteractor *HanaInteractorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HanaInteractorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HanaInteractor.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaInteractorOwnershipTransferred)
				if err := _HanaInteractor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HanaInteractor *HanaInteractorFilterer) ParseOwnershipTransferred(log types.Log) (*HanaInteractorOwnershipTransferred, error) {
	event := new(HanaInteractorOwnershipTransferred)
	if err := _HanaInteractor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
