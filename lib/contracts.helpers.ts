import {
  HanaConnectorBase,
  HanaConnectorBase__factory as HanaConnectorBaseFactory,
  HanaConnectorEth,
  HanaConnectorEth__factory as HanaConnectorEthFactory,
  HanaConnectorNonEth,
  HanaConnectorNonEth__factory as HanaConnectorNonEthFactory,
  HanaEth,
  HanaEth__factory as HanaEthFactory,
  HanaInteractorMock,
  HanaInteractorMock__factory as HanaInteractorMockFactory,
  HanaNonEth,
  HanaNonEth__factory as HanaNonEthFactory,
  HanaReceiverMock,
  HanaReceiverMock__factory as HanaReceiverMockFactory,
  HanaTokenConsumerUniV2,
  HanaTokenConsumerUniV2__factory as HanaTokenConsumerUniV2Factory,
  HanaTokenConsumerUniV3,
  HanaTokenConsumerUniV3__factory as HanaTokenConsumerUniV3Factory,
  ImmutableCreate2Factory,
  ImmutableCreate2Factory__factory,
} from "@typechain-types";
import { BaseContract, ContractFactory } from "ethers";
import { ethers } from "hardhat";

export const isEthNetworkName = (networkName: string) =>
  networkName === "eth-localnet" || networkName === "goerli" || networkName === "eth-mainnet";

export const deployHanaConnectorBase = async ({ args }: { args: Parameters<HanaConnectorBaseFactory["deploy"]> }) => {
  const Factory = (await ethers.getContractFactory("HanaConnectorBase")) as HanaConnectorBaseFactory;

  const hanaConnectorContract = (await Factory.deploy(...args)) as HanaConnectorBase;

  await hanaConnectorContract.deployed();

  return hanaConnectorContract;
};

export const deployHanaConnectorEth = async ({ args }: { args: Parameters<HanaConnectorEthFactory["deploy"]> }) => {
  const Factory = (await ethers.getContractFactory("HanaConnectorEth")) as HanaConnectorEthFactory;

  const hanaConnectorContract = (await Factory.deploy(...args)) as HanaConnectorEth;

  await hanaConnectorContract.deployed();

  return hanaConnectorContract;
};

export const deployHanaConnectorNonEth = async ({
  args,
}: {
  args: Parameters<HanaConnectorNonEthFactory["deploy"]>;
}) => {
  const Factory = (await ethers.getContractFactory("HanaConnectorNonEth")) as HanaConnectorNonEthFactory;

  const hanaConnectorContract = (await Factory.deploy(...args)) as HanaConnectorNonEth;

  await hanaConnectorContract.deployed();

  return hanaConnectorContract;
};

export const deployHanaReceiverMock = async () => {
  const Factory = (await ethers.getContractFactory("HanaReceiverMock")) as HanaReceiverMockFactory;

  const hanaReceiverMock = (await Factory.deploy()) as HanaReceiverMock;

  await hanaReceiverMock.deployed();

  return hanaReceiverMock;
};

export const deployHanaEth = async ({ args }: { args: Parameters<HanaEthFactory["deploy"]> }) => {
  const Factory = (await ethers.getContractFactory("HanaEth")) as HanaEthFactory;

  const hanaEthContract = (await Factory.deploy(...args)) as HanaEth;

  await hanaEthContract.deployed();

  return hanaEthContract;
};

export const deployHanaNonEth = async ({ args }: { args: Parameters<HanaNonEthFactory["deploy"]> }) => {
  const Factory = (await ethers.getContractFactory("HanaNonEth")) as HanaNonEthFactory;

  const hanaNonEthContract = (await Factory.deploy(...args)) as HanaNonEth;

  await hanaNonEthContract.deployed();

  return hanaNonEthContract;
};

export const deployImmutableCreate2Factory = async () => {
  const Factory = (await ethers.getContractFactory("ImmutableCreate2Factory")) as ImmutableCreate2Factory__factory;

  const ImmutableCreate2FactoryContract = (await Factory.deploy()) as ImmutableCreate2Factory;

  await ImmutableCreate2FactoryContract.deployed();

  return ImmutableCreate2FactoryContract;
};

export const getHanaConnectorEth = async (params: GetContractParams<HanaConnectorEthFactory>) =>
  getContract<HanaConnectorEthFactory, HanaConnectorEth>({
    contractName: "HanaConnectorEth",
    ...params,
  });

export const getHanaConnectorNonEth = async (params: GetContractParams<HanaConnectorNonEthFactory>) =>
  getContract<HanaConnectorNonEthFactory, HanaConnectorNonEth>({
    contractName: "HanaConnectorNonEth",
    ...params,
  });

export const getHanaFactoryNonEth = async (params: GetContractParams<HanaNonEthFactory>) =>
  await getContract<HanaNonEthFactory, HanaNonEth>({
    contractName: "HanaNonEth",
    ...params,
  });

export const getHanaFactoryEth = async (params: GetContractParams<HanaEthFactory>) =>
  await getContract<HanaEthFactory, HanaEth>({
    contractName: "HanaNonEth",
    ...params,
  });

export const getHanaInteractorMock = async (hanaToken: string) =>
  getContract<HanaInteractorMockFactory, HanaInteractorMock>({
    contractName: "HanaInteractorMock",
    deployParams: [hanaToken],
  });

export const getHanaTokenConsumerUniV2Strategy = async (params: GetContractParams<HanaTokenConsumerUniV2Factory>) =>
  getContract<HanaTokenConsumerUniV2Factory, HanaTokenConsumerUniV2>({
    contractName: "HanaTokenConsumerUniV2",
    ...params,
  });

export const getHanaTokenConsumerUniV3Strategy = async (params: GetContractParams<HanaTokenConsumerUniV3Factory>) =>
  getContract<HanaTokenConsumerUniV3Factory, HanaTokenConsumerUniV3>({
    contractName: "HanaTokenConsumerUniV3",
    ...params,
  });

export type GetContractParams<Factory extends ContractFactory> =
  | {
      deployParams: Parameters<Factory["deploy"]>;
      existingContractAddress?: null;
    }
  | {
      deployParams?: null;
      existingContractAddress: string;
    };

export const getContract = async <Factory extends ContractFactory, Contract extends BaseContract>({
  contractName,
  deployParams,
  existingContractAddress,
}: GetContractParams<Factory> & {
  contractName: string;
}): Promise<Contract> => {
  const ContractFactory = (await ethers.getContractFactory(contractName)) as Factory;

  const isGetExistingContract = Boolean(existingContractAddress);
  if (isGetExistingContract) {
    console.log("Getting existing contract from address:", existingContractAddress);
    return ContractFactory.attach(existingContractAddress!) as Contract;
  }

  const contract = (await ContractFactory.deploy(...deployParams!)) as Contract;
  await contract.deployed();

  return contract;
};
