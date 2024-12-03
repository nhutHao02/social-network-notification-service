package notification

import (
	"context"

	"github.com/nhutHao02/social-network-notification-service/internal/domain/entity"
)

type NotificationQueryRepository interface {
}

type NotificationCommandRepository interface {
	SaveNotificaion(ctx context.Context, entityModel entity.Notification) (string, error)
}
