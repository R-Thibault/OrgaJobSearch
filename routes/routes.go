package routes

import (
	"github.com/R-Thibault/OrgaJobSearch/config"
	"github.com/R-Thibault/OrgaJobSearch/controllers"
	"github.com/R-Thibault/OrgaJobSearch/repository"
	"github.com/R-Thibault/OrgaJobSearch/services"
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

	// Public route for signing in
	authController := controllers.NewAuthController(services.NewUserService(repository.NewUserRepository(config.DB)))
	router.POST("/sign-in", authController.SignIn)

	// Public route for signing up
	userController := controllers.NewUserController(services.NewUserService(repository.NewUserRepository(config.DB)))
	router.POST("/sign-up", userController.SignUp)
}
