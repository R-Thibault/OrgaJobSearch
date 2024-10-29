package utils

import "time"

type JWTTokenGeneratorUtilInterface interface {
	GenerateJWTToken(tokenType *string, body *string, expirationTime time.Time) (JWTToken string, err error)
}
