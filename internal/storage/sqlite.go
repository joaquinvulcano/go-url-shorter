package storage

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(dbPath string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Crear la tabla si no existe
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			short_url TEXT PRIMARY KEY,
			long_url TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return nil, err
	}

	return &SQLiteStore{db: db}, nil
}

func (s *SQLiteStore) SaveURL(ctx context.Context, shortURL, longURL string) error {
	_, err := s.db.ExecContext(ctx, 
		"INSERT INTO urls (short_url, long_url) VALUES (?, ?)",
		shortURL, longURL)
	return err
}

func (s *SQLiteStore) GetURL(ctx context.Context, shortURL string) (string, error) {
	var longURL string
	err := s.db.QueryRowContext(ctx, 
		"SELECT long_url FROM urls WHERE short_url = ?",
		shortURL).Scan(&longURL)
	if err == sql.ErrNoRows {
		return "", ErrURLNotFound
	}
	return longURL, err
}

func (s *SQLiteStore) Close() error {
	return s.db.Close()
}
