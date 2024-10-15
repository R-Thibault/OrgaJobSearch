package services

import "github.com/R-Thibault/OrgaJobSearch/models"

type UserServiceInterface interface {
	RegisterUser(creds models.Credentials) error
	GetUserByEmail(email string) (*models.User, error)
}
