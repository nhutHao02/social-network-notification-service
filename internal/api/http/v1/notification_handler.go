package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nhutHao02/social-network-common-service/request"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-common-service/utils/token"
	"github.com/nhutHao02/social-network-notification-service/internal/application"
	"github.com/nhutHao02/social-network-notification-service/internal/domain/model"
	"github.com/nhutHao02/social-network-notification-service/pkg/common"
	"github.com/nhutHao02/social-network-notification-service/pkg/constants"
	"github.com/nhutHao02/social-network-notification-service/pkg/websocket"
	"go.uber.org/zap"
)

type NotificationHandler struct {
	notifService application.NotificationService
}

func NewNotificationHandler(notifService application.NotificationService) *NotificationHandler {
	return &NotificationHandler{notifService: notifService}
}

// NotificationWSHandler godoc
//
//	@Summary		NotificationWSHandler
//	@Description	NotificationWSHandler send messages to user.
//	@Tags			Notification
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer <your_token>"
//	@Param			userID			query		int							true	"User ID"
//	@Success		101				{string}	string						"WebSocket connection established"
//	@Failure		default			{object}	common.Response{data=nil}	"failure"
//	@Router			/ws/notification [get]
func (h *NotificationHandler) NotificationWSHandler(c *gin.Context) {
	var req model.NotifWSReq

	if err := request.GetQueryParamsFromUrl(c, &req); err != nil {
		return
	}

	userID, err := token.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error(), constants.ConnectNotificationWSFailure))
		return
	}

	if userID != int(req.UserID) {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(constants.InvalidUserID, constants.ConnectNotificationWSFailure))
		return
	}

	token, err := token.GetTokenString(c)
	if err != nil {
		logger.Error("ChatHandler-MessageWebSocketHandler: get token from request error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error(), constants.ConnectNotificationWSFailure))
		return
	}

	req.Token = token

	// Upgrade HTTP connection to WebSocket
	conn, err := websocket.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("Error when upgrade HTTP connection to WebSocket", zap.Error(err))
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error(), constants.ConnectNotificationWSFailure))
		return
	}

	h.notifService.NotificationWS(c.Request.Context(), conn, req)
}
