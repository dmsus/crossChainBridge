// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import "../src/BridgeEthereum.sol";
import "../src/TokenEthereum.sol";

contract DeployEthereum is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
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
