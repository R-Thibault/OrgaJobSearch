package utils

import "time"

type JWTTokenGeneratorServiceInterface interface {
	GenerateJWTToken(userID *uint, email string, expirationTime time.Time) (JWTToken string, err error)
}
