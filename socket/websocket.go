package socket

import (
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type WebSocketClient struct {
	Conns map[*websocket.Conn]bool
	mu    sync.Mutex
}

func NewWebSocketClient() *WebSocketClient {
	return &WebSocketClient{
		Conns: make(map[*websocket.Conn]bool),
	}
}

func (c *WebSocketClient) Add(conn *websocket.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Conns[conn] = true
}

func (c *WebSocketClient) Remove(conn *websocket.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.Conns, conn)
}

func (c *WebSocketClient) Broadcast(message []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for conn := range c.Conns {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("web socket error: " + err.Error())
			conn.Close()
			delete(c.Conns, conn)
		}
	}
}
