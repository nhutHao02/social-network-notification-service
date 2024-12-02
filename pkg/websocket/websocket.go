package websocket

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/model"
	"go.uber.org/zap"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Socket struct {
	//          map[userID]*websocket.Conn
	connections map[string]*websocket.Conn
	mu          sync.RWMutex
}

func NewSocket() *Socket {
	return &Socket{
		connections: make(map[string]*websocket.Conn),
	}
}

// Add connection
func (s *Socket) AddConnection(userID string, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.connections[userID] = conn
}

// Remove connection
func (s *Socket) RemoveConnection(userID string, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.connections, userID)
	conn.Close()
}

// Broadcast message to all connections
func (s *Socket) Broadcast(userID string, message model.OutgoingMessageWSRes) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for uid, conn := range s.connections {
		if uid != userID {
			if err := conn.WriteJSON(message); err != nil {
				logger.Error("Socket-Broadcast: Error sending message", zap.Error(err))
				s.RemoveConnection(userID, conn)
			}
		}
	}
}
