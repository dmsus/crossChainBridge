package security

import (
    "fmt"
    "net/http"
    "strconv"
    "strings"
    "time"

    "github.com/sirupsen/logrus"
)

// SecurityMiddleware provides security features for HTTP API
type SecurityMiddleware struct {
    rateLimiter    *RateLimiter
    replayProtector *ReplayProtector
    logger         *logrus.Logger
    config         SecurityConfig
}

type SecurityConfig struct {
    EnableRateLimiting   bool
    RequestsPerMinute    int
    BurstSize           int
    TimestampWindow     time.Duration
    BlockedIPs          []string
}

func NewSecurityMiddleware(rateLimiter *RateLimiter, replayProtector *ReplayProtector, config SecurityConfig, logger *logrus.Logger) *SecurityMiddleware {
    return &SecurityMiddleware{
        rateLimiter:    rateLimiter,
        replayProtector: replayProtector,
        logger:         logger,
        config:         config,
    }
}

// RateLimitMiddleware applies rate limiting to HTTP handlers
func (sm *SecurityMiddleware) RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !sm.config.EnableRateLimiting {
            next.ServeHTTP(w, r)
            return
        }

        clientIP := GetClientIP(r.RemoteAddr)
        
        // Check if IP is blocked
        if sm.isIPBlocked(clientIP) {
            sm.logger.Warnf("Blocked IP attempted access: %s", clientIP)
            http.Error(w, "Access denied", http.StatusForbidden)
            return
        }

        // Apply rate limiting
        if !sm.rateLimiter.AllowIP(clientIP, 1) {
            sm.logger.Warnf("Rate limit exceeded for IP: %s", clientIP)
            w.Header().Set("Retry-After", "60")
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }

        next.ServeHTTP(w, r)
    })
}

// SecurityHeadersMiddleware adds security headers to responses
func (sm *SecurityMiddleware) SecurityHeadersMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Add security headers
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
        
        // Content Security Policy
        csp := "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'"
        w.Header().Set("Content-Security-Policy", csp)

        next.ServeHTTP(w, r)
    })
}

// InputValidationMiddleware validates and sanitizes input
func (sm *SecurityMiddleware) InputValidationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Validate content type for POST requests
        if r.Method == http.MethodPost {
            contentType := r.Header.Get("Content-Type")
            if !strings.Contains(contentType, "application/json") {
                http.Error(w, "Unsupported content type", http.StatusUnsupportedMediaType)
                return
            }
        }

        // Validate and sanitize query parameters
        if err := sm.validateQueryParams(r); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func (sm *SecurityMiddleware) isIPBlocked(ip string) bool {
    for _, blockedIP := range sm.config.BlockedIPs {
        if ip == blockedIP {
            return true
        }
    }
    return false
}

func (sm *SecurityMiddleware) validateQueryParams(r *http.Request) error {
    query := r.URL.Query()
    
    // Validate numeric parameters
    if limit := query.Get("limit"); limit != "" {
        if _, err := strconv.Atoi(limit); err != nil {
            return fmt.Errorf("invalid limit parameter")
        }
    }
    
    if offset := query.Get("offset"); offset != "" {
        if _, err := strconv.Atoi(offset); err != nil {
            return fmt.Errorf("invalid offset parameter")
        }
    }

    // Validate address parameters (basic format check)
    if address := query.Get("address"); address != "" {
        if !strings.HasPrefix(address, "0x") || len(address) != 42 {
            return fmt.Errorf("invalid address format")
        }
    }

    return nil
}

// LogSecurityEvent logs security-related events
func (sm *SecurityMiddleware) LogSecurityEvent(eventType, message, ip string, severity logrus.Level) {
    fields := logrus.Fields{
        "type":      eventType,
        "ip":        ip,
        "timestamp": time.Now().UTC(),
    }

    switch severity {
    case logrus.WarnLevel:
        sm.logger.WithFields(fields).Warn(message)
    case logrus.ErrorLevel:
        sm.logger.WithFields(fields).Error(message)
    default:
        sm.logger.WithFields(fields).Info(message)
    }
}
