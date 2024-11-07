package controllers

import (
	"log"
	"net/http"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	otpServices "github.com/R-Thibault/OrgaJobSearch/backend/services/otp_services"
	registrationservices "github.com/R-Thibault/OrgaJobSearch/backend/services/registration_services"
	tokenService "github.com/R-Thibault/OrgaJobSearch/backend/services/token_services"
	userServices "github.com/R-Thibault/OrgaJobSearch/backend/services/user_services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService         userServices.UserServiceInterface
	OTPService          otpServices.OTPServiceInterface
	tokenService        tokenService.TokenServiceInterface
	Registrationservice registrationservices.RegistrationServiceInterface
}

func NewUserController(
	UserService userServices.UserServiceInterface,
	OTPService otpServices.OTPServiceInterface,
	tokenService tokenService.TokenServiceInterface,
	Registrationservice registrationservices.RegistrationServiceInterface) *UserController {
	return &UserController{
		UserService:         UserService,
		OTPService:          OTPService,
		tokenService:        tokenService,
		Registrationservice: Registrationservice}
}

func (u *UserController) SignUp(c *gin.Context) {
	// Parse the request body to extract credentials
	var creds models.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		// If the input is invalid, respond with an error
		log.Printf("ERROR : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	err := u.Registrationservice.UserRegistration(creds)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "User registration successful !"})

}

func (u *UserController) MyProfile(c *gin.Context) {
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

	existingUser, err := u.UserService.GetUserByUUID(userUUIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserUUID do not match a user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userEmail": existingUser.Email, "userFirstName": existingUser.FirstName, "userLastName": existingUser.LastName})
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var datas models.UserProfileUpdate
	if err := c.ShouldBindJSON(&datas); err != nil {
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

	existingUser, err := u.UserService.GetUserByUUID(userUUIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserUUID do not match a user"})
		return
	}

	updateErr := u.UserService.UpdateUser(*existingUser, datas)
	if updateErr != nil {
		c.JSON(http.StatusConflict, gin.H{"error": updateErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User update successful"})
}
