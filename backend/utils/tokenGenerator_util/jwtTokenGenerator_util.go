package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/config"
	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	"github.com/dgrijalva/jwt-go"
)

type JWTTokenGeneratorUtil struct {
}

func NewJWTTokenGeneratorUtil() *JWTTokenGeneratorUtil {
	return &JWTTokenGeneratorUtil{}
}

// GenerateJWTToken generates a JWT token with the specified type, body, and expiration time.
// It uses the JWT key from the configuration to sign the token.
//
// Parameters:
//   - tokenType: A string representing the type of the token.
//   - body: A string representing the body of the token.
//   - expirationTime: A time.Time value representing the expiration time of the token.
//
// Returns:
//   - JWTToken: A string representing the generated JWT token.
//   - err: An error if the token generation or signing fails.
func (u *JWTTokenGeneratorUtil) GenerateJWTToken(tokenType string, body string, expirationTime time.Time) (JWTToken string, err error) {

	var jwtKey = []byte(config.GetConfig("JWT_KEY"))
	// Set expiration
	// expirationTime := time.Now().Add(60 * time.Minute)

	token := models.JWTToken{
		TokenType: &tokenType,
		Body:      &body,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Sign the token
	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	tokenString, err := generatedToken.SignedString(jwtKey)
	if err != nil {
		fmt.Printf("Failed to sign the token: %v\n", err)
		return "", errors.New("Failed to sign the token")
	}
	return tokenString, nil
}
