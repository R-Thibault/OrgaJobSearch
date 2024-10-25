package tests

import (
	"testing"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	mockRepo "github.com/R-Thibault/OrgaJobSearch/backend/repository/mocks"
	tokenServices "github.com/R-Thibault/OrgaJobSearch/backend/services/token_services"
	userServices "github.com/R-Thibault/OrgaJobSearch/backend/services/user_services"
	mockUtil "github.com/R-Thibault/OrgaJobSearch/backend/utils/mocks"
	tokenGeneratorUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/tokenGenerator_util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestInvitationSignup_EmailExist(t *testing.T) {
	mockRepo := new(mockRepo.UserRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	userService := userServices.NewUserService(mockRepo, mockHashingService)

	userInvitation := models.UserInvitation{
		Email: "existing@example.com",
	}

	mockRepo.On("GetUserByEmail", userInvitation.Email).Return(&models.User{Email: userInvitation.Email}, nil)

	err := userService.PreRegisterUser(userInvitation.Email, nil)
	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestInvitationSignup_UserPreRegisteredCorrectly(t *testing.T) {
	mockRepo := new(mockRepo.UserRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	userService := userServices.NewUserService(mockRepo, mockHashingService)

	userInvitation := models.UserInvitation{
		UserID: uint(10),
		Email:  "existing@example.com",
	}
	user := &models.User{
		Model: gorm.Model{
			ID: 1,
		},
		Email: userInvitation.Email,
	}

	mockRepo.On("GetUserByEmail", userInvitation.Email).Return(nil, nil)
	mockRepo.On("PreRegisterUser", mock.AnythingOfType("models.User")).Return(nil)
	mockRepo.On("GetUserByID", userInvitation.UserID).Return(user, nil)

	err := userService.PreRegisterUser(userInvitation.Email, &userInvitation.UserID)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "GetUserByEmail", userInvitation.Email)
	mockRepo.AssertExpectations(t)
}

func TestInvitationSignup_VerifyTokenFail(t *testing.T) {
	var tokenGenerator tokenGeneratorUtils.JWTTokenGeneratorServiceInterface = tokenGeneratorUtils.NewJWTTokenGeneratorService()
	tokenService := tokenServices.NewTokenService()
	expirationTime := time.Now().Add(-1 * time.Hour)
	tokenType := "personnalInvitation"
	newUUID := uuid.New().String()

	tokenString, err := tokenGenerator.GenerateJWTToken(&tokenType, &newUUID, expirationTime)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := tokenService.VerifyToken(tokenString)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "token is expired")
	assert.Nil(t, claims)

}

func TestInvitationSignup_VerifyTokenPass(t *testing.T) {
	var tokenGenerator tokenGeneratorUtils.JWTTokenGeneratorServiceInterface = tokenGeneratorUtils.NewJWTTokenGeneratorService()
	tokenService := tokenServices.NewTokenService()
	expirationTime := time.Now().Add(1 * time.Hour)
	invitationType := "personnalInvitation"
	newUUID := uuid.New().String()

	tokenString, err := tokenGenerator.GenerateJWTToken(&invitationType, &newUUID, expirationTime)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := tokenService.VerifyToken(tokenString)

	assert.NoError(t, err)
	assert.Equal(t, invitationType, *claims.TokenType)
	assert.Equal(t, newUUID, claims.Body)
	assert.Equal(t, expirationTime.Unix(), claims.ExpiresAt)
}
