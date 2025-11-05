import { useState, useEffect } from 'react';
import { ethers } from 'ethers';
import { CONFIG } from '../config/networks';

export const ContractDebugger = ({ account }) => {
  const [debugInfo, setDebugInfo] = useState('');

  useEffect(() => {
    const checkContracts = async () => {
      if (!account) return;
      
      let info = '🔍 CONTRACT DEBUG INFO:\\n\\n';
      
      try {
        const provider = new ethers.JsonRpcProvider(CONFIG.ETHEREUM.rpcUrl);
        
        // 1. Проверяем Token контракт
        info += 'TOKEN CONTRACT:\\n';
        const token = new ethers.Contract(
          CONFIG.ETHEREUM.tokenAddress,
          ['function balanceOf(address) view returns (uint256)', 'function decimals() view returns (uint8)'],
          provider
        );
        
        const tokenCode = await provider.getCode(CONFIG.ETHEREUM.tokenAddress);
        const balance = await token.balanceOf(account);
        const decimals = await token.decimals();
        const allowance = await token.allowance(account, CONFIG.ETHEREUM.bridgeAddress);
        
        info += `✅ Code deployed: ${tokenCode.length > 2}\\n`;
        info += `✅ Balance: ${ethers.formatUnits(balance, decimals)} TEST\\n`;
        info += `✅ Decimals: ${decimals}\\n`;
        info += `✅ Allowance for bridge: ${ethers.formatUnits(allowance, decimals)} TEST\\n\\n`;
        
        // 2. Проверяем Bridge контракт
        info += 'BRIDGE CONTRACT:\\n';
        const bridgeCode = await provider.getCode(CONFIG.ETHEREUM.bridgeAddress);
        info += `✅ Code deployed: ${bridgeCode.length > 2}\\n`;
        
        if (bridgeCode.length > 2) {
          const bridge = new ethers.Contract(
            CONFIG.ETHEREUM.bridgeAddress,
            ['function getTargetChain() view returns (uint256)'],
            provider
          );
          
          try {
            const targetChain = await bridge.getTargetChain();
            info += `✅ Target chain: ${targetChain}\\n`;
          } catch (e) {
            info += `⚠️ Cannot read targetChain: ${e.message}\\n`;
          }
        }
        
        info += `\\n🔧 SUGGESTIONS:\\n`;
        info += `• Check if bridge accepts this token\\n`;
        info += `• Try different lockTokens parameters\\n`;
        info += `• Verify contract ABI matches deployed version\\n`;
        
      } catch (error) {
        info += `❌ Debug error: ${error.message}\\n`;
      }
      
      setDebugInfo(info);
    };

    checkContracts();
  }, [account]);

  return (
    <div className="p-4 bg-yellow-100 border border-yellow-400 rounded-lg mt-4">
      <h3 className="font-bold text-yellow-800 mb-2">Contract Debugger</h3>
      <pre className="text-sm whitespace-pre-wrap text-yellow-700">{debugInfo}</pre>
    </div>
  );
};
