package router

import (
	"github.com/Amitkumarsharma7082/product-api/internal/products"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Routes
	app.Get("/products", products.GetProducts)
	app.Get("/products/:id", products.GetProduct)
	app.Post("/products", products.CreateProduct)
	app.Delete("/products/:id", products.DeleteProduct)
}
