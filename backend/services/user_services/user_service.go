package services

import (
	"errors"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	userRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/user_repository"
	hashingUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/hash_util"
)

type UserService struct {
	UserRepo     userRepository.UserRepositoryInterface
	hashingUtils hashingUtils.HashingServiceInterface
}

func NewUserService(UserRepo userRepository.UserRepositoryInterface, hashingUtils hashingUtils.HashingServiceInterface) *UserService {
	return &UserService{UserRepo: UserRepo, hashingUtils: hashingUtils}
}

var _ UserServiceInterface = &UserService{}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.UserRepo.GetUserByEmail(email)
}

func (s *UserService) EmailValidation(email string) error {
	return s.UserRepo.ValidateEmail(email)
}

func (s *UserService) GetUserByID(userID uint) (*models.User, error) {
	return s.UserRepo.GetUserByID(userID)
}

func (s *UserService) GetUserByUUID(userUUID string) (*models.User, error) {
	return s.UserRepo.GetUserByUUID(userUUID)
}

func (s *UserService) UpdateUser(existingUser models.User, updatedUserDatas models.UserProfileUpdate) error {
	if updatedUserDatas.UserLastName == "" || updatedUserDatas.UserFirstName == "" {
		return errors.New("UserName can't be empty")
	}
	if updatedUserDatas.Email == "" {
		return errors.New("Email can't be empty")
	}
	if existingUser.Email != updatedUserDatas.Email {
		return errors.New("Emails do not match")
	}
	return s.UserRepo.UpdateUser(existingUser.ID, updatedUserDatas)
}
