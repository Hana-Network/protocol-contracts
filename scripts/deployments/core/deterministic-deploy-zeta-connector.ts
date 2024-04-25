import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { getSaltNumber } from "../../../lib/contracts.constants";
import { isEthNetworkName } from "../../../lib/contracts.helpers";
import {
  deployContractToAddress,
  saltToHex,
} from "../../../lib/ImmutableCreate2Factory/ImmutableCreate2Factory.helpers";
import { HanaConnectorEth__factory, HanaConnectorNonEth__factory } from "../../../typechain-types";

export const deterministicDeployHanaConnector = async () => {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;
  const initialBalance = await signer.getBalance();

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const hanaTokenAddress = getAddress("hanaToken", network.name);
  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);
  const immutableCreate2FactoryAddress = getAddress("immutableCreate2Factory", network.name);

  const saltNumber = getSaltNumber("hanaConnector", network.name);
  const saltStr = BigNumber.from(saltNumber).toHexString();

  const salthex = saltToHex(saltStr, DEPLOYER_ADDRESS);
  const constructorTypes = ["address", "address", "address", "address"];
  const constructorArgs = [hanaTokenAddress, tssAddress, tssUpdaterAddress, tssUpdaterAddress];

  let contractBytecode;
  if (isEthNetworkName(network.name)) {
    contractBytecode = HanaConnectorEth__factory.bytecode;
  } else {
    contractBytecode = HanaConnectorNonEth__factory.bytecode;
  }

  const { address } = await deployContractToAddress({
    constructorArgs,
    constructorTypes,
    contractBytecode,
    factoryAddress: immutableCreate2FactoryAddress,
    salt: salthex,
    signer,
  });

  const finalBalance = await signer.getBalance();
  console.log("Deployed HanaConnector. Address:", address);
  console.log("Constructor Args", constructorArgs);
  console.log("ETH spent:", initialBalance.sub(finalBalance).toString());

  return address;
};

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployHanaConnector()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
