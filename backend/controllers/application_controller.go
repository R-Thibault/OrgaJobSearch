package controllers

import (
	"net/http"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	applicationservices "github.com/R-Thibault/OrgaJobSearch/backend/services/application_services"
	userServices "github.com/R-Thibault/OrgaJobSearch/backend/services/user_services"
	"github.com/gin-gonic/gin"
)

type ApplicationController struct {
	UserService        userServices.UserServiceInterface
	ApplicationService applicationservices.ApplicationServiceInterface
}

func NewApplicationController(UserService userServices.UserServiceInterface, ApplicationService applicationservices.ApplicationServiceInterface) *ApplicationController {
	return &ApplicationController{UserService: UserService, ApplicationService: ApplicationService}
}

func (app *ApplicationController) SaveApplication(c *gin.Context) {
	var appData models.Application
	if err := c.ShouldBindJSON(&appData); err != nil {
		// If the input is invalid, respond with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	userUUID, exists := c.Get("userUUID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserUUID not Found in context"})
		return
	}
	userUUIDStr, ok := userUUID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserUUID in context is not a a valid string"})
		return
	}

	existingUser, err := app.UserService.GetUserByUUID(userUUIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserUUID do not match a user"})
		return
	}
	appErr := app.ApplicationService.SaveApplication(existingUser.ID, appData)
	if appErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Saving application failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Application saved successfully"})
}
