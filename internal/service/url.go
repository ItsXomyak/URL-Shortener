package service

import (
	"crypto/sha256"
	"encoding/base64"

	"url-shortener/internal/storage"
)

func Shorten(longURL string) (string, error) {
	hash := sha256.Sum256([]byte(longURL))
	shortURL := base64.URLEncoding.EncodeToString(hash[:])[:8]

	err := storage.SaveURL(shortURL, longURL) 
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

