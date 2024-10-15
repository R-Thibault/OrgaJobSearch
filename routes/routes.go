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
	hashingService := utils.NewHashingService()

	// Initialize the user service with the repository and hashing service
	userService := services.NewUserService(userRepository, hashingService)

	// Public route for signing in
	authController := controllers.NewAuthController(userService, hashingService)
	router.POST("/sign-in", authController.SignIn)

	// Public route for signing up
	userController := controllers.NewUserController(userService)
	router.POST("/sign-up", userController.SignUp)
}
