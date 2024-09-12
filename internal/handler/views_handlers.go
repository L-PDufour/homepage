package handler

import (
	"database/sql"
	"fmt"
	"homepage/internal/auth"
	"homepage/internal/database"
	"homepage/internal/middleware"
	"homepage/internal/views"
	"net/http"
	"strconv"
)

func (h *Handler) AdminAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.AdminAuthPage().Render(r.Context(), w)
}

func RedirectToLogin(w http.ResponseWriter, r *http.Request) {
	loginURL := fmt.Sprintf("https://%s.cloudflareaccess.com/cdn-cgi/access/login?redirect_url=%s", auth.CfTeamDomain, "https://dev.lpdufour.xyz/admin/auth")
	w.Header().Set("Cache-Control", "no-store, must-revalidate")
	http.Redirect(w, r, loginURL, http.StatusTemporaryRedirect)
}

func (h *Handler) Admin(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromContext(r.Context())
	if user == nil {
		RedirectToLogin(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "private, max-age=0, no-cache")
	views.Adminpage().Render(r.Context(), w)
}

func (h *Handler) ViewContentHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromContext(r.Context())
	isAdmin := user != nil && user.IsAdmin
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	content, err := h.DB.GetContentById(r.Context(), int32(idInt))
	if err != nil {
		http.Error(w, "Content not found", http.StatusNotFound)
		return
	}

	isEditing := r.URL.Query().Get("edit") == "true"

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Trigger", "contentLoaded")
	}
	views.UnifiedContent(content, isAdmin, isEditing).Render(r.Context(), w)
}

func (h *Handler) BlogPage(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromContext(r.Context())
	isAdmin := user != nil && user.IsAdmin
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.Blog(isAdmin).Render(r.Context(), w)
}

func (h *Handler) BioHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromContext(r.Context())
	isAdmin := user != nil && user.IsAdmin

	content, err := h.DB.GetContentByTitle(r.Context(), database.GetContentByTitleParams{
		Type:  "about",
		Title: "bio",
	})
	if err != nil {
		content = database.Content{
			ID:    0,
			Type:  "about",
			Title: "bio",
			Markdown: sql.NullString{
				String: "No content available",
				Valid:  true,
			},
		}
	}

	isEditing := r.URL.Query().Get("edit") == "true"

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.Bio(content, isAdmin, isEditing).Render(r.Context(), w)
}

func (h *Handler) ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromContext(r.Context())
	isAdmin := user != nil && user.IsAdmin

	projects, err := h.DB.GetContentsByType(r.Context(), "project")
	if err != nil {
		http.Error(w, "Error fetching projects", http.StatusInternalServerError)
		return
	}

	editingID, _ := strconv.Atoi(r.URL.Query().Get("edit"))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.Projects(projects, isAdmin, editingID).Render(r.Context(), w)
}
