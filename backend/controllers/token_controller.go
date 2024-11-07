package controllers

import (
	tokenService "github.com/R-Thibault/OrgaJobSearch/backend/services/token_services"
	"github.com/gin-gonic/gin"
)

type TokenController struct {
	tokenService tokenService.TokenServiceInterface
}

// NewTokenController creates a new instance of AuthController
func NewTokenController(tokenService tokenService.TokenServiceInterface) *AuthController {
	return &AuthController{
		tokenService: tokenService,
	}
}

func (a *AuthController) VerifyToken(c *gin.Context) {
	// var tokenString models.TokenRequest
	// if err := c.ShouldBindJSON(&tokenString); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	// 	return
	// }
	// token, err := a.tokenService.VerifyToken(tokenString.Token)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization"})
	// 	return
	// }

}
