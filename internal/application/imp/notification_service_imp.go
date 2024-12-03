package imp

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/nhutHao02/social-network-common-service/rabbitmq"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-notification-service/internal/application"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/entity"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/interface/notification"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/model"
	"github.com/nhutHao02/social-network-notification-service/pkg/constants"
	ws "github.com/nhutHao02/social-network-notification-service/pkg/websocket"
	grpcUser "github.com/nhutHao02/social-network-user-service/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type notificationService struct {
	queryRepo   notification.NotificationQueryRepository
	commandRepo notification.NotificationCommandRepository
	userClient  grpcUser.UserServiceClient
	ws          *ws.Socket
	rabbitmq    *rabbitmq.RabbitMQ
}

// NotificationWS implements application.NotificationService.
func (n *notificationService) NotificationWS(ctx context.Context, conn *websocket.Conn, req model.NotifWSReq) {
	userIDKey := strconv.Itoa(int(req.UserID))
	// Add connection
	n.ws.AddConnection(userIDKey, conn)

	// Listen for connection close event
	defer n.ws.RemoveConnection(userIDKey, conn)

	// consume notificaiton
	msgs, err := n.rabbitmq.ConsumeMessages()
	if err != nil {
		logger.Error("Failed to register a consumer: %s", zap.Error(err))
		return
	}

	for msg := range msgs {
		var notification model.Notification
		if err := json.Unmarshal(msg.Body, &notification); err != nil {
			logger.Error("Error unmarshalling message from notification", zap.Error(err))
			continue
		}
		// get message
		notification.Message = notification.Type.Message()

		// save to db
		entityModel := entity.Notification{
			UserID:    notification.UserID,
			AuthorID:  notification.AuthorID,
			Message:   notification.Message,
			Type:      notification.Type,
			CreatedAt: notification.CreatedAt,
		}
		notifID, err := n.commandRepo.SaveNotificaion(ctx, entityModel)

		// pass notificationID to notification
		notification.ID = notifID

		// get author info
		// Create context with metadata
		md := metadata.Pairs("authorization", constants.BearerString+req.Token)
		ctxx := metadata.NewOutgoingContext(ctx, md)

		authorInfo, err := n.userClient.GetUserInfo(ctxx, &grpcUser.GetUserRequest{UserID: notification.AuthorID})
		if err != nil {
			logger.Error("tweetService-CommentWebSocket: Error get UserInfo, call grpcUser to server error", zap.Error(err))
		}

		// pass author info to notification
		notification.AuthorInfo = &model.UserInfo{
			ID:       &authorInfo.Id,
			Email:    &authorInfo.Email,
			FullName: &authorInfo.FullName,
			UrlAvt:   &authorInfo.UrlAvt,
		}

		// broadcast to connection
		n.ws.Broadcast(userIDKey, notification)
	}
}

func NewNotificationService(
	queryRepo notification.NotificationQueryRepository,
	commandRepo notification.NotificationCommandRepository,
	userClient grpcUser.UserServiceClient,
	ws *ws.Socket,
	rabbitmq *rabbitmq.RabbitMQ,
) application.NotificationService {
	return &notificationService{queryRepo: queryRepo, commandRepo: commandRepo, userClient: userClient, ws: ws, rabbitmq: rabbitmq}
}
