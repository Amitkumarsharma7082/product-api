package server

import (
	"github.com/Amitkumarsharma7082/product-api/internal/products"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	// Routes
	api.Get("/products", products.GetProducts)
	api.Get("/products/:id", products.GetProduct)
	api.Post("/products", products.CreateProduct)
	api.Put("/products/:id", products.UpdateProduct)
	api.Delete("/products/:id", products.DeleteProduct)
}
