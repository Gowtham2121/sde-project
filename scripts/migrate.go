package main

import (
	"log"
	"product-management-system/config"
)

func main() {
	config.LoadConfig()
	config.ConnectDatabase()

	db := config.DB

	// Create users table
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}
	log.Println("Users table created successfully!")

	// Create products table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id) ON DELETE CASCADE,
		product_name VARCHAR(255) NOT NULL,
		product_description TEXT,
		product_images TEXT[],
		compressed_product_images TEXT[],
		product_price DECIMAL(10, 2) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		log.Fatalf("Error creating products table: %v", err)
	}
	log.Println("Products table created successfully!")
}
