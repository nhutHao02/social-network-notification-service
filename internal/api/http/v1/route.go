package v1

import (
	"net/http"

	_ "github.com/nhutHao02/social-network-notification-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MapRoutes(
	router *gin.Engine,
	notifHandler *NotificationHandler,
) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		// v1.Use(middleware.JwtAuthMiddleware(logger.GetDefaultLogger()))
		// {
		// 	vChat := v1.Group("/notif")

		// 	vSocket := v1.Group("/ws")
		// 	vSocket.GET("", )
		// }

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
