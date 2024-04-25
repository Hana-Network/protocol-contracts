import { BigNumber } from "ethers";
import { parseEther } from "ethers/lib/utils";

import {
  BNB_BSC,
  bscConnectorAddress,
  bscERC20CustodyAddress,
  bscTokenAddress,
  BTC_BTC,
  ETH_ETH,
  ethConnectorAddress,
  ethERC20CustodyAddress,
  ethTokenAddress,
} from "./bytecode.constants";
import {
  compareBytecode,
  encodeAddress,
  encodeNumber,
  getDeployedBytecode,
  getEtherscanBytecode,
  getHanaNodeBytecode,
  removeImmutableAddress,
  removeImmutableNumber,
} from "./bytecode.helpers";

const checkEthConnectorBytecode = async () => {
  const remoteBytecode = await getEtherscanBytecode("eth", ethConnectorAddress);
  const cleanRemoteBytecode = removeImmutableAddress(remoteBytecode, encodeAddress(ethTokenAddress));
  const deployedBytecode = await getDeployedBytecode("HanaConnector.eth.sol/HanaConnectorEth", "evm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkEthCustodyBytecode = async () => {
  const remoteBytecode = await getEtherscanBytecode("eth", ethERC20CustodyAddress);
  let cleanRemoteBytecode = removeImmutableAddress(remoteBytecode, encodeAddress(ethTokenAddress));
  cleanRemoteBytecode = removeImmutableNumber(cleanRemoteBytecode, encodeNumber(parseEther("1000"))); // hanaMaxFee
  const deployedBytecode = await getDeployedBytecode("ERC20Custody.sol/ERC20Custody", "evm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkBscConnectorBytecode = async () => {
  const remoteBytecode = await getEtherscanBytecode("bsc", bscConnectorAddress);
  const cleanRemoteBytecode = removeImmutableAddress(remoteBytecode, encodeAddress(bscTokenAddress));
  const deployedBytecode = await getDeployedBytecode("HanaConnector.non-eth.sol/HanaConnectorNonEth", "evm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkBscCustodyBytecode = async () => {
  const remoteBytecode = await getEtherscanBytecode("bsc", bscERC20CustodyAddress);
  let cleanRemoteBytecode = removeImmutableAddress(remoteBytecode, encodeAddress(bscTokenAddress));
  cleanRemoteBytecode = removeImmutableNumber(cleanRemoteBytecode, encodeNumber(parseEther("1000"))); // hanaMaxFee
  const deployedBytecode = await getDeployedBytecode("ERC20Custody.sol/ERC20Custody", "evm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkHRC20ETHBytecode = async () => {
  const remoteBytecode = await getHanaNodeBytecode(ETH_ETH);
  let cleanRemoteBytecode = removeImmutableNumber(remoteBytecode, encodeNumber(BigNumber.from("1"))); // ETH CHAIN ID
  cleanRemoteBytecode = removeImmutableNumber(cleanRemoteBytecode, encodeNumber(BigNumber.from("1"))); // Gas COIN TYPE
  const deployedBytecode = await getDeployedBytecode("HRC20.sol/HRC20", "hevm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkHRC20BTCBytecode = async () => {
  const remoteBytecode = await getHanaNodeBytecode(BTC_BTC);
  let cleanRemoteBytecode = removeImmutableNumber(remoteBytecode, encodeNumber(BigNumber.from("8332"))); // BTC CHAIN ID
  cleanRemoteBytecode = removeImmutableNumber(cleanRemoteBytecode, encodeNumber(BigNumber.from("1"))); // Gas COIN TYPE
  const deployedBytecode = await getDeployedBytecode("HRC20.sol/HRC20", "hevm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkHRC20BSCBytecode = async () => {
  const remoteBytecode = await getHanaNodeBytecode(BNB_BSC);
  let cleanRemoteBytecode = removeImmutableNumber(remoteBytecode, encodeNumber(BigNumber.from("56"))); // BSC CHAIN ID
  cleanRemoteBytecode = removeImmutableNumber(cleanRemoteBytecode, encodeNumber(BigNumber.from("1"))); // Gas COIN TYPE
  const deployedBytecode = await getDeployedBytecode("HRC20.sol/HRC20", "hevm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkBytecode = async () => {
  // ETH
  await checkEthConnectorBytecode();
  await checkEthCustodyBytecode();
  // BSC
  await checkBscConnectorBytecode();
  await checkBscCustodyBytecode();
  // HEVM
  await checkHRC20ETHBytecode();
  await checkHRC20BTCBytecode();
  await checkHRC20BSCBytecode();
};

checkBytecode()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
