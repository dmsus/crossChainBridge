// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-contracts/contracts/access/Ownable.sol";
import "openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";
import "openzeppelin-contracts/contracts/utils/cryptography/EIP712.sol";

/**
 * @title BridgePolygon
 * @dev Cross-chain bridge contract for releasing tokens on Polygon Amoy
 * @dev Uses EIP-712 for secure signature verification
 */
contract BridgePolygon is Ownable, EIP712 {
    using ECDSA for bytes32;
    
    IERC20 public immutable token;
    address public immutable signer;
    mapping(uint256 => bool) public processedNonces;
    bool public paused;
    
    // EIP-712 typehash - исправленная версия
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
        paused = false;
        _transferOwnership(msg.sender);
    }
    
    function releaseTokens(
        address user,
        uint256 amount,
        uint256 nonce,
        bytes calldata signature
    ) external {
        if (paused) revert BridgeIsPaused();
        if (user == address(0)) revert ZeroAddress();
        if (amount == 0) revert ZeroAmount();
        if (processedNonces[nonce]) revert NonceAlreadyProcessed();
        
        // Создаем хэш структуры данных для подписи
        bytes32 structHash = keccak256(
            abi.encode(
                _RELEASE_TYPEHASH,
                user,
                amount,
                nonce
            )
        );
        
        bytes32 digest = _hashTypedDataV4(structHash);
        address recoveredSigner = ECDSA.recover(digest, signature);
        
        if (recoveredSigner != signer) revert InvalidSignature();
        
        processedNonces[nonce] = true;
        
        bool success = token.transfer(user, amount);
        require(success, "Token transfer failed");
        
        emit TokensReleased(user, amount, nonce);
    }
    
    function pauseBridge() external onlyOwner {
        paused = true;
        emit BridgePaused(msg.sender);
    }
    
    function unpauseBridge() external onlyOwner {
        paused = false;
        emit BridgeUnpaused(msg.sender);
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
    
    // Helper function for tests - создает digest для подписи
    function getDigest(address user, uint256 amount, uint256 nonce) external view returns (bytes32) {
        bytes32 structHash = keccak256(
            abi.encode(
                _RELEASE_TYPEHASH,
                user,
                amount,
                nonce
            )
        );
        return _hashTypedDataV4(structHash);
    }
}
