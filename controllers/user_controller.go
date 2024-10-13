package controllers

import (
	"net/http"

	"github.com/R-Thibault/OrgaJobSearch/models"
	"github.com/R-Thibault/OrgaJobSearch/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (u *UserController) SignUp(c *gin.Context) {
	// Parse the request body to extract credentials
	var creds models.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		// If the input is invalid, respond with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Call the service to register the user
	err := u.service.RegisterUser(creds)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	// Respond with succes if no errors
	c.JSON(http.StatusOK, gin.H{"message": "User successfully registered"})
}
