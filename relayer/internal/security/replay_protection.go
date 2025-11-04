package security

import (
    "context"
    "database/sql"
    "fmt"
    "time"

    "github.com/dmsus/crossChainBridge/relayer/pkg/database"
    "github.com/sirupsen/logrus"
)

// ReplayProtector предотвращает replay-атаки через nonce tracking
type ReplayProtector struct {
    db     *database.Repository
    logger *logrus.Logger
}

// NewReplayProtector создает новый replay protector
func NewReplayProtector(dbRepo *database.Repository, logger *logrus.Logger) *ReplayProtector {
    return &ReplayProtector{
        db:     dbRepo,
        logger: logger,
    }
}

// ValidateNonceSequence проверяет последовательность nonce для пользователя
func (rp *ReplayProtector) ValidateNonceSequence(ctx context.Context, userAddress string, chainID int64, nonce int64) (bool, error) {
    // Получаем последний использованный nonce из БД
    lastNonce, err := rp.getLastNonce(ctx, userAddress, chainID)
    if err != nil {
        return false, fmt.Errorf("failed to get last nonce: %w", err)
    }

    // Проверяем последовательность (должен быть lastNonce + 1)
    if nonce != lastNonce+1 {
        rp.logger.Warnf("Invalid nonce sequence for user %s: expected %d, got %d", 
            userAddress, lastNonce+1, nonce)
        return false, nil
    }

    // Обновляем nonce в БД
    if err := rp.updateNonce(ctx, userAddress, chainID, nonce); err != nil {
        return false, fmt.Errorf("failed to update nonce: %w", err)
    }

    rp.logger.Debugf("Valid nonce sequence for user %s: %d -> %d", userAddress, lastNonce, nonce)
    return true, nil
}

// ValidateTimestamp проверяет временную метку для предотвращения replay атак
func (rp *ReplayProtector) ValidateTimestamp(ctx context.Context, timestamp int64, window time.Duration) bool {
    now := time.Now().Unix()
    windowSeconds := int64(window.Seconds())
    
    // Проверяем что timestamp в пределах допустимого окна
    if timestamp < now-windowSeconds || timestamp > now+windowSeconds {
        rp.logger.Warnf("Invalid timestamp: %d (now: %d, window: %ds)", timestamp, now, windowSeconds)
        return false
    }
    
    return true
}

// getLastNonce возвращает последний использованный nonce для пользователя и сети
func (rp *ReplayProtector) getLastNonce(ctx context.Context, userAddress string, chainID int64) (int64, error) {
    var lastNonce int64
    
    query := `
        SELECT last_nonce FROM nonce_tracking 
        WHERE user_address = $1 AND chain_id = $2
    `
    
    err := rp.db.DB().QueryRowContext(ctx, query, userAddress, chainID).Scan(&lastNonce)
    if err != nil {
        if err == sql.ErrNoRows {
            // Если записи нет, начинаем с 0
            return 0, nil
        }
        return 0, err
    }
    
    return lastNonce, nil
}

// updateNonce обновляет nonce для пользователя и сети
func (rp *ReplayProtector) updateNonce(ctx context.Context, userAddress string, chainID int64, nonce int64) error {
    query := `
        INSERT INTO nonce_tracking (user_address, chain_id, last_nonce, updated_at)
        VALUES ($1, $2, $3, NOW())
        ON CONFLICT (user_address, chain_id) 
        DO UPDATE SET last_nonce = $3, updated_at = NOW()
    `
    
    _, err := rp.db.DB().ExecContext(ctx, query, userAddress, chainID, nonce)
    return err
}

// CleanupExpiredNonces очищает устаревшие записи nonce
func (rp *ReplayProtector) CleanupExpiredNonces(ctx context.Context, retentionPeriod time.Duration) error {
    query := `
        DELETE FROM nonce_tracking 
        WHERE updated_at < NOW() - INTERVAL '1 day' * $1
    `
    
    days := int(retentionPeriod.Hours() / 24)
    _, err := rp.db.DB().ExecContext(ctx, query, days)
    return err
}
