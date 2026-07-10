package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

// Load đọc cấu hình từ .env (nếu có) và biến môi trường.
// Connection string KHÔNG được hardcode — luôn lấy từ DATABASE_URL.
func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("không tìm thấy file .env, dùng biến môi trường hệ thống")
	}

	cfg := Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        "10000",
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("thiếu biến môi trường DATABASE_URL")
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg
}
