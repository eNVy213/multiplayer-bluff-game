package game

import (
	"errors"
	"sync"

	"github.com/eNVy213/multiplayer-bluff-game/internal/network"
)

type Table struct {
	ID         string
	Players    []*Player `json:"players"`
	Mutex      sync.Mutex
	Broadcast  chan network.Message
	MaxPlayers int
	GameState  *BluffGame
	//Players       []*Player `json:"players"`
	CurrentPlayer *Player `json:"current_player"`
}

type Table struct {
	Players       []*Player `json:"players"`
	CurrentPlayer int       `json:"current_player"` // Index of the current player
}

func (t *Table) PlayCard(playerID string, card string) error {
	player := t.Players[t.CurrentPlayer]
	if player.ID != playerID {
		return errors.New("not your turn")
	}

	// Check if the player has the card
	cardIndex := -1
	for i, c := range player.Hand {
		if c == card {
			cardIndex = i
			break
		}
	}
	if cardIndex == -1 {
		return errors.New("card not in hand")
	}

	// Remove card from player's hand
	player.Hand = append(player.Hand[:cardIndex], player.Hand[cardIndex+1:]...)

	// Broadcast message
	message := network.Message{
		Type:    "PlayCard",
		Payload: map[string]interface{}{"playerID": playerID, "card": card},
	}
	t.BroadcastMessage(message)

	// Update turn
	t.CurrentPlayer = (t.CurrentPlayer + 1) % len(t.Players)
	return nil
}

func (t *Table) NextPlayer() {
	if len(t.Players) == 0 {
		return
	}

	currentIndex := -1
	for i, player := range t.Players {
		if player == t.CurrentPlayer {
			currentIndex = i
			break
		}
	}

	if currentIndex == -1 || currentIndex+1 >= len(t.Players) {
		t.CurrentPlayer = t.Players[0]
	} else {
		t.CurrentPlayer = t.Players[currentIndex+1]
	}
}

func NewTable(id string, maxPlayers int) *Table {
	return &Table{
		ID:         id,
		Players:    make([]*Player, 0),
		Broadcast:  make(chan network.Message),
		MaxPlayers: maxPlayers,
		GameState:  NewBluffGame(),
	}
}

func (t *Table) AddPlayer(player *Player) bool {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	if len(t.Players) >= t.MaxPlayers {
		return false
	}

	t.Players = append(t.Players, player)
	t.Broadcast <- network.Message{
		Type: network.PlayerJoined,
		Data: map[string]string{"player_id": player.ID},
	}
	return true
}

func (t *Table) RemovePlayer(playerID string) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	for i, p := range t.Players {
		if p.ID == playerID {
			t.Players = append(t.Players[:i], t.Players[i+1:]...)
			t.Broadcast <- network.Message{
				Type: network.PlayerLeft,
				Data: map[string]string{"player_id": playerID},
			}
			return
		}
	}
}

func (t *Table) BroadcastMessage(message network.Message) {
	for _, player := range t.Players {
		if player.Connection != nil {
			player.Connection.Send <- message
		}
	}
}
