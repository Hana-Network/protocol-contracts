// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

interface HanaInterfaces {
    /**
     * @dev Use SendInput to interact with the Connector: connector.send(SendInput)
     */
    struct SendInput {
        /// @dev Chain id of the destination chain. More about chain ids https://docs.hana.network/learn/glossary#chain-id
        uint256 destinationChainId;
        /// @dev Address receiving the message on the destination chain (expressed in bytes since it can be non-EVM)
        bytes destinationAddress;
        /// @dev Gas limit for the destination chain's transaction
        uint256 destinationGasLimit;
        /// @dev An encoded, arbitrary message to be parsed by the destination contract
        bytes message;
        /// @dev ZETA to be sent cross-chain + HanaNetwork gas fees + destination chain gas fees (expressed in ZETA)
        uint256 hanaValueAndGas;
        /// @dev Optional parameters for the HanaNetwork protocol
        bytes hanaParams;
    }

    /**
     * @dev Our Connector calls onHanaMessage with this struct as argument
     */
    struct HanaMessage {
        bytes hanaTxSenderAddress;
        uint256 sourceChainId;
        address destinationAddress;
        /// @dev Remaining ZETA from hanaValueAndGas after subtracting HanaNetwork gas fees and destination gas fees
        uint256 hanaValue;
        bytes message;
    }

    /**
     * @dev Our Connector calls onHanaRevert with this struct as argument
     */
    struct HanaRevert {
        address hanaTxSenderAddress;
        uint256 sourceChainId;
        bytes destinationAddress;
        uint256 destinationChainId;
        /// @dev Equals to: hanaValueAndGas - HanaNetwork gas fees - destination chain gas fees - source chain revert tx gas fees
        uint256 remainingHanaValue;
        bytes message;
    }
}

interface HanaConnector {
    /**
     * @dev Sending value and data cross-chain is as easy as calling connector.send(SendInput)
     */
    function send(HanaInterfaces.SendInput calldata input) external;
}

interface HanaReceiver {
    /**
     * @dev onHanaMessage is called when a cross-chain message reaches a contract
     */
    function onHanaMessage(HanaInterfaces.HanaMessage calldata hanaMessage) external;

    /**
     * @dev onHanaRevert is called when a cross-chain message reverts.
     * It's useful to rollback to the original state
     */
    function onHanaRevert(HanaInterfaces.HanaRevert calldata hanaRevert) external;
}

/**
 * @dev HanaTokenConsumer makes it easier to handle the following situations:
 *   - Getting Hana using native coin (to pay for destination gas while using `connector.send`)
 *   - Getting Hana using a token (to pay for destination gas while using `connector.send`)
 *   - Getting native coin using Hana (to return unused destination gas when `onHanaRevert` is executed)
 *   - Getting a token using Hana (to return unused destination gas when `onHanaRevert` is executed)
 * @dev The interface can be implemented using different strategies, like UniswapV2, UniswapV3, etc
 */
interface HanaTokenConsumer {
    event EthExchangedForHana(uint256 amountIn, uint256 amountOut);
    event TokenExchangedForHana(address token, uint256 amountIn, uint256 amountOut);
    event HanaExchangedForEth(uint256 amountIn, uint256 amountOut);
    event HanaExchangedForToken(address token, uint256 amountIn, uint256 amountOut);

    function getHanaFromEth(address destinationAddress, uint256 minAmountOut) external payable returns (uint256);

    function getHanaFromToken(
        address destinationAddress,
        uint256 minAmountOut,
        address inputToken,
        uint256 inputTokenAmount
    ) external returns (uint256);

    function getEthFromHana(
        address destinationAddress,
        uint256 minAmountOut,
        uint256 hanaTokenAmount
    ) external returns (uint256);

    function getTokenFromHana(
        address destinationAddress,
        uint256 minAmountOut,
        address outputToken,
        uint256 hanaTokenAmount
    ) external returns (uint256);

    function hasHanaLiquidity() external view returns (bool);
}

interface HanaCommonErrors {
    error InvalidAddress();
}
