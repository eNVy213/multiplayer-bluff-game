package test

import (
	"testing"

	"github.com/eNVy213/multiplayer-bluff-game/internal/game"
)

func TestBluffGame(t *testing.T) {
	gameInstance := game.NewBluffGame()

	if len(gameInstance.Deck) != 52 {
		t.Errorf("Expected deck size 52, got %d", len(gameInstance.Deck))
	}

	player := &game.Player{ID: "player1", Name: "Test Player"}
	err := gameInstance.AddPlayer(player)
	if err != nil {
		t.Errorf("Failed to add player: %v", err)
	}

	if len(gameInstance.Players) != 1 {
		t.Errorf("Expected 1 player, got %d", len(gameInstance.Players))
	}
}
