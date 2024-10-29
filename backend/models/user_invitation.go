package models

type UserInvitation struct {
	Email          string `json:"email"`
	InvitationType string `json:"invitationType"`
}

type GlobalInvitation struct {
	InvitationType string `json:"invitationType"`
}
