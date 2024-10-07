package configuration

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT          string
	HOST          string
	DB_HOST       string
	DATABASE_NAME string
	DB_USER       string
	DB_PASSWORD   string
	DB_PORT       string
	SSL_MODE      string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	return &Config{
		PORT:          getEnv("PORT", "8000"),
		HOST:          getEnv("HOST", "localhost"),
		DB_HOST:       getEnv("DB_HOST", "localhost"),
		DB_USER:       getEnv("DB_USER", "user"),
		DATABASE_NAME: getEnv("DATABASE_NAME", "messaging_api"),
		DB_PASSWORD:   getEnv("DB_PASSWORD", ""),
		DB_PORT:       getEnv("DB_PORT", "5432"),
		SSL_MODE:      getEnv("SSL_MODE", "localhost"),
	}, nil
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
