// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-contracts/contracts/security/ReentrancyGuard.sol";
import "openzeppelin-contracts/contracts/access/Ownable.sol";

contract BridgeEthereum is ReentrancyGuard, Ownable {
    IERC20 public immutable token;
    uint256 public nonce;
    bool public paused;
    
    event TokensLocked(
        address indexed user,
        uint256 amount,
        uint256 indexed nonce,
        address targetAddress,
        uint256 targetChainId
    );
    
    event BridgePaused(address indexed admin);
    event BridgeUnpaused(address indexed admin);
    
    error BridgeIsPaused();
    error ZeroAmount();
    error ZeroAddress();
    error ZeroChainId();
    error InsufficientBalance();
    error InsufficientAllowance();
    
    constructor(address _token) {
        if (_token == address(0)) revert ZeroAddress();
        token = IERC20(_token);
        paused = false;
        _transferOwnership(msg.sender); // Правильный способ установить владельца
    }
    
    function lockTokens(
        uint256 amount,
        address targetAddress,
        uint256 targetChainId
    ) external nonReentrant {
        if (paused) revert BridgeIsPaused();
        if (amount == 0) revert ZeroAmount();
        if (targetAddress == address(0)) revert ZeroAddress();
        if (targetChainId == 0) revert ZeroChainId();
        
        if (token.balanceOf(msg.sender) < amount) revert InsufficientBalance();
        if (token.allowance(msg.sender, address(this)) < amount) revert InsufficientAllowance();
        
        bool success = token.transferFrom(msg.sender, address(this), amount);
        require(success, "Token transfer failed");
        
        nonce++;
        emit TokensLocked(
            msg.sender,
            amount,
            nonce,
            targetAddress,
            targetChainId
        );
    }
    
    function pauseBridge() external onlyOwner {
        paused = true;
        emit BridgePaused(msg.sender);
    }
    
    function unpauseBridge() external onlyOwner {
        paused = false;
        emit BridgeUnpaused(msg.sender);
    }
    
    function getNonce() external view returns (uint256) { return nonce; }
    function isPaused() external view returns (bool) { return paused; }
}
