package handler

import (
	"database/sql"
	"fmt"
	"homepage/internal/auth"
	"homepage/internal/database"
	"homepage/internal/middleware"
	"homepage/internal/models"
	"homepage/internal/views"
	"net/http"
	"strconv"
)

type Handler struct {
	DB *database.Queries
}

func NewHandler(db *database.Queries) *Handler {
	return &Handler{
		DB: db,
	}
}
func ParseContentType(s string) (database.ContentType, error) {
	switch s {
	case "blog", "project", "bio":
		return database.ContentType(s), nil
	default:
		return "", fmt.Errorf("invalid content type: %s", s)
	}
}

func (h *Handler) ServeResume() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		views.ResumeTemplate().Render(r.Context(), w)
	}
}

func (h *Handler) UnifiedView(contentType database.ContentType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := middleware.GetUserFromContext(r.Context())
		isAdmin := user != nil && user.IsAdmin

		contents, err := h.DB.GetContentsByType(r.Context(), contentType)
		if err != nil {
			http.Error(w, "Error fetching content", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		props := models.ContentViewProps{
			Contents:    contents,
			ContentType: contentType,
			IsAdmin:     isAdmin,
			IsEditing:   false,
		}
		component := views.UnifiedContentView(props)
		component.Render(r.Context(), w)
	}
}

func (h *Handler) AdminAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		views.AdminAuthPage().Render(r.Context(), w)
	}
}

func redirectToLogin(w http.ResponseWriter, r *http.Request) {
	loginURL := fmt.Sprintf("https://%s.cloudflareaccess.com/cdn-cgi/access/login?redirect_url=%s", auth.CfTeamDomain, "https://dev.lpdufour.xyz/admin/auth")
	w.Header().Set("Cache-Control", "no-store, must-revalidate")
	http.Redirect(w, r, loginURL, http.StatusTemporaryRedirect)
}

func (h *Handler) Admin(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromContext(r.Context())
	if user == nil {
		redirectToLogin(w, r)
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
	props := models.ContentItemProps{
		Content:   content,
		IsAdmin:   isAdmin,
		IsEditing: isEditing,
	}
	views.UnifiedContent(props).Render(r.Context(), w)
}

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

	props := models.ContentItemProps{
		Content:   content,
		IsAdmin:   isAdmin,
		IsEditing: false,
	}
	views.UnifiedContent(props).Render(r.Context(), w)
}

func (h *Handler) ListContentHandler(w http.ResponseWriter, r *http.Request) {
	contentTypeStr := r.URL.Query().Get("type")
	contentType, err := ParseContentType(contentTypeStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, _ := middleware.GetUserFromContext(r.Context())
	isAdmin := user != nil && user.IsAdmin

	contents, err := h.DB.GetContentsByType(r.Context(), contentType)
	if err != nil {
		http.Error(w, "Failed to fetch contents", http.StatusInternalServerError)
		return
	}

	props := models.ContentViewProps{
		Contents:    contents,
		ContentType: contentType,
		IsAdmin:     isAdmin,
		IsEditing:   false,
	}

	component := views.UnifiedContentView(props)
	component.Render(r.Context(), w)
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
	isAdmin := user.IsAdmin
	content, err := h.DB.GetContentById(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Content not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	props := models.ContentItemProps{
		Content:   content,
		IsAdmin:   isAdmin,
		IsEditing: true,
	}
	views.UnifiedContent(props).Render(r.Context(), w)
}

func (h *Handler) UpdateContentHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromContext(r.Context())
	if user == nil || !user.IsAdmin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	isAdmin := user.IsAdmin
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(r.Form.Get("id"))
	content := database.UpdateContentParams{
		ID:          int32(id),
		ContentType: database.ContentType(r.Form.Get("type")),
		Title:       r.Form.Get("title"),
		Markdown:    sql.NullString{String: r.Form.Get("content"), Valid: true},
	}

	updatedContent, err := h.DB.UpdateContent(r.Context(), content)
	if err != nil {
		http.Error(w, "Failed to update content", http.StatusInternalServerError)
		return
	}

	props := models.ContentItemProps{
		Content:   updatedContent,
		IsAdmin:   isAdmin,
		IsEditing: false,
	}
	views.UnifiedContent(props).Render(r.Context(), w)
}

func (h *Handler) DeleteContentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid content ID", http.StatusBadRequest)
		return
	}

	err = h.DB.DeleteContent(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Content not found", http.StatusNotFound)
		return
	}

	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Trigger", "contentDeleted")

		return
	}
}

func (h *Handler) NewContentFormHandler(w http.ResponseWriter, r *http.Request) {
	views.NewContentForm().Render(r.Context(), w)
}

func (h *Handler) CreateContentHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	contentType := database.ContentType(r.Form.Get("type"))
	title := r.FormValue("title")
	markdown := r.FormValue("markdown")
	imageUrl := r.FormValue("image_url")
	link := r.FormValue("link")

	newContent := database.CreateContentParams{
		ContentType: contentType,
		Title:       title,
		Markdown:    sql.NullString{String: markdown, Valid: true},
		ImageUrl:    sql.NullString{String: imageUrl, Valid: imageUrl != ""},
		Link:        sql.NullString{String: link, Valid: link != ""},
	}

	_, err = h.DB.CreateContent(r.Context(), newContent)
	if err != nil {
		http.Error(w, "Unable to create content", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "contentCreated")
}

func (h *Handler) GetFullContent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, _ := strconv.Atoi(id)
	content, err := h.DB.GetContentById(r.Context(), int32(idInt))
	if err != nil {
		http.Error(w, "Error fetching content", http.StatusInternalServerError)
		return
	}

	views.FullContentSection(content).Render(r.Context(), w)
}

func (h *Handler) GetTruncatedContent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, _ := strconv.Atoi(id)
	content, err := h.DB.GetContentById(r.Context(), int32(idInt))
	if err != nil {
		http.Error(w, "Error fetching content", http.StatusInternalServerError)
		return
	}

	views.ContentSection(content).Render(r.Context(), w)
}
