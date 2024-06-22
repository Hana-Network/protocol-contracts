// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "./interfaces/ConnectorErrors.sol";
import "./HanaConnector.base.sol";
import "./interfaces/HanaInterfaces.sol";

/**
 * @dev ETH implementation of HanaConnector.
 * This contract manages interactions between TSS and different chains.
 * This version is only for Ethereum network because in the other chains we mint and burn and in this one we lock and unlock.
 */
contract HanaConnectorEth is HanaConnectorBase {
    constructor(
        address hanaToken_,
        address tssAddress_,
        address tssAddressUpdater_,
        address pauserAddress_
    ) HanaConnectorBase(hanaToken_, tssAddress_, tssAddressUpdater_, pauserAddress_) {}

    function getLockedAmount() external view returns (uint256) {
        return IERC20(hanaToken).balanceOf(address(this));
    }

    /**
     * @dev Entrypoint to send data through HanaNetwork
     * This call locks the token on the contract and emits an event with all the data needed by the protocol.
     */
    function send(HanaInterfaces.SendInput calldata input) external override whenNotPaused {
        bool success = IERC20(hanaToken).transferFrom(msg.sender, address(this), input.hanaValueAndGas);
        if (!success) revert HanaTranserError();

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
     * Transfers the Hana tokens to destination and calls onHanaMessage if it's needed.
     */
    function onReceive(
        bytes calldata hanaTxSenderAddress,
        uint256 sourceChainId,
        address destinationAddress,
        uint256 hanaValue,
        bytes calldata message,
        bytes32 internalSendHash
    ) external override onlyTssAddress {
        bool success = IERC20(hanaToken).transfer(destinationAddress, hanaValue);
        if (!success) revert HanaTranserError();

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
     * Transfers the Hana tokens to destination and calls onHanaRevert if it's needed.
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
        bool success = IERC20(hanaToken).transfer(hanaTxSenderAddress, remainingHanaValue);
        if (!success) revert HanaTranserError();

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
