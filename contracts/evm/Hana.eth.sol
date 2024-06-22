// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @dev HanaEth is an implementation of OpenZeppelin's ERC20
 */
contract HanaEth is ERC20("Hana", "HANA") {
    constructor(address creator, uint256 initialSupply) {
        _mint(creator, initialSupply * (10 ** uint256(decimals())));
    }
}
