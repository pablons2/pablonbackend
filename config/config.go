package config

import (
	"os"
)

// Configuration holds all config settings
type Configuration struct {
	JWTSecret string
	Port      string
}

// GetConfig returns the application configuration
func GetConfig() Configuration {
	jwtSecret := getEnvOrDefault("JWT_SECRET", "sua-chave-secreta")
	port := getEnvOrDefault("PORT", "3000")

	return Configuration{
		JWTSecret: jwtSecret,
		Port:      port,
	}
}

// getEnvOrDefault gets an environment variable or returns a default value
func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
