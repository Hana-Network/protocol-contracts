# HanaNetwork Protocol Contracts

This repository contains HanaNetwork protocol contracts: Solidity source code,
generated Go bindings, deployed contract addresses and helper utilities.

## Importing Protocol Contracts

As a dApp developer, you can install the protocol contracts package into your
project:

```
yarn add --dev @zetachain/protocol-contracts
```

Getting the TSS address on BSC testnet:

```ts
import { getAddress } from "@zetachain/protocol-contracts";

getAddress("tss", "hana_testnet");
```

Getting a ZRC-20 BSC USDT on HanaNetwork Mainnet Beta:

```ts
import { getAddress } from "@zetachain/protocol-contracts";

getAddress("hrc20", "hana_mainnet", "USDT.BSC");
```

The third argument (symbol) is only used when querying ZRC-20 addresses to
specify which token address is needed.

To view a table of all contracts visit the [Contract Addresses](https://www.hana.network/docs/reference/contracts/) page in the docs.

Importing
[`HanaInterfaces`](https://www.hana.network/docs/developers/cross-chain-messaging/connector/)
and `HanaInteractor` for cross-chain messaging:

```solidity
import "@zetachain/protocol-contracts/contracts/evm/interfaces/HanaInterfaces.sol";
import "@zetachain/protocol-contracts/contracts/evm/tools/HanaInteractor.sol";
```

Importing [HRC20](https://www.hana.network/docs/developers/tokens/hrc20/)
and the [system
contract](https://www.hana.network/docs/developers/omnichain/system-contract/)
for omni-chain smart contracts:

```solidity
import "@zetachain/protocol-contracts/contracts/hevm/interfaces/IHRC20.sol";
import "@zetachain/protocol-contracts/contracts/hevm/interfaces/hContract.sol";
import "@zetachain/protocol-contracts/contracts/hevm/SystemContract.sol";
```

## Prerequisites for Development

Before you can contribute to this project, you must have the following installed:

- [Node.js](https://nodejs.org/)
- [Yarn](https://yarnpkg.com/)
- [jq](https://stedolan.github.io/jq/)
- [abigen](https://geth.ethereum.org/docs/tools/abigen)

## Compiling Contracts

To compile the contracts, run the following command:

```
yarn compile
```

This will compile the Solidity contracts and output the resulting JSON artifacts
to the `artifacts` directory.

## Generating Go Bindings and Contract Addresses

To generate Go bindings for the Solidity contracts and fetch, run the following command:

```
yarn generate
```

This will use `abigen` to generate Go bindings for the contracts and output the
resulting Go files to the `pkg` directory.

## Contributing

If you would like to contribute to this project, please fork the repository and
submit a pull request. All contributions are welcome!
