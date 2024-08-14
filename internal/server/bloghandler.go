package server

import (
	"bytes"
	"context"
	"homepage/cmd/web"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current working directory: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Construct the absolute path
	mdPath := filepath.Join(cwd, "cmd", "web", "content.md")
	log.Printf("Current working directory: %s", cwd)
	log.Printf("Attempting to read file at: %s", mdPath)
	// Read the Markdown file
	mdContent, err := os.ReadFile(mdPath)
	if err != nil {
		log.Printf("Error reading Markdown file at %s: %v", mdPath, err)
		http.Error(w, "Error reading Markdown file", http.StatusInternalServerError)
		return
	}

	// Convert Markdown to HTML
	var buf bytes.Buffer
	if err := goldmark.Convert(mdContent, &buf); err != nil {
		log.Printf("Error converting Markdown to HTML: %v", err)
		http.Error(w, "Error converting Markdown to HTML", http.StatusInternalServerError)
		return
	}

	// Set the converted HTML as a context value
	ctx := context.WithValue(r.Context(), "blogContent", buf.String())

	// Render the Blog component
	err = web.Blog().Render(ctx, w)
	if err != nil {
		log.Printf("Error rendering Blog component: %v", err)
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
		return
	}
}
