// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanainteractormock

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

// HanaInteractorMockMetaData contains all meta data concerning the HanaInteractorMock contract.
var HanaInteractorMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hanaConnectorAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidHanaMessageCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidHanaRevertCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connector\",\"outputs\":[{\"internalType\":\"contractHanaConnector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"interactorsByChainId\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"hanaTxSenderAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hanaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structHanaInterfaces.HanaMessage\",\"name\":\"hanaMessage\",\"type\":\"tuple\"}],\"name\":\"onHanaMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingHanaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structHanaInterfaces.HanaRevert\",\"name\":\"hanaRevert\",\"type\":\"tuple\"}],\"name\":\"onHanaRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"contractAddress\",\"type\":\"bytes\"}],\"name\":\"setInteractorByChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200123c3803806200123c833981810160405281019062000037919062000228565b80620000586200004c6200010760201b60201c565b6200010f60201b60201c565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620000c0576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b46608081815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250505050620002ad565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556200014a816200014d60201b620005bd1760201c565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050620002228162000293565b92915050565b6000602082840312156200024157620002406200028e565b5b6000620002518482850162000211565b91505092915050565b600062000267826200026e565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200029e816200025a565b8114620002aa57600080fd5b50565b60805160a05160601c610f5f620002dd600039600081816104050152610683015260006102d30152610f5f6000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806383f3084f1161006657806383f3084f1461011f5780638da5cb5b1461013d578063b204be931461015b578063e30c397814610177578063f2fde38b146101955761009e565b80632618143f146100a35780634204cf9b146100d35780634fd3f7d7146100ef578063715018a61461010b57806379ba509714610115575b600080fd5b6100bd60048036038101906100b891906109ea565b6101b1565b6040516100ca9190610c03565b60405180910390f35b6100ed60048036038101906100e891906109a1565b610251565b005b61010960048036038101906101049190610a17565b610332565b005b610113610362565b005b61011d610376565b005b610127610403565b6040516101349190610c25565b60405180910390f35b610145610427565b6040516101529190610be8565b60405180910390f35b61017560048036038101906101709190610958565b610450565b005b61017f6104e6565b60405161018c9190610be8565b60405180910390f35b6101af60048036038101906101aa919061092b565b610510565b005b600260205280600052604060002060009150905080546101d090610de4565b80601f01602080910402602001604051908101604052809291908181526020018280546101fc90610de4565b80156102495780601f1061021e57610100808354040283529160200191610249565b820191906000526020600020905b81548152906001019060200180831161022c57829003601f168201915b505050505081565b8061025a610681565b3073ffffffffffffffffffffffffffffffffffffffff16816000016020810190610284919061092b565b73ffffffffffffffffffffffffffffffffffffffff16146102d1576040517fb53f9b7500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000081602001351461032e576040517fb53f9b7500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b61033a610713565b818160026000868152602001908152602001600020919061035c9291906107ca565b50505050565b61036a610713565b6103746000610791565b565b60006103806107c2565b90508073ffffffffffffffffffffffffffffffffffffffff166103a16104e6565b73ffffffffffffffffffffffffffffffffffffffff16146103f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ee90610c40565b60405180910390fd5b61040081610791565b50565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b80610459610681565b600260008260200135815260200190815260200160002060405161047d9190610bd1565b60405180910390208180600001906104959190610c80565b6040516104a3929190610bb8565b6040518091039020146104e2576040517f39b0327e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610518610713565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16610578610427565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461071157336040517fcbd9d2e00000000000000000000000000000000000000000000000000000000081526004016107089190610be8565b60405180910390fd5b565b61071b6107c2565b73ffffffffffffffffffffffffffffffffffffffff16610739610427565b73ffffffffffffffffffffffffffffffffffffffff161461078f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078690610c60565b60405180910390fd5b565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556107bf816105bd565b50565b600033905090565b8280546107d690610de4565b90600052602060002090601f0160209004810192826107f8576000855561083f565b82601f1061081157803560ff191683800117855561083f565b8280016001018555821561083f579182015b8281111561083e578235825591602001919060010190610823565b5b50905061084c9190610850565b5090565b5b80821115610869576000816000905550600101610851565b5090565b60008135905061087c81610efb565b92915050565b60008083601f84011261089857610897610e4a565b5b8235905067ffffffffffffffff8111156108b5576108b4610e45565b5b6020830191508360018202830111156108d1576108d0610e5e565b5b9250929050565b600060a082840312156108ee576108ed610e54565b5b81905092915050565b600060c0828403121561090d5761090c610e54565b5b81905092915050565b60008135905061092581610f12565b92915050565b60006020828403121561094157610940610e6d565b5b600061094f8482850161086d565b91505092915050565b60006020828403121561096e5761096d610e6d565b5b600082013567ffffffffffffffff81111561098c5761098b610e68565b5b610998848285016108d8565b91505092915050565b6000602082840312156109b7576109b6610e6d565b5b600082013567ffffffffffffffff8111156109d5576109d4610e68565b5b6109e1848285016108f7565b91505092915050565b600060208284031215610a00576109ff610e6d565b5b6000610a0e84828501610916565b91505092915050565b600080600060408486031215610a3057610a2f610e6d565b5b6000610a3e86828701610916565b935050602084013567ffffffffffffffff811115610a5f57610a5e610e68565b5b610a6b86828701610882565b92509250509250925092565b610a8081610d30565b82525050565b6000610a928385610d14565b9350610a9f838584610da2565b82840190509392505050565b6000610ab682610cf8565b610ac08185610d03565b9350610ad0818560208601610db1565b610ad981610e72565b840191505092915050565b60008154610af181610de4565b610afb8186610d14565b94506001821660008114610b165760018114610b2757610b5a565b60ff19831686528186019350610b5a565b610b3085610ce3565b60005b83811015610b5257815481890152600182019150602081019050610b33565b838801955050505b50505092915050565b610b6c81610d6c565b82525050565b6000610b7f602983610d1f565b9150610b8a82610e83565b604082019050919050565b6000610ba2602083610d1f565b9150610bad82610ed2565b602082019050919050565b6000610bc5828486610a86565b91508190509392505050565b6000610bdd8284610ae4565b915081905092915050565b6000602082019050610bfd6000830184610a77565b92915050565b60006020820190508181036000830152610c1d8184610aab565b905092915050565b6000602082019050610c3a6000830184610b63565b92915050565b60006020820190508181036000830152610c5981610b72565b9050919050565b60006020820190508181036000830152610c7981610b95565b9050919050565b60008083356001602003843603038112610c9d57610c9c610e59565b5b80840192508235915067ffffffffffffffff821115610cbf57610cbe610e4f565b5b602083019250600182023603831315610cdb57610cda610e63565b5b509250929050565b60008190508160005260206000209050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610d3b82610d42565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610d7782610d7e565b9050919050565b6000610d8982610d90565b9050919050565b6000610d9b82610d42565b9050919050565b82818337600083830152505050565b60005b83811015610dcf578082015181840152602081019050610db4565b83811115610dde576000848401525b50505050565b60006002820490506001821680610dfc57607f821691505b60208210811415610e1057610e0f610e16565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c6532537465703a2063616c6c6572206973206e6f74207468652060008201527f6e6577206f776e65720000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b610f0481610d30565b8114610f0f57600080fd5b50565b610f1b81610d62565b8114610f2657600080fd5b5056fea264697066735822122020f14fa0348e19c5df9fd10d59e79fd804b9b46f752f1fe0627f57f16413203064736f6c63430008070033",
}

// HanaInteractorMockABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaInteractorMockMetaData.ABI instead.
var HanaInteractorMockABI = HanaInteractorMockMetaData.ABI

// HanaInteractorMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HanaInteractorMockMetaData.Bin instead.
var HanaInteractorMockBin = HanaInteractorMockMetaData.Bin

// DeployHanaInteractorMock deploys a new Ethereum contract, binding an instance of HanaInteractorMock to it.
func DeployHanaInteractorMock(auth *bind.TransactOpts, backend bind.ContractBackend, hanaConnectorAddress common.Address) (common.Address, *types.Transaction, *HanaInteractorMock, error) {
	parsed, err := HanaInteractorMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HanaInteractorMockBin), backend, hanaConnectorAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HanaInteractorMock{HanaInteractorMockCaller: HanaInteractorMockCaller{contract: contract}, HanaInteractorMockTransactor: HanaInteractorMockTransactor{contract: contract}, HanaInteractorMockFilterer: HanaInteractorMockFilterer{contract: contract}}, nil
}

// HanaInteractorMock is an auto generated Go binding around an Ethereum contract.
type HanaInteractorMock struct {
	HanaInteractorMockCaller     // Read-only binding to the contract
	HanaInteractorMockTransactor // Write-only binding to the contract
	HanaInteractorMockFilterer   // Log filterer for contract events
}

// HanaInteractorMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaInteractorMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaInteractorMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaInteractorMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaInteractorMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaInteractorMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaInteractorMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaInteractorMockSession struct {
	Contract     *HanaInteractorMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HanaInteractorMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaInteractorMockCallerSession struct {
	Contract *HanaInteractorMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// HanaInteractorMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaInteractorMockTransactorSession struct {
	Contract     *HanaInteractorMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// HanaInteractorMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaInteractorMockRaw struct {
	Contract *HanaInteractorMock // Generic contract binding to access the raw methods on
}

// HanaInteractorMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaInteractorMockCallerRaw struct {
	Contract *HanaInteractorMockCaller // Generic read-only contract binding to access the raw methods on
}

// HanaInteractorMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaInteractorMockTransactorRaw struct {
	Contract *HanaInteractorMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaInteractorMock creates a new instance of HanaInteractorMock, bound to a specific deployed contract.
func NewHanaInteractorMock(address common.Address, backend bind.ContractBackend) (*HanaInteractorMock, error) {
	contract, err := bindHanaInteractorMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorMock{HanaInteractorMockCaller: HanaInteractorMockCaller{contract: contract}, HanaInteractorMockTransactor: HanaInteractorMockTransactor{contract: contract}, HanaInteractorMockFilterer: HanaInteractorMockFilterer{contract: contract}}, nil
}

// NewHanaInteractorMockCaller creates a new read-only instance of HanaInteractorMock, bound to a specific deployed contract.
func NewHanaInteractorMockCaller(address common.Address, caller bind.ContractCaller) (*HanaInteractorMockCaller, error) {
	contract, err := bindHanaInteractorMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorMockCaller{contract: contract}, nil
}

// NewHanaInteractorMockTransactor creates a new write-only instance of HanaInteractorMock, bound to a specific deployed contract.
func NewHanaInteractorMockTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaInteractorMockTransactor, error) {
	contract, err := bindHanaInteractorMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorMockTransactor{contract: contract}, nil
}

// NewHanaInteractorMockFilterer creates a new log filterer instance of HanaInteractorMock, bound to a specific deployed contract.
func NewHanaInteractorMockFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaInteractorMockFilterer, error) {
	contract, err := bindHanaInteractorMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorMockFilterer{contract: contract}, nil
}

// bindHanaInteractorMock binds a generic wrapper to an already deployed contract.
func bindHanaInteractorMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaInteractorMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaInteractorMock *HanaInteractorMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaInteractorMock.Contract.HanaInteractorMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaInteractorMock *HanaInteractorMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.HanaInteractorMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaInteractorMock *HanaInteractorMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.HanaInteractorMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaInteractorMock *HanaInteractorMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaInteractorMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaInteractorMock *HanaInteractorMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaInteractorMock *HanaInteractorMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.contract.Transact(opts, method, params...)
}

// Connector is a free data retrieval call binding the contract method 0x83f3084f.
//
// Solidity: function connector() view returns(address)
func (_HanaInteractorMock *HanaInteractorMockCaller) Connector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaInteractorMock.contract.Call(opts, &out, "connector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Connector is a free data retrieval call binding the contract method 0x83f3084f.
//
// Solidity: function connector() view returns(address)
func (_HanaInteractorMock *HanaInteractorMockSession) Connector() (common.Address, error) {
	return _HanaInteractorMock.Contract.Connector(&_HanaInteractorMock.CallOpts)
}

// Connector is a free data retrieval call binding the contract method 0x83f3084f.
//
// Solidity: function connector() view returns(address)
func (_HanaInteractorMock *HanaInteractorMockCallerSession) Connector() (common.Address, error) {
	return _HanaInteractorMock.Contract.Connector(&_HanaInteractorMock.CallOpts)
}

// InteractorsByChainId is a free data retrieval call binding the contract method 0x2618143f.
//
// Solidity: function interactorsByChainId(uint256 ) view returns(bytes)
func (_HanaInteractorMock *HanaInteractorMockCaller) InteractorsByChainId(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _HanaInteractorMock.contract.Call(opts, &out, "interactorsByChainId", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// InteractorsByChainId is a free data retrieval call binding the contract method 0x2618143f.
//
// Solidity: function interactorsByChainId(uint256 ) view returns(bytes)
func (_HanaInteractorMock *HanaInteractorMockSession) InteractorsByChainId(arg0 *big.Int) ([]byte, error) {
	return _HanaInteractorMock.Contract.InteractorsByChainId(&_HanaInteractorMock.CallOpts, arg0)
}

// InteractorsByChainId is a free data retrieval call binding the contract method 0x2618143f.
//
// Solidity: function interactorsByChainId(uint256 ) view returns(bytes)
func (_HanaInteractorMock *HanaInteractorMockCallerSession) InteractorsByChainId(arg0 *big.Int) ([]byte, error) {
	return _HanaInteractorMock.Contract.InteractorsByChainId(&_HanaInteractorMock.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HanaInteractorMock *HanaInteractorMockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaInteractorMock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HanaInteractorMock *HanaInteractorMockSession) Owner() (common.Address, error) {
	return _HanaInteractorMock.Contract.Owner(&_HanaInteractorMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HanaInteractorMock *HanaInteractorMockCallerSession) Owner() (common.Address, error) {
	return _HanaInteractorMock.Contract.Owner(&_HanaInteractorMock.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_HanaInteractorMock *HanaInteractorMockCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaInteractorMock.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_HanaInteractorMock *HanaInteractorMockSession) PendingOwner() (common.Address, error) {
	return _HanaInteractorMock.Contract.PendingOwner(&_HanaInteractorMock.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_HanaInteractorMock *HanaInteractorMockCallerSession) PendingOwner() (common.Address, error) {
	return _HanaInteractorMock.Contract.PendingOwner(&_HanaInteractorMock.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_HanaInteractorMock *HanaInteractorMockTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaInteractorMock.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_HanaInteractorMock *HanaInteractorMockSession) AcceptOwnership() (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.AcceptOwnership(&_HanaInteractorMock.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_HanaInteractorMock *HanaInteractorMockTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.AcceptOwnership(&_HanaInteractorMock.TransactOpts)
}

// OnHanaMessage is a paid mutator transaction binding the contract method 0xb204be93.
//
// Solidity: function onHanaMessage((bytes,uint256,address,uint256,bytes) hanaMessage) returns()
func (_HanaInteractorMock *HanaInteractorMockTransactor) OnHanaMessage(opts *bind.TransactOpts, hanaMessage HanaInterfacesHanaMessage) (*types.Transaction, error) {
	return _HanaInteractorMock.contract.Transact(opts, "onHanaMessage", hanaMessage)
}

// OnHanaMessage is a paid mutator transaction binding the contract method 0xb204be93.
//
// Solidity: function onHanaMessage((bytes,uint256,address,uint256,bytes) hanaMessage) returns()
func (_HanaInteractorMock *HanaInteractorMockSession) OnHanaMessage(hanaMessage HanaInterfacesHanaMessage) (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.OnHanaMessage(&_HanaInteractorMock.TransactOpts, hanaMessage)
}

// OnHanaMessage is a paid mutator transaction binding the contract method 0xb204be93.
//
// Solidity: function onHanaMessage((bytes,uint256,address,uint256,bytes) hanaMessage) returns()
func (_HanaInteractorMock *HanaInteractorMockTransactorSession) OnHanaMessage(hanaMessage HanaInterfacesHanaMessage) (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.OnHanaMessage(&_HanaInteractorMock.TransactOpts, hanaMessage)
}

// OnHanaRevert is a paid mutator transaction binding the contract method 0x4204cf9b.
//
// Solidity: function onHanaRevert((address,uint256,bytes,uint256,uint256,bytes) hanaRevert) returns()
func (_HanaInteractorMock *HanaInteractorMockTransactor) OnHanaRevert(opts *bind.TransactOpts, hanaRevert HanaInterfacesHanaRevert) (*types.Transaction, error) {
	return _HanaInteractorMock.contract.Transact(opts, "onHanaRevert", hanaRevert)
}

// OnHanaRevert is a paid mutator transaction binding the contract method 0x4204cf9b.
//
// Solidity: function onHanaRevert((address,uint256,bytes,uint256,uint256,bytes) hanaRevert) returns()
func (_HanaInteractorMock *HanaInteractorMockSession) OnHanaRevert(hanaRevert HanaInterfacesHanaRevert) (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.OnHanaRevert(&_HanaInteractorMock.TransactOpts, hanaRevert)
}

// OnHanaRevert is a paid mutator transaction binding the contract method 0x4204cf9b.
//
// Solidity: function onHanaRevert((address,uint256,bytes,uint256,uint256,bytes) hanaRevert) returns()
func (_HanaInteractorMock *HanaInteractorMockTransactorSession) OnHanaRevert(hanaRevert HanaInterfacesHanaRevert) (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.OnHanaRevert(&_HanaInteractorMock.TransactOpts, hanaRevert)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HanaInteractorMock *HanaInteractorMockTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaInteractorMock.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HanaInteractorMock *HanaInteractorMockSession) RenounceOwnership() (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.RenounceOwnership(&_HanaInteractorMock.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HanaInteractorMock *HanaInteractorMockTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.RenounceOwnership(&_HanaInteractorMock.TransactOpts)
}

// SetInteractorByChainId is a paid mutator transaction binding the contract method 0x4fd3f7d7.
//
// Solidity: function setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) returns()
func (_HanaInteractorMock *HanaInteractorMockTransactor) SetInteractorByChainId(opts *bind.TransactOpts, destinationChainId *big.Int, contractAddress []byte) (*types.Transaction, error) {
	return _HanaInteractorMock.contract.Transact(opts, "setInteractorByChainId", destinationChainId, contractAddress)
}

// SetInteractorByChainId is a paid mutator transaction binding the contract method 0x4fd3f7d7.
//
// Solidity: function setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) returns()
func (_HanaInteractorMock *HanaInteractorMockSession) SetInteractorByChainId(destinationChainId *big.Int, contractAddress []byte) (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.SetInteractorByChainId(&_HanaInteractorMock.TransactOpts, destinationChainId, contractAddress)
}

// SetInteractorByChainId is a paid mutator transaction binding the contract method 0x4fd3f7d7.
//
// Solidity: function setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) returns()
func (_HanaInteractorMock *HanaInteractorMockTransactorSession) SetInteractorByChainId(destinationChainId *big.Int, contractAddress []byte) (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.SetInteractorByChainId(&_HanaInteractorMock.TransactOpts, destinationChainId, contractAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HanaInteractorMock *HanaInteractorMockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HanaInteractorMock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HanaInteractorMock *HanaInteractorMockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.TransferOwnership(&_HanaInteractorMock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HanaInteractorMock *HanaInteractorMockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HanaInteractorMock.Contract.TransferOwnership(&_HanaInteractorMock.TransactOpts, newOwner)
}

// HanaInteractorMockOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the HanaInteractorMock contract.
type HanaInteractorMockOwnershipTransferStartedIterator struct {
	Event *HanaInteractorMockOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *HanaInteractorMockOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaInteractorMockOwnershipTransferStarted)
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
		it.Event = new(HanaInteractorMockOwnershipTransferStarted)
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
func (it *HanaInteractorMockOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaInteractorMockOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaInteractorMockOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the HanaInteractorMock contract.
type HanaInteractorMockOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_HanaInteractorMock *HanaInteractorMockFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HanaInteractorMockOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HanaInteractorMock.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorMockOwnershipTransferStartedIterator{contract: _HanaInteractorMock.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_HanaInteractorMock *HanaInteractorMockFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *HanaInteractorMockOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HanaInteractorMock.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaInteractorMockOwnershipTransferStarted)
				if err := _HanaInteractorMock.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_HanaInteractorMock *HanaInteractorMockFilterer) ParseOwnershipTransferStarted(log types.Log) (*HanaInteractorMockOwnershipTransferStarted, error) {
	event := new(HanaInteractorMockOwnershipTransferStarted)
	if err := _HanaInteractorMock.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaInteractorMockOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HanaInteractorMock contract.
type HanaInteractorMockOwnershipTransferredIterator struct {
	Event *HanaInteractorMockOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HanaInteractorMockOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaInteractorMockOwnershipTransferred)
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
		it.Event = new(HanaInteractorMockOwnershipTransferred)
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
func (it *HanaInteractorMockOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaInteractorMockOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaInteractorMockOwnershipTransferred represents a OwnershipTransferred event raised by the HanaInteractorMock contract.
type HanaInteractorMockOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HanaInteractorMock *HanaInteractorMockFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HanaInteractorMockOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HanaInteractorMock.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HanaInteractorMockOwnershipTransferredIterator{contract: _HanaInteractorMock.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HanaInteractorMock *HanaInteractorMockFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HanaInteractorMockOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HanaInteractorMock.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaInteractorMockOwnershipTransferred)
				if err := _HanaInteractorMock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_HanaInteractorMock *HanaInteractorMockFilterer) ParseOwnershipTransferred(log types.Log) (*HanaInteractorMockOwnershipTransferred, error) {
	event := new(HanaInteractorMockOwnershipTransferred)
	if err := _HanaInteractorMock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
