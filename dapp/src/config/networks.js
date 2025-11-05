export const NETWORKS = {
  ethereum: {
    chainId: "0xaa36a7", // 11155111
    chainName: "Ethereum Sepolia", 
    rpcUrls: ["https://sepolia.infura.io/v3/afa8f63b32a84f508f554b798b247453"],
    blockExplorerUrls: ["https://sepolia.etherscan.io"],
    nativeCurrency: {
      name: "Sepolia ETH",
      symbol: "ETH",
      decimals: 18
    }
  },
  polygon: {
    chainId: "0x13882", // 80002
    chainName: "Polygon Amoy",
    rpcUrls: ["https://polygon-amoy.infura.io/v3/afa8f63b32a84f508f554b798b247453"], 
    blockExplorerUrls: ["https://amoy.polygonscan.com"],
    nativeCurrency: {
      name: "Polygon MATIC",
      symbol: "MATIC", 
      decimals: 18
    }
  }
};

// RPC URLs для ethers.js
export const RPC_URLS = {
  ethereum: "https://sepolia.infura.io/v3/afa8f63b32a84f508f554b798b247453",
  polygon: "https://polygon-amoy.infura.io/v3/afa8f63b32a84f508f554b798b247453"
};
