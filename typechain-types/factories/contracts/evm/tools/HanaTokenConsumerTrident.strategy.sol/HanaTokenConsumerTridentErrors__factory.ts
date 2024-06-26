/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  HanaTokenConsumerTridentErrors,
  HanaTokenConsumerTridentErrorsInterface,
} from "../../../../../contracts/evm/tools/HanaTokenConsumerTrident.strategy.sol/HanaTokenConsumerTridentErrors";

const _abi = [
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
    name: "ReentrancyError",
    type: "error",
  },
] as const;

export class HanaTokenConsumerTridentErrors__factory {
  static readonly abi = _abi;
  static createInterface(): HanaTokenConsumerTridentErrorsInterface {
    return new utils.Interface(_abi) as HanaTokenConsumerTridentErrorsInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): HanaTokenConsumerTridentErrors {
    return new Contract(
      address,
      _abi,
      signerOrProvider
    ) as HanaTokenConsumerTridentErrors;
  }
}
