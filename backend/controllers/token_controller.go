package controllers

import (
	"net/http"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	invitationServices "github.com/R-Thibault/OrgaJobSearch/backend/services/invitation_services"
	tokenService "github.com/R-Thibault/OrgaJobSearch/backend/services/token_services"
	"github.com/gin-gonic/gin"
)

type TokenController struct {
	tokenService      tokenService.TokenServiceInterface
	invitationService invitationServices.InvitationServiceInterface
}

// NewTokenController creates a new instance of AuthController
func NewTokenController(tokenService tokenService.TokenServiceInterface, invitationService invitationServices.InvitationServiceInterface) *AuthController {
	return &AuthController{
		tokenService:      tokenService,
		invitationService: invitationService,
	}
}

func (a *AuthController) VerifyInvitationToken(c *gin.Context) {
	var tokenString models.TokenRequest
	if err := c.ShouldBindJSON(&tokenString); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	token, err := a.tokenService.VerifyToken(tokenString.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization"})
		return
	}
	switch *token.TokenType {
	case "PersonalInvitation":
		email, err := a.invitationService.VerifyPersonnalInvitationTokenData(*token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"email":     email,
			"tokenType": *token.TokenType})
	case "GlobalInvitation":
		err := a.invitationService.VerifyGlobalInvitationTokenData(*token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"tokenType": *token.TokenType})
	default:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
}
