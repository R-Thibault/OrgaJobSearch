package services

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

type UserServiceInterface interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(userID uint) (*models.User, error)
	EmailValidation(email string) error
}
