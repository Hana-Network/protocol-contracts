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

interface WHANA {
    function transferFrom(address src, address dst, uint wad) external returns (bool);

    function withdraw(uint wad) external;
}

contract HanaConnectorHEVM is HanaInterfaces {
    /// @notice Contract custom errors.
    error OnlyWHANA();
    error WHANATransferFailed();
    error OnlyFungibleModule();
    error FailedHanaSent();

    /// @notice Fungible module address.
    address public constant FUNGIBLE_MODULE_ADDRESS = payable(0x735b14BB79463307AAcBED86DAf3322B1e6226aB);
    /// @notice WHANA token address.
    address public whana;

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
    event SetWHANA(address whana_);

    constructor(address _whana) {
        whana = _whana;
    }

    /// @dev Receive function to receive HANA from WETH9.withdraw().
    receive() external payable {
        if (msg.sender != whana) revert OnlyWHANA();
    }

    /**
     * @dev Sends HANA and bytes messages (to execute it) crosschain.
     * @param input, SendInput struct, checkout above.
     */
    function send(HanaInterfaces.SendInput calldata input) external {
        // Transfer whana to "fungible" module, which will be burnt by the protocol post processing via hooks.
        if (!WHANA(whana).transferFrom(msg.sender, address(this), input.hanaValueAndGas)) revert WHANATransferFailed();
        WHANA(whana).withdraw(input.hanaValueAndGas);
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
     * @dev Sends HANA and bytes messages (to execute it) crosschain.
     * @param whana_, new WHANA address.
     */
    function setWhanaAddress(address whana_) external {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert OnlyFungibleModule();
        whana = whana_;
        emit SetWHANA(whana_);
    }
}
