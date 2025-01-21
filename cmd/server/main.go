package main

import (
	"log"
	"multiplayer-bluff-game/config"
	"multiplayer-bluff-game/internal/network"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("./config/config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize WebSocket server
	server := network.NewServer(cfg.ServerAddress)

	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
