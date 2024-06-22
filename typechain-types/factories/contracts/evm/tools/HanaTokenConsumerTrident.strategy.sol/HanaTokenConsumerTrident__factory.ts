/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../../common";
import type {
  HanaTokenConsumerTrident,
  HanaTokenConsumerTridentInterface,
} from "../../../../../contracts/evm/tools/HanaTokenConsumerTrident.strategy.sol/HanaTokenConsumerTrident";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "hanaToken_",
        type: "address",
      },
      {
        internalType: "address",
        name: "uniswapV3Router_",
        type: "address",
      },
      {
        internalType: "address",
        name: "WETH9Address_",
        type: "address",
      },
      {
        internalType: "address",
        name: "poolFactory_",
        type: "address",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    inputs: [],
    name: "ErrorSendingETH",
    type: "error",
  },
  {
    inputs: [],
    name: "InputCantBeZero",
    type: "error",
  },
  {
    inputs: [],
    name: "InvalidAddress",
    type: "error",
  },
  {
    inputs: [],
    name: "ReentrancyError",
    type: "error",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint256",
        name: "amountIn",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amountOut",
        type: "uint256",
      },
    ],
    name: "EthExchangedForHana",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint256",
        name: "amountIn",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amountOut",
        type: "uint256",
      },
    ],
    name: "HanaExchangedForEth",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "token",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amountIn",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amountOut",
        type: "uint256",
      },
    ],
    name: "HanaExchangedForToken",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "token",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amountIn",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amountOut",
        type: "uint256",
      },
    ],
    name: "TokenExchangedForHana",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "destinationAddress",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "minAmountOut",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "hanaTokenAmount",
        type: "uint256",
      },
    ],
    name: "getEthFromHana",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "destinationAddress",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "minAmountOut",
        type: "uint256",
      },
    ],
    name: "getHanaFromEth",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "payable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "destinationAddress",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "minAmountOut",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "inputToken",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "inputTokenAmount",
        type: "uint256",
      },
    ],
    name: "getHanaFromToken",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "destinationAddress",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "minAmountOut",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "outputToken",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "hanaTokenAmount",
        type: "uint256",
      },
    ],
    name: "getTokenFromHana",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "hanaToken",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "hasHanaLiquidity",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "poolFactory",
    outputs: [
      {
        internalType: "contract ConcentratedLiquidityPoolFactory",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "tridentRouter",
    outputs: [
      {
        internalType: "contract IPoolRouter",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    stateMutability: "payable",
    type: "receive",
  },
] as const;

const _bytecode =
  "0x6101006040523480156200001257600080fd5b5060405162002b5a38038062002b5a833981810160405281019062000038919062000245565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480620000a05750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b80620000d85750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b80620001105750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b1562000148576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508273ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b81525050505050506200030a565b6000815190506200023f81620002f0565b92915050565b60008060008060808587031215620002625762000261620002eb565b5b600062000272878288016200022e565b945050602062000285878288016200022e565b935050604062000298878288016200022e565b9250506060620002ab878288016200022e565b91505092959194509250565b6000620002c482620002cb565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620002fb81620002b7565b81146200030757600080fd5b50565b60805160601c60a05160601c60c05160601c60e05160601c61274662000414600039600081816103d3015281816106330152818161074d01528181610bd401528181610cc0015281816111fd01526113090152600081816103180152818161053e0152818161089901528181610b1901528181610f1c015281816110540152818161118201526115450152600081816102ce0152818161033a01528181610386015281816104980152818161072101528181610acf01528181610b3b01528181610b8701528181610ea70152818161103001526112d90152600081816103a70152818161070001528181610ba801528181610c8f015281816111d101526112b801526127466000f3fe60806040526004361061007f5760003560e01c80635c2fec9a1161004e5780635c2fec9a1461014e5780635e694a921461018b57806364b5528a146101b657806365d6edc8146101e157610086565b8063246d567e1461008b578063291d55f7146100b65780634219dc40146100f357806345413df71461011e57610086565b3661008657005b600080fd5b34801561009757600080fd5b506100a061021e565b6040516100ad91906121ea565b60405180910390f35b3480156100c257600080fd5b506100dd60048036038101906100d89190611cb7565b610223565b6040516100ea919061231a565b60405180910390f35b3480156100ff57600080fd5b50610108610631565b6040516101159190612205565b60405180910390f35b61013860048036038101906101339190611c10565b610655565b604051610145919061231a565b60405180910390f35b34801561015a57600080fd5b5061017560048036038101906101709190611c50565b61098c565b604051610182919061231a565b60405180910390f35b34801561019757600080fd5b506101a061102e565b6040516101ad91906120ca565b60405180910390f35b3480156101c257600080fd5b506101cb611052565b6040516101d89190612220565b60405180910390f35b3480156101ed57600080fd5b5061020860048036038101906102039190611c50565b611076565b604051610215919061231a565b60405180910390f35b600090565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561028b576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008214156102c6576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103133330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661163d909392919063ffffffff16565b61037e7f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166116c69092919063ffffffff16565b6000806103cb7f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611824565b9150915060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128484600060016040518563ffffffff1660e01b8152600401610432949392919061210e565b60006040518083038186803b15801561044a57600080fd5b505afa15801561045e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906104879190611d0a565b905060006040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168152602001878152602001888152602001836000815181106104f2576104f1612532565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff16815260200160011515815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c07f5c32836040518263ffffffff1660e01b815260040161059591906122ff565b602060405180830381600087803b1580156105af57600080fd5b505af11580156105c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e79190611d80565b90507fd94ef7f26aa4584e0bb757c8b45e0663569b29a3a8a0fc85ed49a8bb95ed306a878260405161061a929190612335565b60405180910390a180955050505050509392505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156106bd576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60003414156106f8576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806107457f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611824565b9150915060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128484600060016040518563ffffffff1660e01b81526004016107ac949392919061210e565b60006040518083038186803b1580156107c457600080fd5b505afa1580156107d8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906108019190611d0a565b905060006040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020013481526020018781526020018360008151811061084d5761084c612532565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200160001515815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c07f5c3234846040518363ffffffff1660e01b81526004016108f191906122ff565b6020604051808303818588803b15801561090a57600080fd5b505af115801561091e573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906109439190611d80565b90507f877de62ce9cb129480429d4e4ac5b3ce3b8b85638de13f817d02b9cea87a0f363482604051610976929190612335565b60405180910390a1809550505050505092915050565b60008060009054906101000a900460ff16156109d4576040517f29f745a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016000806101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161480610a555750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610a8c576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610ac7576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b143330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661163d909392919063ffffffff16565b610b7f7f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166116c69092919063ffffffff16565b600080610bcc7f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611824565b9150915060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128484600060016040518563ffffffff1660e01b8152600401610c33949392919061210e565b60006040518083038186803b158015610c4b57600080fd5b505afa158015610c5f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610c889190611d0a565b9050610cb47f000000000000000000000000000000000000000000000000000000000000000087611824565b809350819450505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128585600060016040518563ffffffff1660e01b8152600401610d1f949392919061210e565b60006040518083038186803b158015610d3757600080fd5b505afa158015610d4b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610d749190611d0a565b90506000600267ffffffffffffffff811115610d9357610d92612561565b5b604051908082528060200260200182016040528015610dc15781602001602082028036833780820191505090505b50905082600081518110610dd857610dd7612532565b5b602002602001015181600081518110610df457610df3612532565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505081600081518110610e4257610e41612532565b5b602002602001015181600181518110610e5e57610e5d612532565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060006040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018b81526020018381526020018c73ffffffffffffffffffffffffffffffffffffffff16815260200160001515815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663363a9dba836040518263ffffffff1660e01b8152600401610f7391906122dd565b602060405180830381600087803b158015610f8d57600080fd5b505af1158015610fa1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fc59190611d80565b90507f532789829182476cbd3cc6523b623cc2504e5a0b348dfdc5527baf726e6d55af8a8a83604051610ffa939291906121b3565b60405180910390a18097505050505050505060008060006101000a81548160ff021916908315150217905550949350505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614806110de5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15611115576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415611150576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61117d3330848673ffffffffffffffffffffffffffffffffffffffff1661163d909392919063ffffffff16565b6111c87f0000000000000000000000000000000000000000000000000000000000000000838573ffffffffffffffffffffffffffffffffffffffff166116c69092919063ffffffff16565b6000806111f5857f0000000000000000000000000000000000000000000000000000000000000000611824565b9150915060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128484600060016040518563ffffffff1660e01b815260040161125c949392919061210e565b60006040518083038186803b15801561127457600080fd5b505afa158015611288573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906112b19190611d0a565b90506112fd7f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611824565b809350819450505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128585600060016040518563ffffffff1660e01b8152600401611368949392919061210e565b60006040518083038186803b15801561138057600080fd5b505afa158015611394573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906113bd9190611d0a565b90506000600267ffffffffffffffff8111156113dc576113db612561565b5b60405190808252806020026020018201604052801561140a5781602001602082028036833780820191505090505b5090508260008151811061142157611420612532565b5b60200260200101518160008151811061143d5761143c612532565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508160008151811061148b5761148a612532565b5b6020026020010151816001815181106114a7576114a6612532565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060006040518060c001604052808a73ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018b81526020018381526020018c73ffffffffffffffffffffffffffffffffffffffff16815260200160001515815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663363a9dba836040518263ffffffff1660e01b815260040161159c91906122dd565b602060405180830381600087803b1580156115b657600080fd5b505af11580156115ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ee9190611d80565b90507fc7f70cbba54984322735d319a149c7bccec0ab310897fbc85b00f59d4486ce608a8a83604051611623939291906121b3565b60405180910390a180975050505050505050949350505050565b6116c0846323b872dd60e01b85858560405160240161165e93929190612153565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611874565b50505050565b600081148061175f575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b815260040161170d9291906120e5565b60206040518083038186803b15801561172557600080fd5b505afa158015611739573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061175d9190611d80565b145b61179e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611795906122bd565b60405180910390fd5b61181f8363095ea7b360e01b84846040516024016117bd92919061218a565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611874565b505050565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610156118665783839150915061186d565b8284915091505b9250929050565b60006118d6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661193b9092919063ffffffff16565b905060008151111561193657808060200190518101906118f69190611d53565b611935576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161192c9061229d565b60405180910390fd5b5b505050565b606061194a8484600085611953565b90509392505050565b606082471015611998576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161198f9061225d565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516119c191906120b3565b60006040518083038185875af1925050503d80600081146119fe576040519150601f19603f3d011682016040523d82523d6000602084013e611a03565b606091505b5091509150611a1487838387611a20565b92505050949350505050565b60608315611a8357600083511415611a7b57611a3b85611a96565b611a7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a719061227d565b60405180910390fd5b5b829050611a8e565b611a8d8383611ab9565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082511115611acc5781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b00919061223b565b60405180910390fd5b6000611b1c611b1784612383565b61235e565b90508083825260208201905082856020860282011115611b3f57611b3e612595565b5b60005b85811015611b6f5781611b558882611b8e565b845260208401935060208301925050600181019050611b42565b5050509392505050565b600081359050611b88816126cb565b92915050565b600081519050611b9d816126cb565b92915050565b600082601f830112611bb857611bb7612590565b5b8151611bc8848260208601611b09565b91505092915050565b600081519050611be0816126e2565b92915050565b600081359050611bf5816126f9565b92915050565b600081519050611c0a816126f9565b92915050565b60008060408385031215611c2757611c2661259f565b5b6000611c3585828601611b79565b9250506020611c4685828601611be6565b9150509250929050565b60008060008060808587031215611c6a57611c6961259f565b5b6000611c7887828801611b79565b9450506020611c8987828801611be6565b9350506040611c9a87828801611b79565b9250506060611cab87828801611be6565b91505092959194509250565b600080600060608486031215611cd057611ccf61259f565b5b6000611cde86828701611b79565b9350506020611cef86828701611be6565b9250506040611d0086828701611be6565b9150509250925092565b600060208284031215611d2057611d1f61259f565b5b600082015167ffffffffffffffff811115611d3e57611d3d61259a565b5b611d4a84828501611ba3565b91505092915050565b600060208284031215611d6957611d6861259f565b5b6000611d7784828501611bd1565b91505092915050565b600060208284031215611d9657611d9561259f565b5b6000611da484828501611bfb565b91505092915050565b6000611db98383611dc5565b60208301905092915050565b611dce8161241a565b82525050565b611ddd8161241a565b82525050565b6000611dee826123bf565b611df881856123ed565b9350611e03836123af565b8060005b83811015611e34578151611e1b8882611dad565b9750611e26836123e0565b925050600181019050611e07565b5085935050505092915050565b611e4a8161242c565b82525050565b611e598161242c565b82525050565b6000611e6a826123ca565b611e7481856123fe565b9350611e848185602086016124ce565b80840191505092915050565b611e9981612462565b82525050565b611ea881612474565b82525050565b611eb781612486565b82525050565b611ec681612498565b82525050565b6000611ed7826123d5565b611ee18185612409565b9350611ef18185602086016124ce565b611efa816125a4565b840191505092915050565b6000611f12602683612409565b9150611f1d826125b5565b604082019050919050565b6000611f35601d83612409565b9150611f4082612604565b602082019050919050565b6000611f58602a83612409565b9150611f638261262d565b604082019050919050565b6000611f7b603683612409565b9150611f868261267c565b604082019050919050565b600060c083016000830151611fa96000860182611dc5565b506020830151611fbc6020860182612095565b506040830151611fcf6040860182612095565b5060608301518482036060860152611fe78282611de3565b9150506080830151611ffc6080860182611dc5565b5060a083015161200f60a0860182611e41565b508091505092915050565b60c0820160008201516120306000850182611dc5565b5060208201516120436020850182612095565b5060408201516120566040850182612095565b5060608201516120696060850182611dc5565b50608082015161207c6080850182611dc5565b5060a082015161208f60a0850182611e41565b50505050565b61209e81612458565b82525050565b6120ad81612458565b82525050565b60006120bf8284611e5f565b915081905092915050565b60006020820190506120df6000830184611dd4565b92915050565b60006040820190506120fa6000830185611dd4565b6121076020830184611dd4565b9392505050565b60006080820190506121236000830187611dd4565b6121306020830186611dd4565b61213d6040830185611eae565b61214a6060830184611ebd565b95945050505050565b60006060820190506121686000830186611dd4565b6121756020830185611dd4565b61218260408301846120a4565b949350505050565b600060408201905061219f6000830185611dd4565b6121ac60208301846120a4565b9392505050565b60006060820190506121c86000830186611dd4565b6121d560208301856120a4565b6121e260408301846120a4565b949350505050565b60006020820190506121ff6000830184611e50565b92915050565b600060208201905061221a6000830184611e90565b92915050565b60006020820190506122356000830184611e9f565b92915050565b600060208201905081810360008301526122558184611ecc565b905092915050565b6000602082019050818103600083015261227681611f05565b9050919050565b6000602082019050818103600083015261229681611f28565b9050919050565b600060208201905081810360008301526122b681611f4b565b9050919050565b600060208201905081810360008301526122d681611f6e565b9050919050565b600060208201905081810360008301526122f78184611f91565b905092915050565b600060c082019050612314600083018461201a565b92915050565b600060208201905061232f60008301846120a4565b92915050565b600060408201905061234a60008301856120a4565b61235760208301846120a4565b9392505050565b6000612368612379565b90506123748282612501565b919050565b6000604051905090565b600067ffffffffffffffff82111561239e5761239d612561565b5b602082029050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061242582612438565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061246d826124aa565b9050919050565b600061247f826124aa565b9050919050565b600061249182612458565b9050919050565b60006124a382612458565b9050919050565b60006124b5826124bc565b9050919050565b60006124c782612438565b9050919050565b60005b838110156124ec5780820151818401526020810190506124d1565b838111156124fb576000848401525b50505050565b61250a826125a4565b810181811067ffffffffffffffff8211171561252957612528612561565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60008201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000602082015250565b6126d48161241a565b81146126df57600080fd5b50565b6126eb8161242c565b81146126f657600080fd5b50565b61270281612458565b811461270d57600080fd5b5056fea2646970667358221220c9655d5354cea5bf908706039b3daff223316a838736f7f18a41bdfed9ab84f464736f6c63430008070033";

type HanaTokenConsumerTridentConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: HanaTokenConsumerTridentConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class HanaTokenConsumerTrident__factory extends ContractFactory {
  constructor(...args: HanaTokenConsumerTridentConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    hanaToken_: PromiseOrValue<string>,
    uniswapV3Router_: PromiseOrValue<string>,
    WETH9Address_: PromiseOrValue<string>,
    poolFactory_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<HanaTokenConsumerTrident> {
    return super.deploy(
      hanaToken_,
      uniswapV3Router_,
      WETH9Address_,
      poolFactory_,
      overrides || {}
    ) as Promise<HanaTokenConsumerTrident>;
  }
  override getDeployTransaction(
    hanaToken_: PromiseOrValue<string>,
    uniswapV3Router_: PromiseOrValue<string>,
    WETH9Address_: PromiseOrValue<string>,
    poolFactory_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(
      hanaToken_,
      uniswapV3Router_,
      WETH9Address_,
      poolFactory_,
      overrides || {}
    );
  }
  override attach(address: string): HanaTokenConsumerTrident {
    return super.attach(address) as HanaTokenConsumerTrident;
  }
  override connect(signer: Signer): HanaTokenConsumerTrident__factory {
    return super.connect(signer) as HanaTokenConsumerTrident__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): HanaTokenConsumerTridentInterface {
    return new utils.Interface(_abi) as HanaTokenConsumerTridentInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): HanaTokenConsumerTrident {
    return new Contract(
      address,
      _abi,
      signerOrProvider
    ) as HanaTokenConsumerTrident;
  }
}
