import { MaxUint256 } from "@ethersproject/constants";
import { parseEther, parseUnits } from "@ethersproject/units";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import {
  IERC20,
  IERC20__factory,
  UniswapV2Router02__factory,
  HanaTokenConsumer,
  HanaTokenConsumerUniV2,
  HanaTokenConsumerUniV3,
} from "@typechain-types";
import chai, { expect } from "chai";
import { BigNumber } from "ethers";
import { ethers } from "hardhat";
import { getNonHanaAddress } from "lib";

import { getTestAddress } from "../lib/address.helpers";
import {
  deployHanaNonEth,
  getHanaTokenConsumerUniV2Strategy,
  getHanaTokenConsumerUniV3Strategy,
} from "../lib/contracts.helpers";
import { parseHanaConsumerLog } from "./test.helpers";

chai.should();

describe("HanaTokenConsumer tests", () => {
  let uniswapV2RouterAddr: string;
  let uniswapV3RouterAddr: string;
  let USDCAddr: string;

  let hanaTokenConsumerUniV2: HanaTokenConsumerUniV2;
  let hanaTokenConsumerUniV3: HanaTokenConsumerUniV3;
  let hanaTokenNonEthAddress: string;
  let hanaTokenNonEth: IERC20;

  let accounts: SignerWithAddress[];
  let tssUpdater: SignerWithAddress;
  let tssSigner: SignerWithAddress;
  let randomSigner: SignerWithAddress;

  const getNow = async () => {
    const block = await ethers.provider.getBlock("latest");
    return block.timestamp;
  };

  const swapToken = async (signer: SignerWithAddress, tokenAddress: string, expectedAmount: BigNumber) => {
    const uniswapRouter = UniswapV2Router02__factory.connect(uniswapV2RouterAddr, signer);
    const WETH = await uniswapRouter.WETH();
    const path = [WETH, tokenAddress];
    const tx = await uniswapRouter
      .connect(signer)
      .swapETHForExactTokens(expectedAmount, path, signer.address, (await getNow()) + 360, { value: parseEther("10") });
    await tx.wait();
  };

  beforeEach(async () => {
    accounts = await ethers.getSigners();
    [tssUpdater, tssSigner, randomSigner] = accounts;

    hanaTokenNonEth = await deployHanaNonEth({
      args: [tssSigner.address, tssUpdater.address],
    });

    const DAI = getTestAddress("dai", "eth_mainnet");

    USDCAddr = getTestAddress("usdc", "eth_mainnet");

    uniswapV2RouterAddr = getNonHanaAddress("uniswapV2Router02", "eth_mainnet");

    const UNI_FACTORY_V3 = getNonHanaAddress("uniswapV3Factory", "eth_mainnet");

    const UNI_ROUTER_V3 = getNonHanaAddress("uniswapV3Router", "eth_mainnet");

    const WETH9 = getNonHanaAddress("weth9", "eth_mainnet");

    // For testing purposes we use an existing uni v3 pool
    await swapToken(tssUpdater, DAI, parseEther("10000"));
    await swapToken(randomSigner, DAI, parseEther("10000"));
    await swapToken(randomSigner, DAI, parseEther("10000"));
    await swapToken(randomSigner, DAI, parseEther("10000"));
    await swapToken(randomSigner, DAI, parseEther("10000"));

    hanaTokenNonEthAddress = DAI;
    hanaTokenNonEth = IERC20__factory.connect(hanaTokenNonEthAddress, tssSigner);

    hanaTokenConsumerUniV2 = await getHanaTokenConsumerUniV2Strategy({
      deployParams: [hanaTokenNonEthAddress, uniswapV2RouterAddr],
    });

    uniswapV3RouterAddr = UNI_ROUTER_V3;
    hanaTokenConsumerUniV3 = await getHanaTokenConsumerUniV3Strategy({
      deployParams: [hanaTokenNonEthAddress, uniswapV3RouterAddr, UNI_FACTORY_V3, WETH9, 3000, 3000],
    });
  });

  describe("getHanaFromEth", () => {
    const shouldGetHanaFromETH = async (hanaTokenConsumer: HanaTokenConsumer) => {
      const initialHanaBalance = await hanaTokenNonEth.balanceOf(randomSigner.address);
      const tx = await hanaTokenConsumer.getHanaFromEth(randomSigner.address, 1, { value: parseEther("1") });

      const result = await tx.wait();
      const eventNames = parseHanaConsumerLog(result.logs);
      expect(eventNames.filter((e) => e === "EthExchangedForHana")).to.have.lengthOf(1);

      const finalHanaBalance = await hanaTokenNonEth.balanceOf(randomSigner.address);
      expect(finalHanaBalance).to.be.gt(initialHanaBalance);
    };

    it("Should get hana from eth using UniV2", async () => {
      const hanaTokenConsumer = hanaTokenConsumerUniV2.connect(randomSigner);
      await shouldGetHanaFromETH(hanaTokenConsumer);
    });

    it("Should get hana from eth using UniV3", async () => {
      const hanaTokenConsumer = hanaTokenConsumerUniV3.connect(randomSigner);
      await shouldGetHanaFromETH(hanaTokenConsumer);
    });
  });

  describe("getHanaFromToken", () => {
    const shouldGetHanaFromToken = async (hanaTokenConsumer: HanaTokenConsumer) => {
      const USDCContract = IERC20__factory.connect(USDCAddr, randomSigner);
      await swapToken(randomSigner, USDCAddr, parseUnits("10000", 6));

      const initialHanaBalance = await hanaTokenNonEth.balanceOf(randomSigner.address);
      const tx1 = await USDCContract.approve(hanaTokenConsumer.address, MaxUint256);
      await tx1.wait();

      const tx2 = await hanaTokenConsumer.getHanaFromToken(randomSigner.address, 1, USDCAddr, parseUnits("100", 6));
      const result = await tx2.wait();

      const eventNames = parseHanaConsumerLog(result.logs);
      expect(eventNames.filter((e) => e === "TokenExchangedForHana")).to.have.lengthOf(1);

      const finalHanaBalance = await hanaTokenNonEth.balanceOf(randomSigner.address);
      expect(finalHanaBalance).to.be.gt(initialHanaBalance);
    };

    it("Should get hana from token using UniV2", async () => {
      const hanaTokenConsumer = hanaTokenConsumerUniV2.connect(randomSigner);
      await shouldGetHanaFromToken(hanaTokenConsumer);
    });

    it("Should get hana from token using UniV3", async () => {
      const hanaTokenConsumer = hanaTokenConsumerUniV3.connect(randomSigner);
      await shouldGetHanaFromToken(hanaTokenConsumer);
    });
  });

  describe("getEthFromHana", () => {
    const shouldGetETHFromHana = async (hanaTokenConsumer: HanaTokenConsumer) => {
      const initialEthBalance = await ethers.provider.getBalance(randomSigner.address);
      const tx1 = await hanaTokenNonEth.connect(randomSigner).approve(hanaTokenConsumer.address, MaxUint256);
      await tx1.wait();

      const tx2 = await hanaTokenConsumer.getEthFromHana(randomSigner.address, 1, parseUnits("5000", 18));
      const result = await tx2.wait();

      const eventNames = parseHanaConsumerLog(result.logs);
      expect(eventNames.filter((e) => e === "HanaExchangedForEth")).to.have.lengthOf(1);

      const finalEthBalance = await ethers.provider.getBalance(randomSigner.address);
      expect(finalEthBalance).to.be.gt(initialEthBalance);
    };

    it("Should get eth from hana using UniV2", async () => {
      const hanaTokenConsumer = hanaTokenConsumerUniV2.connect(randomSigner);
      await shouldGetETHFromHana(hanaTokenConsumer);
    });

    it("Should get eth from hana using UniV3", async () => {
      const hanaTokenConsumer = hanaTokenConsumerUniV3.connect(randomSigner);

      await shouldGetETHFromHana(hanaTokenConsumer);
    });
  });

  describe("getTokenFromHana", () => {
    const shouldGetTokenFromHana = async (hanaTokenConsumer: HanaTokenConsumer) => {
      const USDCContract = IERC20__factory.connect(USDCAddr, randomSigner);

      const initialTokenBalance = await USDCContract.balanceOf(randomSigner.address);
      const tx1 = await hanaTokenNonEth.connect(randomSigner).approve(hanaTokenConsumer.address, MaxUint256);
      await tx1.wait();

      const tx2 = await hanaTokenConsumer.getTokenFromHana(randomSigner.address, 1, USDCAddr, parseUnits("5000", 18));
      const result = await tx2.wait();

      const eventNames = parseHanaConsumerLog(result.logs);
      expect(eventNames.filter((e) => e === "HanaExchangedForToken")).to.have.lengthOf(1);

      const finalTokenBalance = await USDCContract.balanceOf(randomSigner.address);
      expect(finalTokenBalance).to.be.gt(initialTokenBalance);
    };

    it("Should get token from hana using UniV2", async () => {
      const hanaTokenConsumer = hanaTokenConsumerUniV2.connect(randomSigner);
      await shouldGetTokenFromHana(hanaTokenConsumer);
    });

    it("Should get token from hana using UniV3", async () => {
      const hanaTokenConsumer = hanaTokenConsumerUniV3.connect(randomSigner);
      await shouldGetTokenFromHana(hanaTokenConsumer);
    });
  });
});
