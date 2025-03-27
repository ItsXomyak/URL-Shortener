package storage

func SaveURL(shortURL, longURL string) error {
	_, err := db.Exec("INSERT INTO urls (short_url, original_url) VALUES ($1, $2)", shortURL, longURL)
	return err
}
