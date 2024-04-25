import uniswapV2Router from "@uniswap/v2-periphery/build/IUniswapV2Router02.json";
import SwapRouter from "@uniswap/v3-periphery/artifacts/contracts/SwapRouter.sol/SwapRouter.json";
import { getEndpoints } from "@zetachain/networks";
import axios, { AxiosResponse } from "axios";
import { task } from "hardhat/config";
import { HardhatRuntimeEnvironment } from "hardhat/types";
import { isEqual } from "lodash";

import { HanaConnectorBase__factory } from "../typechain-types";
import { ERC20Custody__factory } from "../typechain-types/factories/contracts/evm/ERC20Custody__factory";
import { SystemContract__factory } from "../typechain-types/factories/contracts/hevm/SystemContract.sol/SystemContract__factory";
import hana_mainnet_addresses from "./addresses.mainnet.json";
import hana_testnet_addresses from "./addresses.testnet.json";

declare const hre: any;

type Network = "zeta_mainnet" | "zeta_testnet";
// type Network = "hana_mainnet" | "hana_testnet";

type AddressDetails = {
  address: string;
  category: string;
  chain_id: number;
  chain_name: string;
  type: string;
};

const api = {
  zeta_mainnet: {
    evm: "https://zetachain-evm.blockpi.network/v1/rpc/public",
    rpc: "https://zetachain.blockpi.network/lcd/v1/public",
  },
  // hana_mainnet: {
  //   evm: "https://hananetwork-evm.blockpi.network/v1/rpc/public",
  //   rpc: "https://hananetwork.blockpi.network/lcd/v1/public",
  // },
  zeta_testnet: {
    evm: "https://zetachain-athens-evm.blockpi.network/v1/rpc/public",
    rpc: "https://zetachain-athens.blockpi.network/lcd/v1/public",
  },
  // hana_testnet: {
  //   evm: "https://hananetwork-sakura-evm.blockpi.network/v1/rpc/public",
  //   rpc: "https://hananetwork-sakura.blockpi.network/lcd/v1/public",
  // },
};

const fetchChains = async (network: Network) => {
  const URL = `${api[network].rpc}/zeta-chain/observer/supportedChains`;
  try {
    const response: AxiosResponse<any> = await axios.get(URL);

    if (response.status === 200) {
      return response.data.chains;
    } else {
      console.error("Error fetching chains:", response.status, response.statusText);
    }
  } catch (error) {
    console.error("Error fetching chains:", error);
  }
};

const fetchTssData = async (chains: any, addresses: any, network: Network) => {
  const bitcoinChainID = network === "zeta_mainnet" ? "8332" : "18332";
  const URL = `${api[network].rpc}/zeta-chain/observer/get_tss_address/${bitcoinChainID}`;
  try {
    const tssResponse: AxiosResponse<any> = await axios.get(URL);

    if (tssResponse.status === 200) {
      chains.forEach((chain: any) => {
        const { btc, eth } = tssResponse.data;
        if (chain.chain_name === "zeta_testnet") return;
        addresses.push({
          address: ["btc_testnet", "btc_mainnet"].includes(chain.chain_name) ? btc : eth,
          category: "omnichain",
          chain_id: parseInt(chain.chain_id),
          chain_name: chain.chain_name,
          type: "tss",
        });
      });
    } else {
      console.error("Error fetching TSS data:", tssResponse.status, tssResponse.statusText);
    }
  } catch (error) {
    console.error("Error fetching TSS data:", error);
  }
};

const fetchSystemContract = async (addresses: any, network: Network) => {
  const chain_id = network === "zeta_mainnet" ? 8787 : 8788;
  const URL = `${api[network].rpc}/zeta-chain/fungible/system_contract`;
  try {
    const systemContractResponse: AxiosResponse<any> = await axios.get(URL);

    if (systemContractResponse.status === 200) {
      addresses.push({
        address: systemContractResponse.data.SystemContract.system_contract,
        category: "omnichain",
        chain_id,
        chain_name: network,
        type: "systemContract",
      });
      addresses.push({
        address: systemContractResponse.data.SystemContract.connector_zevm,
        category: "messaging",
        chain_id,
        chain_name: network,
        type: "connector",
      });
    } else {
      console.error("Error fetching system contract:", systemContractResponse.statusText);
    }
  } catch (error) {
    console.error("Error fetching system contract:", error);
  }
};

const fetchForeignCoinsData = async (chains: any, addresses: any, network: Network) => {
  const chain_id = network === "zeta_mainnet" ? 8787 : 8788;
  const URL = `${api[network].rpc}/zeta-chain/fungible/foreign_coins`;
  try {
    const foreignCoinsResponse: AxiosResponse<any> = await axios.get(URL);
    if (foreignCoinsResponse.status === 200) {
      foreignCoinsResponse.data.foreignCoins.forEach((token: any) => {
        addresses.push({
          address: token.zrc20_contract_address,
          asset: token.asset,
          category: "omnichain",
          chain_id,
          chain_name: network,
          coin_type: token.coin_type.toLowerCase(),
          decimals: 18,
          description: token.name,
          foreign_chain_id: token.foreign_chain_id,
          symbol: token.symbol,
          type: "zrc20",
        });
      });
    } else {
      console.error("Error fetching foreign coins data:", foreignCoinsResponse.status, foreignCoinsResponse.statusText);
    }
  } catch (error) {
    console.error("Error fetching foreign coins data:", error);
  }
};

const fetchAthensAddresses = async (addresses: any, hre: any, network: Network) => {
  const chain_id = network === "zeta_mainnet" ? 8787 : 8788;
  const systemContract = addresses.find((a: any) => {
    return a.chain_name === network && a.type === "systemContract";
  })?.address;
  const provider = new hre.ethers.providers.JsonRpcProvider(api[network].evm);
  const sc = SystemContract__factory.connect(systemContract, provider);
  const common = {
    category: "omnichain",
    chain_id,
    chain_name: network,
  };
  try {
    addresses.push({ ...common, address: await sc.uniswapv2FactoryAddress(), type: "uniswapV2Factory" });
    addresses.push({ ...common, address: await sc.wHanaContractAddress(), type: "zetaToken" });
    addresses.push({ ...common, address: await sc.uniswapv2Router02Address(), type: "uniswapV2Router02" });
    // addresses.push({ ...common, address: await sc.zetaConnectorZEVMAddress(), type: "zetaConnectorZEVM" });
    addresses.push({ ...common, address: await sc.FUNGIBLE_MODULE_ADDRESS(), type: "fungibleModule" });
  } catch (error) {
    console.error("Error fetching addresses from ZetaChain:", error);
  }
};

const fetchChainSpecificAddresses = async (chains: any, addresses: any, network: Network) => {
  await Promise.all(
    chains.map(async (chain: any) => {
      return axios
        .get(`${api[network].rpc}/zeta-chain/observer/get_chain_params_for_chain/${chain.chain_id}`)
        .then(({ data }) => {
          const hanaToken = data.chain_params.hana_token_contract_address;
          if (hanaToken && hanaToken != "0x0000000000000000000000000000000000000000") {
            addresses.push({
              address: hanaToken,
              category: "messaging",
              chain_id: parseInt(chain.chain_id),
              chain_name: chain.chain_name,
              type: "hanaToken",
            });
          }
          const connector = data.chain_params.connector_contract_address;
          if (connector && connector != "0x0000000000000000000000000000000000000000") {
            addresses.push({
              address: connector,
              category: "messaging",
              chain_id: parseInt(chain.chain_id),
              chain_name: chain.chain_name,
              type: "connector",
            });
          }
          const erc20Custody = data.chain_params.erc20_custody_contract_address;
          if (erc20Custody && erc20Custody != "0x0000000000000000000000000000000000000000") {
            addresses.push({
              address: data.chain_params.erc20_custody_contract_address,
              category: "omnichain",
              chain_id: parseInt(chain.chain_id),
              chain_name: chain.chain_name,
              type: "erc20Custody",
            });
          }
        });
    })
  );
};

const fetchTSSUpdater = async (chains: any, addresses: any) => {
  await Promise.all(
    chains.map(async (chain: any) => {
      const erc20Custody = addresses.find((a: any) => {
        return a.chain_name === chain.chain_name && a.type === "erc20Custody";
      })?.address;
      if (erc20Custody) {
        if (["18332", "8332"].includes(chain.chain_id)) return;
        const rpc = getEndpoints("evm", chain.chain_name)[0]?.url;
        const provider = new hre.ethers.providers.JsonRpcProvider(rpc);
        const custody = ERC20Custody__factory.connect(erc20Custody, provider);
        return custody.TSSAddressUpdater().then((address: string) => {
          addresses.push({
            address,
            category: "omnichain",
            chain_id: parseInt(chain.chain_id),
            chain_name: chain.chain_name,
            type: "tssUpdater",
          });
        });
      }
    })
  );
};

const fetchPauser = async (chains: any, addresses: any) => {
  await Promise.all(
    chains.map(async (chain: any) => {
      const erc20Custody = addresses.find((a: any) => {
        return a.chain_name === chain.chain_name && a.type === "connector";
      })?.address;
      if (erc20Custody) {
        if (["18332", "8332", "8788", "8787"].includes(chain.chain_id)) return;
        const rpc = getEndpoints("evm", chain.chain_name)[0]?.url;
        const provider = new hre.ethers.providers.JsonRpcProvider(rpc);
        const connector = HanaConnectorBase__factory.connect(erc20Custody, provider);
        return connector.pauserAddress().then((address: string) => {
          addresses.push({
            address,
            category: "messaging",
            chain_id: parseInt(chain.chain_id),
            chain_name: chain.chain_name,
            type: "pauser",
          });
        });
      }
    })
  );
};

const fetchFactoryV2 = async (addresses: any, hre: any, network: Network) => {
  const routers = addresses.filter((a: any) => a.type === "uniswapV2Router02");

  for (const router of routers) {
    const rpc = getEndpoints("evm", router.chain_name)[0]?.url;
    const provider = new hre.ethers.providers.JsonRpcProvider(rpc);
    const routerContract = new hre.ethers.Contract(router.address, uniswapV2Router.abi, provider);

    try {
      const wethAddress = await routerContract.WETH();
      const factoryAddress = await routerContract.factory();

      // Skip HanaNetwork as we've already added HANA token
      if (router.chain_id !== 8787 && router.chain_id !== 8788) {
        addresses.push({
          address: wethAddress,
          category: "messaging",
          chain_id: router.chain_id,
          chain_name: router.chain_name,
          type: "weth9",
        });
      }

      addresses.push({
        address: factoryAddress,
        category: "messaging",
        chain_id: router.chain_id,
        chain_name: router.chain_name,
        type: "uniswapV2Factory",
      });
    } catch (error) {
      console.error(`Error fetching factory and WETH for router v2 ${router.address}:`, error);
    }
  }
};

const fetchFactoryV3 = async (addresses: any, hre: any, network: Network) => {
  const routers = addresses.filter((a: any) => a.type === "uniswapV3Router");

  for (const router of routers) {
    const rpc = getEndpoints("evm", router.chain_name)[0]?.url;
    const provider = new hre.ethers.providers.JsonRpcProvider(rpc);
    const routerContract = new hre.ethers.Contract(router.address, SwapRouter.abi, provider);

    try {
      const wethAddress = await routerContract.WETH9();
      const factoryAddress = await routerContract.factory();

      const wethObj = {
        address: wethAddress,
        category: "messaging",
        chain_id: router.chain_id,
        chain_name: router.chain_name,
        type: "weth9",
      };

      if (!addresses.some((e: any) => isEqual(e, wethObj))) {
        addresses.push(wethObj);
      }

      addresses.push({
        address: factoryAddress,
        category: "messaging",
        chain_id: router.chain_id,
        chain_name: router.chain_name,
        type: "uniswapV3Factory",
      });
    } catch (error) {
      console.error(`Error fetching factory for router v3 ${router.address}:`, error);
    }
  }
};

const main = async (args: any, hre: HardhatRuntimeEnvironment) => {
  let addresses: any = [];

  const n = hre.network.name;
  if (n === "zeta_testnet") {
    addresses.push(...hana_testnet_addresses);
  } else if (n === "zeta_mainnet") {
    addresses.push(...hana_mainnet_addresses);
  } else {
    throw new Error(`Unsupported network: ${n}. Must be 'zeta_testnet' or 'zeta_mainnet'.`);
  }

  const network: Network = n;

  const chains = await fetchChains(network);
  await fetchTssData(chains, addresses, network);
  await fetchSystemContract(addresses, network);
  await fetchForeignCoinsData(chains, addresses, network);
  await fetchAthensAddresses(addresses, hre, network);
  await fetchChainSpecificAddresses(chains, addresses, network);
  await fetchTSSUpdater(chains, addresses);
  await fetchPauser(chains, addresses);
  await fetchFactoryV2(addresses, hre, network);
  await fetchFactoryV3(addresses, hre, network);

  addresses = addresses.sort((a: AddressDetails, b: AddressDetails) => {
    if (a.chain_id !== b.chain_id) {
      return a.chain_id - b.chain_id;
    } else if (a.type !== b.type) {
      return a.type.localeCompare(b.type);
    } else {
      return a.address.localeCompare(b.address);
    }
  });

  console.log(JSON.stringify(addresses, null, 2));
};

task("addresses", "").setAction(main);
