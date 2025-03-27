package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL, err := service.Shorten(req.URL)
	if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

	resp := ShortenResponse{ShortURL: shortURL}

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(resp)
	log.Println("Сокращенная ссылка: ", resp.ShortURL)
}