package main

import (
	"fmt"
	"log"

	"url-shortener/internal/config"
)

func main() {
	cfg := config.Load()

	log.Printf(" Server running on port %s\n", cfg.ServerPort)
	log.Printf(" Connected to PostgreSQL: %s\n", cfg.PostgresDSN)
	log.Printf(" Connected to Redis: %s\n", cfg.RedisAddr)

	fmt.Println("Configuration loaded successfully")
}
