package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/R-Thibault/OrgaJobSearch/backend/controllers"
	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	serviceMocks "github.com/R-Thibault/OrgaJobSearch/backend/services/mocks"
	JWTMock "github.com/R-Thibault/OrgaJobSearch/backend/utils/mocks"
	hashMocks "github.com/R-Thibault/OrgaJobSearch/backend/utils/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestSignIn_SignIn_UserNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(serviceMocks.UserServiceInterface)
	mockHashingService := new(hashMocks.HashingServiceInterface)
	mockJWTTokenGenerator := new(JWTMock.JWTTokenGeneratorUtilInterface)
	mockTokenService := new(serviceMocks.TokenServiceInterface)
	authController := controllers.NewAuthController(mockUserService, mockHashingService, mockTokenService, mockJWTTokenGenerator)

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

	authController.Login(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")
	mockUserService.AssertExpectations(t)
}

func TestSignIn_Successful(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(serviceMocks.UserServiceInterface)
	mockHashingService := new(hashMocks.HashingServiceInterface)
	mockJWTTokenGenerator := new(JWTMock.JWTTokenGeneratorUtilInterface)
	mockTokenService := new(serviceMocks.TokenServiceInterface)
	authController := controllers.NewAuthController(mockUserService, mockHashingService, mockTokenService, mockJWTTokenGenerator)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	creds := models.Credentials{
		Email:    "user@example.com",
		Password: "superPassword1!",
	}
	// expirationTime := time.Now().Add(60 * time.Minute)
	cookieName := "Cookie"
	body, _ := json.Marshal(creds)

	c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")
	hashedPassword := "encodedSalt:encodedHash"
	mockUserService.On("GetUserByEmail", creds.Email).Return(&models.User{Email: creds.Email, HashedPassword: hashedPassword, EmailIsValide: true}, nil)
	mockHashingService.On("CompareHashPassword", creds.Password, hashedPassword).Return(true, nil)

	mockJWTTokenGenerator.On("GenerateJWTToken", &cookieName, mock.AnythingOfType("*string"), mock.Anything).Return("string", nil)
	authController.Login(c)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code 200, but got %v", w.Code)
	assert.Contains(t, w.Body.String(), "Sign in successful")
	mockUserService.AssertExpectations(t)
	mockHashingService.AssertExpectations(t)
	mockJWTTokenGenerator.AssertExpectations(t)
}

func TestSignIn_PasswordDoNotMatch(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(serviceMocks.UserServiceInterface)
	mockHashingService := new(hashMocks.HashingServiceInterface)
	mockJWTTokenGenerator := new(JWTMock.JWTTokenGeneratorUtilInterface)
	mockTokenService := new(serviceMocks.TokenServiceInterface)
	authController := controllers.NewAuthController(mockUserService, mockHashingService, mockTokenService, mockJWTTokenGenerator)

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

	authController.Login(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")
	mockUserService.AssertExpectations(t)
}

func TestLogout_LogoutSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(serviceMocks.UserServiceInterface)
	mockHashingService := new(hashMocks.HashingServiceInterface)
	mockJWTTokenGenerator := new(JWTMock.JWTTokenGeneratorUtilInterface)
	mockTokenService := new(serviceMocks.TokenServiceInterface)
	authController := controllers.NewAuthController(mockUserService, mockHashingService, mockTokenService, mockJWTTokenGenerator)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	creds := models.Credentials{
		Email:    "user@example.com",
		Password: "SuperPassword1!",
	}

	body, _ := json.Marshal(creds)
	c.Request, _ = http.NewRequest(http.MethodPost, "/logout", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")
	c.Request.AddCookie(&http.Cookie{
		Name:  "token",
		Value: "mockTokenValue", // Use a mock token value for the test
		Path:  "/",
	})
	authController.Logout(c)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code 200, but got %v", w.Code)
	assert.Contains(t, w.Body.String(), "Logout successful")
}

func TestLogout_LogoutFail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(serviceMocks.UserServiceInterface)
	mockHashingService := new(hashMocks.HashingServiceInterface)
	mockJWTTokenGenerator := new(JWTMock.JWTTokenGeneratorUtilInterface)
	mockTokenService := new(serviceMocks.TokenServiceInterface)
	authController := controllers.NewAuthController(mockUserService, mockHashingService, mockTokenService, mockJWTTokenGenerator)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// No cookie is set for this test case
	c.Request, _ = http.NewRequest(http.MethodPost, "/logout", nil)
	c.Request.Header.Set("Content-type", "application/json")

	// Call the Logout function
	authController.Logout(c)

	// Assert the response
	assert.Equal(t, http.StatusUnauthorized, w.Code, "Expected status code 401, but got %v", w.Code)
	assert.Contains(t, w.Body.String(), "No token found", "Response body does not contain the expected error message")
}
