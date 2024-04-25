import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { ERC20_CUSTODY_HANA_FEE, ERC20_CUSTODY_HANA_MAX_FEE, getSaltNumber } from "../../../lib/contracts.constants";
import {
  deployContractToAddress,
  saltToHex,
} from "../../../lib/ImmutableCreate2Factory/ImmutableCreate2Factory.helpers";
import { ERC20Custody__factory } from "../../../typechain-types";

export const deterministicDeployERC20Custody = async () => {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;
  const initialBalance = await signer.getBalance();

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const hanaTokenAddress = getAddress("hanaToken", network.name);
  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);
  const immutableCreate2FactoryAddress = getAddress("immutableCreate2Factory", network.name);

  const saltNumber = getSaltNumber("hanaERC20Custody", network.name);
  const saltStr = BigNumber.from(saltNumber).toHexString();

  const hanaFee = ERC20_CUSTODY_HANA_FEE;
  const hanaMaxFee = ERC20_CUSTODY_HANA_MAX_FEE;

  const salthex = saltToHex(saltStr, DEPLOYER_ADDRESS);
  console.log("SaltHex:", salthex);

  const constructorTypes = ["address", "address", "uint256", "uint256", "address"];
  const constructorArgs = [tssAddress, tssUpdaterAddress, hanaFee.toString(), hanaMaxFee.toString(), hanaTokenAddress];
  const contractBytecode = ERC20Custody__factory.bytecode;

  const { address } = await deployContractToAddress({
    constructorArgs,
    constructorTypes,
    contractBytecode,
    factoryAddress: immutableCreate2FactoryAddress,
    salt: salthex,
    signer,
  });

  const finalBalance = await signer.getBalance();
  console.log("Deployed ERC20 Custody. Address:", address);
  console.log("Constructor Args", constructorArgs);
  console.log("ETH spent:", initialBalance.sub(finalBalance).toString());

  return address;
};

if (!process.env.EXECUTE_PROGRAMMATICALLY) {
  deterministicDeployERC20Custody()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}
