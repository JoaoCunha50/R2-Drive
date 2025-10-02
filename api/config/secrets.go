package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	PORT string
	DATABASE_URL string
}

func LoadConfig() *env {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Warning: .env file not found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("Error: DATABASE_URL not found")
	}

	return &env{PORT: port, DATABASE_URL: url}
}
