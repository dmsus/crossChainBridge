// ABI для контрактов (упрощенная версия)
export const BRIDGE_ABI = [
  "function lockTokens(uint256 amount, address targetAddress, uint256 targetChainId) external",
  "function releaseTokens(address user, uint256 amount, uint256 nonce, bytes memory signature) external", 
  "event TokensLocked(address indexed user, uint256 amount, uint256 indexed nonce, address targetAddress, uint256 targetChainId)"
];

export const TOKEN_ABI = [
  "function balanceOf(address account) external view returns (uint256)",
  "function approve(address spender, uint256 amount) external returns (bool)",
  "function transfer(address to, uint256 amount) external returns (bool)"
];

// Адреса контрактов из твоего деплоя
export const CONTRACT_ADDRESSES = {
  ethereum: {
    bridge: "0xc2766cBFc1Dc3E95058bf09c929E7D6226b10187",
    token: "0x5CFdE9C777be47FC4a401c918181DD92BA4c81Cc"
  },
  polygon: {
    bridge: "0xc2766cBFc1Dc3E95058bf09c929E7D6226b10187",
    token: "0x5CFdE9C777be47FC4a401c918181DD92BA4c81Cc"
  }
};
