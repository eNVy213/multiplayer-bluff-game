package game

import (
	"testing"
)

func TestPlayCard(t *testing.T) {
	players := []*Player{
		{ID: "1", Name: "Alice", Hand: []string{"Ace of Spades", "2 of Hearts"}},
		{ID: "2", Name: "Bob", Hand: []string{"King of Hearts"}},
	}

	table := &Table{
		Players:       players,
		CurrentPlayer: players[0],
	}

	err := table.PlayCard("Ace of Spades")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(table.CurrentPlayer.Hand) != 1 {
		t.Fatalf("Expected 1 card in Alice's hand, got %d", len(table.CurrentPlayer.Hand))
	}

	if table.CurrentPlayer != players[1] {
		t.Fatalf("Expected current player to be Bob, got %v", table.CurrentPlayer.Name)
	}
}

func TestPlayCard_NoCurrentPlayer(t *testing.T) {
	table := &Table{}

	err := table.PlayCard("Ace of Spades")
	if err == nil || err.Error() != "no current player set" {
		t.Fatalf("Expected 'no current player set' error, got %v", err)
	}
}

func TestPlayCard_CardNotInHand(t *testing.T) {
	players := []*Player{
		{ID: "1", Name: "Alice", Hand: []string{"2 of Hearts"}},
	}

	table := &Table{
		Players:       players,
		CurrentPlayer: players[0],
	}

	err := table.PlayCard("Ace of Spades")
	if err == nil || err.Error() != "card not in hand" {
		t.Fatalf("Expected 'card not in hand' error, got %v", err)
	}
}
