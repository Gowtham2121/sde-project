package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"product-management-system/models"
	"product-management-system/repository"
	"product-management-system/utils"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateProduct handles POST /products to create a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	log.Printf("Received Product Data: %+v", product)

	if err := repository.CreateProduct(&product); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, product)
}

// GetProductByID handles GET /products/:id to retrieve a product by ID
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Product ID is required")
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := repository.GetProductByID(productID)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, product)
}

// GetProductsByUserID handles GET /products to retrieve products for a specific user
func GetProductsByUserID(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	products, err := repository.GetProductsByUserID(userIDInt)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve products")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, products)
}
