package signer

import (
    "bytes"
    "encoding/hex"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
)

// ABIEncoder кодирует данные для вызова функций контракта
type ABIEncoder struct{}

// NewABIEncoder создает новый энкодер
func NewABIEncoder() *ABIEncoder {
    return &ABIEncoder{}
}

// EncodeReleaseTokensCall кодирует вызов функции releaseTokens
// function releaseTokens(address user, uint256 amount, uint256 nonce, bytes memory signature)
func (e *ABIEncoder) EncodeReleaseTokensCall(user common.Address, amount *big.Int, nonce *big.Int, signature []byte) ([]byte, error) {
    // ABI определение функции releaseTokens
    releaseTokensABI := `[{
        "inputs": [
            {"name": "user", "type": "address"},
            {"name": "amount", "type": "uint256"},
            {"name": "nonce", "type": "uint256"},
            {"name": "signature", "type": "bytes"}
        ],
        "name": "releaseTokens",
        "type": "function"
    }]`

    // Парсим ABI - используем bytes.NewReader
    parsedABI, err := abi.JSON(bytes.NewReader([]byte(releaseTokensABI)))
    if err != nil {
        return nil, err
    }

    // Кодируем данные вызова
    data, err := parsedABI.Pack("releaseTokens", user, amount, nonce, signature)
    if err != nil {
        return nil, err
    }

    log.Printf("📝 ABI encoded releaseTokens call: data=%s", hex.EncodeToString(data))

    return data, nil
}

// DecodeReleaseTokensCall декодирует данные вызова releaseTokens
func (e *ABIEncoder) DecodeReleaseTokensCall(data []byte) (common.Address, *big.Int, *big.Int, []byte, error) {
    releaseTokensABI := `[{
        "inputs": [
            {"name": "user", "type": "address"},
            {"name": "amount", "type": "uint256"},
            {"name": "nonce", "type": "uint256"},
            {"name": "signature", "type": "bytes"}
        ],
        "name": "releaseTokens",
        "type": "function"
    }]`

    parsedABI, err := abi.JSON(bytes.NewReader([]byte(releaseTokensABI)))
    if err != nil {
        return common.Address{}, nil, nil, nil, err
    }

    var user common.Address
    var amount *big.Int
    var nonce *big.Int
    var signature []byte

    // Декодируем данные
    method, err := parsedABI.MethodById(data[:4])
    if err != nil {
        return common.Address{}, nil, nil, nil, err
    }

    values, err := method.Inputs.Unpack(data[4:])
    if err != nil {
        return common.Address{}, nil, nil, nil, err
    }

    if len(values) >= 4 {
        user = values[0].(common.Address)
        amount = values[1].(*big.Int)
        nonce = values[2].(*big.Int)
        signature = values[3].([]byte)
    }

    return user, amount, nonce, signature, nil
}
