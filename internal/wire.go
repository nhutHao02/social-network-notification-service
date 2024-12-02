//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/nhutHao02/social-network-notification-service/config"
	"github.com/nhutHao02/social-network-notification-service/database"
	"github.com/nhutHao02/social-network-notification-service/internal/api"
	"github.com/nhutHao02/social-network-notification-service/internal/api/http"
	"github.com/nhutHao02/social-network-notification-service/internal/api/http/v1"
	"github.com/nhutHao02/social-network-notification-service/internal/application/imp"
	"github.com/nhutHao02/social-network-notification-service/internal/infrastructure/notification"
	"github.com/nhutHao02/social-network-notification-service/pkg/redis"
	ws "github.com/nhutHao02/social-network-notification-service/pkg/websocket"
	grpcUser "github.com/nhutHao02/social-network-user-service/pkg/grpc"
)

var serverSet = wire.NewSet(
	api.NewSerVer,
)

var itemServerSet = wire.NewSet(
	http.NewHTTPServer,
)

var httpHandlerSet = wire.NewSet(
	v1.NewNotificationHandler,
)

var serviceSet = wire.NewSet(
	imp.NewNotificationService,
)

var repositorySet = wire.NewSet(
	notification.NewNotificationCommandRepository,
	notification.NewNotificationQueryRepository,
)

func InitializeServer(
	cfg *config.Config,
	db *database.MongoDbClient,
	rdb *redis.RedisClient,
	userClient grpcUser.UserServiceClient,
	ws *ws.Socket,
) *api.Server {
	wire.Build(serverSet, itemServerSet, httpHandlerSet, serviceSet, repositorySet)
	return &api.Server{}
}
