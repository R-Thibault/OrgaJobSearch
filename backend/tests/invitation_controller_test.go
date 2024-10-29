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
		InvitationType: "GlobalInvitation",
	}

	body, _ := json.Marshal(invitation)

	c.Request, _ = http.NewRequest(http.MethodPost, "/generate-url", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")

	// Set the userUUID in context
	c.Set("userUUID", "valid-uuid")
	// Set up mocks for OTP generation error
	mockUserService.On("GetUserByUUID", "valid-uuid").Return(&models.User{Model: gorm.Model{
		ID: 1,
	}}, nil)
	mockOTPService.On("GenerateOTP", uint(1), invitation.InvitationType).Return("otp123", nil)
	mockOTPService.On("CheckOTPCodeForGlobalInvitation", uint(1), invitation.InvitationType).Return("", nil)
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
	invitation := models.GlobalInvitation{
		InvitationType: "GlobalInvitation",
	}

	body, _ := json.Marshal(invitation)

	c.Request, _ = http.NewRequest(http.MethodPost, "/generate-url", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")

	// Set the userUUID in context
	c.Set("userUUID", "valid-uuid")
	// Mock the behavior for GetUserByID with any user ID to prevent the test from failing
	mockUserService.On("GetUserByUUID", "valid-uuid").Return(nil, errors.New("Invalid user UUID"))

	invitationController.GenerateGlobalURLInvitation(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "UserUUID does not match a user")
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
	c.Set("userUUID", "valid-uuid")
	invitation := models.GlobalInvitation{
		InvitationType: "GlobalInvitation",
	}

	body, _ := json.Marshal(invitation)

	c.Request, _ = http.NewRequest(http.MethodPost, "/generate-url", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")

	// Mock behavior for user not found
	mockUserService.On("GetUserByUUID", "valid-uuid").Return(nil, errors.New("user not found"))

	invitationController.GenerateGlobalURLInvitation(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "UserUUID does not match a user")
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
		InvitationType: "GlobalInvitation",
	}

	body, _ := json.Marshal(invitation)

	c.Request, _ = http.NewRequest(http.MethodPost, "/generate-url", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")

	// Set the userUUID in context
	c.Set("userUUID", "valid-uuid")

	// Set up mocks for OTP generation error
	mockUserService.On("GetUserByUUID", "valid-uuid").Return(&models.User{Model: gorm.Model{
		ID: 1,
	}}, nil)
	mockOTPService.On("CheckOTPCodeForGlobalInvitation", uint(1), "GlobalInvitation").Return("", nil)
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
		InvitationType: "GlobalInvitation",
	}

	body, _ := json.Marshal(invitation)

	c.Request, _ = http.NewRequest(http.MethodPost, "/generate-url", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")

	// Set the userUUID in context
	c.Set("userUUID", "valid-uuid")
	// Set up mocks for OTP generation error
	mockUserService.On("GetUserByUUID", "valid-uuid").Return(&models.User{Model: gorm.Model{
		ID: 1,
	}}, nil)
	mockOTPService.On("CheckOTPCodeForGlobalInvitation", uint(1), "GlobalInvitation").Return("", nil)
	mockOTPService.On("GenerateOTP", uint(1), "GlobalInvitation").Return("otp123", nil)
	mockTokenGenerator.On("GenerateJWTToken", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New("Token generation error"))

	invitationController.GenerateGlobalURLInvitation(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Contains(t, w.Body.String(), "Token generation error")
	mockUserService.AssertExpectations(t)
	mockOTPService.AssertExpectations(t)
	mockTokenGenerator.AssertExpectations(t)
}
