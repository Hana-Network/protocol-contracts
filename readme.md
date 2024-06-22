# HanaNetwork Protocol Contracts

This repository contains the smart contracts for HanaNetwork. The smart contracts
are written in Solidity, and the repository includes scripts to compile the
contracts and generate Go bindings.

## Importing Protocol Contracts

As a dApp developer, you can install the protocol contracts package into your
project:

```
yarn add --dev @hananetwork/protocol-contracts
```

Importing
[`HanaInterfaces`](https://www.hana.network/docs/developers/cross-chain-messaging/connector/)
and `HanaInteractor` for cross-chain messaging:

```solidity
import "@hananetwork/protocol-contracts/contracts/evm/interfaces/HanaInterfaces.sol";
import "@hananetwork/protocol-contracts/contracts/evm/tools/HanaInteractor.sol";
```

Importing [HRC20](https://www.hana.network/docs/developers/concepts/hrc-20/)
and the [system
contract](https://www.hana.network/docs/developers/concepts/system-contract/)
for omni-chain smart contracts:

```solidity
import "@hananetwork/protocol-contracts/contracts/hevm/interfaces/IHRC20.sol";
import "@hananetwork/protocol-contracts/contracts/hevm/interfaces/hContract.sol";
import "@hananetwork/protocol-contracts/contracts/hevm/SystemContract.sol";
```

## Prerequisites

Before you can contribute to this project, you must have the following installed:

- [Node.js](https://nodejs.org/)
- [Yarn](https://yarnpkg.com/)
- [jq](https://stedolan.github.io/jq/)
- [abigen](https://geth.ethereum.org/docs/tools/abigen)

## Getting Started

To get started with this project, you should first clone the repository:

```
git clone https://github.com/hana-network/protocol
```

Once you have cloned the repository, you can navigate to the project directory
and run the following command to install the project dependencies:

```
yarn
```

## Compiling Contracts

To compile the contracts, run the following command:

```
yarn compile
```

This will compile the Solidity contracts and output the resulting JSON artifacts
to the `artifacts` directory.

## Generating Go Bindings

To generate Go bindings for the Solidity contracts, run the following command:

```
yarn generate:go
```

This will use `abigen` to generate Go bindings for the contracts and output the
resulting Go files to the `pkg` directory.

## Contributing

If you would like to contribute to this project, please fork the repository and
submit a pull request. All contributions are welcome!
