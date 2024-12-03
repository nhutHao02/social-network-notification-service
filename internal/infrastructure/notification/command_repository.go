package notification

import (
	"context"

	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-notification-service/config"
	"github.com/nhutHao02/social-network-notification-service/database"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/entity"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/interface/notification"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type notificationCommandRepository struct {
	cfg *config.Config
	db  *database.MongoDbClient
}

// SaveNotificaion implements notification.NotificationCommandRepository.
func (repo *notificationCommandRepository) SaveNotificaion(ctx context.Context, entityModel entity.Notification) (string, error) {
	documentID, err := repo.db.InsertOne(ctx, repo.cfg.Database.DBName, entity.CollectionNotification, entityModel)
	if err != nil {
		logger.Error("notificationCommandRepository-SaveNotificaion: Error inserting document", zap.Error(err))
		return "", err
	}
	return documentID.(primitive.ObjectID).Hex(), nil
}

func NewNotificationCommandRepository(db *database.MongoDbClient, cfg *config.Config) notification.NotificationCommandRepository {
	return &notificationCommandRepository{db: db, cfg: cfg}
}
