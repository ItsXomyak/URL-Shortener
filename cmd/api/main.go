package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"url-shortener/internal/handler"
	"url-shortener/internal/storage"
)

func main() { err := storage.InitDB()
	if err != nil {
    log.Fatal(err)
  }

	err = storage.MigrateDB()
	if err != nil {
    log.Fatal(err)
  }

	r := mux.NewRouter()
 
	r.HandleFunc("/shorten", handler.ShortenURL).Methods("POST") // сокращает урл
	r.HandleFunc("/{short_url}", handler.RedirectURL).Methods("GET") // делает редирект
	r.HandleFunc("/stats/{short_url}", handler.StatsURL).Methods("GET") // стата по урлу

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}