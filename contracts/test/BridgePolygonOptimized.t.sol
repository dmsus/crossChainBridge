// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Test.sol";
import "../src/BridgePolygonOptimized.sol";
import "../src/TestToken.sol";

contract BridgePolygonOptimizedTest is Test {
    BridgePolygonOptimized public bridge;
    TestToken public token;
    
    address owner = address(0x111);
    address signer = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
    address user = address(0x333);
    
    uint256 constant TEST_NONCE = 123;
    uint256 constant TEST_AMOUNT = 100 * 1e18;
    
    // Приватные ключи для тестовых адресов
    uint256 constant SIGNER_PRIVATE_KEY = 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d;
    uint256 constant WRONG_SIGNER_PRIVATE_KEY = 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a;
    
    function setUp() public {
        vm.deal(owner, 10 ether);
        vm.deal(signer, 10 ether);
        vm.deal(user, 10 ether);
        
        vm.prank(owner);
        token = new TestToken("Test Token", "TEST", 18);
        
        vm.prank(owner);
        bridge = new BridgePolygonOptimized(address(token), signer);
        
        // Mint tokens to bridge for testing
        token.mint(address(bridge), 1000 * 1e18);
    }
    
    function testReleaseTokens() public {
        bytes32 digest = bridge.getDigest(user, TEST_AMOUNT, TEST_NONCE);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(SIGNER_PRIVATE_KEY, digest);
        bytes memory signature = abi.encodePacked(r, s, v);
        
        vm.prank(user);
        bridge.releaseTokens(user, TEST_AMOUNT, TEST_NONCE, signature);
        
        assertEq(token.balanceOf(user), TEST_AMOUNT);
        assertTrue(bridge.isNonceProcessed(TEST_NONCE));
    }
    
    function testRevertWhenInvalidSignature() public {
        bytes32 digest = bridge.getDigest(user, TEST_AMOUNT, TEST_NONCE);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(WRONG_SIGNER_PRIVATE_KEY, digest);
        bytes memory signature = abi.encodePacked(r, s, v);
        
        vm.prank(user);
        vm.expectRevert(BridgePolygonOptimized.InvalidSignature.selector);
        bridge.releaseTokens(user, TEST_AMOUNT, TEST_NONCE, signature);
    }
    
    function testPauseUnpause() public {
        vm.prank(owner);
        bridge.pauseBridge();
        assertTrue(bridge.isPaused());
        
        vm.prank(owner);
        bridge.unpauseBridge();
        assertFalse(bridge.isPaused());
    }
    
    function testNonOwnerCannotPause() public {
        vm.prank(user);
        vm.expectRevert("Ownable: caller is not the owner");
        bridge.pauseBridge();
    }
    
    function testBridgeDeploys() public {
        assertEq(address(bridge.token()), address(token));
        assertEq(bridge.signer(), signer);
        assertFalse(bridge.isPaused());
    }
}
