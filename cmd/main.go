package main

import (
	"log"

	"github.com/R-Thibault/OrgaJobSearch/config"
	"github.com/R-Thibault/OrgaJobSearch/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration and database connection
	config.SetupConfig()
	config.InitDB()
	defer config.CloseDB()

	// Create a new Gin engine instance
	r := gin.Default()
	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Adjust this based on your client origin
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Set-Cookie"},
		AllowCredentials: true, // Needed to allow cookies to be passed
	}))

	// Set up application routes
	routes.SetupRoutes(r)
	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur : %v", err)
	}
}
