import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { HANA_INITIAL_SUPPLY } from "../../../lib/contracts.constants";
import { isEthNetworkName } from "../../../lib/contracts.helpers";
import { calculateBestSalt } from "../../../lib/deterministic-deploy.helpers";
import { HanaEth__factory, HanaNonEth__factory } from "../../../typechain-types";

const MAX_ITERATIONS = BigNumber.from(1000000);

export async function deterministicDeployGetSaltHanaToken() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);

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
  await calculateBestSalt(
    MAX_ITERATIONS,
    DEPLOYER_ADDRESS,
    constructorTypes,
    constructorArgs,
    contractBytecode,
    network.name
  );
}

deterministicDeployGetSaltHanaToken()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
