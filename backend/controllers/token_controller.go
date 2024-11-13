package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	mailerService "github.com/R-Thibault/OrgaJobSearch/backend/services"
	otpServices "github.com/R-Thibault/OrgaJobSearch/backend/services/otp_services"
	tokenServices "github.com/R-Thibault/OrgaJobSearch/backend/services/token_services"
	userServices "github.com/R-Thibault/OrgaJobSearch/backend/services/user_services"
	tokenUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/tokenGenerator_util"
	"github.com/gin-gonic/gin"
)

type TokenController struct {
	TokenServices      tokenServices.TokenServiceInterface
	UserServices       userServices.UserServiceInterface
	OTPServices        otpServices.OTPServiceInterface
	TokenGeneratorUtil tokenUtils.JWTTokenGeneratorUtilInterface
	MailerService      mailerService.MailerService
}

// NewTokenController creates a new instance of AuthController
func NewTokenController(
	TokenServices tokenServices.TokenServiceInterface,
	UserServices userServices.UserServiceInterface,
	OTPServices otpServices.OTPServiceInterface,
	TokenGeneratorUtil tokenUtils.JWTTokenGeneratorUtilInterface,
	MailerService mailerService.MailerService) *TokenController {
	return &TokenController{
		TokenServices:      TokenServices,
		UserServices:       UserServices,
		OTPServices:        OTPServices,
		TokenGeneratorUtil: TokenGeneratorUtil,
		MailerService:      MailerService,
	}
}

func (t *TokenController) VerifyResetPasswordToken(c *gin.Context) {
	var tokenString models.TokenRequest
	if err := c.ShouldBindJSON(&tokenString); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	token, err := t.TokenServices.VerifyToken(tokenString.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization"})
		return
	}
	otpType := "resetPassword"
	otpSaved, otpErr := t.OTPServices.VerifyOTPCode(*token.Body, otpType)
	if otpErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}
	user, userErr := t.UserServices.GetUserByID(otpSaved.UserID)
	if userErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user do not exist"})
	}
	c.JSON(http.StatusOK, gin.H{"message": user.Email})
}

func (t *TokenController) SendResetPasswordEmail(c *gin.Context) {
	log.Print("TEST")
	var creds models.ResetPasswordCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	// if creds.Password != creds.ConfirmPassword {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Password and confirm do not match"})
	// 	return
	// }
	existingUser, err := t.UserServices.GetUserByEmail(creds.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't find user"})
		return
	}
	log.Printf("TEST :%v", creds)
	otpType := "resetPassword"
	JWTType := "resetPassword"
	expirationTime := time.Now().Add(2 * time.Hour)
	existingOtp, err := t.OTPServices.CheckAndRefreshOTPCode(existingUser.ID, otpType, expirationTime)
	if err != nil || existingOtp == "" {
		// Generate the OTP with a type "GlobalInvitation"
		otpGenerated, err := t.OTPServices.GenerateOTP(existingUser.ID, otpType, expirationTime)
		log.Printf("NO OTP : %v", otpGenerated)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		expirationTime := time.Now().Add(2 * time.Hour)
		jwtTokenString, err := t.TokenGeneratorUtil.GenerateJWTToken(JWTType, otpGenerated, expirationTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Problem during JWT Token generation"})
			return
		}
		mailerErr := t.MailerService.SendResetPasswordEmail(existingUser.Email, jwtTokenString)
		if mailerErr != nil {
			c.JSON(http.StatusConflict, gin.H{"error": mailerErr})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Email sent"})
	} else {
		// Use the existing valid OTP
		log.Printf("NO OTP : %v", existingOtp)
		expirationTime := time.Now().Add(2 * time.Hour)
		jwtTokenString, err := t.TokenGeneratorUtil.GenerateJWTToken(JWTType, existingOtp, expirationTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Problem during JWT Token generation"})
			return
		}
		mailerErr := t.MailerService.SendResetPasswordEmail(existingUser.Email, jwtTokenString)
		if mailerErr != nil {
			c.JSON(http.StatusConflict, gin.H{"error": mailerErr})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Email sent OTP"})
	}
}
