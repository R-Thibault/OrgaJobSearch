package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/R-Thibault/OrgaJobSearch/backend/config"
	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	"github.com/dgrijalva/jwt-go"
)

type JWTTokenGeneratorService struct {
}

func NewJWTTokenGeneratorService() *JWTTokenGeneratorService {
	return &JWTTokenGeneratorService{}
}

func (u *JWTTokenGeneratorService) GenerateJWTToken(tokenType *string, body *string, expirationTime time.Time) (JWTToken string, err error) {

	var jwtKey = []byte(config.GetConfig("JWT_KEY"))
	// Set expiration
	// expirationTime := time.Now().Add(60 * time.Minute)

	token := models.JWTToken{
		TokenType: tokenType,
		Body:      body,
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
