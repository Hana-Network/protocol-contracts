// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package connectorhevm

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

// HanaConnectorHEVMMetaData contains all meta data concerning the HanaConnectorHEVM contract.
var HanaConnectorHEVMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_whana\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedHanaSent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyFungibleModule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWHANA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WHANATransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceTxOriginAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hanaValueAndGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hanaParams\",\"type\":\"bytes\"}],\"name\":\"HanaSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"whana_\",\"type\":\"address\"}],\"name\":\"SetWHANA\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"hanaValueAndGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"hanaParams\",\"type\":\"bytes\"}],\"internalType\":\"structHanaInterfaces.SendInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whana_\",\"type\":\"address\"}],\"name\":\"setWhanaAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whana\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610a52380380610a528339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610108565b600081519050610087816100f1565b92915050565b6000602082840312156100a3576100a26100ec565b5b60006100b184828501610078565b91505092915050565b60006100c5826100cc565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6100fa816100ba565b811461010557600080fd5b50565b61093b806101176000396000f3fe6080604052600436106100435760003560e01c80633ce4a5bc146100d4578063a4590ea1146100ff578063ca9457a514610128578063ec02690114610153576100cf565b366100cf5760008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146100cd576040517fa5a74d8c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b3480156100e057600080fd5b506100e961017c565b6040516100f691906106e6565b60405180910390f35b34801561010b57600080fd5b50610126600480360381019061012191906105c0565b610194565b005b34801561013457600080fd5b5061013d610287565b60405161014a91906106e6565b60405180910390f35b34801561015f57600080fd5b5061017a6004803603810190610175919061061a565b6102ab565b005b73735b14bb79463307aacbed86daf3322b1e6226ab81565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461020d576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507ffd250d2fc6a8268f30b40917cdfa5b10a3e164e384e03bf4a90192e81ffa27d08160405161027c91906106e6565b60405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd333084608001356040518463ffffffff1660e01b815260040161030c93929190610701565b602060405180830381600087803b15801561032657600080fd5b505af115801561033a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061035e91906105ed565b610394576040517ff8e9c1a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d82608001356040518263ffffffff1660e01b81526004016103f191906107b4565b600060405180830381600087803b15801561040b57600080fd5b505af115801561041f573d6000803e3d6000fd5b50505050600073735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff168260800135604051610461906106d1565b60006040518083038185875af1925050503d806000811461049e576040519150601f19603f3d011682016040523d82523d6000602084013e6104a3565b606091505b50509050806104de576040517f9a543f2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600001353373ffffffffffffffffffffffffffffffffffffffff167f7a0dd478962bca8db7a72b684fbfd3404be69978e5f8f1c8eab2cfe4724ebaee3285806020019061052c91906107cf565b8760800135886040013589806060019061054691906107cf565b8b8060a0019061055691906107cf565b60405161056b99989796959493929190610738565b60405180910390a35050565b600081359050610586816108d7565b92915050565b60008151905061059b816108ee565b92915050565b600060c082840312156105b7576105b66108aa565b5b81905092915050565b6000602082840312156105d6576105d56108be565b5b60006105e484828501610577565b91505092915050565b600060208284031215610603576106026108be565b5b60006106118482850161058c565b91505092915050565b6000602082840312156106305761062f6108be565b5b600082013567ffffffffffffffff81111561064e5761064d6108b9565b5b61065a848285016105a1565b91505092915050565b61066c8161084e565b82525050565b600061067e8385610832565b935061068b838584610896565b610694836108c3565b840190509392505050565b60006106ac600083610843565b91506106b7826108d4565b600082019050919050565b6106cb8161088c565b82525050565b60006106dc8261069f565b9150819050919050565b60006020820190506106fb6000830184610663565b92915050565b60006060820190506107166000830186610663565b6107236020830185610663565b61073060408301846106c2565b949350505050565b600060c08201905061074d600083018c610663565b8181036020830152610760818a8c610672565b905061076f60408301896106c2565b61077c60608301886106c2565b818103608083015261078f818688610672565b905081810360a08301526107a4818486610672565b90509a9950505050505050505050565b60006020820190506107c960008301846106c2565b92915050565b600080833560016020038436030381126107ec576107eb6108af565b5b80840192508235915067ffffffffffffffff82111561080e5761080d6108a5565b5b60208301925060018202360383131561082a576108296108b4565b5b509250929050565b600082825260208201905092915050565b600081905092915050565b60006108598261086c565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b50565b6108e08161084e565b81146108eb57600080fd5b50565b6108f781610860565b811461090257600080fd5b5056fea2646970667358221220b0dda7375a0943a921214b5105c1e1054d898d600d9848abc9eb488ec80a0fd264736f6c63430008070033",
}

// HanaConnectorHEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaConnectorHEVMMetaData.ABI instead.
var HanaConnectorHEVMABI = HanaConnectorHEVMMetaData.ABI

// HanaConnectorHEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HanaConnectorHEVMMetaData.Bin instead.
var HanaConnectorHEVMBin = HanaConnectorHEVMMetaData.Bin

// DeployHanaConnectorHEVM deploys a new Ethereum contract, binding an instance of HanaConnectorHEVM to it.
func DeployHanaConnectorHEVM(auth *bind.TransactOpts, backend bind.ContractBackend, _whana common.Address) (common.Address, *types.Transaction, *HanaConnectorHEVM, error) {
	parsed, err := HanaConnectorHEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HanaConnectorHEVMBin), backend, _whana)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HanaConnectorHEVM{HanaConnectorHEVMCaller: HanaConnectorHEVMCaller{contract: contract}, HanaConnectorHEVMTransactor: HanaConnectorHEVMTransactor{contract: contract}, HanaConnectorHEVMFilterer: HanaConnectorHEVMFilterer{contract: contract}}, nil
}

// HanaConnectorHEVM is an auto generated Go binding around an Ethereum contract.
type HanaConnectorHEVM struct {
	HanaConnectorHEVMCaller     // Read-only binding to the contract
	HanaConnectorHEVMTransactor // Write-only binding to the contract
	HanaConnectorHEVMFilterer   // Log filterer for contract events
}

// HanaConnectorHEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaConnectorHEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorHEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaConnectorHEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorHEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaConnectorHEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorHEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaConnectorHEVMSession struct {
	Contract     *HanaConnectorHEVM // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HanaConnectorHEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaConnectorHEVMCallerSession struct {
	Contract *HanaConnectorHEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// HanaConnectorHEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaConnectorHEVMTransactorSession struct {
	Contract     *HanaConnectorHEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// HanaConnectorHEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaConnectorHEVMRaw struct {
	Contract *HanaConnectorHEVM // Generic contract binding to access the raw methods on
}

// HanaConnectorHEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaConnectorHEVMCallerRaw struct {
	Contract *HanaConnectorHEVMCaller // Generic read-only contract binding to access the raw methods on
}

// HanaConnectorHEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaConnectorHEVMTransactorRaw struct {
	Contract *HanaConnectorHEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaConnectorHEVM creates a new instance of HanaConnectorHEVM, bound to a specific deployed contract.
func NewHanaConnectorHEVM(address common.Address, backend bind.ContractBackend) (*HanaConnectorHEVM, error) {
	contract, err := bindHanaConnectorHEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorHEVM{HanaConnectorHEVMCaller: HanaConnectorHEVMCaller{contract: contract}, HanaConnectorHEVMTransactor: HanaConnectorHEVMTransactor{contract: contract}, HanaConnectorHEVMFilterer: HanaConnectorHEVMFilterer{contract: contract}}, nil
}

// NewHanaConnectorHEVMCaller creates a new read-only instance of HanaConnectorHEVM, bound to a specific deployed contract.
func NewHanaConnectorHEVMCaller(address common.Address, caller bind.ContractCaller) (*HanaConnectorHEVMCaller, error) {
	contract, err := bindHanaConnectorHEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorHEVMCaller{contract: contract}, nil
}

// NewHanaConnectorHEVMTransactor creates a new write-only instance of HanaConnectorHEVM, bound to a specific deployed contract.
func NewHanaConnectorHEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaConnectorHEVMTransactor, error) {
	contract, err := bindHanaConnectorHEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorHEVMTransactor{contract: contract}, nil
}

// NewHanaConnectorHEVMFilterer creates a new log filterer instance of HanaConnectorHEVM, bound to a specific deployed contract.
func NewHanaConnectorHEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaConnectorHEVMFilterer, error) {
	contract, err := bindHanaConnectorHEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorHEVMFilterer{contract: contract}, nil
}

// bindHanaConnectorHEVM binds a generic wrapper to an already deployed contract.
func bindHanaConnectorHEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaConnectorHEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaConnectorHEVM *HanaConnectorHEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaConnectorHEVM.Contract.HanaConnectorHEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaConnectorHEVM *HanaConnectorHEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorHEVM.Contract.HanaConnectorHEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaConnectorHEVM *HanaConnectorHEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaConnectorHEVM.Contract.HanaConnectorHEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaConnectorHEVM *HanaConnectorHEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaConnectorHEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaConnectorHEVM *HanaConnectorHEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorHEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaConnectorHEVM *HanaConnectorHEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaConnectorHEVM.Contract.contract.Transact(opts, method, params...)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_HanaConnectorHEVM *HanaConnectorHEVMCaller) FUNGIBLEMODULEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaConnectorHEVM.contract.Call(opts, &out, "FUNGIBLE_MODULE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_HanaConnectorHEVM *HanaConnectorHEVMSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _HanaConnectorHEVM.Contract.FUNGIBLEMODULEADDRESS(&_HanaConnectorHEVM.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_HanaConnectorHEVM *HanaConnectorHEVMCallerSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _HanaConnectorHEVM.Contract.FUNGIBLEMODULEADDRESS(&_HanaConnectorHEVM.CallOpts)
}

// Whana is a free data retrieval call binding the contract method 0xca9457a5.
//
// Solidity: function whana() view returns(address)
func (_HanaConnectorHEVM *HanaConnectorHEVMCaller) Whana(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaConnectorHEVM.contract.Call(opts, &out, "whana")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Whana is a free data retrieval call binding the contract method 0xca9457a5.
//
// Solidity: function whana() view returns(address)
func (_HanaConnectorHEVM *HanaConnectorHEVMSession) Whana() (common.Address, error) {
	return _HanaConnectorHEVM.Contract.Whana(&_HanaConnectorHEVM.CallOpts)
}

// Whana is a free data retrieval call binding the contract method 0xca9457a5.
//
// Solidity: function whana() view returns(address)
func (_HanaConnectorHEVM *HanaConnectorHEVMCallerSession) Whana() (common.Address, error) {
	return _HanaConnectorHEVM.Contract.Whana(&_HanaConnectorHEVM.CallOpts)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMTransactor) Send(opts *bind.TransactOpts, input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnectorHEVM.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMSession) Send(input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnectorHEVM.Contract.Send(&_HanaConnectorHEVM.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMTransactorSession) Send(input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnectorHEVM.Contract.Send(&_HanaConnectorHEVM.TransactOpts, input)
}

// SetWhanaAddress is a paid mutator transaction binding the contract method 0xa4590ea1.
//
// Solidity: function setWhanaAddress(address whana_) returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMTransactor) SetWhanaAddress(opts *bind.TransactOpts, whana_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorHEVM.contract.Transact(opts, "setWhanaAddress", whana_)
}

// SetWhanaAddress is a paid mutator transaction binding the contract method 0xa4590ea1.
//
// Solidity: function setWhanaAddress(address whana_) returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMSession) SetWhanaAddress(whana_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorHEVM.Contract.SetWhanaAddress(&_HanaConnectorHEVM.TransactOpts, whana_)
}

// SetWhanaAddress is a paid mutator transaction binding the contract method 0xa4590ea1.
//
// Solidity: function setWhanaAddress(address whana_) returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMTransactorSession) SetWhanaAddress(whana_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorHEVM.Contract.SetWhanaAddress(&_HanaConnectorHEVM.TransactOpts, whana_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorHEVM.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMSession) Receive() (*types.Transaction, error) {
	return _HanaConnectorHEVM.Contract.Receive(&_HanaConnectorHEVM.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMTransactorSession) Receive() (*types.Transaction, error) {
	return _HanaConnectorHEVM.Contract.Receive(&_HanaConnectorHEVM.TransactOpts)
}

// HanaConnectorHEVMHanaSentIterator is returned from FilterHanaSent and is used to iterate over the raw logs and unpacked data for HanaSent events raised by the HanaConnectorHEVM contract.
type HanaConnectorHEVMHanaSentIterator struct {
	Event *HanaConnectorHEVMHanaSent // Event containing the contract specifics and raw log

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
func (it *HanaConnectorHEVMHanaSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorHEVMHanaSent)
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
		it.Event = new(HanaConnectorHEVMHanaSent)
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
func (it *HanaConnectorHEVMHanaSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorHEVMHanaSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorHEVMHanaSent represents a HanaSent event raised by the HanaConnectorHEVM contract.
type HanaConnectorHEVMHanaSent struct {
	SourceTxOriginAddress common.Address
	HanaTxSenderAddress   common.Address
	DestinationChainId    *big.Int
	DestinationAddress    []byte
	HanaValueAndGas       *big.Int
	DestinationGasLimit   *big.Int
	Message               []byte
	HanaParams            []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterHanaSent is a free log retrieval operation binding the contract event 0x7a0dd478962bca8db7a72b684fbfd3404be69978e5f8f1c8eab2cfe4724ebaee.
//
// Solidity: event HanaSent(address sourceTxOriginAddress, address indexed hanaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 hanaValueAndGas, uint256 destinationGasLimit, bytes message, bytes hanaParams)
func (_HanaConnectorHEVM *HanaConnectorHEVMFilterer) FilterHanaSent(opts *bind.FilterOpts, hanaTxSenderAddress []common.Address, destinationChainId []*big.Int) (*HanaConnectorHEVMHanaSentIterator, error) {

	var hanaTxSenderAddressRule []interface{}
	for _, hanaTxSenderAddressItem := range hanaTxSenderAddress {
		hanaTxSenderAddressRule = append(hanaTxSenderAddressRule, hanaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _HanaConnectorHEVM.contract.FilterLogs(opts, "HanaSent", hanaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorHEVMHanaSentIterator{contract: _HanaConnectorHEVM.contract, event: "HanaSent", logs: logs, sub: sub}, nil
}

// WatchHanaSent is a free log subscription operation binding the contract event 0x7a0dd478962bca8db7a72b684fbfd3404be69978e5f8f1c8eab2cfe4724ebaee.
//
// Solidity: event HanaSent(address sourceTxOriginAddress, address indexed hanaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 hanaValueAndGas, uint256 destinationGasLimit, bytes message, bytes hanaParams)
func (_HanaConnectorHEVM *HanaConnectorHEVMFilterer) WatchHanaSent(opts *bind.WatchOpts, sink chan<- *HanaConnectorHEVMHanaSent, hanaTxSenderAddress []common.Address, destinationChainId []*big.Int) (event.Subscription, error) {

	var hanaTxSenderAddressRule []interface{}
	for _, hanaTxSenderAddressItem := range hanaTxSenderAddress {
		hanaTxSenderAddressRule = append(hanaTxSenderAddressRule, hanaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _HanaConnectorHEVM.contract.WatchLogs(opts, "HanaSent", hanaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorHEVMHanaSent)
				if err := _HanaConnectorHEVM.contract.UnpackLog(event, "HanaSent", log); err != nil {
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

// ParseHanaSent is a log parse operation binding the contract event 0x7a0dd478962bca8db7a72b684fbfd3404be69978e5f8f1c8eab2cfe4724ebaee.
//
// Solidity: event HanaSent(address sourceTxOriginAddress, address indexed hanaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 hanaValueAndGas, uint256 destinationGasLimit, bytes message, bytes hanaParams)
func (_HanaConnectorHEVM *HanaConnectorHEVMFilterer) ParseHanaSent(log types.Log) (*HanaConnectorHEVMHanaSent, error) {
	event := new(HanaConnectorHEVMHanaSent)
	if err := _HanaConnectorHEVM.contract.UnpackLog(event, "HanaSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorHEVMSetWHANAIterator is returned from FilterSetWHANA and is used to iterate over the raw logs and unpacked data for SetWHANA events raised by the HanaConnectorHEVM contract.
type HanaConnectorHEVMSetWHANAIterator struct {
	Event *HanaConnectorHEVMSetWHANA // Event containing the contract specifics and raw log

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
func (it *HanaConnectorHEVMSetWHANAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorHEVMSetWHANA)
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
		it.Event = new(HanaConnectorHEVMSetWHANA)
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
func (it *HanaConnectorHEVMSetWHANAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorHEVMSetWHANAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorHEVMSetWHANA represents a SetWHANA event raised by the HanaConnectorHEVM contract.
type HanaConnectorHEVMSetWHANA struct {
	Whana common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetWHANA is a free log retrieval operation binding the contract event 0xfd250d2fc6a8268f30b40917cdfa5b10a3e164e384e03bf4a90192e81ffa27d0.
//
// Solidity: event SetWHANA(address whana_)
func (_HanaConnectorHEVM *HanaConnectorHEVMFilterer) FilterSetWHANA(opts *bind.FilterOpts) (*HanaConnectorHEVMSetWHANAIterator, error) {

	logs, sub, err := _HanaConnectorHEVM.contract.FilterLogs(opts, "SetWHANA")
	if err != nil {
		return nil, err
	}
	return &HanaConnectorHEVMSetWHANAIterator{contract: _HanaConnectorHEVM.contract, event: "SetWHANA", logs: logs, sub: sub}, nil
}

// WatchSetWHANA is a free log subscription operation binding the contract event 0xfd250d2fc6a8268f30b40917cdfa5b10a3e164e384e03bf4a90192e81ffa27d0.
//
// Solidity: event SetWHANA(address whana_)
func (_HanaConnectorHEVM *HanaConnectorHEVMFilterer) WatchSetWHANA(opts *bind.WatchOpts, sink chan<- *HanaConnectorHEVMSetWHANA) (event.Subscription, error) {

	logs, sub, err := _HanaConnectorHEVM.contract.WatchLogs(opts, "SetWHANA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorHEVMSetWHANA)
				if err := _HanaConnectorHEVM.contract.UnpackLog(event, "SetWHANA", log); err != nil {
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

// ParseSetWHANA is a log parse operation binding the contract event 0xfd250d2fc6a8268f30b40917cdfa5b10a3e164e384e03bf4a90192e81ffa27d0.
//
// Solidity: event SetWHANA(address whana_)
func (_HanaConnectorHEVM *HanaConnectorHEVMFilterer) ParseSetWHANA(log types.Log) (*HanaConnectorHEVMSetWHANA, error) {
	event := new(HanaConnectorHEVMSetWHANA)
	if err := _HanaConnectorHEVM.contract.UnpackLog(event, "SetWHANA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
