package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/R-Thibault/OrgaJobSearch/backend/controllers"
	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	mailerService "github.com/R-Thibault/OrgaJobSearch/backend/services"
	serviceMocks "github.com/R-Thibault/OrgaJobSearch/backend/services/mocks"
	JWTMock "github.com/R-Thibault/OrgaJobSearch/backend/utils/mocks"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGenerateGlobalURLInvitation_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(serviceMocks.UserServiceInterface)
	mockOTPService := new(serviceMocks.OTPServiceInterface)
	mockTokenGenerator := new(JWTMock.JWTTokenGeneratorUtilInterface)
	mockRegistrationService := new(serviceMocks.RegistrationServiceInterface)
	mailerService := mailerService.NewMailerService()
	invitationController := controllers.NewUserInvitationController(mockUserService, mockTokenGenerator, *mailerService, mockOTPService, mockRegistrationService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	invitation := models.GlobalInvitation{
		UserID: 1,
	}

	body, _ := json.Marshal(invitation)

	c.Request, _ = http.NewRequest(http.MethodPost, "/generate-invitation", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")

	// Set up mocks for the success case
	mockUserService.On("GetUserByID", uint(1)).Return(&models.User{Model: gorm.Model{
		ID: 1,
	}}, nil)
	mockOTPService.On("GenerateOTP", uint(1), "GlobalInvitation").Return("otp123", nil)
	mockTokenGenerator.On("GenerateJWTToken", mock.Anything, mock.Anything, mock.Anything).Return("jwtToken123", nil)

	invitationController.GenerateGlobalURLInvitation(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "http://localhost:3000/sign-up?token=jwtToken123")
	mockUserService.AssertExpectations(t)
	mockOTPService.AssertExpectations(t)
	mockTokenGenerator.AssertExpectations(t)
}

func TestGenerateGlobalURLInvitation_InvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(serviceMocks.UserServiceInterface)
	mockOTPService := new(serviceMocks.OTPServiceInterface)
	mockTokenGenerator := new(JWTMock.JWTTokenGeneratorUtilInterface)
	mockRegistrationService := new(serviceMocks.RegistrationServiceInterface)
	mailerService := mailerService.NewMailerService()
	invitationController := controllers.NewUserInvitationController(mockUserService, mockTokenGenerator, *mailerService, mockOTPService, mockRegistrationService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Empty body to simulate invalid input
	c.Request, _ = http.NewRequest(http.MethodPost, "/generate-invitation", bytes.NewBuffer([]byte("{}")))
	c.Request.Header.Set("Content-type", "application/json")
	// Mock the behavior for GetUserByID with any user ID to prevent the test from failing
	mockUserService.On("GetUserByID", mock.AnythingOfType("uint")).Return(nil, errors.New("Invalid user ID"))

	invitationController.GenerateGlobalURLInvitation(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid request data")
}

func TestGenerateGlobalURLInvitation_UserNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(serviceMocks.UserServiceInterface)
	mockOTPService := new(serviceMocks.OTPServiceInterface)
	mockTokenGenerator := new(JWTMock.JWTTokenGeneratorUtilInterface)
	mockRegistrationService := new(serviceMocks.RegistrationServiceInterface)
	mailerService := mailerService.NewMailerService()
	invitationController := controllers.NewUserInvitationController(mockUserService, mockTokenGenerator, *mailerService, mockOTPService, mockRegistrationService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	invitation := models.GlobalInvitation{
		UserID: 2,
	}

	body, _ := json.Marshal(invitation)

	c.Request, _ = http.NewRequest(http.MethodPost, "/generate-invitation", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")

	// Mock behavior for user not found
	mockUserService.On("GetUserByID", uint(2)).Return(nil, gorm.ErrRecordNotFound)

	invitationController.GenerateGlobalURLInvitation(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid token")
	mockUserService.AssertExpectations(t)
}

func TestGenerateGlobalURLInvitation_OTPGenerationError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(serviceMocks.UserServiceInterface)
	mockOTPService := new(serviceMocks.OTPServiceInterface)
	mockTokenGenerator := new(JWTMock.JWTTokenGeneratorUtilInterface)
	mockRegistrationService := new(serviceMocks.RegistrationServiceInterface)
	mailerService := mailerService.NewMailerService()
	invitationController := controllers.NewUserInvitationController(mockUserService, mockTokenGenerator, *mailerService, mockOTPService, mockRegistrationService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	invitation := models.GlobalInvitation{
		UserID: 1,
	}

	body, _ := json.Marshal(invitation)

	c.Request, _ = http.NewRequest(http.MethodPost, "/generate-invitation", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")

	// Set up mocks for OTP generation error
	mockUserService.On("GetUserByID", uint(1)).Return(&models.User{Model: gorm.Model{
		ID: 1,
	}}, nil)
	mockOTPService.On("GenerateOTP", uint(1), "GlobalInvitation").Return("", errors.New("OTP generation error"))

	invitationController.GenerateGlobalURLInvitation(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Contains(t, w.Body.String(), "OTP generation error")
	mockUserService.AssertExpectations(t)
	mockOTPService.AssertExpectations(t)
}

func TestGenerateGlobalURLInvitation_TokenGenerationError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(serviceMocks.UserServiceInterface)
	mockOTPService := new(serviceMocks.OTPServiceInterface)
	mockTokenGenerator := new(JWTMock.JWTTokenGeneratorUtilInterface)
	mockRegistrationService := new(serviceMocks.RegistrationServiceInterface)
	mailerService := mailerService.NewMailerService()
	invitationController := controllers.NewUserInvitationController(mockUserService, mockTokenGenerator, *mailerService, mockOTPService, mockRegistrationService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	invitation := models.GlobalInvitation{
		UserID: 1,
	}

	body, _ := json.Marshal(invitation)

	c.Request, _ = http.NewRequest(http.MethodPost, "/generate-invitation", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")

	// Set up mocks for token generation error
	mockUserService.On("GetUserByID", uint(1)).Return(&models.User{Model: gorm.Model{
		ID: 1,
	}}, nil)
	mockOTPService.On("GenerateOTP", uint(1), "GlobalInvitation").Return("otp123", nil)
	mockTokenGenerator.On("GenerateJWTToken", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New("Token generation error"))

	invitationController.GenerateGlobalURLInvitation(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Contains(t, w.Body.String(), "Token generation error")
	mockUserService.AssertExpectations(t)
	mockOTPService.AssertExpectations(t)
	mockTokenGenerator.AssertExpectations(t)
}
