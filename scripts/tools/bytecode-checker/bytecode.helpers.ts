const fetch = require("node-fetch");
const fs = require("fs");
const path = require("path");
const { ethers } = require("ethers");
import { BigNumber } from "ethers";

import {
  BSC_ETHERSCAN_API_ENDPOINT,
  bscEtherscanApiKey,
  ETHERSCAN_API_ENDPOINT,
  etherscanApiKey,
  HANA_NODE_ENDPOINT,
} from "./bytecode.constants";

export const encodeNumber = (weiValue: BigNumber) => {
  return ethers.utils.hexZeroPad(weiValue.toHexString(), 32).replace("0x", "");
};

export const encodeAddress = (address: string) => {
  return ethers.utils.hexZeroPad(address, 32).replace("0x", "");
};

export const compareBytecode = (bytecodeA: string, bytecodeB: string) => {
  if (bytecodeA === bytecodeB) {
    console.log("Bytecode matches!");
  } else {
    console.error("Bytecode doesn't match!");
  }
};

export const getEtherscanBytecode = async (network: "bsc" | "eth", contractAddress: string) => {
  const endpoint = network === "bsc" ? BSC_ETHERSCAN_API_ENDPOINT : ETHERSCAN_API_ENDPOINT;
  const apiKey = network === "bsc" ? bscEtherscanApiKey : etherscanApiKey;
  const etherscanApiUrl = `${endpoint}/api?module=proxy&action=eth_getCode&address=${contractAddress}&apikey=${apiKey}`;

  try {
    const response = await fetch(etherscanApiUrl);
    const data = await response.json();
    const bytecode = data.result;
    return bytecode;
  } catch (error) {
    console.error("Error fetching bytecode:", error);
  }
};

// @dev: helper to find differences between two bytecodes when there's a mismatch
export const findDiff = (codeA: string, codeB: string) => {
  console.log("Lengths:", codeA.length, codeB.length);
  for (let i = 0; i < codeA.length; i++) {
    if (codeA[i] !== codeB[i]) {
      console.log(i, codeA.substring(i - 20, i + 40), codeB.substring(i - 20, i + 40));
    }
  }
};

export const getHanaNodeBytecode = async (contractAddress: string) => {
  try {
    const provider = new ethers.providers.JsonRpcProvider(HANA_NODE_ENDPOINT);
    const bytecode = await provider.getCode(contractAddress);
    return bytecode;
  } catch (error) {
    console.error("Error fetching bytecode:", error);
    console.error("Please make sure you are connected to tailscale.");
  }
};

export const removeImmutableAddress = (bytecode: string, pattern: string) => {
  const replacement = encodeAddress(ethers.constants.AddressZero);
  const regex = new RegExp(pattern, "gi");

  bytecode = bytecode.replace(regex, replacement);
  return bytecode;
};

export const removeImmutableNumber = (bytecode: string, pattern: string) => {
  const replacement = encodeNumber(BigNumber.from("0"));
  const regex = new RegExp(pattern, "gi");

  bytecode = bytecode.replace(regex, replacement);
  return bytecode;
};

export const getDeployedBytecode = async (contract: string, kind: "evm" | "hevm") => {
  try {
    const filePath = path.join(__dirname, `../../../artifacts/contracts/${kind}/${contract}.json`);
    const fileContent = fs.readFileSync(filePath, "utf8");
    const jsonContent = JSON.parse(fileContent);
    const deployedBytecode = jsonContent.deployedBytecode;

    return deployedBytecode;
  } catch (error) {
    console.error("Error reading the file or parsing JSON:", error);
  }
};
