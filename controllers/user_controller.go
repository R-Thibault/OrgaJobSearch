package controllers

import (
	"log"
	"net/http"

	"github.com/R-Thibault/OrgaJobSearch/models"
	"github.com/R-Thibault/OrgaJobSearch/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
	OTPService  *services.OTPService
}

func NewUserController(UserService *services.UserService, OTPService *services.OTPService) *UserController {
	return &UserController{UserService: UserService, OTPService: OTPService}
}

func (u *UserController) SendOTP(c *gin.Context) {
	var creds models.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		// If the input is invalid, respond with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	Otp, err := u.OTPService.GenerateOTP(creds.Email)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	log.Printf("%v\n", Otp)
	// Ici logique d'envoie d'email

	c.JSON(http.StatusOK, gin.H{"message": "OTP Successfully send"})
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
	err := u.UserService.RegisterUser(creds)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	// Respond with succes if no errors
	c.JSON(http.StatusOK, gin.H{"message": "User successfully registered"})
}
