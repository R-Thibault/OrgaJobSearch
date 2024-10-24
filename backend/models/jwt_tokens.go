package models

import "github.com/dgrijalva/jwt-go"

type JWTToken struct {
	TokenType *string `json:"tokenType"`
	Email     string  `json:"email"`
	jwt.StandardClaims
}
