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
		UserUUID:       uuid.New().String(),
		UserStatus:     "pending",
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

func (s *UserService) GetUserByID(userID uint) (*models.User, error) {
	return s.UserRepo.GetUserByID(userID)
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

func (s *UserService) JobSeekerRegistration(tokenBody string, creds models.Credentials) error {
	savedUser, err := s.UserRepo.GetUserByUUID(tokenBody)
	if err != nil {
		return errors.New("UUID doesn't match a user")
	}
	// check password requirement
	isMatch := utils.RegexPassword(creds.Password)
	if !isMatch {
		return errors.New("Password doesn't match requirement")
	}
	hashedPassword, hashErr := s.hashingUtils.HashPassword(creds.Password)
	if hashErr != nil {
		return errors.New("Error during password hash")
	}
	savedUser.HashedPassword = hashedPassword
	savedUser.Email = creds.Email
	savedUser.UserUUID = tokenBody
	updateErr := s.UserRepo.UpdateJobSeeker(*savedUser)
	if updateErr != nil {
		return errors.New("Error during user update")
	}
	return nil
}
