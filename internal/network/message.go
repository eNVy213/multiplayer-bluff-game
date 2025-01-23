package network

type MessageType string

const (
	PlayerAction MessageType = "PlayerAction"
	GameUpdate   MessageType = "GameUpdate"
	PlayerJoined MessageType = "PlayerJoined"
	PlayerLeft   MessageType = "PlayerLeft"
)

type Message struct {
	Type MessageType `json:"type"`
	Data interface{} `json:"data"`
}

type PlayerActionPayload struct {
	PlayerID string `json:"player_id"`
	Action   string `json:"action"`
	Card     string `json:"card,omitempty"`
}

type GameUpdatePayload struct {
	State string `json:"state"`
	Data  string `json:"data"`
}
