import { isLocalNetworkName } from "@zetachain/addresses";
import { ethers, network } from "hardhat";

import { isEthNetworkName } from "../../../lib/contracts.helpers";
import { setHanaAddresses } from "../../tools/set-hana-token-addresses";
import { deployHanaToken } from "./deploy-hana-token";
import { deployHanaConnector } from "./deploy-zeta-connector";

async function main() {
  if (isLocalNetworkName(network.name)) {
    const [owner] = await ethers.getSigners();
  }

  const hanaTokenAddress = await deployHanaToken();
  const connectorAddress = await deployHanaConnector();

  /**
   * @description The Eth implementation of Hana token doesn't need any address
   */
  if (isEthNetworkName(network.name)) return;

  /**
   * @description Avoid setting Hana addresses for local network,
   * since it must be done after starting the local Hana node
   */
  if (!isLocalNetworkName(network.name)) {
    await setHanaAddresses(connectorAddress, hanaTokenAddress);
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
