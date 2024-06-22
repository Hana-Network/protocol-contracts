// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanatokenconsumeruniv2

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

// HanaTokenConsumerUniV2MetaData contains all meta data concerning the HanaTokenConsumerUniV2 contract.
var HanaTokenConsumerUniV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hanaToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV2Router_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthExchangedForHana\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"HanaExchangedForEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"HanaExchangedForToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenExchangedForHana\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hanaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getEthFromHana\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"getHanaFromEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getHanaFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hanaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenFromHana\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hanaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasHanaLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620029a0380380620029a083398181016040528101906200003791906200024e565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806200009f5750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15620000d7576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1663ad5c46486040518163ffffffff1660e01b815260040160206040518083038186803b1580156200018c57600080fd5b505afa158015620001a1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001c791906200021c565b73ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250505050620002e8565b6000815190506200021681620002ce565b92915050565b600060208284031215620002355762000234620002c9565b5b6000620002458482850162000205565b91505092915050565b60008060408385031215620002685762000267620002c9565b5b6000620002788582860162000205565b92505060206200028b8582860162000205565b9150509250929050565b6000620002a282620002a9565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620002d98162000295565b8114620002e557600080fd5b50565b60805160601c60a05160601c60c05160601c6125c9620003d7600039600081816102c6015281816104a40152818161063a0152818161094301528181610ba301528181610f0c0152818161117501526114be0152600081816102570152818161045a015281816104c60152818161055a015281816108d201528181610b5901528181610bc501528181610cad01528181610ddc01528181611047015281816112ce015261144c0152600081816101e8015281816105c90152818161086301528181610c0d01528181610d1c01528181610e4b015281816111bf0152818161125f01526113dd01526125c96000f3fe6080604052600436106100555760003560e01c8063246d567e1461005a578063291d55f71461008557806345413df7146100c25780635c2fec9a146100f25780635e694a921461012f57806365d6edc81461015a575b600080fd5b34801561006657600080fd5b5061006f610197565b60405161007c9190611fab565b60405180910390f35b34801561009157600080fd5b506100ac60048036038101906100a79190611c0c565b6103af565b6040516100b99190612098565b60405180910390f35b6100dc60048036038101906100d79190611b65565b610770565b6040516100e99190612098565b60405180910390f35b3480156100fe57600080fd5b5061011960048036038101906101149190611ba5565b610a77565b6040516101269190612098565b60405180910390f35b34801561013b57600080fd5b50610144611045565b6040516101519190611ed0565b60405180910390f35b34801561016657600080fd5b50610181600480360381019061017c9190611ba5565b611069565b60405161018e9190612098565b60405180910390f35b600080600267ffffffffffffffff8111156101b5576101b46123e4565b5b6040519080825280602002602001820160405280156101e35781602001602082028036833780820191505090505b5090507f00000000000000000000000000000000000000000000000000000000000000008160008151811061021b5761021a6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f00000000000000000000000000000000000000000000000000000000000000008160018151811061028a576102896123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d06ca61f6001836040518363ffffffff1660e01b8152600401610320929190611fc6565b60006040518083038186803b15801561033857600080fd5b505afa92505050801561036e57506040513d6000823e3d601f19601f8201168201806040525081019061036b9190611c5f565b60015b61037c5760009150506103ac565b6000816001845161038d9190612294565b8151811061039e5761039d6123b5565b5b602002602001015111925050505b90565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610417576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610452576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61049f3330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166115f7909392919063ffffffff16565b61050a7f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166116809092919063ffffffff16565b6000600267ffffffffffffffff811115610527576105266123e4565b5b6040519080825280602002602001820160405280156105555781602001602082028036833780820191505090505b5090507f00000000000000000000000000000000000000000000000000000000000000008160008151811061058d5761058c6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f0000000000000000000000000000000000000000000000000000000000000000816001815181106105fc576105fb6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318cbafe58587858a60c842610685919061223e565b6040518663ffffffff1660e01b81526004016106a5959493929190612128565b600060405180830381600087803b1580156106bf57600080fd5b505af11580156106d3573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906106fc9190611c5f565b90506000816001845161070f9190612294565b815181106107205761071f6123b5565b5b602002602001015190507fd94ef7f26aa4584e0bb757c8b45e0663569b29a3a8a0fc85ed49a8bb95ed306a858260405161075b9291906120ff565b60405180910390a18093505050509392505050565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156107d8576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000341415610813576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600267ffffffffffffffff8111156108305761082f6123e4565b5b60405190808252806020026020018201604052801561085e5781602001602082028036833780820191505090505b5090507f000000000000000000000000000000000000000000000000000000000000000081600081518110610896576108956123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f000000000000000000000000000000000000000000000000000000000000000081600181518110610905576109046123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16637ff36ab53486858960c84261098e919061223e565b6040518663ffffffff1660e01b81526004016109ad94939291906120b3565b6000604051808303818588803b1580156109c657600080fd5b505af11580156109da573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f82011682018060405250810190610a049190611c5f565b905060008160018451610a179190612294565b81518110610a2857610a276123b5565b5b602002602001015190507f877de62ce9cb129480429d4e4ac5b3ce3b8b85638de13f817d02b9cea87a0f363482604051610a639291906120ff565b60405180910390a180935050505092915050565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161480610adf5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610b16576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610b51576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b9e3330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166115f7909392919063ffffffff16565b610c097f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166116809092919063ffffffff16565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610d8e57600267ffffffffffffffff811115610c7a57610c796123e4565b5b604051908082528060200260200182016040528015610ca85781602001602082028036833780820191505090505b5090507f000000000000000000000000000000000000000000000000000000000000000081600081518110610ce057610cdf6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f000000000000000000000000000000000000000000000000000000000000000081600181518110610d4f57610d4e6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050610f08565b600367ffffffffffffffff811115610da957610da86123e4565b5b604051908082528060200260200182016040528015610dd75781602001602082028036833780820191505090505b5090507f000000000000000000000000000000000000000000000000000000000000000081600081518110610e0f57610e0e6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f000000000000000000000000000000000000000000000000000000000000000081600181518110610e7e57610e7d6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508381600281518110610ecd57610ecc6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166338ed17398588858b60c842610f57919061223e565b6040518663ffffffff1660e01b8152600401610f77959493929190612128565b600060405180830381600087803b158015610f9157600080fd5b505af1158015610fa5573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610fce9190611c5f565b905060008160018451610fe19190612294565b81518110610ff257610ff16123b5565b5b602002602001015190507f532789829182476cbd3cc6523b623cc2504e5a0b348dfdc5527baf726e6d55af86868360405161102f93929190611f74565b60405180910390a1809350505050949350505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614806110d15750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15611108576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415611143576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111703330848673ffffffffffffffffffffffffffffffffffffffff166115f7909392919063ffffffff16565b6111bb7f0000000000000000000000000000000000000000000000000000000000000000838573ffffffffffffffffffffffffffffffffffffffff166116809092919063ffffffff16565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561134057600267ffffffffffffffff81111561122c5761122b6123e4565b5b60405190808252806020026020018201604052801561125a5781602001602082028036833780820191505090505b5090507f000000000000000000000000000000000000000000000000000000000000000081600081518110611292576112916123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f000000000000000000000000000000000000000000000000000000000000000081600181518110611301576113006123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506114ba565b600367ffffffffffffffff81111561135b5761135a6123e4565b5b6040519080825280602002602001820160405280156113895781602001602082028036833780820191505090505b50905083816000815181106113a1576113a06123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f0000000000000000000000000000000000000000000000000000000000000000816001815181106114105761140f6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f00000000000000000000000000000000000000000000000000000000000000008160028151811061147f5761147e6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166338ed17398588858b60c842611509919061223e565b6040518663ffffffff1660e01b8152600401611529959493929190612128565b600060405180830381600087803b15801561154357600080fd5b505af1158015611557573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906115809190611c5f565b9050600081600184516115939190612294565b815181106115a4576115a36123b5565b5b602002602001015190507fc7f70cbba54984322735d319a149c7bccec0ab310897fbc85b00f59d4486ce608686836040516115e193929190611f74565b60405180910390a1809350505050949350505050565b61167a846323b872dd60e01b85858560405160240161161893929190611f14565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506117de565b50505050565b6000811480611719575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b81526004016116c7929190611eeb565b60206040518083038186803b1580156116df57600080fd5b505afa1580156116f3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117179190611cd5565b145b611758576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161174f90612078565b60405180910390fd5b6117d98363095ea7b360e01b8484604051602401611777929190611f4b565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506117de565b505050565b6000611840826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166118a59092919063ffffffff16565b90506000815111156118a057808060200190518101906118609190611ca8565b61189f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161189690612058565b60405180910390fd5b5b505050565b60606118b484846000856118bd565b90509392505050565b606082471015611902576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118f990612018565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161192b9190611eb9565b60006040518083038185875af1925050503d8060008114611968576040519150601f19603f3d011682016040523d82523d6000602084013e61196d565b606091505b509150915061197e8783838761198a565b92505050949350505050565b606083156119ed576000835114156119e5576119a585611a00565b6119e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119db90612038565b60405180910390fd5b5b8290506119f8565b6119f78383611a23565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082511115611a365781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a6a9190611ff6565b60405180910390fd5b6000611a86611a81846121a7565b612182565b90508083825260208201905082856020860282011115611aa957611aa8612418565b5b60005b85811015611ad95781611abf8882611b50565b845260208401935060208301925050600181019050611aac565b5050509392505050565b600081359050611af28161254e565b92915050565b600082601f830112611b0d57611b0c612413565b5b8151611b1d848260208601611a73565b91505092915050565b600081519050611b3581612565565b92915050565b600081359050611b4a8161257c565b92915050565b600081519050611b5f8161257c565b92915050565b60008060408385031215611b7c57611b7b612422565b5b6000611b8a85828601611ae3565b9250506020611b9b85828601611b3b565b9150509250929050565b60008060008060808587031215611bbf57611bbe612422565b5b6000611bcd87828801611ae3565b9450506020611bde87828801611b3b565b9350506040611bef87828801611ae3565b9250506060611c0087828801611b3b565b91505092959194509250565b600080600060608486031215611c2557611c24612422565b5b6000611c3386828701611ae3565b9350506020611c4486828701611b3b565b9250506040611c5586828701611b3b565b9150509250925092565b600060208284031215611c7557611c74612422565b5b600082015167ffffffffffffffff811115611c9357611c9261241d565b5b611c9f84828501611af8565b91505092915050565b600060208284031215611cbe57611cbd612422565b5b6000611ccc84828501611b26565b91505092915050565b600060208284031215611ceb57611cea612422565b5b6000611cf984828501611b50565b91505092915050565b6000611d0e8383611d1a565b60208301905092915050565b611d23816122c8565b82525050565b611d32816122c8565b82525050565b6000611d43826121e3565b611d4d8185612211565b9350611d58836121d3565b8060005b83811015611d89578151611d708882611d02565b9750611d7b83612204565b925050600181019050611d5c565b5085935050505092915050565b611d9f816122da565b82525050565b6000611db0826121ee565b611dba8185612222565b9350611dca818560208601612322565b80840191505092915050565b611ddf81612310565b82525050565b6000611df0826121f9565b611dfa818561222d565b9350611e0a818560208601612322565b611e1381612427565b840191505092915050565b6000611e2b60268361222d565b9150611e3682612438565b604082019050919050565b6000611e4e601d8361222d565b9150611e5982612487565b602082019050919050565b6000611e71602a8361222d565b9150611e7c826124b0565b604082019050919050565b6000611e9460368361222d565b9150611e9f826124ff565b604082019050919050565b611eb381612306565b82525050565b6000611ec58284611da5565b915081905092915050565b6000602082019050611ee56000830184611d29565b92915050565b6000604082019050611f006000830185611d29565b611f0d6020830184611d29565b9392505050565b6000606082019050611f296000830186611d29565b611f366020830185611d29565b611f436040830184611eaa565b949350505050565b6000604082019050611f606000830185611d29565b611f6d6020830184611eaa565b9392505050565b6000606082019050611f896000830186611d29565b611f966020830185611eaa565b611fa36040830184611eaa565b949350505050565b6000602082019050611fc06000830184611d96565b92915050565b6000604082019050611fdb6000830185611dd6565b8181036020830152611fed8184611d38565b90509392505050565b600060208201905081810360008301526120108184611de5565b905092915050565b6000602082019050818103600083015261203181611e1e565b9050919050565b6000602082019050818103600083015261205181611e41565b9050919050565b6000602082019050818103600083015261207181611e64565b9050919050565b6000602082019050818103600083015261209181611e87565b9050919050565b60006020820190506120ad6000830184611eaa565b92915050565b60006080820190506120c86000830187611eaa565b81810360208301526120da8186611d38565b90506120e96040830185611d29565b6120f66060830184611eaa565b95945050505050565b60006040820190506121146000830185611eaa565b6121216020830184611eaa565b9392505050565b600060a08201905061213d6000830188611eaa565b61214a6020830187611eaa565b818103604083015261215c8186611d38565b905061216b6060830185611d29565b6121786080830184611eaa565b9695505050505050565b600061218c61219d565b90506121988282612355565b919050565b6000604051905090565b600067ffffffffffffffff8211156121c2576121c16123e4565b5b602082029050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061224982612306565b915061225483612306565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561228957612288612386565b5b828201905092915050565b600061229f82612306565b91506122aa83612306565b9250828210156122bd576122bc612386565b5b828203905092915050565b60006122d3826122e6565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061231b82612306565b9050919050565b60005b83811015612340578082015181840152602081019050612325565b8381111561234f576000848401525b50505050565b61235e82612427565b810181811067ffffffffffffffff8211171561237d5761237c6123e4565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60008201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000602082015250565b612557816122c8565b811461256257600080fd5b50565b61256e816122da565b811461257957600080fd5b50565b61258581612306565b811461259057600080fd5b5056fea26469706673582212201cf9f4f67085592057da751f9bac80d525340846ae677747ca7cf23bcd2044ba64736f6c63430008070033",
}

// HanaTokenConsumerUniV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaTokenConsumerUniV2MetaData.ABI instead.
var HanaTokenConsumerUniV2ABI = HanaTokenConsumerUniV2MetaData.ABI

// HanaTokenConsumerUniV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HanaTokenConsumerUniV2MetaData.Bin instead.
var HanaTokenConsumerUniV2Bin = HanaTokenConsumerUniV2MetaData.Bin

// DeployHanaTokenConsumerUniV2 deploys a new Ethereum contract, binding an instance of HanaTokenConsumerUniV2 to it.
func DeployHanaTokenConsumerUniV2(auth *bind.TransactOpts, backend bind.ContractBackend, hanaToken_ common.Address, uniswapV2Router_ common.Address) (common.Address, *types.Transaction, *HanaTokenConsumerUniV2, error) {
	parsed, err := HanaTokenConsumerUniV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HanaTokenConsumerUniV2Bin), backend, hanaToken_, uniswapV2Router_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HanaTokenConsumerUniV2{HanaTokenConsumerUniV2Caller: HanaTokenConsumerUniV2Caller{contract: contract}, HanaTokenConsumerUniV2Transactor: HanaTokenConsumerUniV2Transactor{contract: contract}, HanaTokenConsumerUniV2Filterer: HanaTokenConsumerUniV2Filterer{contract: contract}}, nil
}

// HanaTokenConsumerUniV2 is an auto generated Go binding around an Ethereum contract.
type HanaTokenConsumerUniV2 struct {
	HanaTokenConsumerUniV2Caller     // Read-only binding to the contract
	HanaTokenConsumerUniV2Transactor // Write-only binding to the contract
	HanaTokenConsumerUniV2Filterer   // Log filterer for contract events
}

// HanaTokenConsumerUniV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerUniV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerUniV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaTokenConsumerUniV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaTokenConsumerUniV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaTokenConsumerUniV2Session struct {
	Contract     *HanaTokenConsumerUniV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HanaTokenConsumerUniV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaTokenConsumerUniV2CallerSession struct {
	Contract *HanaTokenConsumerUniV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// HanaTokenConsumerUniV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaTokenConsumerUniV2TransactorSession struct {
	Contract     *HanaTokenConsumerUniV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// HanaTokenConsumerUniV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type HanaTokenConsumerUniV2Raw struct {
	Contract *HanaTokenConsumerUniV2 // Generic contract binding to access the raw methods on
}

// HanaTokenConsumerUniV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV2CallerRaw struct {
	Contract *HanaTokenConsumerUniV2Caller // Generic read-only contract binding to access the raw methods on
}

// HanaTokenConsumerUniV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaTokenConsumerUniV2TransactorRaw struct {
	Contract *HanaTokenConsumerUniV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaTokenConsumerUniV2 creates a new instance of HanaTokenConsumerUniV2, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV2(address common.Address, backend bind.ContractBackend) (*HanaTokenConsumerUniV2, error) {
	contract, err := bindHanaTokenConsumerUniV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV2{HanaTokenConsumerUniV2Caller: HanaTokenConsumerUniV2Caller{contract: contract}, HanaTokenConsumerUniV2Transactor: HanaTokenConsumerUniV2Transactor{contract: contract}, HanaTokenConsumerUniV2Filterer: HanaTokenConsumerUniV2Filterer{contract: contract}}, nil
}

// NewHanaTokenConsumerUniV2Caller creates a new read-only instance of HanaTokenConsumerUniV2, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV2Caller(address common.Address, caller bind.ContractCaller) (*HanaTokenConsumerUniV2Caller, error) {
	contract, err := bindHanaTokenConsumerUniV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV2Caller{contract: contract}, nil
}

// NewHanaTokenConsumerUniV2Transactor creates a new write-only instance of HanaTokenConsumerUniV2, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV2Transactor(address common.Address, transactor bind.ContractTransactor) (*HanaTokenConsumerUniV2Transactor, error) {
	contract, err := bindHanaTokenConsumerUniV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV2Transactor{contract: contract}, nil
}

// NewHanaTokenConsumerUniV2Filterer creates a new log filterer instance of HanaTokenConsumerUniV2, bound to a specific deployed contract.
func NewHanaTokenConsumerUniV2Filterer(address common.Address, filterer bind.ContractFilterer) (*HanaTokenConsumerUniV2Filterer, error) {
	contract, err := bindHanaTokenConsumerUniV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV2Filterer{contract: contract}, nil
}

// bindHanaTokenConsumerUniV2 binds a generic wrapper to an already deployed contract.
func bindHanaTokenConsumerUniV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaTokenConsumerUniV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaTokenConsumerUniV2.Contract.HanaTokenConsumerUniV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.Contract.HanaTokenConsumerUniV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.Contract.HanaTokenConsumerUniV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaTokenConsumerUniV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.Contract.contract.Transact(opts, method, params...)
}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Caller) HanaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaTokenConsumerUniV2.contract.Call(opts, &out, "hanaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Session) HanaToken() (common.Address, error) {
	return _HanaTokenConsumerUniV2.Contract.HanaToken(&_HanaTokenConsumerUniV2.CallOpts)
}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2CallerSession) HanaToken() (common.Address, error) {
	return _HanaTokenConsumerUniV2.Contract.HanaToken(&_HanaTokenConsumerUniV2.CallOpts)
}

// HasHanaLiquidity is a free data retrieval call binding the contract method 0x246d567e.
//
// Solidity: function hasHanaLiquidity() view returns(bool)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Caller) HasHanaLiquidity(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HanaTokenConsumerUniV2.contract.Call(opts, &out, "hasHanaLiquidity")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasHanaLiquidity is a free data retrieval call binding the contract method 0x246d567e.
//
// Solidity: function hasHanaLiquidity() view returns(bool)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Session) HasHanaLiquidity() (bool, error) {
	return _HanaTokenConsumerUniV2.Contract.HasHanaLiquidity(&_HanaTokenConsumerUniV2.CallOpts)
}

// HasHanaLiquidity is a free data retrieval call binding the contract method 0x246d567e.
//
// Solidity: function hasHanaLiquidity() view returns(bool)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2CallerSession) HasHanaLiquidity() (bool, error) {
	return _HanaTokenConsumerUniV2.Contract.HasHanaLiquidity(&_HanaTokenConsumerUniV2.CallOpts)
}

// GetEthFromHana is a paid mutator transaction binding the contract method 0x291d55f7.
//
// Solidity: function getEthFromHana(address destinationAddress, uint256 minAmountOut, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Transactor) GetEthFromHana(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.contract.Transact(opts, "getEthFromHana", destinationAddress, minAmountOut, hanaTokenAmount)
}

// GetEthFromHana is a paid mutator transaction binding the contract method 0x291d55f7.
//
// Solidity: function getEthFromHana(address destinationAddress, uint256 minAmountOut, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Session) GetEthFromHana(destinationAddress common.Address, minAmountOut *big.Int, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.Contract.GetEthFromHana(&_HanaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut, hanaTokenAmount)
}

// GetEthFromHana is a paid mutator transaction binding the contract method 0x291d55f7.
//
// Solidity: function getEthFromHana(address destinationAddress, uint256 minAmountOut, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2TransactorSession) GetEthFromHana(destinationAddress common.Address, minAmountOut *big.Int, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.Contract.GetEthFromHana(&_HanaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut, hanaTokenAmount)
}

// GetHanaFromEth is a paid mutator transaction binding the contract method 0x45413df7.
//
// Solidity: function getHanaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Transactor) GetHanaFromEth(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.contract.Transact(opts, "getHanaFromEth", destinationAddress, minAmountOut)
}

// GetHanaFromEth is a paid mutator transaction binding the contract method 0x45413df7.
//
// Solidity: function getHanaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Session) GetHanaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.Contract.GetHanaFromEth(&_HanaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut)
}

// GetHanaFromEth is a paid mutator transaction binding the contract method 0x45413df7.
//
// Solidity: function getHanaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2TransactorSession) GetHanaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.Contract.GetHanaFromEth(&_HanaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut)
}

// GetHanaFromToken is a paid mutator transaction binding the contract method 0x65d6edc8.
//
// Solidity: function getHanaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Transactor) GetHanaFromToken(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.contract.Transact(opts, "getHanaFromToken", destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetHanaFromToken is a paid mutator transaction binding the contract method 0x65d6edc8.
//
// Solidity: function getHanaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Session) GetHanaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.Contract.GetHanaFromToken(&_HanaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetHanaFromToken is a paid mutator transaction binding the contract method 0x65d6edc8.
//
// Solidity: function getHanaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2TransactorSession) GetHanaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.Contract.GetHanaFromToken(&_HanaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetTokenFromHana is a paid mutator transaction binding the contract method 0x5c2fec9a.
//
// Solidity: function getTokenFromHana(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Transactor) GetTokenFromHana(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.contract.Transact(opts, "getTokenFromHana", destinationAddress, minAmountOut, outputToken, hanaTokenAmount)
}

// GetTokenFromHana is a paid mutator transaction binding the contract method 0x5c2fec9a.
//
// Solidity: function getTokenFromHana(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Session) GetTokenFromHana(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.Contract.GetTokenFromHana(&_HanaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut, outputToken, hanaTokenAmount)
}

// GetTokenFromHana is a paid mutator transaction binding the contract method 0x5c2fec9a.
//
// Solidity: function getTokenFromHana(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 hanaTokenAmount) returns(uint256)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2TransactorSession) GetTokenFromHana(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, hanaTokenAmount *big.Int) (*types.Transaction, error) {
	return _HanaTokenConsumerUniV2.Contract.GetTokenFromHana(&_HanaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut, outputToken, hanaTokenAmount)
}

// HanaTokenConsumerUniV2EthExchangedForHanaIterator is returned from FilterEthExchangedForHana and is used to iterate over the raw logs and unpacked data for EthExchangedForHana events raised by the HanaTokenConsumerUniV2 contract.
type HanaTokenConsumerUniV2EthExchangedForHanaIterator struct {
	Event *HanaTokenConsumerUniV2EthExchangedForHana // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerUniV2EthExchangedForHanaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerUniV2EthExchangedForHana)
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
		it.Event = new(HanaTokenConsumerUniV2EthExchangedForHana)
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
func (it *HanaTokenConsumerUniV2EthExchangedForHanaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerUniV2EthExchangedForHanaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerUniV2EthExchangedForHana represents a EthExchangedForHana event raised by the HanaTokenConsumerUniV2 contract.
type HanaTokenConsumerUniV2EthExchangedForHana struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthExchangedForHana is a free log retrieval operation binding the contract event 0x877de62ce9cb129480429d4e4ac5b3ce3b8b85638de13f817d02b9cea87a0f36.
//
// Solidity: event EthExchangedForHana(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Filterer) FilterEthExchangedForHana(opts *bind.FilterOpts) (*HanaTokenConsumerUniV2EthExchangedForHanaIterator, error) {

	logs, sub, err := _HanaTokenConsumerUniV2.contract.FilterLogs(opts, "EthExchangedForHana")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV2EthExchangedForHanaIterator{contract: _HanaTokenConsumerUniV2.contract, event: "EthExchangedForHana", logs: logs, sub: sub}, nil
}

// WatchEthExchangedForHana is a free log subscription operation binding the contract event 0x877de62ce9cb129480429d4e4ac5b3ce3b8b85638de13f817d02b9cea87a0f36.
//
// Solidity: event EthExchangedForHana(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Filterer) WatchEthExchangedForHana(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerUniV2EthExchangedForHana) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumerUniV2.contract.WatchLogs(opts, "EthExchangedForHana")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerUniV2EthExchangedForHana)
				if err := _HanaTokenConsumerUniV2.contract.UnpackLog(event, "EthExchangedForHana", log); err != nil {
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
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Filterer) ParseEthExchangedForHana(log types.Log) (*HanaTokenConsumerUniV2EthExchangedForHana, error) {
	event := new(HanaTokenConsumerUniV2EthExchangedForHana)
	if err := _HanaTokenConsumerUniV2.contract.UnpackLog(event, "EthExchangedForHana", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaTokenConsumerUniV2HanaExchangedForEthIterator is returned from FilterHanaExchangedForEth and is used to iterate over the raw logs and unpacked data for HanaExchangedForEth events raised by the HanaTokenConsumerUniV2 contract.
type HanaTokenConsumerUniV2HanaExchangedForEthIterator struct {
	Event *HanaTokenConsumerUniV2HanaExchangedForEth // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerUniV2HanaExchangedForEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerUniV2HanaExchangedForEth)
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
		it.Event = new(HanaTokenConsumerUniV2HanaExchangedForEth)
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
func (it *HanaTokenConsumerUniV2HanaExchangedForEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerUniV2HanaExchangedForEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerUniV2HanaExchangedForEth represents a HanaExchangedForEth event raised by the HanaTokenConsumerUniV2 contract.
type HanaTokenConsumerUniV2HanaExchangedForEth struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHanaExchangedForEth is a free log retrieval operation binding the contract event 0xd94ef7f26aa4584e0bb757c8b45e0663569b29a3a8a0fc85ed49a8bb95ed306a.
//
// Solidity: event HanaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Filterer) FilterHanaExchangedForEth(opts *bind.FilterOpts) (*HanaTokenConsumerUniV2HanaExchangedForEthIterator, error) {

	logs, sub, err := _HanaTokenConsumerUniV2.contract.FilterLogs(opts, "HanaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV2HanaExchangedForEthIterator{contract: _HanaTokenConsumerUniV2.contract, event: "HanaExchangedForEth", logs: logs, sub: sub}, nil
}

// WatchHanaExchangedForEth is a free log subscription operation binding the contract event 0xd94ef7f26aa4584e0bb757c8b45e0663569b29a3a8a0fc85ed49a8bb95ed306a.
//
// Solidity: event HanaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Filterer) WatchHanaExchangedForEth(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerUniV2HanaExchangedForEth) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumerUniV2.contract.WatchLogs(opts, "HanaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerUniV2HanaExchangedForEth)
				if err := _HanaTokenConsumerUniV2.contract.UnpackLog(event, "HanaExchangedForEth", log); err != nil {
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
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Filterer) ParseHanaExchangedForEth(log types.Log) (*HanaTokenConsumerUniV2HanaExchangedForEth, error) {
	event := new(HanaTokenConsumerUniV2HanaExchangedForEth)
	if err := _HanaTokenConsumerUniV2.contract.UnpackLog(event, "HanaExchangedForEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaTokenConsumerUniV2HanaExchangedForTokenIterator is returned from FilterHanaExchangedForToken and is used to iterate over the raw logs and unpacked data for HanaExchangedForToken events raised by the HanaTokenConsumerUniV2 contract.
type HanaTokenConsumerUniV2HanaExchangedForTokenIterator struct {
	Event *HanaTokenConsumerUniV2HanaExchangedForToken // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerUniV2HanaExchangedForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerUniV2HanaExchangedForToken)
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
		it.Event = new(HanaTokenConsumerUniV2HanaExchangedForToken)
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
func (it *HanaTokenConsumerUniV2HanaExchangedForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerUniV2HanaExchangedForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerUniV2HanaExchangedForToken represents a HanaExchangedForToken event raised by the HanaTokenConsumerUniV2 contract.
type HanaTokenConsumerUniV2HanaExchangedForToken struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHanaExchangedForToken is a free log retrieval operation binding the contract event 0x532789829182476cbd3cc6523b623cc2504e5a0b348dfdc5527baf726e6d55af.
//
// Solidity: event HanaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Filterer) FilterHanaExchangedForToken(opts *bind.FilterOpts) (*HanaTokenConsumerUniV2HanaExchangedForTokenIterator, error) {

	logs, sub, err := _HanaTokenConsumerUniV2.contract.FilterLogs(opts, "HanaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV2HanaExchangedForTokenIterator{contract: _HanaTokenConsumerUniV2.contract, event: "HanaExchangedForToken", logs: logs, sub: sub}, nil
}

// WatchHanaExchangedForToken is a free log subscription operation binding the contract event 0x532789829182476cbd3cc6523b623cc2504e5a0b348dfdc5527baf726e6d55af.
//
// Solidity: event HanaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Filterer) WatchHanaExchangedForToken(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerUniV2HanaExchangedForToken) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumerUniV2.contract.WatchLogs(opts, "HanaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerUniV2HanaExchangedForToken)
				if err := _HanaTokenConsumerUniV2.contract.UnpackLog(event, "HanaExchangedForToken", log); err != nil {
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
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Filterer) ParseHanaExchangedForToken(log types.Log) (*HanaTokenConsumerUniV2HanaExchangedForToken, error) {
	event := new(HanaTokenConsumerUniV2HanaExchangedForToken)
	if err := _HanaTokenConsumerUniV2.contract.UnpackLog(event, "HanaExchangedForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaTokenConsumerUniV2TokenExchangedForHanaIterator is returned from FilterTokenExchangedForHana and is used to iterate over the raw logs and unpacked data for TokenExchangedForHana events raised by the HanaTokenConsumerUniV2 contract.
type HanaTokenConsumerUniV2TokenExchangedForHanaIterator struct {
	Event *HanaTokenConsumerUniV2TokenExchangedForHana // Event containing the contract specifics and raw log

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
func (it *HanaTokenConsumerUniV2TokenExchangedForHanaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaTokenConsumerUniV2TokenExchangedForHana)
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
		it.Event = new(HanaTokenConsumerUniV2TokenExchangedForHana)
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
func (it *HanaTokenConsumerUniV2TokenExchangedForHanaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaTokenConsumerUniV2TokenExchangedForHanaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaTokenConsumerUniV2TokenExchangedForHana represents a TokenExchangedForHana event raised by the HanaTokenConsumerUniV2 contract.
type HanaTokenConsumerUniV2TokenExchangedForHana struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangedForHana is a free log retrieval operation binding the contract event 0xc7f70cbba54984322735d319a149c7bccec0ab310897fbc85b00f59d4486ce60.
//
// Solidity: event TokenExchangedForHana(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Filterer) FilterTokenExchangedForHana(opts *bind.FilterOpts) (*HanaTokenConsumerUniV2TokenExchangedForHanaIterator, error) {

	logs, sub, err := _HanaTokenConsumerUniV2.contract.FilterLogs(opts, "TokenExchangedForHana")
	if err != nil {
		return nil, err
	}
	return &HanaTokenConsumerUniV2TokenExchangedForHanaIterator{contract: _HanaTokenConsumerUniV2.contract, event: "TokenExchangedForHana", logs: logs, sub: sub}, nil
}

// WatchTokenExchangedForHana is a free log subscription operation binding the contract event 0xc7f70cbba54984322735d319a149c7bccec0ab310897fbc85b00f59d4486ce60.
//
// Solidity: event TokenExchangedForHana(address token, uint256 amountIn, uint256 amountOut)
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Filterer) WatchTokenExchangedForHana(opts *bind.WatchOpts, sink chan<- *HanaTokenConsumerUniV2TokenExchangedForHana) (event.Subscription, error) {

	logs, sub, err := _HanaTokenConsumerUniV2.contract.WatchLogs(opts, "TokenExchangedForHana")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaTokenConsumerUniV2TokenExchangedForHana)
				if err := _HanaTokenConsumerUniV2.contract.UnpackLog(event, "TokenExchangedForHana", log); err != nil {
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
func (_HanaTokenConsumerUniV2 *HanaTokenConsumerUniV2Filterer) ParseTokenExchangedForHana(log types.Log) (*HanaTokenConsumerUniV2TokenExchangedForHana, error) {
	event := new(HanaTokenConsumerUniV2TokenExchangedForHana)
	if err := _HanaTokenConsumerUniV2.contract.UnpackLog(event, "TokenExchangedForHana", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
