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

func (t *Table) PlayCard(playerID string, card Card) error {
	for _, player := range t.Players {
		if player.ID == playerID {
			for i, c := range player.Hand {
				if c == card {
					// Remove the card from the player's hand
					player.Hand = append(player.Hand[:i], player.Hand[i+1:]...)
					return nil
				}
			}
			return errors.New("card not found in player's hand")
		}
	}
	return errors.New("player not found")
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
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	for _, player := range t.Players {
		if player.Connection != nil {
			player.Connection.Send <- message
		}
	}
}
