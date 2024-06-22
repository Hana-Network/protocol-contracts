// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/interfaces/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@uniswap/v3-periphery/contracts/interfaces/IQuoter.sol";

import "../interfaces/HanaInterfaces.sol";
import "./interfaces/TridentConcentratedLiquidityPoolFactory.sol";
import "./interfaces/TridentIPoolRouter.sol";

interface HanaTokenConsumerTridentErrors {
    error InputCantBeZero();

    error ErrorSendingETH();

    error ReentrancyError();
}

interface WETH9 {
    function deposit() external payable;

    function withdraw(uint256 wad) external;

    function depositTo(address to) external payable;

    function withdrawTo(address payable to, uint256 value) external;
}

/**
 * @dev Trident strategy for HanaTokenConsumer
 */
contract HanaTokenConsumerTrident is HanaTokenConsumer, HanaTokenConsumerTridentErrors {
    using SafeERC20 for IERC20;
    uint256 internal constant MAX_DEADLINE = 200;

    address internal immutable WETH9Address;
    address public immutable hanaToken;

    IPoolRouter public immutable tridentRouter;
    ConcentratedLiquidityPoolFactory public immutable poolFactory;

    bool internal _locked;

    constructor(address hanaToken_, address uniswapV3Router_, address WETH9Address_, address poolFactory_) {
        if (
            hanaToken_ == address(0) ||
            uniswapV3Router_ == address(0) ||
            WETH9Address_ == address(0) ||
            poolFactory_ == address(0)
        ) revert HanaCommonErrors.InvalidAddress();

        hanaToken = hanaToken_;
        tridentRouter = IPoolRouter(uniswapV3Router_);
        poolFactory = ConcentratedLiquidityPoolFactory(poolFactory_);
        WETH9Address = WETH9Address_;
    }

    modifier nonReentrant() {
        if (_locked) revert ReentrancyError();
        _locked = true;
        _;
        _locked = false;
    }

    receive() external payable {}

    function getPair(address token0, address token1) internal pure returns (address, address) {
        if (token0 < token1) return (token0, token1);

        return (token1, token0);
    }

    function getHanaFromEth(
        address destinationAddress,
        uint256 minAmountOut
    ) external payable override returns (uint256) {
        if (destinationAddress == address(0)) revert HanaCommonErrors.InvalidAddress();
        if (msg.value == 0) revert InputCantBeZero();

        (address token0, address token1) = getPair(WETH9Address, hanaToken);
        address[] memory pairPools = poolFactory.getPools(token0, token1, 0, 1);

        IPoolRouter.ExactInputSingleParams memory params = IPoolRouter.ExactInputSingleParams({
            tokenIn: address(0),
            amountIn: msg.value,
            amountOutMinimum: minAmountOut,
            pool: pairPools[0],
            to: destinationAddress,
            unwrap: false
        });

        uint256 amountOut = tridentRouter.exactInputSingle{value: msg.value}(params);

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
        IERC20(inputToken).safeApprove(address(tridentRouter), inputTokenAmount);

        (address token0, address token1) = getPair(inputToken, WETH9Address);
        address[] memory pairPools1 = poolFactory.getPools(token0, token1, 0, 1);

        (token0, token1) = getPair(WETH9Address, hanaToken);
        address[] memory pairPools2 = poolFactory.getPools(token0, token1, 0, 1);

        address[] memory path = new address[](2);
        path[0] = pairPools1[0];
        path[1] = pairPools2[0];

        IPoolRouter.ExactInputParams memory params = IPoolRouter.ExactInputParams({
            tokenIn: inputToken,
            amountIn: inputTokenAmount,
            amountOutMinimum: minAmountOut,
            path: path,
            to: destinationAddress,
            unwrap: false
        });

        uint256 amountOut = tridentRouter.exactInput(params);

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
        IERC20(hanaToken).safeApprove(address(tridentRouter), hanaTokenAmount);

        (address token0, address token1) = getPair(hanaToken, WETH9Address);
        address[] memory pairPools = poolFactory.getPools(token0, token1, 0, 1);

        IPoolRouter.ExactInputSingleParams memory params = IPoolRouter.ExactInputSingleParams({
            tokenIn: hanaToken,
            amountIn: hanaTokenAmount,
            amountOutMinimum: minAmountOut,
            pool: pairPools[0],
            to: destinationAddress,
            unwrap: true
        });

        uint256 amountOut = tridentRouter.exactInputSingle(params);

        emit HanaExchangedForEth(hanaTokenAmount, amountOut);

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
        IERC20(hanaToken).safeApprove(address(tridentRouter), hanaTokenAmount);

        (address token0, address token1) = getPair(hanaToken, WETH9Address);
        address[] memory pairPools1 = poolFactory.getPools(token0, token1, 0, 1);

        (token0, token1) = getPair(WETH9Address, outputToken);
        address[] memory pairPools2 = poolFactory.getPools(token0, token1, 0, 1);

        address[] memory path = new address[](2);
        path[0] = pairPools1[0];
        path[1] = pairPools2[0];

        IPoolRouter.ExactInputParams memory params = IPoolRouter.ExactInputParams({
            tokenIn: hanaToken,
            amountIn: hanaTokenAmount,
            amountOutMinimum: minAmountOut,
            path: path,
            to: destinationAddress,
            unwrap: false
        });

        uint256 amountOut = tridentRouter.exactInput(params);

        emit HanaExchangedForToken(outputToken, hanaTokenAmount, amountOut);
        return amountOut;
    }

    function hasHanaLiquidity() external view override returns (bool) {
        //@TODO: Implement
        return false;
    }
}
