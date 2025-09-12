package controllers

import (
	"strconv"

	"github.com/Amitkumarsharma7082/product-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

// All
func GetProducts(c *fiber.Ctx) error {
	products := make([]models.Product, 0, len(models.Products))
	for _, product := range models.Products {
		products = append(products, product)
	}
	return c.JSON(products)
}

// Product by id
func GetProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	product, exists := models.Products[id]
	if !exists {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(product)
}

// create
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	models.LastID++
	product.ID = models.LastID
	models.Products[product.ID] = product

	return c.Status(201).JSON(product)
}

// delete
func DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Check if user exists
	if _, exists := models.Products[id]; !exists {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}

	delete(models.Products, id)
	return c.Status(204).SendString("")
}
