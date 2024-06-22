// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

/**
 * @dev Interfaces of SystemContract and HRC20 to make easier to import.
 */
interface ISystem {
    function FUNGIBLE_MODULE_ADDRESS() external view returns (address);

    function wHanaContractAddress() external view returns (address);

    function uniswapv2FactoryAddress() external view returns (address);

    function gasPriceByChainId(uint256 chainID) external view returns (uint256);

    function gasCoinHRC20ByChainId(uint256 chainID) external view returns (address);

    function gasHanaPoolByChainId(uint256 chainID) external view returns (address);
}

interface IHRC20 {
    function totalSupply() external view returns (uint256);

    function balanceOf(address account) external view returns (uint256);

    function transfer(address recipient, uint256 amount) external returns (bool);

    function allowance(address owner, address spender) external view returns (uint256);

    function approve(address spender, uint256 amount) external returns (bool);

    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);

    function deposit(address to, uint256 amount) external returns (bool);

    function withdraw(bytes memory to, uint256 amount) external returns (bool);

    function withdrawGasFee() external view returns (address, uint256);

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    event Deposit(bytes from, address indexed to, uint256 value);
    event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee);
    event UpdatedSystemContract(address systemContract);
    event UpdatedGasLimit(uint256 gasLimit);
    event UpdatedProtocolFlatFee(uint256 protocolFlatFee);
}

abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }
}

interface IHRC20Metadata is IHRC20 {
    function name() external view returns (string memory);

    function symbol() external view returns (string memory);

    function decimals() external view returns (uint8);
}

/// @dev Coin types for HRC20. Hana value should not be used.
enum CoinType {
    Hana,
    Gas,
    ERC20
}

/**
 * @dev Any HanaNetwork Contract must implement this interface to allow SystemContract to interact with.
 * This is only required if the contract wants to interact with other chains.
 */
interface hContract {
    function onCrossChainCall(address hrc20, uint256 amount, bytes calldata message) external;
}
