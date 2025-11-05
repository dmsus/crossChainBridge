import { useState, useEffect } from 'react';
import { ethers } from 'ethers';
import { CONFIG } from '../config/networks';

export const WalletConnector = ({ onAccountChange }) => {
  const [account, setAccount] = useState('');
  const [chainId, setChainId] = useState('');
  const [networkName, setNetworkName] = useState('');

  const addAmoyNetwork = async () => {
    try {
      await window.ethereum.request({
        method: 'wallet_addEthereumChain',
        params: [{
          chainId: CONFIG.POLYGON.chainId,
          chainName: 'Polygon Amoy Testnet',
          rpcUrls: [CONFIG.POLYGON.rpcUrl],
          blockExplorerUrls: ['https://amoy.polygonscan.com/'],
          nativeCurrency: {
            name: 'MATIC',
            symbol: 'MATIC',
            decimals: 18
          }
        }]
      });
    } catch (error) {
      console.error('Error adding Amoy network:', error);
    }
  };

  const switchToSepolia = async () => {
    try {
      await window.ethereum.request({
        method: 'wallet_switchEthereumChain',
        params: [{ chainId: CONFIG.ETHEREUM.chainId }],
      });
    } catch (error) {
      console.error('Error switching to Sepolia:', error);
    }
  };

  const switchToAmoy = async () => {
    try {
      await window.ethereum.request({
        method: 'wallet_switchEthereumChain', 
        params: [{ chainId: CONFIG.POLYGON.chainId }],
      });
    } catch (switchError) {
      // Если сеть не добавлена, добавляем её
      if (switchError.code === 4902) {
        await addAmoyNetwork();
      }
    }
  };

  const switchToMainnet = async () => {
    try {
      await window.ethereum.request({
        method: 'wallet_switchEthereumChain',
        params: [{ chainId: '0x1' }],
      });
    } catch (error) {
      console.error('Error switching to Mainnet:', error);
    }
  };

  const connectWallet = async () => {
    if (window.ethereum) {
      try {
        const accounts = await window.ethereum.request({
          method: 'eth_requestAccounts'
        });
        const account = accounts[0];
        setAccount(account);
        onAccountChange(account);
        
        const provider = new ethers.BrowserProvider(window.ethereum);
        const network = await provider.getNetwork();
        updateNetworkInfo(network.chainId.toString());

        // Слушаем изменения аккаунта и сети
        window.ethereum.on('accountsChanged', (accounts) => {
          const newAccount = accounts[0] || '';
          setAccount(newAccount);
          onAccountChange(newAccount);
        });

        window.ethereum.on('chainChanged', (chainId) => {
          updateNetworkInfo(parseInt(chainId).toString());
        });

      } catch (error) {
        console.error('Error connecting wallet:', error);
      }
    } else {
      alert('Please install MetaMask!');
    }
  };

  const updateNetworkInfo = (chainId) => {
    setChainId(chainId);
    setNetworkName(getNetworkName(chainId));
  };

  const getNetworkName = (chainId) => {
    switch(chainId) {
      case '1': return 'Ethereum Mainnet';
      case '11155111': return 'Sepolia Testnet';
      case '80002': return 'Polygon Amoy';
      default: return `Unknown Network (${chainId})`;
    }
  };

  return (
    <div className="p-6 bg-gradient-to-r from-blue-500 to-purple-600 rounded-lg shadow-lg text-white">
      <h2 className="text-2xl font-bold mb-4">🔗 Wallet Connection</h2>
      {!account ? (
        <button
          onClick={connectWallet}
          className="bg-white text-blue-600 px-6 py-3 rounded-lg font-bold hover:bg-gray-100 transition-colors"
        >
          🔌 Connect MetaMask
        </button>
      ) : (
        <div className="space-y-4">
          <div className="bg-white/20 p-4 rounded-lg">
            <p className="font-semibold">✅ Connected: {account.slice(0, 6)}...{account.slice(-4)}</p>
            <p className="text-sm opacity-90">Network: {networkName} (ID: {chainId})</p>
          </div>
          
          <div className="flex flex-wrap gap-2">
            <button
              onClick={switchToMainnet}
              className="bg-gray-600 text-white px-3 py-2 rounded text-sm hover:bg-gray-700 transition-colors"
            >
              🏠 Mainnet
            </button>
            <button
              onClick={switchToSepolia}
              className="bg-purple-600 text-white px-3 py-2 rounded text-sm hover:bg-purple-700 transition-colors"
            >
              🔵 Sepolia
            </button>
            <button
              onClick={switchToAmoy} 
              className="bg-indigo-600 text-white px-3 py-2 rounded text-sm hover:bg-indigo-700 transition-colors"
            >
              🟣 Polygon Amoy
            </button>
          </div>
        </div>
      )}
    </div>
  );
};
