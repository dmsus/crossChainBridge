// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import "../src/BridgePolygonOptimized.sol";
import "../src/TokenPolygon.sol";

contract DeployPolygonOptimized is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.rememberKey(deployerPrivateKey);
        
        vm.startBroadcast(deployer);

        // Деплоим оптимизированный токен
        TokenPolygon token = new TokenPolygon();
        console.log("TokenPolygonOptimized deployed at:", address(token));

        // Деплоим оптимизированный мост с подписантом
        address signer = vm.envAddress("SIGNER_ADDRESS");
        BridgePolygonOptimized bridge = new BridgePolygonOptimized(address(token), signer);
        console.log("BridgePolygonOptimized deployed at:", address(bridge));

        vm.stopBroadcast();
    }
}
