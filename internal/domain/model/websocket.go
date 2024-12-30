package model

import "time"

type NotifWSReq struct {
	UserID int64  `form:"userID"`
	Token  string `form:"token"`
}

type IncomingMessageWSReq struct {
	Message string `json:"message"`
}

type OutgoingMessageWSRes struct {
	ID string `json:"id"`
	// Sender    *UserInfo `json:"sender"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}
