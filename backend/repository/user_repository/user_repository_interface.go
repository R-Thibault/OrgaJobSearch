package repository

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

// UserRepositoryInterface defines the methods for interacting with users in the database.
type UserRepositoryInterface interface {
	SaveUser(user models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(ID uint) (*models.User, error)
	ValidateEmail(email string) error
	PreRegisterUser(user models.User) error
	GetUserByUUID(uuid string) (*models.User, error)
}
