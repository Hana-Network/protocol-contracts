/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  HanaInteractorMock,
  HanaInteractorMockInterface,
} from "../../../../contracts/evm/testing/HanaInteractorMock";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "hanaConnectorAddress",
        type: "address",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    inputs: [],
    name: "InvalidAddress",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "caller",
        type: "address",
      },
    ],
    name: "InvalidCaller",
    type: "error",
  },
  {
    inputs: [],
    name: "InvalidDestinationChainId",
    type: "error",
  },
  {
    inputs: [],
    name: "InvalidHanaMessageCall",
    type: "error",
  },
  {
    inputs: [],
    name: "InvalidHanaRevertCall",
    type: "error",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "previousOwner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "OwnershipTransferStarted",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "previousOwner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "OwnershipTransferred",
    type: "event",
  },
  {
    inputs: [],
    name: "acceptOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "connector",
    outputs: [
      {
        internalType: "contract HanaConnector",
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
    name: "interactorsByChainId",
    outputs: [
      {
        internalType: "bytes",
        name: "",
        type: "bytes",
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
            name: "hanaTxSenderAddress",
            type: "bytes",
          },
          {
            internalType: "uint256",
            name: "sourceChainId",
            type: "uint256",
          },
          {
            internalType: "address",
            name: "destinationAddress",
            type: "address",
          },
          {
            internalType: "uint256",
            name: "hanaValue",
            type: "uint256",
          },
          {
            internalType: "bytes",
            name: "message",
            type: "bytes",
          },
        ],
        internalType: "struct HanaInterfaces.HanaMessage",
        name: "hanaMessage",
        type: "tuple",
      },
    ],
    name: "onHanaMessage",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "address",
            name: "hanaTxSenderAddress",
            type: "address",
          },
          {
            internalType: "uint256",
            name: "sourceChainId",
            type: "uint256",
          },
          {
            internalType: "bytes",
            name: "destinationAddress",
            type: "bytes",
          },
          {
            internalType: "uint256",
            name: "destinationChainId",
            type: "uint256",
          },
          {
            internalType: "uint256",
            name: "remainingHanaValue",
            type: "uint256",
          },
          {
            internalType: "bytes",
            name: "message",
            type: "bytes",
          },
        ],
        internalType: "struct HanaInterfaces.HanaRevert",
        name: "hanaRevert",
        type: "tuple",
      },
    ],
    name: "onHanaRevert",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "owner",
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
    name: "pendingOwner",
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
    name: "renounceOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "destinationChainId",
        type: "uint256",
      },
      {
        internalType: "bytes",
        name: "contractAddress",
        type: "bytes",
      },
    ],
    name: "setInteractorByChainId",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "transferOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const;

const _bytecode =
  "0x60c06040523480156200001157600080fd5b50604051620011dc380380620011dc833981810160405281019062000037919062000228565b80620000586200004c6200010760201b60201c565b6200010f60201b60201c565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620000c0576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b46608081815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250505050620002ad565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556200014a816200014d60201b620005601760201c565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050620002228162000293565b92915050565b6000602082840312156200024157620002406200028e565b5b6000620002518482850162000211565b91505092915050565b600062000267826200026e565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200029e816200025a565b8114620002aa57600080fd5b50565b60805160a05160601c610f02620002da600039600081816103a80152610626015260005050610f026000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806383f3084f1161006657806383f3084f1461011f5780638da5cb5b1461013d578063b204be931461015b578063e30c397814610177578063f2fde38b146101955761009e565b80632618143f146100a35780634204cf9b146100d35780634fd3f7d7146100ef578063715018a61461010b57806379ba509714610115575b600080fd5b6100bd60048036038101906100b8919061098d565b6101b1565b6040516100ca9190610ba6565b60405180910390f35b6100ed60048036038101906100e89190610944565b610251565b005b610109600480360381019061010491906109ba565b6102d5565b005b610113610305565b005b61011d610319565b005b6101276103a6565b6040516101349190610bc8565b60405180910390f35b6101456103ca565b6040516101529190610b8b565b60405180910390f35b610175600480360381019061017091906108fb565b6103f3565b005b61017f610489565b60405161018c9190610b8b565b60405180910390f35b6101af60048036038101906101aa91906108ce565b6104b3565b005b600260205280600052604060002060009150905080546101d090610d87565b80601f01602080910402602001604051908101604052809291908181526020018280546101fc90610d87565b80156102495780601f1061021e57610100808354040283529160200191610249565b820191906000526020600020905b81548152906001019060200180831161022c57829003601f168201915b505050505081565b8061025a610624565b3073ffffffffffffffffffffffffffffffffffffffff1681600001602081019061028491906108ce565b73ffffffffffffffffffffffffffffffffffffffff16146102d1576040517fb53f9b7500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b6102dd6106b6565b81816002600086815260200190815260200160002091906102ff92919061076d565b50505050565b61030d6106b6565b6103176000610734565b565b6000610323610765565b90508073ffffffffffffffffffffffffffffffffffffffff16610344610489565b73ffffffffffffffffffffffffffffffffffffffff161461039a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039190610be3565b60405180910390fd5b6103a381610734565b50565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b806103fc610624565b60026000826020013581526020019081526020016000206040516104209190610b74565b60405180910390208180600001906104389190610c23565b604051610446929190610b5b565b604051809103902014610485576040517f39b0327e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6104bb6106b6565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff1661051b6103ca565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106b457336040517fcbd9d2e00000000000000000000000000000000000000000000000000000000081526004016106ab9190610b8b565b60405180910390fd5b565b6106be610765565b73ffffffffffffffffffffffffffffffffffffffff166106dc6103ca565b73ffffffffffffffffffffffffffffffffffffffff1614610732576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072990610c03565b60405180910390fd5b565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905561076281610560565b50565b600033905090565b82805461077990610d87565b90600052602060002090601f01602090048101928261079b57600085556107e2565b82601f106107b457803560ff19168380011785556107e2565b828001600101855582156107e2579182015b828111156107e15782358255916020019190600101906107c6565b5b5090506107ef91906107f3565b5090565b5b8082111561080c5760008160009055506001016107f4565b5090565b60008135905061081f81610e9e565b92915050565b60008083601f84011261083b5761083a610ded565b5b8235905067ffffffffffffffff81111561085857610857610de8565b5b60208301915083600182028301111561087457610873610e01565b5b9250929050565b600060a0828403121561089157610890610df7565b5b81905092915050565b600060c082840312156108b0576108af610df7565b5b81905092915050565b6000813590506108c881610eb5565b92915050565b6000602082840312156108e4576108e3610e10565b5b60006108f284828501610810565b91505092915050565b60006020828403121561091157610910610e10565b5b600082013567ffffffffffffffff81111561092f5761092e610e0b565b5b61093b8482850161087b565b91505092915050565b60006020828403121561095a57610959610e10565b5b600082013567ffffffffffffffff81111561097857610977610e0b565b5b6109848482850161089a565b91505092915050565b6000602082840312156109a3576109a2610e10565b5b60006109b1848285016108b9565b91505092915050565b6000806000604084860312156109d3576109d2610e10565b5b60006109e1868287016108b9565b935050602084013567ffffffffffffffff811115610a0257610a01610e0b565b5b610a0e86828701610825565b92509250509250925092565b610a2381610cd3565b82525050565b6000610a358385610cb7565b9350610a42838584610d45565b82840190509392505050565b6000610a5982610c9b565b610a638185610ca6565b9350610a73818560208601610d54565b610a7c81610e15565b840191505092915050565b60008154610a9481610d87565b610a9e8186610cb7565b94506001821660008114610ab95760018114610aca57610afd565b60ff19831686528186019350610afd565b610ad385610c86565b60005b83811015610af557815481890152600182019150602081019050610ad6565b838801955050505b50505092915050565b610b0f81610d0f565b82525050565b6000610b22602983610cc2565b9150610b2d82610e26565b604082019050919050565b6000610b45602083610cc2565b9150610b5082610e75565b602082019050919050565b6000610b68828486610a29565b91508190509392505050565b6000610b808284610a87565b915081905092915050565b6000602082019050610ba06000830184610a1a565b92915050565b60006020820190508181036000830152610bc08184610a4e565b905092915050565b6000602082019050610bdd6000830184610b06565b92915050565b60006020820190508181036000830152610bfc81610b15565b9050919050565b60006020820190508181036000830152610c1c81610b38565b9050919050565b60008083356001602003843603038112610c4057610c3f610dfc565b5b80840192508235915067ffffffffffffffff821115610c6257610c61610df2565b5b602083019250600182023603831315610c7e57610c7d610e06565b5b509250929050565b60008190508160005260206000209050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610cde82610ce5565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610d1a82610d21565b9050919050565b6000610d2c82610d33565b9050919050565b6000610d3e82610ce5565b9050919050565b82818337600083830152505050565b60005b83811015610d72578082015181840152602081019050610d57565b83811115610d81576000848401525b50505050565b60006002820490506001821680610d9f57607f821691505b60208210811415610db357610db2610db9565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c6532537465703a2063616c6c6572206973206e6f74207468652060008201527f6e6577206f776e65720000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b610ea781610cd3565b8114610eb257600080fd5b50565b610ebe81610d05565b8114610ec957600080fd5b5056fea2646970667358221220dc5fe29b15afa442da221d4dce506aaf3e811e3636a37d7a6ac81b2d05d1d1b964736f6c63430008070033";

type HanaInteractorMockConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: HanaInteractorMockConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class HanaInteractorMock__factory extends ContractFactory {
  constructor(...args: HanaInteractorMockConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    hanaConnectorAddress: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<HanaInteractorMock> {
    return super.deploy(
      hanaConnectorAddress,
      overrides || {}
    ) as Promise<HanaInteractorMock>;
  }
  override getDeployTransaction(
    hanaConnectorAddress: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(hanaConnectorAddress, overrides || {});
  }
  override attach(address: string): HanaInteractorMock {
    return super.attach(address) as HanaInteractorMock;
  }
  override connect(signer: Signer): HanaInteractorMock__factory {
    return super.connect(signer) as HanaInteractorMock__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): HanaInteractorMockInterface {
    return new utils.Interface(_abi) as HanaInteractorMockInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): HanaInteractorMock {
    return new Contract(address, _abi, signerOrProvider) as HanaInteractorMock;
  }
}
