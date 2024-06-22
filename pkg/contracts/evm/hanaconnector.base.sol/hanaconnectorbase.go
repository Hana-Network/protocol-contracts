// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanaconnector

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

// HanaConnectorBaseMetaData contains all meta data concerning the HanaConnectorBase contract.
var HanaConnectorBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hanaToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tssAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tssAddressUpdater_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauserAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTss\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssOrUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"ExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HanaTranserError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hanaTxSenderAddress\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hanaValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"HanaReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingHanaValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"HanaReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceTxOriginAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hanaValueAndGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hanaParams\",\"type\":\"bytes\"}],\"name\":\"HanaSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updaterAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTssAddress\",\"type\":\"address\"}],\"name\":\"PauserAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTssAddress\",\"type\":\"address\"}],\"name\":\"TSSAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"hanaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"hanaTxSenderAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hanaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"onReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingHanaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"onRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceTssAddressUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"hanaValueAndGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"hanaParams\",\"type\":\"bytes\"}],\"internalType\":\"structHanaInterfaces.SendInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pauserAddress_\",\"type\":\"address\"}],\"name\":\"updatePauserAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tssAddress_\",\"type\":\"address\"}],\"name\":\"updateTssAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620012c3380380620012c383398181016040528101906200003791906200027c565b60008060006101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480620000b95750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b80620000f15750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b80620001295750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b1562000161576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b8152505082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505062000341565b600081519050620002768162000327565b92915050565b6000806000806080858703121562000299576200029862000322565b5b6000620002a98782880162000265565b9450506020620002bc8782880162000265565b9350506040620002cf8782880162000265565b9250506060620002e28782880162000265565b91505092959194509250565b6000620002fb8262000302565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200033281620002ee565b81146200033e57600080fd5b50565b60805160601c610f6362000360600039600061031e0152610f636000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80636128480f1161008c5780639122c344116100665780639122c344146101a2578063942a5e16146101be578063ec026901146101da578063f7fb869b146101f6576100cf565b80636128480f14610172578063779e3b631461018e5780638456cb5914610198576100cf565b806329dd214d146100d4578063328a01d0146100f05780633f4ba83a1461010e5780635b112591146101185780635c975abb146101365780635e694a9214610154575b600080fd5b6100ee60048036038101906100e99190610bfa565b610214565b005b6100f861021e565b6040516101059190610d76565b60405180910390f35b610116610244565b005b6101206102e0565b60405161012d9190610d76565b60405180910390f35b61013e610306565b60405161014b9190610dba565b60405180910390f35b61015c61031c565b6040516101699190610d76565b60405180910390f35b61018c60048036038101906101879190610aeb565b610340565b005b6101966104b6565b005b6101a0610636565b005b6101bc60048036038101906101b79190610aeb565b6106d2565b005b6101d860048036038101906101d39190610b18565b6108a4565b005b6101f460048036038101906101ef9190610cc9565b6108af565b005b6101fe6108b2565b60405161020b9190610d76565b60405180910390f35b5050505050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102d657336040517f4677a0d30000000000000000000000000000000000000000000000000000000081526004016102cd9190610d76565b60405180910390fd5b6102de6108d8565b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900460ff16905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103d257336040517f4677a0d30000000000000000000000000000000000000000000000000000000081526004016103c99190610d76565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610439576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d039733826040516104ab929190610d91565b60405180910390a150565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461054857336040517fe700765e00000000000000000000000000000000000000000000000000000000815260040161053f9190610d76565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156105d1576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106c857336040517f4677a0d30000000000000000000000000000000000000000000000000000000081526004016106bf9190610d76565b60405180910390fd5b6106d061093a565b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415801561077e5750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b156107c057336040517fcdfcef970000000000000000000000000000000000000000000000000000000081526004016107b79190610d76565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610827576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff3382604051610899929190610d91565b60405180910390a150565b505050505050505050565b50565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6108e061099c565b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6109236109e5565b6040516109309190610d76565b60405180910390a1565b6109426109ed565b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586109856109e5565b6040516109929190610d76565b60405180910390a1565b6109a4610306565b6109e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109da90610dd5565b60405180910390fd5b565b600033905090565b6109f5610306565b15610a35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2c90610df5565b60405180910390fd5b565b600081359050610a4681610ee8565b92915050565b600081359050610a5b81610eff565b92915050565b60008083601f840112610a7757610a76610e7d565b5b8235905067ffffffffffffffff811115610a9457610a93610e78565b5b602083019150836001820283011115610ab057610aaf610e87565b5b9250929050565b600060c08284031215610acd57610acc610e82565b5b81905092915050565b600081359050610ae581610f16565b92915050565b600060208284031215610b0157610b00610e91565b5b6000610b0f84828501610a37565b91505092915050565b600080600080600080600080600060e08a8c031215610b3a57610b39610e91565b5b6000610b488c828d01610a37565b9950506020610b598c828d01610ad6565b98505060408a013567ffffffffffffffff811115610b7a57610b79610e8c565b5b610b868c828d01610a61565b97509750506060610b998c828d01610ad6565b9550506080610baa8c828d01610ad6565b94505060a08a013567ffffffffffffffff811115610bcb57610bca610e8c565b5b610bd78c828d01610a61565b935093505060c0610bea8c828d01610a4c565b9150509295985092959850929598565b60008060008060008060008060c0898b031215610c1a57610c19610e91565b5b600089013567ffffffffffffffff811115610c3857610c37610e8c565b5b610c448b828c01610a61565b98509850506020610c578b828c01610ad6565b9650506040610c688b828c01610a37565b9550506060610c798b828c01610ad6565b945050608089013567ffffffffffffffff811115610c9a57610c99610e8c565b5b610ca68b828c01610a61565b935093505060a0610cb98b828c01610a4c565b9150509295985092959890939650565b600060208284031215610cdf57610cde610e91565b5b600082013567ffffffffffffffff811115610cfd57610cfc610e8c565b5b610d0984828501610ab7565b91505092915050565b610d1b81610e26565b82525050565b610d2a81610e38565b82525050565b6000610d3d601483610e15565b9150610d4882610e96565b602082019050919050565b6000610d60601083610e15565b9150610d6b82610ebf565b602082019050919050565b6000602082019050610d8b6000830184610d12565b92915050565b6000604082019050610da66000830185610d12565b610db36020830184610d12565b9392505050565b6000602082019050610dcf6000830184610d21565b92915050565b60006020820190508181036000830152610dee81610d30565b9050919050565b60006020820190508181036000830152610e0e81610d53565b9050919050565b600082825260208201905092915050565b6000610e3182610e4e565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b610ef181610e26565b8114610efc57600080fd5b50565b610f0881610e44565b8114610f1357600080fd5b50565b610f1f81610e6e565b8114610f2a57600080fd5b5056fea2646970667358221220e70ae91aaf383673bff278d79014abc0db26e424f09ef233ec7a0275fe41468564736f6c63430008070033",
}

// HanaConnectorBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaConnectorBaseMetaData.ABI instead.
var HanaConnectorBaseABI = HanaConnectorBaseMetaData.ABI

// HanaConnectorBaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HanaConnectorBaseMetaData.Bin instead.
var HanaConnectorBaseBin = HanaConnectorBaseMetaData.Bin

// DeployHanaConnectorBase deploys a new Ethereum contract, binding an instance of HanaConnectorBase to it.
func DeployHanaConnectorBase(auth *bind.TransactOpts, backend bind.ContractBackend, hanaToken_ common.Address, tssAddress_ common.Address, tssAddressUpdater_ common.Address, pauserAddress_ common.Address) (common.Address, *types.Transaction, *HanaConnectorBase, error) {
	parsed, err := HanaConnectorBaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HanaConnectorBaseBin), backend, hanaToken_, tssAddress_, tssAddressUpdater_, pauserAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HanaConnectorBase{HanaConnectorBaseCaller: HanaConnectorBaseCaller{contract: contract}, HanaConnectorBaseTransactor: HanaConnectorBaseTransactor{contract: contract}, HanaConnectorBaseFilterer: HanaConnectorBaseFilterer{contract: contract}}, nil
}

// HanaConnectorBase is an auto generated Go binding around an Ethereum contract.
type HanaConnectorBase struct {
	HanaConnectorBaseCaller     // Read-only binding to the contract
	HanaConnectorBaseTransactor // Write-only binding to the contract
	HanaConnectorBaseFilterer   // Log filterer for contract events
}

// HanaConnectorBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaConnectorBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaConnectorBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaConnectorBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaConnectorBaseSession struct {
	Contract     *HanaConnectorBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HanaConnectorBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaConnectorBaseCallerSession struct {
	Contract *HanaConnectorBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// HanaConnectorBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaConnectorBaseTransactorSession struct {
	Contract     *HanaConnectorBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// HanaConnectorBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaConnectorBaseRaw struct {
	Contract *HanaConnectorBase // Generic contract binding to access the raw methods on
}

// HanaConnectorBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaConnectorBaseCallerRaw struct {
	Contract *HanaConnectorBaseCaller // Generic read-only contract binding to access the raw methods on
}

// HanaConnectorBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaConnectorBaseTransactorRaw struct {
	Contract *HanaConnectorBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaConnectorBase creates a new instance of HanaConnectorBase, bound to a specific deployed contract.
func NewHanaConnectorBase(address common.Address, backend bind.ContractBackend) (*HanaConnectorBase, error) {
	contract, err := bindHanaConnectorBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorBase{HanaConnectorBaseCaller: HanaConnectorBaseCaller{contract: contract}, HanaConnectorBaseTransactor: HanaConnectorBaseTransactor{contract: contract}, HanaConnectorBaseFilterer: HanaConnectorBaseFilterer{contract: contract}}, nil
}

// NewHanaConnectorBaseCaller creates a new read-only instance of HanaConnectorBase, bound to a specific deployed contract.
func NewHanaConnectorBaseCaller(address common.Address, caller bind.ContractCaller) (*HanaConnectorBaseCaller, error) {
	contract, err := bindHanaConnectorBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorBaseCaller{contract: contract}, nil
}

// NewHanaConnectorBaseTransactor creates a new write-only instance of HanaConnectorBase, bound to a specific deployed contract.
func NewHanaConnectorBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaConnectorBaseTransactor, error) {
	contract, err := bindHanaConnectorBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorBaseTransactor{contract: contract}, nil
}

// NewHanaConnectorBaseFilterer creates a new log filterer instance of HanaConnectorBase, bound to a specific deployed contract.
func NewHanaConnectorBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaConnectorBaseFilterer, error) {
	contract, err := bindHanaConnectorBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorBaseFilterer{contract: contract}, nil
}

// bindHanaConnectorBase binds a generic wrapper to an already deployed contract.
func bindHanaConnectorBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaConnectorBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaConnectorBase *HanaConnectorBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaConnectorBase.Contract.HanaConnectorBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaConnectorBase *HanaConnectorBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.HanaConnectorBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaConnectorBase *HanaConnectorBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.HanaConnectorBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaConnectorBase *HanaConnectorBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaConnectorBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaConnectorBase *HanaConnectorBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaConnectorBase *HanaConnectorBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.contract.Transact(opts, method, params...)
}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaConnectorBase *HanaConnectorBaseCaller) HanaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaConnectorBase.contract.Call(opts, &out, "hanaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaConnectorBase *HanaConnectorBaseSession) HanaToken() (common.Address, error) {
	return _HanaConnectorBase.Contract.HanaToken(&_HanaConnectorBase.CallOpts)
}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaConnectorBase *HanaConnectorBaseCallerSession) HanaToken() (common.Address, error) {
	return _HanaConnectorBase.Contract.HanaToken(&_HanaConnectorBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HanaConnectorBase *HanaConnectorBaseCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HanaConnectorBase.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HanaConnectorBase *HanaConnectorBaseSession) Paused() (bool, error) {
	return _HanaConnectorBase.Contract.Paused(&_HanaConnectorBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HanaConnectorBase *HanaConnectorBaseCallerSession) Paused() (bool, error) {
	return _HanaConnectorBase.Contract.Paused(&_HanaConnectorBase.CallOpts)
}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_HanaConnectorBase *HanaConnectorBaseCaller) PauserAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaConnectorBase.contract.Call(opts, &out, "pauserAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_HanaConnectorBase *HanaConnectorBaseSession) PauserAddress() (common.Address, error) {
	return _HanaConnectorBase.Contract.PauserAddress(&_HanaConnectorBase.CallOpts)
}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_HanaConnectorBase *HanaConnectorBaseCallerSession) PauserAddress() (common.Address, error) {
	return _HanaConnectorBase.Contract.PauserAddress(&_HanaConnectorBase.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_HanaConnectorBase *HanaConnectorBaseCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaConnectorBase.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_HanaConnectorBase *HanaConnectorBaseSession) TssAddress() (common.Address, error) {
	return _HanaConnectorBase.Contract.TssAddress(&_HanaConnectorBase.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_HanaConnectorBase *HanaConnectorBaseCallerSession) TssAddress() (common.Address, error) {
	return _HanaConnectorBase.Contract.TssAddress(&_HanaConnectorBase.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_HanaConnectorBase *HanaConnectorBaseCaller) TssAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaConnectorBase.contract.Call(opts, &out, "tssAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_HanaConnectorBase *HanaConnectorBaseSession) TssAddressUpdater() (common.Address, error) {
	return _HanaConnectorBase.Contract.TssAddressUpdater(&_HanaConnectorBase.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_HanaConnectorBase *HanaConnectorBaseCallerSession) TssAddressUpdater() (common.Address, error) {
	return _HanaConnectorBase.Contract.TssAddressUpdater(&_HanaConnectorBase.CallOpts)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes hanaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 hanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactor) OnReceive(opts *bind.TransactOpts, hanaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, hanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorBase.contract.Transact(opts, "onReceive", hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes hanaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 hanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorBase *HanaConnectorBaseSession) OnReceive(hanaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, hanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.OnReceive(&_HanaConnectorBase.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes hanaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 hanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactorSession) OnReceive(hanaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, hanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.OnReceive(&_HanaConnectorBase.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address hanaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingHanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactor) OnRevert(opts *bind.TransactOpts, hanaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingHanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorBase.contract.Transact(opts, "onRevert", hanaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingHanaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address hanaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingHanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorBase *HanaConnectorBaseSession) OnRevert(hanaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingHanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.OnRevert(&_HanaConnectorBase.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingHanaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address hanaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingHanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactorSession) OnRevert(hanaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingHanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.OnRevert(&_HanaConnectorBase.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingHanaValue, message, internalSendHash)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorBase.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HanaConnectorBase *HanaConnectorBaseSession) Pause() (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.Pause(&_HanaConnectorBase.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactorSession) Pause() (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.Pause(&_HanaConnectorBase.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactor) RenounceTssAddressUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorBase.contract.Transact(opts, "renounceTssAddressUpdater")
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_HanaConnectorBase *HanaConnectorBaseSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.RenounceTssAddressUpdater(&_HanaConnectorBase.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactorSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.RenounceTssAddressUpdater(&_HanaConnectorBase.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactor) Send(opts *bind.TransactOpts, input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnectorBase.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnectorBase *HanaConnectorBaseSession) Send(input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.Send(&_HanaConnectorBase.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactorSession) Send(input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.Send(&_HanaConnectorBase.TransactOpts, input)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorBase.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HanaConnectorBase *HanaConnectorBaseSession) Unpause() (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.Unpause(&_HanaConnectorBase.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactorSession) Unpause() (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.Unpause(&_HanaConnectorBase.TransactOpts)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactor) UpdatePauserAddress(opts *bind.TransactOpts, pauserAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorBase.contract.Transact(opts, "updatePauserAddress", pauserAddress_)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_HanaConnectorBase *HanaConnectorBaseSession) UpdatePauserAddress(pauserAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.UpdatePauserAddress(&_HanaConnectorBase.TransactOpts, pauserAddress_)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactorSession) UpdatePauserAddress(pauserAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.UpdatePauserAddress(&_HanaConnectorBase.TransactOpts, pauserAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactor) UpdateTssAddress(opts *bind.TransactOpts, tssAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorBase.contract.Transact(opts, "updateTssAddress", tssAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_HanaConnectorBase *HanaConnectorBaseSession) UpdateTssAddress(tssAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.UpdateTssAddress(&_HanaConnectorBase.TransactOpts, tssAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_HanaConnectorBase *HanaConnectorBaseTransactorSession) UpdateTssAddress(tssAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorBase.Contract.UpdateTssAddress(&_HanaConnectorBase.TransactOpts, tssAddress_)
}

// HanaConnectorBaseHanaReceivedIterator is returned from FilterHanaReceived and is used to iterate over the raw logs and unpacked data for HanaReceived events raised by the HanaConnectorBase contract.
type HanaConnectorBaseHanaReceivedIterator struct {
	Event *HanaConnectorBaseHanaReceived // Event containing the contract specifics and raw log

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
func (it *HanaConnectorBaseHanaReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorBaseHanaReceived)
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
		it.Event = new(HanaConnectorBaseHanaReceived)
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
func (it *HanaConnectorBaseHanaReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorBaseHanaReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorBaseHanaReceived represents a HanaReceived event raised by the HanaConnectorBase contract.
type HanaConnectorBaseHanaReceived struct {
	HanaTxSenderAddress []byte
	SourceChainId       *big.Int
	DestinationAddress  common.Address
	HanaValue           *big.Int
	Message             []byte
	InternalSendHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterHanaReceived is a free log retrieval operation binding the contract event 0xb29e5f376a4a399e83dce3b4e7f71bc32f1ab109ccf94574f8f4870cf4aa341f.
//
// Solidity: event HanaReceived(bytes hanaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 hanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) FilterHanaReceived(opts *bind.FilterOpts, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (*HanaConnectorBaseHanaReceivedIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _HanaConnectorBase.contract.FilterLogs(opts, "HanaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorBaseHanaReceivedIterator{contract: _HanaConnectorBase.contract, event: "HanaReceived", logs: logs, sub: sub}, nil
}

// WatchHanaReceived is a free log subscription operation binding the contract event 0xb29e5f376a4a399e83dce3b4e7f71bc32f1ab109ccf94574f8f4870cf4aa341f.
//
// Solidity: event HanaReceived(bytes hanaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 hanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) WatchHanaReceived(opts *bind.WatchOpts, sink chan<- *HanaConnectorBaseHanaReceived, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _HanaConnectorBase.contract.WatchLogs(opts, "HanaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorBaseHanaReceived)
				if err := _HanaConnectorBase.contract.UnpackLog(event, "HanaReceived", log); err != nil {
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

// ParseHanaReceived is a log parse operation binding the contract event 0xb29e5f376a4a399e83dce3b4e7f71bc32f1ab109ccf94574f8f4870cf4aa341f.
//
// Solidity: event HanaReceived(bytes hanaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 hanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) ParseHanaReceived(log types.Log) (*HanaConnectorBaseHanaReceived, error) {
	event := new(HanaConnectorBaseHanaReceived)
	if err := _HanaConnectorBase.contract.UnpackLog(event, "HanaReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorBaseHanaRevertedIterator is returned from FilterHanaReverted and is used to iterate over the raw logs and unpacked data for HanaReverted events raised by the HanaConnectorBase contract.
type HanaConnectorBaseHanaRevertedIterator struct {
	Event *HanaConnectorBaseHanaReverted // Event containing the contract specifics and raw log

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
func (it *HanaConnectorBaseHanaRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorBaseHanaReverted)
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
		it.Event = new(HanaConnectorBaseHanaReverted)
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
func (it *HanaConnectorBaseHanaRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorBaseHanaRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorBaseHanaReverted represents a HanaReverted event raised by the HanaConnectorBase contract.
type HanaConnectorBaseHanaReverted struct {
	HanaTxSenderAddress common.Address
	SourceChainId       *big.Int
	DestinationChainId  *big.Int
	DestinationAddress  []byte
	RemainingHanaValue  *big.Int
	Message             []byte
	InternalSendHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterHanaReverted is a free log retrieval operation binding the contract event 0xa0589272400cb97d69e64e37e4dbd1f2aaa79d5b3aa0688e6f2d992ccf500921.
//
// Solidity: event HanaReverted(address hanaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingHanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) FilterHanaReverted(opts *bind.FilterOpts, destinationChainId []*big.Int, internalSendHash [][32]byte) (*HanaConnectorBaseHanaRevertedIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _HanaConnectorBase.contract.FilterLogs(opts, "HanaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorBaseHanaRevertedIterator{contract: _HanaConnectorBase.contract, event: "HanaReverted", logs: logs, sub: sub}, nil
}

// WatchHanaReverted is a free log subscription operation binding the contract event 0xa0589272400cb97d69e64e37e4dbd1f2aaa79d5b3aa0688e6f2d992ccf500921.
//
// Solidity: event HanaReverted(address hanaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingHanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) WatchHanaReverted(opts *bind.WatchOpts, sink chan<- *HanaConnectorBaseHanaReverted, destinationChainId []*big.Int, internalSendHash [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _HanaConnectorBase.contract.WatchLogs(opts, "HanaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorBaseHanaReverted)
				if err := _HanaConnectorBase.contract.UnpackLog(event, "HanaReverted", log); err != nil {
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

// ParseHanaReverted is a log parse operation binding the contract event 0xa0589272400cb97d69e64e37e4dbd1f2aaa79d5b3aa0688e6f2d992ccf500921.
//
// Solidity: event HanaReverted(address hanaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingHanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) ParseHanaReverted(log types.Log) (*HanaConnectorBaseHanaReverted, error) {
	event := new(HanaConnectorBaseHanaReverted)
	if err := _HanaConnectorBase.contract.UnpackLog(event, "HanaReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorBaseHanaSentIterator is returned from FilterHanaSent and is used to iterate over the raw logs and unpacked data for HanaSent events raised by the HanaConnectorBase contract.
type HanaConnectorBaseHanaSentIterator struct {
	Event *HanaConnectorBaseHanaSent // Event containing the contract specifics and raw log

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
func (it *HanaConnectorBaseHanaSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorBaseHanaSent)
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
		it.Event = new(HanaConnectorBaseHanaSent)
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
func (it *HanaConnectorBaseHanaSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorBaseHanaSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorBaseHanaSent represents a HanaSent event raised by the HanaConnectorBase contract.
type HanaConnectorBaseHanaSent struct {
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
func (_HanaConnectorBase *HanaConnectorBaseFilterer) FilterHanaSent(opts *bind.FilterOpts, hanaTxSenderAddress []common.Address, destinationChainId []*big.Int) (*HanaConnectorBaseHanaSentIterator, error) {

	var hanaTxSenderAddressRule []interface{}
	for _, hanaTxSenderAddressItem := range hanaTxSenderAddress {
		hanaTxSenderAddressRule = append(hanaTxSenderAddressRule, hanaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _HanaConnectorBase.contract.FilterLogs(opts, "HanaSent", hanaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorBaseHanaSentIterator{contract: _HanaConnectorBase.contract, event: "HanaSent", logs: logs, sub: sub}, nil
}

// WatchHanaSent is a free log subscription operation binding the contract event 0x7a0dd478962bca8db7a72b684fbfd3404be69978e5f8f1c8eab2cfe4724ebaee.
//
// Solidity: event HanaSent(address sourceTxOriginAddress, address indexed hanaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 hanaValueAndGas, uint256 destinationGasLimit, bytes message, bytes hanaParams)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) WatchHanaSent(opts *bind.WatchOpts, sink chan<- *HanaConnectorBaseHanaSent, hanaTxSenderAddress []common.Address, destinationChainId []*big.Int) (event.Subscription, error) {

	var hanaTxSenderAddressRule []interface{}
	for _, hanaTxSenderAddressItem := range hanaTxSenderAddress {
		hanaTxSenderAddressRule = append(hanaTxSenderAddressRule, hanaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _HanaConnectorBase.contract.WatchLogs(opts, "HanaSent", hanaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorBaseHanaSent)
				if err := _HanaConnectorBase.contract.UnpackLog(event, "HanaSent", log); err != nil {
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
func (_HanaConnectorBase *HanaConnectorBaseFilterer) ParseHanaSent(log types.Log) (*HanaConnectorBaseHanaSent, error) {
	event := new(HanaConnectorBaseHanaSent)
	if err := _HanaConnectorBase.contract.UnpackLog(event, "HanaSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorBasePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the HanaConnectorBase contract.
type HanaConnectorBasePausedIterator struct {
	Event *HanaConnectorBasePaused // Event containing the contract specifics and raw log

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
func (it *HanaConnectorBasePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorBasePaused)
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
		it.Event = new(HanaConnectorBasePaused)
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
func (it *HanaConnectorBasePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorBasePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorBasePaused represents a Paused event raised by the HanaConnectorBase contract.
type HanaConnectorBasePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) FilterPaused(opts *bind.FilterOpts) (*HanaConnectorBasePausedIterator, error) {

	logs, sub, err := _HanaConnectorBase.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &HanaConnectorBasePausedIterator{contract: _HanaConnectorBase.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *HanaConnectorBasePaused) (event.Subscription, error) {

	logs, sub, err := _HanaConnectorBase.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorBasePaused)
				if err := _HanaConnectorBase.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) ParsePaused(log types.Log) (*HanaConnectorBasePaused, error) {
	event := new(HanaConnectorBasePaused)
	if err := _HanaConnectorBase.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorBasePauserAddressUpdatedIterator is returned from FilterPauserAddressUpdated and is used to iterate over the raw logs and unpacked data for PauserAddressUpdated events raised by the HanaConnectorBase contract.
type HanaConnectorBasePauserAddressUpdatedIterator struct {
	Event *HanaConnectorBasePauserAddressUpdated // Event containing the contract specifics and raw log

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
func (it *HanaConnectorBasePauserAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorBasePauserAddressUpdated)
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
		it.Event = new(HanaConnectorBasePauserAddressUpdated)
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
func (it *HanaConnectorBasePauserAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorBasePauserAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorBasePauserAddressUpdated represents a PauserAddressUpdated event raised by the HanaConnectorBase contract.
type HanaConnectorBasePauserAddressUpdated struct {
	UpdaterAddress common.Address
	NewTssAddress  common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPauserAddressUpdated is a free log retrieval operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address updaterAddress, address newTssAddress)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) FilterPauserAddressUpdated(opts *bind.FilterOpts) (*HanaConnectorBasePauserAddressUpdatedIterator, error) {

	logs, sub, err := _HanaConnectorBase.contract.FilterLogs(opts, "PauserAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &HanaConnectorBasePauserAddressUpdatedIterator{contract: _HanaConnectorBase.contract, event: "PauserAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchPauserAddressUpdated is a free log subscription operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address updaterAddress, address newTssAddress)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) WatchPauserAddressUpdated(opts *bind.WatchOpts, sink chan<- *HanaConnectorBasePauserAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _HanaConnectorBase.contract.WatchLogs(opts, "PauserAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorBasePauserAddressUpdated)
				if err := _HanaConnectorBase.contract.UnpackLog(event, "PauserAddressUpdated", log); err != nil {
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

// ParsePauserAddressUpdated is a log parse operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address updaterAddress, address newTssAddress)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) ParsePauserAddressUpdated(log types.Log) (*HanaConnectorBasePauserAddressUpdated, error) {
	event := new(HanaConnectorBasePauserAddressUpdated)
	if err := _HanaConnectorBase.contract.UnpackLog(event, "PauserAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorBaseTSSAddressUpdatedIterator is returned from FilterTSSAddressUpdated and is used to iterate over the raw logs and unpacked data for TSSAddressUpdated events raised by the HanaConnectorBase contract.
type HanaConnectorBaseTSSAddressUpdatedIterator struct {
	Event *HanaConnectorBaseTSSAddressUpdated // Event containing the contract specifics and raw log

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
func (it *HanaConnectorBaseTSSAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorBaseTSSAddressUpdated)
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
		it.Event = new(HanaConnectorBaseTSSAddressUpdated)
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
func (it *HanaConnectorBaseTSSAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorBaseTSSAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorBaseTSSAddressUpdated represents a TSSAddressUpdated event raised by the HanaConnectorBase contract.
type HanaConnectorBaseTSSAddressUpdated struct {
	HanaTxSenderAddress common.Address
	NewTssAddress       common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTSSAddressUpdated is a free log retrieval operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address hanaTxSenderAddress, address newTssAddress)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) FilterTSSAddressUpdated(opts *bind.FilterOpts) (*HanaConnectorBaseTSSAddressUpdatedIterator, error) {

	logs, sub, err := _HanaConnectorBase.contract.FilterLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &HanaConnectorBaseTSSAddressUpdatedIterator{contract: _HanaConnectorBase.contract, event: "TSSAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchTSSAddressUpdated is a free log subscription operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address hanaTxSenderAddress, address newTssAddress)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) WatchTSSAddressUpdated(opts *bind.WatchOpts, sink chan<- *HanaConnectorBaseTSSAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _HanaConnectorBase.contract.WatchLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorBaseTSSAddressUpdated)
				if err := _HanaConnectorBase.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
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

// ParseTSSAddressUpdated is a log parse operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address hanaTxSenderAddress, address newTssAddress)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) ParseTSSAddressUpdated(log types.Log) (*HanaConnectorBaseTSSAddressUpdated, error) {
	event := new(HanaConnectorBaseTSSAddressUpdated)
	if err := _HanaConnectorBase.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorBaseUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the HanaConnectorBase contract.
type HanaConnectorBaseUnpausedIterator struct {
	Event *HanaConnectorBaseUnpaused // Event containing the contract specifics and raw log

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
func (it *HanaConnectorBaseUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorBaseUnpaused)
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
		it.Event = new(HanaConnectorBaseUnpaused)
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
func (it *HanaConnectorBaseUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorBaseUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorBaseUnpaused represents a Unpaused event raised by the HanaConnectorBase contract.
type HanaConnectorBaseUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) FilterUnpaused(opts *bind.FilterOpts) (*HanaConnectorBaseUnpausedIterator, error) {

	logs, sub, err := _HanaConnectorBase.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &HanaConnectorBaseUnpausedIterator{contract: _HanaConnectorBase.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *HanaConnectorBaseUnpaused) (event.Subscription, error) {

	logs, sub, err := _HanaConnectorBase.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorBaseUnpaused)
				if err := _HanaConnectorBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HanaConnectorBase *HanaConnectorBaseFilterer) ParseUnpaused(log types.Log) (*HanaConnectorBaseUnpaused, error) {
	event := new(HanaConnectorBaseUnpaused)
	if err := _HanaConnectorBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
