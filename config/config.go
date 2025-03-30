package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIURL string
	APIKey string
	DBURL  string
}

func LoadConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Không thể tải file .env: %v", err)
	}

	return Config{
		APIURL: os.Getenv("API_URL"),
		APIKey: os.Getenv("API_KEY"),
		DBURL:  os.Getenv("DB_URL"),
	}
}
