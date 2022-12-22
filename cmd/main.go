package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/noah415/jwt-go-server/internal/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading the .env file")
	} else {
		fmt.Println(".env file successfully loaded")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.InitRouter()
}
