package services

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

type UserServiceInterface interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(userID uint) (*models.User, error)
	EmailValidation(email string) error
	GetUserByUUID(userUUID string) (*models.User, error)
	UpdateUser(existingUser models.User, updatedUserDatas models.UserProfileUpdate) error
	ResetPassword(user models.User, claims models.JWTToken, newPassword string) error
}
