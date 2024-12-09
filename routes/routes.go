package routes

import (
	"product-management-system/controllers"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up the API routes
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id:[0-9]+}", controllers.GetProductByID).Methods("GET")
	r.HandleFunc("/products", controllers.GetProductsByUserID).Methods("GET")
}
