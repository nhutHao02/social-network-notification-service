package application

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/model"
)

type NotificationService interface {
	NotificationWS(ctx context.Context, conn *websocket.Conn, req model.NotifWSReq)
}
