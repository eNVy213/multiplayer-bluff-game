package main

import (
	"github.com/eNVy213/multiplayer-bluff-game/config"
	"github.com/eNVy213/multiplayer-bluff-game/internal/network"
	"github.com/eNVy213/multiplayer-bluff-game/internal/storage"
	"github.com/eNVy213/multiplayer-bluff-game/internal/utils"
)

func main() {
	// Initialize logger
	utils.InitializeLogger()

	// Load configuration
	cfg, err := config.LoadConfig("./config/config.json")
	if err != nil {
		utils.Logger.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize PostgreSQL
	db, err := storage.NewPostgresDB("localhost", "5432", "user", "password", "bluff_game")
	if err != nil {
		utils.Logger.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	// Initialize WebSocket server
	server := network.NewServer(cfg.ServerAddress)

	utils.Logger.Printf("Starting server on %s", cfg.ServerAddress)
	if err := server.Start(); err != nil {
		utils.Logger.Fatalf("Failed to start server: %v", err)
	}
}
