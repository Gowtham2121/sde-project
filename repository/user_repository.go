package repository

import (
	"log"
	"product-management-system/config"
	"product-management-system/models"
)

// CreateUser inserts a new user into the database
func CreateUser(user *models.User) error {
	query := `
        INSERT INTO users (name, email, created_at)
        VALUES ($1, $2, CURRENT_TIMESTAMP)
        RETURNING id;
    `
	err := config.DB.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}

// GetUserByID retrieves a user by their ID
func GetUserByID(userID int) (*models.User, error) {
	user := &models.User{}
	query := `
        SELECT id, name, email, created_at
        FROM users
        WHERE id = $1;
    `
	row := config.DB.QueryRow(query, userID)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		log.Printf("Error fetching user by ID: %v", err)
		return nil, err
	}
	return user, nil
}
