package services

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

type UserServiceInterface interface {
	RegisterUser(creds models.Credentials) error
	GetUserByEmail(email string) (*models.User, error)
	EmailValidation(email string) error
}
