package tokenservices

import "github.com/R-Thibault/OrgaJobSearch/backend/models"

type TokenServiceInterface interface {
	VerifyToken(tokenString string) (*models.JWTToken, error)
}
