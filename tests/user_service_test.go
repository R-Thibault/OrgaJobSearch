package tests

import (
	"testing"

	"github.com/R-Thibault/OrgaJobSearch/models"
	"github.com/R-Thibault/OrgaJobSearch/repository/mocks"
	mockRepo "github.com/R-Thibault/OrgaJobSearch/repository/mocks"
	"github.com/R-Thibault/OrgaJobSearch/services"
	mockUtil "github.com/R-Thibault/OrgaJobSearch/utils/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser_UserAlreadyExists(t *testing.T) {
	mockRepo := new(mockRepo.UserRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	userService := services.NewUserService(mockRepo, mockHashingService)

	creds := models.Credentials{
		Email:    "existing@example.com",
		Password: "superPassword1!",
	}

	// Setup mock expectations to simulate an existing user
	mockRepo.On("GetUserByEmail", creds.Email).Return(&models.User{Email: creds.Email}, nil)

	// Execute the function
	err := userService.RegisterUser(creds)

	//Assertions
	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestRegisterUser_PasswordRegexCheckFail(t *testing.T) {
	mockRepo := new(mockRepo.UserRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	UserService := services.NewUserService(mockRepo, mockHashingService)

	creds := models.Credentials{
		Email:    "test@example.com",
		Password: "superPassword1",
	}
	mockRepo.On("GetUserByEmail", creds.Email).Return(nil, nil)
	//Execute function
	err := UserService.RegisterUser(creds)

	//Assertions
	assert.Error(t, err)
	assert.Equal(t, "Password doesn't match requirement", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestRegisterUser_PasswordRegexCheckPass(t *testing.T) {
	//Setup the mock repository
	mockRepo := new(mocks.UserRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	userService := services.NewUserService(mockRepo, mockHashingService)

	// Define credentials
	creds := models.Credentials{
		Email:    "test@example.com",
		Password: "Password1!",
	}
	hashedPassword := "encodedSalt:encodedHash"
	// Mock the repository behavior
	mockRepo.On("GetUserByEmail", creds.Email).Return(nil, nil)
	mockHashingService.On("HashPassword", creds.Password).Return(hashedPassword, nil)
	mockRepo.On("SaveUser", mock.AnythingOfType("models.User")).Return(nil)

	// Execute the function
	err := userService.RegisterUser(creds)

	//Assertions
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "GetUserByEmail", creds.Email)
	mockRepo.AssertCalled(t, "SaveUser", mock.AnythingOfType("models.User"))
}
