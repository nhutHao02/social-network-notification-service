package notification

import (
	"context"

	"github.com/nhutHao02/social-network-notification-service/internal/domain/entity"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/model"
)

type NotificationQueryRepository interface {
	GetNotifByUserID(ctx context.Context, req model.GetNotifByUserIDReq) ([]model.GetNotifByUserIDRes, uint64, error)
}

type NotificationCommandRepository interface {
	SaveNotificaion(ctx context.Context, entityModel entity.Notification) (string, error)
}
