package services

import (
	"errors"

	"github.com/R-Thibault/OrgaJobSearch/models"
	userRepository "github.com/R-Thibault/OrgaJobSearch/repository/user_repository"
	"github.com/R-Thibault/OrgaJobSearch/utils"
	hashingUtils "github.com/R-Thibault/OrgaJobSearch/utils/hash_util"
)

type UserService struct {
	UserRrepo    userRepository.UserRepositoryInterface
	hashingUtils hashingUtils.HashingServiceInterface
}

func NewUserService(UserRrepo userRepository.UserRepositoryInterface, hashingUtils hashingUtils.HashingServiceInterface) *UserService {
	return &UserService{UserRrepo: UserRrepo, hashingUtils: hashingUtils}
}

var _ UserServiceInterface = &UserService{}

func (s *UserService) RegisterUser(creds models.Credentials) error {
	//Check if a user with same email exists
	existingUser, _ := s.UserRrepo.GetUserByEmail(creds.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}
	// check password requirement
	isMatch := utils.RegexPassword(creds.Password)
	if !isMatch {
		return errors.New("Password doesn't match requirement")
	}

	//Hash user's password
	hashedPassword, err := s.hashingUtils.HashPassword(creds.Password)
	if err != nil {
		return err
	}

	// Prepare user object
	user := models.User{
		Email:          creds.Email,
		HashedPassword: string(hashedPassword),
	}
	// Save the user
	return s.UserRrepo.SaveUser(user)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.UserRrepo.GetUserByEmail(email)
}

func (s *UserService) EmailValidation(email string) error {
	return s.UserRrepo.ValidateEmail(email)
}
