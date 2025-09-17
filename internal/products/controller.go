package products

import (
	"github.com/Amitkumarsharma7082/product-api/database"
	"github.com/Amitkumarsharma7082/product-api/models"
	"github.com/gofiber/fiber/v2"
)

// All
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	database.DB.Find(&products)
	return c.JSON(products)
}

// Product by id
func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	result := database.DB.First(&product, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(product)
}

// create
func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Create(&product)
	return c.JSON(product)
}

// Put
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	database.DB.Save(&product)
	return c.JSON(product)
}

// delete
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	result := database.DB.Delete(&models.Product{}, id)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(fiber.Map{"message": "Product deleted successfully"})
}
