package routes

import (
	"github.com/R-Thibault/OrgaJobSearch/backend/config"
	"github.com/R-Thibault/OrgaJobSearch/backend/controllers"
	otpRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/otp_repository"
	userRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/user_repository"
	"github.com/R-Thibault/OrgaJobSearch/backend/services"
	invitationServices "github.com/R-Thibault/OrgaJobSearch/backend/services/invitation_services"
	otpServices "github.com/R-Thibault/OrgaJobSearch/backend/services/otp_services"
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
	GenerateTokenService := tokenUtils.NewJWTTokenGeneratorService()
	OTPGeneratorService := otpGeneratorUtils.NewOtpGeneratorService()

	// Initialize Serivces
	UserService := userServices.NewUserService(userRepository, hashingService)
	OTPService := otpServices.NewOTPService(userRepository, OTPRepository, OTPGeneratorService)
	TokenService := tokenService.NewTokenService()
	MailerService := services.NewMailerService()
	invitationService := invitationServices.NewInvitationService(userRepository, OTPService)

	// Initialize Controllers
	authController := controllers.NewAuthController(UserService, hashingService, TokenService, invitationService, GenerateTokenService)
	userController := controllers.NewUserController(UserService, OTPService, TokenService)
	OTPcontroller := controllers.NewOTPController(OTPService, MailerService, UserService)
	userInvitationController := controllers.NewUserInvitationController(UserService, GenerateTokenService, *MailerService, OTPService)

	// Public route for signing in
	router.POST("/login", authController.SignIn)

	// Public route to check token from url invitation
	router.POST("/verify-token", authController.VerifyInvitationToken)

	// Public route for signing up
	router.POST("/sign-up", userController.SignUp)

	// Public route to generate OTP
	router.POST("/generate-otp", OTPcontroller.GenerateOTPForSignUp)

	// Public ( will be protected) route for send User invitation
	router.POST("/send-user-invitation", userInvitationController.SendJobSeekerInvitation)

	// Public ( will be protected) route for send Global invitation
	router.POST("/generate-url", userInvitationController.GenerateGlobalURLInvitation)

	//Public route for sending OTP
	router.POST("/send-otp", OTPcontroller.SendOTP)

	router.POST("/verify-otp", OTPcontroller.ValidateOTP)

}
