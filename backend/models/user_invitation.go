package models

type UserInvitation struct {
	Email  string `json:"email"`
	UserID uint   `json:"userID"`
}
