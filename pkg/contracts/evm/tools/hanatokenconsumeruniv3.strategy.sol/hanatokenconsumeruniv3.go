// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanatokenconsumeruniv3

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

// HanaTokenConsumerUniV3MetaData contains all meta data concerning the HanaTokenConsumerUniV3 contract.
var HanaTokenConsumerUniV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hanaToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV3Router_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV3Factory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WETH9Address_\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"hanaPoolFee_\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"tokenPoolFee_\",\"type\":\"uint24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrorSendingETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthExchangedForHana\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"HanaExchangedForEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"HanaExchangedForToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenExchangedForHana\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH9Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hanaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getEthFromHana\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"getHanaFromEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getHanaFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hanaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenFromHana\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hanaPoolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hanaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasHanaLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPoolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV3Factory\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV3Router\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b50604051620029833803806200298383398181016040528101906200003891906200028a565b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161480620000a05750600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16145b80620000d85750600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16145b80620001105750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b1562000148576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508473ffffffffffffffffffffffffffffffffffffffff166101008173ffffffffffffffffffffffffffffffffffffffff1660601b815250508373ffffffffffffffffffffffffffffffffffffffff166101208173ffffffffffffffffffffffffffffffffffffffff1660601b815250508273ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508162ffffff1660808162ffffff1660e81b815250508062ffffff1660a08162ffffff1660e81b81525050505050505050620003a2565b6000815190506200026d816200036e565b92915050565b600081519050620002848162000388565b92915050565b60008060008060008060c08789031215620002aa57620002a962000369565b5b6000620002ba89828a016200025c565b9650506020620002cd89828a016200025c565b9550506040620002e089828a016200025c565b9450506060620002f389828a016200025c565b93505060806200030689828a0162000273565b92505060a06200031989828a0162000273565b9150509295509295509295565b600062000333826200033a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062ffffff82169050919050565b600080fd5b620003798162000326565b81146200038557600080fd5b50565b62000393816200035a565b81146200039f57600080fd5b50565b60805160e81c60a05160e81c60c05160601c60e05160601c6101005160601c6101205160601c6124ad620004d6600039600081816102e90152610c200152600081816105ca0152818161073c0152818161095d01528181610b2e01528181610dcf01528181610f2d0152818161118e01526112cc01526000818161034601528181610580015281816105ec0152818161064001528181610a6e01528181610d8501528181610df101528181610e440152818161106001526112470152600081816103250152818161067c015281816107e901528181610a3201528181610e860152818161120501526113c1015260008181610ea70152818161103c01526111e40152600081816102c201528181610367015281816106b801528181610aaa01528181610e65015261122601526124ad6000f3fe6080604052600436106100a05760003560e01c80635b549182116100645780635b5491821461019a5780635c2fec9a146101c55780635d9dfdde146102025780635e694a921461022d57806365d6edc814610258578063c469cf1414610295576100a7565b8063207c1044146100ac578063246d567e146100d7578063291d55f7146101025780632c76d7a61461013f57806345413df71461016a576100a7565b366100a757005b600080fd5b3480156100b857600080fd5b506100c16102c0565b6040516100ce9190612087565b60405180910390f35b3480156100e357600080fd5b506100ec6102e4565b6040516100f99190611f56565b60405180910390f35b34801561010e57600080fd5b50610129600480360381019061012491906119b1565b6104d5565b60405161013691906120a2565b60405180910390f35b34801561014b57600080fd5b5061015461095b565b6040516101619190611f71565b60405180910390f35b610184600480360381019061017f919061190a565b61097f565b60405161019191906120a2565b60405180910390f35b3480156101a657600080fd5b506101af610c1e565b6040516101bc9190611f8c565b60405180910390f35b3480156101d157600080fd5b506101ec60048036038101906101e7919061194a565b610c42565b6040516101f991906120a2565b60405180910390f35b34801561020e57600080fd5b5061021761103a565b6040516102249190612087565b60405180910390f35b34801561023957600080fd5b5061024261105e565b60405161024f9190611e44565b60405180910390f35b34801561026457600080fd5b5061027f600480360381019061027a919061194a565b611082565b60405161028c91906120a2565b60405180910390f35b3480156102a157600080fd5b506102aa6113bf565b6040516102b79190611e44565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b6000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631698ee827f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006040518463ffffffff1660e01b81526004016103a493929190611e88565b60206040518083038186803b1580156103bc57600080fd5b505afa1580156103d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f491906118dd565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156104355760009150506104d2565b600081905060008173ffffffffffffffffffffffffffffffffffffffff16631a6865026040518163ffffffff1660e01b815260040160206040518083038186803b15801561048257600080fd5b505afa158015610496573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ba9190611a31565b6fffffffffffffffffffffffffffffffff1611925050505b90565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561053d576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610578576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105c53330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166113e3909392919063ffffffff16565b6106307f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661146c9092919063ffffffff16565b60006040518061010001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000062ffffff1681526020013073ffffffffffffffffffffffffffffffffffffffff16815260200160c84261070a9190612129565b8152602001848152602001858152602001600073ffffffffffffffffffffffffffffffffffffffff16815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663414bf389836040518263ffffffff1660e01b8152600401610793919061206b565b602060405180830381600087803b1580156107ad57600080fd5b505af11580156107c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e59190611a5e565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d826040518263ffffffff1660e01b815260040161084091906120a2565b600060405180830381600087803b15801561085a57600080fd5b505af115801561086e573d6000803e3d6000fd5b505050507fd94ef7f26aa4584e0bb757c8b45e0663569b29a3a8a0fc85ed49a8bb95ed306a84826040516108a39291906120bd565b60405180910390a160008673ffffffffffffffffffffffffffffffffffffffff16826040516108d190611e2f565b60006040518083038185875af1925050503d806000811461090e576040519150601f19603f3d011682016040523d82523d6000602084013e610913565b606091505b505090508061094e576040517f3794aeaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8193505050509392505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156109e7576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000341415610a22576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006040518061010001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000062ffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff16815260200160c842610afc9190612129565b8152602001348152602001848152602001600073ffffffffffffffffffffffffffffffffffffffff16815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663414bf38934846040518363ffffffff1660e01b8152600401610b86919061206b565b6020604051808303818588803b158015610b9f57600080fd5b505af1158015610bb3573d6000803e3d6000fd5b50505050506040513d601f19601f82011682018060405250810190610bd89190611a5e565b90507f877de62ce9cb129480429d4e4ac5b3ce3b8b85638de13f817d02b9cea87a0f363482604051610c0b9291906120bd565b60405180910390a1809250505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008060009054906101000a900460ff1615610c8a576040517f29f745a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016000806101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161480610d0b5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610d42576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610d7d576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610dca3330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166113e3909392919063ffffffff16565b610e357f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661146c9092919063ffffffff16565b60006040518060a001604052807f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000089604051602001610edb959493929190611db9565b60405160208183030381529060405281526020018773ffffffffffffffffffffffffffffffffffffffff16815260200160c842610f189190612129565b815260200184815260200186815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c04b8d59836040518263ffffffff1660e01b8152600401610f849190612049565b602060405180830381600087803b158015610f9e57600080fd5b505af1158015610fb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fd69190611a5e565b90507f532789829182476cbd3cc6523b623cc2504e5a0b348dfdc5527baf726e6d55af85858360405161100b93929190611f1f565b60405180910390a1809250505060008060006101000a81548160ff021916908315150217905550949350505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614806110ea5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15611121576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082141561115c576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111893330848673ffffffffffffffffffffffffffffffffffffffff166113e3909392919063ffffffff16565b6111d47f0000000000000000000000000000000000000000000000000000000000000000838573ffffffffffffffffffffffffffffffffffffffff1661146c9092919063ffffffff16565b60006040518060a00160405280857f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060405160200161127a959493929190611db9565b60405160208183030381529060405281526020018773ffffffffffffffffffffffffffffffffffffffff16815260200160c8426112b79190612129565b815260200184815260200186815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c04b8d59836040518263ffffffff1660e01b81526004016113239190612049565b602060405180830381600087803b15801561133d57600080fd5b505af1158015611351573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113759190611a5e565b90507fc7f70cbba54984322735d319a149c7bccec0ab310897fbc85b00f59d4486ce608585836040516113aa93929190611f1f565b60405180910390a18092505050949350505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b611466846323b872dd60e01b85858560405160240161140493929190611ebf565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506115ca565b50505050565b6000811480611505575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b81526004016114b3929190611e5f565b60206040518083038186803b1580156114cb57600080fd5b505afa1580156114df573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115039190611a5e565b145b611544576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161153b90612029565b60405180910390fd5b6115c58363095ea7b360e01b8484604051602401611563929190611ef6565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506115ca565b505050565b600061162c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166116919092919063ffffffff16565b905060008151111561168c578080602001905181019061164c9190611a04565b61168b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161168290612009565b60405180910390fd5b5b505050565b60606116a084846000856116a9565b90509392505050565b6060824710156116ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116e590611fc9565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516117179190611e18565b60006040518083038185875af1925050503d8060008114611754576040519150601f19603f3d011682016040523d82523d6000602084013e611759565b606091505b509150915061176a87838387611776565b92505050949350505050565b606083156117d9576000835114156117d157611791856117ec565b6117d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117c790611fe9565b60405180910390fd5b5b8290506117e4565b6117e3838361180f565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156118225781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118569190611fa7565b60405180910390fd5b60008135905061186e8161241b565b92915050565b6000815190506118838161241b565b92915050565b60008151905061189881612432565b92915050565b6000815190506118ad81612449565b92915050565b6000813590506118c281612460565b92915050565b6000815190506118d781612460565b92915050565b6000602082840312156118f3576118f26122d2565b5b600061190184828501611874565b91505092915050565b60008060408385031215611921576119206122d2565b5b600061192f8582860161185f565b9250506020611940858286016118b3565b9150509250929050565b60008060008060808587031215611964576119636122d2565b5b60006119728782880161185f565b9450506020611983878288016118b3565b93505060406119948782880161185f565b92505060606119a5878288016118b3565b91505092959194509250565b6000806000606084860312156119ca576119c96122d2565b5b60006119d88682870161185f565b93505060206119e9868287016118b3565b92505060406119fa868287016118b3565b9150509250925092565b600060208284031215611a1a57611a196122d2565b5b6000611a2884828501611889565b91505092915050565b600060208284031215611a4757611a466122d2565b5b6000611a558482850161189e565b91505092915050565b600060208284031215611a7457611a736122d2565b5b6000611a82848285016118c8565b91505092915050565b611a948161217f565b82525050565b611aa38161217f565b82525050565b611aba611ab58261217f565b61226d565b82525050565b611ac981612191565b82525050565b6000611ada826120e6565b611ae481856120fc565b9350611af481856020860161223a565b611afd816122d7565b840191505092915050565b6000611b13826120e6565b611b1d818561210d565b9350611b2d81856020860161223a565b80840191505092915050565b611b42816121f2565b82525050565b611b5181612204565b82525050565b6000611b62826120f1565b611b6c8185612118565b9350611b7c81856020860161223a565b611b85816122d7565b840191505092915050565b6000611b9d602683612118565b9150611ba882612302565b604082019050919050565b6000611bc060008361210d565b9150611bcb82612351565b600082019050919050565b6000611be3601d83612118565b9150611bee82612354565b602082019050919050565b6000611c06602a83612118565b9150611c118261237d565b604082019050919050565b6000611c29603683612118565b9150611c34826123cc565b604082019050919050565b600060a0830160008301518482036000860152611c5c8282611acf565b9150506020830151611c716020860182611a8b565b506040830151611c846040860182611d9b565b506060830151611c976060860182611d9b565b506080830151611caa6080860182611d9b565b508091505092915050565b61010082016000820151611ccc6000850182611a8b565b506020820151611cdf6020850182611a8b565b506040820151611cf26040850182611d66565b506060820151611d056060850182611a8b565b506080820151611d186080850182611d9b565b5060a0820151611d2b60a0850182611d9b565b5060c0820151611d3e60c0850182611d9b565b5060e0820151611d5160e0850182611d57565b50505050565b611d60816121b9565b82525050565b611d6f816121d9565b82525050565b611d7e816121d9565b82525050565b611d95611d90826121d9565b612291565b82525050565b611da4816121e8565b82525050565b611db3816121e8565b82525050565b6000611dc58288611aa9565b601482019150611dd58287611d84565b600382019150611de58286611aa9565b601482019150611df58285611d84565b600382019150611e058284611aa9565b6014820191508190509695505050505050565b6000611e248284611b08565b915081905092915050565b6000611e3a82611bb3565b9150819050919050565b6000602082019050611e596000830184611a9a565b92915050565b6000604082019050611e746000830185611a9a565b611e816020830184611a9a565b9392505050565b6000606082019050611e9d6000830186611a9a565b611eaa6020830185611a9a565b611eb76040830184611d75565b949350505050565b6000606082019050611ed46000830186611a9a565b611ee16020830185611a9a565b611eee6040830184611daa565b949350505050565b6000604082019050611f0b6000830185611a9a565b611f186020830184611daa565b9392505050565b6000606082019050611f346000830186611a9a565b611f416020830185611daa565b611f4e6040830184611daa565b949350505050565b6000602082019050611f6b6000830184611ac0565b92915050565b6000602082019050611f866000830184611b39565b92915050565b6000602082019050611fa16000830184611b48565b92915050565b60006020820190508181036000830152611fc18184611b57565b905092915050565b60006020820190508181036000830152611fe281611b90565b9050919050565b6000602082019050818103600083015261200281611bd6565b9050919050565b6000602082019050818103600083015261202281611bf9565b9050919050565b6000602082019050818103600083015261204281611c1c565b9050919050565b600060208201905081810360008301526120638184611c3f565b905092915050565b6000610100820190506120816000830184611cb5565b92915050565b600060208201905061209c6000830184611d75565b92915050565b60006020820190506120b76000830184611daa565b92915050565b60006040820190506120d26000830185611daa565b6120df6020830184611daa565b9392505050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000612134826121e8565b915061213f836121e8565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115612174576121736122a3565b5b828201905092915050565b600061218a826121b9565b9050919050565b60008115159050919050565b60006fffffffffffffffffffffffffffffffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062ffffff82169050919050565b6000819050919050565b60006121fd82612216565b9050919050565b600061220f82612216565b9050919050565b600061222182612228565b9050919050565b6000612233826121b9565b9050919050565b60005b8381101561225857808201518184015260208101905061223d565b83811115612267576000848401525b50505050565b60006122788261227f565b9050919050565b600061228a826122f5565b9050919050565b600061229c826122e8565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6000601f19601f8301169050919050565b60008160e81b9050919050565b60008160601b9050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60008201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000602082015250565b6124248161217f565b811461242f57600080fd5b50565b61243b81612191565b811461244657600080fd5b50565b6124528161219d565b811461245d57600080fd5b50565b612469816121e8565b811461247457600080fd5b5056fea2646970667358221220d470d9357a7028924aaa804a693be283279bbcdf526037197988ae6b4c0d52c564736f6c63430008070033",
}

// HanaTokenConsumerUniV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaTokenConsumerUniV3MetaData.ABI instead.
var HanaTokenConsumerUniV3ABI = HanaTokenConsumerUniV3MetaData.ABI

// HanaTokenConsumerUniV3Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HanaTokenConsumerUniV3MetaData.Bin instead.
var HanaTokenConsumerUniV3Bin = HanaTokenConsumerUniV3MetaData.Bin

// DeployHanaTokenConsumerUniV3 deploys a new Ethereum contract, binding an instance of HanaTokenConsumerUniV3 to it.
func DeployHanaTokenConsumerUniV3(auth *bind.TransactOpts, backend bind.ContractBackend, hanaToken_ common.Address, uniswapV3Router_ common.Address, uniswapV3Factory_ common.Address, WETH9Address_ common.Address, hanaPoolFee_ *big.Int, tokenPoolFee_ *big.Int) (common.Address, *types.Transaction, *HanaTokenConsumerUniV3, error) {
	parsed, err := HanaTokenConsumerUniV3MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HanaTokenConsumerUniV3Bin), backend, hanaToken_, uniswapV3Router_, uniswapV3Factory_, WETH9Address_, hanaPoolFee_, tokenPoolFee_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HanaTokenConsumerUniV3{HanaTokenConsumerUniV3Caller: HanaTokenConsumerUniV3Caller{contract: contract}, HanaTokenConsumerUniV3Transactor: HanaTokenConsumerUniV3Transactor{contract: contract}, HanaTokenConsumerUniV3Filterer: HanaTokenConsumerUniV3Filterer{contract: contract}}, nil
}

// HanaTokenConsumerUniV3 is an auto generated Go binding around an Ethereum contract.
type HanaTokenConsumerUniV3 struct {
	HanaTokenConsumerUniV3Caller     // Read-only binding to the contract
	HanaTokenConsumerUniV3Transactor // Write-only binding to the contract
	HanaTokenConsumerUniV3Filterer   // Log filterer for contract events
}

// HanaTokenConsumerUniV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerUniV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerUniV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaTokenConsumerUniV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerUniV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaTokenConsumerUniV3Session struct {
	Contract     *HanaTokenConsumerUniV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HanaTokenConsumerUniV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaTokenConsumerUniV3CallerSession struct {
	Contract *HanaTokenConsumerUniV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// HanaTokenConsumerUniV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaTokenConsumerUniV3TransactorSession struct {
	Contract     *HanaTokenConsumerUniV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// HanaTokenConsumerUniV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type HanaTokenConsumerUniV3Raw struct {
	Contract *HanaTokenConsumerUniV3 // Generic contract binding to access the raw methods on
}

// HanaTokenConsumerUniV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV3CallerRaw struct {
	Contract *HanaTokenConsumerUniV3Caller // Generic read-only contract binding to access the raw methods on
}

// HanaTokenConsumerUniV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV3TransactorRaw struct {
	Contract *HanaTokenConsumerUniV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaTokenConsumerUniV3 creates a new instance of HanaTokenConsumerUniV3, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV3(address common.Address, backend bind.ContractBackend) (*HanaTokenConsumerUniV3, error) {
	contract, err := bindHanaTokenConsumerUniV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV3{HanaTokenConsumerUniV3Caller: HanaTokenConsumerUniV3Caller{contract: contract}, HanaTokenConsumerUniV3Transactor: HanaTokenConsumerUniV3Transactor{contract: contract}, HanaTokenConsumerUniV3Filterer: HanaTokenConsumerUniV3Filterer{contract: contract}}, nil
}

// NewHanaTokenConsumerUniV3Caller creates a new read-only instance of HanaTokenConsumerUniV3, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV3Caller(address common.Address, caller bind.ContractCaller) (*HanaTokenConsumerUniV3Caller, error) {
	contract, err := bindHanaTokenConsumerUniV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV3Caller{contract: contract}, nil
}

// NewHanaTokenConsumerUniV3Transactor creates a new write-only instance of HanaTokenConsumerUniV3, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV3Transactor(address common.Address, transactor bind.ContractTransactor) (*HanaTokenConsumerUniV3Transactor, error) {
	contract, err := bindHanaTokenConsumerUniV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV3Transactor{contract: contract}, nil
}

// NewHanaTokenConsumerUniV3Filterer creates a new log filterer instance of HanaTokenConsumerUniV3, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV3Filterer(address common.Address, filterer bind.ContractFilterer) (*HanaTokenConsumerUniV3Filterer, error) {
	contract, err := bindHanaTokenConsumerUniV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV3Filterer{contract: contract}, nil
}

// bindHanaTokenConsumerUniV3 binds a generic wrapper to an already deployed contract.
func bindHanaTokenConsumerUniV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaTokenConsumerUniV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaTokenConsumerUniV3.Contract.HanaTokenConsumerUniV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.Contract.HanaTokenConsumerUniV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.Contract.HanaTokenConsumerUniV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaTokenConsumerUniV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.Contract.contract.Transact(opts, method, params...)
}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Caller) WETH9Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaTokenConsumerUniV3.contract.Call(opts, &out, "WETH9Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Session) WETH9Address() (common.Address, error) {
	return _HanaTokenConsumerUniV3.Contract.WETH9Address(&_HanaTokenConsumerUniV3.CallOpts)
}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3CallerSession) WETH9Address() (common.Address, error) {
	return _HanaTokenConsumerUniV3.Contract.WETH9Address(&_HanaTokenConsumerUniV3.CallOpts)
}

// HanaPoolFee is a free data retrieval call binding the contract method 0x207c1044.
//
// Solidity: function hanaPoolFee() view returns(uint24)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Caller) HanaPoolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HanaTokenConsumerUniV3.contract.Call(opts, &out, "hanaPoolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HanaPoolFee is a free data retrieval call binding the contract method 0x207c1044.
//
// Solidity: function hanaPoolFee() view returns(uint24)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Session) HanaPoolFee() (*big.Int, error) {
	return _HanaTokenConsumerUniV3.Contract.HanaPoolFee(&_HanaTokenConsumerUniV3.CallOpts)
}

// HanaPoolFee is a free data retrieval call binding the contract method 0x207c1044.
//
// Solidity: function hanaPoolFee() view returns(uint24)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3CallerSession) HanaPoolFee() (*big.Int, error) {
	return _HanaTokenConsumerUniV3.Contract.HanaPoolFee(&_HanaTokenConsumerUniV3.CallOpts)
}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Caller) HanaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaTokenConsumerUniV3.contract.Call(opts, &out, "hanaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Session) HanaToken() (common.Address, error) {
	return _HanaTokenConsumerUniV3.Contract.HanaToken(&_HanaTokenConsumerUniV3.CallOpts)
}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3CallerSession) HanaToken() (common.Address, error) {
	return _HanaTokenConsumerUniV3.Contract.HanaToken(&_HanaTokenConsumerUniV3.CallOpts)
}

// HasHanaLiquidity is a free data retrieval call binding the contract method 0x246d567e.
//
// Solidity: function hasHanaLiquidity() view returns(bool)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Caller) HasHanaLiquidity(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HanaTokenConsumerUniV3.contract.Call(opts, &out, "hasHanaLiquidity")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasHanaLiquidity is a free data retrieval call binding the contract method 0x246d567e.
//
// Solidity: function hasHanaLiquidity() view returns(bool)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Session) HasHanaLiquidity() (bool, error) {
	return _HanaTokenConsumerUniV3.Contract.HasHanaLiquidity(&_HanaTokenConsumerUniV3.CallOpts)
}

// HasHanaLiquidity is a free data retrieval call binding the contract method 0x246d567e.
//
// Solidity: function hasHanaLiquidity() view returns(bool)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3CallerSession) HasHanaLiquidity() (bool, error) {
	return _HanaTokenConsumerUniV3.Contract.HasHanaLiquidity(&_HanaTokenConsumerUniV3.CallOpts)
}

// TokenPoolFee is a free data retrieval call binding the contract method 0x5d9dfdde.
//
// Solidity: function tokenPoolFee() view returns(uint24)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Caller) TokenPoolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HanaTokenConsumerUniV3.contract.Call(opts, &out, "tokenPoolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenPoolFee is a free data retrieval call binding the contract method 0x5d9dfdde.
//
// Solidity: function tokenPoolFee() view returns(uint24)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Session) TokenPoolFee() (*big.Int, error) {
	return _HanaTokenConsumerUniV3.Contract.TokenPoolFee(&_HanaTokenConsumerUniV3.CallOpts)
}

// TokenPoolFee is a free data retrieval call binding the contract method 0x5d9dfdde.
//
// Solidity: function tokenPoolFee() view returns(uint24)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3CallerSession) TokenPoolFee() (*big.Int, error) {
	return _HanaTokenConsumerUniV3.Contract.TokenPoolFee(&_HanaTokenConsumerUniV3.CallOpts)
}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Caller) UniswapV3Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaTokenConsumerUniV3.contract.Call(opts, &out, "uniswapV3Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Session) UniswapV3Factory() (common.Address, error) {
	return _HanaTokenConsumerUniV3.Contract.UniswapV3Factory(&_HanaTokenConsumerUniV3.CallOpts)
}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3CallerSession) UniswapV3Factory() (common.Address, error) {
	return _HanaTokenConsumerUniV3.Contract.UniswapV3Factory(&_HanaTokenConsumerUniV3.CallOpts)
}

// UniswapV3Router is a free data retrieval call binding the contract method 0x2c76d7a6.
//
// Solidity: function uniswapV3Router() view returns(address)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Caller) UniswapV3Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaTokenConsumerUniV3.contract.Call(opts, &out, "uniswapV3Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV3Router is a free data retrieval call binding the contract method 0x2c76d7a6.
//
// Solidity: function uniswapV3Router() view returns(address)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Session) UniswapV3Router() (common.Address, error) {
	return _HanaTokenConsumerUniV3.Contract.UniswapV3Router(&_HanaTokenConsumerUniV3.CallOpts)
}

// UniswapV3Router is a free data retrieval call binding the contract method 0x2c76d7a6.
//
// Solidity: function uniswapV3Router() view returns(address)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3CallerSession) UniswapV3Router() (common.Address, error) {
	return _HanaTokenConsumerUniV3.Contract.UniswapV3Router(&_HanaTokenConsumerUniV3.CallOpts)
}

// GetEthFromHana is a paid mutator transaction binding the contract method 0x291d55f7.
//
// Solidity: function getEthFromHana(address destinationAddress, uint256 minAmountOut, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Transactor) GetEthFromHana(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.contract.Transact(opts, "getEthFromHana", destinationAddress, minAmountOut, hanaTokenAmount)
}

// GetEthFromHana is a paid mutator transaction binding the contract method 0x291d55f7.
//
// Solidity: function getEthFromHana(address destinationAddress, uint256 minAmountOut, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Session) GetEthFromHana(destinationAddress common.Address, minAmountOut *big.Int, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.Contract.GetEthFromHana(&_HanaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, hanaTokenAmount)
}

// GetEthFromHana is a paid mutator transaction binding the contract method 0x291d55f7.
//
// Solidity: function getEthFromHana(address destinationAddress, uint256 minAmountOut, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3TransactorSession) GetEthFromHana(destinationAddress common.Address, minAmountOut *big.Int, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.Contract.GetEthFromHana(&_HanaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, hanaTokenAmount)
}

// GetHanaFromEth is a paid mutator transaction binding the contract method 0x45413df7.
//
// Solidity: function getHanaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Transactor) GetHanaFromEth(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.contract.Transact(opts, "getHanaFromEth", destinationAddress, minAmountOut)
}

// GetHanaFromEth is a paid mutator transaction binding the contract method 0x45413df7.
//
// Solidity: function getHanaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Session) GetHanaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.Contract.GetHanaFromEth(&_HanaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut)
}

// GetHanaFromEth is a paid mutator transaction binding the contract method 0x45413df7.
//
// Solidity: function getHanaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3TransactorSession) GetHanaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.Contract.GetHanaFromEth(&_HanaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut)
}

// GetHanaFromToken is a paid mutator transaction binding the contract method 0x65d6edc8.
//
// Solidity: function getHanaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Transactor) GetHanaFromToken(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.contract.Transact(opts, "getHanaFromToken", destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetHanaFromToken is a paid mutator transaction binding the contract method 0x65d6edc8.
//
// Solidity: function getHanaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Session) GetHanaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.Contract.GetHanaFromToken(&_HanaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetHanaFromToken is a paid mutator transaction binding the contract method 0x65d6edc8.
//
// Solidity: function getHanaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3TransactorSession) GetHanaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.Contract.GetHanaFromToken(&_HanaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetTokenFromHana is a paid mutator transaction binding the contract method 0x5c2fec9a.
//
// Solidity: function getTokenFromHana(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Transactor) GetTokenFromHana(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.contract.Transact(opts, "getTokenFromHana", destinationAddress, minAmountOut, outputToken, hanaTokenAmount)
}

// GetTokenFromHana is a paid mutator transaction binding the contract method 0x5c2fec9a.
//
// Solidity: function getTokenFromHana(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Session) GetTokenFromHana(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.Contract.GetTokenFromHana(&_HanaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, outputToken, hanaTokenAmount)
}

// GetTokenFromHana is a paid mutator transaction binding the contract method 0x5c2fec9a.
//
// Solidity: function getTokenFromHana(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3TransactorSession) GetTokenFromHana(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.Contract.GetTokenFromHana(&_HanaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, outputToken, hanaTokenAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Session) Receive() (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.Contract.Receive(&_HanaTokenConsumerUniV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3TransactorSession) Receive() (*types.Transaction, error) {
	return _HanaTokenConsumerUniV3.Contract.Receive(&_HanaTokenConsumerUniV3.TransactOpts)
}

// HanaTokenConsumerUniV3EthExchangedForHanaIterator is returned from FilterEthExchangedForHana and is used to iterate over the raw logs and unpacked data for EthExchangedForHana events raised by the HanaTokenConsumerUniV3 contract.
type HanaTokenConsumerUniV3EthExchangedForHanaIterator struct {
	Event *HanaTokenConsumerUniV3EthExchangedForHana // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerUniV3EthExchangedForHanaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerUniV3EthExchangedForHana)
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
		it.Event = new(HanaTokenConsumerUniV3EthExchangedForHana)
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
func (it *HanaTokenConsumerUniV3EthExchangedForHanaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerUniV3EthExchangedForHanaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerUniV3EthExchangedForHana represents a EthExchangedForHana event raised by the HanaTokenConsumerUniV3 contract.
type HanaTokenConsumerUniV3EthExchangedForHana struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthExchangedForHana is a free log retrieval operation binding the contract event 0x877de62ce9cb129480429d4e4ac5b3ce3b8b85638de13f817d02b9cea87a0f36.
//
// Solidity: event EthExchangedForHana(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Filterer) FilterEthExchangedForHana(opts *bind.FilterOpts) (*HanaTokenConsumerUniV3EthExchangedForHanaIterator, error) {

	logs, sub, err := _HanaTokenConsumerUniV3.contract.FilterLogs(opts, "EthExchangedForHana")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV3EthExchangedForHanaIterator{contract: _HanaTokenConsumerUniV3.contract, event: "EthExchangedForHana", logs: logs, sub: sub}, nil
}

// WatchEthExchangedForHana is a free log subscription operation binding the contract event 0x877de62ce9cb129480429d4e4ac5b3ce3b8b85638de13f817d02b9cea87a0f36.
//
// Solidity: event EthExchangedForHana(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Filterer) WatchEthExchangedForHana(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerUniV3EthExchangedForHana) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumerUniV3.contract.WatchLogs(opts, "EthExchangedForHana")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerUniV3EthExchangedForHana)
				if err := _HanaTokenConsumerUniV3.contract.UnpackLog(event, "EthExchangedForHana", log); err != nil {
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
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Filterer) ParseEthExchangedForHana(log types.Log) (*HanaTokenConsumerUniV3EthExchangedForHana, error) {
	event := new(HanaTokenConsumerUniV3EthExchangedForHana)
	if err := _HanaTokenConsumerUniV3.contract.UnpackLog(event, "EthExchangedForHana", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaTokenConsumerUniV3HanaExchangedForEthIterator is returned from FilterHanaExchangedForEth and is used to iterate over the raw logs and unpacked data for HanaExchangedForEth events raised by the HanaTokenConsumerUniV3 contract.
type HanaTokenConsumerUniV3HanaExchangedForEthIterator struct {
	Event *HanaTokenConsumerUniV3HanaExchangedForEth // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerUniV3HanaExchangedForEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerUniV3HanaExchangedForEth)
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
		it.Event = new(HanaTokenConsumerUniV3HanaExchangedForEth)
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
func (it *HanaTokenConsumerUniV3HanaExchangedForEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerUniV3HanaExchangedForEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerUniV3HanaExchangedForEth represents a HanaExchangedForEth event raised by the HanaTokenConsumerUniV3 contract.
type HanaTokenConsumerUniV3HanaExchangedForEth struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHanaExchangedForEth is a free log retrieval operation binding the contract event 0xd94ef7f26aa4584e0bb757c8b45e0663569b29a3a8a0fc85ed49a8bb95ed306a.
//
// Solidity: event HanaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Filterer) FilterHanaExchangedForEth(opts *bind.FilterOpts) (*HanaTokenConsumerUniV3HanaExchangedForEthIterator, error) {

	logs, sub, err := _HanaTokenConsumerUniV3.contract.FilterLogs(opts, "HanaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV3HanaExchangedForEthIterator{contract: _HanaTokenConsumerUniV3.contract, event: "HanaExchangedForEth", logs: logs, sub: sub}, nil
}

// WatchHanaExchangedForEth is a free log subscription operation binding the contract event 0xd94ef7f26aa4584e0bb757c8b45e0663569b29a3a8a0fc85ed49a8bb95ed306a.
//
// Solidity: event HanaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Filterer) WatchHanaExchangedForEth(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerUniV3HanaExchangedForEth) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumerUniV3.contract.WatchLogs(opts, "HanaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerUniV3HanaExchangedForEth)
				if err := _HanaTokenConsumerUniV3.contract.UnpackLog(event, "HanaExchangedForEth", log); err != nil {
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
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Filterer) ParseHanaExchangedForEth(log types.Log) (*HanaTokenConsumerUniV3HanaExchangedForEth, error) {
	event := new(HanaTokenConsumerUniV3HanaExchangedForEth)
	if err := _HanaTokenConsumerUniV3.contract.UnpackLog(event, "HanaExchangedForEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaTokenConsumerUniV3HanaExchangedForTokenIterator is returned from FilterHanaExchangedForToken and is used to iterate over the raw logs and unpacked data for HanaExchangedForToken events raised by the HanaTokenConsumerUniV3 contract.
type HanaTokenConsumerUniV3HanaExchangedForTokenIterator struct {
	Event *HanaTokenConsumerUniV3HanaExchangedForToken // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerUniV3HanaExchangedForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerUniV3HanaExchangedForToken)
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
		it.Event = new(HanaTokenConsumerUniV3HanaExchangedForToken)
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
func (it *HanaTokenConsumerUniV3HanaExchangedForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerUniV3HanaExchangedForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerUniV3HanaExchangedForToken represents a HanaExchangedForToken event raised by the HanaTokenConsumerUniV3 contract.
type HanaTokenConsumerUniV3HanaExchangedForToken struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHanaExchangedForToken is a free log retrieval operation binding the contract event 0x532789829182476cbd3cc6523b623cc2504e5a0b348dfdc5527baf726e6d55af.
//
// Solidity: event HanaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Filterer) FilterHanaExchangedForToken(opts *bind.FilterOpts) (*HanaTokenConsumerUniV3HanaExchangedForTokenIterator, error) {

	logs, sub, err := _HanaTokenConsumerUniV3.contract.FilterLogs(opts, "HanaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV3HanaExchangedForTokenIterator{contract: _HanaTokenConsumerUniV3.contract, event: "HanaExchangedForToken", logs: logs, sub: sub}, nil
}

// WatchHanaExchangedForToken is a free log subscription operation binding the contract event 0x532789829182476cbd3cc6523b623cc2504e5a0b348dfdc5527baf726e6d55af.
//
// Solidity: event HanaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Filterer) WatchHanaExchangedForToken(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerUniV3HanaExchangedForToken) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumerUniV3.contract.WatchLogs(opts, "HanaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerUniV3HanaExchangedForToken)
				if err := _HanaTokenConsumerUniV3.contract.UnpackLog(event, "HanaExchangedForToken", log); err != nil {
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
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Filterer) ParseHanaExchangedForToken(log types.Log) (*HanaTokenConsumerUniV3HanaExchangedForToken, error) {
	event := new(HanaTokenConsumerUniV3HanaExchangedForToken)
	if err := _HanaTokenConsumerUniV3.contract.UnpackLog(event, "HanaExchangedForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaTokenConsumerUniV3TokenExchangedForHanaIterator is returned from FilterTokenExchangedForHana and is used to iterate over the raw logs and unpacked data for TokenExchangedForHana events raised by the HanaTokenConsumerUniV3 contract.
type HanaTokenConsumerUniV3TokenExchangedForHanaIterator struct {
	Event *HanaTokenConsumerUniV3TokenExchangedForHana // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerUniV3TokenExchangedForHanaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerUniV3TokenExchangedForHana)
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
		it.Event = new(HanaTokenConsumerUniV3TokenExchangedForHana)
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
func (it *HanaTokenConsumerUniV3TokenExchangedForHanaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerUniV3TokenExchangedForHanaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerUniV3TokenExchangedForHana represents a TokenExchangedForHana event raised by the HanaTokenConsumerUniV3 contract.
type HanaTokenConsumerUniV3TokenExchangedForHana struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangedForHana is a free log retrieval operation binding the contract event 0xc7f70cbba54984322735d319a149c7bccec0ab310897fbc85b00f59d4486ce60.
//
// Solidity: event TokenExchangedForHana(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Filterer) FilterTokenExchangedForHana(opts *bind.FilterOpts) (*HanaTokenConsumerUniV3TokenExchangedForHanaIterator, error) {

	logs, sub, err := _HanaTokenConsumerUniV3.contract.FilterLogs(opts, "TokenExchangedForHana")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV3TokenExchangedForHanaIterator{contract: _HanaTokenConsumerUniV3.contract, event: "TokenExchangedForHana", logs: logs, sub: sub}, nil
}

// WatchTokenExchangedForHana is a free log subscription operation binding the contract event 0xc7f70cbba54984322735d319a149c7bccec0ab310897fbc85b00f59d4486ce60.
//
// Solidity: event TokenExchangedForHana(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Filterer) WatchTokenExchangedForHana(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerUniV3TokenExchangedForHana) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumerUniV3.contract.WatchLogs(opts, "TokenExchangedForHana")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerUniV3TokenExchangedForHana)
				if err := _HanaTokenConsumerUniV3.contract.UnpackLog(event, "TokenExchangedForHana", log); err != nil {
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
func (_HanaTokenConsumerUniV3 *HanaTokenConsumerUniV3Filterer) ParseTokenExchangedForHana(log types.Log) (*HanaTokenConsumerUniV3TokenExchangedForHana, error) {
	event := new(HanaTokenConsumerUniV3TokenExchangedForHana)
	if err := _HanaTokenConsumerUniV3.contract.UnpackLog(event, "TokenExchangedForHana", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
