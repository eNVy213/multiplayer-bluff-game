package network

import (
	"log"

	"github.com/gorilla/websocket"
)

type Connection struct {
	Conn *websocket.Conn
	Send chan Message
}

func NewConnection(conn *websocket.Conn) *Connection {
	return &Connection{
		Conn: conn,
		Send: make(chan Message, 256),
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
	for message := range c.Send {
		if err := c.Conn.WriteJSON(message); err != nil {
			break
		}
	}
	c.Conn.Close()
}
