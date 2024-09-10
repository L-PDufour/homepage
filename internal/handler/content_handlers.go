package handler

import (
	"database/sql"
	"homepage/internal/database"
	"homepage/internal/views"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetContentHandler(w http.ResponseWriter, r *http.Request) {
	typeStr := r.URL.Query().Get("type")
	titleStr := r.URL.Query().Get("title")

	contentDB, err := h.DB.GetContentByTitle(r.Context(), database.GetContentByTitleParams{
		Type:  typeStr,
		Title: titleStr,
	})
	if err != nil {
		http.Error(w, "Content not found", http.StatusNotFound)
		return
	}
	if !contentDB.Markdown.Valid {
		http.Error(w, "Content is empty", http.StatusNoContent)
		return
	}

	sanitizedHTML, err := h.MD.ConvertAndSanitize(contentDB.Markdown.String)
	if err != nil {
		http.Error(w, "Error processing content", http.StatusInternalServerError)
		return
	}

	// Render only the content section
	views.ContentSection(sanitizedHTML).Render(r.Context(), w)
}

func (h *Handler) EditContentHandler(w http.ResponseWriter, r *http.Request) {
	typeStr := r.URL.Query().Get("type")
	titleStr := r.URL.Query().Get("title")

	contentDB, err := h.DB.GetContentByTitle(r.Context(), database.GetContentByTitleParams{
		Type:  typeStr,
		Title: titleStr,
	})
	if err != nil {
		http.Error(w, "Content not found", http.StatusNotFound)
		return
	}

	views.EditContentSection(typeStr, titleStr, contentDB.Markdown.String, contentDB.ID).Render(r.Context(), w)
}

func (h *Handler) UpdateContentHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	typeStr := r.Form.Get("type")
	titleStr := r.Form.Get("title")
	newContent := r.Form.Get("content")
	idStr := r.Form.Get("id")
	log.Println("update")
	log.Println(idStr)
	id, _ := strconv.ParseInt(idStr, 10, 32)
	// Convert newContent to sql.NullString
	nullableContent := sql.NullString{
		String: newContent,
		Valid:  newContent != "",
	}

	// Update content in the database
	err := h.DB.UpdateContent(r.Context(), database.UpdateContentParams{
		Type:     typeStr,
		Title:    titleStr,
		Markdown: nullableContent,
		ID:       int32(id),
	})
	if err != nil {
		http.Error(w, "Error updating content", http.StatusInternalServerError)
		return
	}

	// Return the updated content to render
	views.ContentSection(nullableContent.String).Render(r.Context(), w)
}
