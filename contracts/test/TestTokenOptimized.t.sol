// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Test.sol";
import "../src/TestTokenOptimized.sol";

contract TestTokenOptimizedTest is Test {
    TestTokenOptimized public token;
    
    address owner = address(0x111);
    address user1 = address(0x222);
    address user2 = address(0x333);
    
    uint256 constant MINT_AMOUNT = 1000 * 1e18;
    uint256 constant TRANSFER_AMOUNT = 100 * 1e18;
    
    function setUp() public {
        vm.prank(owner);
        token = new TestTokenOptimized("Test Token", "TEST", 18);
    }
    
    function test_Mint() public {
        uint256 initialSupply = token.totalSupply();
        
        vm.prank(owner);
        token.mint(user1, MINT_AMOUNT);
        
        assertEq(token.totalSupply(), initialSupply + MINT_AMOUNT);
        assertEq(token.balanceOf(user1), MINT_AMOUNT);
    }
    
    function test_Transfer() public {
        vm.prank(owner);
        token.mint(user1, MINT_AMOUNT);
        
        vm.prank(user1);
        bool success = token.transfer(user2, TRANSFER_AMOUNT);
        
        assertTrue(success);
        assertEq(token.balanceOf(user1), MINT_AMOUNT - TRANSFER_AMOUNT);
        assertEq(token.balanceOf(user2), TRANSFER_AMOUNT);
    }
    
    function test_TransferFrom() public {
        vm.prank(owner);
        token.mint(user1, MINT_AMOUNT);
        
        vm.prank(user1);
        token.approve(address(this), TRANSFER_AMOUNT);
        
        bool success = token.transferFrom(user1, user2, TRANSFER_AMOUNT);
        
        assertTrue(success);
        assertEq(token.balanceOf(user1), MINT_AMOUNT - TRANSFER_AMOUNT);
        assertEq(token.balanceOf(user2), TRANSFER_AMOUNT);
        assertEq(token.allowance(user1, address(this)), 0);
    }
    
    function test_Decimals() public {
        assertEq(token.decimals(), 18);
    }
}
