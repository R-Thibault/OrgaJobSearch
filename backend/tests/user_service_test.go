package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	mockRepo "github.com/R-Thibault/OrgaJobSearch/backend/repository/mocks"
	registrationservices "github.com/R-Thibault/OrgaJobSearch/backend/services/registration_services"
	userservices "github.com/R-Thibault/OrgaJobSearch/backend/services/user_services"
	mockUtil "github.com/R-Thibault/OrgaJobSearch/backend/utils/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestRegisterUser_UserAlreadyExists(t *testing.T) {
	mockUserRepo := new(mockRepo.UserRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	registrationService := registrationservices.NewRegistrationService(mockUserRepo, mockHashingService)

	creds := models.Credentials{
		Email:    "existing@example.com",
		Password: "superPassword1!",
	}

	// Setup mock expectations to simulate an existing user
	mockUserRepo.On("GetUserByEmail", creds.Email).Return(&models.User{Email: creds.Email}, nil)

	// Execute the function
	err := registrationService.UserRegistration(creds)

	//Assertions
	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())
	mockUserRepo.AssertExpectations(t)
}

func TestRegisterUser_PasswordRegexCheckFail(t *testing.T) {
	mockUserRepo := new(mockRepo.UserRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	registrationService := registrationservices.NewRegistrationService(mockUserRepo, mockHashingService)

	creds := models.Credentials{
		Email:    "test@example.com",
		Password: "superPassword1",
	}
	mockUserRepo.On("GetUserByEmail", creds.Email).Return(nil, nil)
	//Execute function
	err := registrationService.UserRegistration(creds)

	//Assertions
	assert.Error(t, err)
	assert.Equal(t, "Password doesn't match requirement", err.Error())
	mockUserRepo.AssertExpectations(t)
}

func TestRegisterUser_PasswordRegexCheckPass(t *testing.T) {
	mockUserRepo := new(mockRepo.UserRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	registrationService := registrationservices.NewRegistrationService(mockUserRepo, mockHashingService)

	// Define credentials
	creds := models.Credentials{
		Email:    "test@example.com",
		Password: "Password1!",
	}
	hashedPassword := "encodedSalt:encodedHash"

	mockUserRepo.On("GetUserByEmail", creds.Email).Return(nil, nil)
	mockHashingService.On("HashPassword", creds.Password).Return(hashedPassword, nil)
	mockUserRepo.On("SaveUser", mock.AnythingOfType("models.User")).Return(nil)

	// Execute the function
	err := registrationService.UserRegistration(creds)

	//Assertions
	assert.NoError(t, err)
	mockUserRepo.AssertCalled(t, "GetUserByEmail", creds.Email)
	mockUserRepo.AssertCalled(t, "SaveUser", mock.AnythingOfType("models.User"))
}

func TestUser_ResetPasswordSuccess(t *testing.T) {
	mockUserRepo := new(mockRepo.UserRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	mockOTPRepo := new(mockRepo.OTPRepositoryInterface)
	userService := userservices.NewUserService(mockUserRepo, mockOTPRepo, mockHashingService)

	email := "existing@example.com"
	newPassword := "Example1!"
	hashedPassword := "encodedSalt:encodedHash"
	userUUID := "valid-uuid"
	tokenType := "resetPassword"
	bodyContent := "valid-otp"
	jwtToken := models.JWTToken{TokenType: &tokenType, Body: &bodyContent}
	user := &models.User{
		Model: gorm.Model{
			ID: 1,
		},
		UserUUID: userUUID,
		Email:    email}

	mockOTPRepo.On("GetOTPCodeByUserIDandType", user.ID, tokenType).Return(&models.OTP{OtpCode: bodyContent, OtpType: tokenType, OtpExpiration: time.Now().Add(time.Hour)}, nil)

	mockHashingService.On("HashPassword", newPassword).Return(hashedPassword, nil)
	mockUserRepo.On("UpdateUserPassword", user.ID, hashedPassword).Return(nil)

	err := userService.ResetPassword(*user, jwtToken, newPassword)
	//Assertions
	assert.NoError(t, err)
	mockHashingService.AssertCalled(t, "HashPassword", newPassword)
	mockUserRepo.AssertCalled(t, "UpdateUserPassword", user.ID, hashedPassword)

}

func TestUser_ResetPasswordFail(t *testing.T) {
	mockUserRepo := new(mockRepo.UserRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	mockOTPRepo := new(mockRepo.OTPRepositoryInterface)
	userService := userservices.NewUserService(mockUserRepo, mockOTPRepo, mockHashingService)

	email := "existing@example.com"
	newPassword := "Example1!"
	hashedPassword := "encodedSalt:encodedHash"
	tokenType := "resetPassword"
	bodyContent := "valid-otp"
	jwtToken := models.JWTToken{TokenType: &tokenType, Body: &bodyContent}
	userUUID := "valid-uuid"
	user := &models.User{
		Model: gorm.Model{
			ID: 1,
		},
		UserUUID: userUUID,
		Email:    email}

	mockOTPRepo.On("GetOTPCodeByUserIDandType", user.ID, tokenType).Return(&models.OTP{OtpCode: bodyContent, OtpType: tokenType, OtpExpiration: time.Now().Add(time.Hour)}, nil)

	mockHashingService.On("HashPassword", newPassword).Return(hashedPassword, nil)
	mockUserRepo.On("UpdateUserPassword", user.ID, hashedPassword).Return(errors.New("user not found"))

	err := userService.ResetPassword(*user, jwtToken, newPassword)

	//Assertions
	assert.Error(t, err)
	assert.Equal(t, "Update user failed", err.Error())
	mockHashingService.AssertCalled(t, "HashPassword", newPassword)
	mockUserRepo.AssertExpectations(t)

}
