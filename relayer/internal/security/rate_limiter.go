package security

import (
    "net"
    "sync"
    "time"

    "github.com/sirupsen/logrus"
)

// RateLimiter implements token bucket algorithm
type RateLimiter struct {
    mu          sync.RWMutex
    limits      map[string]*tokenBucket
    config      RateLimitConfig
    logger      *logrus.Logger
}

type RateLimitConfig struct {
    DefaultRequestsPerMinute int
    DefaultBurstSize         int
    CleanupInterval          time.Duration
}

type tokenBucket struct {
    tokens     float64
    capacity   float64
    lastRefill time.Time
    rate       float64 // tokens per second
}

func NewRateLimiter(config RateLimitConfig, logger *logrus.Logger) *RateLimiter {
    rl := &RateLimiter{
        limits: make(map[string]*tokenBucket),
        config: config,
        logger: logger,
    }

    // Start cleanup goroutine
    go rl.cleanupExpiredBuckets()

    return rl
}

// Allow checks if request is allowed for given key
func (rl *RateLimiter) Allow(key string) bool {
    return rl.AllowN(key, 1)
}

// AllowN checks if N requests are allowed for given key
func (rl *RateLimiter) AllowN(key string, n int) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()

    bucket, exists := rl.limits[key]
    if !exists {
        bucket = rl.newTokenBucket(key)
        rl.limits[key] = bucket
    }

    return bucket.allow(n)
}

// AllowIP checks if request is allowed for IP address
func (rl *RateLimiter) AllowIP(ip string, n int) bool {
    return rl.AllowN("ip:"+ip, n)
}

// AllowUser checks if request is allowed for user address
func (rl *RateLimiter) AllowUser(userAddress string, n int) bool {
    return rl.AllowN("user:"+userAddress, n)
}

func (rl *RateLimiter) newTokenBucket(key string) *tokenBucket {
    ratePerSecond := float64(rl.config.DefaultRequestsPerMinute) / 60.0
    return &tokenBucket{
        tokens:     float64(rl.config.DefaultBurstSize),
        capacity:   float64(rl.config.DefaultBurstSize),
        lastRefill: time.Now(),
        rate:       ratePerSecond,
    }
}

func (tb *tokenBucket) allow(n int) bool {
    now := time.Now()
    elapsed := now.Sub(tb.lastRefill).Seconds()
    tb.lastRefill = now

    // Refill tokens
    tb.tokens += elapsed * tb.rate
    if tb.tokens > tb.capacity {
        tb.tokens = tb.capacity
    }

    // Check if enough tokens
    if tb.tokens >= float64(n) {
        tb.tokens -= float64(n)
        return true
    }
    return false
}

func (rl *RateLimiter) cleanupExpiredBuckets() {
    ticker := time.NewTicker(rl.config.CleanupInterval)
    defer ticker.Stop()

    for range ticker.C {
        rl.mu.Lock()
        for key, bucket := range rl.limits {
            // Remove buckets that haven't been used in 1 hour
            if time.Since(bucket.lastRefill) > time.Hour {
                delete(rl.limits, key)
            }
        }
        rl.mu.Unlock()
    }
}

// GetClientIP extracts client IP from request
func GetClientIP(remoteAddr string) string {
    host, _, err := net.SplitHostPort(remoteAddr)
    if err != nil {
        return remoteAddr
    }
    return host
}
