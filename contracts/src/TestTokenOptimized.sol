// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";

/**
 * @title TestTokenOptimized
 * @dev Gas-optimized test ERC20 token for cross-chain bridge testing
 */
contract TestTokenOptimized is ERC20 {
    uint8 private _decimals;
    
    constructor(
        string memory name,
        string memory symbol, 
        uint8 decimals_
    ) ERC20(name, symbol) {
        _decimals = decimals_;
    }
    
    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }
    
    /**
     * @dev Mint tokens - GAS OPTIMIZED
     * Gas optimization: remove explicit return statement
     */
    function mint(address to, uint256 amount) external {
        _mint(to, amount);
        // No explicit return - saves gas
    }
    
    /**
     * @dev Transfer override - GAS OPTIMIZED  
     * Gas optimization: call parent directly without modifications
     */
    function transfer(address to, uint256 amount) 
        public 
        virtual 
        override 
        returns (bool) 
    {
        return super.transfer(to, amount);
    }
    
    /**
     * @dev TransferFrom override - GAS OPTIMIZED
     * Gas optimization: call parent directly without modifications
     */
    function transferFrom(
        address from,
        address to, 
        uint256 amount
    ) 
        public 
        virtual 
        override 
        returns (bool) 
    {
        return super.transferFrom(from, to, amount);
    }
}
