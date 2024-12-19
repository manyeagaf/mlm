package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// AppConfig stores app configuration variables
var AppConfig struct {
	Port string
}

// InitConfig loads environment variables
func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values.")
	}

	AppConfig.Port = os.Getenv("PORT")
	if AppConfig.Port == "" {
		AppConfig.Port = "8080" // Default port
	}
}
