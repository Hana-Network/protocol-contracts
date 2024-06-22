// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hananonethinterface

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

// HanaNonEthInterfaceMetaData contains all meta data concerning the HanaNonEthInterface contract.
var HanaNonEthInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mintee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HanaNonEthInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaNonEthInterfaceMetaData.ABI instead.
var HanaNonEthInterfaceABI = HanaNonEthInterfaceMetaData.ABI

// HanaNonEthInterface is an auto generated Go binding around an Ethereum contract.
type HanaNonEthInterface struct {
	HanaNonEthInterfaceCaller     // Read-only binding to the contract
	HanaNonEthInterfaceTransactor // Write-only binding to the contract
	HanaNonEthInterfaceFilterer   // Log filterer for contract events
}

// HanaNonEthInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaNonEthInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaNonEthInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaNonEthInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaNonEthInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaNonEthInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaNonEthInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaNonEthInterfaceSession struct {
	Contract     *HanaNonEthInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HanaNonEthInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaNonEthInterfaceCallerSession struct {
	Contract *HanaNonEthInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// HanaNonEthInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaNonEthInterfaceTransactorSession struct {
	Contract     *HanaNonEthInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// HanaNonEthInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaNonEthInterfaceRaw struct {
	Contract *HanaNonEthInterface // Generic contract binding to access the raw methods on
}

// HanaNonEthInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaNonEthInterfaceCallerRaw struct {
	Contract *HanaNonEthInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// HanaNonEthInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaNonEthInterfaceTransactorRaw struct {
	Contract *HanaNonEthInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaNonEthInterface creates a new instance of HanaNonEthInterface, bound to a specific deployed contract.
func NewHanaNonEthInterface(address common.Address, backend bind.ContractBackend) (*HanaNonEthInterface, error) {
	contract, err := bindHanaNonEthInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaNonEthInterface{HanaNonEthInterfaceCaller: HanaNonEthInterfaceCaller{contract: contract}, HanaNonEthInterfaceTransactor: HanaNonEthInterfaceTransactor{contract: contract}, HanaNonEthInterfaceFilterer: HanaNonEthInterfaceFilterer{contract: contract}}, nil
}

// NewHanaNonEthInterfaceCaller creates a new read-only instance of HanaNonEthInterface, bound to a specific deployed contract.
func NewHanaNonEthInterfaceCaller(address common.Address, caller bind.ContractCaller) (*HanaNonEthInterfaceCaller, error) {
	contract, err := bindHanaNonEthInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaNonEthInterfaceCaller{contract: contract}, nil
}

// NewHanaNonEthInterfaceTransactor creates a new write-only instance of HanaNonEthInterface, bound to a specific deployed contract.
func NewHanaNonEthInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaNonEthInterfaceTransactor, error) {
	contract, err := bindHanaNonEthInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaNonEthInterfaceTransactor{contract: contract}, nil
}

// NewHanaNonEthInterfaceFilterer creates a new log filterer instance of HanaNonEthInterface, bound to a specific deployed contract.
func NewHanaNonEthInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaNonEthInterfaceFilterer, error) {
	contract, err := bindHanaNonEthInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaNonEthInterfaceFilterer{contract: contract}, nil
}

// bindHanaNonEthInterface binds a generic wrapper to an already deployed contract.
func bindHanaNonEthInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaNonEthInterfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaNonEthInterface *HanaNonEthInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaNonEthInterface.Contract.HanaNonEthInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaNonEthInterface *HanaNonEthInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaNonEthInterface.Contract.HanaNonEthInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaNonEthInterface *HanaNonEthInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaNonEthInterface.Contract.HanaNonEthInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaNonEthInterface *HanaNonEthInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaNonEthInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaNonEthInterface *HanaNonEthInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaNonEthInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaNonEthInterface *HanaNonEthInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaNonEthInterface.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HanaNonEthInterface *HanaNonEthInterfaceCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HanaNonEthInterface.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HanaNonEthInterface *HanaNonEthInterfaceSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _HanaNonEthInterface.Contract.Allowance(&_HanaNonEthInterface.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HanaNonEthInterface *HanaNonEthInterfaceCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _HanaNonEthInterface.Contract.Allowance(&_HanaNonEthInterface.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HanaNonEthInterface *HanaNonEthInterfaceCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HanaNonEthInterface.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HanaNonEthInterface *HanaNonEthInterfaceSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HanaNonEthInterface.Contract.BalanceOf(&_HanaNonEthInterface.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HanaNonEthInterface *HanaNonEthInterfaceCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HanaNonEthInterface.Contract.BalanceOf(&_HanaNonEthInterface.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HanaNonEthInterface *HanaNonEthInterfaceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HanaNonEthInterface.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HanaNonEthInterface *HanaNonEthInterfaceSession) TotalSupply() (*big.Int, error) {
	return _HanaNonEthInterface.Contract.TotalSupply(&_HanaNonEthInterface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HanaNonEthInterface *HanaNonEthInterfaceCallerSession) TotalSupply() (*big.Int, error) {
	return _HanaNonEthInterface.Contract.TotalSupply(&_HanaNonEthInterface.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_HanaNonEthInterface *HanaNonEthInterfaceTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HanaNonEthInterface.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_HanaNonEthInterface *HanaNonEthInterfaceSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HanaNonEthInterface.Contract.Approve(&_HanaNonEthInterface.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_HanaNonEthInterface *HanaNonEthInterfaceTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HanaNonEthInterface.Contract.Approve(&_HanaNonEthInterface.TransactOpts, spender, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_HanaNonEthInterface *HanaNonEthInterfaceTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HanaNonEthInterface.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_HanaNonEthInterface *HanaNonEthInterfaceSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HanaNonEthInterface.Contract.BurnFrom(&_HanaNonEthInterface.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_HanaNonEthInterface *HanaNonEthInterfaceTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HanaNonEthInterface.Contract.BurnFrom(&_HanaNonEthInterface.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_HanaNonEthInterface *HanaNonEthInterfaceTransactor) Mint(opts *bind.TransactOpts, mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaNonEthInterface.contract.Transact(opts, "mint", mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_HanaNonEthInterface *HanaNonEthInterfaceSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaNonEthInterface.Contract.Mint(&_HanaNonEthInterface.TransactOpts, mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_HanaNonEthInterface *HanaNonEthInterfaceTransactorSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaNonEthInterface.Contract.Mint(&_HanaNonEthInterface.TransactOpts, mintee, value, internalSendHash)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_HanaNonEthInterface *HanaNonEthInterfaceTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HanaNonEthInterface.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_HanaNonEthInterface *HanaNonEthInterfaceSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HanaNonEthInterface.Contract.Transfer(&_HanaNonEthInterface.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_HanaNonEthInterface *HanaNonEthInterfaceTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HanaNonEthInterface.Contract.Transfer(&_HanaNonEthInterface.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_HanaNonEthInterface *HanaNonEthInterfaceTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HanaNonEthInterface.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_HanaNonEthInterface *HanaNonEthInterfaceSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HanaNonEthInterface.Contract.TransferFrom(&_HanaNonEthInterface.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_HanaNonEthInterface *HanaNonEthInterfaceTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HanaNonEthInterface.Contract.TransferFrom(&_HanaNonEthInterface.TransactOpts, from, to, amount)
}

// HanaNonEthInterfaceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the HanaNonEthInterface contract.
type HanaNonEthInterfaceApprovalIterator struct {
	Event *HanaNonEthInterfaceApproval // Event containing the contract specifics and raw log

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
func (it *HanaNonEthInterfaceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaNonEthInterfaceApproval)
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
		it.Event = new(HanaNonEthInterfaceApproval)
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
func (it *HanaNonEthInterfaceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaNonEthInterfaceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaNonEthInterfaceApproval represents a Approval event raised by the HanaNonEthInterface contract.
type HanaNonEthInterfaceApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HanaNonEthInterface *HanaNonEthInterfaceFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*HanaNonEthInterfaceApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _HanaNonEthInterface.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &HanaNonEthInterfaceApprovalIterator{contract: _HanaNonEthInterface.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HanaNonEthInterface *HanaNonEthInterfaceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HanaNonEthInterfaceApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _HanaNonEthInterface.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaNonEthInterfaceApproval)
				if err := _HanaNonEthInterface.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HanaNonEthInterface *HanaNonEthInterfaceFilterer) ParseApproval(log types.Log) (*HanaNonEthInterfaceApproval, error) {
	event := new(HanaNonEthInterfaceApproval)
	if err := _HanaNonEthInterface.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaNonEthInterfaceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the HanaNonEthInterface contract.
type HanaNonEthInterfaceTransferIterator struct {
	Event *HanaNonEthInterfaceTransfer // Event containing the contract specifics and raw log

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
func (it *HanaNonEthInterfaceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaNonEthInterfaceTransfer)
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
		it.Event = new(HanaNonEthInterfaceTransfer)
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
func (it *HanaNonEthInterfaceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaNonEthInterfaceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaNonEthInterfaceTransfer represents a Transfer event raised by the HanaNonEthInterface contract.
type HanaNonEthInterfaceTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HanaNonEthInterface *HanaNonEthInterfaceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HanaNonEthInterfaceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HanaNonEthInterface.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HanaNonEthInterfaceTransferIterator{contract: _HanaNonEthInterface.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HanaNonEthInterface *HanaNonEthInterfaceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HanaNonEthInterfaceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HanaNonEthInterface.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaNonEthInterfaceTransfer)
				if err := _HanaNonEthInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HanaNonEthInterface *HanaNonEthInterfaceFilterer) ParseTransfer(log types.Log) (*HanaNonEthInterfaceTransfer, error) {
	event := new(HanaNonEthInterfaceTransfer)
	if err := _HanaNonEthInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
