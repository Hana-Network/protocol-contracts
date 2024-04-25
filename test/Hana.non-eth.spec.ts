import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { HanaConnectorNonEth, HanaNonEth, HanaReceiverMock } from "@typechain-types";
import { expect } from "chai";
import { parseEther } from "ethers/lib/utils";
import { ethers } from "hardhat";

import { deployHanaConnectorNonEth, deployHanaNonEth, deployHanaReceiverMock } from "../lib/contracts.helpers";

describe("HanaNonEth tests", () => {
  let hanaTokenNonEthContract: HanaNonEth;
  let hanaReceiverMockContract: HanaReceiverMock;
  let hanaConnectorNonEthContract: HanaConnectorNonEth;
  let tssUpdater: SignerWithAddress;
  let tssSigner: SignerWithAddress;
  let randomSigner: SignerWithAddress;
  let pauserSigner: SignerWithAddress;

  const tssUpdaterApproveConnectorNonEth = async () => {
    await (await hanaTokenNonEthContract.approve(hanaConnectorNonEthContract.address, parseEther("100000"))).wait();
  };

  const mint100kHanaNonEth = async (transferTo: string) => {
    const hana100k = parseEther("100000");

    await (
      await hanaConnectorNonEthContract
        .connect(tssSigner)
        .onReceive(randomSigner.address, 1, transferTo, hana100k, [], ethers.constants.HashZero)
    ).wait();
  };

  const transfer100kHanaNonEth = async (transferTo: string) => {
    await mint100kHanaNonEth(tssUpdater.address);

    await (await hanaTokenNonEthContract.connect(tssUpdater).transfer(transferTo, 100_000)).wait();
  };

  beforeEach(async () => {
    const accounts = await ethers.getSigners();
    [tssUpdater, tssSigner, randomSigner, pauserSigner] = accounts;

    hanaTokenNonEthContract = await deployHanaNonEth({
      args: [tssSigner.address, tssUpdater.address],
    });

    hanaReceiverMockContract = await deployHanaReceiverMock();
    hanaConnectorNonEthContract = await deployHanaConnectorNonEth({
      args: [hanaTokenNonEthContract.address, tssSigner.address, tssUpdater.address, pauserSigner.address],
    });

    await hanaTokenNonEthContract.updateTssAndConnectorAddresses(
      tssSigner.address,
      hanaConnectorNonEthContract.address
    );

    await mint100kHanaNonEth(tssUpdater.address);
  });

  describe("updateTssAndConnectorAddresses", () => {
    it("Should revert if the caller is not tssAddressUpdater or TSS", async () => {
      expect(
        hanaTokenNonEthContract
          .connect(randomSigner)
          .updateTssAndConnectorAddresses(randomSigner.address, hanaConnectorNonEthContract.address)
      ).to.be.revertedWith(`CallerIsNotTssOrUpdater("${randomSigner.address}")`);
    });

    it("Should change the addresses if the caller is tssAddressUpdater", async () => {
      await (
        await hanaTokenNonEthContract.updateTssAndConnectorAddresses(randomSigner.address, randomSigner.address)
      ).wait();

      expect(await hanaTokenNonEthContract.tssAddress()).to.equal(randomSigner.address);
      expect(await hanaTokenNonEthContract.connectorAddress()).to.equal(randomSigner.address);
    });

    it("Should change the addresses if the caller is TSS", async () => {
      await (
        await hanaTokenNonEthContract
          .connect(tssSigner)
          .updateTssAndConnectorAddresses(randomSigner.address, randomSigner.address)
      ).wait();

      expect(await hanaTokenNonEthContract.tssAddress()).to.equal(randomSigner.address);
      expect(await hanaTokenNonEthContract.connectorAddress()).to.equal(randomSigner.address);
    });
  });

  describe("renounceTssAddressUpdater", () => {
    it("Should revert if the caller is not tssAddressUpdater", async () => {
      expect(hanaTokenNonEthContract.connect(randomSigner).renounceTssAddressUpdater()).to.be.revertedWith(
        `CallerIsNotTssUpdater("${randomSigner.address}")`
      );
    });

    it("Should change tssAddressUpdater to tssAddress if the caller is tssAddressUpdater", async () => {
      await (await hanaTokenNonEthContract.renounceTssAddressUpdater()).wait();

      expect(await hanaTokenNonEthContract.tssAddressUpdater()).to.equal(tssSigner.address);
    });
  });

  describe("mint", () => {
    it("Should revert if the caller is not the Connector contract", async () => {
      expect(
        hanaTokenNonEthContract.connect(randomSigner).mint(tssUpdater.address, 100_000, ethers.constants.AddressZero)
      ).to.be.revertedWith(`CallerIsNotConnector("${randomSigner.address}")`);
    });

    it("Should emit `Minted` on success", async () => {
      const hanaMintedFilter = hanaTokenNonEthContract.filters.Minted();

      const tx = await hanaConnectorNonEthContract
        .connect(tssSigner)
        .onReceive(
          randomSigner.address,
          1,
          hanaReceiverMockContract.address,
          1000,
          new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          ethers.constants.HashZero
        );
      const receipt = await tx.wait();

      const events = await hanaTokenNonEthContract.queryFilter(hanaMintedFilter, receipt.blockHash);
      expect(events.length).to.equal(1);
    });
  });

  describe("burnFrom", () => {
    it("Should revert if the caller is not the Connector contract", async () => {
      expect(hanaTokenNonEthContract.connect(randomSigner).burnFrom(tssUpdater.address, 100_000)).to.be.revertedWith(
        `CallerIsNotConnector("${randomSigner.address}")`
      );
    });

    it("Should emit `Burnt` on success", async () => {
      await tssUpdaterApproveConnectorNonEth();
      const hanaBurntFilter = hanaTokenNonEthContract.filters.Burnt();

      const tx = await hanaConnectorNonEthContract.send({
        destinationAddress: randomSigner.address,
        destinationChainId: 1,
        destinationGasLimit: 2500000,
        message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        hanaValueAndGas: 1000,
      });
      const receipt = await tx.wait();

      const events = await hanaTokenNonEthContract.queryFilter(hanaBurntFilter, receipt.blockHash);
      expect(events.length).to.equal(1);
    });
  });
});
