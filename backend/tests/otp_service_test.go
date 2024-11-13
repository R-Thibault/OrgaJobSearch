package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	mockRepo "github.com/R-Thibault/OrgaJobSearch/backend/repository/mocks"
	otpServices "github.com/R-Thibault/OrgaJobSearch/backend/services/otp_services"
	mockUtil "github.com/R-Thibault/OrgaJobSearch/backend/utils/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestGenerateOTP_UserNotFound(t *testing.T) {
	mockRepoUser := new(mockRepo.UserRepositoryInterface)
	mockRepoOTP := new(mockRepo.OTPRepositoryInterface)
	mockUtilOTP := new(mockUtil.OtpGeneratorServiceInterface)
	otpService := otpServices.NewOTPService(mockRepoUser, mockRepoOTP, mockUtilOTP)

	userID := uint(1)
	otpType := "emailValidation"
	expirationTime := time.Now().Add(48 * time.Hour)
	// Setup mock expectation
	mockRepoUser.On("GetUserByID", userID).Return(nil, errors.New("user not found"))

	// Execute function
	_, err := otpService.GenerateOTP(userID, otpType, expirationTime)

	//Assertions
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	mockRepoUser.AssertExpectations(t)
}

func TestGenerateOTP_Success(t *testing.T) {
	mockRepoUser := new(mockRepo.UserRepositoryInterface)
	mockRepoOTP := new(mockRepo.OTPRepositoryInterface)
	mockUtilOTP := new(mockUtil.OtpGeneratorServiceInterface)
	otpService := otpServices.NewOTPService(mockRepoUser, mockRepoOTP, mockUtilOTP)

	email := "existinguser@example.com"
	otp := "123456"
	user := &models.User{Model: gorm.Model{
		ID: 1,
	}, Email: email}
	otpType := "emailValidation"
	expirationTime := time.Now().Add(48 * time.Hour)
	// Setup mock expectation
	mockRepoUser.On("GetUserByID", user.ID).Return(user, nil)
	mockUtilOTP.On("GenerateOTP", user, otpType, expirationTime).Return(models.OTP{
		UserID:        user.ID,
		OtpCode:       otp,
		OtpExpiration: time.Now().Add(60 * time.Minute),
		OtpType:       "login",
		OtpAttempts:   0,
		IsValid:       true,
	})
	mockRepoOTP.On("SaveOTP", mock.AnythingOfType("models.OTP")).Return(otp, nil)

	// Execute function
	generatedOTP, err := otpService.GenerateOTP(user.ID, otpType, expirationTime)

	//Assertions
	assert.NoError(t, err)
	assert.Equal(t, otp, generatedOTP)
	mockRepoOTP.AssertCalled(t, "SaveOTP", mock.AnythingOfType("models.OTP"))
}

func TestVerifyOTP_Sucess(t *testing.T) {
	mockRepoUser := new(mockRepo.UserRepositoryInterface)
	mockRepoOTP := new(mockRepo.OTPRepositoryInterface)
	mockUtilOTP := new(mockUtil.OtpGeneratorServiceInterface)
	otpService := otpServices.NewOTPService(mockRepoUser, mockRepoOTP, mockUtilOTP)

	otp := "123456"
	otpType := "emailValidation"
	email := "existinguser@example.com"
	user := &models.User{Model: gorm.Model{
		ID: 1,
	}, Email: email}

	validOTP := &models.OTP{
		OtpCode:       otp,
		OtpType:       "emailValidation",
		OtpExpiration: time.Now().Add(10 * time.Minute),
	}

	// Setup mock expectation
	mockRepoUser.On("GetUserByEmail", email).Return(user, nil)
	mockRepoOTP.On("GetOTPCodeByUserIDandType", user.ID, validOTP.OtpType).Return(validOTP, nil)

	// Execute function
	err := otpService.VerifyOTPGiven(email, otpType, otp)

	// Assertions
	assert.NoError(t, err)
	mockRepoOTP.AssertCalled(t, "GetOTPCodeByUserIDandType", user.ID, validOTP.OtpType)
}

func TestVerifyOTP_Fail_IncorrectOTP(t *testing.T) {
	mockRepoUser := new(mockRepo.UserRepositoryInterface)
	mockRepoOTP := new(mockRepo.OTPRepositoryInterface)
	mockUtilOTP := new(mockUtil.OtpGeneratorServiceInterface)
	otpService := otpServices.NewOTPService(mockRepoUser, mockRepoOTP, mockUtilOTP)

	email := "existinguser@example.com"
	user := &models.User{
		Model: gorm.Model{
			ID: 1,
		}, Email: email}

	otpType := "emailValidation"
	incorrectOTP := "incorrectOTP"
	validOTP := &models.OTP{
		UserID:        user.ID,
		OtpCode:       "123456",
		OtpType:       "emailValidation",                // The correct OTP for this user
		OtpExpiration: time.Now().Add(10 * time.Minute), // OTP is valid
	}
	// Setup mock expectation
	mockRepoUser.On("GetUserByEmail", email).Return(user, nil)
	mockRepoOTP.On("GetOTPCodeByUserIDandType", user.ID, validOTP.OtpType).Return(validOTP, nil)

	// Execute function
	err := otpService.VerifyOTPGiven(email, otpType, incorrectOTP)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "OTP codes do not match", err.Error())
	// Ensure all expectations were met
	mockRepoUser.AssertExpectations(t)
	mockRepoOTP.AssertExpectations(t)
}

func TestVerifyOTP_Fail_ExpiredOTP(t *testing.T) {
	mockRepoUser := new(mockRepo.UserRepositoryInterface)
	mockRepoOTP := new(mockRepo.OTPRepositoryInterface)
	mockUtilOTP := new(mockUtil.OtpGeneratorServiceInterface)
	otpService := otpServices.NewOTPService(mockRepoUser, mockRepoOTP, mockUtilOTP)

	email := "existinguser@example.com"
	user := &models.User{
		Model: gorm.Model{
			ID: 1,
		}, Email: email}

	invalidOTP := &models.OTP{
		UserID:        user.ID,
		OtpCode:       "123456",
		OtpType:       "emailValidation",
		OtpExpiration: time.Now().Add(-10 * time.Minute),
	}
	// Setup mock expectation
	mockRepoUser.On("GetUserByEmail", email).Return(user, nil)
	mockRepoOTP.On("GetOTPCodeByUserIDandType", user.ID, invalidOTP.OtpType).Return(invalidOTP, nil)

	// Execute function
	err := otpService.VerifyOTPGiven(email, invalidOTP.OtpType, invalidOTP.OtpCode)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "OTP expired", err.Error())
	// Ensure all expectations were met
	mockRepoUser.AssertExpectations(t)
	mockRepoOTP.AssertExpectations(t)
}
