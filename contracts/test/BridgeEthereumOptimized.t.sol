// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Test.sol";
import "../src/BridgeEthereumOptimized.sol";
import "../src/TestToken.sol";

contract BridgeEthereumOptimizedTest is Test {
    BridgeEthereumOptimized public bridge;
    TestToken public token;
    
    address owner = address(0x111);
    address user = address(0x222);
    address target = address(0x333);
    
    uint256 constant TEST_AMOUNT = 100 * 1e18;
    uint256 constant TARGET_CHAIN_ID = 80002;
    
    function setUp() public {
        vm.deal(owner, 10 ether);
        vm.deal(user, 10 ether);
        
        vm.prank(owner);
        token = new TestToken("Test Token", "TEST", 18);
        
        vm.prank(owner);
        bridge = new BridgeEthereumOptimized(address(token));
        
        // Setup user with tokens and allowance
        token.mint(user, 1000 * 1e18);
        vm.prank(user);
        token.approve(address(bridge), type(uint256).max);
    }
    
    function test_LockTokens() public {
        uint256 initialNonce = bridge.nonce();
        uint256 initialBridgeBalance = token.balanceOf(address(bridge));
        
        vm.prank(user);
        bridge.lockTokens(TEST_AMOUNT, target, TARGET_CHAIN_ID);
        
        assertEq(bridge.nonce(), initialNonce + 1);
        assertEq(token.balanceOf(address(bridge)), initialBridgeBalance + TEST_AMOUNT);
        assertEq(token.balanceOf(user), 1000 * 1e18 - TEST_AMOUNT);
    }
    
    function test_PauseUnpause() public {
        vm.prank(owner);
        bridge.pauseBridge();
        assertTrue(bridge.isPaused());
        
        vm.prank(owner);
        bridge.unpauseBridge();
        assertFalse(bridge.isPaused());
    }
    
    function test_RevertWhen_ZeroAmount() public {
        vm.prank(user);
        vm.expectRevert(BridgeEthereumOptimized.ZeroAmount.selector);
        bridge.lockTokens(0, target, TARGET_CHAIN_ID);
    }
    
    function test_RevertWhen_ZeroAddress() public {
        vm.prank(user);
        vm.expectRevert(BridgeEthereumOptimized.ZeroAddress.selector);
        bridge.lockTokens(TEST_AMOUNT, address(0), TARGET_CHAIN_ID);
    }
    
    function test_RevertWhen_ZeroChainId() public {
        vm.prank(user);
        vm.expectRevert(BridgeEthereumOptimized.ZeroChainId.selector);
        bridge.lockTokens(TEST_AMOUNT, target, 0);
    }
    
    function test_RevertWhen_InsufficientBalance() public {
        vm.prank(user);
        vm.expectRevert(BridgeEthereumOptimized.InsufficientBalance.selector);
        bridge.lockTokens(10000 * 1e18, target, TARGET_CHAIN_ID);
    }
}
