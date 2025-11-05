import { useState, useEffect } from 'react';

export const TransactionStatus = ({ txHash, account }) => {
  const [status, setStatus] = useState('pending');
  const [polygonTx, setPolygonTx] = useState('');

  useEffect(() => {
    if (!txHash) return;

    // Здесь будет интеграция с Go relayer API
    const checkStatus = async () => {
      try {
        // Временная заглушка - в реальности запрос к relayer API
        setTimeout(() => {
          setStatus('processed');
          setPolygonTx('0x...polygon-tx-hash'); // Заглушка
        }, 30000); // Через 30 секунд
      } catch (error) {
        console.error('Status check error:', error);
      }
    };

    checkStatus();
    const interval = setInterval(checkStatus, 10000);
    return () => clearInterval(interval);
  }, [txHash]);

  const getStatusMessage = () => {
    switch(status) {
      case 'pending':
        return '⏳ Waiting for relayer to process...';
      case 'processed':
        return '✅ Tokens released on Polygon!';
      default:
        return '🔄 Processing...';
    }
  };

  if (!txHash) return null;

  return (
    <div className="p-4 bg-purple-500 rounded-lg text-white mt-4">
      <h3 className="font-bold mb-2">🔄 Transaction Status</h3>
      <p>{getStatusMessage()}</p>
      <div className="text-sm mt-2 space-y-1">
        <p>Ethereum Tx: <a href={`https://sepolia.etherscan.io/tx/${txHash}`} target="_blank" className="underline">{txHash.slice(0,10)}...{txHash.slice(-8)}</a></p>
        {polygonTx && (
          <p>Polygon Tx: <a href={`https://amoy.polygonscan.com/tx/${polygonTx}`} target="_blank" className="underline">View on Explorer</a></p>
        )}
      </div>
    </div>
  );
};
