package middleware

import (
    "net/http"
    "strings"

    "github.com/dmsus/crossChainBridge/relayer/internal/security"
)

// SecurityMiddleware объединяет все security middleware
type SecurityMiddleware struct {
    rateLimiter *security.RateLimiter
}

// NewSecurityMiddleware создает новый security middleware
func NewSecurityMiddleware(rateLimiter *security.RateLimiter) *SecurityMiddleware {
    return &SecurityMiddleware{
        rateLimiter: rateLimiter,
    }
}

// RateLimitMiddleware добавляет rate limiting
func (sm *SecurityMiddleware) RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        clientIP := getClientIP(r)
        
        // Пропускаем health checks и metrics
        if r.URL.Path == "/health" || r.URL.Path == "/metrics" {
            next.ServeHTTP(w, r)
            return
        }

        if !sm.rateLimiter.AllowIP(clientIP, 1) {
            respondWithError(w, http.StatusTooManyRequests, "RateLimitExceeded", "Too many requests")
            return
        }

        next.ServeHTTP(w, r)
    })
}

// SecurityHeadersMiddleware добавляет security headers
func SecurityHeadersMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
        
        next.ServeHTTP(w, r)
    })
}

// getClientIP извлекает реальный IP клиента
func getClientIP(r *http.Request) string {
    // Проверяем заголовки прокси
    if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
        return strings.Split(ip, ",")[0]
    }
    if ip := r.Header.Get("X-Real-IP"); ip != "" {
        return ip
    }
    return strings.Split(r.RemoteAddr, ":")[0]
}

// AdminOnlyMiddleware ограничивает доступ только для администраторов
func AdminOnlyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // В реальной реализации здесь была бы проверка ролей
        // Сейчас просто проверяем что это admin endpoint
        if strings.HasPrefix(r.URL.Path, "/api/admin") {
            apiKey := GetAPIKeyFromContext(r.Context())
            if apiKey == "" {
                respondWithError(w, http.StatusForbidden, "Forbidden", "Admin access required")
                return
            }
        }
        
        next.ServeHTTP(w, r)
    })
}
