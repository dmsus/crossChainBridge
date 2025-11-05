import { useState, useEffect } from 'react';
import { ethers } from 'ethers';
import { CONFIG } from '../config/networks';

export const TokenBalance = ({ account }) => {
  const [balances, setBalances] = useState({
    sepolia: '0',
    amoy: '0'
  });
  const [loading, setLoading] = useState(false);

  const fetchBalances = async () => {
    if (!account) return;
    
    setLoading(true);
    try {
      // Баланс в Sepolia
      const sepoliaProvider = new ethers.JsonRpcProvider(CONFIG.ETHEREUM.rpcUrl);
      const sepoliaToken = new ethers.Contract(
        CONFIG.ETHEREUM.tokenAddress,
        ['function balanceOf(address) view returns (uint256)'],
        sepoliaProvider
      );
      const sepoliaBalance = await sepoliaToken.balanceOf(account);

      // Баланс в Amoy
      const amoyProvider = new ethers.JsonRpcProvider(CONFIG.POLYGON.rpcUrl);
      const amoyToken = new ethers.Contract(
        CONFIG.POLYGON.tokenAddress,
        ['function balanceOf(address) view returns (uint256)'],
        amoyProvider
      );
      const amoyBalance = await amoyToken.balanceOf(account);

      setBalances({
        sepolia: ethers.formatUnits(sepoliaBalance, 18),
        amoy: ethers.formatUnits(amoyBalance, 18)
      });
    } catch (error) {
      console.error('Error fetching balances:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchBalances();
  }, [account]);

  return (
    <div className="p-6 bg-white rounded-lg shadow-md mt-4">
      <h2 className="text-xl font-bold mb-4">Token Balances</h2>
      
      {loading ? (
        <p className="text-gray-600">Loading balances...</p>
      ) : (
        <div className="grid grid-cols-2 gap-4">
          <div className="text-center p-4 bg-blue-50 rounded">
            <h3 className="font-semibold text-blue-700">Sepolia</h3>
            <p className="text-2xl font-bold mt-2">{parseFloat(balances.sepolia).toFixed(4)}</p>
            <p className="text-sm text-gray-600">TEST Tokens</p>
          </div>
          
          <div className="text-center p-4 bg-purple-50 rounded">
            <h3 className="font-semibold text-purple-700">Polygon Amoy</h3>
            <p className="text-2xl font-bold mt-2">{parseFloat(balances.amoy).toFixed(4)}</p>
            <p className="text-sm text-gray-600">TEST Tokens</p>
          </div>
        </div>
      )}
      
      <button
        onClick={fetchBalances}
        className="mt-4 bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600"
      >
        Refresh Balances
      </button>
    </div>
  );
};
