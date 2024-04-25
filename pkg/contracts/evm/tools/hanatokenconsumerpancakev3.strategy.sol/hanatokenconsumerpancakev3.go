// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanatokenconsumerpancakev3

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

// HanaTokenConsumerPancakeV3MetaData contains all meta data concerning the HanaTokenConsumerPancakeV3 contract.
var HanaTokenConsumerPancakeV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hanaToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pancakeV3Router_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV3Factory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WETH9Address_\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"hanaPoolFee_\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"tokenPoolFee_\",\"type\":\"uint24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrorSendingETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthExchangedForHana\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"HanaExchangedForEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"HanaExchangedForToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenExchangedForHana\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH9Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hanaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getEthFromHana\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"getHanaFromEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getHanaFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hanaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenFromHana\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hanaPoolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hanaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasHanaLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pancakeV3Router\",\"outputs\":[{\"internalType\":\"contractISwapRouterPancake\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPoolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV3Factory\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b506040516200288c3803806200288c83398181016040528101906200003891906200028a565b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161480620000a05750600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16145b80620000d85750600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16145b80620001105750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b1562000148576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508473ffffffffffffffffffffffffffffffffffffffff166101008173ffffffffffffffffffffffffffffffffffffffff1660601b815250508373ffffffffffffffffffffffffffffffffffffffff166101208173ffffffffffffffffffffffffffffffffffffffff1660601b815250508273ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508162ffffff1660808162ffffff1660e81b815250508062ffffff1660a08162ffffff1660e81b81525050505050505050620003a2565b6000815190506200026d816200036e565b92915050565b600081519050620002848162000388565b92915050565b60008060008060008060c08789031215620002aa57620002a962000369565b5b6000620002ba89828a016200025c565b9650506020620002cd89828a016200025c565b9550506040620002e089828a016200025c565b9450506060620002f389828a016200025c565b93505060806200030689828a0162000273565b92505060a06200031989828a0162000273565b9150509295509295509295565b600062000333826200033a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062ffffff82169050919050565b600080fd5b620003798162000326565b81146200038557600080fd5b50565b62000393816200035a565b81146200039f57600080fd5b50565b60805160e81c60a05160e81c60c05160601c60e05160601c6101005160601c6101205160601c6123b6620004d6600039600081816102e90152610bd60152600081816105ca0152818161072901528181610ae401528181610d8501528181610ed1015281816111320152818161125e015261135301526000818161034601528181610580015281816105ec0152818161063f01528181610a3601528181610d3b01528181610da701528181610dfa0152818161100401526111eb0152600081816103250152818161067b015281816107d6015281816109fa01528181610e3c015281816111a90152611377015260008181610e5d01528181610fe001526111880152600081816102c201528181610367015281816106b701528181610a7201528181610e1b01526111ca01526123b66000f3fe6080604052600436106100a05760003560e01c80635c2fec9a116100645780635c2fec9a1461019a5780635d9dfdde146101d75780635e694a921461020257806365d6edc81461022d578063c27745dd1461026a578063c469cf1414610295576100a7565b8063207c1044146100ac578063246d567e146100d7578063291d55f71461010257806345413df71461013f5780635b5491821461016f576100a7565b366100a757005b600080fd5b3480156100b857600080fd5b506100c16102c0565b6040516100ce9190612015565b60405180910390f35b3480156100e357600080fd5b506100ec6102e4565b6040516100f99190611ee5565b60405180910390f35b34801561010e57600080fd5b5061012960048036038101906101249190611967565b6104d5565b6040516101369190612030565b60405180910390f35b610159600480360381019061015491906118c0565b610948565b6040516101669190612030565b60405180910390f35b34801561017b57600080fd5b50610184610bd4565b6040516101919190611f1b565b60405180910390f35b3480156101a657600080fd5b506101c160048036038101906101bc9190611900565b610bf8565b6040516101ce9190612030565b60405180910390f35b3480156101e357600080fd5b506101ec610fde565b6040516101f99190612015565b60405180910390f35b34801561020e57600080fd5b50610217611002565b6040516102249190611dd3565b60405180910390f35b34801561023957600080fd5b50610254600480360381019061024f9190611900565b611026565b6040516102619190612030565b60405180910390f35b34801561027657600080fd5b5061027f611351565b60405161028c9190611f00565b60405180910390f35b3480156102a157600080fd5b506102aa611375565b6040516102b79190611dd3565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b6000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631698ee827f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006040518463ffffffff1660e01b81526004016103a493929190611e17565b60206040518083038186803b1580156103bc57600080fd5b505afa1580156103d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f49190611893565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156104355760009150506104d2565b600081905060008173ffffffffffffffffffffffffffffffffffffffff16631a6865026040518163ffffffff1660e01b815260040160206040518083038186803b15801561048257600080fd5b505afa158015610496573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ba91906119e7565b6fffffffffffffffffffffffffffffffff1611925050505b90565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561053d576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610578576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105c53330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611399909392919063ffffffff16565b6106307f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166114229092919063ffffffff16565b60006040518060e001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000062ffffff1681526020013073ffffffffffffffffffffffffffffffffffffffff168152602001848152602001858152602001600073ffffffffffffffffffffffffffffffffffffffff16815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166304e45aaf836040518263ffffffff1660e01b81526004016107809190611ffa565b602060405180830381600087803b15801561079a57600080fd5b505af11580156107ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d29190611a14565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d826040518263ffffffff1660e01b815260040161082d9190612030565b600060405180830381600087803b15801561084757600080fd5b505af115801561085b573d6000803e3d6000fd5b505050507fd94ef7f26aa4584e0bb757c8b45e0663569b29a3a8a0fc85ed49a8bb95ed306a848260405161089092919061204b565b60405180910390a160008673ffffffffffffffffffffffffffffffffffffffff16826040516108be90611dbe565b60006040518083038185875af1925050503d80600081146108fb576040519150601f19603f3d011682016040523d82523d6000602084013e610900565b606091505b505090508061093b576040517f3794aeaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8193505050509392505050565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156109b0576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60003414156109eb576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006040518060e001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000062ffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001348152602001848152602001600073ffffffffffffffffffffffffffffffffffffffff16815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166304e45aaf34846040518363ffffffff1660e01b8152600401610b3c9190611ffa565b6020604051808303818588803b158015610b5557600080fd5b505af1158015610b69573d6000803e3d6000fd5b50505050506040513d601f19601f82011682018060405250810190610b8e9190611a14565b90507f877de62ce9cb129480429d4e4ac5b3ce3b8b85638de13f817d02b9cea87a0f363482604051610bc192919061204b565b60405180910390a1809250505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008060009054906101000a900460ff1615610c40576040517f29f745a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016000806101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161480610cc15750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610cf8576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610d33576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d803330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611399909392919063ffffffff16565b610deb7f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166114229092919063ffffffff16565b600060405180608001604052807f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000089604051602001610e91959493929190611d48565b60405160208183030381529060405281526020018773ffffffffffffffffffffffffffffffffffffffff16815260200184815260200186815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b858183f836040518263ffffffff1660e01b8152600401610f289190611fd8565b602060405180830381600087803b158015610f4257600080fd5b505af1158015610f56573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f7a9190611a14565b90507f532789829182476cbd3cc6523b623cc2504e5a0b348dfdc5527baf726e6d55af858583604051610faf93929190611eae565b60405180910390a1809250505060008060006101000a81548160ff021916908315150217905550949350505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16148061108e5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b156110c5576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415611100576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61112d3330848673ffffffffffffffffffffffffffffffffffffffff16611399909392919063ffffffff16565b6111787f0000000000000000000000000000000000000000000000000000000000000000838573ffffffffffffffffffffffffffffffffffffffff166114229092919063ffffffff16565b60006040518060800160405280857f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060405160200161121e959493929190611d48565b60405160208183030381529060405281526020018773ffffffffffffffffffffffffffffffffffffffff16815260200184815260200186815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b858183f836040518263ffffffff1660e01b81526004016112b59190611fd8565b602060405180830381600087803b1580156112cf57600080fd5b505af11580156112e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113079190611a14565b90507fc7f70cbba54984322735d319a149c7bccec0ab310897fbc85b00f59d4486ce6085858360405161133c93929190611eae565b60405180910390a18092505050949350505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b61141c846323b872dd60e01b8585856040516024016113ba93929190611e4e565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611580565b50505050565b60008114806114bb575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b8152600401611469929190611dee565b60206040518083038186803b15801561148157600080fd5b505afa158015611495573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114b99190611a14565b145b6114fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114f190611fb8565b60405180910390fd5b61157b8363095ea7b360e01b8484604051602401611519929190611e85565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611580565b505050565b60006115e2826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166116479092919063ffffffff16565b9050600081511115611642578080602001905181019061160291906119ba565b611641576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161163890611f98565b60405180910390fd5b5b505050565b6060611656848460008561165f565b90509392505050565b6060824710156116a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161169b90611f58565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516116cd9190611da7565b60006040518083038185875af1925050503d806000811461170a576040519150601f19603f3d011682016040523d82523d6000602084013e61170f565b606091505b50915091506117208783838761172c565b92505050949350505050565b6060831561178f5760008351141561178757611747856117a2565b611786576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161177d90611f78565b60405180910390fd5b5b82905061179a565b61179983836117c5565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156117d85781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161180c9190611f36565b60405180910390fd5b60008135905061182481612324565b92915050565b60008151905061183981612324565b92915050565b60008151905061184e8161233b565b92915050565b60008151905061186381612352565b92915050565b60008135905061187881612369565b92915050565b60008151905061188d81612369565b92915050565b6000602082840312156118a9576118a86121db565b5b60006118b78482850161182a565b91505092915050565b600080604083850312156118d7576118d66121db565b5b60006118e585828601611815565b92505060206118f685828601611869565b9150509250929050565b6000806000806080858703121561191a576119196121db565b5b600061192887828801611815565b945050602061193987828801611869565b935050604061194a87828801611815565b925050606061195b87828801611869565b91505092959194509250565b6000806000606084860312156119805761197f6121db565b5b600061198e86828701611815565b935050602061199f86828701611869565b92505060406119b086828701611869565b9150509250925092565b6000602082840312156119d0576119cf6121db565b5b60006119de8482850161183f565b91505092915050565b6000602082840312156119fd576119fc6121db565b5b6000611a0b84828501611854565b91505092915050565b600060208284031215611a2a57611a296121db565b5b6000611a388482850161187e565b91505092915050565b611a4a816120b7565b82525050565b611a59816120b7565b82525050565b611a70611a6b826120b7565b6121a5565b82525050565b611a7f816120c9565b82525050565b6000611a9082612074565b611a9a818561208a565b9350611aaa818560208601612172565b611ab3816121e0565b840191505092915050565b6000611ac982612074565b611ad3818561209b565b9350611ae3818560208601612172565b80840191505092915050565b611af88161212a565b82525050565b611b078161213c565b82525050565b6000611b188261207f565b611b2281856120a6565b9350611b32818560208601612172565b611b3b816121e0565b840191505092915050565b6000611b536026836120a6565b9150611b5e8261220b565b604082019050919050565b6000611b7660008361209b565b9150611b818261225a565b600082019050919050565b6000611b99601d836120a6565b9150611ba48261225d565b602082019050919050565b6000611bbc602a836120a6565b9150611bc782612286565b604082019050919050565b6000611bdf6036836120a6565b9150611bea826122d5565b604082019050919050565b60006080830160008301518482036000860152611c128282611a85565b9150506020830151611c276020860182611a41565b506040830151611c3a6040860182611d2a565b506060830151611c4d6060860182611d2a565b508091505092915050565b60e082016000820151611c6e6000850182611a41565b506020820151611c816020850182611a41565b506040820151611c946040850182611cf5565b506060820151611ca76060850182611a41565b506080820151611cba6080850182611d2a565b5060a0820151611ccd60a0850182611d2a565b5060c0820151611ce060c0850182611ce6565b50505050565b611cef816120f1565b82525050565b611cfe81612111565b82525050565b611d0d81612111565b82525050565b611d24611d1f82612111565b6121c9565b82525050565b611d3381612120565b82525050565b611d4281612120565b82525050565b6000611d548288611a5f565b601482019150611d648287611d13565b600382019150611d748286611a5f565b601482019150611d848285611d13565b600382019150611d948284611a5f565b6014820191508190509695505050505050565b6000611db38284611abe565b915081905092915050565b6000611dc982611b69565b9150819050919050565b6000602082019050611de86000830184611a50565b92915050565b6000604082019050611e036000830185611a50565b611e106020830184611a50565b9392505050565b6000606082019050611e2c6000830186611a50565b611e396020830185611a50565b611e466040830184611d04565b949350505050565b6000606082019050611e636000830186611a50565b611e706020830185611a50565b611e7d6040830184611d39565b949350505050565b6000604082019050611e9a6000830185611a50565b611ea76020830184611d39565b9392505050565b6000606082019050611ec36000830186611a50565b611ed06020830185611d39565b611edd6040830184611d39565b949350505050565b6000602082019050611efa6000830184611a76565b92915050565b6000602082019050611f156000830184611aef565b92915050565b6000602082019050611f306000830184611afe565b92915050565b60006020820190508181036000830152611f508184611b0d565b905092915050565b60006020820190508181036000830152611f7181611b46565b9050919050565b60006020820190508181036000830152611f9181611b8c565b9050919050565b60006020820190508181036000830152611fb181611baf565b9050919050565b60006020820190508181036000830152611fd181611bd2565b9050919050565b60006020820190508181036000830152611ff28184611bf5565b905092915050565b600060e08201905061200f6000830184611c58565b92915050565b600060208201905061202a6000830184611d04565b92915050565b60006020820190506120456000830184611d39565b92915050565b60006040820190506120606000830185611d39565b61206d6020830184611d39565b9392505050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006120c2826120f1565b9050919050565b60008115159050919050565b60006fffffffffffffffffffffffffffffffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062ffffff82169050919050565b6000819050919050565b60006121358261214e565b9050919050565b60006121478261214e565b9050919050565b600061215982612160565b9050919050565b600061216b826120f1565b9050919050565b60005b83811015612190578082015181840152602081019050612175565b8381111561219f576000848401525b50505050565b60006121b0826121b7565b9050919050565b60006121c2826121fe565b9050919050565b60006121d4826121f1565b9050919050565b600080fd5b6000601f19601f8301169050919050565b60008160e81b9050919050565b60008160601b9050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60008201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000602082015250565b61232d816120b7565b811461233857600080fd5b50565b612344816120c9565b811461234f57600080fd5b50565b61235b816120d5565b811461236657600080fd5b50565b61237281612120565b811461237d57600080fd5b5056fea2646970667358221220e4efdcb4eaf7f4f87bdfd1ad667fc0b5373f826a8d5c7122b20273aa6d5dd74264736f6c63430008070033",
}

// HanaTokenConsumerPancakeV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaTokenConsumerPancakeV3MetaData.ABI instead.
var HanaTokenConsumerPancakeV3ABI = HanaTokenConsumerPancakeV3MetaData.ABI

// HanaTokenConsumerPancakeV3Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HanaTokenConsumerPancakeV3MetaData.Bin instead.
var HanaTokenConsumerPancakeV3Bin = HanaTokenConsumerPancakeV3MetaData.Bin

// DeployHanaTokenConsumerPancakeV3 deploys a new Ethereum contract, binding an instance of HanaTokenConsumerPancakeV3 to it.
func DeployHanaTokenConsumerPancakeV3(auth *bind.TransactOpts, backend bind.ContractBackend, hanaToken_ common.Address, pancakeV3Router_ common.Address, uniswapV3Factory_ common.Address, WETH9Address_ common.Address, hanaPoolFee_ *big.Int, tokenPoolFee_ *big.Int) (common.Address, *types.Transaction, *HanaTokenConsumerPancakeV3, error) {
	parsed, err := HanaTokenConsumerPancakeV3MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HanaTokenConsumerPancakeV3Bin), backend, hanaToken_, pancakeV3Router_, uniswapV3Factory_, WETH9Address_, hanaPoolFee_, tokenPoolFee_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HanaTokenConsumerPancakeV3{HanaTokenConsumerPancakeV3Caller: HanaTokenConsumerPancakeV3Caller{contract: contract}, HanaTokenConsumerPancakeV3Transactor: HanaTokenConsumerPancakeV3Transactor{contract: contract}, HanaTokenConsumerPancakeV3Filterer: HanaTokenConsumerPancakeV3Filterer{contract: contract}}, nil
}

// HanaTokenConsumerPancakeV3 is an auto generated Go binding around an Ethereum contract.
type HanaTokenConsumerPancakeV3 struct {
	HanaTokenConsumerPancakeV3Caller     // Read-only binding to the contract
	HanaTokenConsumerPancakeV3Transactor // Write-only binding to the contract
	HanaTokenConsumerPancakeV3Filterer   // Log filterer for contract events
}

// HanaTokenConsumerPancakeV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type HanaTokenConsumerPancakeV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerPancakeV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaTokenConsumerPancakeV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerPancakeV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaTokenConsumerPancakeV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerPancakeV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaTokenConsumerPancakeV3Session struct {
	Contract     *HanaTokenConsumerPancakeV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// HanaTokenConsumerPancakeV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaTokenConsumerPancakeV3CallerSession struct {
	Contract *HanaTokenConsumerPancakeV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// HanaTokenConsumerPancakeV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaTokenConsumerPancakeV3TransactorSession struct {
	Contract     *HanaTokenConsumerPancakeV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// HanaTokenConsumerPancakeV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type HanaTokenConsumerPancakeV3Raw struct {
	Contract *HanaTokenConsumerPancakeV3 // Generic contract binding to access the raw methods on
}

// HanaTokenConsumerPancakeV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaTokenConsumerPancakeV3CallerRaw struct {
	Contract *HanaTokenConsumerPancakeV3Caller // Generic read-only contract binding to access the raw methods on
}

// HanaTokenConsumerPancakeV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaTokenConsumerPancakeV3TransactorRaw struct {
	Contract *HanaTokenConsumerPancakeV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaTokenConsumerPancakeV3 creates a new instance of HanaTokenConsumerPancakeV3, bound to a specific deployed contract.
func NewHanaTokenConsumerPancakeV3(address common.Address, backend bind.ContractBackend) (*HanaTokenConsumerPancakeV3, error) {
	contract, err := bindHanaTokenConsumerPancakeV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerPancakeV3{HanaTokenConsumerPancakeV3Caller: HanaTokenConsumerPancakeV3Caller{contract: contract}, HanaTokenConsumerPancakeV3Transactor: HanaTokenConsumerPancakeV3Transactor{contract: contract}, HanaTokenConsumerPancakeV3Filterer: HanaTokenConsumerPancakeV3Filterer{contract: contract}}, nil
}

// NewHanaTokenConsumerPancakeV3Caller creates a new read-only instance of HanaTokenConsumerPancakeV3, bound to a specific deployed contract.
func NewHanaTokenConsumerPancakeV3Caller(address common.Address, caller bind.ContractCaller) (*HanaTokenConsumerPancakeV3Caller, error) {
	contract, err := bindHanaTokenConsumerPancakeV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerPancakeV3Caller{contract: contract}, nil
}

// NewHanaTokenConsumerPancakeV3Transactor creates a new write-only instance of HanaTokenConsumerPancakeV3, bound to a specific deployed contract.
func NewHanaTokenConsumerPancakeV3Transactor(address common.Address, transactor bind.ContractTransactor) (*HanaTokenConsumerPancakeV3Transactor, error) {
	contract, err := bindHanaTokenConsumerPancakeV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerPancakeV3Transactor{contract: contract}, nil
}

// NewHanaTokenConsumerPancakeV3Filterer creates a new log filterer instance of HanaTokenConsumerPancakeV3, bound to a specific deployed contract.
func NewHanaTokenConsumerPancakeV3Filterer(address common.Address, filterer bind.ContractFilterer) (*HanaTokenConsumerPancakeV3Filterer, error) {
	contract, err := bindHanaTokenConsumerPancakeV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerPancakeV3Filterer{contract: contract}, nil
}

// bindHanaTokenConsumerPancakeV3 binds a generic wrapper to an already deployed contract.
func bindHanaTokenConsumerPancakeV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaTokenConsumerPancakeV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaTokenConsumerPancakeV3.Contract.HanaTokenConsumerPancakeV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.Contract.HanaTokenConsumerPancakeV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.Contract.HanaTokenConsumerPancakeV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaTokenConsumerPancakeV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.Contract.contract.Transact(opts, method, params...)
}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Caller) WETH9Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaTokenConsumerPancakeV3.contract.Call(opts, &out, "WETH9Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Session) WETH9Address() (common.Address, error) {
	return _HanaTokenConsumerPancakeV3.Contract.WETH9Address(&_HanaTokenConsumerPancakeV3.CallOpts)
}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3CallerSession) WETH9Address() (common.Address, error) {
	return _HanaTokenConsumerPancakeV3.Contract.WETH9Address(&_HanaTokenConsumerPancakeV3.CallOpts)
}

// HanaPoolFee is a free data retrieval call binding the contract method 0x207c1044.
//
// Solidity: function hanaPoolFee() view returns(uint24)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Caller) HanaPoolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HanaTokenConsumerPancakeV3.contract.Call(opts, &out, "hanaPoolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HanaPoolFee is a free data retrieval call binding the contract method 0x207c1044.
//
// Solidity: function hanaPoolFee() view returns(uint24)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Session) HanaPoolFee() (*big.Int, error) {
	return _HanaTokenConsumerPancakeV3.Contract.HanaPoolFee(&_HanaTokenConsumerPancakeV3.CallOpts)
}

// HanaPoolFee is a free data retrieval call binding the contract method 0x207c1044.
//
// Solidity: function hanaPoolFee() view returns(uint24)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3CallerSession) HanaPoolFee() (*big.Int, error) {
	return _HanaTokenConsumerPancakeV3.Contract.HanaPoolFee(&_HanaTokenConsumerPancakeV3.CallOpts)
}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Caller) HanaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaTokenConsumerPancakeV3.contract.Call(opts, &out, "hanaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Session) HanaToken() (common.Address, error) {
	return _HanaTokenConsumerPancakeV3.Contract.HanaToken(&_HanaTokenConsumerPancakeV3.CallOpts)
}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3CallerSession) HanaToken() (common.Address, error) {
	return _HanaTokenConsumerPancakeV3.Contract.HanaToken(&_HanaTokenConsumerPancakeV3.CallOpts)
}

// HasHanaLiquidity is a free data retrieval call binding the contract method 0x246d567e.
//
// Solidity: function hasHanaLiquidity() view returns(bool)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Caller) HasHanaLiquidity(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HanaTokenConsumerPancakeV3.contract.Call(opts, &out, "hasHanaLiquidity")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasHanaLiquidity is a free data retrieval call binding the contract method 0x246d567e.
//
// Solidity: function hasHanaLiquidity() view returns(bool)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Session) HasHanaLiquidity() (bool, error) {
	return _HanaTokenConsumerPancakeV3.Contract.HasHanaLiquidity(&_HanaTokenConsumerPancakeV3.CallOpts)
}

// HasHanaLiquidity is a free data retrieval call binding the contract method 0x246d567e.
//
// Solidity: function hasHanaLiquidity() view returns(bool)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3CallerSession) HasHanaLiquidity() (bool, error) {
	return _HanaTokenConsumerPancakeV3.Contract.HasHanaLiquidity(&_HanaTokenConsumerPancakeV3.CallOpts)
}

// PancakeV3Router is a free data retrieval call binding the contract method 0xc27745dd.
//
// Solidity: function pancakeV3Router() view returns(address)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Caller) PancakeV3Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaTokenConsumerPancakeV3.contract.Call(opts, &out, "pancakeV3Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PancakeV3Router is a free data retrieval call binding the contract method 0xc27745dd.
//
// Solidity: function pancakeV3Router() view returns(address)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Session) PancakeV3Router() (common.Address, error) {
	return _HanaTokenConsumerPancakeV3.Contract.PancakeV3Router(&_HanaTokenConsumerPancakeV3.CallOpts)
}

// PancakeV3Router is a free data retrieval call binding the contract method 0xc27745dd.
//
// Solidity: function pancakeV3Router() view returns(address)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3CallerSession) PancakeV3Router() (common.Address, error) {
	return _HanaTokenConsumerPancakeV3.Contract.PancakeV3Router(&_HanaTokenConsumerPancakeV3.CallOpts)
}

// TokenPoolFee is a free data retrieval call binding the contract method 0x5d9dfdde.
//
// Solidity: function tokenPoolFee() view returns(uint24)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Caller) TokenPoolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HanaTokenConsumerPancakeV3.contract.Call(opts, &out, "tokenPoolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenPoolFee is a free data retrieval call binding the contract method 0x5d9dfdde.
//
// Solidity: function tokenPoolFee() view returns(uint24)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Session) TokenPoolFee() (*big.Int, error) {
	return _HanaTokenConsumerPancakeV3.Contract.TokenPoolFee(&_HanaTokenConsumerPancakeV3.CallOpts)
}

// TokenPoolFee is a free data retrieval call binding the contract method 0x5d9dfdde.
//
// Solidity: function tokenPoolFee() view returns(uint24)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3CallerSession) TokenPoolFee() (*big.Int, error) {
	return _HanaTokenConsumerPancakeV3.Contract.TokenPoolFee(&_HanaTokenConsumerPancakeV3.CallOpts)
}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Caller) UniswapV3Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaTokenConsumerPancakeV3.contract.Call(opts, &out, "uniswapV3Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Session) UniswapV3Factory() (common.Address, error) {
	return _HanaTokenConsumerPancakeV3.Contract.UniswapV3Factory(&_HanaTokenConsumerPancakeV3.CallOpts)
}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3CallerSession) UniswapV3Factory() (common.Address, error) {
	return _HanaTokenConsumerPancakeV3.Contract.UniswapV3Factory(&_HanaTokenConsumerPancakeV3.CallOpts)
}

// GetEthFromHana is a paid mutator transaction binding the contract method 0x291d55f7.
//
// Solidity: function getEthFromHana(address destinationAddress, uint256 minAmountOut, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Transactor) GetEthFromHana(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.contract.Transact(opts, "getEthFromHana", destinationAddress, minAmountOut, hanaTokenAmount)
}

// GetEthFromHana is a paid mutator transaction binding the contract method 0x291d55f7.
//
// Solidity: function getEthFromHana(address destinationAddress, uint256 minAmountOut, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Session) GetEthFromHana(destinationAddress common.Address, minAmountOut *big.Int, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.Contract.GetEthFromHana(&_HanaTokenConsumerPancakeV3.TransactOpts, destinationAddress, minAmountOut, hanaTokenAmount)
}

// GetEthFromHana is a paid mutator transaction binding the contract method 0x291d55f7.
//
// Solidity: function getEthFromHana(address destinationAddress, uint256 minAmountOut, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3TransactorSession) GetEthFromHana(destinationAddress common.Address, minAmountOut *big.Int, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.Contract.GetEthFromHana(&_HanaTokenConsumerPancakeV3.TransactOpts, destinationAddress, minAmountOut, hanaTokenAmount)
}

// GetHanaFromEth is a paid mutator transaction binding the contract method 0x45413df7.
//
// Solidity: function getHanaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Transactor) GetHanaFromEth(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.contract.Transact(opts, "getHanaFromEth", destinationAddress, minAmountOut)
}

// GetHanaFromEth is a paid mutator transaction binding the contract method 0x45413df7.
//
// Solidity: function getHanaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Session) GetHanaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.Contract.GetHanaFromEth(&_HanaTokenConsumerPancakeV3.TransactOpts, destinationAddress, minAmountOut)
}

// GetHanaFromEth is a paid mutator transaction binding the contract method 0x45413df7.
//
// Solidity: function getHanaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3TransactorSession) GetHanaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.Contract.GetHanaFromEth(&_HanaTokenConsumerPancakeV3.TransactOpts, destinationAddress, minAmountOut)
}

// GetHanaFromToken is a paid mutator transaction binding the contract method 0x65d6edc8.
//
// Solidity: function getHanaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Transactor) GetHanaFromToken(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.contract.Transact(opts, "getHanaFromToken", destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetHanaFromToken is a paid mutator transaction binding the contract method 0x65d6edc8.
//
// Solidity: function getHanaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Session) GetHanaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.Contract.GetHanaFromToken(&_HanaTokenConsumerPancakeV3.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetHanaFromToken is a paid mutator transaction binding the contract method 0x65d6edc8.
//
// Solidity: function getHanaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3TransactorSession) GetHanaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.Contract.GetHanaFromToken(&_HanaTokenConsumerPancakeV3.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetTokenFromHana is a paid mutator transaction binding the contract method 0x5c2fec9a.
//
// Solidity: function getTokenFromHana(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Transactor) GetTokenFromHana(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.contract.Transact(opts, "getTokenFromHana", destinationAddress, minAmountOut, outputToken, hanaTokenAmount)
}

// GetTokenFromHana is a paid mutator transaction binding the contract method 0x5c2fec9a.
//
// Solidity: function getTokenFromHana(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Session) GetTokenFromHana(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.Contract.GetTokenFromHana(&_HanaTokenConsumerPancakeV3.TransactOpts, destinationAddress, minAmountOut, outputToken, hanaTokenAmount)
}

// GetTokenFromHana is a paid mutator transaction binding the contract method 0x5c2fec9a.
//
// Solidity: function getTokenFromHana(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3TransactorSession) GetTokenFromHana(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.Contract.GetTokenFromHana(&_HanaTokenConsumerPancakeV3.TransactOpts, destinationAddress, minAmountOut, outputToken, hanaTokenAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Session) Receive() (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.Contract.Receive(&_HanaTokenConsumerPancakeV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3TransactorSession) Receive() (*types.Transaction, error) {
	return _HanaTokenConsumerPancakeV3.Contract.Receive(&_HanaTokenConsumerPancakeV3.TransactOpts)
}

// HanaTokenConsumerPancakeV3EthExchangedForHanaIterator is returned from FilterEthExchangedForHana and is used to iterate over the raw logs and unpacked data for EthExchangedForHana events raised by the HanaTokenConsumerPancakeV3 contract.
type HanaTokenConsumerPancakeV3EthExchangedForHanaIterator struct {
	Event *HanaTokenConsumerPancakeV3EthExchangedForHana // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerPancakeV3EthExchangedForHanaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerPancakeV3EthExchangedForHana)
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
		it.Event = new(HanaTokenConsumerPancakeV3EthExchangedForHana)
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
func (it *HanaTokenConsumerPancakeV3EthExchangedForHanaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerPancakeV3EthExchangedForHanaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerPancakeV3EthExchangedForHana represents a EthExchangedForHana event raised by the HanaTokenConsumerPancakeV3 contract.
type HanaTokenConsumerPancakeV3EthExchangedForHana struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthExchangedForHana is a free log retrieval operation binding the contract event 0x877de62ce9cb129480429d4e4ac5b3ce3b8b85638de13f817d02b9cea87a0f36.
//
// Solidity: event EthExchangedForHana(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Filterer) FilterEthExchangedForHana(opts *bind.FilterOpts) (*HanaTokenConsumerPancakeV3EthExchangedForHanaIterator, error) {

	logs, sub, err := _HanaTokenConsumerPancakeV3.contract.FilterLogs(opts, "EthExchangedForHana")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerPancakeV3EthExchangedForHanaIterator{contract: _HanaTokenConsumerPancakeV3.contract, event: "EthExchangedForHana", logs: logs, sub: sub}, nil
}

// WatchEthExchangedForHana is a free log subscription operation binding the contract event 0x877de62ce9cb129480429d4e4ac5b3ce3b8b85638de13f817d02b9cea87a0f36.
//
// Solidity: event EthExchangedForHana(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Filterer) WatchEthExchangedForHana(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerPancakeV3EthExchangedForHana) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumerPancakeV3.contract.WatchLogs(opts, "EthExchangedForHana")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerPancakeV3EthExchangedForHana)
				if err := _HanaTokenConsumerPancakeV3.contract.UnpackLog(event, "EthExchangedForHana", log); err != nil {
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
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Filterer) ParseEthExchangedForHana(log types.Log) (*HanaTokenConsumerPancakeV3EthExchangedForHana, error) {
	event := new(HanaTokenConsumerPancakeV3EthExchangedForHana)
	if err := _HanaTokenConsumerPancakeV3.contract.UnpackLog(event, "EthExchangedForHana", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaTokenConsumerPancakeV3HanaExchangedForEthIterator is returned from FilterHanaExchangedForEth and is used to iterate over the raw logs and unpacked data for HanaExchangedForEth events raised by the HanaTokenConsumerPancakeV3 contract.
type HanaTokenConsumerPancakeV3HanaExchangedForEthIterator struct {
	Event *HanaTokenConsumerPancakeV3HanaExchangedForEth // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerPancakeV3HanaExchangedForEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerPancakeV3HanaExchangedForEth)
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
		it.Event = new(HanaTokenConsumerPancakeV3HanaExchangedForEth)
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
func (it *HanaTokenConsumerPancakeV3HanaExchangedForEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerPancakeV3HanaExchangedForEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerPancakeV3HanaExchangedForEth represents a HanaExchangedForEth event raised by the HanaTokenConsumerPancakeV3 contract.
type HanaTokenConsumerPancakeV3HanaExchangedForEth struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHanaExchangedForEth is a free log retrieval operation binding the contract event 0xd94ef7f26aa4584e0bb757c8b45e0663569b29a3a8a0fc85ed49a8bb95ed306a.
//
// Solidity: event HanaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Filterer) FilterHanaExchangedForEth(opts *bind.FilterOpts) (*HanaTokenConsumerPancakeV3HanaExchangedForEthIterator, error) {

	logs, sub, err := _HanaTokenConsumerPancakeV3.contract.FilterLogs(opts, "HanaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerPancakeV3HanaExchangedForEthIterator{contract: _HanaTokenConsumerPancakeV3.contract, event: "HanaExchangedForEth", logs: logs, sub: sub}, nil
}

// WatchHanaExchangedForEth is a free log subscription operation binding the contract event 0xd94ef7f26aa4584e0bb757c8b45e0663569b29a3a8a0fc85ed49a8bb95ed306a.
//
// Solidity: event HanaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Filterer) WatchHanaExchangedForEth(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerPancakeV3HanaExchangedForEth) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumerPancakeV3.contract.WatchLogs(opts, "HanaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerPancakeV3HanaExchangedForEth)
				if err := _HanaTokenConsumerPancakeV3.contract.UnpackLog(event, "HanaExchangedForEth", log); err != nil {
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
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Filterer) ParseHanaExchangedForEth(log types.Log) (*HanaTokenConsumerPancakeV3HanaExchangedForEth, error) {
	event := new(HanaTokenConsumerPancakeV3HanaExchangedForEth)
	if err := _HanaTokenConsumerPancakeV3.contract.UnpackLog(event, "HanaExchangedForEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaTokenConsumerPancakeV3HanaExchangedForTokenIterator is returned from FilterHanaExchangedForToken and is used to iterate over the raw logs and unpacked data for HanaExchangedForToken events raised by the HanaTokenConsumerPancakeV3 contract.
type HanaTokenConsumerPancakeV3HanaExchangedForTokenIterator struct {
	Event *HanaTokenConsumerPancakeV3HanaExchangedForToken // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerPancakeV3HanaExchangedForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerPancakeV3HanaExchangedForToken)
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
		it.Event = new(HanaTokenConsumerPancakeV3HanaExchangedForToken)
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
func (it *HanaTokenConsumerPancakeV3HanaExchangedForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerPancakeV3HanaExchangedForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerPancakeV3HanaExchangedForToken represents a HanaExchangedForToken event raised by the HanaTokenConsumerPancakeV3 contract.
type HanaTokenConsumerPancakeV3HanaExchangedForToken struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHanaExchangedForToken is a free log retrieval operation binding the contract event 0x532789829182476cbd3cc6523b623cc2504e5a0b348dfdc5527baf726e6d55af.
//
// Solidity: event HanaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Filterer) FilterHanaExchangedForToken(opts *bind.FilterOpts) (*HanaTokenConsumerPancakeV3HanaExchangedForTokenIterator, error) {

	logs, sub, err := _HanaTokenConsumerPancakeV3.contract.FilterLogs(opts, "HanaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerPancakeV3HanaExchangedForTokenIterator{contract: _HanaTokenConsumerPancakeV3.contract, event: "HanaExchangedForToken", logs: logs, sub: sub}, nil
}

// WatchHanaExchangedForToken is a free log subscription operation binding the contract event 0x532789829182476cbd3cc6523b623cc2504e5a0b348dfdc5527baf726e6d55af.
//
// Solidity: event HanaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Filterer) WatchHanaExchangedForToken(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerPancakeV3HanaExchangedForToken) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumerPancakeV3.contract.WatchLogs(opts, "HanaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerPancakeV3HanaExchangedForToken)
				if err := _HanaTokenConsumerPancakeV3.contract.UnpackLog(event, "HanaExchangedForToken", log); err != nil {
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
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Filterer) ParseHanaExchangedForToken(log types.Log) (*HanaTokenConsumerPancakeV3HanaExchangedForToken, error) {
	event := new(HanaTokenConsumerPancakeV3HanaExchangedForToken)
	if err := _HanaTokenConsumerPancakeV3.contract.UnpackLog(event, "HanaExchangedForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaTokenConsumerPancakeV3TokenExchangedForHanaIterator is returned from FilterTokenExchangedForHana and is used to iterate over the raw logs and unpacked data for TokenExchangedForHana events raised by the HanaTokenConsumerPancakeV3 contract.
type HanaTokenConsumerPancakeV3TokenExchangedForHanaIterator struct {
	Event *HanaTokenConsumerPancakeV3TokenExchangedForHana // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerPancakeV3TokenExchangedForHanaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerPancakeV3TokenExchangedForHana)
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
		it.Event = new(HanaTokenConsumerPancakeV3TokenExchangedForHana)
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
func (it *HanaTokenConsumerPancakeV3TokenExchangedForHanaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerPancakeV3TokenExchangedForHanaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerPancakeV3TokenExchangedForHana represents a TokenExchangedForHana event raised by the HanaTokenConsumerPancakeV3 contract.
type HanaTokenConsumerPancakeV3TokenExchangedForHana struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangedForHana is a free log retrieval operation binding the contract event 0xc7f70cbba54984322735d319a149c7bccec0ab310897fbc85b00f59d4486ce60.
//
// Solidity: event TokenExchangedForHana(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Filterer) FilterTokenExchangedForHana(opts *bind.FilterOpts) (*HanaTokenConsumerPancakeV3TokenExchangedForHanaIterator, error) {

	logs, sub, err := _HanaTokenConsumerPancakeV3.contract.FilterLogs(opts, "TokenExchangedForHana")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerPancakeV3TokenExchangedForHanaIterator{contract: _HanaTokenConsumerPancakeV3.contract, event: "TokenExchangedForHana", logs: logs, sub: sub}, nil
}

// WatchTokenExchangedForHana is a free log subscription operation binding the contract event 0xc7f70cbba54984322735d319a149c7bccec0ab310897fbc85b00f59d4486ce60.
//
// Solidity: event TokenExchangedForHana(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Filterer) WatchTokenExchangedForHana(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerPancakeV3TokenExchangedForHana) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumerPancakeV3.contract.WatchLogs(opts, "TokenExchangedForHana")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerPancakeV3TokenExchangedForHana)
				if err := _HanaTokenConsumerPancakeV3.contract.UnpackLog(event, "TokenExchangedForHana", log); err != nil {
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
func (_HanaTokenConsumerPancakeV3 *HanaTokenConsumerPancakeV3Filterer) ParseTokenExchangedForHana(log types.Log) (*HanaTokenConsumerPancakeV3TokenExchangedForHana, error) {
	event := new(HanaTokenConsumerPancakeV3TokenExchangedForHana)
	if err := _HanaTokenConsumerPancakeV3.contract.UnpackLog(event, "TokenExchangedForHana", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
