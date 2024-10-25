package controllers

import (
	"net/http"

	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	otpServices "github.com/R-Thibault/OrgaJobSearch/backend/services/otp_services"
	tokenService "github.com/R-Thibault/OrgaJobSearch/backend/services/token_services"
	userServices "github.com/R-Thibault/OrgaJobSearch/backend/services/user_services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService  *userServices.UserService
	OTPService   *otpServices.OTPService
	tokenService tokenService.TokenServiceInterface
}

func NewUserController(UserService *userServices.UserService, OTPService *otpServices.OTPService, tokenService tokenService.TokenServiceInterface) *UserController {
	return &UserController{UserService: UserService, OTPService: OTPService, tokenService: tokenService}
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
	c.JSON(http.StatusOK, gin.H{"message": creds.Email})
}

func (u *UserController) MyProfil(c *gin.Context) {
	//extract cookie
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentification required"})

		return
	}
	tokenClaims, err := u.tokenService.VerifyToken(cookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization"})
		return
	}
	user, err := u.UserService.UserRepo.GetUserByUUID(*tokenClaims.Body)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"userName": user.Name, "email": user.Email, "role": user.Roles})
}
