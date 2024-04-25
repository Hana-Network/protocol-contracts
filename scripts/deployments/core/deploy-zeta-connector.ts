import { Contract } from "ethers";
import { network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib/address.tools";

import { deployHanaConnectorEth, deployHanaConnectorNonEth, isEthNetworkName } from "../../../lib/contracts.helpers";

export async function deployHanaConnector() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const hanaTokenAddress = getAddress("hanaToken", network.name);
  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);

  let contract: Contract;
  console.log(`Deploying HanaConnector to ${network.name}`);

  if (isEthNetworkName(network.name)) {
    contract = await deployHanaConnectorEth({
      args: [hanaTokenAddress, tssAddress, tssUpdaterAddress, tssUpdaterAddress],
    });
  } else {
    contract = await deployHanaConnectorNonEth({
      args: [hanaTokenAddress, tssAddress, tssUpdaterAddress, tssUpdaterAddress],
    });
  }

  console.log("Deployed HanaConnector. Address:", contract.address);
  return contract.address;
}
