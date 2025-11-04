package testcases

import (
    "context"
    "testing"
    
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/stretchr/testify/assert"
)

func TestBasicConnection(t *testing.T) {
    // Простой тест подключения к блокчейну
    client, err := ethclient.Dial("http://localhost:8545")
    assert.NoError(t, err)
    
    block, err := client.BlockByNumber(context.Background(), nil)
    assert.NoError(t, err)
    
    t.Logf("✅ Connected to Ethereum! Block: %d", block.Number())
}
