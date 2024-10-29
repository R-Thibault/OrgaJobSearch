package controllers

import (
	"encoding/json"
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
	JWTTokenGenerator JWTTokenGenerator.JWTTokenGeneratorUtilInterface
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(service userServices.UserServiceInterface, hashingUtils hashingUtils.HashingServiceInterface, tokenService tokenService.TokenServiceInterface, invitationService invitationServices.InvitationServiceInterface, JWTTokenGenerator JWTTokenGenerator.JWTTokenGeneratorUtilInterface) *AuthController {
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

	// Create JWT Token
	expirationTime := time.Now().Add(24 * time.Hour)
	tokenType := "Cookie"
	userUUID := string(existingUser.UserUUID)
	userRole := existingUser.Roles
	bodyContent := map[string]string{
		"userUUID": userUUID,
		"userRole": fmt.Sprintf("%v", userRole),
	}

	// Encode `bodyContent` to JSON
	bodyBytes, err := json.Marshal(bodyContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode body content"})
		return
	}
	bodyString := string(bodyBytes)
	tokenString, err := a.JWTTokenGenerator.GenerateJWTToken(&tokenType, &bodyString, expirationTime)
	if err != nil {
		fmt.Printf("Failed to sign the token: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.SetCookie("token", tokenString, 36000, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Sign in successful"})

}

func (a *AuthController) Logout(c *gin.Context) {

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
