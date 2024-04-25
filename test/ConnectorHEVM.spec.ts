import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { WETH9, HanaConnectorHEVM, HanaReceiverMock } from "@typechain-types";
import { expect } from "chai";
import { parseEther } from "ethers/lib/utils";
import { ethers } from "hardhat";

import { FUNGIBLE_MODULE_ADDRESS } from "./test.helpers";

const hre = require("hardhat");

describe("ConnectorHEVM tests", () => {
  let hanaTokenContract: WETH9;
  let hanaConnectorHEVM: HanaConnectorHEVM;
  let hanaReceiverMockContract: HanaReceiverMock;

  let owner: SignerWithAddress;
  let fungibleModuleSigner: SignerWithAddress;
  let addrs: SignerWithAddress[];
  let randomSigner: SignerWithAddress;

  beforeEach(async () => {
    [owner, randomSigner, ...addrs] = await ethers.getSigners();

    // Impersonate the fungible module account
    await hre.network.provider.request({
      method: "hardhat_impersonateAccount",
      params: [FUNGIBLE_MODULE_ADDRESS],
    });

    // Get a signer for the fungible module account
    fungibleModuleSigner = await ethers.getSigner(FUNGIBLE_MODULE_ADDRESS);
    hre.network.provider.send("hardhat_setBalance", [FUNGIBLE_MODULE_ADDRESS, parseEther("1000000").toHexString()]);

    const WHANAFactory = await ethers.getContractFactory("contracts/hevm/WHANA.sol:WETH9");
    hanaTokenContract = (await WHANAFactory.deploy()) as WETH9;

    const HanaConnectorHEVMFactory = await ethers.getContractFactory("HanaConnectorHEVM");
    hanaConnectorHEVM = (await HanaConnectorHEVMFactory.connect(owner).deploy(
      hanaTokenContract.address
    )) as HanaConnectorHEVM;

    const HanaReceiverMockContractactory = await ethers.getContractFactory("HanaReceiverMock");
    hanaReceiverMockContract = (await HanaReceiverMockContractactory.connect(owner).deploy()) as HanaReceiverMock;
  });

  describe("HanaConnectorHEVM", () => {
    it("Should revert if the hanaTxSender has no enough hana", async () => {
      await hanaTokenContract.connect(randomSigner).approve(hanaConnectorHEVM.address, 100_000);

      const tx = hanaConnectorHEVM.connect(randomSigner).send({
        destinationAddress: randomSigner.address,
        destinationChainId: 1,
        destinationGasLimit: 2500000,
        message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        hanaValueAndGas: 1000,
      });

      // @dev: As we use the standard WETH contract, there's no error message for not enough balance
      await expect(tx).to.be.reverted;
    });

    it("Should revert if the hanaTxSender didn't allow HanaConnector to spend Hana token", async () => {
      await hanaTokenContract.deposit({ value: 100_000 });
      await hanaTokenContract.transfer(randomSigner.address, 100_000);

      const balance = await hanaTokenContract.balanceOf(randomSigner.address);
      expect(balance.toString()).to.equal("100000");

      const tx = hanaConnectorHEVM.send({
        destinationAddress: randomSigner.address,
        destinationChainId: 1,
        destinationGasLimit: 2500000,
        message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        hanaValueAndGas: 1000,
      });

      // @dev: As we use the standard WETH contract, there's no error message for not enough balance
      await expect(tx).to.be.reverted;

      await hanaTokenContract.connect(randomSigner).transfer(owner.address, balance);
    });

    it("Should emit `HanaSent` on success", async () => {
      const tx = await hanaConnectorHEVM.send({
        destinationAddress: randomSigner.address,
        destinationChainId: 1,
        destinationGasLimit: 2500000,
        message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        hanaValueAndGas: 0,
      });

      expect(tx)
        .to.emit(hanaConnectorHEVM, "HanaSent")
        .withArgs(owner.address, owner.address, 1, randomSigner.address, 0, 2500000, "hello", "hello");
    });

    it("Should transfer value and gas to fungible address", async () => {
      const hanaValueAndGas = 1000;
      await hanaTokenContract.approve(hanaConnectorHEVM.address, hanaValueAndGas);
      await hanaTokenContract.deposit({ value: hanaValueAndGas });

      const balanceBefore = await ethers.provider.getBalance(fungibleModuleSigner.address);

      await hanaConnectorHEVM.send({
        destinationAddress: randomSigner.address,
        destinationChainId: 1,
        destinationGasLimit: 2500000,
        message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        hanaValueAndGas,
      });

      const balanceAfter = await ethers.provider.getBalance(fungibleModuleSigner.address);
      expect(balanceAfter.sub(balanceBefore)).to.equal(hanaValueAndGas);
    });

    it("Should update whana address if is call from fungible address", async () => {
      const WHANAFactory = await ethers.getContractFactory("contracts/hevm/WHANA.sol:WETH9");
      const newHanaTokenContract = (await WHANAFactory.deploy()) as WETH9;

      const tx = hanaConnectorHEVM.connect(fungibleModuleSigner).setWhanaAddress(newHanaTokenContract.address);
      await expect(tx).to.emit(hanaConnectorHEVM, "SetWHANA").withArgs(newHanaTokenContract.address);

      expect(await hanaConnectorHEVM.whana()).to.equal(newHanaTokenContract.address);
    });

    it("Should revert if try to update whana address from other address", async () => {
      const WHANAFactory = await ethers.getContractFactory("contracts/hevm/WHANA.sol:WETH9");
      const newHanaTokenContract = (await WHANAFactory.deploy()) as WETH9;

      const tx = hanaConnectorHEVM.setWhanaAddress(newHanaTokenContract.address);
      await expect(tx).to.be.revertedWith("OnlyFungibleModule");
    });

    it("Should transfer to the receiver address", async () => {
      // read eth balance of hanaConnectorEVM
      const initialHanaBalanceConnector = await ethers.provider.getBalance(hanaConnectorHEVM.address);
      const initialBalanceConnector = await hanaTokenContract.balanceOf(hanaConnectorHEVM.address);
      const initialBalanceReceiver = await hanaTokenContract.balanceOf(hanaReceiverMockContract.address);
      expect(initialHanaBalanceConnector.toString()).to.equal("0");
      expect(initialBalanceConnector.toString()).to.equal("0");
      expect(initialBalanceReceiver.toString()).to.equal("0");

      const tx = await hanaConnectorHEVM
        .connect(fungibleModuleSigner)
        .onReceive(
          randomSigner.address,
          1,
          hanaReceiverMockContract.address,
          1000,
          new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          ethers.constants.HashZero,
          { value: 1000 }
        );

      await expect(tx)
        .to.emit(hanaReceiverMockContract, "MockOnHanaMessage")
        .withArgs(hanaReceiverMockContract.address);

      const finalHanaBalanceConnector = await ethers.provider.getBalance(hanaConnectorHEVM.address);
      const finalBalanceConnector = await hanaTokenContract.balanceOf(hanaConnectorHEVM.address);
      const finalBalanceReceiver = await hanaTokenContract.balanceOf(hanaReceiverMockContract.address);

      expect(finalHanaBalanceConnector.toString()).to.equal("0");
      expect(finalBalanceConnector.toString()).to.equal("0");
      expect(finalBalanceReceiver.toString()).to.equal("1000");
    });

    it("Should call onRevert to the original address", async () => {
      // read eth balance of hanaConnectorEVM
      const initialHanaBalanceConnector = await ethers.provider.getBalance(hanaConnectorHEVM.address);
      const initialBalanceConnector = await hanaTokenContract.balanceOf(hanaConnectorHEVM.address);
      const initialBalanceReceiver = await hanaTokenContract.balanceOf(hanaReceiverMockContract.address);
      expect(initialHanaBalanceConnector.toString()).to.equal("0");
      expect(initialBalanceConnector.toString()).to.equal("0");
      expect(initialBalanceReceiver.toString()).to.equal("0");

      const tx = await hanaConnectorHEVM
        .connect(fungibleModuleSigner)
        .onRevert(
          hanaReceiverMockContract.address,
          1,
          randomSigner.address,
          5,
          1000,
          new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          ethers.constants.HashZero,
          { value: 1000 }
        );

      await expect(tx).to.emit(hanaReceiverMockContract, "MockOnHanaRevert").withArgs(hanaReceiverMockContract.address);

      const finalHanaBalanceConnector = await ethers.provider.getBalance(hanaConnectorHEVM.address);
      const finalBalanceConnector = await hanaTokenContract.balanceOf(hanaConnectorHEVM.address);
      const finalBalanceReceiver = await hanaTokenContract.balanceOf(hanaReceiverMockContract.address);

      expect(finalHanaBalanceConnector.toString()).to.equal("0");
      expect(finalBalanceConnector.toString()).to.equal("0");
      expect(finalBalanceReceiver.toString()).to.equal("1000");
    });

    it("Should accept HANA from fungible module", async () => {
      const initialHanaBalanceConnector = await ethers.provider.getBalance(hanaConnectorHEVM.address);
      expect(initialHanaBalanceConnector.toString()).to.equal("0");

      await fungibleModuleSigner.sendTransaction({
        to: hanaConnectorHEVM.address,
        value: 1000,
      });

      // read eth balance of hanaConnectorEVM
      const finalHanaBalanceConnector = await ethers.provider.getBalance(hanaConnectorHEVM.address);
      expect(finalHanaBalanceConnector.toString()).to.equal("1000");
    });

    it("Should reject HANA from other address", async () => {
      await fungibleModuleSigner.sendTransaction({
        to: randomSigner.address,
        value: 1000,
      });

      const tx = randomSigner.sendTransaction({
        to: hanaConnectorHEVM.address,
        value: 1000,
      });

      await expect(tx).to.be.revertedWith("OnlyWHANAOrFungible");
    });

    it("Should reject if call onReceive from other than fungible module", async () => {
      const tx = hanaConnectorHEVM.onReceive(
        randomSigner.address,
        1,
        hanaReceiverMockContract.address,
        1000,
        new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        ethers.constants.HashZero,
        { value: 1000 }
      );

      await expect(tx).to.be.revertedWith("OnlyFungibleModule");
    });

    it("Should reject if call onRevert from other than fungible module", async () => {
      const tx = hanaConnectorHEVM.onRevert(
        hanaReceiverMockContract.address,
        1,
        randomSigner.address,
        5,
        1000,
        new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        ethers.constants.HashZero,
        { value: 1000 }
      );

      await expect(tx).to.be.revertedWith("OnlyFungibleModule");
    });

    it("Should revert if value is not the same as declared in onReceive args", async () => {
      const tx = hanaConnectorHEVM
        .connect(fungibleModuleSigner)
        .onReceive(
          randomSigner.address,
          1,
          hanaReceiverMockContract.address,
          1000,
          new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          ethers.constants.HashZero,
          { value: 900 }
        );

      await expect(tx).to.be.revertedWith("WrongValue");
    });

    it("Should revert if value is not the same as declared in onRevert args", async () => {
      const tx = hanaConnectorHEVM
        .connect(fungibleModuleSigner)
        .onRevert(
          hanaReceiverMockContract.address,
          1,
          randomSigner.address,
          5,
          1000,
          new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          ethers.constants.HashZero,
          { value: 900 }
        );

      await expect(tx).to.be.revertedWith("WrongValue");
    });
  });
});
