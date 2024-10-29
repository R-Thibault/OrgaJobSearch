package registrationservices

import (
	"errors"
	"fmt"
	"log"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	roleRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/role_repository"
	userRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/user_repository"
	"github.com/R-Thibault/OrgaJobSearch/backend/utils"
	hashingUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/hash_util"
	"github.com/google/uuid"
)

type RegistrationService struct {
	UserRepo       userRepository.UserRepositoryInterface
	HashingUtils   hashingUtils.HashingServiceInterface
	RoleRepository roleRepository.RoleRepositoryInterface
}

func NewRegistrationService(
	UserRepo userRepository.UserRepositoryInterface,
	HashingUtils hashingUtils.HashingServiceInterface,
	RoleRepository roleRepository.RoleRepositoryInterface) *RegistrationService {
	return &RegistrationService{
		UserRepo:       UserRepo,
		HashingUtils:   HashingUtils,
		RoleRepository: RoleRepository}
}

func (s *RegistrationService) RegisterCareerCoach(creds models.Credentials) error {
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
	hashedPassword, err := s.HashingUtils.HashPassword(creds.Password)
	if err != nil {
		return err
	}
	role, roleErr := s.RoleRepository.GetRoleByName("CareerCoach")
	if roleErr != nil {
		return fmt.Errorf("failed to get role: %w", roleErr)
	}
	// Prepare user object
	user := models.User{
		Email:          creds.Email,
		HashedPassword: string(hashedPassword),
		UserUUID:       uuid.New().String(),
		UserStatus:     "pending",
		Roles:          []models.Role{*role},
	}
	// Save the user
	return s.UserRepo.SaveUser(user)
}

func (s *RegistrationService) JobSeekerRegistration(tokenBody string, creds models.Credentials) error {
	savedUser, err := s.UserRepo.GetUserByUUID(tokenBody)
	if err != nil {
		return errors.New("UUID doesn't match a user")
	}
	// check password requirement
	isMatch := utils.RegexPassword(creds.Password)
	if !isMatch {
		return errors.New("Password doesn't match requirement")
	}
	hashedPassword, hashErr := s.HashingUtils.HashPassword(creds.Password)
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

func (s *RegistrationService) PreRegisterJobSeeker(email string, careerSuportID *uint) (*models.User, error) {
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
	role, err := s.RoleRepository.GetRoleByName("JobSeeker")
	if err != nil {
		return nil, fmt.Errorf("failed to get role: %w", err)
	}
	user := models.User{
		Email:           email,
		UserStatus:      "pre-register",
		UserUUID:        uuid.New().String(),
		CareerSupportID: careerSuportID,
		Roles:           []models.Role{*role},
	}
	savedUser, err := s.UserRepo.PreRegisterJobSeeker(user)
	if err != nil {
		return nil, errors.New("Error during user pre-registration")
	}
	return savedUser, nil
}
