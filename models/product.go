package models

import "time"

// Main Product structure
type Product struct {
	ID            int                `json:"id" gorm:"primaryKey"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	Price         float64            `json:"price"`
	CategoryID    int                `json:"category_id"`
	Category      Category           `json:"category" gorm:"foreignKey:CategoryID"`
	SubcategoryID int                `json:"subcategory_id"`
	Subcategory   Subcategory        `json:"subcategory" gorm:"foreignKey:SubcategoryID"`
	Brand         string             `json:"brand"`
	SKU           string             `json:"sku"`
	Attributes    []ProductAttribute `json:"attributes" gorm:"foreignKey:ProductID"`
	InStock       bool               `json:"in_stock"`
	StockQuantity int                `json:"stock_quantity"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
}

// Top-level categories (Electronics, Home & Furniture, etc.)
type Category struct {
	ID            int           `json:"id" gorm:"primaryKey"`
	Name          string        `json:"name"`
	Slug          string        `json:"slug"`
	Description   string        `json:"description"`
	Subcategories []Subcategory `json:"subcategories" gorm:"foreignKey:CategoryID"`
	IsActive      bool          `json:"is_active"`
	DisplayOrder  int           `json:"display_order"`
}

// Subcategories (Audio, Camera, Computer, etc. under Electronics)
type Subcategory struct {
	ID           int           `json:"id" gorm:"primaryKey"`
	CategoryID   int           `json:"category_id"`
	Name         string        `json:"name"`
	Slug         string        `json:"slug"`
	Description  string        `json:"description"`
	ProductTypes []ProductType `json:"product_types" gorm:"foreignKey:SubcategoryID"`
	IsActive     bool          `json:"is_active"`
	DisplayOrder int           `json:"display_order"`
}

// Product types (Bluetooth Headphones, Wired Headphones, etc. under Audio)
type ProductType struct {
	ID            int         `json:"id" gorm:"primaryKey"`
	SubcategoryID int         `json:"subcategory_id"`
	Name          string      `json:"name"`
	Slug          string      `json:"slug"`
	Description   string      `json:"description"`
	Attributes    []Attribute `json:"attributes" gorm:"foreignKey:ProductTypeID"`
	IsActive      bool        `json:"is_active"`
	DisplayOrder  int         `json:"display_order"`
}

// Product attributes template (defines what attributes products of this type have)
type Attribute struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	ProductTypeID int    `json:"product_type_id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	DataType      string `json:"data_type"` // string, number, boolean, etc.
	IsRequired    bool   `json:"is_required"`
	DisplayOrder  int    `json:"display_order"`
}

// Product-specific attribute values
type ProductAttribute struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	ProductID    int    `json:"product_id"`
	Name         string `json:"name"`
	Value        string `json:"value"`
	DisplayOrder int    `json:"display_order"`
}
