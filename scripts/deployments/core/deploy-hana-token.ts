import { Contract } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { ZETA_INITIAL_SUPPLY } from "../../../lib/contracts.constants";
import { deployHanaEth, deployHanaNonEth, isEthNetworkName } from "../../../lib/contracts.helpers";

export async function deployHanaToken() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);

  let contract: Contract;

  if (isEthNetworkName(network.name)) {
    contract = await deployHanaEth({
      args: [DEPLOYER_ADDRESS, ZETA_INITIAL_SUPPLY],
    });
  } else {
    contract = await deployHanaNonEth({
      args: [tssAddress, tssUpdaterAddress],
    });
  }

  console.log("Deployed Hana to:", contract.address);
  return contract.address;
}
