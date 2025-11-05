package security

import (
    "github.com/sirupsen/logrus"
)

// SecurityComponents объединяет все security компоненты
type SecurityComponents struct {
    RateLimiter     *RateLimiter
    ReplayProtector *ReplayProtector
    Middleware      *SecurityMiddleware
}

// SecurityConfig содержит конфигурацию безопасности
type SecurityConfig struct {
    EnableRateLimiting bool
    RequestsPerMinute  int
    BurstSize          int
    TimestampWindow    int64
    BlockedIPs         []string
}

// RateLimitConfig содержит конфигурацию rate limiting
type RateLimitConfig struct {
    DefaultRequestsPerMinute int
    DefaultBurstSize         int
    CleanupInterval          int
}
