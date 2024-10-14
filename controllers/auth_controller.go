package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/config"
	"github.com/R-Thibault/OrgaJobSearch/models"
	"github.com/R-Thibault/OrgaJobSearch/services"
	"github.com/R-Thibault/OrgaJobSearch/utils"
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
	service *services.UserService
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(service *services.UserService) *AuthController {
	return &AuthController{service: service}
}

// SignIn handles the sign-in process
func (a *AuthController) SignIn(c *gin.Context) {
	var creds models.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}
	// check password requirement
	isMatch := utils.RegexPassword(creds.Password)
	if !isMatch {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
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

	// Verify the password
	isMatch, err := utils.CompareHashPassword(creds.Password, existingUser.HashedPassword)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if isMatch {
		fmt.Println("Password matches!")
		// Create JWT Token
		expirationTime := time.Now().Add(35 * time.Minute) // Extend the expiration time to 15 minutes, For demos purpose
		claims := &Claims{
			Email: creds.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		// Sign the token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
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
