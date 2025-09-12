package main

import (
	"log"

	"github.com/Amitkumarsharma7082/product-api/database"
	"github.com/Amitkumarsharma7082/product-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Initialize DB
	database.InitializeDatabase()

	// Setup routes
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
