// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanareceivermock

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

// HanaReceiverMockMetaData contains all meta data concerning the HanaReceiverMock contract.
var HanaReceiverMockMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"}],\"name\":\"MockOnHanaMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"}],\"name\":\"MockOnHanaRevert\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"hanaTxSenderAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hanaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structHanaInterfaces.HanaMessage\",\"name\":\"hanaMessage\",\"type\":\"tuple\"}],\"name\":\"onHanaMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingHanaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structHanaInterfaces.HanaRevert\",\"name\":\"hanaRevert\",\"type\":\"tuple\"}],\"name\":\"onHanaRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506102d5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634204cf9b1461003b578063b204be9314610057575b600080fd5b610055600480360381019061005091906101d4565b610073565b005b610071600480360381019061006c919061018b565b6100bf565b005b7f038b3b35b75bcab7cbeac0fe3b4c1b6dd122dff390d20be23e270a89d7d7550c8160000160208101906100a7919061015e565b6040516100b4919061022c565b60405180910390a150565b7f5918e0a29c5ec5a8cb80256eba14d26ff045577fee0e40f47b15f0cbdd8593518160400160208101906100f3919061015e565b604051610100919061022c565b60405180910390a150565b60008135905061011a81610288565b92915050565b600060a0828403121561013657610135610279565b5b81905092915050565b600060c0828403121561015557610154610279565b5b81905092915050565b60006020828403121561017457610173610283565b5b60006101828482850161010b565b91505092915050565b6000602082840312156101a1576101a0610283565b5b600082013567ffffffffffffffff8111156101bf576101be61027e565b5b6101cb84828501610120565b91505092915050565b6000602082840312156101ea576101e9610283565b5b600082013567ffffffffffffffff8111156102085761020761027e565b5b6102148482850161013f565b91505092915050565b61022681610247565b82525050565b6000602082019050610241600083018461021d565b92915050565b600061025282610259565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b600080fd5b600080fd5b61029181610247565b811461029c57600080fd5b5056fea2646970667358221220b979fb2e6713570cb53b79bcd103aa8c32f927cf7d380b5543504e7b85fbc90a64736f6c63430008070033",
}

// HanaReceiverMockABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaReceiverMockMetaData.ABI instead.
var HanaReceiverMockABI = HanaReceiverMockMetaData.ABI

// HanaReceiverMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HanaReceiverMockMetaData.Bin instead.
var HanaReceiverMockBin = HanaReceiverMockMetaData.Bin

// DeployHanaReceiverMock deploys a new Ethereum contract, binding an instance of HanaReceiverMock to it.
func DeployHanaReceiverMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HanaReceiverMock, error) {
	parsed, err := HanaReceiverMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HanaReceiverMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HanaReceiverMock{HanaReceiverMockCaller: HanaReceiverMockCaller{contract: contract}, HanaReceiverMockTransactor: HanaReceiverMockTransactor{contract: contract}, HanaReceiverMockFilterer: HanaReceiverMockFilterer{contract: contract}}, nil
}

// HanaReceiverMock is an auto generated Go binding around an Ethereum contract.
type HanaReceiverMock struct {
	HanaReceiverMockCaller     // Read-only binding to the contract
	HanaReceiverMockTransactor // Write-only binding to the contract
	HanaReceiverMockFilterer   // Log filterer for contract events
}

// HanaReceiverMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaReceiverMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaReceiverMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaReceiverMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaReceiverMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaReceiverMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaReceiverMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaReceiverMockSession struct {
	Contract     *HanaReceiverMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HanaReceiverMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaReceiverMockCallerSession struct {
	Contract *HanaReceiverMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// HanaReceiverMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaReceiverMockTransactorSession struct {
	Contract     *HanaReceiverMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// HanaReceiverMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaReceiverMockRaw struct {
	Contract *HanaReceiverMock // Generic contract binding to access the raw methods on
}

// HanaReceiverMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaReceiverMockCallerRaw struct {
	Contract *HanaReceiverMockCaller // Generic read-only contract binding to access the raw methods on
}

// HanaReceiverMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaReceiverMockTransactorRaw struct {
	Contract *HanaReceiverMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaReceiverMock creates a new instance of HanaReceiverMock, bound to a specific deployed contract.
func NewHanaReceiverMock(address common.Address, backend bind.ContractBackend) (*HanaReceiverMock, error) {
	contract, err := bindHanaReceiverMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaReceiverMock{HanaReceiverMockCaller: HanaReceiverMockCaller{contract: contract}, HanaReceiverMockTransactor: HanaReceiverMockTransactor{contract: contract}, HanaReceiverMockFilterer: HanaReceiverMockFilterer{contract: contract}}, nil
}

// NewHanaReceiverMockCaller creates a new read-only instance of HanaReceiverMock, bound to a specific deployed contract.
func NewHanaReceiverMockCaller(address common.Address, caller bind.ContractCaller) (*HanaReceiverMockCaller, error) {
	contract, err := bindHanaReceiverMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaReceiverMockCaller{contract: contract}, nil
}

// NewHanaReceiverMockTransactor creates a new write-only instance of HanaReceiverMock, bound to a specific deployed contract.
func NewHanaReceiverMockTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaReceiverMockTransactor, error) {
	contract, err := bindHanaReceiverMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaReceiverMockTransactor{contract: contract}, nil
}

// NewHanaReceiverMockFilterer creates a new log filterer instance of HanaReceiverMock, bound to a specific deployed contract.
func NewHanaReceiverMockFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaReceiverMockFilterer, error) {
	contract, err := bindHanaReceiverMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaReceiverMockFilterer{contract: contract}, nil
}

// bindHanaReceiverMock binds a generic wrapper to an already deployed contract.
func bindHanaReceiverMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaReceiverMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaReceiverMock *HanaReceiverMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaReceiverMock.Contract.HanaReceiverMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaReceiverMock *HanaReceiverMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaReceiverMock.Contract.HanaReceiverMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaReceiverMock *HanaReceiverMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaReceiverMock.Contract.HanaReceiverMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaReceiverMock *HanaReceiverMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaReceiverMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaReceiverMock *HanaReceiverMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaReceiverMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaReceiverMock *HanaReceiverMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaReceiverMock.Contract.contract.Transact(opts, method, params...)
}

// OnHanaMessage is a paid mutator transaction binding the contract method 0xb204be93.
//
// Solidity: function onHanaMessage((bytes,uint256,address,uint256,bytes) hanaMessage) returns()
func (_HanaReceiverMock *HanaReceiverMockTransactor) OnHanaMessage(opts *bind.TransactOpts, hanaMessage HanaInterfacesHanaMessage) (*types.Transaction, error) {
	return _HanaReceiverMock.contract.Transact(opts, "onHanaMessage", hanaMessage)
}

// OnHanaMessage is a paid mutator transaction binding the contract method 0xb204be93.
//
// Solidity: function onHanaMessage((bytes,uint256,address,uint256,bytes) hanaMessage) returns()
func (_HanaReceiverMock *HanaReceiverMockSession) OnHanaMessage(hanaMessage HanaInterfacesHanaMessage) (*types.Transaction, error) {
	return _HanaReceiverMock.Contract.OnHanaMessage(&_HanaReceiverMock.TransactOpts, hanaMessage)
}

// OnHanaMessage is a paid mutator transaction binding the contract method 0xb204be93.
//
// Solidity: function onHanaMessage((bytes,uint256,address,uint256,bytes) hanaMessage) returns()
func (_HanaReceiverMock *HanaReceiverMockTransactorSession) OnHanaMessage(hanaMessage HanaInterfacesHanaMessage) (*types.Transaction, error) {
	return _HanaReceiverMock.Contract.OnHanaMessage(&_HanaReceiverMock.TransactOpts, hanaMessage)
}

// OnHanaRevert is a paid mutator transaction binding the contract method 0x4204cf9b.
//
// Solidity: function onHanaRevert((address,uint256,bytes,uint256,uint256,bytes) hanaRevert) returns()
func (_HanaReceiverMock *HanaReceiverMockTransactor) OnHanaRevert(opts *bind.TransactOpts, hanaRevert HanaInterfacesHanaRevert) (*types.Transaction, error) {
	return _HanaReceiverMock.contract.Transact(opts, "onHanaRevert", hanaRevert)
}

// OnHanaRevert is a paid mutator transaction binding the contract method 0x4204cf9b.
//
// Solidity: function onHanaRevert((address,uint256,bytes,uint256,uint256,bytes) hanaRevert) returns()
func (_HanaReceiverMock *HanaReceiverMockSession) OnHanaRevert(hanaRevert HanaInterfacesHanaRevert) (*types.Transaction, error) {
	return _HanaReceiverMock.Contract.OnHanaRevert(&_HanaReceiverMock.TransactOpts, hanaRevert)
}

// OnHanaRevert is a paid mutator transaction binding the contract method 0x4204cf9b.
//
// Solidity: function onHanaRevert((address,uint256,bytes,uint256,uint256,bytes) hanaRevert) returns()
func (_HanaReceiverMock *HanaReceiverMockTransactorSession) OnHanaRevert(hanaRevert HanaInterfacesHanaRevert) (*types.Transaction, error) {
	return _HanaReceiverMock.Contract.OnHanaRevert(&_HanaReceiverMock.TransactOpts, hanaRevert)
}

// HanaReceiverMockMockOnHanaMessageIterator is returned from FilterMockOnHanaMessage and is used to iterate over the raw logs and unpacked data for MockOnHanaMessage events raised by the HanaReceiverMock contract.
type HanaReceiverMockMockOnHanaMessageIterator struct {
	Event *HanaReceiverMockMockOnHanaMessage // Event containing the contract specifics and raw log

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
func (it *HanaReceiverMockMockOnHanaMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaReceiverMockMockOnHanaMessage)
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
		it.Event = new(HanaReceiverMockMockOnHanaMessage)
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
func (it *HanaReceiverMockMockOnHanaMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaReceiverMockMockOnHanaMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaReceiverMockMockOnHanaMessage represents a MockOnHanaMessage event raised by the HanaReceiverMock contract.
type HanaReceiverMockMockOnHanaMessage struct {
	DestinationAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMockOnHanaMessage is a free log retrieval operation binding the contract event 0x5918e0a29c5ec5a8cb80256eba14d26ff045577fee0e40f47b15f0cbdd859351.
//
// Solidity: event MockOnHanaMessage(address destinationAddress)
func (_HanaReceiverMock *HanaReceiverMockFilterer) FilterMockOnHanaMessage(opts *bind.FilterOpts) (*HanaReceiverMockMockOnHanaMessageIterator, error) {

	logs, sub, err := _HanaReceiverMock.contract.FilterLogs(opts, "MockOnHanaMessage")
	if err != nil {
		return nil, err
	}
	return &HanaReceiverMockMockOnHanaMessageIterator{contract: _HanaReceiverMock.contract, event: "MockOnHanaMessage", logs: logs, sub: sub}, nil
}

// WatchMockOnHanaMessage is a free log subscription operation binding the contract event 0x5918e0a29c5ec5a8cb80256eba14d26ff045577fee0e40f47b15f0cbdd859351.
//
// Solidity: event MockOnHanaMessage(address destinationAddress)
func (_HanaReceiverMock *HanaReceiverMockFilterer) WatchMockOnHanaMessage(opts *bind.WatchOpts, sink chan<- *HanaReceiverMockMockOnHanaMessage) (event.Subscription, error) {

	logs, sub, err := _HanaReceiverMock.contract.WatchLogs(opts, "MockOnHanaMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaReceiverMockMockOnHanaMessage)
				if err := _HanaReceiverMock.contract.UnpackLog(event, "MockOnHanaMessage", log); err != nil {
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

// ParseMockOnHanaMessage is a log parse operation binding the contract event 0x5918e0a29c5ec5a8cb80256eba14d26ff045577fee0e40f47b15f0cbdd859351.
//
// Solidity: event MockOnHanaMessage(address destinationAddress)
func (_HanaReceiverMock *HanaReceiverMockFilterer) ParseMockOnHanaMessage(log types.Log) (*HanaReceiverMockMockOnHanaMessage, error) {
	event := new(HanaReceiverMockMockOnHanaMessage)
	if err := _HanaReceiverMock.contract.UnpackLog(event, "MockOnHanaMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaReceiverMockMockOnHanaRevertIterator is returned from FilterMockOnHanaRevert and is used to iterate over the raw logs and unpacked data for MockOnHanaRevert events raised by the HanaReceiverMock contract.
type HanaReceiverMockMockOnHanaRevertIterator struct {
	Event *HanaReceiverMockMockOnHanaRevert // Event containing the contract specifics and raw log

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
func (it *HanaReceiverMockMockOnHanaRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaReceiverMockMockOnHanaRevert)
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
		it.Event = new(HanaReceiverMockMockOnHanaRevert)
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
func (it *HanaReceiverMockMockOnHanaRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaReceiverMockMockOnHanaRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaReceiverMockMockOnHanaRevert represents a MockOnHanaRevert event raised by the HanaReceiverMock contract.
type HanaReceiverMockMockOnHanaRevert struct {
	HanaTxSenderAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMockOnHanaRevert is a free log retrieval operation binding the contract event 0x038b3b35b75bcab7cbeac0fe3b4c1b6dd122dff390d20be23e270a89d7d7550c.
//
// Solidity: event MockOnHanaRevert(address hanaTxSenderAddress)
func (_HanaReceiverMock *HanaReceiverMockFilterer) FilterMockOnHanaRevert(opts *bind.FilterOpts) (*HanaReceiverMockMockOnHanaRevertIterator, error) {

	logs, sub, err := _HanaReceiverMock.contract.FilterLogs(opts, "MockOnHanaRevert")
	if err != nil {
		return nil, err
	}
	return &HanaReceiverMockMockOnHanaRevertIterator{contract: _HanaReceiverMock.contract, event: "MockOnHanaRevert", logs: logs, sub: sub}, nil
}

// WatchMockOnHanaRevert is a free log subscription operation binding the contract event 0x038b3b35b75bcab7cbeac0fe3b4c1b6dd122dff390d20be23e270a89d7d7550c.
//
// Solidity: event MockOnHanaRevert(address hanaTxSenderAddress)
func (_HanaReceiverMock *HanaReceiverMockFilterer) WatchMockOnHanaRevert(opts *bind.WatchOpts, sink chan<- *HanaReceiverMockMockOnHanaRevert) (event.Subscription, error) {

	logs, sub, err := _HanaReceiverMock.contract.WatchLogs(opts, "MockOnHanaRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaReceiverMockMockOnHanaRevert)
				if err := _HanaReceiverMock.contract.UnpackLog(event, "MockOnHanaRevert", log); err != nil {
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

// ParseMockOnHanaRevert is a log parse operation binding the contract event 0x038b3b35b75bcab7cbeac0fe3b4c1b6dd122dff390d20be23e270a89d7d7550c.
//
// Solidity: event MockOnHanaRevert(address hanaTxSenderAddress)
func (_HanaReceiverMock *HanaReceiverMockFilterer) ParseMockOnHanaRevert(log types.Log) (*HanaReceiverMockMockOnHanaRevert, error) {
	event := new(HanaReceiverMockMockOnHanaRevert)
	if err := _HanaReceiverMock.contract.UnpackLog(event, "MockOnHanaRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
