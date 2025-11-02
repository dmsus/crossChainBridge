package signer

import (
    "crypto/ecdsa"
    "encoding/hex"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/common/math"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// EIP712Signer реализует EIP-712 подписи для кросс-чейн верификации
type EIP712Signer struct {
    privateKey *ecdsa.PrivateKey
    domain     apitypes.TypedDataDomain
}

// BridgeMessage представляет сообщение для подписи моста
type BridgeMessage struct {
    User          common.Address `json:"user"`
    Amount        *big.Int       `json:"amount"`
    Nonce         *big.Int       `json:"nonce"`
    TargetChainID *big.Int       `json:"targetChainId"`
}

// Config конфигурация для EIP712Signer
type Config struct {
    PrivateKey    string
    BridgeAddress common.Address
    ChainID       *big.Int
}

// NewEIP712Signer создает новый экземпляр подписывателя
func NewEIP712Signer(cfg Config) (*EIP712Signer, error) {
    privateKey, err := crypto.HexToECDSA(cfg.PrivateKey[2:]) // убираем 0x префикс
    if err != nil {
        return nil, fmt.Errorf("failed to parse private key: %v", err)
    }

    // Создаем EIP-712 domain
    domain := apitypes.TypedDataDomain{
        Name:              "CrossChainBridge",
        Version:           "1",
        ChainId:           math.NewHexOrDecimal256(cfg.ChainID.Int64()),
        VerifyingContract: cfg.BridgeAddress.Hex(),
        Salt:              "", // Опционально
    }

    log.Printf("✅ EIP712Signer initialized: chainID=%s, bridge=%s", 
        cfg.ChainID.String(), cfg.BridgeAddress.Hex())

    return &EIP712Signer{
        privateKey: privateKey,
        domain:     domain,
    }, nil
}

// SignBridgeMessage создает EIP-712 подпись для сообщения моста
func (s *EIP712Signer) SignBridgeMessage(message *BridgeMessage) ([]byte, error) {
    log.Printf("🔐 Signing bridge message: user=%s, amount=%s, nonce=%s, targetChain=%s",
        message.User.Hex(), message.Amount.String(), message.Nonce.String(), message.TargetChainID.String())

    // Создаем typed data для EIP-712
    typedData := apitypes.TypedData{
        Types: apitypes.Types{
            "EIP712Domain": []apitypes.Type{
                {Name: "name", Type: "string"},
                {Name: "version", Type: "string"},
                {Name: "chainId", Type: "uint256"},
                {Name: "verifyingContract", Type: "address"},
            },
            "BridgeTransfer": []apitypes.Type{
                {Name: "user", Type: "address"},
                {Name: "amount", Type: "uint256"},
                {Name: "nonce", Type: "uint256"},
                {Name: "targetChainId", Type: "uint256"},
            },
        },
        PrimaryType: "BridgeTransfer",
        Domain:      s.domain,
        Message: apitypes.TypedDataMessage{
            "user":          message.User.Hex(),
            "amount":        math.NewHexOrDecimal256(message.Amount.Int64()),
            "nonce":         math.NewHexOrDecimal256(message.Nonce.Int64()),
            "targetChainId": math.NewHexOrDecimal256(message.TargetChainID.Int64()),
        },
    }

    // Вычисляем хеш typed data
    hash, err := s.hashTypedData(typedData)
    if err != nil {
        return nil, fmt.Errorf("failed to hash typed data: %v", err)
    }

    // Подписываем хеш
    signature, err := crypto.Sign(hash, s.privateKey)
    if err != nil {
        return nil, fmt.Errorf("failed to sign hash: %v", err)
    }

    // В ECDSA нужно установить v = 27 или 28 (recovery ID)
    if signature[64] < 27 {
        signature[64] += 27
    }

    log.Printf("✅ Message signed successfully: signature=%s", hex.EncodeToString(signature))

    return signature, nil
}

// hashTypedData вычисляет хеш EIP-712 typed data
func (s *EIP712Signer) hashTypedData(typedData apitypes.TypedData) ([]byte, error) {
    domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
    if err != nil {
        return nil, err
    }

    hashStruct, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
    if err != nil {
        return nil, err
    }

    rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(hashStruct)))
    return crypto.Keccak256(rawData), nil
}

// VerifySignature проверяет EIP-712 подпись
func (s *EIP712Signer) VerifySignature(message *BridgeMessage, signature []byte) (bool, common.Address, error) {
    // Создаем typed data (такой же как при подписи)
    typedData := apitypes.TypedData{
        Types: apitypes.Types{
            "EIP712Domain": []apitypes.Type{
                {Name: "name", Type: "string"},
                {Name: "version", Type: "string"},
                {Name: "chainId", Type: "uint256"},
                {Name: "verifyingContract", Type: "address"},
            },
            "BridgeTransfer": []apitypes.Type{
                {Name: "user", Type: "address"},
                {Name: "amount", Type: "uint256"},
                {Name: "nonce", Type: "uint256"},
                {Name: "targetChainId", Type: "uint256"},
            },
        },
        PrimaryType: "BridgeTransfer",
        Domain:      s.domain,
        Message: apitypes.TypedDataMessage{
            "user":          message.User.Hex(),
            "amount":        math.NewHexOrDecimal256(message.Amount.Int64()),
            "nonce":         math.NewHexOrDecimal256(message.Nonce.Int64()),
            "targetChainId": math.NewHexOrDecimal256(message.TargetChainID.Int64()),
        },
    }

    // Вычисляем хеш
    hash, err := s.hashTypedData(typedData)
    if err != nil {
        return false, common.Address{}, err
    }

    // Восстанавливаем публичный ключ из подписи
    if signature[64] >= 27 {
        signature[64] -= 27
    }

    pubKey, err := crypto.SigToPub(hash, signature)
    if err != nil {
        return false, common.Address{}, err
    }

    // Получаем адрес из публичного ключа
    recoveredAddr := crypto.PubkeyToAddress(*pubKey)

    // Проверяем что подписант совпадает с user в сообщении
    isValid := recoveredAddr == message.User

    log.Printf("🔍 Signature verification: valid=%t, signer=%s, expected=%s",
        isValid, recoveredAddr.Hex(), message.User.Hex())

    return isValid, recoveredAddr, nil
}

// GetPublicAddress возвращает адрес соответствующий приватному ключу
func (s *EIP712Signer) GetPublicAddress() common.Address {
    publicKey := s.privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return common.Address{}
    }
    return crypto.PubkeyToAddress(*publicKeyECDSA)
}

// GenerateTestKey генерирует тестовый приватный ключ (для тестов)
func GenerateTestKey() (*ecdsa.PrivateKey, error) {
    return crypto.GenerateKey()
}
