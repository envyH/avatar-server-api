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
	// Chỉ load file .env khi chạy local (khi biến môi trường RENDER không tồn tại)
	if os.Getenv("RENDER") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("Không tìm thấy file .env, sử dụng biến môi trường hệ thống")
		}
	}

	return Config{
		APIURL: os.Getenv("API_URL"),
		APIKey: os.Getenv("API_KEY"),
		DBURL:  os.Getenv("DB_URL"),
	}
}
