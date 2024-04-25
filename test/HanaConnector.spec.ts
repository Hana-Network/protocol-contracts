import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import {
  HanaConnectorBase,
  HanaConnectorEth,
  HanaConnectorNonEth,
  HanaEth,
  HanaNonEth,
  HanaReceiverMock,
} from "@typechain-types";
import { expect } from "chai";
import { parseEther } from "ethers/lib/utils";
import { ethers } from "hardhat";

import {
  deployHanaConnectorBase,
  deployHanaConnectorEth,
  deployHanaConnectorNonEth,
  deployHanaEth,
  deployHanaNonEth,
  deployHanaReceiverMock,
} from "../lib/contracts.helpers";

describe("HanaConnector tests", () => {
  let hanaTokenEthContract: HanaEth;
  let hanaTokenNonEthContract: HanaNonEth;
  let hanaConnectorBaseContract: HanaConnectorBase;
  let hanaConnectorEthContract: HanaConnectorEth;
  let hanaReceiverMockContract: HanaReceiverMock;
  let hanaConnectorNonEthContract: HanaConnectorNonEth;

  let tssUpdater: SignerWithAddress;
  let tssSigner: SignerWithAddress;
  let randomSigner: SignerWithAddress;
  let pauserSigner: SignerWithAddress;

  const tssUpdaterApproveConnectorEth = async () => {
    await (await hanaTokenEthContract.approve(hanaConnectorEthContract.address, parseEther("100000"))).wait();
  };

  const tssUpdaterApproveConnectorNonEth = async () => {
    await (await hanaTokenNonEthContract.approve(hanaConnectorNonEthContract.address, parseEther("100000"))).wait();
  };

  const transfer100kHanaEth = async (transferTo: string) => {
    await (await hanaTokenEthContract.transfer(transferTo, 100_000)).wait();
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

    hanaTokenEthContract = await deployHanaEth({
      args: [tssUpdater.address, 100_000],
    });

    hanaTokenNonEthContract = await deployHanaNonEth({
      args: [tssSigner.address, tssUpdater.address],
    });

    hanaReceiverMockContract = await deployHanaReceiverMock();
    hanaConnectorBaseContract = await deployHanaConnectorBase({
      args: [hanaTokenEthContract.address, tssSigner.address, tssUpdater.address, pauserSigner.address],
    });
    hanaConnectorEthContract = await deployHanaConnectorEth({
      args: [hanaTokenEthContract.address, tssSigner.address, tssUpdater.address, pauserSigner.address],
    });
    hanaConnectorNonEthContract = await deployHanaConnectorNonEth({
      args: [hanaTokenNonEthContract.address, tssSigner.address, tssUpdater.address, pauserSigner.address],
    });

    await hanaTokenNonEthContract.updateTssAndConnectorAddresses(
      tssSigner.address,
      hanaConnectorNonEthContract.address
    );

    await mint100kHanaNonEth(tssUpdater.address);
  });

  describe("HanaConnector.base", () => {
    describe("updateTssAddress", () => {
      it("Should revert if the caller is not TSS or TSS updater", async () => {
        await expect(hanaConnectorBaseContract.connect(randomSigner).updateTssAddress(randomSigner.address))
          .to.revertedWith("CallerIsNotTssOrUpdater")
          .withArgs(randomSigner.address);
      });

      it("Should revert if the new TSS address is invalid", async () => {
        await expect(
          hanaConnectorBaseContract.updateTssAddress("0x0000000000000000000000000000000000000000")
        ).to.revertedWith("InvalidAddress");
      });

      it("Should change the TSS address if called by TSS", async () => {
        await (await hanaConnectorBaseContract.connect(tssSigner).updateTssAddress(randomSigner.address)).wait();

        const address = await hanaConnectorBaseContract.tssAddress();

        expect(address).to.equal(randomSigner.address);
      });

      it("Should change the TSS address if called by TSS updater", async () => {
        await (await hanaConnectorBaseContract.updateTssAddress(randomSigner.address)).wait();

        const address = await hanaConnectorBaseContract.tssAddress();

        expect(address).to.equal(randomSigner.address);
      });
    });

    describe("updatePauserAddress", () => {
      it("Should revert if the caller is not the Pauser", async () => {
        await expect(hanaConnectorBaseContract.connect(randomSigner).updatePauserAddress(randomSigner.address))
          .to.revertedWith("CallerIsNotPauser")
          .withArgs(randomSigner.address);
      });

      it("Should revert if the new Pauser address is invalid", async () => {
        await expect(
          hanaConnectorBaseContract
            .connect(pauserSigner)
            .updatePauserAddress("0x0000000000000000000000000000000000000000")
        ).to.revertedWith("InvalidAddress");
      });

      it("Should change the Pauser address if called by Pauser", async () => {
        await (await hanaConnectorBaseContract.connect(pauserSigner).updatePauserAddress(randomSigner.address)).wait();

        const address = await hanaConnectorBaseContract.pauserAddress();

        expect(address).to.equal(randomSigner.address);
      });

      it("Should emit `PauserAddressUpdated` on success", async () => {
        const pauserAddressUpdatedFilter = hanaConnectorBaseContract.filters.PauserAddressUpdated();

        const tx = await hanaConnectorBaseContract.connect(pauserSigner).updatePauserAddress(randomSigner.address);
        const receipt = await tx.wait();

        const address = await hanaConnectorBaseContract.pauserAddress();

        expect(address).to.equal(randomSigner.address);

        const event = await hanaConnectorBaseContract.queryFilter(pauserAddressUpdatedFilter, receipt.blockHash);
        expect(event.length).to.equal(1);
      });
    });

    describe("pause, unpause", () => {
      it("Should revert if not called by the Pauser", async () => {
        await expect(hanaConnectorBaseContract.connect(randomSigner).pause())
          .to.revertedWith("CallerIsNotPauser")
          .withArgs(randomSigner.address);

        await expect(hanaConnectorBaseContract.connect(randomSigner).unpause())
          .to.revertedWith("CallerIsNotPauser")
          .withArgs(randomSigner.address);
      });

      it("Should pause if called by the Pauser", async () => {
        await (await hanaConnectorBaseContract.connect(pauserSigner).pause()).wait();
        const paused1 = await hanaConnectorBaseContract.paused();
        expect(paused1).to.equal(true);

        await (await hanaConnectorBaseContract.connect(pauserSigner).unpause()).wait();
        const paused2 = await hanaConnectorBaseContract.paused();
        expect(paused2).to.equal(false);
      });
    });
  });

  describe("HanaConnector.eth", () => {
    describe("send", () => {
      it("Should revert if the contract is paused", async () => {
        await (await hanaConnectorEthContract.connect(pauserSigner).pause()).wait();
        const paused1 = await hanaConnectorEthContract.paused();
        expect(paused1).to.equal(true);

        await expect(
          hanaConnectorEthContract.send({
            destinationAddress: randomSigner.address,
            destinationChainId: 1,
            destinationGasLimit: 2500000,
            hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            hanaValueAndGas: 1000,
            message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          })
        ).to.revertedWith("Pausable: paused");
      });

      it("Should revert if the hanaTxSender has no enough hana", async () => {
        await (
          await hanaTokenEthContract.connect(randomSigner).approve(hanaConnectorEthContract.address, 100_000)
        ).wait();

        await expect(
          hanaConnectorEthContract.connect(randomSigner).send({
            destinationAddress: randomSigner.address,
            destinationChainId: 1,
            destinationGasLimit: 2500000,
            hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            hanaValueAndGas: 1000,
            message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          })
        ).to.revertedWith("ERC20: transfer amount exceeds balance");
      });

      it("Should revert if the hanaTxSender didn't allow HanaConnector to spend Hana token", async () => {
        await expect(
          hanaConnectorEthContract.send({
            destinationAddress: randomSigner.address,
            destinationChainId: 1,
            destinationGasLimit: 2500000,
            hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            hanaValueAndGas: 1000,
            message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          })
        ).to.revertedWith("ERC20: insufficient allowance");
      });

      it("Should transfer Hana token from the hanaTxSender account to the Connector contract", async () => {
        const initialBalanceDeployer = await hanaTokenEthContract.balanceOf(tssUpdater.address);
        const initialBalanceConnector = await hanaTokenEthContract.balanceOf(hanaConnectorEthContract.address);

        expect(initialBalanceDeployer.toString()).to.equal("100000000000000000000000");
        expect(initialBalanceConnector.toString()).to.equal("0");

        await tssUpdaterApproveConnectorEth();

        await (
          await hanaConnectorEthContract.send({
            destinationAddress: randomSigner.address,
            destinationChainId: 1,
            destinationGasLimit: 2500000,
            hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            hanaValueAndGas: 1000,
            message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          })
        ).wait();

        const finalBalanceDeployer = await hanaTokenEthContract.balanceOf(tssUpdater.address);
        const finalBalanceConnector = await hanaTokenEthContract.balanceOf(hanaConnectorEthContract.address);

        expect(finalBalanceDeployer.toString()).to.equal("99999999999999999999000");
        expect(finalBalanceConnector.toString()).to.equal("1000");
      });

      it("Should emit `HanaSent` on success", async () => {
        const hanaSentFilter = hanaConnectorEthContract.filters.HanaSent();

        const tx = await hanaConnectorEthContract.send({
          destinationAddress: randomSigner.address,
          destinationChainId: 1,
          destinationGasLimit: 2500000,
          hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          hanaValueAndGas: 0,
          message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        });
        const receipt = await tx.wait();
        const events = await hanaConnectorEthContract.queryFilter(hanaSentFilter, receipt.blockHash);
        expect(events.length).to.equal(1);
      });

      it("Should emit `HanaSent` with tx.origin as the first parameter", async () => {
        const hanaSentFilter = hanaConnectorEthContract.filters.HanaSent();

        const tx = await hanaConnectorEthContract.connect(randomSigner).send({
          destinationAddress: randomSigner.address,
          destinationChainId: 1,
          destinationGasLimit: 2500000,
          hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          hanaValueAndGas: 0,
          message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        });

        const receipt = await tx.wait();

        const events = await hanaConnectorEthContract.queryFilter(hanaSentFilter, receipt.blockHash);
        expect(events[0].args[0].toString()).to.equal(randomSigner.address);
      });
    });

    describe("onReceive", () => {
      it("Should not revert if the contract is paused", async () => {
        await (await hanaConnectorEthContract.connect(pauserSigner).pause()).wait();
        const paused1 = await hanaConnectorEthContract.paused();
        expect(paused1).to.equal(true);

        await expect(
          hanaConnectorEthContract
            .connect(tssSigner)
            .onReceive(
              tssUpdater.address,
              1,
              hanaReceiverMockContract.address,
              0,
              new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
              ethers.constants.HashZero
            )
        ).to.be.not.reverted;
      });

      it("Should revert if not called by TSS address", async () => {
        await expect(
          hanaConnectorEthContract.onReceive(
            tssUpdater.address,
            1,
            randomSigner.address,
            1000,
            new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            ethers.constants.HashZero
          )
        )
          .to.revertedWith("CallerIsNotTss")
          .withArgs(tssUpdater.address);
      });

      it("Should revert if Hana transfer fails", async () => {
        await expect(
          hanaConnectorEthContract
            .connect(tssSigner)
            .onReceive(
              randomSigner.address,
              1,
              randomSigner.address,
              1000,
              new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
              ethers.constants.HashZero
            )
        ).to.revertedWith("ERC20: transfer amount exceeds balance");
      });

      it("Should transfer to the receiver address", async () => {
        await transfer100kHanaEth(hanaConnectorEthContract.address);

        const initialBalanceConnector = await hanaTokenEthContract.balanceOf(hanaConnectorEthContract.address);
        const initialBalanceReceiver = await hanaTokenEthContract.balanceOf(hanaReceiverMockContract.address);
        expect(initialBalanceConnector.toString()).to.equal("100000");
        expect(initialBalanceReceiver.toString()).to.equal("0");

        await (
          await hanaConnectorEthContract
            .connect(tssSigner)
            .onReceive(
              randomSigner.address,
              1,
              hanaReceiverMockContract.address,
              1000,
              new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
              ethers.constants.HashZero
            )
        ).wait();

        const finalBalanceConnector = await hanaTokenEthContract.balanceOf(hanaConnectorEthContract.address);
        const finalBalanceReceiver = await hanaTokenEthContract.balanceOf(hanaReceiverMockContract.address);

        expect(finalBalanceConnector.toString()).to.equal("99000");
        expect(finalBalanceReceiver.toString()).to.equal("1000");
      });

      it("Should emit `HanaReceived` on success", async () => {
        await transfer100kHanaEth(hanaConnectorEthContract.address);

        const hanaReceivedFilter = hanaConnectorEthContract.filters.HanaReceived();

        const tx = await hanaConnectorEthContract
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

        const events = await hanaConnectorEthContract.queryFilter(hanaReceivedFilter, receipt.blockHash);
        expect(events.length).to.equal(1);
      });
    });

    describe("onRevert", () => {
      it("Should revert if the contract is paused", async () => {
        await (await hanaConnectorEthContract.connect(pauserSigner).pause()).wait();
        const paused1 = await hanaConnectorEthContract.paused();
        expect(paused1).to.equal(true);

        await expect(
          hanaConnectorEthContract.onRevert(
            randomSigner.address,
            1,
            randomSigner.address,
            2,
            1000,
            new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            ethers.constants.HashZero
          )
        ).to.revertedWith("Pausable: paused");
      });

      it("Should revert if not called by TSS address", async () => {
        await expect(
          hanaConnectorEthContract.onRevert(
            randomSigner.address,
            1,
            tssUpdater.address,
            1,
            1000,
            new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            ethers.constants.HashZero
          )
        )
          .to.revertedWith("CallerIsNotTss")
          .withArgs(tssUpdater.address);
      });

      it("Should transfer to the hanaTxSender address", async () => {
        await transfer100kHanaEth(hanaConnectorEthContract.address);

        const initialBalanceConnector = await hanaTokenEthContract.balanceOf(hanaConnectorEthContract.address);
        const initialBalanceHanaTxSender = await hanaTokenEthContract.balanceOf(hanaReceiverMockContract.address);
        expect(initialBalanceConnector.toString()).to.equal("100000");
        expect(initialBalanceHanaTxSender.toString()).to.equal("0");

        await (
          await hanaConnectorEthContract
            .connect(tssSigner)
            .onRevert(
              hanaReceiverMockContract.address,
              1,
              randomSigner.address,
              1,
              1000,
              new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
              ethers.constants.HashZero
            )
        ).wait();

        const finalBalanceConnector = await hanaTokenEthContract.balanceOf(hanaConnectorEthContract.address);
        const finalBalanceHanaTxSender = await hanaTokenEthContract.balanceOf(hanaReceiverMockContract.address);

        expect(finalBalanceConnector.toString()).to.equal("99000");
        expect(finalBalanceHanaTxSender.toString()).to.equal("1000");
      });

      it("Should emit `HanaReverted` on success", async () => {
        await transfer100kHanaEth(hanaConnectorEthContract.address);

        const hanaRevertedFilter = hanaConnectorEthContract.filters.HanaReverted();

        const tx = await hanaConnectorEthContract
          .connect(tssSigner)
          .onRevert(
            hanaReceiverMockContract.address,
            1,
            randomSigner.address,
            1,
            1000,
            new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            ethers.constants.HashZero
          );
        const receipt = await tx.wait();

        const events = await hanaConnectorEthContract.queryFilter(hanaRevertedFilter, receipt.blockHash);
        expect(events.length).to.equal(1);
      });
    });
  });

  describe("HanaConnector.non-eth", () => {
    describe("send", () => {
      it("Should revert if the contract is paused", async () => {
        await (await hanaConnectorNonEthContract.connect(pauserSigner).pause()).wait();
        const paused1 = await hanaConnectorNonEthContract.paused();
        expect(paused1).to.equal(true);

        await expect(
          hanaConnectorNonEthContract.send({
            destinationAddress: randomSigner.address,
            destinationChainId: 1,
            destinationGasLimit: 2500000,
            hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            hanaValueAndGas: 1000,
            message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          })
        ).to.revertedWith("Pausable: paused");
      });

      it("Should revert if the hanaTxSender has no enough hana", async () => {
        await (
          await hanaTokenEthContract.connect(randomSigner).approve(hanaConnectorEthContract.address, 100_000)
        ).wait();

        await expect(
          hanaConnectorNonEthContract.connect(randomSigner).send({
            destinationAddress: randomSigner.address,
            destinationChainId: 1,
            destinationGasLimit: 2500000,
            hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            hanaValueAndGas: 1000,
            message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          })
        ).to.revertedWith("ERC20: insufficient allowance");
      });

      it("Should revert if the hanaTxSender didn't allow HanaConnector to spend Hana token", async () => {
        await expect(
          hanaConnectorNonEthContract.send({
            destinationAddress: randomSigner.address,
            destinationChainId: 1,
            destinationGasLimit: 2500000,
            hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            hanaValueAndGas: 1000,
            message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          })
        ).to.revertedWith("ERC20: insufficient allowance");
      });

      it("Should burn Hana token from the hanaTxSender account", async () => {
        const initialBalanceDeployer = await hanaTokenNonEthContract.balanceOf(tssUpdater.address);
        expect(initialBalanceDeployer.toString()).to.equal(parseEther("100000"));

        await tssUpdaterApproveConnectorNonEth();

        await (
          await hanaConnectorNonEthContract.send({
            destinationAddress: randomSigner.address,
            destinationChainId: 1,
            destinationGasLimit: 2500000,
            hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            hanaValueAndGas: parseEther("1"),
            message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          })
        ).wait();

        const finalBalanceDeployer = await hanaTokenNonEthContract.balanceOf(tssUpdater.address);
        expect(finalBalanceDeployer.toString()).to.equal(parseEther("99999"));
      });

      it("Should emit `HanaSent` on success", async () => {
        const hanaSentFilter = hanaConnectorNonEthContract.filters.HanaSent();

        const tx = await hanaConnectorNonEthContract.send({
          destinationAddress: randomSigner.address,
          destinationChainId: 1,
          destinationGasLimit: 2500000,
          hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          hanaValueAndGas: 0,
          message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        });
        const receipt = await tx.wait();

        const events = await hanaConnectorNonEthContract.queryFilter(hanaSentFilter, receipt.blockHash);
        expect(events.length).to.equal(1);
      });

      it("Should emit `HanaSent` with tx.origin as the first parameter", async () => {
        const hanaSentFilter = hanaConnectorNonEthContract.filters.HanaSent();

        const tx = await hanaConnectorNonEthContract.connect(randomSigner).send({
          destinationAddress: randomSigner.address,
          destinationChainId: 1,
          destinationGasLimit: 2500000,
          hanaParams: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
          hanaValueAndGas: 0,
          message: new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
        });
        const receipt = await tx.wait();

        const events = await hanaConnectorNonEthContract.queryFilter(hanaSentFilter, receipt.blockHash);
        expect(events[0].args[0].toString()).to.equal(randomSigner.address);
      });
    });

    describe("onReceive", () => {
      it("Should not revert if the contract is paused", async () => {
        await (await hanaConnectorNonEthContract.connect(pauserSigner).pause()).wait();
        const paused1 = await hanaConnectorNonEthContract.paused();
        expect(paused1).to.equal(true);

        await expect(
          hanaConnectorNonEthContract
            .connect(tssSigner)
            .onReceive(
              tssUpdater.address,
              1,
              hanaReceiverMockContract.address,
              0,
              new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
              ethers.constants.HashZero
            )
        ).to.be.not.reverted;
      });

      it("Should revert if not called by TSS address", async () => {
        await expect(
          hanaConnectorNonEthContract.onReceive(
            tssUpdater.address,
            1,
            randomSigner.address,
            1000,
            new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            ethers.constants.HashZero
          )
        )
          .to.revertedWith("CallerIsNotTss")
          .withArgs(tssUpdater.address);
      });

      it("Should revert if mint fails", async () => {
        /**
         * Update TSS and Connector addresses so minting fails
         */
        await hanaTokenNonEthContract.updateTssAndConnectorAddresses(tssSigner.address, randomSigner.address);

        await expect(
          hanaConnectorNonEthContract
            .connect(tssSigner)
            .onReceive(
              randomSigner.address,
              1,
              hanaReceiverMockContract.address,
              1000,
              new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
              ethers.constants.HashZero
            )
        )
          .to.revertedWith("CallerIsNotConnector")
          .withArgs(hanaConnectorNonEthContract.address);
      });

      it("Should mint on the receiver address", async () => {
        const initialBalanceReceiver = await hanaTokenNonEthContract.balanceOf(hanaReceiverMockContract.address);
        expect(initialBalanceReceiver.toString()).to.equal("0");

        await (
          await hanaConnectorNonEthContract
            .connect(tssSigner)
            .onReceive(
              randomSigner.address,
              1,
              hanaReceiverMockContract.address,
              1000,
              new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
              ethers.constants.HashZero
            )
        ).wait();

        const finalBalanceReceiver = await hanaTokenNonEthContract.balanceOf(hanaReceiverMockContract.address);

        expect(finalBalanceReceiver.toString()).to.equal("1000");
      });

      it("Should emit `HanaReceived` on success", async () => {
        const hanaReceivedFilter = hanaConnectorNonEthContract.filters.HanaReceived();

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

        const events = await hanaConnectorNonEthContract.queryFilter(hanaReceivedFilter, receipt.blockHash);
        expect(events.length).to.equal(1);
      });
    });

    describe("onRevert", () => {
      it("Should revert if the contract is paused", async () => {
        await (await hanaConnectorNonEthContract.connect(pauserSigner).pause()).wait();
        const paused1 = await hanaConnectorNonEthContract.paused();
        expect(paused1).to.equal(true);

        await expect(
          hanaConnectorNonEthContract.onRevert(
            randomSigner.address,
            1,
            randomSigner.address,
            2,
            1000,
            new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            ethers.constants.HashZero
          )
        ).to.revertedWith("Pausable: paused");
      });

      it("Should revert if not called by TSS address", async () => {
        await expect(
          hanaConnectorNonEthContract.onRevert(
            randomSigner.address,
            1,
            tssUpdater.address,
            1,
            1000,
            new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            ethers.constants.HashZero
          )
        )
          .to.revertedWith("CallerIsNotTss")
          .withArgs(tssUpdater.address);
      });

      it("Should mint on the hanaTxSender address", async () => {
        const initialBalanceHanaTxSender = await hanaTokenNonEthContract.balanceOf(hanaReceiverMockContract.address);
        expect(initialBalanceHanaTxSender.toString()).to.equal("0");

        await (
          await hanaConnectorNonEthContract
            .connect(tssSigner)
            .onRevert(
              hanaReceiverMockContract.address,
              1,
              randomSigner.address,
              1,
              1000,
              new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
              ethers.constants.HashZero
            )
        ).wait();

        const finalBalanceHanaTxSender = await hanaTokenNonEthContract.balanceOf(hanaReceiverMockContract.address);
        expect(finalBalanceHanaTxSender.toString()).to.equal("1000");
      });

      it("Should emit `HanaReverted` on success", async () => {
        await transfer100kHanaNonEth(hanaConnectorNonEthContract.address);

        const hanaRevertedFilter = hanaConnectorNonEthContract.filters.HanaReverted();

        const tx = await hanaConnectorNonEthContract
          .connect(tssSigner)
          .onRevert(
            hanaReceiverMockContract.address,
            1,
            randomSigner.address,
            1,
            1000,
            new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
            ethers.constants.HashZero
          );
        const receipt = await tx.wait();

        const events = await hanaConnectorNonEthContract.queryFilter(hanaRevertedFilter, receipt.blockHash);
        expect(events.length).to.equal(1);
      });
    });

    describe("MaxSupply", () => {
      describe("setMaxSupply", () => {
        it("Should revert if the caller is not the TSS address", async () => {
          await expect(hanaConnectorNonEthContract.connect(randomSigner).setMaxSupply(0))
            .to.revertedWith("CallerIsNotTss")
            .withArgs(randomSigner.address);
        });

        it("Should revert if want to mint more than MaxSupply", async () => {
          await hanaConnectorNonEthContract.connect(tssSigner).setMaxSupply(999);
          await expect(
            hanaConnectorNonEthContract
              .connect(tssSigner)
              .onReceive(
                randomSigner.address,
                1,
                hanaReceiverMockContract.address,
                1000,
                new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
                ethers.constants.HashZero
              )
          )
            .to.revertedWith("ExceedsMaxSupply")
            .withArgs(999);
        });
      });

      describe("onReceive, onRevert (mint)", () => {
        it("Should mint if total supply + supply to add < max supply", async () => {
          const supplyToAdd = 1000;
          const initialSupply = await hanaTokenNonEthContract.totalSupply();

          await hanaConnectorNonEthContract.connect(tssSigner).setMaxSupply(initialSupply.add(supplyToAdd));

          await expect(
            hanaConnectorNonEthContract
              .connect(tssSigner)
              .onReceive(
                randomSigner.address,
                1,
                hanaReceiverMockContract.address,
                supplyToAdd,
                new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
                ethers.constants.HashZero
              )
          ).to.be.not.reverted;

          const finalSupply = await hanaTokenNonEthContract.totalSupply();

          expect(finalSupply).to.eq(initialSupply.add(supplyToAdd));

          await expect(
            hanaConnectorNonEthContract
              .connect(tssSigner)
              .onReceive(
                randomSigner.address,
                1,
                hanaReceiverMockContract.address,
                1,
                new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
                ethers.constants.HashZero
              )
          )
            .to.revertedWith("ExceedsMaxSupply")
            .withArgs(initialSupply.add(supplyToAdd));

          await expect(
            hanaConnectorNonEthContract
              .connect(tssSigner)
              .onRevert(
                randomSigner.address,
                1,
                randomSigner.address,
                2,
                1000,
                new ethers.utils.AbiCoder().encode(["string"], ["hello"]),
                ethers.constants.HashZero
              )
          )
            .to.revertedWith("ExceedsMaxSupply")
            .withArgs(initialSupply.add(supplyToAdd));
        });
      });
    });
  });
});
