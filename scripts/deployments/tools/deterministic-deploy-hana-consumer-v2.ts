import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, getNonHanaAddress, isProtocolNetworkName } from "lib";

import { ZETA_CONSUMER_SALT_NUMBER } from "../../../lib/contracts.constants";
import {
  deployContractToAddress,
  saltToHex,
} from "../../../lib/ImmutableCreate2Factory/ImmutableCreate2Factory.helpers";
import { HanaTokenConsumerUniV2__factory } from "../../../typechain-types";

export async function deterministicDeployHanaConsumer() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const hanaTokenAddress = getAddress("hanaToken", network.name);
  const immutableCreate2FactoryAddress = getAddress("immutableCreate2Factory", network.name);

  const uniswapV2Router02 = getNonHanaAddress("uniswapV2Router02", network.name);

  const saltNumber = ZETA_CONSUMER_SALT_NUMBER;
  const saltStr = BigNumber.from(saltNumber).toHexString();

  const salthex = saltToHex(saltStr, DEPLOYER_ADDRESS);
  const constructorTypes = ["address", "address"];
  const constructorArgs = [hanaTokenAddress, uniswapV2Router02];

  const contractBytecode = HanaTokenConsumerUniV2__factory.bytecode;

  const { address } = await deployContractToAddress({
    constructorArgs,
    constructorTypes,
    contractBytecode,
    factoryAddress: immutableCreate2FactoryAddress,
    salt: salthex,
    signer,
  });

  console.log("Deployed HanaConsumer. Address:", address);
}

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployHanaConsumer()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
