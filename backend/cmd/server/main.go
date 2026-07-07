package main

import (
	"context"
	"log"

	"vanthu-backend/internal/config"
	"vanthu-backend/internal/db"
	"vanthu-backend/internal/router"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()
	pool, err := db.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("kết nối database thất bại: %v", err)
	}
	defer pool.Close()

	r := router.New(pool)

	log.Printf("server đang chạy tại :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server dừng với lỗi: %v", err)
	}
}
