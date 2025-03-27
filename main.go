package main

import (
	"log"

	"url-shortener/internal/storage"
)

func main() {
	err := storage.InitDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	err = storage.MigrateDB()
	if err != nil {
		log.Fatalf("Ошибка миграции БД: %v", err)
	}

	log.Println("Сервис запущен")
}
