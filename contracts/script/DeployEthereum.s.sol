// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import "../src/BridgeEthereum.sol";
import "../src/TokenEthereum.sol";

contract DeployEthereum is Script {
    function run() external {
        // ФИКС: Используем envString вместо envUint
        string memory privateKeyStr = vm.envString("PRIVATE_KEY");
        uint256 deployerPrivateKey = vm.parseUint(privateKeyStr);
        
        vm.startBroadcast(deployerPrivateKey);

        // Deploy token
        TokenEthereum token = new TokenEthereum();
        console.log("TokenEthereum deployed at:", address(token));

        // Deploy bridge
        BridgeEthereum bridge = new BridgeEthereum(address(token));
        console.log("BridgeEthereum deployed at:", address(bridge));

        vm.stopBroadcast();
    }
}
