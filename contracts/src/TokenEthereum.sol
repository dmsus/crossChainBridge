// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";

/**
 * @title TokenEthereum
 * @dev ERC20 token for Ethereum Sepolia bridge testing
 */
contract TokenEthereum is ERC20 {
    uint8 private _decimals;
    
    constructor() ERC20("Bridge Token Ethereum", "BRG-ETH") {
        _decimals = 18;
        _mint(msg.sender, 1000000 * 10 ** _decimals);
    }
    
    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }
    
    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}
