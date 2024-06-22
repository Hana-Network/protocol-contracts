// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/interfaces/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";

import "../interfaces/HanaInterfaces.sol";

interface HanaTokenConsumerUniV2Errors {
    error InputCantBeZero();
}

/**
 * @dev Uniswap V2 strategy for HanaTokenConsumer
 */
contract HanaTokenConsumerUniV2 is HanaTokenConsumer, HanaTokenConsumerUniV2Errors {
    using SafeERC20 for IERC20;
    uint256 internal constant MAX_DEADLINE = 200;

    address internal immutable wETH;
    address public immutable hanaToken;

    IUniswapV2Router02 internal immutable uniswapV2Router;

    constructor(address hanaToken_, address uniswapV2Router_) {
        if (hanaToken_ == address(0) || uniswapV2Router_ == address(0)) revert HanaCommonErrors.InvalidAddress();

        hanaToken = hanaToken_;
        uniswapV2Router = IUniswapV2Router02(uniswapV2Router_);
        wETH = IUniswapV2Router02(uniswapV2Router_).WETH();
    }

    function getHanaFromEth(
        address destinationAddress,
        uint256 minAmountOut
    ) external payable override returns (uint256) {
        if (destinationAddress == address(0)) revert HanaCommonErrors.InvalidAddress();
        if (msg.value == 0) revert InputCantBeZero();

        address[] memory path = new address[](2);
        path[0] = wETH;
        path[1] = hanaToken;

        uint256[] memory amounts = uniswapV2Router.swapExactETHForTokens{value: msg.value}(
            minAmountOut,
            path,
            destinationAddress,
            block.timestamp + MAX_DEADLINE
        );

        uint256 amountOut = amounts[path.length - 1];

        emit EthExchangedForHana(msg.value, amountOut);
        return amountOut;
    }

    function getHanaFromToken(
        address destinationAddress,
        uint256 minAmountOut,
        address inputToken,
        uint256 inputTokenAmount
    ) external override returns (uint256) {
        if (destinationAddress == address(0) || inputToken == address(0)) revert HanaCommonErrors.InvalidAddress();
        if (inputTokenAmount == 0) revert InputCantBeZero();

        IERC20(inputToken).safeTransferFrom(msg.sender, address(this), inputTokenAmount);
        IERC20(inputToken).safeApprove(address(uniswapV2Router), inputTokenAmount);

        address[] memory path;
        if (inputToken == wETH) {
            path = new address[](2);
            path[0] = wETH;
            path[1] = hanaToken;
        } else {
            path = new address[](3);
            path[0] = inputToken;
            path[1] = wETH;
            path[2] = hanaToken;
        }

        uint256[] memory amounts = uniswapV2Router.swapExactTokensForTokens(
            inputTokenAmount,
            minAmountOut,
            path,
            destinationAddress,
            block.timestamp + MAX_DEADLINE
        );
        uint256 amountOut = amounts[path.length - 1];

        emit TokenExchangedForHana(inputToken, inputTokenAmount, amountOut);
        return amountOut;
    }

    function getEthFromHana(
        address destinationAddress,
        uint256 minAmountOut,
        uint256 hanaTokenAmount
    ) external override returns (uint256) {
        if (destinationAddress == address(0)) revert HanaCommonErrors.InvalidAddress();
        if (hanaTokenAmount == 0) revert InputCantBeZero();

        IERC20(hanaToken).safeTransferFrom(msg.sender, address(this), hanaTokenAmount);
        IERC20(hanaToken).safeApprove(address(uniswapV2Router), hanaTokenAmount);

        address[] memory path = new address[](2);
        path[0] = hanaToken;
        path[1] = wETH;

        uint256[] memory amounts = uniswapV2Router.swapExactTokensForETH(
            hanaTokenAmount,
            minAmountOut,
            path,
            destinationAddress,
            block.timestamp + MAX_DEADLINE
        );

        uint256 amountOut = amounts[path.length - 1];

        emit HanaExchangedForEth(hanaTokenAmount, amountOut);
        return amountOut;
    }

    function getTokenFromHana(
        address destinationAddress,
        uint256 minAmountOut,
        address outputToken,
        uint256 hanaTokenAmount
    ) external override returns (uint256) {
        if (destinationAddress == address(0) || outputToken == address(0)) revert HanaCommonErrors.InvalidAddress();
        if (hanaTokenAmount == 0) revert InputCantBeZero();

        IERC20(hanaToken).safeTransferFrom(msg.sender, address(this), hanaTokenAmount);
        IERC20(hanaToken).safeApprove(address(uniswapV2Router), hanaTokenAmount);

        address[] memory path;
        if (outputToken == wETH) {
            path = new address[](2);
            path[0] = hanaToken;
            path[1] = wETH;
        } else {
            path = new address[](3);
            path[0] = hanaToken;
            path[1] = wETH;
            path[2] = outputToken;
        }

        uint256[] memory amounts = uniswapV2Router.swapExactTokensForTokens(
            hanaTokenAmount,
            minAmountOut,
            path,
            destinationAddress,
            block.timestamp + MAX_DEADLINE
        );

        uint256 amountOut = amounts[path.length - 1];

        emit HanaExchangedForToken(outputToken, hanaTokenAmount, amountOut);
        return amountOut;
    }

    function hasHanaLiquidity() external view override returns (bool) {
        address[] memory path = new address[](2);
        path[0] = wETH;
        path[1] = hanaToken;

        try uniswapV2Router.getAmountsOut(1, path) returns (uint256[] memory amounts) {
            return amounts[path.length - 1] > 0;
        } catch {
            return false;
        }
    }
}
