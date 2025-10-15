package configs

import (
	"os"
	"strings"
	"time"
)

// Config holds all configuration for the application
type Config struct {
	// Server
	Port    string
	GinMode string

	// MongoDB
	MongoURI string

	// JWT
	JWTSecret     string
	JWTExpiration time.Duration

	// Gemini AI
	GeminiAPIKey string

	// Payments
	PaymentsProvider string

	// CORS
	AllowedOrigins []string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	jwtExpStr := os.Getenv("JWT_EXPIRATION")
	if jwtExpStr == "" {
		jwtExpStr = "24h"
	}
	jwtExp, _ := time.ParseDuration(jwtExpStr)

	originsStr := os.Getenv("ALLOWED_ORIGINS")
	if originsStr == "" {
		originsStr = "http://localhost:5173,http://localhost:3000"
	}
	origins := strings.Split(originsStr, ",")

	return &Config{
		Port:             getEnv("PORT", "3000"),
		GinMode:          getEnv("GIN_MODE", "debug"),
		MongoURI:         getEnv("MONGODB_URI", "mongodb://localhost:27017/fastspot"),
		JWTSecret:        getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		JWTExpiration:    jwtExp,
		GeminiAPIKey:     getEnv("GEMINI_API_KEY", ""),
		PaymentsProvider: getEnv("PAYMENTS_PROVIDER", "stub"),
		AllowedOrigins:   origins,
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
