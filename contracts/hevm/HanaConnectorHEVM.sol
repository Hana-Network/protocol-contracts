// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;
import "./interfaces/IWHANA.sol";

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
        /// @dev HANA to be sent cross-chain + HanaNetwork gas fees + destination chain gas fees (expressed in HANA)
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
        /// @dev Remaining HANA from hanaValueAndGas after subtracting HanaNetwork gas fees and destination gas fees
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

contract HanaConnectorHEVM {
    /// @notice WHANA token address.
    address public whana;

    /// @notice Contract custom errors.
    error OnlyWHANAOrFungible();
    error WHANATransferFailed();
    error OnlyFungibleModule();
    error FailedHanaSent();
    error WrongValue();

    /// @notice Fungible module address.
    address public constant FUNGIBLE_MODULE_ADDRESS = payable(0x735b14BB79463307AAcBED86DAf3322B1e6226aB);

    event SetWHANA(address whana_);

    event HanaSent(
        address sourceTxOriginAddress,
        address indexed hanaTxSenderAddress,
        uint256 indexed destinationChainId,
        bytes destinationAddress,
        uint256 hanaValueAndGas,
        uint256 destinationGasLimit,
        bytes message,
        bytes hanaParams
    );

    event HanaReceived(
        bytes hanaTxSenderAddress,
        uint256 indexed sourceChainId,
        address indexed destinationAddress,
        uint256 hanaValue,
        bytes message,
        bytes32 indexed internalSendHash
    );

    event HanaReverted(
        address hanaTxSenderAddress,
        uint256 sourceChainId,
        uint256 indexed destinationChainId,
        bytes destinationAddress,
        uint256 remainingHanaValue,
        bytes message,
        bytes32 indexed internalSendHash
    );

    /**
     * @dev Modifier to restrict actions to fungible module.
     */
    modifier onlyFungibleModule() {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert OnlyFungibleModule();
        _;
    }

    constructor(address whana_) {
        whana = whana_;
    }

    /// @dev Receive function to receive HANA from WETH9.withdraw().
    receive() external payable {
        if (msg.sender != whana && msg.sender != FUNGIBLE_MODULE_ADDRESS) revert OnlyWHANAOrFungible();
    }

    function setWhanaAddress(address whana_) external onlyFungibleModule {
        whana = whana_;
        emit SetWHANA(whana_);
    }

    /**
     * @dev Sends HANA and bytes messages (to execute it) crosschain.
     * @param input, SendInput struct, checkout above.
     */
    function send(HanaInterfaces.SendInput calldata input) external {
        // Transfer whana to "fungible" module, which will be burnt by the protocol post processing via hooks.
        if (!IWETH9(whana).transferFrom(msg.sender, address(this), input.hanaValueAndGas)) revert WHANATransferFailed();
        IWETH9(whana).withdraw(input.hanaValueAndGas);
        (bool sent, ) = FUNGIBLE_MODULE_ADDRESS.call{value: input.hanaValueAndGas}("");
        if (!sent) revert FailedHanaSent();
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
     * This method can be called only by Fungible Module.
     * Transfer the Hana tokens to destination and calls onHanaMessage if it's needed.
     * To perform the transfer wrap the new tokens
     */
    function onReceive(
        bytes calldata hanaTxSenderAddress,
        uint256 sourceChainId,
        address destinationAddress,
        uint256 hanaValue,
        bytes calldata message,
        bytes32 internalSendHash
    ) external payable onlyFungibleModule {
        if (msg.value != hanaValue) revert WrongValue();
        IWETH9(whana).deposit{value: hanaValue}();
        if (!IWETH9(whana).transferFrom(address(this), destinationAddress, hanaValue)) revert WHANATransferFailed();

        if (message.length > 0) {
            HanaReceiver(destinationAddress).onHanaMessage(
                HanaInterfaces.HanaMessage(hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message)
            );
        }

        emit HanaReceived(hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message, internalSendHash);
    }

    /**
     * @dev Handler to receive errors from other chain.
     * This method can be called only by Fungible Module.
     * Transfer the Hana tokens to destination and calls onHanaRevert if it's needed.
     */
    function onRevert(
        address hanaTxSenderAddress,
        uint256 sourceChainId,
        bytes calldata destinationAddress,
        uint256 destinationChainId,
        uint256 remainingHanaValue,
        bytes calldata message,
        bytes32 internalSendHash
    ) external payable onlyFungibleModule {
        if (msg.value != remainingHanaValue) revert WrongValue();
        IWETH9(whana).deposit{value: remainingHanaValue}();
        if (!IWETH9(whana).transferFrom(address(this), hanaTxSenderAddress, remainingHanaValue))
            revert WHANATransferFailed();

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
