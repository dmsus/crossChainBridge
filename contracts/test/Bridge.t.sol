// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Test.sol";
import "../src/BridgeEthereum.sol";
import "../src/TestToken.sol";

contract BridgeTest is Test {
    BridgeEthereum public bridge;
    TestToken public token;
    
    address owner = address(0x111);
    address user = address(0x222);
    address targetUser = address(0x333);
    
    uint256 constant POLYGON_AMOY_CHAIN_ID = 80002;
    
    function setUp() public {
        vm.deal(owner, 10 ether);
        vm.deal(user, 10 ether);
        
        vm.prank(owner);
        token = new TestToken("Test Token", "TEST", 18);
        
        vm.prank(owner);
        bridge = new BridgeEthereum(address(token));
        
        vm.prank(owner);
        token.transfer(user, 1000 * 1e18);
        
        vm.prank(user);
        token.approve(address(bridge), type(uint256).max);
    }
    
    function test_LockTokens() public {
        uint256 lockAmount = 100 * 1e18;
        
        vm.expectEmit(true, true, false, true);
        // Правильный emit - из контракта BridgeEthereum
        emit BridgeEthereum.TokensLocked(user, lockAmount, 1, targetUser, POLYGON_AMOY_CHAIN_ID);
        
        vm.prank(user);
        bridge.lockTokens(lockAmount, targetUser, POLYGON_AMOY_CHAIN_ID);
        
        assertEq(bridge.nonce(), 1);
        assertEq(token.balanceOf(user), 900 * 1e18);
        assertEq(token.balanceOf(address(bridge)), 100 * 1e18);
    }
    
    function test_RevertWhen_ZeroAmount() public {
        vm.prank(user);
        vm.expectRevert(BridgeEthereum.ZeroAmount.selector);
        bridge.lockTokens(0, targetUser, POLYGON_AMOY_CHAIN_ID);
    }
    
    function test_RevertWhen_ZeroAddress() public {
        vm.prank(user);
        vm.expectRevert(BridgeEthereum.ZeroAddress.selector);
        bridge.lockTokens(100 * 1e18, address(0), POLYGON_AMOY_CHAIN_ID);
    }
    
    function test_RevertWhen_ZeroChainId() public {
        vm.prank(user);
        vm.expectRevert(BridgeEthereum.ZeroChainId.selector);
        bridge.lockTokens(100 * 1e18, targetUser, 0);
    }
    
    function test_RevertWhen_InsufficientBalance() public {
        vm.prank(user);
        vm.expectRevert(BridgeEthereum.InsufficientBalance.selector);
        bridge.lockTokens(5000 * 1e18, targetUser, POLYGON_AMOY_CHAIN_ID);
    }
    
    function test_PauseUnpause() public {
        vm.prank(owner);
        bridge.pauseBridge();
        assertEq(bridge.isPaused(), true);
        
        vm.prank(user);
        vm.expectRevert(BridgeEthereum.BridgeIsPaused.selector);
        bridge.lockTokens(100 * 1e18, targetUser, POLYGON_AMOY_CHAIN_ID);
        
        vm.prank(owner);
        bridge.unpauseBridge();
        assertEq(bridge.isPaused(), false);
        
        vm.prank(user);
        bridge.lockTokens(100 * 1e18, targetUser, POLYGON_AMOY_CHAIN_ID);
    }
}
