package handler

import (
	"database/sql"
	"homepage/internal/database"
	"homepage/internal/middleware"
	"homepage/internal/views"
	"net/http"
	"strconv"
)

func (h *Handler) GetContentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid content ID", http.StatusBadRequest)
		return
	}

	content, err := h.DB.GetContentById(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Content not found", http.StatusNotFound)
		return
	}

	user, _ := middleware.GetUserFromContext(r.Context())
	isAdmin := user != nil && user.IsAdmin

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.UnifiedContent(content, isAdmin, false).Render(r.Context(), w)
}

func (h *Handler) EditContentHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromContext(r.Context())
	if user == nil || !user.IsAdmin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid content ID", http.StatusBadRequest)
		return
	}

	content, err := h.DB.GetContentById(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Content not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.UnifiedContent(content, true, true).Render(r.Context(), w)
}

func (h *Handler) UpdateContentHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromContext(r.Context())
	if user == nil || !user.IsAdmin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(r.Form.Get("id"))
	content := database.UpdateContentParams{
		ID:       int32(id),
		Type:     r.Form.Get("type"),
		Title:    r.Form.Get("title"),
		Markdown: sql.NullString{String: r.Form.Get("content"), Valid: true},
	}

	updatedContent, err := h.DB.UpdateContent(r.Context(), content)
	if err != nil {
		http.Error(w, "Failed to update content", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.UnifiedContent(updatedContent, true, false).Render(r.Context(), w)
}
