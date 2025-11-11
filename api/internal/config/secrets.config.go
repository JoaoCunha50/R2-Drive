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
	ADMIN_PASSWORD string
	ADMIN_USERNAME string
	ADMIN_EMAIL string
	JWT_SECRET string
}

var EnvVariables *env

func init() {
	EnvVariables = LoadConfig()
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

	pass := os.Getenv("ADMIN_PASSWORD") 
	if pass == "" {
		log.Fatal("Error: Env variable missing")
	}

	user := os.Getenv("ADMIN_USERNAME")
	if user == "" {
		log.Fatal("Error: Env variable missing")
	}

	email := os.Getenv("ADMIN_EMAIL")
	if email == "" {
		log.Fatal("Error: Env variable missing")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("Error: Env variable missing")
	}

	return &env{PORT: port, DATABASE_URL: url, ADMIN_PASSWORD: pass, ADMIN_USERNAME: user, ADMIN_EMAIL: email, JWT_SECRET: jwtSecret}
}
