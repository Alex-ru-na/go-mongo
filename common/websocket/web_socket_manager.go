package websocket

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketManager struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (m *WebSocketManager) AddClient(conn *websocket.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.clients[conn] = true
	log.Println("New client connected")
}

func (m *WebSocketManager) RemoveClient(conn *websocket.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.clients, conn)
	log.Println("Client disconnected")
}

func (m *WebSocketManager) BroadcastMessage(message string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for client := range m.clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Printf("Error sending message to client: %v", err)
			client.Close()
			delete(m.clients, client)
		}
	}
}
