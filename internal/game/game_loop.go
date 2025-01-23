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
	close(g.Quit)
}

func (g *GameLoop) ExecuteTurn() {
	currentPlayer := g.Table.CurrentPlayer

	if currentPlayer == nil {
		fmt.Println("No current player set")
		return
	}

	fmt.Printf("It's %s's turn\n", currentPlayer.ID)

	// Simulate player action
	if len(currentPlayer.Hand) > 0 {
		playedCard := currentPlayer.Hand[0]
		err := g.Table.PlayCard(currentPlayer.ID, playedCard)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}

		fmt.Printf("%s played %s\n", currentPlayer.ID, playedCard)
	} else {
		fmt.Printf("%s has no cards left\n", currentPlayer.ID)
	}

	// Check for victory condition
	if g.CheckVictory() {
		fmt.Printf("%s wins the game!\n", currentPlayer.ID)
		g.Stop()
	} else {
		g.Table.NextPlayer()
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
