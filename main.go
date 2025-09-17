package main

import (
	"log"

	"github.com/Amitkumarsharma7082/product-api/database"
	"github.com/Amitkumarsharma7082/product-api/models"
	"github.com/Amitkumarsharma7082/product-api/server"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Initialize DB
	database.Connect()

	// Auto-migrate Book model to create the table
	database.DB.AutoMigrate(&models.Product{})

	// Setup routes
	server.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
