package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Ethereum struct {
		RPCURL      string `mapstructure:"rpc_url"`
		WSURL       string `mapstructure:"ws_url"`
		BridgeAddr  string `mapstructure:"bridge_addr"`
		PrivateKey  string `mapstructure:"private_key"`
	} `mapstructure:"ethereum"`
	
	Polygon struct {
		RPCURL      string `mapstructure:"rpc_url"`
		WSURL       string `mapstructure:"ws_url"`
		BridgeAddr  string `mapstructure:"bridge_addr"`
		PrivateKey  string `mapstructure:"private_key"`
	} `mapstructure:"polygon"`
	
	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`
	
	Service struct {
		LogLevel    string `mapstructure:"log_level"`
		MetricsPort string `mapstructure:"metrics_port"`
		APIPort     string `mapstructure:"api_port"`
	} `mapstructure:"service"`
}

func LoadConfig(env string) (*Config, error) {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configFile := ".env." + env + ".yaml"
	fullPath := filepath.Join(wd, configFile)
	
	log.Printf("Loading config from: %s", fullPath)
	
	viper.SetConfigFile(fullPath)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No config file found at %s, error: %v", fullPath, err)
		log.Printf("Using environment variables only")
	} else {
		log.Printf("✅ Config file loaded successfully: %s", fullPath)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
