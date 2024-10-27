package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	mailerService "github.com/R-Thibault/OrgaJobSearch/backend/services"
	otpServices "github.com/R-Thibault/OrgaJobSearch/backend/services/otp_services"
	userServices "github.com/R-Thibault/OrgaJobSearch/backend/services/user_services"
	tokenUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/tokenGenerator_util"
	"github.com/gin-gonic/gin"
)

type UserInvitationController struct {
	UserService        userServices.UserServiceInterface
	TokenGeneratorUtil tokenUtils.JWTTokenGeneratorServiceInterface
	MailerService      mailerService.MailerService
	otpServices        otpServices.OTPServiceInterface
}

func NewUserInvitationController(UserService userServices.UserServiceInterface, TokenGeneratorUtil tokenUtils.JWTTokenGeneratorServiceInterface, MailerService mailerService.MailerService, otpServices otpServices.OTPServiceInterface) *UserInvitationController {
	return &UserInvitationController{UserService: UserService, TokenGeneratorUtil: TokenGeneratorUtil, MailerService: MailerService, otpServices: otpServices}
}

func (u *UserInvitationController) SendJobSeekerInvitation(c *gin.Context) {
	var userInvitation models.UserInvitation
	if err := c.ShouldBindJSON(&userInvitation); err != nil {
		// If the input is invalid, respond with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	savedUser, err := u.UserService.PreRegisterJobSeeker(userInvitation.Email, &userInvitation.UserID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	userInvitation.InvitationType = "PersonalInvitation"
	// Set expiration time for token
	expirationTime := time.Now().Add(8 * time.Hour)
	// Generate Token here
	jwtTokenString, err := u.TokenGeneratorUtil.GenerateJWTToken(&userInvitation.InvitationType, &savedUser.UserUUID, expirationTime)

	// Build email template with url + tokenstring and send it
	mailerErr := u.MailerService.SendJobSeekerSignUpInvitation(userInvitation.Email, jwtTokenString)
	if mailerErr != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Email sent"})

}

func (u *UserInvitationController) GenerateGlobalURLInvitation(c *gin.Context) {
	var invitation models.GlobalInvitation
	if err := c.ShouldBindJSON(&invitation); err != nil || invitation.UserID == 0 {
		// If the input is invalid, respond with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	user, err := u.UserService.GetUserByID(invitation.UserID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	//Generate the OTP with a type "GlobalInvitation"
	otpType := "GlobalInvitation"
	otpGenerated, err := u.otpServices.GenerateOTP(user.ID, otpType)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	// then add the otp to the JWTToken
	invitation.InvitationType = "GlobalInvitation"
	// Set expiration time for token
	expirationTime := time.Now().Add(8 * time.Hour)
	// Generate Token here
	jwtTokenString, err := u.TokenGeneratorUtil.GenerateJWTToken(&invitation.InvitationType, &otpGenerated, expirationTime)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	url := fmt.Sprintf(`http://localhost:3000/sign-up?token=%s`, jwtTokenString)
	c.JSON(http.StatusOK, gin.H{"url": url})
}
