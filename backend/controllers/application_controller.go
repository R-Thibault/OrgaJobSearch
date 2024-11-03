package controllers

import (
	"log"
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
		log.Printf("APPDATA: %v", appData)
		log.Printf("ERROR: %v", err)
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

func (app *ApplicationController) UpdateApplication(c *gin.Context) {
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
	UpdatedApplication, err := app.ApplicationService.UpdateApplication(existingUser.ID, appData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during application update"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"UpdatedApplication": UpdatedApplication})
}

func (app *ApplicationController) GetApplicationByID(c *gin.Context) {
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

	application, appErr := app.ApplicationService.GetApplicationByID(existingUser.ID, appData.ID)
	if appErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't get application informations"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": application})
}

func (app *ApplicationController) GetApplicationsByUserID(c *gin.Context) {
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
	applications, err := app.ApplicationService.GetApplicationsByUserID(existingUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't find applications for this user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": applications})
}

func (app *ApplicationController) DeleteApplication(c *gin.Context) {
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
	deleteErr := app.ApplicationService.DeleteApplication(existingUser.ID, appData.ID)
	if deleteErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during application supression"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "application delete successfully"})
}
