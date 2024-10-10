package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	return godotenv.Load(".env.local")
}

func GetEnv(key string, fallback string) string {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}
	return value
}
