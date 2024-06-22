import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { isEthNetworkName } from "../../../lib/contracts.helpers";
import { calculateBestSalt } from "../../../lib/deterministic-deploy.helpers";
import { HanaConnectorEth__factory, HanaConnectorNonEth__factory } from "../../../typechain-types";

const MAX_ITERATIONS = BigNumber.from(100000);

export async function deterministicDeployGetSaltHanaConnector() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const hanaTokenAddress = getAddress("hanaToken", network.name);
  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);

  // @todo: decide which address use as pauser
  const constructorTypes = ["address", "address", "address", "address"];
  const constructorArgs = [hanaTokenAddress, tssAddress, tssUpdaterAddress, tssUpdaterAddress];

  let contractBytecode;

  if (isEthNetworkName(network.name)) {
    contractBytecode = HanaConnectorEth__factory.bytecode;
  } else {
    contractBytecode = HanaConnectorNonEth__factory.bytecode;
  }

  calculateBestSalt(MAX_ITERATIONS, DEPLOYER_ADDRESS, constructorTypes, constructorArgs, contractBytecode);
}

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployGetSaltHanaConnector()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
