package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	config     map[string]string
	configOnce sync.Once
)

func LoadConfig() map[string]string {
	configOnce.Do(func() {
		// Load .env file
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, using system environment variables.")
		}

		config = map[string]string{
			"MONGO_URI": getEnv("MONGO_URI", "mongodb://localhost:270171"),
			"MONGO_DB":  getEnv("MONGO_DB", "testdbgo"),
			"APP_PORT":  getEnv("APP_PORT", "8080"),
		}
	})
	return config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetConfig(key string) string {
	return LoadConfig()[key]
}
