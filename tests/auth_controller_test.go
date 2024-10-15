package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/R-Thibault/OrgaJobSearch/controllers"
	"github.com/R-Thibault/OrgaJobSearch/models"
	userMocks "github.com/R-Thibault/OrgaJobSearch/services/mocks"
	hashMocks "github.com/R-Thibault/OrgaJobSearch/utils/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestSignIn_SignIn_UserNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(userMocks.UserServiceInterface)
	mockHashingService := new(hashMocks.HashingServiceInterface)
	authController := controllers.NewAuthController(mockUserService, mockHashingService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	creds := models.Credentials{
		Email:    "test@example.com",
		Password: "superPassword1!",
	}

	body, _ := json.Marshal(creds)

	c.Request, _ = http.NewRequest(http.MethodPost, "/sign-in", bytes.NewBuffer(body))
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
	authController := controllers.NewAuthController(mockUserService, mockHashingService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	creds := models.Credentials{
		Email:    "user@example.com",
		Password: "superPassword1!",
	}

	body, _ := json.Marshal(creds)

	c.Request, _ = http.NewRequest(http.MethodPost, "/sign-in", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")
	hashedPassword := "encodedSalt:encodedHash"
	mockUserService.On("GetUserByEmail", creds.Email).Return(&models.User{Email: creds.Email, HashedPassword: hashedPassword}, nil)
	mockHashingService.On("CompareHashPassword", creds.Password, hashedPassword).Return(true, nil)

	authController.SignIn(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Sign in successful")
	mockUserService.AssertExpectations(t)
}

func TestSignIn_PasswordDoNotMatch(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(userMocks.UserServiceInterface)
	mockHashingService := new(hashMocks.HashingServiceInterface)
	authController := controllers.NewAuthController(mockUserService, mockHashingService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	creds := models.Credentials{
		Email:    "user@example.com",
		Password: "SuperPassword1!",
	}

	body, _ := json.Marshal(creds)

	c.Request, _ = http.NewRequest(http.MethodPost, "/sign-in", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")
	hashedPassword := "encodedSalt:encodedHash"
	mockUserService.On("GetUserByEmail", creds.Email).Return(&models.User{Email: creds.Email, HashedPassword: hashedPassword}, nil)
	mockHashingService.On("CompareHashPassword", creds.Password, hashedPassword).Return(false, nil)

	authController.SignIn(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")
	mockUserService.AssertExpectations(t)
}
