// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanaconnectorhevm

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whana_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedHanaSent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyFungibleModule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWHANAOrFungible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WHANATransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongValue\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hanaTxSenderAddress\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hanaValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"HanaReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingHanaValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"HanaReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceTxOriginAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hanaValueAndGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hanaParams\",\"type\":\"bytes\"}],\"name\":\"HanaSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"whana_\",\"type\":\"address\"}],\"name\":\"SetWHANA\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"hanaTxSenderAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hanaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"onReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingHanaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"onRevert\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"hanaValueAndGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"hanaParams\",\"type\":\"bytes\"}],\"internalType\":\"structHanaInterfaces.SendInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whana_\",\"type\":\"address\"}],\"name\":\"setWhanaAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whana\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200177e3803806200177e833981810160405281019062000037919062000095565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200011a565b6000815190506200008f8162000100565b92915050565b600060208284031215620000ae57620000ad620000fb565b5b6000620000be848285016200007e565b91505092915050565b6000620000d482620000db565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200010b81620000c7565b81146200011757600080fd5b50565b611654806200012a6000396000f3fe6080604052600436106100595760003560e01c806329dd214d146101385780633ce4a5bc14610154578063942a5e161461017f578063a4590ea11461019b578063ca9457a5146101c4578063ec026901146101ef57610133565b366101335760008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141580156100fa575073735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610131576040517fcaca94fd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b610152600480360381019061014d9190610f78565b610218565b005b34801561016057600080fd5b506101696105ce565b6040516101769190611277565b60405180910390f35b61019960048036038101906101949190610e69565b6105e6565b005b3480156101a757600080fd5b506101c260048036038101906101bd9190610e3c565b610990565b005b3480156101d057600080fd5b506101d9610a83565b6040516101e69190611277565b60405180910390f35b3480156101fb57600080fd5b5061021660048036038101906102119190611047565b610aa7565b005b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610291576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8334146102ca576040517f98d4901c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b15801561033257600080fd5b505af1158015610346573d6000803e3d6000fd5b505050505060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3087876040518463ffffffff1660e01b81526004016103a893929190611292565b602060405180830381600087803b1580156103c257600080fd5b505af11580156103d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103fa9190610f4b565b610430576040517ff8e9c1a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600083839050111561056c578473ffffffffffffffffffffffffffffffffffffffff1663b204be936040518060a001604052808b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018981526020018873ffffffffffffffffffffffffffffffffffffffff16815260200187815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152506040518263ffffffff1660e01b815260040161053991906113f3565b600060405180830381600087803b15801561055357600080fd5b505af1158015610567573d6000803e3d6000fd5b505050505b808573ffffffffffffffffffffffffffffffffffffffff16877fb29e5f376a4a399e83dce3b4e7f71bc32f1ab109ccf94574f8f4870cf4aa341f8b8b8989896040516105bc9594939291906113aa565b60405180910390a45050505050505050565b73735b14bb79463307aacbed86daf3322b1e6226ab81565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461065f576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b833414610698576040517f98d4901c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b15801561070057600080fd5b505af1158015610714573d6000803e3d6000fd5b505050505060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd308b876040518463ffffffff1660e01b815260040161077693929190611292565b602060405180830381600087803b15801561079057600080fd5b505af11580156107a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c89190610f4b565b6107fe576040517ff8e9c1a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000838390501115610940578873ffffffffffffffffffffffffffffffffffffffff16634204cf9b6040518060c001604052808c73ffffffffffffffffffffffffffffffffffffffff1681526020018b81526020018a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200188815260200187815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152506040518263ffffffff1660e01b815260040161090d9190611415565b600060405180830381600087803b15801561092757600080fd5b505af115801561093b573d6000803e3d6000fd5b505050505b80857fa0589272400cb97d69e64e37e4dbd1f2aaa79d5b3aa0688e6f2d992ccf5009218b8b8b8b8a8a8a60405161097d9796959493929190611345565b60405180910390a3505050505050505050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a09576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507ffd250d2fc6a8268f30b40917cdfa5b10a3e164e384e03bf4a90192e81ffa27d081604051610a789190611277565b60405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd333084608001356040518463ffffffff1660e01b8152600401610b0893929190611292565b602060405180830381600087803b158015610b2257600080fd5b505af1158015610b36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5a9190610f4b565b610b90576040517ff8e9c1a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d82608001356040518263ffffffff1660e01b8152600401610bed9190611437565b600060405180830381600087803b158015610c0757600080fd5b505af1158015610c1b573d6000803e3d6000fd5b50505050600073735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff168260800135604051610c5d90611262565b60006040518083038185875af1925050503d8060008114610c9a576040519150601f19603f3d011682016040523d82523d6000602084013e610c9f565b606091505b5050905080610cda576040517f9a543f2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600001353373ffffffffffffffffffffffffffffffffffffffff167f7a0dd478962bca8db7a72b684fbfd3404be69978e5f8f1c8eab2cfe4724ebaee32858060200190610d289190611452565b87608001358860400135898060600190610d429190611452565b8b8060a00190610d529190611452565b604051610d67999897969594939291906112c9565b60405180910390a35050565b600081359050610d82816115c2565b92915050565b600081519050610d97816115d9565b92915050565b600081359050610dac816115f0565b92915050565b60008083601f840112610dc857610dc7611586565b5b8235905067ffffffffffffffff811115610de557610de4611581565b5b602083019150836001820283011115610e0157610e0061159a565b5b9250929050565b600060c08284031215610e1e57610e1d611590565b5b81905092915050565b600081359050610e3681611607565b92915050565b600060208284031215610e5257610e516115a9565b5b6000610e6084828501610d73565b91505092915050565b600080600080600080600080600060e08a8c031215610e8b57610e8a6115a9565b5b6000610e998c828d01610d73565b9950506020610eaa8c828d01610e27565b98505060408a013567ffffffffffffffff811115610ecb57610eca6115a4565b5b610ed78c828d01610db2565b97509750506060610eea8c828d01610e27565b9550506080610efb8c828d01610e27565b94505060a08a013567ffffffffffffffff811115610f1c57610f1b6115a4565b5b610f288c828d01610db2565b935093505060c0610f3b8c828d01610d9d565b9150509295985092959850929598565b600060208284031215610f6157610f606115a9565b5b6000610f6f84828501610d88565b91505092915050565b60008060008060008060008060c0898b031215610f9857610f976115a9565b5b600089013567ffffffffffffffff811115610fb657610fb56115a4565b5b610fc28b828c01610db2565b98509850506020610fd58b828c01610e27565b9650506040610fe68b828c01610d73565b9550506060610ff78b828c01610e27565b945050608089013567ffffffffffffffff811115611018576110176115a4565b5b6110248b828c01610db2565b935093505060a06110378b828c01610d9d565b9150509295985092959890939650565b60006020828403121561105d5761105c6115a9565b5b600082013567ffffffffffffffff81111561107b5761107a6115a4565b5b61108784828501610e08565b91505092915050565b611099816114ed565b82525050565b6110a8816114ed565b82525050565b60006110ba83856114d1565b93506110c783858461153f565b6110d0836115ae565b840190509392505050565b60006110e6826114b5565b6110f081856114c0565b935061110081856020860161154e565b611109816115ae565b840191505092915050565b60006111216000836114e2565b915061112c826115bf565b600082019050919050565b600060a083016000830151848203600086015261115482826110db565b91505060208301516111696020860182611244565b50604083015161117c6040860182611090565b50606083015161118f6060860182611244565b50608083015184820360808601526111a782826110db565b9150508091505092915050565b600060c0830160008301516111cc6000860182611090565b5060208301516111df6020860182611244565b50604083015184820360408601526111f782826110db565b915050606083015161120c6060860182611244565b50608083015161121f6080860182611244565b5060a083015184820360a086015261123782826110db565b9150508091505092915050565b61124d81611535565b82525050565b61125c81611535565b82525050565b600061126d82611114565b9150819050919050565b600060208201905061128c600083018461109f565b92915050565b60006060820190506112a7600083018661109f565b6112b4602083018561109f565b6112c16040830184611253565b949350505050565b600060c0820190506112de600083018c61109f565b81810360208301526112f1818a8c6110ae565b90506113006040830189611253565b61130d6060830188611253565b81810360808301526113208186886110ae565b905081810360a08301526113358184866110ae565b90509a9950505050505050505050565b600060a08201905061135a600083018a61109f565b6113676020830189611253565b818103604083015261137a8187896110ae565b90506113896060830186611253565b818103608083015261139c8184866110ae565b905098975050505050505050565b600060608201905081810360008301526113c58187896110ae565b90506113d46020830186611253565b81810360408301526113e78184866110ae565b90509695505050505050565b6000602082019050818103600083015261140d8184611137565b905092915050565b6000602082019050818103600083015261142f81846111b4565b905092915050565b600060208201905061144c6000830184611253565b92915050565b6000808335600160200384360303811261146f5761146e611595565b5b80840192508235915067ffffffffffffffff8211156114915761149061158b565b5b6020830192506001820236038313156114ad576114ac61159f565b5b509250929050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60006114f882611515565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561156c578082015181840152602081019050611551565b8381111561157b576000848401525b50505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b50565b6115cb816114ed565b81146115d657600080fd5b50565b6115e2816114ff565b81146115ed57600080fd5b50565b6115f98161150b565b811461160457600080fd5b50565b61161081611535565b811461161b57600080fd5b5056fea2646970667358221220253c6ad797453974f646ff43de9eab2560ab45a94d677ff4dcf1a7548431af5b64736f6c63430008070033",
}

// HanaConnectorHEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaConnectorHEVMMetaData.ABI instead.
var HanaConnectorHEVMABI = HanaConnectorHEVMMetaData.ABI

// HanaConnectorHEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HanaConnectorHEVMMetaData.Bin instead.
var HanaConnectorHEVMBin = HanaConnectorHEVMMetaData.Bin

// DeployHanaConnectorHEVM deploys a new Ethereum contract, binding an instance of HanaConnectorHEVM to it.
func DeployHanaConnectorHEVM(auth *bind.TransactOpts, backend bind.ContractBackend, whana_ common.Address) (common.Address, *types.Transaction, *HanaConnectorHEVM, error) {
	parsed, err := HanaConnectorHEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HanaConnectorHEVMBin), backend, whana_)
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

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes hanaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 hanaValue, bytes message, bytes32 internalSendHash) payable returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMTransactor) OnReceive(opts *bind.TransactOpts, hanaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, hanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorHEVM.contract.Transact(opts, "onReceive", hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes hanaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 hanaValue, bytes message, bytes32 internalSendHash) payable returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMSession) OnReceive(hanaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, hanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorHEVM.Contract.OnReceive(&_HanaConnectorHEVM.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes hanaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 hanaValue, bytes message, bytes32 internalSendHash) payable returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMTransactorSession) OnReceive(hanaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, hanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorHEVM.Contract.OnReceive(&_HanaConnectorHEVM.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address hanaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingHanaValue, bytes message, bytes32 internalSendHash) payable returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMTransactor) OnRevert(opts *bind.TransactOpts, hanaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingHanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorHEVM.contract.Transact(opts, "onRevert", hanaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingHanaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address hanaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingHanaValue, bytes message, bytes32 internalSendHash) payable returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMSession) OnRevert(hanaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingHanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorHEVM.Contract.OnRevert(&_HanaConnectorHEVM.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingHanaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address hanaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingHanaValue, bytes message, bytes32 internalSendHash) payable returns()
func (_HanaConnectorHEVM *HanaConnectorHEVMTransactorSession) OnRevert(hanaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingHanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorHEVM.Contract.OnRevert(&_HanaConnectorHEVM.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingHanaValue, message, internalSendHash)
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

// HanaConnectorHEVMHanaReceivedIterator is returned from FilterHanaReceived and is used to iterate over the raw logs and unpacked data for HanaReceived events raised by the HanaConnectorHEVM contract.
type HanaConnectorHEVMHanaReceivedIterator struct {
	Event *HanaConnectorHEVMHanaReceived // Event containing the contract specifics and raw log

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
func (it *HanaConnectorHEVMHanaReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorHEVMHanaReceived)
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
		it.Event = new(HanaConnectorHEVMHanaReceived)
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
func (it *HanaConnectorHEVMHanaReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorHEVMHanaReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorHEVMHanaReceived represents a HanaReceived event raised by the HanaConnectorHEVM contract.
type HanaConnectorHEVMHanaReceived struct {
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
func (_HanaConnectorHEVM *HanaConnectorHEVMFilterer) FilterHanaReceived(opts *bind.FilterOpts, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (*HanaConnectorHEVMHanaReceivedIterator, error) {

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

	logs, sub, err := _HanaConnectorHEVM.contract.FilterLogs(opts, "HanaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorHEVMHanaReceivedIterator{contract: _HanaConnectorHEVM.contract, event: "HanaReceived", logs: logs, sub: sub}, nil
}

// WatchHanaReceived is a free log subscription operation binding the contract event 0xb29e5f376a4a399e83dce3b4e7f71bc32f1ab109ccf94574f8f4870cf4aa341f.
//
// Solidity: event HanaReceived(bytes hanaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 hanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorHEVM *HanaConnectorHEVMFilterer) WatchHanaReceived(opts *bind.WatchOpts, sink chan<- *HanaConnectorHEVMHanaReceived, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _HanaConnectorHEVM.contract.WatchLogs(opts, "HanaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorHEVMHanaReceived)
				if err := _HanaConnectorHEVM.contract.UnpackLog(event, "HanaReceived", log); err != nil {
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
func (_HanaConnectorHEVM *HanaConnectorHEVMFilterer) ParseHanaReceived(log types.Log) (*HanaConnectorHEVMHanaReceived, error) {
	event := new(HanaConnectorHEVMHanaReceived)
	if err := _HanaConnectorHEVM.contract.UnpackLog(event, "HanaReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorHEVMHanaRevertedIterator is returned from FilterHanaReverted and is used to iterate over the raw logs and unpacked data for HanaReverted events raised by the HanaConnectorHEVM contract.
type HanaConnectorHEVMHanaRevertedIterator struct {
	Event *HanaConnectorHEVMHanaReverted // Event containing the contract specifics and raw log

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
func (it *HanaConnectorHEVMHanaRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorHEVMHanaReverted)
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
		it.Event = new(HanaConnectorHEVMHanaReverted)
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
func (it *HanaConnectorHEVMHanaRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorHEVMHanaRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorHEVMHanaReverted represents a HanaReverted event raised by the HanaConnectorHEVM contract.
type HanaConnectorHEVMHanaReverted struct {
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
func (_HanaConnectorHEVM *HanaConnectorHEVMFilterer) FilterHanaReverted(opts *bind.FilterOpts, destinationChainId []*big.Int, internalSendHash [][32]byte) (*HanaConnectorHEVMHanaRevertedIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _HanaConnectorHEVM.contract.FilterLogs(opts, "HanaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorHEVMHanaRevertedIterator{contract: _HanaConnectorHEVM.contract, event: "HanaReverted", logs: logs, sub: sub}, nil
}

// WatchHanaReverted is a free log subscription operation binding the contract event 0xa0589272400cb97d69e64e37e4dbd1f2aaa79d5b3aa0688e6f2d992ccf500921.
//
// Solidity: event HanaReverted(address hanaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingHanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorHEVM *HanaConnectorHEVMFilterer) WatchHanaReverted(opts *bind.WatchOpts, sink chan<- *HanaConnectorHEVMHanaReverted, destinationChainId []*big.Int, internalSendHash [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _HanaConnectorHEVM.contract.WatchLogs(opts, "HanaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorHEVMHanaReverted)
				if err := _HanaConnectorHEVM.contract.UnpackLog(event, "HanaReverted", log); err != nil {
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
func (_HanaConnectorHEVM *HanaConnectorHEVMFilterer) ParseHanaReverted(log types.Log) (*HanaConnectorHEVMHanaReverted, error) {
	event := new(HanaConnectorHEVMHanaReverted)
	if err := _HanaConnectorHEVM.contract.UnpackLog(event, "HanaReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
