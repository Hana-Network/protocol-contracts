import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { getSaltNumber, HANA_INITIAL_SUPPLY } from "../../../lib/contracts.constants";
import { isEthNetworkName } from "../../../lib/contracts.helpers";
import {
  deployContractToAddress,
  saltToHex,
} from "../../../lib/ImmutableCreate2Factory/ImmutableCreate2Factory.helpers";
import { HanaEth__factory, HanaNonEth__factory } from "../../../typechain-types";

export const deterministicDeployHanaToken = async () => {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;
  const initialBalance = await signer.getBalance();

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);
  const immutableCreate2FactoryAddress = getAddress("immutableCreate2Factory", network.name);

  const saltNumber = getSaltNumber("hanaToken", network.name);
  const saltStr = BigNumber.from(saltNumber).toHexString();

  const salthex = saltToHex(saltStr, DEPLOYER_ADDRESS);
  console.log("SaltHex:", salthex);

  let constructorTypes;
  let constructorArgs;
  let contractBytecode;

  if (isEthNetworkName(network.name)) {
    constructorTypes = ["address", "uint256"];
    constructorArgs = [DEPLOYER_ADDRESS, HANA_INITIAL_SUPPLY.toString()];
    contractBytecode = HanaEth__factory.bytecode;
  } else {
    constructorTypes = ["address", "address"];
    constructorArgs = [tssAddress, tssUpdaterAddress];
    contractBytecode = HanaNonEth__factory.bytecode;
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
  console.log("Deployed hanaToken. Address:", address);
  console.log("Constructor Args", constructorArgs);
  console.log("ETH spent:", initialBalance.sub(finalBalance).toString());

  return address;
};

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployHanaToken()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
