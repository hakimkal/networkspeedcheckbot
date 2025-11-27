package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl            string
	Port             string
	TelegramApiToken string
}

// LoadConfig loads environment variables from .env file
func LoadConfig() (*Config, error) {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := &Config{
		DBUrl:            os.Getenv("DB_URL"),
		Port:             os.Getenv("PORT"),
		TelegramApiToken: os.Getenv("TELEGRAM_API_TOKEN"),
	}

	// Set default port if not set
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	// Validation
	if cfg.DBUrl == "" {
		return nil, fmt.Errorf("DB_URL is required")
	}

	if strings.TrimSpace(cfg.TelegramApiToken) == "" {
		return nil, fmt.Errorf("TELEGRAM_API_TOKEN is required")
	}

	return cfg, nil

}
