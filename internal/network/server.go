package network

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	Upgrader websocket.Upgrader
	Clients  map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		Clients: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()
	s.Clients[conn] = true

	for {
		var msg string
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			delete(s.Clients, conn)
			break
		}
		fmt.Println("Received message:", msg)
	}
}

func (s *Server) BroadcastMessage(message string) {
	for client := range s.Clients {
		err := client.WriteJSON(message)
		if err != nil {
			log.Println("Error broadcasting message:", err)
			client.Close()
			delete(s.Clients, client)
		}
	}
}

func (s *Server) Start(addr string) error {
	http.HandleFunc("/ws", s.HandleConnections)
	log.Println("Server starting on", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}
	return nil
}
