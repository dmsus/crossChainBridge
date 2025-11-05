import React, { useState, useEffect } from 'react';

export function TransactionHistory({ account }) {
  const [transactions, setTransactions] = useState([]);

  // В реальном приложении здесь был бы запрос к Go relayer API
  useEffect(() => {
    if (account) {
      // Заглушка - в production здесь реальные данные из БД
      const mockTransactions = [
        {
          id: 1,
          hash: '0x3d573d25b7b94cab421adddd9cd4981c6ccc1ba27fc933691b5110095c289f3d',
          type: 'lock',
          amount: '0.1',
          fromChain: 'Ethereum',
          toChain: 'Polygon',
          status: 'completed',
          timestamp: new Date().toISOString()
        }
      ];
      setTransactions(mockTransactions);
    }
  }, [account]);

  const getStatusColor = (status) => {
    switch (status) {
      case 'completed': return 'text-green-600 bg-green-100';
      case 'pending': return 'text-orange-600 bg-orange-100';
      case 'failed': return 'text-red-600 bg-red-100';
      default: return 'text-gray-600 bg-gray-100';
    }
  };

  const openExplorer = (hash) => {
    window.open(`https://sepolia.etherscan.io/tx/${hash}`, '_blank');
  };

  if (!account) {
    return (
      <div>
        <h3 className="text-lg font-semibold text-gray-800 mb-4">
          📜 Transaction History
        </h3>
        <div className="text-center py-8 bg-gray-50 rounded-lg">
          <p className="text-gray-500">Connect wallet to see transaction history</p>
        </div>
      </div>
    );
  }

  return (
    <div>
      <div className="flex justify-between items-center mb-4">
        <h3 className="text-lg font-semibold text-gray-800">
          📜 Transaction History
        </h3>
        <span className="text-sm text-gray-500">
          {transactions.length} transactions
        </span>
      </div>

      {transactions.length === 0 ? (
        <div className="text-center py-8 bg-gray-50 rounded-lg">
          <p className="text-gray-500">No transactions yet</p>
          <p className="text-sm text-gray-400 mt-1">
            Your bridge transactions will appear here
          </p>
        </div>
      ) : (
        <div className="space-y-3 max-h-64 overflow-y-auto">
          {transactions.map((tx) => (
            <div
              key={tx.id}
              className="border border-gray-200 rounded-lg p-3 hover:bg-gray-50 cursor-pointer"
              onClick={() => openExplorer(tx.hash)}
            >
              <div className="flex justify-between items-start mb-2">
                <div>
                  <span className="font-medium text-gray-800">
                    {tx.amount} TEST
                  </span>
                  <div className="text-sm text-gray-600">
                    {tx.fromChain} → {tx.toChain}
                  </div>
                </div>
                <span className={`text-xs px-2 py-1 rounded-full ${getStatusColor(tx.status)}`}>
                  {tx.status}
                </span>
              </div>
              <div className="text-xs text-gray-500 flex justify-between">
                <span>Lock Tokens</span>
                <span>{new Date(tx.timestamp).toLocaleTimeString()}</span>
              </div>
            </div>
          ))}
        </div>
      )}

      <div className="text-xs text-gray-400 mt-3">
        💡 Click on transaction to view on Etherscan
      </div>
    </div>
  );
}
