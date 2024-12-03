package notification

import (
	"context"

	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-notification-service/config"
	"github.com/nhutHao02/social-network-notification-service/database"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/entity"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/interface/notification"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type notificationQueryRepository struct {
	cfg *config.Config
	db  *database.MongoDbClient
}

// GetNotifByUserID implements notification.NotificationQueryRepository.
func (repo *notificationQueryRepository) GetNotifByUserID(ctx context.Context, req model.GetNotifByUserIDReq) ([]model.GetNotifByUserIDRes, uint64, error) {
	var res []model.GetNotifByUserIDRes

	filter := bson.M{"user_id": req.UserID}

	opts := options.Find()
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	opts.SetSkip((req.Page - 1) * req.Limit)
	opts.SetLimit(req.Limit)

	err := repo.db.FindMany(ctx, repo.cfg.Database.DBName, entity.CollectionNotification, filter, &res, opts)
	if err != nil {
		logger.Error("notificationQueryRepository-GetNotifByUserID: FindMany message error", zap.Error(err))
		return res, 0, err
	}

	totalCount, err := repo.db.CountDocuments(ctx, repo.cfg.Database.DBName, entity.CollectionNotification, filter)
	if err != nil {
		logger.Error("notificationQueryRepository-GetNotifByUserID: Count Documents error", zap.Error(err))
		return res, 0, err
	}

	return res, uint64(totalCount), nil
}

func NewNotificationQueryRepository(db *database.MongoDbClient, cfg *config.Config) notification.NotificationQueryRepository {
	return &notificationQueryRepository{db: db, cfg: cfg}
}
