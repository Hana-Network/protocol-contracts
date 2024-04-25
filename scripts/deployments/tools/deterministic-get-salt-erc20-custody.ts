import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";
import { ERC20_CUSTODY_HANA_FEE, ERC20_CUSTODY_HANA_MAX_FEE } from "lib/contracts.constants";

import { calculateBestSalt } from "../../../lib/deterministic-deploy.helpers";
import { ERC20Custody__factory } from "../../../typechain-types";

const MAX_ITERATIONS = BigNumber.from(1000000);

export const deterministicDeployGetSaltERC20Custody = async () => {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const hanaTokenAddress = getAddress("hanaToken", network.name);
  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);

  const hanaFee = ERC20_CUSTODY_HANA_FEE;
  const hanaMaxFee = ERC20_CUSTODY_HANA_MAX_FEE;

  const constructorTypes = ["address", "address", "uint256", "uint256", "address"];
  const constructorArgs = [tssAddress, tssUpdaterAddress, hanaFee.toString(), hanaMaxFee.toString(), hanaTokenAddress];
  const contractBytecode = ERC20Custody__factory.bytecode;

  await calculateBestSalt(
    MAX_ITERATIONS,
    DEPLOYER_ADDRESS,
    constructorTypes,
    constructorArgs,
    contractBytecode,
    network.name
  );
};

deterministicDeployGetSaltERC20Custody()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
