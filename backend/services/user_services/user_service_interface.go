package services

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

type UserServiceInterface interface {
	RegisterUser(creds models.Credentials) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(userID uint) (*models.User, error)
	PreRegisterUser(email string, careerSuportID *uint) (*models.User, error)
	EmailValidation(email string) error
	JobSeekerRegistration(tokenBody string, creds models.Credentials) error
}
