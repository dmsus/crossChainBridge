// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import "../src/BridgePolygon.sol";
import "../src/TokenPolygon.sol";

contract DeployPolygon is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address relayer = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8; // test address
        
        vm.startBroadcast(deployerPrivateKey);

        // Deploy token
        TokenPolygon token = new TokenPolygon();
        console.log("TokenPolygon deployed at:", address(token));

        // Deploy bridge
        BridgePolygon bridge = new BridgePolygon(address(token), relayer);
        console.log("BridgePolygon deployed at:", address(bridge));

        vm.stopBroadcast();
    }
}
