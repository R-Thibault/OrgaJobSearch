package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/config"
	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	invitationServices "github.com/R-Thibault/OrgaJobSearch/backend/services/invitation_services"
	tokenService "github.com/R-Thibault/OrgaJobSearch/backend/services/token_services"
	userServices "github.com/R-Thibault/OrgaJobSearch/backend/services/user_services"
	hashingUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/hash_util"
	JWTTokenGenerator "github.com/R-Thibault/OrgaJobSearch/backend/utils/tokenGenerator_util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var jwtKey = []byte(config.GetConfig("JWT_KEY"))

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// AuthController handles authentication-related requests
type AuthController struct {
	service           userServices.UserServiceInterface
	tokenService      tokenService.TokenServiceInterface
	invitationService invitationServices.InvitationServiceInterface
	hashingUtils      hashingUtils.HashingServiceInterface
	JWTTokenGenerator JWTTokenGenerator.JWTTokenGeneratorServiceInterface
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(service userServices.UserServiceInterface, hashingUtils hashingUtils.HashingServiceInterface, tokenService tokenService.TokenServiceInterface, invitationService invitationServices.InvitationServiceInterface, JWTTokenGenerator JWTTokenGenerator.JWTTokenGeneratorServiceInterface) *AuthController {
	return &AuthController{
		service:           service,
		hashingUtils:      hashingUtils,
		tokenService:      tokenService,
		invitationService: invitationService,
		JWTTokenGenerator: JWTTokenGenerator,
	}
}

// SignIn handles the login process
func (a *AuthController) SignIn(c *gin.Context) {
	var creds models.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	// User fetching logic
	existingUser, err := a.service.GetUserByEmail(creds.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	// check if user validate is email
	if existingUser.EmailIsValide == false {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	// Verify the password
	isMatch, err := a.hashingUtils.CompareHashPassword(creds.Password, existingUser.HashedPassword)
	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if isMatch {
		fmt.Println("Password matches!")
		// Create JWT Token
		expirationTime := time.Now().Add(24 * time.Hour)
		tokenType := "Cookie"
		tokenString, err := a.JWTTokenGenerator.GenerateJWTToken(&tokenType, nil, expirationTime)
		if err != nil {
			fmt.Printf("Failed to sign the token: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// Set the token in a cookie
		cookie := &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Path:     "/",
			Expires:  time.Now().Add(15 * time.Minute),
			HttpOnly: true,
			Secure:   false,                 // Set to true in production (HTTPS required if SameSite=None)
			SameSite: http.SameSiteNoneMode, // Required for cross-origin cookies
		}
		http.SetCookie(c.Writer, cookie)

		c.JSON(http.StatusOK, gin.H{"message": "Sign in successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
}

func (a *AuthController) VerifyInvitationToken(c *gin.Context) {
	var tokenString models.TokenRequest
	if err := c.ShouldBindJSON(&tokenString); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	token, err := a.tokenService.VerifyToken(tokenString.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization"})
		return
	}
	switch *token.TokenType {
	case "PersonalInvitation":
		email, err := a.invitationService.VerifyPersonnalInvitationTokenData(*token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"email":     email,
			"tokenType": *token.TokenType})
	case "GlobalInvitation":
		err := a.invitationService.VerifyGlobalInvitationTokenData(*token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"tokenType": *token.TokenType})
	default:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
}
