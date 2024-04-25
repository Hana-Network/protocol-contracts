// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/interfaces/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";
import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";
import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";

import "../interfaces/HanaInterfaces.sol";

interface HanaTokenConsumerUniV3Errors {
    error InputCantBeZero();

    error ErrorSendingETH();

    error ReentrancyError();
}

interface WETH9 {
    function withdraw(uint256 wad) external;
}

/**
 * @dev Uniswap V3 strategy for HanaTokenConsumer
 */
contract HanaTokenConsumerUniV3 is HanaTokenConsumer, HanaTokenConsumerUniV3Errors {
    using SafeERC20 for IERC20;
    uint256 internal constant MAX_DEADLINE = 200;

    uint24 public immutable hanaPoolFee;
    uint24 public immutable tokenPoolFee;

    address public immutable WETH9Address;
    address public immutable hanaToken;

    ISwapRouter public immutable uniswapV3Router;
    IUniswapV3Factory public immutable uniswapV3Factory;

    bool internal _locked;

    constructor(
        address hanaToken_,
        address uniswapV3Router_,
        address uniswapV3Factory_,
        address WETH9Address_,
        uint24 hanaPoolFee_,
        uint24 tokenPoolFee_
    ) {
        if (
            hanaToken_ == address(0) ||
            uniswapV3Router_ == address(0) ||
            uniswapV3Factory_ == address(0) ||
            WETH9Address_ == address(0)
        ) revert HanaCommonErrors.InvalidAddress();

        hanaToken = hanaToken_;
        uniswapV3Router = ISwapRouter(uniswapV3Router_);
        uniswapV3Factory = IUniswapV3Factory(uniswapV3Factory_);
        WETH9Address = WETH9Address_;
        hanaPoolFee = hanaPoolFee_;
        tokenPoolFee = tokenPoolFee_;
    }

    modifier nonReentrant() {
        if (_locked) revert ReentrancyError();
        _locked = true;
        _;
        _locked = false;
    }

    receive() external payable {}

    function getHanaFromEth(
        address destinationAddress,
        uint256 minAmountOut
    ) external payable override returns (uint256) {
        if (destinationAddress == address(0)) revert HanaCommonErrors.InvalidAddress();
        if (msg.value == 0) revert InputCantBeZero();

        ISwapRouter.ExactInputSingleParams memory params = ISwapRouter.ExactInputSingleParams({
            deadline: block.timestamp + MAX_DEADLINE,
            tokenIn: WETH9Address,
            tokenOut: hanaToken,
            fee: hanaPoolFee,
            recipient: destinationAddress,
            amountIn: msg.value,
            amountOutMinimum: minAmountOut,
            sqrtPriceLimitX96: 0
        });

        uint256 amountOut = uniswapV3Router.exactInputSingle{value: msg.value}(params);

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
        IERC20(inputToken).safeApprove(address(uniswapV3Router), inputTokenAmount);

        ISwapRouter.ExactInputParams memory params = ISwapRouter.ExactInputParams({
            deadline: block.timestamp + MAX_DEADLINE,
            path: abi.encodePacked(inputToken, tokenPoolFee, WETH9Address, hanaPoolFee, hanaToken),
            recipient: destinationAddress,
            amountIn: inputTokenAmount,
            amountOutMinimum: minAmountOut
        });

        uint256 amountOut = uniswapV3Router.exactInput(params);

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
        IERC20(hanaToken).safeApprove(address(uniswapV3Router), hanaTokenAmount);

        ISwapRouter.ExactInputSingleParams memory params = ISwapRouter.ExactInputSingleParams({
            deadline: block.timestamp + MAX_DEADLINE,
            tokenIn: hanaToken,
            tokenOut: WETH9Address,
            fee: hanaPoolFee,
            recipient: address(this),
            amountIn: hanaTokenAmount,
            amountOutMinimum: minAmountOut,
            sqrtPriceLimitX96: 0
        });

        uint256 amountOut = uniswapV3Router.exactInputSingle(params);

        WETH9(WETH9Address).withdraw(amountOut);

        emit HanaExchangedForEth(hanaTokenAmount, amountOut);

        (bool sent, ) = destinationAddress.call{value: amountOut}("");
        if (!sent) revert ErrorSendingETH();

        return amountOut;
    }

    function getTokenFromHana(
        address destinationAddress,
        uint256 minAmountOut,
        address outputToken,
        uint256 hanaTokenAmount
    ) external override nonReentrant returns (uint256) {
        if (destinationAddress == address(0) || outputToken == address(0)) revert HanaCommonErrors.InvalidAddress();
        if (hanaTokenAmount == 0) revert InputCantBeZero();

        IERC20(hanaToken).safeTransferFrom(msg.sender, address(this), hanaTokenAmount);
        IERC20(hanaToken).safeApprove(address(uniswapV3Router), hanaTokenAmount);

        ISwapRouter.ExactInputParams memory params = ISwapRouter.ExactInputParams({
            deadline: block.timestamp + MAX_DEADLINE,
            path: abi.encodePacked(hanaToken, hanaPoolFee, WETH9Address, tokenPoolFee, outputToken),
            recipient: destinationAddress,
            amountIn: hanaTokenAmount,
            amountOutMinimum: minAmountOut
        });

        uint256 amountOut = uniswapV3Router.exactInput(params);

        emit HanaExchangedForToken(outputToken, hanaTokenAmount, amountOut);
        return amountOut;
    }

    function hasHanaLiquidity() external view override returns (bool) {
        address poolAddress = uniswapV3Factory.getPool(WETH9Address, hanaToken, hanaPoolFee);

        if (poolAddress == address(0)) {
            return false;
        }

        //@dev: if pool does exist, get its liquidity
        IUniswapV3Pool pool = IUniswapV3Pool(poolAddress);
        return pool.liquidity() > 0;
    }
}
