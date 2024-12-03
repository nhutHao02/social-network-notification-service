package notification

import (
	"github.com/nhutHao02/social-network-notification-service/config"
	"github.com/nhutHao02/social-network-notification-service/database"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/interface/notification"
)

type notificationQueryRepository struct {
	cfg *config.Config
	db  *database.MongoDbClient
}

func NewNotificationQueryRepository(db *database.MongoDbClient, cfg *config.Config) notification.NotificationQueryRepository {
	return &notificationQueryRepository{db: db, cfg: cfg}
}
