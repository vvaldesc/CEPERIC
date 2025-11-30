package config

import (
	"os"
)

type Config struct {
	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// Firebase
	FirebaseProjectID   string
	FirebaseCredentials string

	// App
	Environment    string
	AllowedOrigins string
	Port           string
}

func New() *Config {
	return &Config{
		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "ceperic_user"),
		DBPassword: getEnv("DB_PASSWORD", "ceperic_pass"),
		DBName:     getEnv("DB_NAME", "ceperic_db"),

		// Firebase
		FirebaseProjectID:   getEnv("FIREBASE_PROJECT_ID", "ceperic-68bcd"),
		FirebaseCredentials: getEnv("FIREBASE_CREDENTIALS", ""),

		// App
		Environment:    getEnv("ENVIRONMENT", "development"),
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "http://localhost:4200,https://ceperic.web.app"),
		Port:           getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
