package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Load enviroment from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in enviroment")
	}
	// Serve static files (React build)
	fs := http.FileServer(http.Dir("./frontend/build"))
	http.Handle("/", fs)
	log.Printf("Server is running on port %v", port)
	server := &http.Server{
		Handler: nil,
		Addr:    ":" + port,
	}
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
