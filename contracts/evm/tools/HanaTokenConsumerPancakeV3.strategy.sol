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

interface ISwapRouterPancake is IUniswapV3SwapCallback {
    struct ExactInputSingleParams {
        address tokenIn;
        address tokenOut;
        uint24 fee;
        address recipient;
        uint256 amountIn;
        uint256 amountOutMinimum;
        uint160 sqrtPriceLimitX96;
    }

    /// @notice Swaps `amountIn` of one token for as much as possible of another token
    /// @param params The parameters necessary for the swap, encoded as `ExactInputSingleParams` in calldata
    /// @return amountOut The amount of the received token
    function exactInputSingle(ExactInputSingleParams calldata params) external payable returns (uint256 amountOut);

    struct ExactInputParams {
        bytes path;
        address recipient;
        uint256 amountIn;
        uint256 amountOutMinimum;
    }

    /// @notice Swaps `amountIn` of one token for as much as possible of another along the specified path
    /// @param params The parameters necessary for the multi-hop swap, encoded as `ExactInputParams` in calldata
    /// @return amountOut The amount of the received token
    function exactInput(ExactInputParams calldata params) external payable returns (uint256 amountOut);
}

/**
 * @dev Uniswap V3 strategy for HanaTokenConsumer
 */
contract HanaTokenConsumerPancakeV3 is HanaTokenConsumer, HanaTokenConsumerUniV3Errors {
    using SafeERC20 for IERC20;
    uint256 internal constant MAX_DEADLINE = 200;

    uint24 public immutable hanaPoolFee;
    uint24 public immutable tokenPoolFee;

    address public immutable WETH9Address;
    address public immutable hanaToken;

    ISwapRouterPancake public immutable pancakeV3Router;
    IUniswapV3Factory public immutable uniswapV3Factory;

    bool internal _locked;

    constructor(
        address hanaToken_,
        address pancakeV3Router_,
        address uniswapV3Factory_,
        address WETH9Address_,
        uint24 hanaPoolFee_,
        uint24 tokenPoolFee_
    ) {
        if (
            hanaToken_ == address(0) ||
            pancakeV3Router_ == address(0) ||
            uniswapV3Factory_ == address(0) ||
            WETH9Address_ == address(0)
        ) revert HanaCommonErrors.InvalidAddress();

        hanaToken = hanaToken_;
        pancakeV3Router = ISwapRouterPancake(pancakeV3Router_);
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

        ISwapRouterPancake.ExactInputSingleParams memory params = ISwapRouterPancake.ExactInputSingleParams({
            tokenIn: WETH9Address,
            tokenOut: hanaToken,
            fee: hanaPoolFee,
            recipient: destinationAddress,
            amountIn: msg.value,
            amountOutMinimum: minAmountOut,
            sqrtPriceLimitX96: 0
        });

        uint256 amountOut = pancakeV3Router.exactInputSingle{value: msg.value}(params);

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
        IERC20(inputToken).safeApprove(address(pancakeV3Router), inputTokenAmount);

        ISwapRouterPancake.ExactInputParams memory params = ISwapRouterPancake.ExactInputParams({
            path: abi.encodePacked(inputToken, tokenPoolFee, WETH9Address, hanaPoolFee, hanaToken),
            recipient: destinationAddress,
            amountIn: inputTokenAmount,
            amountOutMinimum: minAmountOut
        });

        uint256 amountOut = pancakeV3Router.exactInput(params);

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
        IERC20(hanaToken).safeApprove(address(pancakeV3Router), hanaTokenAmount);

        ISwapRouterPancake.ExactInputSingleParams memory params = ISwapRouterPancake.ExactInputSingleParams({
            tokenIn: hanaToken,
            tokenOut: WETH9Address,
            fee: hanaPoolFee,
            recipient: address(this),
            amountIn: hanaTokenAmount,
            amountOutMinimum: minAmountOut,
            sqrtPriceLimitX96: 0
        });

        uint256 amountOut = pancakeV3Router.exactInputSingle(params);

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
        IERC20(hanaToken).safeApprove(address(pancakeV3Router), hanaTokenAmount);

        ISwapRouterPancake.ExactInputParams memory params = ISwapRouterPancake.ExactInputParams({
            path: abi.encodePacked(hanaToken, hanaPoolFee, WETH9Address, tokenPoolFee, outputToken),
            recipient: destinationAddress,
            amountIn: hanaTokenAmount,
            amountOutMinimum: minAmountOut
        });

        uint256 amountOut = pancakeV3Router.exactInput(params);

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
