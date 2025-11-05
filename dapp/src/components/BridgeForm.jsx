import React, { useState } from 'react';
import { useBridge } from '../hooks/useBridge';

export function BridgeForm({ account, provider, signer }) {
  const [amount, setAmount] = useState('');
  const { balances, isLoading, transferTokens } = useBridge(account, provider, signer);

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!amount || parseFloat(amount) <= 0) {
      alert('Please enter a valid amount');
      return;
    }

    if (parseFloat(amount) > parseFloat(balances.ethereum)) {
      alert('Insufficient balance');
      return;
    }

    try {
      const result = await transferTokens(amount);
      if (result) {
        setAmount('');
        alert('Transfer initiated! Check transaction history.');
      }
    } catch (error) {
      console.error('Transfer error:', error);
    }
  };

  const setMaxAmount = () => {
    if (balances.ethereum && parseFloat(balances.ethereum) > 0) {
      setAmount(parseFloat(balances.ethereum).toFixed(4));
    }
  };

  if (!account) {
    return (
      <div>
        <h3 className="text-lg font-semibold text-gray-800 mb-4">
          🌉 Bridge Tokens
        </h3>
        <div className="text-center py-8 bg-gray-50 rounded-lg">
          <p className="text-gray-500">Please connect your wallet to use the bridge</p>
        </div>
      </div>
    );
  }

  return (
    <div>
      <h3 className="text-lg font-semibold text-gray-800 mb-4">
        🌉 Bridge Tokens to Polygon
      </h3>
      
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Amount to Transfer
          </label>
          <div className="relative">
            <input
              type="number"
              step="0.0001"
              min="0"
              value={amount}
              onChange={(e) => setAmount(e.target.value)}
              placeholder="0.0"
              className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              disabled={isLoading}
            />
            <button
              type="button"
              onClick={setMaxAmount}
              className="absolute right-2 top-1/2 transform -translate-y-1/2 bg-blue-100 text-blue-600 px-2 py-1 rounded text-sm hover:bg-blue-200"
            >
              MAX
            </button>
          </div>
          <div className="flex justify-between text-sm text-gray-500 mt-1">
            <span>Available: {parseFloat(balances.ethereum).toFixed(4)} TEST</span>
            <span>→ Polygon Amoy</span>
          </div>
        </div>

        <button
          type="submit"
          disabled={isLoading || !amount || parseFloat(amount) <= 0}
          className="w-full bg-blue-500 hover:bg-blue-600 disabled:bg-gray-400 text-white font-medium py-3 px-4 rounded-lg transition duration-200 flex items-center justify-center"
        >
          {isLoading ? (
            <>
              <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
              Processing...
            </>
          ) : (
            'Transfer to Polygon'
          )}
        </button>

        <div className="text-xs text-gray-500 space-y-1">
          <p>🔒 Tokens will be locked in Ethereum and released in Polygon</p>
          <p>⏱️ Transfer takes 2-5 minutes to complete</p>
        </div>
      </form>
    </div>
  );
}
