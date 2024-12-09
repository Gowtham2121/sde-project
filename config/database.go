package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// ConnectDatabase connects to the PostgreSQL database using environment variables
func ConnectDatabase() {
	// Build the connection string using environment variables
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),     // Correct variable name for host
		os.Getenv("DB_PORT"),     // Correct variable name for port
		os.Getenv("DB_USER"),     // Correct variable name for user
		os.Getenv("DB_PASSWORD"), // Correct variable name for password
		os.Getenv("DB_NAME"),     // Correct variable name for database name
	)

	// Debug: print the connection string (for troubleshooting purposes)
	log.Println("Connection string:", dsn)

	// Open the connection to the database
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Verify the connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	log.Println("Database connected successfully!")
}
