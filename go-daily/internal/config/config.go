package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	// Instagram API credentials
	InstagramAccessToken string
	InstagramAccountID   string

	// Data paths
	LibrariesPath string
	PostedPath    string

	// Image generation settings
	ImageBasePath string

	// Environment
	Environment string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists (for local development)
	_ = godotenv.Load()

	cfg := &Config{
		InstagramAccessToken: os.Getenv("INSTAGRAM_ACCESS_TOKEN"),
		InstagramAccountID:   os.Getenv("INSTAGRAM_ACCOUNT_ID"),
		LibrariesPath:        getEnvOrDefault("LIBRARIES_PATH", "data/libraries.json"),
		PostedPath:           getEnvOrDefault("POSTED_PATH", "data/posted.json"),
		ImageBasePath:        getEnvOrDefault("IMAGE_BASE_PATH", "internal/image/assets/base.png"),
		Environment:          getEnvOrDefault("ENVIRONMENT", "development"),
	}

	// Validate required fields
	if cfg.InstagramAccessToken == "" {
		return nil, fmt.Errorf("INSTAGRAM_ACCESS_TOKEN is required")
	}
	if cfg.InstagramAccountID == "" {
		return nil, fmt.Errorf("INSTAGRAM_ACCOUNT_ID is required")
	}

	return cfg, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
