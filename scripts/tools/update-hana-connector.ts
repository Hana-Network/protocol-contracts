import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { getHanaFactoryNonEth } from "../../lib/contracts.helpers";

async function updateHanaConnector() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const [, tssUpdaterSigner] = await ethers.getSigners();

  const hanaTokenAddress = getAddress("hanaToken", network.name);
  const tssAddress = getAddress("tss", network.name);
  const connectorAddress = getAddress("connector", network.name);

  const contract = (
    await getHanaFactoryNonEth({ deployParams: null, existingContractAddress: hanaTokenAddress })
  ).connect(tssUpdaterSigner);

  await (await contract.updateTssAndConnectorAddresses(tssAddress, connectorAddress)).wait();

  console.log(`Updated TSS address to ${tssAddress}.`);
  console.log(`Updated Connector address to ${connectorAddress}.`);
}

updateHanaConnector()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
