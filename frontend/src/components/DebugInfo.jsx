import { useEffect, useState } from 'react';
import { ethers } from 'ethers';
import { CONFIG } from '../config/networks';

export const DebugInfo = ({ account }) => {
  const [debug, setDebug] = useState('');

  useEffect(() => {
    const checkContracts = async () => {
      if (!account) return;
      
      let debugInfo = 'Debug Information:\\n';
      
      try {
        // Проверяем Sepolia контракт
        const sepoliaProvider = new ethers.JsonRpcProvider(CONFIG.ETHEREUM.rpcUrl);
        const sepoliaToken = new ethers.Contract(
          CONFIG.ETHEREUM.tokenAddress,
          ['function balanceOf(address) view returns (uint256)'],
          sepoliaProvider
        );
        
        const code = await sepoliaProvider.getCode(CONFIG.ETHEREUM.tokenAddress);
        debugInfo += `Sepolia token code length: ${code.length}\\n`;
        
        const balance = await sepoliaToken.balanceOf(account);
        debugInfo += `Sepolia balance: ${ethers.formatUnits(balance, 18)}\\n`;
        
      } catch (error) {
        debugInfo += `Sepolia error: ${error.message}\\n`;
      }

      setDebug(debugInfo);
    };

    checkContracts();
  }, [account]);

  return (
    <div className="p-4 bg-yellow-100 rounded-lg mt-4">
      <h3 className="font-bold text-yellow-800">Debug Info</h3>
      <pre className="text-sm whitespace-pre-wrap">{debug}</pre>
    </div>
  );
};
