package utils

import "time"

type JWTTokenGeneratorServiceInterface interface {
	GenerateJWTToken(tokenType *string, body *string, expirationTime time.Time) (JWTToken string, err error)
}
