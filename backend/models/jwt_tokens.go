package models

import "github.com/dgrijalva/jwt-go"

type JWTToken struct {
	TokenType *string `json:"tokenType"`
	Body      *string `json:"body"`
	jwt.StandardClaims
}
