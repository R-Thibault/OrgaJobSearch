package routes

import (
	"github.com/R-Thibault/OrgaJobSearch/config"
	"github.com/R-Thibault/OrgaJobSearch/controllers"
	otpRepository "github.com/R-Thibault/OrgaJobSearch/repository/otp_repository"
	userRepository "github.com/R-Thibault/OrgaJobSearch/repository/user_repository"
	"github.com/R-Thibault/OrgaJobSearch/services"
	otpServices "github.com/R-Thibault/OrgaJobSearch/services/otp_services"
	userServices "github.com/R-Thibault/OrgaJobSearch/services/user_services"
	hashingUtils "github.com/R-Thibault/OrgaJobSearch/utils/hash_util"
	otpGeneratorUtils "github.com/R-Thibault/OrgaJobSearch/utils/otpGenerator_util"

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
	userRepository := userRepository.NewUserRepository(config.DB)
	OTPRepository := otpRepository.NewOTPRepository(config.DB)
	hashingService := hashingUtils.NewHashingService()
	OTPGeneratorService := otpGeneratorUtils.NewOtpGeneratorService()

	// Initialize the user service with the repository and hashing service
	UserService := userServices.NewUserService(userRepository, hashingService)
	OTPService := otpServices.NewOTPService(userRepository, OTPRepository, OTPGeneratorService)
	MailerService := services.NewMailerService()

	// Public route for signing in
	authController := controllers.NewAuthController(UserService, hashingService)
	router.POST("/sign-in", authController.SignIn)

	// Public route for signing up
	userController := controllers.NewUserController(UserService, OTPService)
	router.POST("/sign-up", userController.SignUp)

	// Public route to generate OTP
	OTPcontroller := controllers.NewOTPController(OTPService, MailerService, UserService)
	router.POST("/generate-otp", OTPcontroller.GenerateOTP)

	//Public route for sending OTP
	router.POST("/send-otp", OTPcontroller.SendOTP)

	router.POST("/verify-otp", OTPcontroller.ValidateOTP)
}
