package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/masoudsadeghiDev/url-shortener/internal/services/api/routes"
)

func main() {

	// Load enviroment from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in environment")
	}

	router := routes.NewRoutes()
	log.Printf("Server is running on port %v", port)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
