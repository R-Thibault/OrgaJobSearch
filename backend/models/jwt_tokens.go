package models

import "github.com/dgrijalva/jwt-go"

type JWTToken struct {
	UserID *uint  `json:"userID"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
