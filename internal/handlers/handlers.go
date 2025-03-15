package handlers

import (
	"context"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"url-shortener/internal/models"
)

type URLStore interface {
	SaveURL(ctx context.Context, shortURL, longURL string) error
	GetURL(ctx context.Context, shortURL string) (string, error)
}

type Handler struct {
	store     URLStore
	templates *template.Template
}

func NewHandler(store URLStore, templateDir string) (*Handler, error) {
	templates, err := template.ParseGlob(filepath.Join(templateDir, "*.html"))
	if err != nil {
		return nil, err
	}

	return &Handler{
		store:     store,
		templates: templates,
	}, nil
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	h.templates.ExecuteTemplate(w, "index.html", nil)
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortURL, err := models.GenerateShortURL(6)
	if err != nil {
		http.Error(w, "Error generating short URL", http.StatusInternalServerError)
		return
	}

	err = h.store.SaveURL(r.Context(), shortURL, longURL)
	if err != nil {
		http.Error(w, "Error saving URL", http.StatusInternalServerError)
		return
	}

	data := struct {
		ShortURL string
		LongURL  string
	}{
		ShortURL: shortURL,
		LongURL:  longURL,
	}

	h.templates.ExecuteTemplate(w, "result.html", data)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	// Extraer el código de la URL corta del path
	shortURL := strings.TrimPrefix(r.URL.Path, "/")
	
	// Si es una ruta especial, no procesar
	if shortURL == "" || shortURL == "shorten" || strings.HasPrefix(shortURL, "static/") {
		http.NotFound(w, r)
		return
	}

	longURL, err := h.store.GetURL(r.Context(), shortURL)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
