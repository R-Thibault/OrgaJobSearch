package main

// © Rossa Thibault 2024. Tous droits réservés.
// Ce code est la propriété de Rossa Thibault et ne peut être utilisé,
// distribué ou modifié sans autorisation explicite.
import (
	"log"

	"github.com/R-Thibault/OrgaJobSearch/backend/config"
	"github.com/R-Thibault/OrgaJobSearch/backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration and database connection
	config.SetupConfig()
	config.InitDB()
	defer config.CloseDB()

	if err := config.LoadData(config.GetDB()); err != nil {
		log.Fatalf("Error loading initial data: %v", err)
	}

	if err := config.SeedDatabaseWithUsers(config.GetDB()); err != nil {
		log.Fatalf("Error loading initial data: %v", err)
	}
	if err := config.SeedDatabaseWithApplications(config.GetDB()); err != nil {
		log.Fatalf("Error loading initial data: %v", err)
	}
	// Create a new Gin engine instance
	r := gin.Default()
	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Adjust this based on your client origin
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Cookie"},
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
