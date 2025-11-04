// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import "../src/BridgeEthereumOptimized.sol";
import "../src/TokenEthereum.sol";

contract DeployEthereumOptimized is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
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
