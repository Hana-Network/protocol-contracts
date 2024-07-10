# HanaNetwork Protocol Contracts

This repository houses the HanaNetwork protocol contracts, including Solidity source code, 
generated Go bindings, deployed contract addresses, and helper utilities.

## Importing Protocol Contracts

To incorporate the protocol contracts into your dApp project:

```bash
yarn add --dev @hananetwork/protocol-contracts
```

### Usage Examples

#### Retrieving TSS Address (BSC Testnet)

```typescript
import { getAddress } from "@hananetwork/protocol-contracts";

getAddress("tss", "hana_testnet");
```

#### Fetching HRC-20 BSC USDT Address (HanaNetwork Mainnet Beta)

```typescript
import { getAddress } from "@hananetwork/protocol-contracts";

const usdtAddress = getAddress("hrc20", "hana_mainnet", "USDT.BSC");
```

Note: The third argument (symbol) is specific to HRC-20 address queries.

For a comprehensive list of contract addresses, refer to the [Contract Addresses](https://docs.hana.network/docs/reference/contracts/) documentation.

## Development Prerequisites

Ensure you have the following tools installed:

- [Node.js](https://nodejs.org/)
- [Yarn](https://yarnpkg.com/)
- [jq](https://stedolan.github.io/jq/)
- [abigen](https://geth.ethereum.org/docs/tools/abigen)

## Development Workflow

### Compiling Contracts

To compile Solidity contracts:

```bash
yarn compile
```

This command generates JSON artifacts in the `artifacts` directory.

### Generating Go Bindings and Contract Addresses

To create Go bindings for the Solidity contracts:

```bash
yarn generate
```

This utilizes `abigen` to produce Go files in the `pkg` directory.
