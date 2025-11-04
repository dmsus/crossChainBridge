// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-contracts/contracts/security/ReentrancyGuard.sol";
import "openzeppelin-contracts/contracts/access/Ownable.sol";

contract BridgeEthereumOptimized is ReentrancyGuard, Ownable {
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
        // paused = false; // default value, saves gas
        _transferOwnership(msg.sender);
    }
    
    /**
     * @dev Lock tokens for cross-chain transfer - GAS OPTIMIZED
     * Gas optimizations:
     * - Check cheap operations first (paused, zero checks)
     * - Remove redundant balance/allowance variables
     * - Direct transferFrom without success check
     * - Pre-increment nonce for gas efficiency
     */
    function lockTokens(
        uint256 amount,
        address targetAddress,
        uint256 targetChainId
    ) external nonReentrant {
        // Gas optimization: check cheap operations first
        if (paused) revert BridgeIsPaused();
        if (amount == 0) revert ZeroAmount();
        if (targetAddress == address(0)) revert ZeroAddress();
        if (targetChainId == 0) revert ZeroChainId();
        
        // Gas optimization: inline balance and allowance checks
        if (token.balanceOf(msg.sender) < amount) revert InsufficientBalance();
        if (token.allowance(msg.sender, address(this)) < amount) revert InsufficientAllowance();
        
        // Gas optimization: direct transferFrom (reverts on failure)
        token.transferFrom(msg.sender, address(this), amount);
        
        // Gas optimization: pre-increment nonce
        emit TokensLocked(
            msg.sender,
            amount,
            ++nonce, // Pre-increment saves gas
            targetAddress,
            targetChainId
        );
    }
    
    /**
     * @dev Pause bridge operations (only owner) - GAS OPTIMIZED
     */
    function pauseBridge() external onlyOwner {
        if (!paused) { // Gas optimization: skip if already paused
            paused = true;
            emit BridgePaused(msg.sender);
        }
    }
    
    /**
     * @dev Unpause bridge operations (only owner) - GAS OPTIMIZED
     */
    function unpauseBridge() external onlyOwner {
        if (paused) { // Gas optimization: skip if already unpaused
            paused = false;
            emit BridgeUnpaused(msg.sender);
        }
    }
    
    function getNonce() external view returns (uint256) { return nonce; }
    function isPaused() external view returns (bool) { return paused; }
}
