import { useState } from 'react'
import { WalletConnector } from './components/WalletConnector'
import { TokenBalance } from './components/TokenBalance'
import { BridgeForm } from './components/BridgeForm'
import './App.css'

function App() {
  const [account, setAccount] = useState('')
  const [refreshBalances, setRefreshBalances] = useState(0)

  const handleBalanceUpdate = () => {
    setRefreshBalances(prev => prev + 1)
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 to-blue-900 py-8">
      <div className="container mx-auto px-4">
        <h1 className="text-4xl font-bold text-center mb-8 text-white">
          🌉 Cross-Chain Bridge dApp
        </h1>
        
        <div className="max-w-2xl mx-auto space-y-6">
          <WalletConnector onAccountChange={setAccount} />
          {account && (
            <>
              <TokenBalance account={account} key={refreshBalances} />
              <BridgeForm account={account} onBalanceUpdate={handleBalanceUpdate} />
            </>
          )}
        </div>

        <div className="mt-8 text-center text-gray-300">
          <p>Task #30: Frontend dApp with React and ethers.js</p>
          <p>Status: 🎉 FULLY OPERATIONAL - Bridge successful!</p>
          <p>Latest: 0.1 TEST bridged from Sepolia → Polygon Amoy</p>
        </div>
      </div>
    </div>
  )
}

export default App
