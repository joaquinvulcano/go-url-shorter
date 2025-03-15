package main

import (
	"log"
	"net/http"
	"path/filepath"
	"url-shortener/internal/handlers"
	"url-shortener/internal/storage"
)

func main() {
	// Initialize JSON storage
	dataDir := filepath.Join(".", "data")
	store, err := storage.NewJSONStore(dataDir)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	// Initialize handlers
	handler, err := handlers.NewHandler(store, "templates")
	if err != nil {
		log.Fatalf("Failed to initialize handlers: %v", err)
	}

	// Create a new router
	mux := http.NewServeMux()
	
	// Set up routes
	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/shorten", handler.Shorten)
	
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// Handle redirects for short URLs (cualquier ruta que no sea / o /shorten)
	mux.HandleFunc("/{shortURL}", handler.Redirect)

	// Start server
	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
