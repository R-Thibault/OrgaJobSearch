package controllers

import (
	"net/http"
	"time"

	"github.com/R-Thibault/Go----Boilerplate-.git/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(config.GetConfig("JWT_KEY"))

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type AuthController struct{}

// SignIn handles the sign-in process
func (a *AuthController) SignIn(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	// User fetching logic here
	expectedPasswordHash := "$2a$12$ExampleHashFromBCryptForDemoPurposesOnly"

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(expectedPasswordHash), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create JWT Token
	expirationTime := time.Now().Add(15 * time.Minute) // Extend the expiration time to 15 minutes, For demos purpose
	claims := &Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Sign the token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Set the token in a cookie
	c.SetCookie("token", tokenString, int(expirationTime.Unix()-time.Now().Unix()), "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Sign in succesfull"})
}
