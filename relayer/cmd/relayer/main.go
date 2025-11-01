package main

import (
	"log"
	"os"

	"github.com/dmsus/crossChainBridge/relayer/internal/config"
)

func main() {
	env := "local"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("🚀 Starting Cross-Chain Bridge Relayer in %s mode", env)
	log.Printf("📡 Ethereum RPC: %s", cfg.Ethereum.RPCURL)
	log.Printf("📡 Polygon RPC: %s", cfg.Polygon.RPCURL)
	log.Printf("🗄️  Database: %s@%s:%s", cfg.Database.User, cfg.Database.Host, cfg.Database.Port)
	log.Printf("📊 Metrics port: %s", cfg.Service.MetricsPort)
	
	// TODO: Initialize and start services
	log.Println("✅ Relayer initialized successfully - ready for development!")
	
	// Block main thread
	select {}
}
