package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	requiredEnvVars := []string{
		"SERVER_PORT",
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
		"DB_USER",
		"JWT_SECRET",
		"JWT_REFRESH_SECRET",
		"JWT_EXPIRED_AT",
		"JWT_REFRESH_EXPIRED_AT",
		"REDIS_HOST",
	}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("Required environment variable %s is not set", envVar)
			panic("Required environment variable is not set")
		}
	}

	log.Println("Configuration Loaded successfully")
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
