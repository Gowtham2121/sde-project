package main

import (
	"log"
	"net/http"
	"product-management-system/config"
	"product-management-system/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDatabase()

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
