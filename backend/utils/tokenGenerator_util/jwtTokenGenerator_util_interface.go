package utils

import "time"

type JWTTokenGeneratorServiceInterface interface {
	GenerateJWTToken(tokenType *string, uuid *string, expirationTime time.Time) (JWTToken string, err error)
}
