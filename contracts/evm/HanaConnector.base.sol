// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/security/Pausable.sol";

import "./interfaces/ConnectorErrors.sol";
import "./interfaces/HanaInterfaces.sol";

/**
 * @dev Main abstraction of HanaConnector.
 * This contract manages interactions between TSS and different chains.
 * There's an instance of this contract on each chain supported by HanaNetwork.
 */
contract HanaConnectorBase is ConnectorErrors, Pausable {
    address public immutable hanaToken;

    /**
     * @dev Multisig contract to pause incoming transactions.
     * The responsibility of pausing outgoing transactions is left to the protocol for more flexibility.
     */
    address public pauserAddress;

    /**
     * @dev Collectively held by HanaNetwork validators.
     */
    address public tssAddress;

    /**
     * @dev This address will start pointing to a multisig contract, then it will become the TSS address itself.
     */
    address public tssAddressUpdater;

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

    event TSSAddressUpdated(address callerAddress, address newTssAddress);

    event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress);

    event PauserAddressUpdated(address callerAddress, address newTssAddress);

    /**
     * @dev Constructor requires initial addresses.
     * hanaToken address is the only immutable one, while others can be updated.
     */
    constructor(address hanaToken_, address tssAddress_, address tssAddressUpdater_, address pauserAddress_) {
        if (
            hanaToken_ == address(0) ||
            tssAddress_ == address(0) ||
            tssAddressUpdater_ == address(0) ||
            pauserAddress_ == address(0)
        ) {
            revert HanaCommonErrors.InvalidAddress();
        }

        hanaToken = hanaToken_;
        tssAddress = tssAddress_;
        tssAddressUpdater = tssAddressUpdater_;
        pauserAddress = pauserAddress_;
    }

    /**
     * @dev Modifier to restrict actions to pauser address.
     */
    modifier onlyPauser() {
        if (msg.sender != pauserAddress) revert CallerIsNotPauser(msg.sender);
        _;
    }

    /**
     * @dev Modifier to restrict actions to TSS address.
     */
    modifier onlyTssAddress() {
        if (msg.sender != tssAddress) revert CallerIsNotTss(msg.sender);
        _;
    }

    /**
     * @dev Modifier to restrict actions to TSS updater address.
     */
    modifier onlyTssUpdater() {
        if (msg.sender != tssAddressUpdater) revert CallerIsNotTssUpdater(msg.sender);
        _;
    }

    /**
     * @dev Update the pauser address. The only address allowed to do that is the current pauser.
     */
    function updatePauserAddress(address pauserAddress_) external onlyPauser {
        if (pauserAddress_ == address(0)) revert HanaCommonErrors.InvalidAddress();

        pauserAddress = pauserAddress_;

        emit PauserAddressUpdated(msg.sender, pauserAddress_);
    }

    /**
     * @dev Update the TSS address. The address can be updated by the TSS updater or the TSS address itself.
     */
    function updateTssAddress(address tssAddress_) external {
        if (msg.sender != tssAddress && msg.sender != tssAddressUpdater) revert CallerIsNotTssOrUpdater(msg.sender);
        if (tssAddress_ == address(0)) revert HanaCommonErrors.InvalidAddress();

        tssAddress = tssAddress_;

        emit TSSAddressUpdated(msg.sender, tssAddress_);
    }

    /**
     * @dev Changes the ownership of tssAddressUpdater to be the one held by the HanaNetwork TSS Signer nodes.
     */
    function renounceTssAddressUpdater() external onlyTssUpdater {
        if (tssAddress == address(0)) revert HanaCommonErrors.InvalidAddress();

        tssAddressUpdater = tssAddress;
        emit TSSAddressUpdaterUpdated(msg.sender, tssAddressUpdater);
    }

    /**
     * @dev Pause the input (send) transactions.
     */

    function pause() external onlyPauser {
        _pause();
    }

    /**
     * @dev Unpause the contract to allow transactions again.
     */

    function unpause() external onlyPauser {
        _unpause();
    }

    /**
     * @dev Entrypoint to send data and value through HanaNetwork.
     */
    function send(HanaInterfaces.SendInput calldata input) external virtual {}

    /**
     * @dev Handler to receive data from other chain.
     * This method can be called only by TSS. Access validation is in implementation.
     */
    function onReceive(
        bytes calldata hanaTxSenderAddress,
        uint256 sourceChainId,
        address destinationAddress,
        uint256 hanaValue,
        bytes calldata message,
        bytes32 internalSendHash
    ) external virtual {}

    /**
     * @dev Handler to receive errors from other chain.
     * This method can be called only by TSS. Access validation is in implementation.
     */
    function onRevert(
        address hanaTxSenderAddress,
        uint256 sourceChainId,
        bytes calldata destinationAddress,
        uint256 destinationChainId,
        uint256 remainingHanaValue,
        bytes calldata message,
        bytes32 internalSendHash
    ) external virtual {}
}
