package tokenservices

import (
	"fmt"

	"github.com/R-Thibault/OrgaJobSearch/backend/config"
	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	"github.com/dgrijalva/jwt-go"
)

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (s *TokenService) VerifyToken(tokenString string) (*models.JWTToken, error) {
	// load JWT key from configuration
	jwtKey := []byte(config.GetConfig("JWT_KEY"))

	// Parse token with claims of type JWTToken
	parsedToken, err := jwt.ParseWithClaims(tokenString, &models.JWTToken{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		// provid secret key used to sign the token
		return jwtKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}
	// Validate the token and ensure it's valid
	if claims, ok := parsedToken.Claims.(*models.JWTToken); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
