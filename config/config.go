package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DbURL string
var AppPort string

func LoadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load variables from .env file
	DbURL = os.Getenv("DB_URL")
	AppPort = os.Getenv("APP_PORT")

	if DbURL == "" {
		log.Fatal("DB_URL is not set in .env file")
	}
	if AppPort == "" {
		AppPort = "3000" // Default port
	}
}
