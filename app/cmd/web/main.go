package main

import (
	"log"
	"os"

	_ "bzhspback.fr/breizhsport/docs" // This is where swag will generate docs
	"bzhspback.fr/breizhsport/internal/api/v1"
	"bzhspback.fr/breizhsport/internal/database"
	"bzhspback.fr/breizhsport/internal/seeders"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           BreizhSport Product Service API
// @version         1.0
// @description     This is the product service API for BreizhSport e-commerce platform
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.breizhsport.fr/support
// @contact.email  support@breizhsport.fr

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /api/v1
func main() {
	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Check if the seed command was provided
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		log.Println("Starting database seeding...")

		db := database.GetDB()

		// Run seeders in the correct order
		if err := seeders.SeedBrands(db); err != nil {
			log.Fatalf("Error seeding brands: %v", err)
		}
		log.Println("Brands seeded successfully")

		if err := seeders.SeedCategories(db); err != nil {
			log.Fatalf("Error seeding categories: %v", err)
		}
		log.Println("Categories seeded successfully")

		if err := seeders.SeedProducts(db); err != nil {
			log.Fatalf("Error seeding products: %v", err)
		}
		log.Println("Products seeded successfully")

		log.Println("Database seeding completed successfully")
		return
	}

	r := gin.Default()

	// Swagger documentation endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialize routes
	api.InitRoutes(r)

	// Start server
	r.Run(":8081")
}
