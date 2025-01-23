package game

import (
	"sync"

	"github.com/your_project/internal/network"
)

type Table struct {
	ID         string
	Players    []*Player
	Mutex      sync.Mutex
	Broadcast  chan network.Message
	MaxPlayers int
	GameState  *BluffGame
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
		Type:    network.PlayerJoined,
		Payload: map[string]string{"player_id": player.ID},
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
				Type:    network.PlayerLeft,
				Payload: map[string]string{"player_id": playerID},
			}
			return
		}
	}
}

func (t *Table) BroadcastMessage(message network.Message) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	for _, player := range t.Players {
		player.Connection.Send <- message
	}
}
