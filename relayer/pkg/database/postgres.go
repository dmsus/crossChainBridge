package database

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/lib/pq"
)

// Config конфигурация для PostgreSQL
type Config struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

// DefaultConfig возвращает конфигурацию по умолчанию
func DefaultConfig() Config {
    return Config{
        Host:     "localhost",
        Port:     "5432",
        User:     "bridge_user",
        Password: "bridge_password",
        DBName:   "bridge_local",
        SSLMode:  "disable",
    }
}

// NewPostgresDB создает новое подключение к PostgreSQL
func NewPostgresDB(cfg Config) (*sql.DB, error) {
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %v", err)
    }

    // Настраиваем connection pool
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)

    // Проверяем подключение
    if err := db.Ping(); err != nil {
        db.Close()
        return nil, fmt.Errorf("failed to ping database: %v", err)
    }

    log.Println("✅ PostgreSQL connection established")

    return db, nil
}

// SetupDatabase создает и возвращает репозиторий
func SetupDatabase(cfg Config) (*Repository, error) {
    db, err := NewPostgresDB(cfg)
    if err != nil {
        return nil, err
    }

    return NewRepository(db), nil
}
