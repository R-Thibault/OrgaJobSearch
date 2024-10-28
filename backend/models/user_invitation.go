package models

type UserInvitation struct {
	Email          string `json:"email"`
	InvitationType string `json:"invitationType"`
}

type GlobalInvitation struct {
	UserID         uint   `json:"userID"`
	InvitationType string `json:"invitationType"`
}
