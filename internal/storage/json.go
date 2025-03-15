package storage

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

type URLEntry struct {
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}

type JSONStore struct {
	filepath string
	urls    map[string]string
	mutex   sync.RWMutex
}

func NewJSONStore(dataDir string) (*JSONStore, error) {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	filepath := filepath.Join(dataDir, "urls.json")
	store := &JSONStore{
		filepath: filepath,
		urls:    make(map[string]string),
	}

	// Cargar datos existentes si el archivo existe
	if _, err := os.Stat(filepath); !os.IsNotExist(err) {
		data, err := os.ReadFile(filepath)
		if err != nil {
			return nil, err
		}

		var entries []URLEntry
		if err := json.Unmarshal(data, &entries); err != nil {
			return nil, err
		}

		for _, entry := range entries {
			store.urls[entry.ShortURL] = entry.LongURL
		}
	}

	return store, nil
}

func (s *JSONStore) SaveURL(ctx context.Context, shortURL, longURL string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.urls[shortURL] = longURL

	// Convertir el mapa a slice para guardar
	var entries []URLEntry
	for short, long := range s.urls {
		entries = append(entries, URLEntry{
			ShortURL: short,
			LongURL:  long,
		})
	}

	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filepath, data, 0644)
}

func (s *JSONStore) GetURL(ctx context.Context, shortURL string) (string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if longURL, exists := s.urls[shortURL]; exists {
		return longURL, nil
	}
	return "", ErrURLNotFound
}

func (s *JSONStore) Close() error {
	return nil
}
