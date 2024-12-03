package model

import (
	"time"

	"github.com/nhutHao02/social-network-notification-service/pkg/constants"
)

type Notification struct {
	ID         string                `json:"id"`
	UserID     int64                 `json:"userId"`
	AuthorID   int64                 `json:"authorId"`
	AuthorInfo *UserInfo             `json:"authorInfo"`
	Message    string                `json:"message"`
	Type       constants.ActionTweet `json:"type"`
	CreatedAt  time.Time             `json:"createdAt"`
}
