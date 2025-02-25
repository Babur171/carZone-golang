package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds all configuration settings.
type Config struct {
	Port    string
	BaseURL string
	DBURL   string
}

// Global variable to store the loaded configuration.
var AppConfig *Config

// LoadConfig loads environment variables into the Config struct.
func LoadConfig() {
	// Load .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	// Assign values to AppConfig
	AppConfig = &Config{
		Port:    getEnv("PORT", "8080"),
		BaseURL: getEnv("BASE_URL", "http://localhost:8080"),
		DBURL:   getEnv("DATABASE_URL", "postgres://user:pass@localhost:5432/db"),
	}
}

// getEnv fetches an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
