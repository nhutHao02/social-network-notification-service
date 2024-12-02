package imp

import (
	"github.com/nhutHao02/social-network-notification-service/internal/application"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/interface/notification"
	ws "github.com/nhutHao02/social-network-notification-service/pkg/websocket"
	grpcUser "github.com/nhutHao02/social-network-user-service/pkg/grpc"
)

type notificationService struct {
	queryRepo   notification.NotificationQueryRepository
	commandRepo notification.NotificationCommandRepository
	userClient  grpcUser.UserServiceClient
	ws          *ws.Socket
}

// NotificationService implements application.NotificationService.
func (n *notificationService) NotificationService() {
	panic("unimplemented")
}

func NewNotificationService(
	queryRepo notification.NotificationQueryRepository,
	commandRepo notification.NotificationCommandRepository,
	userClient grpcUser.UserServiceClient,
	ws *ws.Socket,
) application.NotificationService {
	return &notificationService{queryRepo: queryRepo, commandRepo: commandRepo, userClient: userClient, ws: ws}
}
