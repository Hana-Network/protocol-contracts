// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "../interfaces/hContract.sol";
import "../interfaces/IHRC20.sol";

interface SystemContractErrors {
    error CallerIsNotFungibleModule();

    error InvalidTarget();

    error CantBeIdenticalAddresses();

    error CantBeZeroAddress();
}

contract SystemContractMock is SystemContractErrors {
    mapping(uint256 => uint256) public gasPriceByChainId;
    mapping(uint256 => address) public gasCoinHRC20ByChainId;
    mapping(uint256 => address) public gasHanaPoolByChainId;

    address public wHanaContractAddress;
    address public uniswapv2FactoryAddress;
    address public uniswapv2Router02Address;

    event SystemContractDeployed();
    event SetGasPrice(uint256, uint256);
    event SetGasCoin(uint256, address);
    event SetGasHanaPool(uint256, address);
    event SetWHana(address);

    constructor(address whana_, address uniswapv2Factory_, address uniswapv2Router02_) {
        wHanaContractAddress = whana_;
        uniswapv2FactoryAddress = uniswapv2Factory_;
        uniswapv2Router02Address = uniswapv2Router02_;
        emit SystemContractDeployed();
    }

    // fungible module updates the gas price oracle periodically
    function setGasPrice(uint256 chainID, uint256 price) external {
        gasPriceByChainId[chainID] = price;
        emit SetGasPrice(chainID, price);
    }

    function setGasCoinHRC20(uint256 chainID, address hrc20) external {
        gasCoinHRC20ByChainId[chainID] = hrc20;
        emit SetGasCoin(chainID, hrc20);
    }

    function setWHANAContractAddress(address addr) external {
        wHanaContractAddress = addr;
        emit SetWHana(wHanaContractAddress);
    }

    // returns sorted token addresses, used to handle return values from pairs sorted in this order
    function sortTokens(address tokenA, address tokenB) internal pure returns (address token0, address token1) {
        if (tokenA == tokenB) revert CantBeIdenticalAddresses();
        (token0, token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
        if (token0 == address(0)) revert CantBeZeroAddress();
    }

    function uniswapv2PairFor(address factory, address tokenA, address tokenB) public pure returns (address pair) {
        (address token0, address token1) = sortTokens(tokenA, tokenB);
        pair = address(
            uint160(
                uint256(
                    keccak256(
                        abi.encodePacked(
                            hex"ff",
                            factory,
                            keccak256(abi.encodePacked(token0, token1)),
                            hex"96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f" // init code hash
                        )
                    )
                )
            )
        );
    }

    function onCrossChainCall(address target, address hrc20, uint256 amount, bytes calldata message) external {
        hContext memory context = hContext({sender: msg.sender, origin: "", chainID: block.chainid});
        IHRC20(hrc20).transfer(target, amount);
        hContract(target).onCrossChainCall(context, hrc20, amount, message);
    }
}
