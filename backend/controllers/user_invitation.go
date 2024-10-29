package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	mailerService "github.com/R-Thibault/OrgaJobSearch/backend/services"
	otpServices "github.com/R-Thibault/OrgaJobSearch/backend/services/otp_services"
	registrationservices "github.com/R-Thibault/OrgaJobSearch/backend/services/registration_services"
	userServices "github.com/R-Thibault/OrgaJobSearch/backend/services/user_services"
	tokenUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/tokenGenerator_util"
	"github.com/gin-gonic/gin"
)

type UserInvitationController struct {
	UserService         userServices.UserServiceInterface
	TokenGeneratorUtil  tokenUtils.JWTTokenGeneratorUtilInterface
	MailerService       mailerService.MailerService
	otpServices         otpServices.OTPServiceInterface
	Registrationservice registrationservices.RegistrationServiceInterface
}

func NewUserInvitationController(
	UserService userServices.UserServiceInterface,
	TokenGeneratorUtil tokenUtils.JWTTokenGeneratorUtilInterface,
	MailerService mailerService.MailerService,
	otpServices otpServices.OTPServiceInterface,
	Registrationservice registrationservices.RegistrationServiceInterface) *UserInvitationController {
	return &UserInvitationController{
		UserService:         UserService,
		TokenGeneratorUtil:  TokenGeneratorUtil,
		MailerService:       MailerService,
		otpServices:         otpServices,
		Registrationservice: Registrationservice}
}

func (u *UserInvitationController) SendJobSeekerInvitation(c *gin.Context) {
	var userInvitation models.UserInvitation
	if err := c.ShouldBindJSON(&userInvitation); err != nil {
		// If the input is invalid, respond with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	userUUID, exists := c.Get("userUUID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserUUID not Found in context"})
		return
	}

	userUUIDStr, ok := userUUID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserUUID in context is not a a valid string"})
		return
	}
	existingUser, err := u.UserService.GetUserByUUID(userUUIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserUUID do not match a user"})
		return
	}

	savedUser, err := u.Registrationservice.PreRegisterJobSeeker(userInvitation.Email, &existingUser.ID)
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
	if err := c.ShouldBindJSON(&invitation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	userUUID, exists := c.Get("userUUID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserUUID not found in context"})
		return
	}

	userUUIDStr, ok := userUUID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserUUID in context is not a valid string"})
		return
	}

	existingUser, err := u.UserService.GetUserByUUID(userUUIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserUUID does not match a user"})
		return
	}

	// Check if an OTP already exists and is still valid
	otpType := "GlobalInvitation"
	existingOtp, err := u.otpServices.CheckOTPCodeForGlobalInvitation(existingUser.ID, otpType)
	if err != nil || existingOtp == "" {
		// Generate the OTP with a type "GlobalInvitation"
		otpGenerated, err := u.otpServices.GenerateOTP(existingUser.ID, otpType)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		invitation.InvitationType = otpType
		expirationTime := time.Now().Add(8 * time.Hour)

		jwtTokenString, err := u.TokenGeneratorUtil.GenerateJWTToken(&invitation.InvitationType, &otpGenerated, expirationTime)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		url := fmt.Sprintf("http://localhost:3000/sign-up?token=%s", jwtTokenString)
		c.JSON(http.StatusOK, gin.H{"url": url})
	} else {
		// Use the existing valid OTP
		invitation.InvitationType = otpType
		expirationTime := time.Now().Add(8 * time.Hour)

		jwtTokenString, err := u.TokenGeneratorUtil.GenerateJWTToken(&invitation.InvitationType, &existingOtp, expirationTime)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		url := fmt.Sprintf("http://localhost:3000/sign-up?token=%s", jwtTokenString)
		c.JSON(http.StatusOK, gin.H{"url": url})
	}
}
