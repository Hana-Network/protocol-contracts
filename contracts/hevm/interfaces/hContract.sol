// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

interface hContract {
    function onCrossChainCall(address hrc20, uint256 amount, bytes calldata message) external;
}
