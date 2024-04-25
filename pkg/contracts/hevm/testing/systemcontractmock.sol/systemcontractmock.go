// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package systemcontractmock

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

// SystemContractMockMetaData contains all meta data concerning the SystemContractMock contract.
var SystemContractMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whana_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapv2Factory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapv2Router02_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerIsNotFungibleModule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CantBeIdenticalAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CantBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTarget\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SetGasCoin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SetGasHanaPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SetGasPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SetWHana\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SystemContractDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gasCoinHRC20ByChainId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gasHanaPoolByChainId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gasPriceByChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"onCrossChainCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"hrc20\",\"type\":\"address\"}],\"name\":\"setGasCoinHRC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setWHANAContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapv2FactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"uniswapv2PairFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapv2Router02Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wHanaContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620010d7380380620010d7833981810160405281019062000037919062000146565b82600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac560405160405180910390a1505050620001f5565b6000815190506200014081620001db565b92915050565b600080600060608486031215620001625762000161620001d6565b5b600062000172868287016200012f565b935050602062000185868287016200012f565b925050604062000198868287016200012f565b9150509250925092565b6000620001af82620001b6565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620001e681620001a2565b8114620001f257600080fd5b50565b610ed280620002056000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063c63585cc11610071578063c63585cc1461013c578063d7fd7afb1461016c578063d936a0121461019c578063d9e9b4c0146101ba578063db2131e5146101ea578063f197e0e114610208576100a9565b806306fd10b4146100ae5780631668a3f0146100ca5780633c669d55146100e6578063842da36d14610102578063a7cb050714610120575b600080fd5b6100c860048036038101906100c3919061097a565b610238565b005b6100e460048036038101906100df9190610818565b6102c7565b005b61010060048036038101906100fb9190610898565b610364565b005b61010a6104b1565b6040516101179190610bce565b60405180910390f35b61013a600480360381019061013591906109ba565b6104d7565b005b61015660048036038101906101519190610845565b61052b565b6040516101639190610bce565b60405180910390f35b6101866004803603810190610181919061094d565b61059d565b6040516101939190610c67565b60405180910390f35b6101a46105b5565b6040516101b19190610bce565b60405180910390f35b6101d460048036038101906101cf919061094d565b6105db565b6040516101e19190610bce565b60405180910390f35b6101f261060e565b6040516101ff9190610bce565b60405180910390f35b610222600480360381019061021d919061094d565b610634565b60405161022f9190610bce565b60405180910390f35b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d82826040516102bb929190610c82565b60405180910390a15050565b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f495b3e78b13285a35b3924be7ea2c0f6363b2e777d484e0e67274f778b63914b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516103599190610bce565b60405180910390a150565b600060405180606001604052806040518060200160405280600081525081526020013373ffffffffffffffffffffffffffffffffffffffff1681526020014681525090508473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb87866040518363ffffffff1660e01b81526004016103e3929190610be9565b602060405180830381600087803b1580156103fd57600080fd5b505af1158015610411573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104359190610920565b508573ffffffffffffffffffffffffffffffffffffffff1663de43156e82878787876040518663ffffffff1660e01b8152600401610477959493929190610c12565b600060405180830381600087803b15801561049157600080fd5b505af11580156104a5573d6000803e3d6000fd5b50505050505050505050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600080848152602001908152602001600020819055507f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d828260405161051f929190610cab565b60405180910390a15050565b600080600061053a8585610667565b91509150858282604051602001610552929190610b60565b60405160208183030381529060405280519060200120604051602001610579929190610b8c565b6040516020818303038152906040528051906020012060001c925050509392505050565b60006020528060005260406000206000915090505481565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156106d0576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161061070a57828461070d565b83835b8092508193505050600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561077c576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b60008135905061079281610e57565b92915050565b6000815190506107a781610e6e565b92915050565b60008083601f8401126107c3576107c2610dd3565b5b8235905067ffffffffffffffff8111156107e0576107df610dce565b5b6020830191508360018202830111156107fc576107fb610dd8565b5b9250929050565b60008135905061081281610e85565b92915050565b60006020828403121561082e5761082d610de2565b5b600061083c84828501610783565b91505092915050565b60008060006060848603121561085e5761085d610de2565b5b600061086c86828701610783565b935050602061087d86828701610783565b925050604061088e86828701610783565b9150509250925092565b6000806000806000608086880312156108b4576108b3610de2565b5b60006108c288828901610783565b95505060206108d388828901610783565b94505060406108e488828901610803565b935050606086013567ffffffffffffffff81111561090557610904610ddd565b5b610911888289016107ad565b92509250509295509295909350565b60006020828403121561093657610935610de2565b5b600061094484828501610798565b91505092915050565b60006020828403121561096357610962610de2565b5b600061097184828501610803565b91505092915050565b6000806040838503121561099157610990610de2565b5b600061099f85828601610803565b92505060206109b085828601610783565b9150509250929050565b600080604083850312156109d1576109d0610de2565b5b60006109df85828601610803565b92505060206109f085828601610803565b9150509250929050565b610a0381610d0c565b82525050565b610a1281610d0c565b82525050565b610a29610a2482610d0c565b610da0565b82525050565b610a40610a3b82610d2a565b610db2565b82525050565b6000610a528385610cf0565b9350610a5f838584610d5e565b610a6883610de7565b840190509392505050565b6000610a7e82610cd4565b610a888185610cdf565b9350610a98818560208601610d6d565b610aa181610de7565b840191505092915050565b6000610ab9602083610d01565b9150610ac482610e05565b602082019050919050565b6000610adc600183610d01565b9150610ae782610e2e565b600182019050919050565b60006060830160008301518482036000860152610b0f8282610a73565b9150506020830151610b2460208601826109fa565b506040830151610b376040860182610b42565b508091505092915050565b610b4b81610d54565b82525050565b610b5a81610d54565b82525050565b6000610b6c8285610a18565b601482019150610b7c8284610a18565b6014820191508190509392505050565b6000610b9782610acf565b9150610ba38285610a18565b601482019150610bb38284610a2f565b602082019150610bc282610aac565b91508190509392505050565b6000602082019050610be36000830184610a09565b92915050565b6000604082019050610bfe6000830185610a09565b610c0b6020830184610b51565b9392505050565b60006080820190508181036000830152610c2c8188610af2565b9050610c3b6020830187610a09565b610c486040830186610b51565b8181036060830152610c5b818486610a46565b90509695505050505050565b6000602082019050610c7c6000830184610b51565b92915050565b6000604082019050610c976000830185610b51565b610ca46020830184610a09565b9392505050565b6000604082019050610cc06000830185610b51565b610ccd6020830184610b51565b9392505050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000610d1782610d34565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610d8b578082015181840152602081019050610d70565b83811115610d9a576000848401525b50505050565b6000610dab82610dbc565b9050919050565b6000819050919050565b6000610dc782610df8565b9050919050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f600082015250565b7fff00000000000000000000000000000000000000000000000000000000000000600082015250565b610e6081610d0c565b8114610e6b57600080fd5b50565b610e7781610d1e565b8114610e8257600080fd5b50565b610e8e81610d54565b8114610e9957600080fd5b5056fea26469706673582212205a5d316e4d53d521f03ef4eee95c1c93fd5972253bb8776bd940a22249969c0c64736f6c63430008070033",
}

// SystemContractMockABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemContractMockMetaData.ABI instead.
var SystemContractMockABI = SystemContractMockMetaData.ABI

// SystemContractMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemContractMockMetaData.Bin instead.
var SystemContractMockBin = SystemContractMockMetaData.Bin

// DeploySystemContractMock deploys a new Ethereum contract, binding an instance of SystemContractMock to it.
func DeploySystemContractMock(auth *bind.TransactOpts, backend bind.ContractBackend, whana_ common.Address, uniswapv2Factory_ common.Address, uniswapv2Router02_ common.Address) (common.Address, *types.Transaction, *SystemContractMock, error) {
	parsed, err := SystemContractMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemContractMockBin), backend, whana_, uniswapv2Factory_, uniswapv2Router02_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemContractMock{SystemContractMockCaller: SystemContractMockCaller{contract: contract}, SystemContractMockTransactor: SystemContractMockTransactor{contract: contract}, SystemContractMockFilterer: SystemContractMockFilterer{contract: contract}}, nil
}

// SystemContractMock is an auto generated Go binding around an Ethereum contract.
type SystemContractMock struct {
	SystemContractMockCaller     // Read-only binding to the contract
	SystemContractMockTransactor // Write-only binding to the contract
	SystemContractMockFilterer   // Log filterer for contract events
}

// SystemContractMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemContractMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemContractMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemContractMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemContractMockSession struct {
	Contract     *SystemContractMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SystemContractMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemContractMockCallerSession struct {
	Contract *SystemContractMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SystemContractMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemContractMockTransactorSession struct {
	Contract     *SystemContractMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SystemContractMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemContractMockRaw struct {
	Contract *SystemContractMock // Generic contract binding to access the raw methods on
}

// SystemContractMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemContractMockCallerRaw struct {
	Contract *SystemContractMockCaller // Generic read-only contract binding to access the raw methods on
}

// SystemContractMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemContractMockTransactorRaw struct {
	Contract *SystemContractMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemContractMock creates a new instance of SystemContractMock, bound to a specific deployed contract.
func NewSystemContractMock(address common.Address, backend bind.ContractBackend) (*SystemContractMock, error) {
	contract, err := bindSystemContractMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemContractMock{SystemContractMockCaller: SystemContractMockCaller{contract: contract}, SystemContractMockTransactor: SystemContractMockTransactor{contract: contract}, SystemContractMockFilterer: SystemContractMockFilterer{contract: contract}}, nil
}

// NewSystemContractMockCaller creates a new read-only instance of SystemContractMock, bound to a specific deployed contract.
func NewSystemContractMockCaller(address common.Address, caller bind.ContractCaller) (*SystemContractMockCaller, error) {
	contract, err := bindSystemContractMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemContractMockCaller{contract: contract}, nil
}

// NewSystemContractMockTransactor creates a new write-only instance of SystemContractMock, bound to a specific deployed contract.
func NewSystemContractMockTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemContractMockTransactor, error) {
	contract, err := bindSystemContractMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemContractMockTransactor{contract: contract}, nil
}

// NewSystemContractMockFilterer creates a new log filterer instance of SystemContractMock, bound to a specific deployed contract.
func NewSystemContractMockFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemContractMockFilterer, error) {
	contract, err := bindSystemContractMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemContractMockFilterer{contract: contract}, nil
}

// bindSystemContractMock binds a generic wrapper to an already deployed contract.
func bindSystemContractMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SystemContractMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemContractMock *SystemContractMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemContractMock.Contract.SystemContractMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemContractMock *SystemContractMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SystemContractMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemContractMock *SystemContractMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SystemContractMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemContractMock *SystemContractMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemContractMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemContractMock *SystemContractMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContractMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemContractMock *SystemContractMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemContractMock.Contract.contract.Transact(opts, method, params...)
}

// GasCoinHRC20ByChainId is a free data retrieval call binding the contract method 0xf197e0e1.
//
// Solidity: function gasCoinHRC20ByChainId(uint256 ) view returns(address)
func (_SystemContractMock *SystemContractMockCaller) GasCoinHRC20ByChainId(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SystemContractMock.contract.Call(opts, &out, "gasCoinHRC20ByChainId", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasCoinHRC20ByChainId is a free data retrieval call binding the contract method 0xf197e0e1.
//
// Solidity: function gasCoinHRC20ByChainId(uint256 ) view returns(address)
func (_SystemContractMock *SystemContractMockSession) GasCoinHRC20ByChainId(arg0 *big.Int) (common.Address, error) {
	return _SystemContractMock.Contract.GasCoinHRC20ByChainId(&_SystemContractMock.CallOpts, arg0)
}

// GasCoinHRC20ByChainId is a free data retrieval call binding the contract method 0xf197e0e1.
//
// Solidity: function gasCoinHRC20ByChainId(uint256 ) view returns(address)
func (_SystemContractMock *SystemContractMockCallerSession) GasCoinHRC20ByChainId(arg0 *big.Int) (common.Address, error) {
	return _SystemContractMock.Contract.GasCoinHRC20ByChainId(&_SystemContractMock.CallOpts, arg0)
}

// GasHanaPoolByChainId is a free data retrieval call binding the contract method 0xd9e9b4c0.
//
// Solidity: function gasHanaPoolByChainId(uint256 ) view returns(address)
func (_SystemContractMock *SystemContractMockCaller) GasHanaPoolByChainId(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SystemContractMock.contract.Call(opts, &out, "gasHanaPoolByChainId", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasHanaPoolByChainId is a free data retrieval call binding the contract method 0xd9e9b4c0.
//
// Solidity: function gasHanaPoolByChainId(uint256 ) view returns(address)
func (_SystemContractMock *SystemContractMockSession) GasHanaPoolByChainId(arg0 *big.Int) (common.Address, error) {
	return _SystemContractMock.Contract.GasHanaPoolByChainId(&_SystemContractMock.CallOpts, arg0)
}

// GasHanaPoolByChainId is a free data retrieval call binding the contract method 0xd9e9b4c0.
//
// Solidity: function gasHanaPoolByChainId(uint256 ) view returns(address)
func (_SystemContractMock *SystemContractMockCallerSession) GasHanaPoolByChainId(arg0 *big.Int) (common.Address, error) {
	return _SystemContractMock.Contract.GasHanaPoolByChainId(&_SystemContractMock.CallOpts, arg0)
}

// GasPriceByChainId is a free data retrieval call binding the contract method 0xd7fd7afb.
//
// Solidity: function gasPriceByChainId(uint256 ) view returns(uint256)
func (_SystemContractMock *SystemContractMockCaller) GasPriceByChainId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SystemContractMock.contract.Call(opts, &out, "gasPriceByChainId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasPriceByChainId is a free data retrieval call binding the contract method 0xd7fd7afb.
//
// Solidity: function gasPriceByChainId(uint256 ) view returns(uint256)
func (_SystemContractMock *SystemContractMockSession) GasPriceByChainId(arg0 *big.Int) (*big.Int, error) {
	return _SystemContractMock.Contract.GasPriceByChainId(&_SystemContractMock.CallOpts, arg0)
}

// GasPriceByChainId is a free data retrieval call binding the contract method 0xd7fd7afb.
//
// Solidity: function gasPriceByChainId(uint256 ) view returns(uint256)
func (_SystemContractMock *SystemContractMockCallerSession) GasPriceByChainId(arg0 *big.Int) (*big.Int, error) {
	return _SystemContractMock.Contract.GasPriceByChainId(&_SystemContractMock.CallOpts, arg0)
}

// Uniswapv2FactoryAddress is a free data retrieval call binding the contract method 0xd936a012.
//
// Solidity: function uniswapv2FactoryAddress() view returns(address)
func (_SystemContractMock *SystemContractMockCaller) Uniswapv2FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContractMock.contract.Call(opts, &out, "uniswapv2FactoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uniswapv2FactoryAddress is a free data retrieval call binding the contract method 0xd936a012.
//
// Solidity: function uniswapv2FactoryAddress() view returns(address)
func (_SystemContractMock *SystemContractMockSession) Uniswapv2FactoryAddress() (common.Address, error) {
	return _SystemContractMock.Contract.Uniswapv2FactoryAddress(&_SystemContractMock.CallOpts)
}

// Uniswapv2FactoryAddress is a free data retrieval call binding the contract method 0xd936a012.
//
// Solidity: function uniswapv2FactoryAddress() view returns(address)
func (_SystemContractMock *SystemContractMockCallerSession) Uniswapv2FactoryAddress() (common.Address, error) {
	return _SystemContractMock.Contract.Uniswapv2FactoryAddress(&_SystemContractMock.CallOpts)
}

// Uniswapv2PairFor is a free data retrieval call binding the contract method 0xc63585cc.
//
// Solidity: function uniswapv2PairFor(address factory, address tokenA, address tokenB) pure returns(address pair)
func (_SystemContractMock *SystemContractMockCaller) Uniswapv2PairFor(opts *bind.CallOpts, factory common.Address, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _SystemContractMock.contract.Call(opts, &out, "uniswapv2PairFor", factory, tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uniswapv2PairFor is a free data retrieval call binding the contract method 0xc63585cc.
//
// Solidity: function uniswapv2PairFor(address factory, address tokenA, address tokenB) pure returns(address pair)
func (_SystemContractMock *SystemContractMockSession) Uniswapv2PairFor(factory common.Address, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _SystemContractMock.Contract.Uniswapv2PairFor(&_SystemContractMock.CallOpts, factory, tokenA, tokenB)
}

// Uniswapv2PairFor is a free data retrieval call binding the contract method 0xc63585cc.
//
// Solidity: function uniswapv2PairFor(address factory, address tokenA, address tokenB) pure returns(address pair)
func (_SystemContractMock *SystemContractMockCallerSession) Uniswapv2PairFor(factory common.Address, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _SystemContractMock.Contract.Uniswapv2PairFor(&_SystemContractMock.CallOpts, factory, tokenA, tokenB)
}

// Uniswapv2Router02Address is a free data retrieval call binding the contract method 0x842da36d.
//
// Solidity: function uniswapv2Router02Address() view returns(address)
func (_SystemContractMock *SystemContractMockCaller) Uniswapv2Router02Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContractMock.contract.Call(opts, &out, "uniswapv2Router02Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uniswapv2Router02Address is a free data retrieval call binding the contract method 0x842da36d.
//
// Solidity: function uniswapv2Router02Address() view returns(address)
func (_SystemContractMock *SystemContractMockSession) Uniswapv2Router02Address() (common.Address, error) {
	return _SystemContractMock.Contract.Uniswapv2Router02Address(&_SystemContractMock.CallOpts)
}

// Uniswapv2Router02Address is a free data retrieval call binding the contract method 0x842da36d.
//
// Solidity: function uniswapv2Router02Address() view returns(address)
func (_SystemContractMock *SystemContractMockCallerSession) Uniswapv2Router02Address() (common.Address, error) {
	return _SystemContractMock.Contract.Uniswapv2Router02Address(&_SystemContractMock.CallOpts)
}

// WHanaContractAddress is a free data retrieval call binding the contract method 0xdb2131e5.
//
// Solidity: function wHanaContractAddress() view returns(address)
func (_SystemContractMock *SystemContractMockCaller) WHanaContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContractMock.contract.Call(opts, &out, "wHanaContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WHanaContractAddress is a free data retrieval call binding the contract method 0xdb2131e5.
//
// Solidity: function wHanaContractAddress() view returns(address)
func (_SystemContractMock *SystemContractMockSession) WHanaContractAddress() (common.Address, error) {
	return _SystemContractMock.Contract.WHanaContractAddress(&_SystemContractMock.CallOpts)
}

// WHanaContractAddress is a free data retrieval call binding the contract method 0xdb2131e5.
//
// Solidity: function wHanaContractAddress() view returns(address)
func (_SystemContractMock *SystemContractMockCallerSession) WHanaContractAddress() (common.Address, error) {
	return _SystemContractMock.Contract.WHanaContractAddress(&_SystemContractMock.CallOpts)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0x3c669d55.
//
// Solidity: function onCrossChainCall(address target, address hrc20, uint256 amount, bytes message) returns()
func (_SystemContractMock *SystemContractMockTransactor) OnCrossChainCall(opts *bind.TransactOpts, target common.Address, hrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _SystemContractMock.contract.Transact(opts, "onCrossChainCall", target, hrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0x3c669d55.
//
// Solidity: function onCrossChainCall(address target, address hrc20, uint256 amount, bytes message) returns()
func (_SystemContractMock *SystemContractMockSession) OnCrossChainCall(target common.Address, hrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _SystemContractMock.Contract.OnCrossChainCall(&_SystemContractMock.TransactOpts, target, hrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0x3c669d55.
//
// Solidity: function onCrossChainCall(address target, address hrc20, uint256 amount, bytes message) returns()
func (_SystemContractMock *SystemContractMockTransactorSession) OnCrossChainCall(target common.Address, hrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _SystemContractMock.Contract.OnCrossChainCall(&_SystemContractMock.TransactOpts, target, hrc20, amount, message)
}

// SetGasCoinHRC20 is a paid mutator transaction binding the contract method 0x06fd10b4.
//
// Solidity: function setGasCoinHRC20(uint256 chainID, address hrc20) returns()
func (_SystemContractMock *SystemContractMockTransactor) SetGasCoinHRC20(opts *bind.TransactOpts, chainID *big.Int, hrc20 common.Address) (*types.Transaction, error) {
	return _SystemContractMock.contract.Transact(opts, "setGasCoinHRC20", chainID, hrc20)
}

// SetGasCoinHRC20 is a paid mutator transaction binding the contract method 0x06fd10b4.
//
// Solidity: function setGasCoinHRC20(uint256 chainID, address hrc20) returns()
func (_SystemContractMock *SystemContractMockSession) SetGasCoinHRC20(chainID *big.Int, hrc20 common.Address) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SetGasCoinHRC20(&_SystemContractMock.TransactOpts, chainID, hrc20)
}

// SetGasCoinHRC20 is a paid mutator transaction binding the contract method 0x06fd10b4.
//
// Solidity: function setGasCoinHRC20(uint256 chainID, address hrc20) returns()
func (_SystemContractMock *SystemContractMockTransactorSession) SetGasCoinHRC20(chainID *big.Int, hrc20 common.Address) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SetGasCoinHRC20(&_SystemContractMock.TransactOpts, chainID, hrc20)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xa7cb0507.
//
// Solidity: function setGasPrice(uint256 chainID, uint256 price) returns()
func (_SystemContractMock *SystemContractMockTransactor) SetGasPrice(opts *bind.TransactOpts, chainID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _SystemContractMock.contract.Transact(opts, "setGasPrice", chainID, price)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xa7cb0507.
//
// Solidity: function setGasPrice(uint256 chainID, uint256 price) returns()
func (_SystemContractMock *SystemContractMockSession) SetGasPrice(chainID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SetGasPrice(&_SystemContractMock.TransactOpts, chainID, price)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xa7cb0507.
//
// Solidity: function setGasPrice(uint256 chainID, uint256 price) returns()
func (_SystemContractMock *SystemContractMockTransactorSession) SetGasPrice(chainID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SetGasPrice(&_SystemContractMock.TransactOpts, chainID, price)
}

// SetWHANAContractAddress is a paid mutator transaction binding the contract method 0x1668a3f0.
//
// Solidity: function setWHANAContractAddress(address addr) returns()
func (_SystemContractMock *SystemContractMockTransactor) SetWHANAContractAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _SystemContractMock.contract.Transact(opts, "setWHANAContractAddress", addr)
}

// SetWHANAContractAddress is a paid mutator transaction binding the contract method 0x1668a3f0.
//
// Solidity: function setWHANAContractAddress(address addr) returns()
func (_SystemContractMock *SystemContractMockSession) SetWHANAContractAddress(addr common.Address) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SetWHANAContractAddress(&_SystemContractMock.TransactOpts, addr)
}

// SetWHANAContractAddress is a paid mutator transaction binding the contract method 0x1668a3f0.
//
// Solidity: function setWHANAContractAddress(address addr) returns()
func (_SystemContractMock *SystemContractMockTransactorSession) SetWHANAContractAddress(addr common.Address) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SetWHANAContractAddress(&_SystemContractMock.TransactOpts, addr)
}

// SystemContractMockSetGasCoinIterator is returned from FilterSetGasCoin and is used to iterate over the raw logs and unpacked data for SetGasCoin events raised by the SystemContractMock contract.
type SystemContractMockSetGasCoinIterator struct {
	Event *SystemContractMockSetGasCoin // Event containing the contract specifics and raw log

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
func (it *SystemContractMockSetGasCoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractMockSetGasCoin)
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
		it.Event = new(SystemContractMockSetGasCoin)
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
func (it *SystemContractMockSetGasCoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractMockSetGasCoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractMockSetGasCoin represents a SetGasCoin event raised by the SystemContractMock contract.
type SystemContractMockSetGasCoin struct {
	Arg0 *big.Int
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetGasCoin is a free log retrieval operation binding the contract event 0xd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d.
//
// Solidity: event SetGasCoin(uint256 arg0, address arg1)
func (_SystemContractMock *SystemContractMockFilterer) FilterSetGasCoin(opts *bind.FilterOpts) (*SystemContractMockSetGasCoinIterator, error) {

	logs, sub, err := _SystemContractMock.contract.FilterLogs(opts, "SetGasCoin")
	if err != nil {
		return nil, err
	}
	return &SystemContractMockSetGasCoinIterator{contract: _SystemContractMock.contract, event: "SetGasCoin", logs: logs, sub: sub}, nil
}

// WatchSetGasCoin is a free log subscription operation binding the contract event 0xd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d.
//
// Solidity: event SetGasCoin(uint256 arg0, address arg1)
func (_SystemContractMock *SystemContractMockFilterer) WatchSetGasCoin(opts *bind.WatchOpts, sink chan<- *SystemContractMockSetGasCoin) (event.Subscription, error) {

	logs, sub, err := _SystemContractMock.contract.WatchLogs(opts, "SetGasCoin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractMockSetGasCoin)
				if err := _SystemContractMock.contract.UnpackLog(event, "SetGasCoin", log); err != nil {
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

// ParseSetGasCoin is a log parse operation binding the contract event 0xd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d.
//
// Solidity: event SetGasCoin(uint256 arg0, address arg1)
func (_SystemContractMock *SystemContractMockFilterer) ParseSetGasCoin(log types.Log) (*SystemContractMockSetGasCoin, error) {
	event := new(SystemContractMockSetGasCoin)
	if err := _SystemContractMock.contract.UnpackLog(event, "SetGasCoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractMockSetGasHanaPoolIterator is returned from FilterSetGasHanaPool and is used to iterate over the raw logs and unpacked data for SetGasHanaPool events raised by the SystemContractMock contract.
type SystemContractMockSetGasHanaPoolIterator struct {
	Event *SystemContractMockSetGasHanaPool // Event containing the contract specifics and raw log

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
func (it *SystemContractMockSetGasHanaPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractMockSetGasHanaPool)
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
		it.Event = new(SystemContractMockSetGasHanaPool)
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
func (it *SystemContractMockSetGasHanaPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractMockSetGasHanaPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractMockSetGasHanaPool represents a SetGasHanaPool event raised by the SystemContractMock contract.
type SystemContractMockSetGasHanaPool struct {
	Arg0 *big.Int
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetGasHanaPool is a free log retrieval operation binding the contract event 0xfde6856eeb152d7c29aff8485d17f698215cf3de7532688f9c59ecedfd2a77c7.
//
// Solidity: event SetGasHanaPool(uint256 arg0, address arg1)
func (_SystemContractMock *SystemContractMockFilterer) FilterSetGasHanaPool(opts *bind.FilterOpts) (*SystemContractMockSetGasHanaPoolIterator, error) {

	logs, sub, err := _SystemContractMock.contract.FilterLogs(opts, "SetGasHanaPool")
	if err != nil {
		return nil, err
	}
	return &SystemContractMockSetGasHanaPoolIterator{contract: _SystemContractMock.contract, event: "SetGasHanaPool", logs: logs, sub: sub}, nil
}

// WatchSetGasHanaPool is a free log subscription operation binding the contract event 0xfde6856eeb152d7c29aff8485d17f698215cf3de7532688f9c59ecedfd2a77c7.
//
// Solidity: event SetGasHanaPool(uint256 arg0, address arg1)
func (_SystemContractMock *SystemContractMockFilterer) WatchSetGasHanaPool(opts *bind.WatchOpts, sink chan<- *SystemContractMockSetGasHanaPool) (event.Subscription, error) {

	logs, sub, err := _SystemContractMock.contract.WatchLogs(opts, "SetGasHanaPool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractMockSetGasHanaPool)
				if err := _SystemContractMock.contract.UnpackLog(event, "SetGasHanaPool", log); err != nil {
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

// ParseSetGasHanaPool is a log parse operation binding the contract event 0xfde6856eeb152d7c29aff8485d17f698215cf3de7532688f9c59ecedfd2a77c7.
//
// Solidity: event SetGasHanaPool(uint256 arg0, address arg1)
func (_SystemContractMock *SystemContractMockFilterer) ParseSetGasHanaPool(log types.Log) (*SystemContractMockSetGasHanaPool, error) {
	event := new(SystemContractMockSetGasHanaPool)
	if err := _SystemContractMock.contract.UnpackLog(event, "SetGasHanaPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractMockSetGasPriceIterator is returned from FilterSetGasPrice and is used to iterate over the raw logs and unpacked data for SetGasPrice events raised by the SystemContractMock contract.
type SystemContractMockSetGasPriceIterator struct {
	Event *SystemContractMockSetGasPrice // Event containing the contract specifics and raw log

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
func (it *SystemContractMockSetGasPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractMockSetGasPrice)
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
		it.Event = new(SystemContractMockSetGasPrice)
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
func (it *SystemContractMockSetGasPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractMockSetGasPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractMockSetGasPrice represents a SetGasPrice event raised by the SystemContractMock contract.
type SystemContractMockSetGasPrice struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetGasPrice is a free log retrieval operation binding the contract event 0x49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d.
//
// Solidity: event SetGasPrice(uint256 arg0, uint256 arg1)
func (_SystemContractMock *SystemContractMockFilterer) FilterSetGasPrice(opts *bind.FilterOpts) (*SystemContractMockSetGasPriceIterator, error) {

	logs, sub, err := _SystemContractMock.contract.FilterLogs(opts, "SetGasPrice")
	if err != nil {
		return nil, err
	}
	return &SystemContractMockSetGasPriceIterator{contract: _SystemContractMock.contract, event: "SetGasPrice", logs: logs, sub: sub}, nil
}

// WatchSetGasPrice is a free log subscription operation binding the contract event 0x49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d.
//
// Solidity: event SetGasPrice(uint256 arg0, uint256 arg1)
func (_SystemContractMock *SystemContractMockFilterer) WatchSetGasPrice(opts *bind.WatchOpts, sink chan<- *SystemContractMockSetGasPrice) (event.Subscription, error) {

	logs, sub, err := _SystemContractMock.contract.WatchLogs(opts, "SetGasPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractMockSetGasPrice)
				if err := _SystemContractMock.contract.UnpackLog(event, "SetGasPrice", log); err != nil {
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

// ParseSetGasPrice is a log parse operation binding the contract event 0x49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d.
//
// Solidity: event SetGasPrice(uint256 arg0, uint256 arg1)
func (_SystemContractMock *SystemContractMockFilterer) ParseSetGasPrice(log types.Log) (*SystemContractMockSetGasPrice, error) {
	event := new(SystemContractMockSetGasPrice)
	if err := _SystemContractMock.contract.UnpackLog(event, "SetGasPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractMockSetWHanaIterator is returned from FilterSetWHana and is used to iterate over the raw logs and unpacked data for SetWHana events raised by the SystemContractMock contract.
type SystemContractMockSetWHanaIterator struct {
	Event *SystemContractMockSetWHana // Event containing the contract specifics and raw log

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
func (it *SystemContractMockSetWHanaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractMockSetWHana)
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
		it.Event = new(SystemContractMockSetWHana)
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
func (it *SystemContractMockSetWHanaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractMockSetWHanaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractMockSetWHana represents a SetWHana event raised by the SystemContractMock contract.
type SystemContractMockSetWHana struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetWHana is a free log retrieval operation binding the contract event 0x495b3e78b13285a35b3924be7ea2c0f6363b2e777d484e0e67274f778b63914b.
//
// Solidity: event SetWHana(address arg0)
func (_SystemContractMock *SystemContractMockFilterer) FilterSetWHana(opts *bind.FilterOpts) (*SystemContractMockSetWHanaIterator, error) {

	logs, sub, err := _SystemContractMock.contract.FilterLogs(opts, "SetWHana")
	if err != nil {
		return nil, err
	}
	return &SystemContractMockSetWHanaIterator{contract: _SystemContractMock.contract, event: "SetWHana", logs: logs, sub: sub}, nil
}

// WatchSetWHana is a free log subscription operation binding the contract event 0x495b3e78b13285a35b3924be7ea2c0f6363b2e777d484e0e67274f778b63914b.
//
// Solidity: event SetWHana(address arg0)
func (_SystemContractMock *SystemContractMockFilterer) WatchSetWHana(opts *bind.WatchOpts, sink chan<- *SystemContractMockSetWHana) (event.Subscription, error) {

	logs, sub, err := _SystemContractMock.contract.WatchLogs(opts, "SetWHana")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractMockSetWHana)
				if err := _SystemContractMock.contract.UnpackLog(event, "SetWHana", log); err != nil {
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

// ParseSetWHana is a log parse operation binding the contract event 0x495b3e78b13285a35b3924be7ea2c0f6363b2e777d484e0e67274f778b63914b.
//
// Solidity: event SetWHana(address arg0)
func (_SystemContractMock *SystemContractMockFilterer) ParseSetWHana(log types.Log) (*SystemContractMockSetWHana, error) {
	event := new(SystemContractMockSetWHana)
	if err := _SystemContractMock.contract.UnpackLog(event, "SetWHana", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractMockSystemContractDeployedIterator is returned from FilterSystemContractDeployed and is used to iterate over the raw logs and unpacked data for SystemContractDeployed events raised by the SystemContractMock contract.
type SystemContractMockSystemContractDeployedIterator struct {
	Event *SystemContractMockSystemContractDeployed // Event containing the contract specifics and raw log

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
func (it *SystemContractMockSystemContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractMockSystemContractDeployed)
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
		it.Event = new(SystemContractMockSystemContractDeployed)
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
func (it *SystemContractMockSystemContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractMockSystemContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractMockSystemContractDeployed represents a SystemContractDeployed event raised by the SystemContractMock contract.
type SystemContractMockSystemContractDeployed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSystemContractDeployed is a free log retrieval operation binding the contract event 0x80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac5.
//
// Solidity: event SystemContractDeployed()
func (_SystemContractMock *SystemContractMockFilterer) FilterSystemContractDeployed(opts *bind.FilterOpts) (*SystemContractMockSystemContractDeployedIterator, error) {

	logs, sub, err := _SystemContractMock.contract.FilterLogs(opts, "SystemContractDeployed")
	if err != nil {
		return nil, err
	}
	return &SystemContractMockSystemContractDeployedIterator{contract: _SystemContractMock.contract, event: "SystemContractDeployed", logs: logs, sub: sub}, nil
}

// WatchSystemContractDeployed is a free log subscription operation binding the contract event 0x80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac5.
//
// Solidity: event SystemContractDeployed()
func (_SystemContractMock *SystemContractMockFilterer) WatchSystemContractDeployed(opts *bind.WatchOpts, sink chan<- *SystemContractMockSystemContractDeployed) (event.Subscription, error) {

	logs, sub, err := _SystemContractMock.contract.WatchLogs(opts, "SystemContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractMockSystemContractDeployed)
				if err := _SystemContractMock.contract.UnpackLog(event, "SystemContractDeployed", log); err != nil {
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

// ParseSystemContractDeployed is a log parse operation binding the contract event 0x80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac5.
//
// Solidity: event SystemContractDeployed()
func (_SystemContractMock *SystemContractMockFilterer) ParseSystemContractDeployed(log types.Log) (*SystemContractMockSystemContractDeployed, error) {
	event := new(SystemContractMockSystemContractDeployed)
	if err := _SystemContractMock.contract.UnpackLog(event, "SystemContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
