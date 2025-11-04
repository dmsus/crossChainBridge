// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import "../src/BridgeEthereumOptimized.sol";
import "../src/TokenEthereum.sol";

contract DeployEthereumOptimizedLocal is Script {
    function run() external {
        // Используем стандартный тестовый приватный ключ из Anvil
        uint256 deployerPrivateKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
        address deployer = vm.rememberKey(deployerPrivateKey);
        
        vm.startBroadcast(deployer);

        // Деплоим оптимизированный токен
        TokenEthereum token = new TokenEthereum();
        console.log("TokenEthereumOptimized deployed at:", address(token));

        // Деплоим оптимизированный мост
        BridgeEthereumOptimized bridge = new BridgeEthereumOptimized(address(token));
        console.log("BridgeEthereumOptimized deployed at:", address(bridge));

        vm.stopBroadcast();
    }
}
