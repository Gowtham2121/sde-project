package models

import "time"

// Product represents a product entity
type Product struct {
	ID                      int       `json:"id"`
	UserID                  int       `json:"user_id"`
	ProductName             string    `json:"product_name"`
	ProductDescription      string    `json:"product_description"`
	ProductImages           []string  `json:"product_images"`
	CompressedProductImages []string  `json:"compressed_product_images"`
	ProductPrice            float64   `json:"product_price"`
	CreatedAt               time.Time `json:"created_at"`
}
