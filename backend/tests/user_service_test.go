package tests

import (
	"testing"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	mockRepo "github.com/R-Thibault/OrgaJobSearch/backend/repository/mocks"
	registrationservices "github.com/R-Thibault/OrgaJobSearch/backend/services/registration_services"
	mockUtil "github.com/R-Thibault/OrgaJobSearch/backend/utils/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser_UserAlreadyExists(t *testing.T) {
	mockUserRepo := new(mockRepo.UserRepositoryInterface)
	mockRoleRepo := new(mockRepo.RoleRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	registrationService := registrationservices.NewRegistrationService(mockUserRepo, mockHashingService, mockRoleRepo)

	creds := models.Credentials{
		Email:    "existing@example.com",
		Password: "superPassword1!",
	}

	// Setup mock expectations to simulate an existing user
	mockUserRepo.On("GetUserByEmail", creds.Email).Return(&models.User{Email: creds.Email}, nil)

	// Execute the function
	err := registrationService.RegisterCareerCoach(creds)

	//Assertions
	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())
	mockUserRepo.AssertExpectations(t)
}

func TestRegisterUser_PasswordRegexCheckFail(t *testing.T) {
	mockUserRepo := new(mockRepo.UserRepositoryInterface)
	mockRoleRepo := new(mockRepo.RoleRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	registrationService := registrationservices.NewRegistrationService(mockUserRepo, mockHashingService, mockRoleRepo)

	creds := models.Credentials{
		Email:    "test@example.com",
		Password: "superPassword1",
	}
	mockUserRepo.On("GetUserByEmail", creds.Email).Return(nil, nil)
	//Execute function
	err := registrationService.RegisterCareerCoach(creds)

	//Assertions
	assert.Error(t, err)
	assert.Equal(t, "Password doesn't match requirement", err.Error())
	mockUserRepo.AssertExpectations(t)
}

func TestRegisterUser_PasswordRegexCheckPass(t *testing.T) {
	mockUserRepo := new(mockRepo.UserRepositoryInterface)
	mockRoleRepo := new(mockRepo.RoleRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	registrationService := registrationservices.NewRegistrationService(mockUserRepo, mockHashingService, mockRoleRepo)

	// Define credentials
	creds := models.Credentials{
		Email:    "test@example.com",
		Password: "Password1!",
	}
	hashedPassword := "encodedSalt:encodedHash"
	jobSeekerRole := &models.Role{
		RoleName: "CareerCoach",
	}
	mockRoleRepo.On("GetRoleByName", "CareerCoach").Return(jobSeekerRole, nil)
	mockUserRepo.On("GetUserByEmail", creds.Email).Return(nil, nil)
	mockHashingService.On("HashPassword", creds.Password).Return(hashedPassword, nil)
	mockUserRepo.On("SaveUser", mock.AnythingOfType("models.User")).Return(nil)

	// Execute the function
	err := registrationService.RegisterCareerCoach(creds)

	//Assertions
	assert.NoError(t, err)
	mockUserRepo.AssertCalled(t, "GetUserByEmail", creds.Email)
	mockUserRepo.AssertCalled(t, "SaveUser", mock.AnythingOfType("models.User"))
}
