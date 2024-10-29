package tests

import (
	"testing"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	mockRepo "github.com/R-Thibault/OrgaJobSearch/backend/repository/mocks"
	registrationservices "github.com/R-Thibault/OrgaJobSearch/backend/services/registration_services"
	tokenServices "github.com/R-Thibault/OrgaJobSearch/backend/services/token_services"
	mockUtil "github.com/R-Thibault/OrgaJobSearch/backend/utils/mocks"
	tokenGeneratorUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/tokenGenerator_util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestInvitationSignup_EmailExist(t *testing.T) {
	mockUserRepo := new(mockRepo.UserRepositoryInterface)
	mockRoleRepo := new(mockRepo.RoleRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	registrationService := registrationservices.NewRegistrationService(mockUserRepo, mockHashingService, mockRoleRepo)

	userInvitation := models.UserInvitation{
		Email: "existing@example.com",
	}

	mockUserRepo.On("GetUserByEmail", userInvitation.Email).Return(&models.User{Email: userInvitation.Email}, nil)

	_, err := registrationService.PreRegisterJobSeeker(userInvitation.Email, nil)
	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())
	mockUserRepo.AssertExpectations(t)
}

func TestInvitationSignup_UserPreRegisteredCorrectly(t *testing.T) {
	mockUserRepo := new(mockRepo.UserRepositoryInterface)
	mockRoleRepo := new(mockRepo.RoleRepositoryInterface)
	mockHashingService := new(mockUtil.HashingServiceInterface)
	registrationService := registrationservices.NewRegistrationService(mockUserRepo, mockHashingService, mockRoleRepo)

	userInvitation := models.UserInvitation{
		Email: "existing@example.com",
	}
	user := &models.User{
		Model: gorm.Model{
			ID: 1,
		},
		UserUUID: uuid.New().String(),
		Email:    userInvitation.Email,
	}
	jobSeekerRole := &models.Role{
		RoleName: "JobSeeker",
	}
	mockRoleRepo.On("GetRoleByName", "JobSeeker").Return(jobSeekerRole, nil)
	mockUserRepo.On("GetUserByEmail", userInvitation.Email).Return(nil, nil)
	mockUserRepo.On("PreRegisterJobSeeker", mock.AnythingOfType("models.User")).Return(user, nil)
	mockUserRepo.On("GetUserByID", user.ID).Return(user, nil)

	savedUser, err := registrationService.PreRegisterJobSeeker(userInvitation.Email, &user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, savedUser)
	mockUserRepo.AssertCalled(t, "GetUserByEmail", userInvitation.Email)
	mockUserRepo.AssertExpectations(t)
	mockRoleRepo.AssertExpectations(t)
}

func TestInvitationSignup_VerifyTokenFail(t *testing.T) {
	var tokenGenerator tokenGeneratorUtils.JWTTokenGeneratorUtilInterface = tokenGeneratorUtils.NewJWTTokenGeneratorUtil()
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
	var tokenGenerator tokenGeneratorUtils.JWTTokenGeneratorUtilInterface = tokenGeneratorUtils.NewJWTTokenGeneratorUtil()
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
	assert.Equal(t, &newUUID, claims.Body)
	assert.Equal(t, expirationTime.Unix(), claims.ExpiresAt)
}
