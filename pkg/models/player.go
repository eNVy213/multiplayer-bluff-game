package models

// Player represents a single player in the game.
type Player struct {
	ID         string   `json:"id"`          // Unique identifier for the player
	Name       string   `json:"name"`        // Display name of the player
	IsReady    bool     `json:"is_ready"`    // Indicates if the player is ready to play
	Score      int      `json:"score"`       // Player's score in the game
	IsTurn     bool     `json:"is_turn"`     // Whether it's the player's turn
	HasBluffed bool     `json:"has_bluffed"` // Tracks if the player has bluffed
	Hand       []string `json:"hand"`
}

// NewPlayer creates a new player with the provided ID and name.
func NewPlayer(id, name string) *Player {
	return &Player{
		ID:         id,
		Name:       name,
		IsReady:    false,
		Score:      0,
		IsTurn:     false,
		HasBluffed: false,
	}
}

// MarkReady sets the player as ready to play.
func (p *Player) MarkReady() {
	p.IsReady = true
}

// UpdateScore increments the player's score.
func (p *Player) UpdateScore(points int) {
	p.Score += points
}
