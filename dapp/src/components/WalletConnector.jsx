import React from 'react';
import { NETWORKS } from '../config/networks';

export function WalletConnector({ account, onConnect, onDisconnect, currentChain, onSwitchNetwork }) {
  const formatAddress = (address) => {
    if (!address) return '';
    return `${address.slice(0, 6)}...${address.slice(-4)}`;
  };

  const getNetworkName = (chainId) => {
    if (!chainId) return 'Unknown';
    
    const network = Object.values(NETWORKS).find(
      net => parseInt(net.chainId, 16).toString() === chainId
    );
    return network ? network.chainName : `Chain ${chainId}`;
  };

  const isCorrectNetwork = () => {
    if (!currentChain) return false;
    const sepoliaChainId = parseInt(NETWORKS.ethereum.chainId, 16).toString();
    return currentChain === sepoliaChainId;
  };

  if (!account) {
    return (
      <div className="bg-white rounded-xl shadow-lg p-6">
        <h3 className="text-lg font-semibold text-gray-800 mb-4">
          🔗 Connect Wallet
        </h3>
        <button
          onClick={onConnect}
          className="w-full bg-blue-500 hover:bg-blue-600 text-white font-medium py-3 px-4 rounded-lg transition duration-200"
        >
          Connect MetaMask
        </button>
        <p className="text-sm text-gray-500 mt-3">
          Connect your wallet to use the bridge
        </p>
      </div>
    );
  }

  return (
    <div className="bg-white rounded-xl shadow-lg p-6">
      <h3 className="text-lg font-semibold text-gray-800 mb-4">
        ✅ Wallet Connected
      </h3>
      
      <div className="space-y-3">
        <div>
          <p className="text-sm text-gray-600">Address</p>
          <p className="font-mono text-sm bg-gray-100 p-2 rounded">
            {formatAddress(account)}
          </p>
        </div>

        <div>
          <p className="text-sm text-gray-600">Network</p>
          <div className="flex items-center justify-between">
            <span className={`text-sm font-medium ${
              isCorrectNetwork() ? 'text-green-600' : 'text-orange-600'
            }`}>
              {getNetworkName(currentChain)}
            </span>
            {!isCorrectNetwork() && (
              <button
                onClick={() => onSwitchNetwork('ethereum')}
                className="text-xs bg-orange-500 hover:bg-orange-600 text-white px-2 py-1 rounded"
              >
                Switch to Sepolia
              </button>
            )}
          </div>
        </div>

        <button
          onClick={onDisconnect}
          className="w-full bg-gray-500 hover:bg-gray-600 text-white font-medium py-2 px-4 rounded-lg transition duration-200 text-sm"
        >
          Disconnect
        </button>
      </div>
    </div>
  );
}
