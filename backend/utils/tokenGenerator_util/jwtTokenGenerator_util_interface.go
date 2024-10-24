package utils

import "time"

type JWTTokenGeneratorServiceInterface interface {
	GenerateJWTToken(tokenType *string, email string, expirationTime time.Time) (JWTToken string, err error)
}
