// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/nhutHao02/social-network-common-service/rabbitmq"
	"github.com/nhutHao02/social-network-notification-service/config"
	"github.com/nhutHao02/social-network-notification-service/database"
	"github.com/nhutHao02/social-network-notification-service/internal/api"
	"github.com/nhutHao02/social-network-notification-service/internal/api/http"
	"github.com/nhutHao02/social-network-notification-service/internal/api/http/v1"
	"github.com/nhutHao02/social-network-notification-service/internal/application/imp"
	"github.com/nhutHao02/social-network-notification-service/internal/infrastructure/notification"
	"github.com/nhutHao02/social-network-notification-service/pkg/redis"
	"github.com/nhutHao02/social-network-notification-service/pkg/websocket"
	"github.com/nhutHao02/social-network-user-service/pkg/grpc"
)

// Injectors from wire.go:

func InitializeServer(cfg *config.Config, db *database.MongoDbClient, rdb *redis.RedisClient, userClient grpc.UserServiceClient, ws *websocket.Socket, rabbitmq2 *rabbitmq.RabbitMQ) *api.Server {
	notificationQueryRepository := notification.NewNotificationQueryRepository(db, cfg)
	notificationCommandRepository := notification.NewNotificationCommandRepository(db, cfg)
	notificationService := imp.NewNotificationService(notificationQueryRepository, notificationCommandRepository, userClient, ws, rabbitmq2)
	notificationHandler := v1.NewNotificationHandler(notificationService)
	httpServer := http.NewHTTPServer(cfg, notificationHandler)
	server := api.NewSerVer(httpServer)
	return server
}

// wire.go:

var serverSet = wire.NewSet(api.NewSerVer)

var itemServerSet = wire.NewSet(http.NewHTTPServer)

var httpHandlerSet = wire.NewSet(v1.NewNotificationHandler)

var serviceSet = wire.NewSet(imp.NewNotificationService)

var repositorySet = wire.NewSet(notification.NewNotificationCommandRepository, notification.NewNotificationQueryRepository)
