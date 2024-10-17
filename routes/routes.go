package routes

import (
	"github.com/R-Thibault/OrgaJobSearch/config"
	"github.com/R-Thibault/OrgaJobSearch/controllers"
	"github.com/R-Thibault/OrgaJobSearch/repository"
	"github.com/R-Thibault/OrgaJobSearch/services"
	"github.com/R-Thibault/OrgaJobSearch/utils"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the API routes
func SetupRoutes(router *gin.Engine) {
	// Define a simple root route for health check
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server Start with succes !",
		})
	})

	// Initialize the repository and the hashing service
	userRepository := repository.NewUserRepository(config.DB)
	OTPRepository := repository.NewOTPRepository(config.DB)
	hashingService := utils.NewHashingService()
	OTPGeneratorService := utils.NewOtpGeneratorService()

	// Initialize the user service with the repository and hashing service
	userService := services.NewUserService(userRepository, hashingService)
	OTPService := services.NewOTPService(userRepository, OTPRepository, OTPGeneratorService)
	MailerService := services.NewMailerService()

	// Public route for signing in
	authController := controllers.NewAuthController(userService, hashingService)
	router.POST("/sign-in", authController.SignIn)

	// Public route for signing up
	userController := controllers.NewUserController(userService, OTPService)
	router.POST("/sign-up", userController.SignUp)

	// Public route to generate OTP
	OTPcontroller := controllers.NewOTPController(OTPService, MailerService)
	router.POST("/generate-otp", OTPcontroller.GenerateOTP)

	//Public route for sending OTP
	router.POST("/send-otp", OTPcontroller.SendOTP)
}
