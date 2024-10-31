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
	rolerepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/role_repository"
	userRepository "github.com/R-Thibault/OrgaJobSearch/backend/repository/user_repository"
	"github.com/R-Thibault/OrgaJobSearch/backend/services"
	applicationservices "github.com/R-Thibault/OrgaJobSearch/backend/services/application_services"
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
	UserRepository := userRepository.NewUserRepository(config.DB)
	OTPRepository := otpRepository.NewOTPRepository(config.DB)
	RoleRepository := rolerepository.NewRoleRepository(config.DB)
	ApplicationRepository := applicationrepository.NewApplicationRepository(config.DB)

	// Initialize Utilities
	HashingService := hashingUtils.NewHashingService()
	GenerateTokenService := tokenUtils.NewJWTTokenGeneratorUtil()
	OTPGeneratorService := otpGeneratorUtils.NewOtpGeneratorService()

	// Initialize Serivces
	UserService := userServices.NewUserService(UserRepository, HashingService)
	OTPService := otpServices.NewOTPService(UserRepository, OTPRepository, OTPGeneratorService)
	TokenService := tokenService.NewTokenService()
	MailerService := services.NewMailerService()
	InvitationService := invitationServices.NewInvitationService(UserRepository, OTPService)
	RegistrationService := registrationservices.NewRegistrationService(UserRepository, HashingService, RoleRepository)
	ApplicationService := applicationservices.NewApplicationService(ApplicationRepository)

	// Initialize Controllers
	AuthController := controllers.NewAuthController(UserService, HashingService, TokenService, InvitationService, GenerateTokenService)
	UserController := controllers.NewUserController(UserService, OTPService, TokenService, RegistrationService)
	OTPcontroller := controllers.NewOTPController(OTPService, MailerService, UserService)
	UserInvitationController := controllers.NewUserInvitationController(UserService, GenerateTokenService, *MailerService, OTPService, RegistrationService)
	ApplicationController := controllers.NewApplicationController(UserService, ApplicationService)

	// Public routes
	router.POST("/login", AuthController.Login)
	router.POST("/logout", AuthController.Logout)
	router.POST("/verify-token", AuthController.VerifyInvitationToken)
	router.POST("/sign-up", UserController.SignUp)
	router.POST("/generate-otp", OTPcontroller.GenerateOTPForSignUp)
	router.POST("/send-otp", OTPcontroller.SendOTP)
	router.POST("/verify-otp", OTPcontroller.ValidateOTP)

	// Protected route
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.POST("/send-user-invitation", middleware.RoleMiddleware("CareerCoach", "CareerSupportManager"), UserInvitationController.SendJobSeekerInvitation)
	protected.POST("/generate-url", middleware.RoleMiddleware("CareerSupportManager"), UserInvitationController.GenerateGlobalURLInvitation)
	protected.GET("/me", middleware.RoleMiddleware("JobSeeker", "CareerCoach", "CareerSupportManager"), UserController.MyProfile)
	protected.POST("/update-user", middleware.RoleMiddleware("JobSeeker", "CareerCoach", "CareerSupportManager"), UserController.UpdateUser)
	protected.POST("/application", middleware.RoleMiddleware("JobSeeker", "CareerCoach", "CareerSupportManager"), ApplicationController.SaveApplication)
}
