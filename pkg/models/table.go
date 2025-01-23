package models

// Table represents a single game table.
type Table struct {
	ID            string    `json:"id"`        // Unique identifier for the table
	Players       []*Player `json:"players"`   // Players at the table
	IsActive      bool      `json:"is_active"` // Whether the table is active
	Round         int       `json:"round"`     // Current round number
	CurrentPlayer *Player   `json:"current_player"`
}

// NewTable creates a new game table with the given ID.
func NewTable(id string) *Table {
	return &Table{
		ID:       id,
		Players:  make([]*Player, 0, 4), // Maximum of 4 players
		IsActive: false,
		Round:    0,
	}
}

// AddPlayer adds a player to the table.
func (t *Table) AddPlayer(player *Player) bool {
	if len(t.Players) >= 4 {
		return false // Table is full
	}
	t.Players = append(t.Players, player)
	return true
}

// StartGame sets the table as active and begins the game.
func (t *Table) StartGame() {
	t.IsActive = true
	t.Round = 1
}

// NextRound advances the game to the next round.
func (t *Table) NextRound() {
	t.Round++
}
