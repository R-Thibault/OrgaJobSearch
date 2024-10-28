package routes

import (
	"github.com/R-Thibault/OrgaJobSearch/backend/config"
	"github.com/R-Thibault/OrgaJobSearch/backend/controllers"
	"github.com/R-Thibault/OrgaJobSearch/backend/middleware"
	otpRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/otp_repository"
	userRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/user_repository"
	"github.com/R-Thibault/OrgaJobSearch/backend/services"
	invitationServices "github.com/R-Thibault/OrgaJobSearch/backend/services/invitation_services"
	otpServices "github.com/R-Thibault/OrgaJobSearch/backend/services/otp_services"
	registrationservices "github.com/R-Thibault/OrgaJobSearch/backend/services/registration_services"
	tokenService "github.com/R-Thibault/OrgaJobSearch/backend/services/token_services"
	userServices "github.com/R-Thibault/OrgaJobSearch/backend/services/user_services"
	hashingUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/hash_util"
	otpGeneratorUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/otpGenerator_util"
	tokenUtils "github.com/R-Thibault/OrgaJobSearch/backend/utils/tokenGenerator_util"

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

	// Initialize repositories
	userRepository := userRepository.NewUserRepository(config.DB)
	OTPRepository := otpRepository.NewOTPRepository(config.DB)

	// Initialize Utilities
	hashingService := hashingUtils.NewHashingService()
	GenerateTokenService := tokenUtils.NewJWTTokenGeneratorUtil()
	OTPGeneratorService := otpGeneratorUtils.NewOtpGeneratorService()

	// Initialize Serivces
	UserService := userServices.NewUserService(userRepository, hashingService)
	OTPService := otpServices.NewOTPService(userRepository, OTPRepository, OTPGeneratorService)
	TokenService := tokenService.NewTokenService()
	MailerService := services.NewMailerService()
	invitationService := invitationServices.NewInvitationService(userRepository, OTPService)
	RegistrationService := registrationservices.NewRegistrationService(userRepository, hashingService)

	// Initialize Controllers
	authController := controllers.NewAuthController(UserService, hashingService, TokenService, invitationService, GenerateTokenService)
	userController := controllers.NewUserController(UserService, OTPService, TokenService, RegistrationService)
	OTPcontroller := controllers.NewOTPController(OTPService, MailerService, UserService)
	userInvitationController := controllers.NewUserInvitationController(UserService, GenerateTokenService, *MailerService, OTPService, RegistrationService)

	// Public routes
	router.POST("/login", authController.Login)
	router.POST("/logout", authController.Logout)
	router.POST("/verify-token", authController.VerifyInvitationToken)
	router.POST("/sign-up", userController.SignUp)
	router.POST("/generate-otp", OTPcontroller.GenerateOTPForSignUp)
	router.POST("/send-otp", OTPcontroller.SendOTP)
	router.POST("/verify-otp", OTPcontroller.ValidateOTP)

	// Protected route
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.POST("/send-user-invitation", userInvitationController.SendJobSeekerInvitation)
	protected.POST("/generate-url", userInvitationController.GenerateGlobalURLInvitation)
	protected.GET("/me", userController.MyProfile)
}
