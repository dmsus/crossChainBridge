import React from 'react';
import { useBridge } from '../hooks/useBridge';

export function TokenBalance({ account, provider, signer }) {
  const { balances, isLoading, refreshBalances } = useBridge(account, provider, signer);

  return (
    <div className="bg-white rounded-xl shadow-lg p-6">
      <div className="flex justify-between items-center mb-4">
        <h3 className="text-lg font-semibold text-gray-800">
          💰 Token Balances
        </h3>
        <button 
          onClick={refreshBalances}
          disabled={isLoading}
          className="text-sm bg-gray-100 hover:bg-gray-200 px-2 py-1 rounded text-gray-600"
        >
          🔄
        </button>
      </div>
      
      <div className="space-y-3">
        <div className="flex justify-between items-center">
          <span className="text-gray-600">Ethereum Sepolia:</span>
          <div className="text-right">
            <span className="font-medium text-gray-800">
              {isLoading ? '...' : parseFloat(balances.ethereum).toFixed(4)}
            </span>
            <div className="text-xs text-gray-400">TEST</div>
          </div>
        </div>
        
        <div className="flex justify-between items-center">
          <span className="text-gray-600">Polygon Amoy:</span>
          <div className="text-right">
            <span className="font-medium text-gray-800">
              {isLoading ? '...' : parseFloat(balances.polygon).toFixed(4)}
            </span>
            <div className="text-xs text-gray-400">TEST</div>
          </div>
        </div>
      </div>
      
      {!account && (
        <p className="text-xs text-gray-400 mt-3">
          Connect wallet to see balances
        </p>
      )}
    </div>
  );
}
