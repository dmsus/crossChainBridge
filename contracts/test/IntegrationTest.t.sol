// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Test.sol";
import "../src/BridgeEthereum.sol";
import "../src/BridgePolygon.sol";
import "../src/TokenEthereum.sol";
import "../src/TokenPolygon.sol";

contract IntegrationTest is Test {
    BridgeEthereum public bridgeEth;
    BridgePolygon public bridgePoly;
    TokenEthereum public tokenEth;
    TokenPolygon public tokenPoly;
    
    address owner = address(0x111);
    address user = address(0x222);
    address relayer = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
    uint256 constant RELAYER_PRIVATE_KEY = 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d;
    
    uint256 constant POLYGON_CHAIN_ID = 80002;
    
    function setUp() public {
        vm.deal(owner, 10 ether);
        vm.deal(user, 10 ether);
        vm.deal(relayer, 10 ether);
        
        // Deploy Ethereum side
        vm.startPrank(owner);
        tokenEth = new TokenEthereum();
        bridgeEth = new BridgeEthereum(address(tokenEth));
        tokenEth.transfer(user, 1000 * 1e18);
        vm.stopPrank();
        
        // Deploy Polygon side
        vm.startPrank(owner);
        tokenPoly = new TokenPolygon();
        bridgePoly = new BridgePolygon(address(tokenPoly), relayer);
        tokenPoly.transfer(address(bridgePoly), 5000 * 1e18);
        vm.stopPrank();
        
        // User approves tokens
        vm.prank(user);
        tokenEth.approve(address(bridgeEth), type(uint256).max);
    }
    
    function test_CompleteCrossChainFlow() public {
        uint256 lockAmount = 100 * 1e18;
        uint256 initialEthBalance = tokenEth.balanceOf(user);
        uint256 initialPolyBalance = tokenPoly.balanceOf(user);
        
        // Step 1: User locks tokens on Ethereum
        vm.prank(user);
        bridgeEth.lockTokens(lockAmount, user, POLYGON_CHAIN_ID);
        
        // Verify Ethereum side
        assertEq(tokenEth.balanceOf(user), initialEthBalance - lockAmount);
        assertEq(tokenEth.balanceOf(address(bridgeEth)), lockAmount);
        assertEq(bridgeEth.nonce(), 1);
        
        // Step 2: Create signature as relayer would
        bytes32 digest = bridgePoly.getDigest(user, lockAmount, 1);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(RELAYER_PRIVATE_KEY, digest);
        bytes memory signature = abi.encodePacked(r, s, v);
        
        // Step 3: Release tokens on Polygon
        vm.prank(user);
        bridgePoly.releaseTokens(user, lockAmount, 1, signature);
        
        // Verify Polygon side
        assertEq(tokenPoly.balanceOf(user), initialPolyBalance + lockAmount);
        assertEq(bridgePoly.isNonceProcessed(1), true);
    }
    
    function test_RevertWhen_NonceAlreadyProcessed() public {
        uint256 lockAmount = 100 * 1e18;
        
        // First release with nonce 99
        bytes32 digest1 = bridgePoly.getDigest(user, lockAmount, 99);
        (uint8 v1, bytes32 r1, bytes32 s1) = vm.sign(RELAYER_PRIVATE_KEY, digest1);
        bytes memory signature1 = abi.encodePacked(r1, s1, v1);
        
        vm.prank(user);
        bridgePoly.releaseTokens(user, lockAmount, 99, signature1);
        
        // Try to release again with the same nonce 99
        bytes32 digest2 = bridgePoly.getDigest(user, lockAmount, 99);
        (uint8 v2, bytes32 r2, bytes32 s2) = vm.sign(RELAYER_PRIVATE_KEY, digest2);
        bytes memory signature2 = abi.encodePacked(r2, s2, v2);
        
        vm.prank(user);
        vm.expectRevert(BridgePolygon.NonceAlreadyProcessed.selector);
        bridgePoly.releaseTokens(user, lockAmount, 99, signature2);
    }
}
