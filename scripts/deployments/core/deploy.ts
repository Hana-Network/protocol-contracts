import { network } from "hardhat";

import { isProtocolNetworkName } from "../../../lib/address.tools";
import { deterministicDeployERC20Custody } from "./deterministic-deploy-erc20-custody";
import { deterministicDeployHanaConnector } from "./deterministic-deploy-hana-connector";
import { deterministicDeployHanaToken } from "./deterministic-deploy-hana-token";

const networkName = network.name;

async function main() {
  if (!isProtocolNetworkName(networkName)) throw new Error("Invalid network name");

  await deterministicDeployHanaToken();
  await deterministicDeployHanaConnector();
  await deterministicDeployERC20Custody();
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
