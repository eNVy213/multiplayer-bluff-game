package game

import (
	"testing"
	"time"
)

func TestServerStart(t *testing.T) {
	server := NewServer("localhost:8080")

	go func() {
		if err := server.Start("localhost:8080"); err != nil {
			t.Fatalf("Failed to start server: %v", err)
		}
	}()

	time.Sleep(1 * time.Second) // Wait for the server to start
}

func TestWebSocketHandler(t *testing.T) {
	server := NewServer("localhost:8080")
	go server.Start("localhost:8080")
	time.Sleep(1 * time.Second)

	// Use a WebSocket client to test the connection
	// Example:
	// ws, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	// if err != nil {
	//     t.Fatalf("Failed to connect to WebSocket server: %v", err)
	// }
	// defer ws.Close()
}
