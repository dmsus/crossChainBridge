package middleware

import (
    "context"
    "encoding/json"
    "net/http"
    "strings"
)

type contextKey string

const (
    APIKeyHeader = "X-API-Key"
    contextKeyAPIKey contextKey = "api_key"
)

// APIKeyAuth middleware для аутентификации по API ключу
type APIKeyAuth struct {
    validKeys map[string]bool
}

// NewAPIKeyAuth создает новый middleware аутентификации
func NewAPIKeyAuth(validKeys []string) *APIKeyAuth {
    keyMap := make(map[string]bool)
    for _, key := range validKeys {
        keyMap[key] = true
    }
    return &APIKeyAuth{validKeys: keyMap}
}

// Middleware функция для проверки API ключа
func (a *APIKeyAuth) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Пропускаем публичные endpoints
        if isPublicEndpoint(r.URL.Path) {
            next.ServeHTTP(w, r)
            return
        }

        apiKey := r.Header.Get(APIKeyHeader)
        if apiKey == "" {
            apiKey = r.URL.Query().Get("api_key")
        }

        if apiKey == "" || !a.validKeys[apiKey] {
            respondWithError(w, http.StatusUnauthorized, "Unauthorized", "Invalid or missing API key")
            return
        }

        // Добавляем API ключ в контекст
        ctx := context.WithValue(r.Context(), contextKeyAPIKey, apiKey)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// GetAPIKeyFromContext извлекает API ключ из контекста
func GetAPIKeyFromContext(ctx context.Context) string {
    if key, ok := ctx.Value(contextKeyAPIKey).(string); ok {
        return key
    }
    return ""
}

// isPublicEndpoint проверяет является ли endpoint публичным
func isPublicEndpoint(path string) bool {
    publicPaths := []string{
        "/health",
        "/metrics", 
        "/api/bridge/status/",
        "/api/bridge/fees",
        "/api/bridge/limits",
        "/api/bridge/tokens",
        "/api/bridge/estimate/",
    }

    for _, publicPath := range publicPaths {
        if strings.HasPrefix(path, publicPath) {
            return true
        }
    }
    return false
}

// respondWithError возвращает ошибку в JSON формате
func respondWithError(w http.ResponseWriter, code int, errorType, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    response := map[string]interface{}{
        "error":   errorType,
        "message": message,
        "code":    code,
    }
    
    if jsonBytes, err := json.Marshal(response); err == nil {
        w.Write(jsonBytes)
    }
}
