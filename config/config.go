package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GIN_MODE    string
	APIURL      string
	APIKey      string
	DBURL       string
	NEON_DB_URL string
}

func LoadConfig() Config {
	// Chỉ load file .env khi chạy local (khi biến môi trường RENDER không tồn tại)
	if os.Getenv("RENDER") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("Không tìm thấy file .env, sử dụng biến môi trường hệ thống")
		}
	}

	// fmt.Println("🚀 Đang tải cấu hình từ biến môi trường...", os.Getenv("NEON_DB_URL"))
	return Config{
		GIN_MODE:    os.Getenv("GIN_MODE"),
		APIURL:      os.Getenv("API_URL"),
		APIKey:      os.Getenv("API_KEY"),
		DBURL:       os.Getenv("DB_URL"),
		NEON_DB_URL: os.Getenv("NEON_DB_URL"),
	}
}
