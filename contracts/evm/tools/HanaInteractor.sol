// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/access/Ownable2Step.sol";

import "../interfaces/HanaInterfaces.sol";
import "../interfaces/HanaInteractorErrors.sol";

abstract contract HanaInteractor is Ownable2Step, HanaInteractorErrors {
    bytes32 constant ZERO_BYTES = keccak256(new bytes(0));
    uint256 internal immutable currentChainId;
    HanaConnector public immutable connector;

    /**
     * @dev Maps a chain id to its corresponding address of the MultiChainSwap contract
     * The address is expressed in bytes to allow non-EVM chains
     * This mapping is useful, mainly, for two reasons:
     *  - Given a chain id, the contract is able to route a transaction to its corresponding address
     *  - To check that the messages (onHanaMessage, onHanaRevert) come from a trusted source
     */
    mapping(uint256 => bytes) public interactorsByChainId;

    modifier isValidMessageCall(HanaInterfaces.HanaMessage calldata hanaMessage) {
        _isValidCaller();
        if (keccak256(hanaMessage.hanaTxSenderAddress) != keccak256(interactorsByChainId[hanaMessage.sourceChainId]))
            revert InvalidHanaMessageCall();
        _;
    }

    modifier isValidRevertCall(HanaInterfaces.HanaRevert calldata hanaRevert) {
        _isValidCaller();
        if (hanaRevert.hanaTxSenderAddress != address(this)) revert InvalidHanaRevertCall();
        if (hanaRevert.sourceChainId != currentChainId) revert InvalidHanaRevertCall();
        _;
    }

    constructor(address hanaConnectorAddress) {
        if (hanaConnectorAddress == address(0)) revert HanaCommonErrors.InvalidAddress();
        currentChainId = block.chainid;
        connector = HanaConnector(hanaConnectorAddress);
    }

    function _isValidCaller() private view {
        if (msg.sender != address(connector)) revert InvalidCaller(msg.sender);
    }

    /**
     * @dev Useful for contracts that inherit from this one
     */
    function _isValidChainId(uint256 chainId) internal view returns (bool) {
        return (keccak256(interactorsByChainId[chainId]) != ZERO_BYTES);
    }

    function setInteractorByChainId(uint256 destinationChainId, bytes calldata contractAddress) external onlyOwner {
        interactorsByChainId[destinationChainId] = contractAddress;
    }
}
