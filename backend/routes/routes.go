package routes

// © Rossa Thibault 2024. Tous droits réservés.
// Ce code est la propriété de Rossa Thibault et ne peut être utilisé,
// distribué ou modifié sans autorisation explicite.
import (
	"github.com/R-Thibault/OrgaJobSearch/backend/config"
	"github.com/R-Thibault/OrgaJobSearch/backend/controllers"
	"github.com/R-Thibault/OrgaJobSearch/backend/middleware"
	applicationrepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/application_repository"
	otpRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/otp_repository"
	userRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/user_repository"
	"github.com/R-Thibault/OrgaJobSearch/backend/services"
	applicationservices "github.com/R-Thibault/OrgaJobSearch/backend/services/application_services"
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
	UserRepository := userRepository.NewUserRepository(config.DB)
	OTPRepository := otpRepository.NewOTPRepository(config.DB)
	ApplicationRepository := applicationrepository.NewApplicationRepository(config.DB)

	// Initialize Utilities
	HashingService := hashingUtils.NewHashingService()
	GenerateTokenService := tokenUtils.NewJWTTokenGeneratorUtil()
	OTPGeneratorService := otpGeneratorUtils.NewOtpGeneratorService()

	// Initialize Serivces
	UserService := userServices.NewUserService(UserRepository, OTPRepository, HashingService)
	OTPService := otpServices.NewOTPService(UserRepository, OTPRepository, OTPGeneratorService)
	TokenService := tokenService.NewTokenService()
	MailerService := services.NewMailerService()
	RegistrationService := registrationservices.NewRegistrationService(UserRepository, HashingService)
	ApplicationService := applicationservices.NewApplicationService(ApplicationRepository)

	// Initialize Controllers
	AuthController := controllers.NewAuthController(UserService, HashingService, TokenService, GenerateTokenService)
	UserController := controllers.NewUserController(UserService, OTPService, TokenService, RegistrationService)
	OTPController := controllers.NewOTPController(OTPService, MailerService, UserService)
	TokenController := controllers.NewTokenController(TokenService, UserService, OTPService, GenerateTokenService, *MailerService)

	ApplicationController := controllers.NewApplicationController(UserService, ApplicationService)

	// Public routes
	router.POST("/login", AuthController.Login)
	router.POST("/logout", AuthController.Logout)
	router.POST("/sign-up", UserController.SignUp)
	router.POST("/generate-otp", OTPController.GenerateOTPForSignUp)
	router.POST("/send-otp", OTPController.SendOTP) // Not use on frontend only called on backend
	router.POST("/verify-otp", OTPController.ValidateOTPForSignUp)
	router.POST("/reset-password", UserController.ResetPassword)
	router.POST("/send-reset-password-link", TokenController.SendResetPasswordEmail)
	router.POST("/verify-reset-password-link", TokenController.VerifyResetPasswordToken)

	// Protected route
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/me", UserController.MyProfile)
	protected.POST("/update-user", UserController.UpdateUser)
	protected.POST("/create-application", ApplicationController.SaveApplication)
	protected.POST("/get-applications-by-user", ApplicationController.GetApplicationsByUserID)
	protected.POST("/update-application")
	protected.POST("/delete-application")        //Soft delete with Gorm
	protected.POST("/update-application-status") //For updating only app status on dashboard
	// protected.POST("/reset-password", UserController.ResetPassword)

}
