import { useState, useEffect } from 'react';
import { ethers } from 'ethers';
import { NETWORKS } from '../config/networks';

export function useWeb3() {
  const [account, setAccount] = useState('');
  const [provider, setProvider] = useState(null);
  const [signer, setSigner] = useState(null);
  const [isConnected, setIsConnected] = useState(false);
  const [currentChain, setCurrentChain] = useState(null);

  useEffect(() => {
    checkConnectedWallet();
    
    if (window.ethereum) {
      window.ethereum.on('accountsChanged', handleAccountsChanged);
      window.ethereum.on('chainChanged', handleChainChanged);
    }

    return () => {
      if (window.ethereum) {
        window.ethereum.removeListener('accountsChanged', handleAccountsChanged);
        window.ethereum.removeListener('chainChanged', handleChainChanged);
      }
    };
  }, []);

  const checkConnectedWallet = async () => {
    if (window.ethereum) {
      try {
        const accounts = await window.ethereum.request({ 
          method: 'eth_accounts' 
        });
        if (accounts.length > 0) {
          setAccount(accounts[0]);
          setIsConnected(true);
          await initProvider();
        }
      } catch (error) {
        console.error('Error checking connected wallet:', error);
      }
    }
  };

  const initProvider = async () => {
    if (window.ethereum) {
      const web3Provider = new ethers.BrowserProvider(window.ethereum);
      setProvider(web3Provider);
      const web3Signer = await web3Provider.getSigner();
      setSigner(web3Signer);
      
      // Получаем текущую сеть
      const network = await web3Provider.getNetwork();
      setCurrentChain(network.chainId.toString());
    }
  };

  const handleAccountsChanged = (accounts) => {
    if (accounts.length === 0) {
      setAccount('');
      setIsConnected(false);
      setProvider(null);
      setSigner(null);
      setCurrentChain(null);
    } else {
      setAccount(accounts[0]);
      initProvider();
    }
  };

  const handleChainChanged = (chainId) => {
    setCurrentChain(parseInt(chainId, 16).toString());
  };

  const connectWallet = async () => {
    if (!window.ethereum) {
      alert('Please install MetaMask!');
      return;
    }

    try {
      const accounts = await window.ethereum.request({ 
        method: 'eth_requestAccounts' 
      });
      setAccount(accounts[0]);
      setIsConnected(true);
      await initProvider();
      return accounts[0];
    } catch (error) {
      console.error('Error connecting wallet:', error);
      if (error.code === 4001) {
        alert('Please connect your wallet to use the bridge');
      }
    }
  };

  const switchNetwork = async (networkKey) => {
    if (!window.ethereum) return false;

    try {
      await window.ethereum.request({
        method: 'wallet_switchEthereumChain',
        params: [{ chainId: NETWORKS[networkKey].chainId }],
      });
      return true;
    } catch (error) {
      if (error.code === 4902) {
        // Сеть не добавлена, добавляем
        try {
          await window.ethereum.request({
            method: 'wallet_addEthereumChain',
            params: [NETWORKS[networkKey]],
          });
          return true;
        } catch (addError) {
          console.error('Error adding network:', addError);
          return false;
        }
      }
      console.error('Error switching network:', error);
      return false;
    }
  };

  const disconnectWallet = () => {
    setAccount('');
    setIsConnected(false);
    setProvider(null);
    setSigner(null);
    setCurrentChain(null);
  };

  return { 
    account, 
    provider, 
    signer, 
    isConnected,
    currentChain,
    connectWallet, 
    disconnectWallet,
    switchNetwork
  };
}
