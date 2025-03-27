package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"url-shortener/internal/service"
	"url-shortener/internal/storage"
)

type RequestBody struct {
	URL string `json:"url"`
}

type ResponseBody struct {
	ShortURL string `json:"short_url,omitempty"`
	Error    string `json:"error,omitempty"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "URL не может быть пустым", http.StatusBadRequest)
		return
	}

	shortURL, err := service.Shorten(req.URL)
	if err != nil {
		http.Error(w, "Ошибка сохранения URL", http.StatusInternalServerError)
		return
	}

	resp := ResponseBody{ShortURL: shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}


func RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["short_url"]

	originalURL, err := storage.GetOriginalURL(shortURL)
	if err != nil {
		http.Error(w, "URL не найден", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func StatsURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["short_url"]

	clicks, err := storage.GetStats(shortURL)
	if err != nil {
		http.Error(w, "URL не найден", http.StatusNotFound)
		return
	}

	resp := map[string]int{"clicks": clicks}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
