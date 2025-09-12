package database

import (
	"time"

	"github.com/Amitkumarsharma7082/product-api/models"
)

type Database struct {
}

func InitializeDatabase() *Database {
	initializeSampleData()
	return &Database{}
}

func initializeSampleData() {
	now := time.Now()

	models.Products[1] = models.Product{
		ID:          1,
		Name:        "Laptop",
		Description: "High-performance gaming laptop",
		Price:       999.99,
		Category:    "Electronics",
		InStock:     true,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	models.Products[2] = models.Product{
		ID:          2,
		Name:        "Smartphone",
		Description: "Latest smartphone with amazing camera",
		Price:       699.99,
		Category:    "Electronics",
		InStock:     true,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	models.Products[3] = models.Product{
		ID:          3,
		Name:        "Headphones",
		Description: "Wireless noise-cancelling headphones",
		Price:       199.99,
		Category:    "Audio",
		InStock:     false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	models.LastID = 3
}
