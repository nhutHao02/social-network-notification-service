package notification

import (
	"github.com/nhutHao02/social-network-notification-service/config"
	"github.com/nhutHao02/social-network-notification-service/database"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/interface/notification"
)

type notificationCommandRepository struct {
	cfg *config.Config
	db  *database.MongoDbClient
}

// Command implements notification.NotificationCommandRepository.
func (n *notificationCommandRepository) Command() {
	panic("unimplemented")
}

func NewNotificationCommandRepository(db *database.MongoDbClient, cfg *config.Config) notification.NotificationCommandRepository {
	return &notificationCommandRepository{db: db, cfg: cfg}
}
