package controllers

import (
	"fmt"
	"log"
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
func (a *AuthController) Login(c *gin.Context) {
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
	if err != nil || !isMatch {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	fmt.Println("Password matches!")
	// Create JWT Token
	expirationTime := time.Now().Add(24 * time.Hour)
	tokenType := "Cookie"
	userUUID := string(existingUser.UserUUID)
	tokenString, err := a.JWTTokenGenerator.GenerateJWTToken(&tokenType, &userUUID, expirationTime)
	if err != nil {
		fmt.Printf("Failed to sign the token: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.SetCookie("token", tokenString, 3600, "/", "localhost", true, false)

	c.JSON(http.StatusOK, gin.H{"message": "Sign in successful"})

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

func (a *AuthController) Logout(c *gin.Context) {
	// Log all cookies present in the request for debugging
	cookies := c.Request.Cookies()
	if len(cookies) == 0 {
		log.Println("No cookies received in the request.")
	} else {
		for _, cookie := range cookies {
			log.Printf("Cookie Name: %s, Cookie Value: %s\n", cookie.Name, cookie.Value)
		}
	}
	cookie, err := c.Cookie("token")
	if err != nil {
		log.Printf("err: %v", err)
		log.Printf("Cookie: %v", cookie)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
		return
	}
	log.Printf("Cookie: %v", cookie)
	// Clear the cookie by setting its expiration date to the past
	c.SetCookie("token", "", -1, "/", "localhost", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
