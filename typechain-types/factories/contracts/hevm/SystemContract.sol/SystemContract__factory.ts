/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  SystemContract,
  SystemContractInterface,
} from "../../../../contracts/hevm/SystemContract.sol/SystemContract";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "whana_",
        type: "address",
      },
      {
        internalType: "address",
        name: "uniswapv2Factory_",
        type: "address",
      },
      {
        internalType: "address",
        name: "uniswapv2Router02_",
        type: "address",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    inputs: [],
    name: "CallerIsNotFungibleModule",
    type: "error",
  },
  {
    inputs: [],
    name: "CantBeIdenticalAddresses",
    type: "error",
  },
  {
    inputs: [],
    name: "CantBeZeroAddress",
    type: "error",
  },
  {
    inputs: [],
    name: "InvalidTarget",
    type: "error",
  },
  {
    inputs: [],
    name: "ZeroAddress",
    type: "error",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    name: "SetConnectorHEVM",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    name: "SetGasCoin",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    name: "SetGasHanaPool",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    name: "SetGasPrice",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    name: "SetWHana",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [],
    name: "SystemContractDeployed",
    type: "event",
  },
  {
    inputs: [],
    name: "FUNGIBLE_MODULE_ADDRESS",
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
    inputs: [
      {
        components: [
          {
            internalType: "bytes",
            name: "origin",
            type: "bytes",
          },
          {
            internalType: "address",
            name: "sender",
            type: "address",
          },
          {
            internalType: "uint256",
            name: "chainID",
            type: "uint256",
          },
        ],
        internalType: "struct hContext",
        name: "context",
        type: "tuple",
      },
      {
        internalType: "address",
        name: "hrc20",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "target",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "message",
        type: "bytes",
      },
    ],
    name: "depositAndCall",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    name: "gasCoinHRC20ByChainId",
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
    inputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    name: "gasHanaPoolByChainId",
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
    inputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    name: "gasPriceByChainId",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "hanaConnectorHEVMAddress",
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
    inputs: [
      {
        internalType: "address",
        name: "addr",
        type: "address",
      },
    ],
    name: "setConnectorHEVMAddress",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "chainID",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "hrc20",
        type: "address",
      },
    ],
    name: "setGasCoinHRC20",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "chainID",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "erc20",
        type: "address",
      },
    ],
    name: "setGasHanaPool",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "chainID",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "price",
        type: "uint256",
      },
    ],
    name: "setGasPrice",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "addr",
        type: "address",
      },
    ],
    name: "setWHANAContractAddress",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "uniswapv2FactoryAddress",
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
    inputs: [
      {
        internalType: "address",
        name: "factory",
        type: "address",
      },
      {
        internalType: "address",
        name: "tokenA",
        type: "address",
      },
      {
        internalType: "address",
        name: "tokenB",
        type: "address",
      },
    ],
    name: "uniswapv2PairFor",
    outputs: [
      {
        internalType: "address",
        name: "pair",
        type: "address",
      },
    ],
    stateMutability: "pure",
    type: "function",
  },
  {
    inputs: [],
    name: "uniswapv2Router02Address",
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
    name: "wHanaContractAddress",
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
] as const;

const _bytecode =
  "0x60c06040523480156200001157600080fd5b50604051620018aa380380620018aa8339818101604052810190620000379190620001ac565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614620000b1576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250507f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac560405160405180910390a15050506200025b565b600081519050620001a68162000241565b92915050565b600080600060608486031215620001c857620001c76200023c565b5b6000620001d88682870162000195565b9350506020620001eb8682870162000195565b9250506040620001fe8682870162000195565b9150509250925092565b600062000215826200021c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200024c8162000208565b81146200025857600080fd5b50565b60805160601c60a05160601c61161c6200028e600039600061086e0152600081816107920152610c41015261161c6000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063c39aca3711610097578063d936a01211610066578063d936a0121461025c578063d9e9b4c01461027a578063db2131e5146102aa578063f197e0e1146102c8576100f5565b8063c39aca37146101c2578063c63585cc146101de578063d42da8701461020e578063d7fd7afb1461022c576100f5565b80633ce4a5bc116100d35780633ce4a5bc1461014e57806379ae782c1461016c578063842da36d14610188578063a7cb0507146101a6576100f5565b806306fd10b4146100fa5780631668a3f0146101165780631e39355514610132575b600080fd5b610114600480360381019061010f919061104f565b6102f8565b005b610130600480360381019061012b9190610ebf565b610400565b005b61014c60048036038101906101479190610ebf565b61057d565b005b6101566106fa565b60405161016391906112b1565b60405180910390f35b6101866004803603810190610181919061104f565b610712565b005b61019061086c565b60405161019d91906112b1565b60405180910390f35b6101c060048036038101906101bb919061108f565b610890565b005b6101dc60048036038101906101d79190610f6c565b61095d565b005b6101f860048036038101906101f39190610eec565b610b8f565b60405161020591906112b1565b60405180910390f35b610216610c01565b60405161022391906112b1565b60405180910390f35b61024660048036038101906102419190611022565b610c27565b604051610253919061134a565b60405180910390f35b610264610c3f565b60405161027191906112b1565b60405180910390f35b610294600480360381019061028f9190611022565b610c63565b6040516102a191906112b1565b60405180910390f35b6102b2610c96565b6040516102bf91906112b1565b60405180910390f35b6102e260048036038101906102dd9190611022565b610cbc565b6040516102ef91906112b1565b60405180910390f35b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610371576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d82826040516103f4929190611365565b60405180910390a15050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610479576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156104e0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f495b3e78b13285a35b3924be7ea2c0f6363b2e777d484e0e67274f778b63914b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405161057291906112b1565b60405180910390a150565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105f6576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561065d576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fbf29203eef941ded6783003162d888cafd3914fdcf81df202d84a247108af57d600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516106ef91906112b1565b60405180910390a150565b73735b14bb79463307aacbed86daf3322b1e6226ab81565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461078b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006107da7f0000000000000000000000000000000000000000000000000000000000000000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684610b8f565b9050806002600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507ffde6856eeb152d7c29aff8485d17f698215cf3de7532688f9c59ecedfd2a77c7838260405161085f929190611365565b60405180910390a1505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610909576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600080848152602001908152602001600020819055507f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d828260405161095192919061138e565b60405180910390a15050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109d6576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161480610a4f57503073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610a86576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff166347e7ef2484866040518363ffffffff1660e01b8152600401610ac19291906112cc565b602060405180830381600087803b158015610adb57600080fd5b505af1158015610aef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b139190610f3f565b508273ffffffffffffffffffffffffffffffffffffffff1663de43156e87878786866040518663ffffffff1660e01b8152600401610b559594939291906112f5565b600060405180830381600087803b158015610b6f57600080fd5b505af1158015610b83573d6000803e3d6000fd5b50505050505050505050565b6000806000610b9e8585610cef565b91509150858282604051602001610bb6929190611243565b60405160208183030381529060405280519060200120604051602001610bdd92919061126f565b6040516020818303038152906040528051906020012060001c925050509392505050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006020528060005260406000206000915090505481565b7f000000000000000000000000000000000000000000000000000000000000000081565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610d58576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610d92578284610d95565b83835b8092508193505050600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610e04576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b600081359050610e1a816115a1565b92915050565b600081519050610e2f816115b8565b92915050565b60008083601f840112610e4b57610e4a61150e565b5b8235905067ffffffffffffffff811115610e6857610e67611509565b5b602083019150836001820283011115610e8457610e8361151d565b5b9250929050565b600060608284031215610ea157610ea0611513565b5b81905092915050565b600081359050610eb9816115cf565b92915050565b600060208284031215610ed557610ed461152c565b5b6000610ee384828501610e0b565b91505092915050565b600080600060608486031215610f0557610f0461152c565b5b6000610f1386828701610e0b565b9350506020610f2486828701610e0b565b9250506040610f3586828701610e0b565b9150509250925092565b600060208284031215610f5557610f5461152c565b5b6000610f6384828501610e20565b91505092915050565b60008060008060008060a08789031215610f8957610f8861152c565b5b600087013567ffffffffffffffff811115610fa757610fa6611522565b5b610fb389828a01610e8b565b9650506020610fc489828a01610e0b565b9550506040610fd589828a01610eaa565b9450506060610fe689828a01610e0b565b935050608087013567ffffffffffffffff81111561100757611006611522565b5b61101389828a01610e35565b92509250509295509295509295565b6000602082840312156110385761103761152c565b5b600061104684828501610eaa565b91505092915050565b600080604083850312156110665761106561152c565b5b600061107485828601610eaa565b925050602061108585828601610e0b565b9150509250929050565b600080604083850312156110a6576110a561152c565b5b60006110b485828601610eaa565b92505060206110c585828601610eaa565b9150509250929050565b6110d881611475565b82525050565b6110e781611475565b82525050565b6110fe6110f982611475565b6114d6565b82525050565b61111561111082611493565b6114e8565b82525050565b600061112783856113b7565b93506111348385846114c7565b61113d83611531565b840190509392505050565b600061115483856113c8565b93506111618385846114c7565b61116a83611531565b840190509392505050565b60006111826020836113d9565b915061118d8261154f565b602082019050919050565b60006111a56001836113d9565b91506111b082611578565b600182019050919050565b6000606083016111ce60008401846113fb565b85830360008701526111e183828461111b565b925050506111f260208401846113e4565b6111ff60208601826110cf565b5061120d604084018461145e565b61121a6040860182611225565b508091505092915050565b61122e816114bd565b82525050565b61123d816114bd565b82525050565b600061124f82856110ed565b60148201915061125f82846110ed565b6014820191508190509392505050565b600061127a82611198565b915061128682856110ed565b6014820191506112968284611104565b6020820191506112a582611175565b91508190509392505050565b60006020820190506112c660008301846110de565b92915050565b60006040820190506112e160008301856110de565b6112ee6020830184611234565b9392505050565b6000608082019050818103600083015261130f81886111bb565b905061131e60208301876110de565b61132b6040830186611234565b818103606083015261133e818486611148565b90509695505050505050565b600060208201905061135f6000830184611234565b92915050565b600060408201905061137a6000830185611234565b61138760208301846110de565b9392505050565b60006040820190506113a36000830185611234565b6113b06020830184611234565b9392505050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60006113f36020840184610e0b565b905092915050565b6000808335600160200384360303811261141857611417611527565b5b83810192508235915060208301925067ffffffffffffffff8211156114405761143f611504565b5b60018202360384131561145657611455611518565b5b509250929050565b600061146d6020840184610eaa565b905092915050565b60006114808261149d565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60006114e1826114f2565b9050919050565b6000819050919050565b60006114fd82611542565b9050919050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f600082015250565b7fff00000000000000000000000000000000000000000000000000000000000000600082015250565b6115aa81611475565b81146115b557600080fd5b50565b6115c181611487565b81146115cc57600080fd5b50565b6115d8816114bd565b81146115e357600080fd5b5056fea264697066735822122005cdc62db4a390de7a4d8dc2116fffe563de6fe9a7513bf2e458fd256654d68464736f6c63430008070033";

type SystemContractConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: SystemContractConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class SystemContract__factory extends ContractFactory {
  constructor(...args: SystemContractConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    whana_: PromiseOrValue<string>,
    uniswapv2Factory_: PromiseOrValue<string>,
    uniswapv2Router02_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<SystemContract> {
    return super.deploy(
      whana_,
      uniswapv2Factory_,
      uniswapv2Router02_,
      overrides || {}
    ) as Promise<SystemContract>;
  }
  override getDeployTransaction(
    whana_: PromiseOrValue<string>,
    uniswapv2Factory_: PromiseOrValue<string>,
    uniswapv2Router02_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(
      whana_,
      uniswapv2Factory_,
      uniswapv2Router02_,
      overrides || {}
    );
  }
  override attach(address: string): SystemContract {
    return super.attach(address) as SystemContract;
  }
  override connect(signer: Signer): SystemContract__factory {
    return super.connect(signer) as SystemContract__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): SystemContractInterface {
    return new utils.Interface(_abi) as SystemContractInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): SystemContract {
    return new Contract(address, _abi, signerOrProvider) as SystemContract;
  }
}
