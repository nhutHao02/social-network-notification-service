package v1

import (
	"github.com/nhutHao02/social-network-notification-service/internal/application"
)

type NotificationHandler struct {
	notifService application.NotificationService
}

func NewNotificationHandler(notifService application.NotificationService) *NotificationHandler {
	return &NotificationHandler{notifService: notifService}
}
