import { useState, useEffect } from 'react';
import { ethers } from 'ethers';
import { CONTRACT_ADDRESSES, TOKEN_ABI, BRIDGE_ABI } from '../config/contracts';
import { RPC_URLS } from '../config/networks';

export function useBridge(account, provider, signer) {
  const [balances, setBalances] = useState({
    ethereum: '0',
    polygon: '0'
  });
  const [isLoading, setIsLoading] = useState(false);

  // Получаем балансы токенов
  const fetchBalances = async () => {
    if (!account) return;

    try {
      // Ethereum баланс
      const ethProvider = new ethers.JsonRpcProvider(RPC_URLS.ethereum);
      const ethTokenContract = new ethers.Contract(
        CONTRACT_ADDRESSES.ethereum.token,
        TOKEN_ABI,
        ethProvider
      );
      const ethBalance = await ethTokenContract.balanceOf(account);

      // Polygon баланс  
      const polyProvider = new ethers.JsonRpcProvider(RPC_URLS.polygon);
      const polyTokenContract = new ethers.Contract(
        CONTRACT_ADDRESSES.polygon.token,
        TOKEN_ABI, 
        polyProvider
      );
      const polyBalance = await polyTokenContract.balanceOf(account);

      setBalances({
        ethereum: ethers.formatUnits(ethBalance, 18),
        polygon: ethers.formatUnits(polyBalance, 18)
      });
    } catch (error) {
      console.error('Error fetching balances:', error);
    }
  };

  // Перевод токенов
  const transferTokens = async (amount) => {
    if (!account || !signer) {
      alert('Please connect your wallet first');
      return;
    }

    setIsLoading(true);
    try {
      const amountWei = ethers.parseUnits(amount, 18);
      
      // 1. Approve токенов для моста
      const tokenContract = new ethers.Contract(
        CONTRACT_ADDRESSES.ethereum.token,
        TOKEN_ABI,
        signer
      );
      
      console.log('Approving tokens...');
      const approveTx = await tokenContract.approve(
        CONTRACT_ADDRESSES.ethereum.bridge,
        amountWei
      );
      await approveTx.wait();
      console.log('Tokens approved');

      // 2. Lock tokens в мосте
      const bridgeContract = new ethers.Contract(
        CONTRACT_ADDRESSES.ethereum.bridge,
        BRIDGE_ABI,
        signer
      );

      console.log('Locking tokens...');
      const lockTx = await bridgeContract.lockTokens(
        amountWei,
        account, // targetAddress (тот же аккаунт в Polygon)
        80002    // Polygon Amoy chainId
      );
      
      const receipt = await lockTx.wait();
      console.log('Tokens locked:', receipt);

      // Обновляем балансы
      await fetchBalances();
      
      return receipt;
    } catch (error) {
      console.error('Transfer error:', error);
      alert(`Transfer failed: ${error.message}`);
    } finally {
      setIsLoading(false);
    }
  };

  // Автоматически обновляем балансы при изменении аккаунта
  useEffect(() => {
    fetchBalances();
  }, [account]);

  return {
    balances,
    isLoading,
    transferTokens,
    refreshBalances: fetchBalances
  };
}
