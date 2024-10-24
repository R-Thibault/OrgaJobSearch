package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/R-Thibault/OrgaJobSearch/backend/controllers"
	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	userMocks "github.com/R-Thibault/OrgaJobSearch/backend/services/mocks"
	JWTMock "github.com/R-Thibault/OrgaJobSearch/backend/utils/mocks"
	hashMocks "github.com/R-Thibault/OrgaJobSearch/backend/utils/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestSignIn_SignIn_UserNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(userMocks.UserServiceInterface)
	mockHashingService := new(hashMocks.HashingServiceInterface)
	mockJWTTokenGenerator := new(JWTMock.JWTTokenGeneratorServiceInterface)
	authController := controllers.NewAuthController(mockUserService, mockHashingService, mockJWTTokenGenerator)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	creds := models.Credentials{
		Email:    "test@example.com",
		Password: "superPassword1!",
	}

	body, _ := json.Marshal(creds)

	c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")

	// Setup mock expectations to simulate an existing user
	mockUserService.On("GetUserByEmail", creds.Email).Return(nil, gorm.ErrRecordNotFound)

	authController.SignIn(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")
	mockUserService.AssertExpectations(t)
}

func TestSignIn_Successful(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(userMocks.UserServiceInterface)
	mockHashingService := new(hashMocks.HashingServiceInterface)
	mockJWTTokenGenerator := new(JWTMock.JWTTokenGeneratorServiceInterface)
	authController := controllers.NewAuthController(mockUserService, mockHashingService, mockJWTTokenGenerator)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	creds := models.Credentials{
		Email:    "user@example.com",
		Password: "superPassword1!",
	}

	body, _ := json.Marshal(creds)

	c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")
	hashedPassword := "encodedSalt:encodedHash"
	mockUserService.On("GetUserByEmail", creds.Email).Return(&models.User{Email: creds.Email, HashedPassword: hashedPassword, EmailIsValide: true}, nil)
	mockHashingService.On("CompareHashPassword", creds.Password, hashedPassword).Return(true, nil)
	mockJWTTokenGenerator.On("GenerateJWTToken", (*string)(nil), creds.Email, mock.Anything).Return("string", nil)
	authController.SignIn(c)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code 200, but got %v", w.Code)
	assert.Contains(t, w.Body.String(), "Sign in successful")
	mockUserService.AssertExpectations(t)
	mockHashingService.AssertExpectations(t)
	mockJWTTokenGenerator.AssertExpectations(t)
}

func TestSignIn_PasswordDoNotMatch(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(userMocks.UserServiceInterface)
	mockHashingService := new(hashMocks.HashingServiceInterface)
	mockJWTTokenGenerator := new(JWTMock.JWTTokenGeneratorServiceInterface)
	authController := controllers.NewAuthController(mockUserService, mockHashingService, mockJWTTokenGenerator)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	creds := models.Credentials{
		Email:    "user@example.com",
		Password: "SuperPassword1!",
	}

	body, _ := json.Marshal(creds)

	c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")
	hashedPassword := "encodedSalt:encodedHash"
	mockUserService.On("GetUserByEmail", creds.Email).Return(&models.User{Email: creds.Email, HashedPassword: hashedPassword}, nil)
	mockHashingService.On("CompareHashPassword", creds.Password, hashedPassword).Return(false, nil)

	authController.SignIn(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")
	mockUserService.AssertExpectations(t)
}
