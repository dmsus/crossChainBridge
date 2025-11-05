import { useState } from 'react';
import { ethers } from 'ethers';
import { CONFIG } from '../config/networks';
import { TransactionStatus } from './TransactionStatus';

export const BridgeForm = ({ account, onBalanceUpdate }) => {
  const [amount, setAmount] = useState('');
  const [loading, setLoading] = useState(false);
  const [status, setStatus] = useState('');
  const [txHash, setTxHash] = useState('');
  const [step, setStep] = useState('');

  const bridgeTokens = async () => {
    if (!amount || !account) return;
    
    setLoading(true);
    setStatus('Starting bridge process...');
    setTxHash('');
    setStep('approve');

    try {
      const provider = new ethers.BrowserProvider(window.ethereum);
      const signer = await provider.getSigner();
      
      // Проверяем что в Sepolia
      const network = await provider.getNetwork();
      if (network.chainId !== 11155111n) {
        setStatus('❌ Please switch to Sepolia network first!');
        setLoading(false);
        return;
      }

      const amountWei = ethers.parseUnits(amount, 18);

      // 1. APPROVE
      setStatus('🔄 Step 1/2: Approving tokens...');
      
      const tokenContract = new ethers.Contract(
        CONFIG.ETHEREUM.tokenAddress,
        CONFIG.ETHEREUM.tokenAbi,
        signer
      );

      const currentAllowance = await tokenContract.allowance(account, CONFIG.ETHEREUM.bridgeAddress);
      
      if (currentAllowance < amountWei) {
        setStatus('🔄 Step 1/2: Sending approve transaction...');
        const approveTx = await tokenContract.approve(CONFIG.ETHEREUM.bridgeAddress, amountWei);
        setTxHash(approveTx.hash);
        
        setStatus('⏳ Step 1/2: Waiting for approve confirmation...');
        await approveTx.wait();
        setStatus('✅ Step 1/2: Tokens approved successfully!');
      } else {
        setStatus('✅ Step 1/2: Allowance already sufficient');
      }

      // 2. LOCK TOKENS
      setStep('lock');
      setStatus('🔄 Step 2/2: Locking tokens in bridge...');
      
      const bridgeContract = new ethers.Contract(
        CONFIG.ETHEREUM.bridgeAddress,
        CONFIG.ETHEREUM.bridgeAbi,
        signer
      );

      setStatus('🔄 Step 2/2: Sending lock transaction...');
      
      const lockTx = await bridgeContract.lockTokens(
        amountWei,           // amount
        account,             // targetAddress
        80002                // targetChainId (Polygon Amoy)
      );
      
      setTxHash(lockTx.hash);
      
      setStatus('⏳ Step 2/2: Waiting for lock confirmation...');
      const receipt = await lockTx.wait();
      
      setStep('complete');
      setStatus('🎉 Tokens locked! Relayer will process to Polygon...');
      setAmount('');
      
      // Обновляем балансы
      if (onBalanceUpdate) {
        setTimeout(onBalanceUpdate, 3000);
      }
      
    } catch (error) {
      console.error('Bridge error:', error);
      setStatus(`❌ Error: ${error.reason || error.message}`);
      setStep('error');
    } finally {
      setLoading(false);
    }
  };

  const maxBalance = () => {
    setAmount('0.1');
  };

  return (
    <div className="p-6 bg-gradient-to-r from-green-400 to-blue-500 rounded-lg shadow-lg text-white">
      <h2 className="text-2xl font-bold mb-4">🌉 Bridge Tokens</h2>
      
      <div className="bg-white/20 p-4 rounded-lg mb-4">
        <div className="flex justify-between items-center mb-2">
          <span className="font-semibold">From:</span>
          <span className="bg-blue-500 px-2 py-1 rounded">Sepolia</span>
        </div>
        <div className="flex justify-between items-center">
          <span className="font-semibold">To:</span>
          <span className="bg-purple-500 px-2 py-1 rounded">Polygon Amoy</span>
        </div>
        
        {step && (
          <div className="mt-3 space-y-2 text-sm">
            <div className={`flex items-center ${step === 'approve' ? 'font-bold' : ''}`}>
              <span className="mr-2">{step === 'approve' ? '🔄' : '✅'}</span>
              <span>1. Approve Tokens</span>
            </div>
            <div className={`flex items-center ${step === 'lock' ? 'font-bold' : step === 'complete' ? 'text-green-300' : ''}`}>
              <span className="mr-2">
                {step === 'lock' ? '🔄' : step === 'complete' ? '✅' : '⏳'}
              </span>
              <span>2. Lock Tokens</span>
            </div>
            {step === 'complete' && (
              <div className="flex items-center text-yellow-300">
                <span className="mr-2">⏳</span>
                <span>3. Relayer Processing</span>
              </div>
            )}
          </div>
        )}
      </div>
      
      <div className="space-y-4">
        <div>
          <label className="block text-sm font-medium mb-2">
            Amount to Bridge
          </label>
          <div className="flex gap-2">
            <input
              type="number"
              step="0.001"
              value={amount}
              onChange={(e) => setAmount(e.target.value)}
              placeholder="0.1"
              className="flex-1 p-3 border border-gray-300 rounded-md text-gray-800 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              disabled={loading}
            />
            <button
              onClick={maxBalance}
              className="bg-yellow-500 text-white px-4 py-3 rounded-md hover:bg-yellow-600 transition-colors"
              disabled={loading}
            >
              TEST
            </button>
          </div>
        </div>

        <button
          onClick={bridgeTokens}
          disabled={loading || !amount}
          className="w-full bg-white text-green-600 py-3 rounded-md font-bold hover:bg-gray-100 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors"
        >
          {loading ? (
            <span className="flex items-center justify-center">
              <div className="animate-spin rounded-full h-5 w-5 border-b-2 border-green-600 mr-2"></div>
              Processing...
            </span>
          ) : (
            '🚀 Start Bridge Process'
          )}
        </button>

        {status && (
          <div className={`p-3 rounded-md ${
            status.includes('✅') || status.includes('🎉') ? 'bg-green-500/20 border border-green-400' : 
            status.includes('❌') ? 'bg-red-500/20 border border-red-400' : 
            'bg-blue-500/20 border border-blue-400'
          }`}>
            <div className="font-semibold">{status}</div>
            {txHash && (
              <a 
                href={`https://sepolia.etherscan.io/tx/${txHash}`}
                target="_blank"
                rel="noopener noreferrer"
                className="text-sm underline mt-1 inline-block"
              >
                🔍 View on Etherscan
              </a>
            )}
          </div>
        )}

        {txHash && step === 'complete' && (
          <TransactionStatus txHash={txHash} account={account} />
        )}
      </div>
    </div>
  );
};
