package router

import (
	"github.com/Amitkumarsharma7082/product-api/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Routes
	app.Get("/products", controllers.GetProducts)
	app.Get("/products/:id", controllers.GetProduct)
	app.Post("/products", controllers.CreateProduct)
	app.Delete("/products/:id", controllers.DeleteProduct)
}
