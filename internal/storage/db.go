package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(" Не удалось загрузить .env файл, используем системные переменные окружения")
	}
}

func InitDB() error {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
		return fmt.Errorf(" Ошибка: отсутствуют данные для подключения к БД")
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf(" Ошибка при открытии БД: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf(" Ошибка при подключении к БД: %v", err)
	}

	log.Println(" подключение к PostgreSQL!")
	return nil
}

func MigrateDB() error {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		short_url TEXT UNIQUE NOT NULL,
		long_url TEXT NOT NULL,
		clicks INT DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("Ошибка при создании таблицы: %v", err)
	}

	log.Println("Таблица 'urls' успешно проверена/создана")
	return nil
}

func IncrementClickCount(shortURL string) error {
	// Выполняем запрос для увеличения счетчика кликов
	_, err := db.Exec("UPDATE urls SET clicks = clicks + 1 WHERE short_url = $1", shortURL)
	if err != nil {
		log.Println("Ошибка при увеличении счетчика кликов:", err)
		return err
	}
	return nil
}


func GetOriginalURL(shortURL string) (string, error) {
	var originalURL string
	err := db.QueryRow("SELECT long_url FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("URL не найден")
		}
		return "", fmt.Errorf("ошибка поиска URL: %w", err)
	}

	_, _ = db.Exec("UPDATE urls SET clicks = clicks + 1 WHERE short_url = $1", shortURL)

	return originalURL, nil
}

func SaveURL(shortURL, longURL string) error {
	_, err := db.Exec(
		"INSERT INTO urls (short_url, long_url) VALUES ($1, $2) ON CONFLICT (short_url) DO NOTHING",
		shortURL, longURL,
	)
	return err
}


func GetStats(shortURL string) (int, error) {
	var clicks int
	err := db.QueryRow("SELECT clicks FROM urls WHERE short_url = $1", shortURL).Scan(&clicks)
	if err != nil {
		log.Println("Ошибка при получении статистики:", err)
		return 0, err
	}
	return clicks, nil
}

func generateShortURL(originalURL string) string {
	return fmt.Sprintf("%x", len(originalURL))
}