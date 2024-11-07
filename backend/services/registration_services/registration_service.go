package registrationservices

import (
	"errors"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	userRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/user_repository"
	"github.com/R-Thibault/OrgaJobSearch/backend/utils"
	hashingUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/hash_util"
	"github.com/google/uuid"
)

type RegistrationService struct {
	UserRepo     userRepository.UserRepositoryInterface
	HashingUtils hashingUtils.HashingServiceInterface
}

func NewRegistrationService(
	UserRepo userRepository.UserRepositoryInterface,
	HashingUtils hashingUtils.HashingServiceInterface,
) *RegistrationService {
	return &RegistrationService{
		UserRepo:     UserRepo,
		HashingUtils: HashingUtils,
	}
}

func (s *RegistrationService) UserRegistration(creds models.Credentials) error {
	existingUser, _ := s.UserRepo.GetUserByEmail(creds.Email)
	if existingUser != nil {
		return errors.New("user already exists")
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

	user := models.User{
		Email:          creds.Email,
		HashedPassword: hashedPassword,
		UserStatus:     "pre-register",
		UserUUID:       uuid.New().String(),
	}
	updateErr := s.UserRepo.SaveUser(user)
	if updateErr != nil {
		return errors.New("Error during user update")
	}
	return nil
}
