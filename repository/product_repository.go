package repository

import (
	"log"
	"product-management-system/config"
	"product-management-system/models"

	"github.com/lib/pq" // Import pq for PostgreSQL array support
)

// CreateProduct inserts a new product record into the database
func CreateProduct(product *models.Product) error {
	query := `
        INSERT INTO products (user_id, product_name, product_description, product_images, compressed_product_images, product_price, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP)
        RETURNING id;
    `
	err := config.DB.QueryRow(query, product.UserID, product.ProductName, product.ProductDescription,
		pq.Array(product.ProductImages), pq.Array(product.CompressedProductImages), product.ProductPrice).Scan(&product.ID)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		return err
	}
	return nil
}

// GetProductByID retrieves a product by its ID
func GetProductByID(productID int) (*models.Product, error) {
	product := &models.Product{}
	query := `
        SELECT id, user_id, product_name, product_description, product_images, compressed_product_images, product_price, created_at
        FROM products
        WHERE id = $1;
    `
	row := config.DB.QueryRow(query, productID)
	err := row.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription,
		pq.Array(&product.ProductImages), pq.Array(&product.CompressedProductImages), &product.ProductPrice, &product.CreatedAt)
	if err != nil {
		log.Printf("Error fetching product by ID: %v", err)
		return nil, err
	}
	return product, nil
}

// GetProductsByUserID retrieves all products for a specific user, with optional filtering
func GetProductsByUserID(userID int) ([]models.Product, error) {
	var products []models.Product
	query := `
        SELECT id, user_id, product_name, product_description, product_images, compressed_product_images, product_price, created_at
        FROM products
        WHERE user_id = $1;
    `
	rows, err := config.DB.Query(query, userID)
	if err != nil {
		log.Printf("Error fetching products by user ID: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription,
			pq.Array(&product.ProductImages), pq.Array(&product.CompressedProductImages), &product.ProductPrice, &product.CreatedAt)
		if err != nil {
			log.Printf("Error scanning product: %v", err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
