package main

import (
	"log"

	"github.com/R-Thibault/Go----Boilerplate-.git/config"
	"github.com/R-Thibault/Go----Boilerplate-.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration and database connection
	config.SetupConfig()
	config.InitDB()
	defer config.CloseDB()

	// Create a new Gin engine instance
	r := gin.Default()

	// Set up application routes
	routes.SetupRoutes(r)
	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur : %v", err)
	}
}
