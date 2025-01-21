package game

import (
	"fmt"
	"time"
)

type GameLoop struct {
	Table *Table
	Quit  chan bool
}

func NewGameLoop(table *Table) *GameLoop {
	return &GameLoop{
		Table: table,
		Quit:  make(chan bool),
	}
}

func (g *GameLoop) Start() {
	fmt.Println("Game loop started")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-g.Quit:
			fmt.Println("Game loop stopped")
			return
		case <-ticker.C:
			g.ExecuteTurn()
		}
	}
}

func (g *GameLoop) Stop() {
	g.Quit <- true
}

func (g *GameLoop) ExecuteTurn() {
	currentPlayer := &g.Table.Players[g.Table.CurrentPlayer]
	fmt.Printf("It's %s's turn\n", currentPlayer.ID)

	// Simulate player action (e.g., play a card or call bluff)
	if len(currentPlayer.Hand) > 0 {
		playedCard := currentPlayer.Hand[0]
		err := g.Table.PlayCard(currentPlayer.ID, playedCard, playedCard.Value)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}

		fmt.Printf("%s played %s of %s\n", currentPlayer.ID, playedCard.Value, playedCard.Suit)
	} else {
		fmt.Printf("%s has no cards left\n", currentPlayer.ID)
	}

	// Check for game end condition
	if g.CheckVictory() {
		fmt.Printf("%s wins the game!\n", currentPlayer.ID)
		g.Stop()
	}
}

func (g *GameLoop) CheckVictory() bool {
	for _, player := range g.Table.Players {
		if len(player.Hand) == 0 {
			return true
		}
	}
	return false
}
