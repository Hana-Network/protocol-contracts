// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

/**
 * @dev Interface with Hana Interactor errors
 */
interface HanaInteractorErrors {
    // @dev Thrown when try to send a message or tokens to a non whitelisted chain
    error InvalidDestinationChainId();

    // @dev Thrown when the caller is invalid. e.g.: if onHanaMessage or onHanaRevert are not called by Connector
    error InvalidCaller(address caller);

    // @dev Thrown when message on onHanaMessage has the wrong format
    error InvalidHanaMessageCall();

    // @dev Thrown when message on onHanaRevert has the wrong format
    error InvalidHanaRevertCall();
}
