package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	manager  *WebSocketManager
	upgrader websocket.Upgrader
}

func NewWebSocketHandler(manager *WebSocketManager) *WebSocketHandler {
	return &WebSocketHandler{
		manager: manager,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Accept all origins
			},
		},
	}
}

func (h *WebSocketHandler) HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	h.manager.AddClient(conn)
	defer h.manager.RemoveClient(conn)

	// Start reading messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		h.manager.BroadcastMessage(string(message))
	}
}
