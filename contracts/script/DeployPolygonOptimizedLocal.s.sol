// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import "../src/BridgePolygonOptimized.sol";
import "../src/TokenPolygon.sol";

contract DeployPolygonOptimizedLocal is Script {
    function run() external {
        // Используем стандартный тестовый приватный ключ из Anvil
        uint256 deployerPrivateKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
        address deployer = vm.rememberKey(deployerPrivateKey);
        
        vm.startBroadcast(deployer);

        // Деплоим оптимизированный токен
        TokenPolygon token = new TokenPolygon();
        console.log("TokenPolygonOptimized deployed at:", address(token));

        // Деплоим оптимизированный мост с подписантом
        address signer = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
        BridgePolygonOptimized bridge = new BridgePolygonOptimized(address(token), signer);
        console.log("BridgePolygonOptimized deployed at:", address(bridge));

        vm.stopBroadcast();
    }
}
