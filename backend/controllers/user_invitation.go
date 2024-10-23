package controllers

import (
	"net/http"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	mailerService "github.com/R-Thibault/OrgaJobSearch/backend/services"
	userServices "github.com/R-Thibault/OrgaJobSearch/backend/services/user_services"
	tokenUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/tokenGenerator_util"
	"github.com/gin-gonic/gin"
)

type UserInvitationController struct {
	UserService        *userServices.UserService
	TokenGeneratorUtil tokenUtils.JWTTokenGeneratorServiceInterface
	MailerService      mailerService.MailerService
}

func NewUserInvitationController(UserService *userServices.UserService, TokenGeneratorUtil tokenUtils.JWTTokenGeneratorServiceInterface, MailerService mailerService.MailerService) *UserInvitationController {
	return &UserInvitationController{UserService: UserService, TokenGeneratorUtil: TokenGeneratorUtil, MailerService: MailerService}
}

func (u *UserInvitationController) SendUserInvitation(c *gin.Context) {
	var userInvitation models.UserInvitation
	if err := c.ShouldBindJSON(&userInvitation); err != nil {
		// If the input is invalid, respond with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	err := u.UserService.PreRegisterUser(userInvitation.Email)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	// Set expiration time for token
	expirationTime := time.Now().Add(8 * time.Hour)
	// Generate Token here
	jwtTokenString, err := u.TokenGeneratorUtil.GenerateJWTToken(&userInvitation.UserID, userInvitation.Email, expirationTime)

	// Build email template with url + tokenstring and send it
	mailerErr := u.MailerService.SendUserSignUpInvitation(userInvitation.Email, jwtTokenString)
	if mailerErr != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Email sent"})

}
