package testcases

import (
    "context"
    "math/big"
    "testing"
    "time"

    "github.com/dmsus/crossChainBridge/tests/e2e/bindings"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

type E2EHappyPathTestSuite struct {
    suite.Suite
    ethClient     *ethclient.Client
    polygonClient *ethclient.Client
    ctx           context.Context
    
    // Контракты Ethereum
    tokenEthereum   *bindings.TokenERC20
    bridgeEthereum  *bindings.BridgeEthereum
    
    // Адреса контрактов
    tokenEthereumAddr  common.Address
    bridgeEthereumAddr common.Address
    
    // Транзакционные опции
    auth *bind.TransactOpts
}

func (s *E2EHappyPathTestSuite) SetupSuite() {
    s.ctx = context.Background()
    
    var err error
    s.ethClient, err = ethclient.Dial("http://localhost:8545")
    assert.NoError(s.T(), err)
    
    s.polygonClient, err = ethclient.Dial("http://localhost:8546") 
    assert.NoError(s.T(), err)
    
    // Адреса контрактов из деплоя
    s.tokenEthereumAddr = common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
    s.bridgeEthereumAddr = common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512")
    
    // Инициализируем транзакционные опции (первый аккаунт из Anvil)
    privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
    assert.NoError(s.T(), err)
    
    s.auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))
    assert.NoError(s.T(), err)
    
    // Инициализируем контракты Ethereum
    s.tokenEthereum, err = bindings.NewTokenERC20(s.tokenEthereumAddr, s.ethClient)
    assert.NoError(s.T(), err)
    
    s.bridgeEthereum, err = bindings.NewBridgeEthereum(s.bridgeEthereumAddr, s.ethClient)
    assert.NoError(s.T(), err)
}

func (s *E2EHappyPathTestSuite) TestBridgeLockTokens() {
    t := s.T()
    
    t.Log("🧪 Testing bridge lock tokens...")
    
    // 1. Получаем текущий баланс
    initialBalance, err := s.tokenEthereum.BalanceOf(&bind.CallOpts{}, s.auth.From)
    assert.NoError(t, err)
    t.Logf("   Initial balance: %s", initialBalance.String())
    
    // 2. Если баланс маленький - минтим
    minBalance := big.NewInt(100000000000000000) // 0.1 токен
    if initialBalance.Cmp(minBalance) < 0 {
        mintAmount := big.NewInt(1000000000000000000) // 1 токен
        mintTx, err := s.tokenEthereum.Mint(s.auth, s.auth.From, mintAmount)
        assert.NoError(t, err)
        s.waitForTransaction(s.ethClient, mintTx.Hash())
        initialBalance = mintAmount
    }
    
    // 3. Определяем сумму для блокировки (10% от баланса)
    lockAmount := new(big.Int).Div(initialBalance, big.NewInt(10))
    t.Logf("1. Locking %s tokens...", lockAmount.String())
    
    // 4. Одобряем бридж тратить токены
    t.Log("2. Approving bridge to spend tokens...")
    approveTx, err := s.tokenEthereum.Approve(s.auth, s.bridgeEthereumAddr, lockAmount)
    assert.NoError(t, err)
    s.waitForTransaction(s.ethClient, approveTx.Hash())
    
    // 5. Проверяем allowance
    allowance, err := s.tokenEthereum.Allowance(&bind.CallOpts{}, s.auth.From, s.bridgeEthereumAddr)
    assert.NoError(t, err)
    assert.Equal(t, lockAmount, allowance, "Allowance should match approved amount")
    t.Logf("   ✅ Allowance set: %s", allowance.String())
    
    // 6. Блокируем токены в бридже
    t.Log("3. Locking tokens in bridge...")
    targetChainId := big.NewInt(31338) // ChainId Polygon сети
    
    lockTx, err := s.bridgeEthereum.LockTokens(s.auth, lockAmount, s.auth.From, targetChainId)
    assert.NoError(t, err)
    s.waitForTransaction(s.ethClient, lockTx.Hash())
    t.Logf("   ✅ Lock transaction: %s", lockTx.Hash().Hex())
    
    // 7. Проверяем что токены заблокированы (баланс уменьшился)
    balanceAfterLock, err := s.tokenEthereum.BalanceOf(&bind.CallOpts{}, s.auth.From)
    assert.NoError(t, err)
    
    expectedBalance := new(big.Int).Sub(initialBalance, lockAmount)
    assert.Equal(t, expectedBalance, balanceAfterLock, "Balance should decrease by lock amount")
    t.Logf("   ✅ Balance after lock: %s (expected: %s)", balanceAfterLock.String(), expectedBalance.String())
    
    // 8. Проверяем nonce увеличился
    nonce, err := s.bridgeEthereum.Nonce(&bind.CallOpts{})
    assert.NoError(t, err)
    assert.True(t, nonce.Cmp(big.NewInt(0)) > 0, "Nonce should be incremented")
    t.Logf("   ✅ Nonce after lock: %s", nonce.String())
    
    // 9. Проверяем что бридж получил токены
    bridgeBalance, err := s.tokenEthereum.BalanceOf(&bind.CallOpts{}, s.bridgeEthereumAddr)
    assert.NoError(t, err)
    t.Logf("   ✅ Bridge balance: %s", bridgeBalance.String())
    
    t.Log("🎉 Bridge lock test completed!")
}

func (s *E2EHappyPathTestSuite) TestTokenTransfer() {
    t := s.T()
    
    t.Log("🧪 Testing token transfer between accounts...")
    
    // Используем второй аккаунт из Anvil для чистого теста
    recipient := common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC")
    transferAmount := big.NewInt(100000000000000000) // 0.1 токен
    
    // 1. Получаем начальные балансы
    initialSenderBalance, err := s.tokenEthereum.BalanceOf(&bind.CallOpts{}, s.auth.From)
    assert.NoError(t, err)
    
    initialRecipientBalance, err := s.tokenEthereum.BalanceOf(&bind.CallOpts{}, recipient)
    assert.NoError(t, err)
    
    t.Logf("1. Transferring %s tokens to %s", transferAmount.String(), recipient.Hex())
    
    // 2. Выполняем трансфер
    transferTx, err := s.tokenEthereum.Transfer(s.auth, recipient, transferAmount)
    assert.NoError(t, err)
    s.waitForTransaction(s.ethClient, transferTx.Hash())
    
    // 3. Проверяем балансы после трансфера
    finalSenderBalance, err := s.tokenEthereum.BalanceOf(&bind.CallOpts{}, s.auth.From)
    assert.NoError(t, err)
    
    finalRecipientBalance, err := s.tokenEthereum.BalanceOf(&bind.CallOpts{}, recipient)
    assert.NoError(t, err)
    
    // Проверяем что баланс отправителя уменьшился на сумму трансфера + газ
    assert.True(t, finalSenderBalance.Cmp(initialSenderBalance) < 0, "Sender balance should decrease")
    
    // Проверяем что баланс получателя увеличился на сумму трансфера
    expectedRecipientBalance := new(big.Int).Add(initialRecipientBalance, transferAmount)
    assert.Equal(t, expectedRecipientBalance, finalRecipientBalance, "Recipient balance should increase by transfer amount")
    
    t.Logf("   ✅ Sender balance: %s -> %s", initialSenderBalance.String(), finalSenderBalance.String())
    t.Logf("   ✅ Recipient balance: %s -> %s", initialRecipientBalance.String(), finalRecipientBalance.String())
    
    t.Log("🎉 Token transfer test completed!")
}

func (s *E2EHappyPathTestSuite) waitForTransaction(client *ethclient.Client, txHash common.Hash) {
    for i := 0; i < 30; i++ {
        _, isPending, err := client.TransactionByHash(s.ctx, txHash)
        if err != nil {
            time.Sleep(1 * time.Second)
            continue
        }
        if !isPending {
            return
        }
        time.Sleep(1 * time.Second)
    }
    s.T().Fatalf("Transaction %s not mined after 30 seconds", txHash.Hex())
}

func TestE2EHappyPath(t *testing.T) {
    suite.Run(t, new(E2EHappyPathTestSuite))
}
