package models

import (
	"crypto/rand"
	"encoding/base64"
)

type URL struct {
	ShortURL string
	LongURL  string
}

// GenerateShortURL generates a random short URL of specified length
func GenerateShortURL(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	
	// Use URL-safe base64 encoding and take only the first 'length' characters
	encoded := base64.URLEncoding.EncodeToString(b)
	return encoded[:length], nil
}
