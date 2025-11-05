import React from 'react';
import { WalletConnector } from './components/WalletConnector';
import { TokenBalance } from './components/TokenBalance';
import { BridgeForm } from './components/BridgeForm';
import { TransactionHistory } from './components/TransactionHistory';
import { useWeb3 } from './hooks/useWeb3';

function App() {
  const { 
    account, 
    provider,
    signer,
    connectWallet, 
    disconnectWallet, 
    currentChain,
    switchNetwork 
  } = useWeb3();

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-500 to-purple-600 py-8 px-4">
      <div className="max-w-4xl mx-auto">
        {/* Header */}
        <div className="text-center mb-8">
          <h1 className="text-4xl font-bold text-white mb-2">
            🌉 Cross-Chain Bridge
          </h1>
          <p className="text-blue-100">
            Transfer tokens between Ethereum Sepolia and Polygon Amoy
          </p>
        </div>

        {/* Main Content */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          {/* Left Sidebar */}
          <div className="lg:col-span-1 space-y-6">
            <WalletConnector 
              account={account}
              currentChain={currentChain}
              onConnect={connectWallet}
              onDisconnect={disconnectWallet}
              onSwitchNetwork={switchNetwork}
            />
            <TokenBalance account={account} provider={provider} signer={signer} />
          </div>

          {/* Main Bridge Interface */}
          <div className="lg:col-span-2 space-y-6">
            <div className="bg-white rounded-xl shadow-lg p-6">
              <BridgeForm account={account} provider={provider} signer={signer} />
            </div>
            
            <div className="bg-white rounded-xl shadow-lg p-6">
              <TransactionHistory account={account} />
            </div>
          </div>
        </div>

        {/* Footer */}
        <div className="text-center mt-8 text-blue-100">
          <p>Built with ❤️ by dmsus | 
            <a href="https://github.com/dmsus/crossChainBridge" 
               className="underline ml-2" target="_blank" rel="noopener">
              View on GitHub
            </a>
          </p>
        </div>
      </div>
    </div>
  );
}

export default App;
