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

// HanaTokenConsumerMetaData contains all meta data concerning the HanaTokenConsumer contract.
var HanaTokenConsumerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthExchangedForHana\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"HanaExchangedForEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"HanaExchangedForToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenExchangedForHana\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hanaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getEthFromHana\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"getHanaFromEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getHanaFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hanaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenFromHana\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasHanaLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// HanaTokenConsumerABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaTokenConsumerMetaData.ABI instead.
var HanaTokenConsumerABI = HanaTokenConsumerMetaData.ABI

// HanaTokenConsumer is an auto generated Go binding around an Ethereum contract.
type HanaTokenConsumer struct {
	HanaTokenConsumerCaller     // Read-only binding to the contract
	HanaTokenConsumerTransactor // Write-only binding to the contract
	HanaTokenConsumerFilterer   // Log filterer for contract events
}

// HanaTokenConsumerCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaTokenConsumerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaTokenConsumerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaTokenConsumerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaTokenConsumerSession struct {
	Contract     *HanaTokenConsumer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HanaTokenConsumerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaTokenConsumerCallerSession struct {
	Contract *HanaTokenConsumerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// HanaTokenConsumerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaTokenConsumerTransactorSession struct {
	Contract     *HanaTokenConsumerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// HanaTokenConsumerRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaTokenConsumerRaw struct {
	Contract *HanaTokenConsumer // Generic contract binding to access the raw methods on
}

// HanaTokenConsumerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaTokenConsumerCallerRaw struct {
	Contract *HanaTokenConsumerCaller // Generic read-only contract binding to access the raw methods on
}

// HanaTokenConsumerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaTokenConsumerTransactorRaw struct {
	Contract *HanaTokenConsumerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaTokenConsumer creates a new instance of HanaTokenConsumer, bound to a specific deployed contract.
func NewHanaTokenConsumer(address common.Address, backend bind.ContractBackend) (*HanaTokenConsumer, error) {
	contract, err := bindHanaTokenConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumer{HanaTokenConsumerCaller: HanaTokenConsumerCaller{contract: contract}, HanaTokenConsumerTransactor: HanaTokenConsumerTransactor{contract: contract}, HanaTokenConsumerFilterer: HanaTokenConsumerFilterer{contract: contract}}, nil
}

// NewHanaTokenConsumerCaller creates a new read-only instance of HanaTokenConsumer, bound to a specific deployed contract.
func NewHanaTokenConsumerCaller(address common.Address, caller bind.ContractCaller) (*HanaTokenConsumerCaller, error) {
	contract, err := bindHanaTokenConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerCaller{contract: contract}, nil
}

// NewHanaTokenConsumerTransactor creates a new write-only instance of HanaTokenConsumer, bound to a specific deployed contract.
func NewHanaTokenConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaTokenConsumerTransactor, error) {
	contract, err := bindHanaTokenConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerTransactor{contract: contract}, nil
}

// NewHanaTokenConsumerFilterer creates a new log filterer instance of HanaTokenConsumer, bound to a specific deployed contract.
func NewHanaTokenConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaTokenConsumerFilterer, error) {
	contract, err := bindHanaTokenConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerFilterer{contract: contract}, nil
}

// bindHanaTokenConsumer binds a generic wrapper to an already deployed contract.
func bindHanaTokenConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaTokenConsumerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaTokenConsumer *HanaTokenConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaTokenConsumer.Contract.HanaTokenConsumerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaTokenConsumer *HanaTokenConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumer.Contract.HanaTokenConsumerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaTokenConsumer *HanaTokenConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaTokenConsumer.Contract.HanaTokenConsumerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaTokenConsumer *HanaTokenConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaTokenConsumer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaTokenConsumer *HanaTokenConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaTokenConsumer *HanaTokenConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaTokenConsumer.Contract.contract.Transact(opts, method, params...)
}

// HasHanaLiquidity is a free data retrieval call binding the contract method 0x246d567e.
//
// Solidity: function hasHanaLiquidity() view returns(bool)
func (_HanaTokenConsumer *HanaTokenConsumerCaller) HasHanaLiquidity(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HanaTokenConsumer.contract.Call(opts, &out, "hasHanaLiquidity")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasHanaLiquidity is a free data retrieval call binding the contract method 0x246d567e.
//
// Solidity: function hasHanaLiquidity() view returns(bool)
func (_HanaTokenConsumer *HanaTokenConsumerSession) HasHanaLiquidity() (bool, error) {
	return _HanaTokenConsumer.Contract.HasHanaLiquidity(&_HanaTokenConsumer.CallOpts)
}

// HasHanaLiquidity is a free data retrieval call binding the contract method 0x246d567e.
//
// Solidity: function hasHanaLiquidity() view returns(bool)
func (_HanaTokenConsumer *HanaTokenConsumerCallerSession) HasHanaLiquidity() (bool, error) {
	return _HanaTokenConsumer.Contract.HasHanaLiquidity(&_HanaTokenConsumer.CallOpts)
}

// GetEthFromHana is a paid mutator transaction binding the contract method 0x291d55f7.
//
// Solidity: function getEthFromHana(address destinationAddress, uint256 minAmountOut, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumer *HanaTokenConsumerTransactor) GetEthFromHana(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumer.contract.Transact(opts, "getEthFromHana", destinationAddress, minAmountOut, hanaTokenAmount)
}

// GetEthFromHana is a paid mutator transaction binding the contract method 0x291d55f7.
//
// Solidity: function getEthFromHana(address destinationAddress, uint256 minAmountOut, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumer *HanaTokenConsumerSession) GetEthFromHana(destinationAddress common.Address, minAmountOut *big.Int, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumer.Contract.GetEthFromHana(&_HanaTokenConsumer.TransactOpts, destinationAddress, minAmountOut, hanaTokenAmount)
}

// GetEthFromHana is a paid mutator transaction binding the contract method 0x291d55f7.
//
// Solidity: function getEthFromHana(address destinationAddress, uint256 minAmountOut, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumer *HanaTokenConsumerTransactorSession) GetEthFromHana(destinationAddress common.Address, minAmountOut *big.Int, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumer.Contract.GetEthFromHana(&_HanaTokenConsumer.TransactOpts, destinationAddress, minAmountOut, hanaTokenAmount)
}

// GetHanaFromEth is a paid mutator transaction binding the contract method 0x45413df7.
//
// Solidity: function getHanaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_HanaTokenConsumer *HanaTokenConsumerTransactor) GetHanaFromEth(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumer.contract.Transact(opts, "getHanaFromEth", destinationAddress, minAmountOut)
}

// GetHanaFromEth is a paid mutator transaction binding the contract method 0x45413df7.
//
// Solidity: function getHanaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_HanaTokenConsumer *HanaTokenConsumerSession) GetHanaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumer.Contract.GetHanaFromEth(&_HanaTokenConsumer.TransactOpts, destinationAddress, minAmountOut)
}

// GetHanaFromEth is a paid mutator transaction binding the contract method 0x45413df7.
//
// Solidity: function getHanaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_HanaTokenConsumer *HanaTokenConsumerTransactorSession) GetHanaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumer.Contract.GetHanaFromEth(&_HanaTokenConsumer.TransactOpts, destinationAddress, minAmountOut)
}

// GetHanaFromToken is a paid mutator transaction binding the contract method 0x65d6edc8.
//
// Solidity: function getHanaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_HanaTokenConsumer *HanaTokenConsumerTransactor) GetHanaFromToken(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumer.contract.Transact(opts, "getHanaFromToken", destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetHanaFromToken is a paid mutator transaction binding the contract method 0x65d6edc8.
//
// Solidity: function getHanaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_HanaTokenConsumer *HanaTokenConsumerSession) GetHanaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumer.Contract.GetHanaFromToken(&_HanaTokenConsumer.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetHanaFromToken is a paid mutator transaction binding the contract method 0x65d6edc8.
//
// Solidity: function getHanaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_HanaTokenConsumer *HanaTokenConsumerTransactorSession) GetHanaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumer.Contract.GetHanaFromToken(&_HanaTokenConsumer.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetTokenFromHana is a paid mutator transaction binding the contract method 0x5c2fec9a.
//
// Solidity: function getTokenFromHana(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumer *HanaTokenConsumerTransactor) GetTokenFromHana(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumer.contract.Transact(opts, "getTokenFromHana", destinationAddress, minAmountOut, outputToken, hanaTokenAmount)
}

// GetTokenFromHana is a paid mutator transaction binding the contract method 0x5c2fec9a.
//
// Solidity: function getTokenFromHana(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumer *HanaTokenConsumerSession) GetTokenFromHana(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumer.Contract.GetTokenFromHana(&_HanaTokenConsumer.TransactOpts, destinationAddress, minAmountOut, outputToken, hanaTokenAmount)
}

// GetTokenFromHana is a paid mutator transaction binding the contract method 0x5c2fec9a.
//
// Solidity: function getTokenFromHana(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumer *HanaTokenConsumerTransactorSession) GetTokenFromHana(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumer.Contract.GetTokenFromHana(&_HanaTokenConsumer.TransactOpts, destinationAddress, minAmountOut, outputToken, hanaTokenAmount)
}

// HanaTokenConsumerEthExchangedForHanaIterator is returned from FilterEthExchangedForHana and is used to iterate over the raw logs and unpacked data for EthExchangedForHana events raised by the HanaTokenConsumer contract.
type HanaTokenConsumerEthExchangedForHanaIterator struct {
	Event *HanaTokenConsumerEthExchangedForHana // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerEthExchangedForHanaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerEthExchangedForHana)
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
		it.Event = new(HanaTokenConsumerEthExchangedForHana)
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
func (it *HanaTokenConsumerEthExchangedForHanaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerEthExchangedForHanaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerEthExchangedForHana represents a EthExchangedForHana event raised by the HanaTokenConsumer contract.
type HanaTokenConsumerEthExchangedForHana struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthExchangedForHana is a free log retrieval operation binding the contract event 0x877de62ce9cb129480429d4e4ac5b3ce3b8b85638de13f817d02b9cea87a0f36.
//
// Solidity: event EthExchangedForHana(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumer *HanaTokenConsumerFilterer) FilterEthExchangedForHana(opts *bind.FilterOpts) (*HanaTokenConsumerEthExchangedForHanaIterator, error) {

	logs, sub, err := _HanaTokenConsumer.contract.FilterLogs(opts, "EthExchangedForHana")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerEthExchangedForHanaIterator{contract: _HanaTokenConsumer.contract, event: "EthExchangedForHana", logs: logs, sub: sub}, nil
}

// WatchEthExchangedForHana is a free log subscription operation binding the contract event 0x877de62ce9cb129480429d4e4ac5b3ce3b8b85638de13f817d02b9cea87a0f36.
//
// Solidity: event EthExchangedForHana(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumer *HanaTokenConsumerFilterer) WatchEthExchangedForHana(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerEthExchangedForHana) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumer.contract.WatchLogs(opts, "EthExchangedForHana")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerEthExchangedForHana)
				if err := _HanaTokenConsumer.contract.UnpackLog(event, "EthExchangedForHana", log); err != nil {
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

// ParseEthExchangedForHana is a log parse operation binding the contract event 0x877de62ce9cb129480429d4e4ac5b3ce3b8b85638de13f817d02b9cea87a0f36.
//
// Solidity: event EthExchangedForHana(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumer *HanaTokenConsumerFilterer) ParseEthExchangedForHana(log types.Log) (*HanaTokenConsumerEthExchangedForHana, error) {
	event := new(HanaTokenConsumerEthExchangedForHana)
	if err := _HanaTokenConsumer.contract.UnpackLog(event, "EthExchangedForHana", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaTokenConsumerHanaExchangedForEthIterator is returned from FilterHanaExchangedForEth and is used to iterate over the raw logs and unpacked data for HanaExchangedForEth events raised by the HanaTokenConsumer contract.
type HanaTokenConsumerHanaExchangedForEthIterator struct {
	Event *HanaTokenConsumerHanaExchangedForEth // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerHanaExchangedForEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerHanaExchangedForEth)
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
		it.Event = new(HanaTokenConsumerHanaExchangedForEth)
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
func (it *HanaTokenConsumerHanaExchangedForEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerHanaExchangedForEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerHanaExchangedForEth represents a HanaExchangedForEth event raised by the HanaTokenConsumer contract.
type HanaTokenConsumerHanaExchangedForEth struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHanaExchangedForEth is a free log retrieval operation binding the contract event 0xd94ef7f26aa4584e0bb757c8b45e0663569b29a3a8a0fc85ed49a8bb95ed306a.
//
// Solidity: event HanaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumer *HanaTokenConsumerFilterer) FilterHanaExchangedForEth(opts *bind.FilterOpts) (*HanaTokenConsumerHanaExchangedForEthIterator, error) {

	logs, sub, err := _HanaTokenConsumer.contract.FilterLogs(opts, "HanaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerHanaExchangedForEthIterator{contract: _HanaTokenConsumer.contract, event: "HanaExchangedForEth", logs: logs, sub: sub}, nil
}

// WatchHanaExchangedForEth is a free log subscription operation binding the contract event 0xd94ef7f26aa4584e0bb757c8b45e0663569b29a3a8a0fc85ed49a8bb95ed306a.
//
// Solidity: event HanaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumer *HanaTokenConsumerFilterer) WatchHanaExchangedForEth(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerHanaExchangedForEth) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumer.contract.WatchLogs(opts, "HanaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerHanaExchangedForEth)
				if err := _HanaTokenConsumer.contract.UnpackLog(event, "HanaExchangedForEth", log); err != nil {
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

// ParseHanaExchangedForEth is a log parse operation binding the contract event 0xd94ef7f26aa4584e0bb757c8b45e0663569b29a3a8a0fc85ed49a8bb95ed306a.
//
// Solidity: event HanaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumer *HanaTokenConsumerFilterer) ParseHanaExchangedForEth(log types.Log) (*HanaTokenConsumerHanaExchangedForEth, error) {
	event := new(HanaTokenConsumerHanaExchangedForEth)
	if err := _HanaTokenConsumer.contract.UnpackLog(event, "HanaExchangedForEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaTokenConsumerHanaExchangedForTokenIterator is returned from FilterHanaExchangedForToken and is used to iterate over the raw logs and unpacked data for HanaExchangedForToken events raised by the HanaTokenConsumer contract.
type HanaTokenConsumerHanaExchangedForTokenIterator struct {
	Event *HanaTokenConsumerHanaExchangedForToken // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerHanaExchangedForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerHanaExchangedForToken)
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
		it.Event = new(HanaTokenConsumerHanaExchangedForToken)
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
func (it *HanaTokenConsumerHanaExchangedForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerHanaExchangedForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerHanaExchangedForToken represents a HanaExchangedForToken event raised by the HanaTokenConsumer contract.
type HanaTokenConsumerHanaExchangedForToken struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHanaExchangedForToken is a free log retrieval operation binding the contract event 0x532789829182476cbd3cc6523b623cc2504e5a0b348dfdc5527baf726e6d55af.
//
// Solidity: event HanaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumer *HanaTokenConsumerFilterer) FilterHanaExchangedForToken(opts *bind.FilterOpts) (*HanaTokenConsumerHanaExchangedForTokenIterator, error) {

	logs, sub, err := _HanaTokenConsumer.contract.FilterLogs(opts, "HanaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerHanaExchangedForTokenIterator{contract: _HanaTokenConsumer.contract, event: "HanaExchangedForToken", logs: logs, sub: sub}, nil
}

// WatchHanaExchangedForToken is a free log subscription operation binding the contract event 0x532789829182476cbd3cc6523b623cc2504e5a0b348dfdc5527baf726e6d55af.
//
// Solidity: event HanaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumer *HanaTokenConsumerFilterer) WatchHanaExchangedForToken(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerHanaExchangedForToken) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumer.contract.WatchLogs(opts, "HanaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerHanaExchangedForToken)
				if err := _HanaTokenConsumer.contract.UnpackLog(event, "HanaExchangedForToken", log); err != nil {
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

// ParseHanaExchangedForToken is a log parse operation binding the contract event 0x532789829182476cbd3cc6523b623cc2504e5a0b348dfdc5527baf726e6d55af.
//
// Solidity: event HanaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumer *HanaTokenConsumerFilterer) ParseHanaExchangedForToken(log types.Log) (*HanaTokenConsumerHanaExchangedForToken, error) {
	event := new(HanaTokenConsumerHanaExchangedForToken)
	if err := _HanaTokenConsumer.contract.UnpackLog(event, "HanaExchangedForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaTokenConsumerTokenExchangedForHanaIterator is returned from FilterTokenExchangedForHana and is used to iterate over the raw logs and unpacked data for TokenExchangedForHana events raised by the HanaTokenConsumer contract.
type HanaTokenConsumerTokenExchangedForHanaIterator struct {
	Event *HanaTokenConsumerTokenExchangedForHana // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerTokenExchangedForHanaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerTokenExchangedForHana)
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
		it.Event = new(HanaTokenConsumerTokenExchangedForHana)
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
func (it *HanaTokenConsumerTokenExchangedForHanaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerTokenExchangedForHanaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerTokenExchangedForHana represents a TokenExchangedForHana event raised by the HanaTokenConsumer contract.
type HanaTokenConsumerTokenExchangedForHana struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangedForHana is a free log retrieval operation binding the contract event 0xc7f70cbba54984322735d319a149c7bccec0ab310897fbc85b00f59d4486ce60.
//
// Solidity: event TokenExchangedForHana(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumer *HanaTokenConsumerFilterer) FilterTokenExchangedForHana(opts *bind.FilterOpts) (*HanaTokenConsumerTokenExchangedForHanaIterator, error) {

	logs, sub, err := _HanaTokenConsumer.contract.FilterLogs(opts, "TokenExchangedForHana")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerTokenExchangedForHanaIterator{contract: _HanaTokenConsumer.contract, event: "TokenExchangedForHana", logs: logs, sub: sub}, nil
}

// WatchTokenExchangedForHana is a free log subscription operation binding the contract event 0xc7f70cbba54984322735d319a149c7bccec0ab310897fbc85b00f59d4486ce60.
//
// Solidity: event TokenExchangedForHana(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumer *HanaTokenConsumerFilterer) WatchTokenExchangedForHana(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerTokenExchangedForHana) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumer.contract.WatchLogs(opts, "TokenExchangedForHana")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerTokenExchangedForHana)
				if err := _HanaTokenConsumer.contract.UnpackLog(event, "TokenExchangedForHana", log); err != nil {
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

// ParseTokenExchangedForHana is a log parse operation binding the contract event 0xc7f70cbba54984322735d319a149c7bccec0ab310897fbc85b00f59d4486ce60.
//
// Solidity: event TokenExchangedForHana(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumer *HanaTokenConsumerFilterer) ParseTokenExchangedForHana(log types.Log) (*HanaTokenConsumerTokenExchangedForHana, error) {
	event := new(HanaTokenConsumerTokenExchangedForHana)
	if err := _HanaTokenConsumer.contract.UnpackLog(event, "TokenExchangedForHana", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
