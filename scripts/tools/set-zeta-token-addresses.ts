import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { HanaNonEth__factory as HanaNonEthFactory } from "../../typechain-types";

export async function setHanaAddresses(connectorAddress: string, hanaTokenAddress: string) {
  const [owner] = await ethers.getSigners();

  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);

  if (owner.address !== tssAddress && owner.address !== tssUpdaterAddress) {
    console.log("Only TSS or TSS Updater can set Hana addresses.");
    console.log("Please execute this step with a valid account.");
    return;
  }

  const [tssSigner] = await ethers.getSigners();

  const contract = HanaNonEthFactory.connect(hanaTokenAddress, tssSigner);

  console.log("Updating");
  console.log("connectorAddress", connectorAddress);
  console.log("hanaTokenAddress", hanaTokenAddress);
  const tx = await contract.updateTssAndConnectorAddresses(tssAddress, connectorAddress);
  await tx.wait();
  console.log("Updated");
}
