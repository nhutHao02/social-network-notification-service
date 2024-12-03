package model

type UserInfo struct {
	ID       *int64  `json:"id"`
	Email    *string `json:"email"`
	FullName *string `json:"fullName"`
	UrlAvt   *string `json:"urlAvt"`
}
