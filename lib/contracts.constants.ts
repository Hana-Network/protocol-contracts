import { parseEther } from "ethers/lib/utils";

export const HANA_INITIAL_SUPPLY = 2_100_000_000;

export const MAX_ETH_ADDRESS = "0xffffffffffffffffffffffffffffffffffffffff";

// dev: this values should be calculated using get-salt script
const SALT_NUMBERS = {
  baobab_testnet: {
    hanaConnector: "71733",
    hanaConsumer: "0",
    hanaERC20Custody: "195084",
    hanaToken: "29265",
  },
  bsc_mainnet: {
    hanaConnector: "71733",
    hanaConsumer: "0",
    hanaERC20Custody: "195084",
    hanaToken: "29265",
  },
  bsc_testnet: {
    hanaConnector: "71733",
    hanaConsumer: "0",
    hanaERC20Custody: "195084",
    hanaToken: "29265",
  },
  btc_testnet: {
    hanaConnector: "",
    hanaConsumer: "",
    hanaERC20Custody: "",
    hanaToken: "",
  },
  eth_mainnet: {
    hanaConnector: "84286",
    hanaConsumer: "0",
    hanaERC20Custody: "926526",
    hanaToken: "0",
  },
  goerli_testnet: {
    hanaConnector: "1414",
    hanaConsumer: "0",
    hanaERC20Custody: "87967",
    hanaToken: "84108",
  },
  hana_testnet: {
    hanaConnector: "71733",
    hanaConsumer: "0",
    hanaERC20Custody: "195084",
    hanaToken: "29265",
  },
  mumbai_testnet: {
    hanaConnector: "71733",
    hanaConsumer: "0",
    hanaERC20Custody: "195084",
    hanaToken: "29265",
  },
};

export const getSaltNumber = (contractName: string, networkName: string) => {
  const saltNumber = SALT_NUMBERS[networkName][contractName];

  if (!saltNumber) {
    throw new Error(`Salt number for ${contractName} on ${networkName} is not defined.`);
  }

  return saltNumber;
};

export const ERC20_CUSTODY_HANA_FEE = "0";
export const ERC20_CUSTODY_HANA_MAX_FEE = parseEther("1000");
