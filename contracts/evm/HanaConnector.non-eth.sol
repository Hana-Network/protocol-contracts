// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "./HanaConnector.base.sol";
import "./interfaces/HanaInterfaces.sol";
import "./interfaces/HanaNonEthInterface.sol";

/**
 * @dev Non ETH implementation of HanaConnector.
 * This contract manages interactions between TSS and different chains.
 * This version is for every chain but Etherum network because in the other chains we mint and burn and in Etherum we lock and unlock
 */
contract HanaConnectorNonEth is HanaConnectorBase {
    uint256 public maxSupply = 2 ** 256 - 1;

    constructor(
        address hanaTokenAddress_,
        address tssAddress_,
        address tssAddressUpdater_,
        address pauserAddress_
    ) HanaConnectorBase(hanaTokenAddress_, tssAddress_, tssAddressUpdater_, pauserAddress_) {}

    function getLockedAmount() external view returns (uint256) {
        return HanaNonEthInterface(hanaToken).balanceOf(address(this));
    }

    function setMaxSupply(uint256 maxSupply_) external onlyTssAddress {
        maxSupply = maxSupply_;
    }

    /**
     * @dev Entry point to send data to protocol
     * This call burn the token and emit an event with all the data needed by the protocol
     */
    function send(HanaInterfaces.SendInput calldata input) external override whenNotPaused {
        HanaNonEthInterface(hanaToken).burnFrom(msg.sender, input.hanaValueAndGas);

        emit HanaSent(
            tx.origin,
            msg.sender,
            input.destinationChainId,
            input.destinationAddress,
            input.hanaValueAndGas,
            input.destinationGasLimit,
            input.message,
            input.hanaParams
        );
    }

    /**
     * @dev Handler to receive data from other chain.
     * This method can be called only by TSS.
     * Transfer the Hana tokens to destination and calls onHanaMessage if it's needed.
     * To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain.
     */
    function onReceive(
        bytes calldata hanaTxSenderAddress,
        uint256 sourceChainId,
        address destinationAddress,
        uint256 hanaValue,
        bytes calldata message,
        bytes32 internalSendHash
    ) external override onlyTssAddress {
        if (hanaValue + HanaNonEthInterface(hanaToken).totalSupply() > maxSupply) revert ExceedsMaxSupply(maxSupply);
        HanaNonEthInterface(hanaToken).mint(destinationAddress, hanaValue, internalSendHash);

        if (message.length > 0) {
            HanaReceiver(destinationAddress).onHanaMessage(
                HanaInterfaces.HanaMessage(hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message)
            );
        }

        emit HanaReceived(hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message, internalSendHash);
    }

    /**
     * @dev Handler to receive errors from other chain.
     * This method can be called only by TSS.
     * Transfer the Hana tokens to destination and calls onHanaRevert if it's needed.
     * To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain.
     */
    function onRevert(
        address hanaTxSenderAddress,
        uint256 sourceChainId,
        bytes calldata destinationAddress,
        uint256 destinationChainId,
        uint256 remainingHanaValue,
        bytes calldata message,
        bytes32 internalSendHash
    ) external override whenNotPaused onlyTssAddress {
        if (remainingHanaValue + HanaNonEthInterface(hanaToken).totalSupply() > maxSupply)
            revert ExceedsMaxSupply(maxSupply);
        HanaNonEthInterface(hanaToken).mint(hanaTxSenderAddress, remainingHanaValue, internalSendHash);

        if (message.length > 0) {
            HanaReceiver(hanaTxSenderAddress).onHanaRevert(
                HanaInterfaces.HanaRevert(
                    hanaTxSenderAddress,
                    sourceChainId,
                    destinationAddress,
                    destinationChainId,
                    remainingHanaValue,
                    message
                )
            );
        }

        emit HanaReverted(
            hanaTxSenderAddress,
            sourceChainId,
            destinationChainId,
            destinationAddress,
            remainingHanaValue,
            message,
            internalSendHash
        );
    }
}
