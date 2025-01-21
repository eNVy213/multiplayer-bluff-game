package test

import (
	"multiplayer-bluff-game/internal/network"
	"net/http/httptest"
	"testing"
)

func TestWebSocketConnection(t *testing.T) {
	server := network.NewServer(":8080")

	go func() {
		if err := server.Start(); err != nil {
			t.Fatalf("Server failed to start: %v", err)
		}
	}()

	defer server.Stop()

	ts := httptest.NewServer(server.Router)
	defer ts.Close()

	wsURL := "ws" + ts.URL[4:] + "/ws"
	_, _, err := network.ConnectToWebSocket(wsURL)
	if err != nil {
		t.Errorf("Failed to establish WebSocket connection: %v", err)
	}
}
