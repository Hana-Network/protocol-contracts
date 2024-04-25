import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { HanaInteractorMock } from "@typechain-types";
import chai, { expect } from "chai";
import { ethers } from "hardhat";

import { getHanaInteractorMock } from "../lib/contracts.helpers";

chai.should();

describe("HanaInteractor tests", () => {
  let hanaInteractorMock: HanaInteractorMock;
  const chainAId = 1;
  const chainBId = 2;

  let accounts: SignerWithAddress[];
  let deployer: SignerWithAddress;
  let crossChainContractB: SignerWithAddress;
  let hanaConnector: SignerWithAddress;

  const encoder = new ethers.utils.AbiCoder();

  beforeEach(async () => {
    accounts = await ethers.getSigners();
    [deployer, crossChainContractB, hanaConnector] = accounts;

    hanaInteractorMock = await getHanaInteractorMock(hanaConnector.address);

    const encodedCrossChainAddressB = ethers.utils.solidityPack(["address"], [crossChainContractB.address]);
    await (await hanaInteractorMock.setInteractorByChainId(chainBId, encodedCrossChainAddressB)).wait();
  });

  describe("onCreate", () => {
    it("Should revert if constructor param is zero address", async () => {
      const Factory = await ethers.getContractFactory("HanaInteractorMock");
      await expect(Factory.deploy(ethers.constants.AddressZero)).to.be.revertedWith("InvalidAddress");
    });

    it("Should revert if the hanaTxSenderAddress it not in interactorsByChainId", async () => {
      await expect(
        hanaInteractorMock.connect(hanaConnector).onHanaMessage({
          destinationAddress: crossChainContractB.address,
          hanaTxSenderAddress: ethers.utils.solidityPack(["address"], [hanaInteractorMock.address]),
          hanaValue: 0,
          message: encoder.encode(["address"], [crossChainContractB.address]),
          sourceChainId: chainBId,
        })
      ).to.be.revertedWith("InvalidHanaMessageCall");
    });
  });

  describe("onHanaMessage", () => {
    it("Should revert if the caller is not HanaConnector", async () => {
      await expect(
        hanaInteractorMock.onHanaMessage({
          destinationAddress: crossChainContractB.address,
          hanaTxSenderAddress: ethers.utils.solidityPack(["address"], [hanaInteractorMock.address]),
          hanaValue: 0,
          message: encoder.encode(["address"], [hanaInteractorMock.address]),
          sourceChainId: chainBId,
        })
      )
        .to.be.revertedWith("InvalidCaller")
        .withArgs(deployer.address);
    });

    it("Should revert if the hanaTxSenderAddress it not in interactorsByChainId", async () => {
      await expect(
        hanaInteractorMock.connect(hanaConnector).onHanaMessage({
          destinationAddress: crossChainContractB.address,
          hanaTxSenderAddress: ethers.utils.solidityPack(["address"], [hanaInteractorMock.address]),
          hanaValue: 0,
          message: encoder.encode(["address"], [crossChainContractB.address]),
          sourceChainId: chainBId,
        })
      ).to.be.revertedWith("InvalidHanaMessageCall");
    });
  });

  describe("onHanaRevert", () => {
    it("Should revert if the caller is not HanaConnector", async () => {
      await expect(
        hanaInteractorMock.onHanaRevert({
          destinationAddress: ethers.utils.solidityPack(["address"], [crossChainContractB.address]),
          destinationChainId: chainBId,
          hanaTxSenderAddress: deployer.address,
          message: encoder.encode(["address"], [hanaInteractorMock.address]),
          remainingHanaValue: 0,
          sourceChainId: chainAId,
        })
      )
        .to.be.revertedWith("InvalidCaller")
        .withArgs(deployer.address);
    });
  });

  describe("transferOwnership", () => {
    it("Should transfer ownership", async () => {
      const randomSigner = accounts[3];
      await hanaInteractorMock.transferOwnership(randomSigner.address);
      await hanaInteractorMock.connect(randomSigner).acceptOwnership();
      await expect(await hanaInteractorMock.owner()).to.be.eq(randomSigner.address);
    });

    it("Should keep the ownership until accept", async () => {
      const randomSigner = accounts[3];
      await hanaInteractorMock.transferOwnership(randomSigner.address);
      await expect(await hanaInteractorMock.owner()).to.be.eq(deployer.address);
    });

    it("Should revert if old owner want to do some action", async () => {
      const randomSigner = accounts[3];
      await hanaInteractorMock.transferOwnership(randomSigner.address);
      await hanaInteractorMock.connect(randomSigner).acceptOwnership();
      await expect(hanaInteractorMock.transferOwnership(randomSigner.address)).to.be.revertedWith(
        "Ownable: caller is not the owner"
      );
    });
  });
});
