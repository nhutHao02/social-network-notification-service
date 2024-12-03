package model

import (
	"time"

	"github.com/nhutHao02/social-network-notification-service/pkg/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type GetNotifByUserIDReq struct {
	UserID int64 `form:"userID"`
	Page   int64 `form:"page"`
	Limit  int64 `form:"limit"`
	Token  string
}

type GetNotifByUserIDRes struct {
	ID         primitive.ObjectID    `json:"id" bson:"_id,omitempty"`
	UserID     int64                 `json:"userID" bson:"user_id"`
	AuthorID   int64                 `json:"authorID" bson:"author_id"`
	AuthorInfo *UserInfo             `json:"authorInfo"`
	Message    string                `json:"message" bson:"message"`
	Type       constants.ActionTweet `json:"type" bson:"type"`
	CreatedAt  time.Time             `json:"createdAt" bson:"created_at"`
}
