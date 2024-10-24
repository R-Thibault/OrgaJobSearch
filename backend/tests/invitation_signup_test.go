package tests

import (
	"testing"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	mockRepo "github.com/R-Thibault/OrgaJobSearch/backend/repository/mocks"
	userServices "github.com/R-Thibault/OrgaJobSearch/backend/services/user_services"
	mockUtil "github.com/R-Thibault/OrgaJobSearch/backend/utils/mocks"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestInvitationSignup_EmailExist(t *testing.T) {
	mockRepo := new(mockRepo.UserRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	userService := userServices.NewUserService(mockRepo, mockHashingService)

	expirationTime := time.Now().Add(1 * time.Hour)
	invitation := models.JWTToken{
		Email: "existing@example.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	mockRepo.On("GetUserByEmail", invitation.Email).Return(&models.User{Email: invitation.Email}, nil)

	err := userService.PreRegisterUser(invitation.Email, nil)
	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestInvitationSignup_UserPreRegisteredCorrectly(t *testing.T) {
	mockRepo := new(mockRepo.UserRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	userService := userServices.NewUserService(mockRepo, mockHashingService)

	expirationTime := time.Now().Add(1 * time.Hour)
	invitationToken := models.JWTToken{
		UserID: func(u uint) *uint { return &u }(10),
		Email:  "existing@example.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	user := &models.User{
		Model: gorm.Model{
			ID: 1,
		},
		Email: invitationToken.Email,
		// You can add more fields here if needed
	}

	mockRepo.On("GetUserByEmail", invitationToken.Email).Return(nil, nil)
	mockRepo.On("PreRegisterUser", mock.AnythingOfType("models.User")).Return(nil)
	mockRepo.On("GetUserByID", *invitationToken.UserID).Return(user, nil)

	err := userService.PreRegisterUser(invitationToken.Email, invitationToken.UserID)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "GetUserByEmail", invitationToken.Email)
	mockRepo.AssertExpectations(t)
}
