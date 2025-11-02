package config

import (
    "fmt"
    "log"
    "os"
    "strings"
    "time"

    "github.com/spf13/viper"
)

// Config представляет полную конфигурацию приложения
type Config struct {
    Environment string         `mapstructure:"environment"`
    Database    DatabaseConfig `mapstructure:"database"`
    Ethereum    EthereumConfig `mapstructure:"ethereum"`
    Polygon     PolygonConfig  `mapstructure:"polygon"`
    Service     ServiceConfig  `mapstructure:"service"`
    Processor   ProcessorConfig `mapstructure:"processor"`
    Monitoring  MonitoringConfig `mapstructure:"monitoring"`
    API         APIConfig      `mapstructure:"api"`
}

// DatabaseConfig конфигурация PostgreSQL
type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     string `mapstructure:"port"`
    Name     string `mapstructure:"name"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    SSLMode  string `mapstructure:"ssl_mode"`
    MaxConns int    `mapstructure:"max_conns"`
}

// EthereumConfig конфигурация Ethereum сети
type EthereumConfig struct {
    RPCURL      string `mapstructure:"rpc_url"`
    WsURL       string `mapstructure:"ws_url"`
    BridgeAddr  string `mapstructure:"bridge_addr"`
    TokenAddr   string `mapstructure:"token_addr"`
    ChainID     int64  `mapstructure:"chain_id"`
    PrivateKey  string `mapstructure:"private_key"`
    BlockConfirmations int `mapstructure:"block_confirmations"`
}

// PolygonConfig конфигурация Polygon сети
type PolygonConfig struct {
    RPCURL      string `mapstructure:"rpc_url"`
    WsURL       string `mapstructure:"ws_url"`
    BridgeAddr  string `mapstructure:"bridge_addr"`
    TokenAddr   string `mapstructure:"token_addr"`
    ChainID     int64  `mapstructure:"chain_id"`
    PrivateKey  string `mapstructure:"private_key"`
    GasLimit    uint64 `mapstructure:"gas_limit"`
}

// ServiceConfig конфигурация сервиса
type ServiceConfig struct {
    LogLevel    string `mapstructure:"log_level"`
    MetricsPort string `mapstructure:"metrics_port"`
    APIPort     string `mapstructure:"api_port"`
}

// ProcessorConfig конфигурация процессора
type ProcessorConfig struct {
    MaxRetries        int           `mapstructure:"max_retries"`
    RetryDelay        time.Duration `mapstructure:"retry_delay"`
    BatchSize         int           `mapstructure:"batch_size"`
    QueueSize         int           `mapstructure:"queue_size"`
    HealthCheckPeriod time.Duration `mapstructure:"health_check_period"`
}

// MonitoringConfig конфигурация мониторинга
type MonitoringConfig struct {
    Enabled     bool   `mapstructure:"enabled"`
    MetricsPort int    `mapstructure:"metrics_port"`
    LogLevel    string `mapstructure:"log_level"`
    LogFormat   string `mapstructure:"log_format"`
}

// APIConfig конфигурация API
type APIConfig struct {
    Enabled     bool   `mapstructure:"enabled"`
    Port        int    `mapstructure:"port"`
    SwaggerEnabled bool `mapstructure:"swagger_enabled"`
    RateLimit   int    `mapstructure:"rate_limit"`
}

// Load загружает конфигурацию для указанного окружения
func Load(env string) (*Config, error) {
    v := viper.New()
    
    // Базовая конфигурация
    v.SetConfigName(fmt.Sprintf(".env.%s", env))
    v.SetConfigType("yaml")
    v.AddConfigPath(".")
    v.AddConfigPath("./configs")
    
    // Автоматическое связывание с переменными окружения
    v.AutomaticEnv()
    v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    
    // Чтение конфигурационного файла
    if err := v.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("failed to read config file: %v", err)
    }
    
    var cfg Config
    if err := v.Unmarshal(&cfg); err != nil {
        return nil, fmt.Errorf("failed to unmarshal config: %v", err)
    }
    
    // Устанавливаем окружение
    cfg.Environment = env
    
    // ВАЖНО: Полная подстановка переменных окружения перед валидацией
    cfg.Ethereum.PrivateKey = os.ExpandEnv(cfg.Ethereum.PrivateKey)
    cfg.Polygon.PrivateKey = os.ExpandEnv(cfg.Polygon.PrivateKey)
    cfg.Database.Password = os.ExpandEnv(cfg.Database.Password)
    
    // Устанавливаем значения по умолчанию
    cfg.setDefaults()
    
    // СТРОГАЯ валидация после подстановки
    if err := cfg.Validate(); err != nil {
        return nil, fmt.Errorf("config validation failed: %v", err)
    }
    
    log.Printf("✅ Config loaded for environment: %s", env)
    return &cfg, nil
}

// setDefaults устанавливает значения по умолчанию
func (c *Config) setDefaults() {
    if c.Database.SSLMode == "" {
        c.Database.SSLMode = "disable"
    }
    if c.Database.MaxConns == 0 {
        c.Database.MaxConns = 25
    }
    if c.Ethereum.ChainID == 0 {
        c.Ethereum.ChainID = 11155111
    }
    if c.Ethereum.BlockConfirmations == 0 {
        c.Ethereum.BlockConfirmations = 3
    }
    if c.Polygon.ChainID == 0 {
        c.Polygon.ChainID = 80002
    }
    if c.Polygon.GasLimit == 0 {
        c.Polygon.GasLimit = 500000
    }
    if c.Processor.MaxRetries == 0 {
        c.Processor.MaxRetries = 3
    }
    if c.Processor.RetryDelay == 0 {
        c.Processor.RetryDelay = 5 * time.Second
    }
    if c.Processor.BatchSize == 0 {
        c.Processor.BatchSize = 10
    }
    if c.Processor.QueueSize == 0 {
        c.Processor.QueueSize = 100
    }
    if c.Processor.HealthCheckPeriod == 0 {
        c.Processor.HealthCheckPeriod = 30 * time.Second
    }
    if c.Monitoring.MetricsPort == 0 {
        c.Monitoring.MetricsPort = 9090
    }
    if c.Monitoring.LogLevel == "" {
        c.Monitoring.LogLevel = "info"
    }
    if c.Monitoring.LogFormat == "" {
        c.Monitoring.LogFormat = "text"
    }
    if c.API.Port == 0 {
        c.API.Port = 8080
    }
    if c.API.RateLimit == 0 {
        c.API.RateLimit = 100
    }
    
    // Сохраняем обратную совместимость
    if c.Service.LogLevel != "" && c.Monitoring.LogLevel == "info" {
        c.Monitoring.LogLevel = c.Service.LogLevel
    }
}

// Validate выполняет строгую валидацию конфигурации
func (c *Config) Validate() error {
    // Проверка обязательных полей
    if c.Environment == "" {
        return fmt.Errorf("environment is required")
    }
    
    if c.Database.Host == "" {
        return fmt.Errorf("database host is required")
    }
    
    if c.Ethereum.RPCURL == "" {
        return fmt.Errorf("ethereum rpc_url is required")
    }
    
    if c.Polygon.RPCURL == "" {
        return fmt.Errorf("polygon rpc_url is required")
    }
    
    // ЖЕСТКАЯ проверка: если остались ${}, значит переменные не подставились
    if strings.Contains(c.Ethereum.PrivateKey, "${") {
        return fmt.Errorf("ethereum private key contains unsubstituted environment variable: %s", c.Ethereum.PrivateKey)
    }
    
    if strings.Contains(c.Polygon.PrivateKey, "${") {
        return fmt.Errorf("polygon private key contains unsubstituted environment variable: %s", c.Polygon.PrivateKey)
    }
    
    // СТРОГАЯ проверка формата приватных ключей
    if !isValidPrivateKey(c.Ethereum.PrivateKey) {
        return fmt.Errorf("invalid ethereum private key format: must be 64 hex characters, got %d", len(c.Ethereum.PrivateKey))
    }
    
    if !isValidPrivateKey(c.Polygon.PrivateKey) {
        return fmt.Errorf("invalid polygon private key format: must be 64 hex characters, got %d", len(c.Polygon.PrivateKey))
    }
    
    return nil
}

// isValidPrivateKey проверяет валидность приватного ключа
func isValidPrivateKey(key string) bool {
    if key == "" {
        return false
    }
    
    cleanKey := strings.TrimPrefix(key, "0x")
    if len(cleanKey) != 64 {
        return false
    }
    
    for _, c := range cleanKey {
        if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
            return false
        }
    }
    return true
}

// GetDatabaseDSN возвращает DSN строку для PostgreSQL
func (c *DatabaseConfig) GetDSN() string {
    return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        c.Host, c.Port, c.User, c.Password, c.Name, c.SSLMode)
}

// IsDevelopment проверяет dev окружение
func (c *Config) IsDevelopment() bool {
    return c.Environment == "dev"
}

// IsStaging проверяет staging окружение
func (c *Config) IsStaging() bool {
    return c.Environment == "staging"
}

// IsProduction проверяет production окружение
func (c *Config) IsProduction() bool {
    return c.Environment == "production"
}

// GetLogLevel возвращает уровень логирования
func (c *Config) GetLogLevel() string {
    if c.Monitoring.LogLevel != "" {
        return c.Monitoring.LogLevel
    }
    if c.Service.LogLevel != "" {
        return c.Service.LogLevel
    }
    return "info"
}
