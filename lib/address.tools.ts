import addresses from "../data/addresses.json";

export declare type HanaProtocolAddress =
  | "connector"
  | "erc20Custody"
  | "hanaToken"
  | "hanaTokenConsumerUniV2"
  | "hanaTokenConsumerUniV3"
  | "immutableCreate2Factory"
  | "tss"
  | "tssUpdater";

export const hanaProtocolAddress: HanaProtocolAddress[] = [
  "connector",
  "erc20Custody",
  "immutableCreate2Factory",
  "tss",
  "tssUpdater",
  "hanaToken",
  "hanaTokenConsumerUniV2",
  "hanaTokenConsumerUniV3",
];
export const isHanaProtocolAddress = (str: string): str is HanaProtocolAddress =>
  hanaProtocolAddress.includes(str as HanaProtocolAddress);

export declare type HanaHEVMAddress =
  | "fungibleModule"
  | "hrc20"
  | "systemContract"
  | "uniswapv2Factory"
  | "uniswapv2Router02";

export declare type HanaProtocolTestNetwork =
  | "baobab_testnet"
  | "bsc_testnet"
  | "btc_testnet"
  | "goerli_testnet"
  | "hana_testnet"
  | "mumbai_testnet";

export const hanaProtocolTestNetworks: HanaProtocolTestNetwork[] = [
  "baobab_testnet",
  "bsc_testnet",
  "btc_testnet",
  "goerli_testnet",
  "mumbai_testnet",
  "hana_testnet",
];

export declare type NonHanaAddress =
  | "uniswapV2Factory"
  | "uniswapV2Router02"
  | "uniswapV3Factory"
  | "uniswapV3Router"
  | "weth9";

export const nonHanaAddress: NonHanaAddress[] = [
  "uniswapV2Factory",
  "uniswapV2Router02",
  "uniswapV3Router",
  "uniswapV3Factory",
  "weth9",
];

export declare type HanaProtocolMainNetwork = "bsc_mainnet" | "eth_mainnet" | "hana_mainnet";
export const hanaProtocolMainNetworks: HanaProtocolMainNetwork[] = ["eth_mainnet", "bsc_mainnet", "hana_mainnet"];

export declare type HanaProtocolNetwork = HanaProtocolMainNetwork | HanaProtocolTestNetwork;
export const hanaProtocolNetworks: HanaProtocolNetwork[] = [...hanaProtocolTestNetworks, ...hanaProtocolMainNetworks];

export declare type HanaProtocolEnviroment = "mainnet" | "testnet";

export const isProtocolNetworkName = (str: string): str is HanaProtocolNetwork =>
  hanaProtocolNetworks.includes(str as HanaProtocolNetwork);

export const isTestnetNetwork = (network: HanaProtocolTestNetwork): boolean => {
  return hanaProtocolTestNetworks.includes(network);
};

export const isMainnetNetwork = (network: HanaProtocolTestNetwork): boolean => {
  return false;
};

// export const getAddress = (address: HanaHEVMAddress | HanaProtocolAddress, network: HanaProtocolNetwork): string => {
//   if (isHanaProtocolAddress(address)) {
//     return (addresses["ccm"] as any)[network][address];
//   }

//   return (addresses["hevm"] as any)[network][address];
// };

// export const getHRC20Address = (network: HanaProtocolNetwork): string => {
//   return (addresses["hevm"] as any)[network]["hrc20"];
// };

export const getNonHanaAddress = (address: NonHanaAddress, network: HanaProtocolNetwork): string => {
  return (addresses["non_hana"] as any)[network][address];
};
