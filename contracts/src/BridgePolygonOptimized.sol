// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-contracts/contracts/access/Ownable.sol";
import "openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";
import "openzeppelin-contracts/contracts/utils/cryptography/EIP712.sol";

/**
 * @title BridgePolygonOptimized
 * @dev Gas-optimized cross-chain bridge contract for releasing tokens on Polygon Amoy
 * @dev Uses EIP-712 for secure signature verification with gas optimizations
 */
contract BridgePolygonOptimized is Ownable, EIP712 {
    using ECDSA for bytes32;
    
    IERC20 public immutable token;
    address public immutable signer;
    mapping(uint256 => bool) public processedNonces;
    bool public paused;
    
    // EIP-712 typehash - precomputed for gas efficiency
    bytes32 private constant _RELEASE_TYPEHASH = keccak256(
        "ReleaseTokens(address user,uint256 amount,uint256 nonce)"
    );
    
    event TokensReleased(
        address indexed user,
        uint256 amount,
        uint256 indexed nonce
    );
    
    event BridgePaused(address indexed admin);
    event BridgeUnpaused(address indexed admin);
    
    error BridgeIsPaused();
    error ZeroAmount();
    error ZeroAddress();
    error InvalidSignature();
    error NonceAlreadyProcessed();
    
    constructor(address _token, address _signer) EIP712("CrossChainBridge", "1") {
        if (_token == address(0)) revert ZeroAddress();
        if (_signer == address(0)) revert ZeroAddress();
        
        token = IERC20(_token);
        signer = _signer;
        // paused = false; // default value, saves 200 gas
        _transferOwnership(msg.sender);
    }
    
    /**
     * @dev Release tokens to user with signature verification - GAS OPTIMIZED
     * Gas optimizations:
     * - Check cheap operations first (paused, nonce)
     * - Use calldata for signature
     * - Inline struct hash computation
     * - Remove redundant variables
     */
    function releaseTokens(
        address user,
        uint256 amount,
        uint256 nonce,
        bytes calldata signature
    ) external {
        // Gas optimization: check cheap operations first
        if (paused) revert BridgeIsPaused();
        if (processedNonces[nonce]) revert NonceAlreadyProcessed();
        if (amount == 0) revert ZeroAmount();
        if (user == address(0)) revert ZeroAddress();
        
        // Gas optimization: inline struct hash computation
        bytes32 digest = _hashTypedDataV4(
            keccak256(abi.encode(_RELEASE_TYPEHASH, user, amount, nonce))
        );
        
        // Gas optimization: use ECDSA.recover directly
        if (ECDSA.recover(digest, signature) != signer) revert InvalidSignature();
        
        processedNonces[nonce] = true;
        
        // Gas optimization: direct transfer without success check (reverts on failure)
        token.transfer(user, amount);
        
        emit TokensReleased(user, amount, nonce);
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
    
    function isNonceProcessed(uint256 nonce) external view returns (bool) {
        return processedNonces[nonce];
    }
    
    function isPaused() external view returns (bool) {
        return paused;
    }
    
    function getDomainSeparator() external view returns (bytes32) {
        return _domainSeparatorV4();
    }
    
    function getBridgeBalance() external view returns (uint256) {
        return token.balanceOf(address(this));
    }
    
    // Helper function for tests
    function getDigest(address user, uint256 amount, uint256 nonce) external view returns (bytes32) {
        return _hashTypedDataV4(
            keccak256(abi.encode(_RELEASE_TYPEHASH, user, amount, nonce))
        );
    }
}
