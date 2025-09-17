package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	DB = db
	fmt.Println("Database connection established")
}

//type Database struct {
//}
//
//func InitializeDatabase() *Database {
//	initializeSampleData()
//	return &Database{}
//}
//
//func initializeSampleData() {
//	now := time.Now()
//
//	models.Products[1] = models.Product{
//		ID:          1,
//		Name:        "Laptop",
//		Description: "High-performance gaming laptop",
//		Price:       999.99,
//		Category:    "Electronics",
//		InStock:     true,
//		CreatedAt:   now,
//		UpdatedAt:   now,
//	}
//
//	models.Products[2] = models.Product{
//		ID:          2,
//		Name:        "Smartphone",
//		Description: "Latest smartphone with amazing camera",
//		Price:       699.99,
//		Category:    "Electronics",
//		InStock:     true,
//		CreatedAt:   now,
//		UpdatedAt:   now,
//	}
//	models.Products[3] = models.Product{
//		ID:          3,
//		Name:        "Headphones",
//		Description: "Wireless noise-cancelling headphones",
//		Price:       199.99,
//		Category:    "Audio",
//		InStock:     false,
//		CreatedAt:   now,
//		UpdatedAt:   now,
//	}
//
//	models.LastID = 3
//}
