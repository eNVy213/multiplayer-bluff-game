package network

import (
	"log"

	"github.com/gorilla/websocket"
)

type Connection struct {
	Conn *websocket.Conn
	Send chan []byte
}

func NewConnection(conn *websocket.Conn) *Connection {
	return &Connection{
		Conn: conn,
		Send: make(chan []byte),
	}
}

func (c *Connection) ReadPump(handleMessage func([]byte)) {
	defer func() {
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		handleMessage(message)
	}
}

func (c *Connection) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
