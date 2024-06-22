import { ethers, network } from "hardhat";
import { getAddress, getNonHanaAddress, isProtocolNetworkName } from "lib";

import { HanaTokenConsumerUniV3__factory } from "../../../typechain-types";

export async function deterministicDeployHanaConsumer() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }
  const accounts = await ethers.getSigners();
  const [signer] = accounts;

  const hanaTokenAddress = getAddress("hanaToken", network.name);

  const uniswapV3Router = getNonHanaAddress("uniswapV3Router", network.name);
  const uniswapV3Factory = getNonHanaAddress("uniswapV3Factory", network.name);
  const WETH9Address = getNonHanaAddress("weth9", network.name);

  const hanaPoolFee = 500;
  const tokenPoolFee = 3000;

  console.log([hanaTokenAddress, uniswapV3Router, uniswapV3Factory, WETH9Address, hanaPoolFee, tokenPoolFee]);

  const Factory = new HanaTokenConsumerUniV3__factory(signer);
  const contract = await Factory.deploy(
    hanaTokenAddress,
    uniswapV3Router,
    uniswapV3Factory,
    WETH9Address,
    hanaPoolFee,
    tokenPoolFee
  );
  await contract.deployed();
  const address = contract.address;

  console.log("Deployed HanaConsumer. Address:", address);
}

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployHanaConsumer()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
