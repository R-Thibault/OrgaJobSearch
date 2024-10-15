package tests

import (
	"testing"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/models"
	mockRepo "github.com/R-Thibault/OrgaJobSearch/repository/mocks"
	"github.com/R-Thibault/OrgaJobSearch/services"
	mockUtil "github.com/R-Thibault/OrgaJobSearch/utils/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGenerateOTP_UserNotFound(t *testing.T) {

	mockRepo := new(mockRepo.UserRepositoryInterface)
	optService := services.NewOTPService(mockRepo)

	email := "nonexistinguser@example.com"

	// Setup mock expectation
	mockRepo.On("GetUserByEmail", email).Return(nil, nil)

	// Execute function
	_, err := otpService.GenerateOTP(email)

	//Assertions
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	mockRepo.AssertExpectations()
}

func TestGenerateOTP_Success(t *testing.T) {

	mockRepo := new(mockRepo.UserRepositoryInterface)
	mockOTPGenerator := new(mockUtil.OTPGeneratorInterface)
	otpService := services.NewOTPService(mockRepo)

	email := "existinguser@example.com"
	otp := "123456"
	user := &models.User{Email: email}

	// Setup mock expectation
	mockRepo.On("GetUserByEmail", email).Return(user, nil)
	mockOTPGenerator.On("GenerateORP", user).Return(otp, nil)
	mockRepo.On("SaveOTP", mock.AnythingOfType("models.OTP")).Return(nil)

	// Execute function
	generatedOTP, err := otpService.GenerateOTP(email)

	//Assertions
	assert.NoError(t, err)
	assert.Equal(t, otp, generatedOTP)
	mockRepo.AssertCalled(t, "SaveOTP", mock.AnythingOfType("models.OTP"))
}

func TestVerifyOTP_Sucess(t *testing.T) {
	mockRepo := new(mockRepo.UserRepositoryInterface)
	otpService := services.NewOTPService(mockRepo)

	otp := "123456"
	email := "existinguser@example.com"
	user := &models.User{Email: email}

	validOTP := &models.OTP{
		OtpCode:       otp,
		OtpExpiration: time.Now().Add(10 * time.Minute),
	}

	// Setup mock expectation
	mockRepo.On("GetUserByEmail", email).Return(user, nil)
	mockRepo.On("GetOTP", user).Return(validOTP, nil)

	// Execute function
	err := otpService.VerifyOTP(email, otp)

	// Assertions
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "GetOTP", user)
}

func TestVerifyOTP_Fail_IncorrectOTP(t *testing.T) {
	mockRepo := new(mockRepo.UserRepositoryInterface)
	otpService := services.NewOTPService(mockRepo)

	otp := "incorrectOTP"
	email := "existinguser@example.com"
	user := &models.User{Email: email}

	validOTP := &models.OTP{
		OtpCode:       "123456",
		OtpExpiration: time.Now().Add(10 * time.Minute),
	}

	// Setup mock expectation
	mockRepo.On("GetUserByEmail", email).Return(user, nil)
	mockRepo.On("GetOTP", user).Return(validOTP, nil)

	// Execute function
	err := otpService.VerifyOTP(email, otp)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "invalid OTP", err.Error())
}

func TestVerifyOTP_Fail_ExpiredOTP(t *testing.T) {
	mockRepo := new(mockRepo.UserRepositoryInterface)
	otpService := services.NewOTPService(mockRepo)

	otp := "123456"
	email := "existinguser@example.com"
	user := &models.User{Email: email}

	validOTP := &models.OTP{
		OtpCode:       otp,
		OtpExpiration: time.Now().Add(-10 * time.Minute),
	}
	// Setup mock expectation
	mockRepo.On("GetUserByEmail", email).Return(user, nil)
	mockRepo.On("GetOTP", user).Return(validOTP, nil)

	// Execute function
	err := otpService.VerifyOTP(email, otp)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "OTP expired", err.Error())
}
