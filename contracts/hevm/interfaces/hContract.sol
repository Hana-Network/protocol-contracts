// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

struct hContext {
    bytes origin;
    address sender;
    uint256 chainID;
}

interface hContract {
    function onCrossChainCall(
        hContext calldata context,
        address hrc20,
        uint256 amount,
        bytes calldata message
    ) external;
}
