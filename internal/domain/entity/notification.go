package entity

import (
	"time"

	"github.com/nhutHao02/social-network-notification-service/pkg/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID        primitive.ObjectID    `bson:"_id,omitempty"`
	UserID    int64                 `bson:"user_id"`
	AuthorID  int64                 `bson:"author_id"`
	Message   string                `bson:"message"`
	Type      constants.ActionTweet `bson:"type"`
	CreatedAt time.Time             `bson:"created_at"`
}
