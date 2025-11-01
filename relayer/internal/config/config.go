package config

import (
    "fmt"
    "log"

    "github.com/spf13/viper"
)

type EthereumConfig struct {
    RPCURL     string `mapstructure:"rpc_url"`
    WsURL      string `mapstructure:"ws_url"`
    BridgeAddr string `mapstructure:"bridge_addr"`
    PrivateKey string `mapstructure:"private_key"`
}

type PolygonConfig struct {
    RPCURL     string `mapstructure:"rpc_url"`
    WsURL      string `mapstructure:"ws_url"`
    BridgeAddr string `mapstructure:"bridge_addr"`
    PrivateKey string `mapstructure:"private_key"`
}

type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     string `mapstructure:"port"`
    Name     string `mapstructure:"name"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
}

type ServiceConfig struct {
    LogLevel    string `mapstructure:"log_level"`
    MetricsPort string `mapstructure:"metrics_port"`
    APIPort     string `mapstructure:"api_port"`
}

type Config struct {
    Ethereum EthereumConfig `mapstructure:"ethereum"`
    Polygon  PolygonConfig  `mapstructure:"polygon"`
    Database DatabaseConfig `mapstructure:"database"`
    Service  ServiceConfig  `mapstructure:"service"`
}

func Load(env string) (*Config, error) {
    v := viper.New()
    
    v.SetConfigName(".env." + env)
    v.SetConfigType("yaml")
    v.AddConfigPath(".")
    
    if err := v.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("failed to read config file: %v", err)
    }
    
    var config Config
    if err := v.Unmarshal(&config); err != nil {
        return nil, fmt.Errorf("failed to unmarshal config: %v", err)
    }
    
    log.Printf("✅ Config loaded for environment: %s", env)
    return &config, nil
}
