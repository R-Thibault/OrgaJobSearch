package routes

import "github.com/gin-gonic/gin"

// SetupRoutes sets up the API routes
func SetupRoutes(router *gin.Engine) {
	// Define a simple root route for health check
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server Start with succes !",
		})
	})
}
