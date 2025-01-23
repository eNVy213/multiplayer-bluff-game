package models

// Message represents a message exchanged over WebSocket.
type Message struct {
	Type    string      `json:"type"`    // Type of message (e.g., "join", "move", "chat")
	Content interface{} `json:"content"` // Actual content of the message
	Sender  string      `json:"sender"`  // ID of the sender
}

// JoinTableRequest represents a request to join a table.
type JoinTableRequest struct {
	PlayerID string `json:"player_id"` // ID of the player requesting to join
	TableID  string `json:"table_id"`  // ID of the table to join
}

// GameAction represents an in-game action by a player.
type GameAction struct {
	PlayerID string `json:"player_id"` // ID of the player performing the action
	Action   string `json:"action"`    // Action type (e.g., "play_card", "bluff")
	Data     string `json:"data"`      // Additional data for the action
}

// ChatMessage represents a chat message sent by a player.
type ChatMessage struct {
	PlayerID string `json:"player_id"` // ID of the player sending the message
	Message  string `json:"message"`   // Chat message content
}
