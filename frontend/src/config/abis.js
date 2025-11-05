// Standard ERC-20 ABI
export const ERC20_ABI = [
  "function name() view returns (string)",
  "function symbol() view returns (string)",
  "function decimals() view returns (uint8)",
  "function totalSupply() view returns (uint256)",
  "function balanceOf(address) view returns (uint256)",
  "function transfer(address to, uint256 amount) returns (bool)",
  "function allowance(address owner, address spender) view returns (uint256)",
  "function approve(address spender, uint256 amount) returns (bool)",
  "function transferFrom(address from, address to, uint256 amount) returns (bool)",
  "event Transfer(address indexed from, address indexed to, uint256 amount)",
  "event Approval(address indexed owner, address indexed spender, uint256 amount)"
];

// Bridge Contract ABI (основные функции)
export const BRIDGE_ABI = [
  "function lockTokens(uint256 amount) external",
  "function lockTokens(uint256 amount, uint256 targetChain) external", 
  "function unlockTokens(address to, uint256 amount, bytes32 originalTxHash) external",
  "function getTargetChain() view returns (uint256)",
  "function token() view returns (address)",
  "function owner() view returns (address)",
  "event TokensLocked(address indexed from, uint256 amount, uint256 targetChain)",
  "event TokensUnlocked(address indexed to, uint256 amount, bytes32 originalTxHash)"
];
