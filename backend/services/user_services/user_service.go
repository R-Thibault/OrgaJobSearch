package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	userRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/user_repository"
	"github.com/R-Thibault/OrgaJobSearch/backend/utils"
	hashingUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/hash_util"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepo     userRepository.UserRepositoryInterface
	hashingUtils hashingUtils.HashingServiceInterface
}

func NewUserService(UserRepo userRepository.UserRepositoryInterface, hashingUtils hashingUtils.HashingServiceInterface) *UserService {
	return &UserService{UserRepo: UserRepo, hashingUtils: hashingUtils}
}

var _ UserServiceInterface = &UserService{}

func (s *UserService) RegisterUser(creds models.Credentials) error {
	//Check if a user with same email exists
	existingUser, _ := s.UserRepo.GetUserByEmail(creds.Email)
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
	return s.UserRepo.SaveUser(user)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.UserRepo.GetUserByEmail(email)
}

func (s *UserService) EmailValidation(email string) error {
	return s.UserRepo.ValidateEmail(email)
}

func (s *UserService) PreRegisterUser(email string, careerSuportID *uint) (*models.User, error) {
	existingUser, _ := s.UserRepo.GetUserByEmail(email)

	if existingUser != nil {
		return nil, errors.New("user already exists")
	}
	log.Printf("careerS ID: %v", careerSuportID)
	// Optionally, verify if the career support user exists
	if careerSuportID != nil {
		careerSupport, err := s.UserRepo.GetUserByID(*careerSuportID)
		if err != nil {
			return nil, fmt.Errorf("failed to find career support: %w", err)
		}
		if careerSupport == nil {
			return nil, errors.New("career support user does not exist")
		}
	}

	user := models.User{
		Email:           email,
		UserUUID:        uuid.New().String(),
		CareerSupportID: careerSuportID,
	}
	savedUser, err := s.UserRepo.PreRegisterUser(user)
	if err != nil {
		return nil, errors.New("Error during user pre-registration")
	}
	return savedUser, nil
}
