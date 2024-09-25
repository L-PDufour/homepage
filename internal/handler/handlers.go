package handler

import (
	"fmt"
	"homepage/internal/auth"
	"homepage/internal/database"
	"homepage/internal/middleware"
	"homepage/internal/models"
	"homepage/internal/service"
	"homepage/internal/views"
	"net/http"
	"strconv"
	// "github.com/a-h/templ"
)

type Handler struct {
	DB      *database.Queries
	Service service.ContentService
}

func NewHandler(db *database.Queries, service service.ContentService) *Handler {
	return &Handler{
		DB:      db,
		Service: service,
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

// func (h *Handler) ViewContentHandler(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	idInt, err := strconv.Atoi(id)
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}
// 	content, err := h.DB.GetContentById(r.Context(), int32(idInt))
// 	if err != nil {
// 		http.Error(w, "Content not found", http.StatusNotFound)
// 		return
// 	}
// 	isEditing := r.URL.Query().Get("edit") == "true"
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	if r.Header.Get("HX-Request") == "true" {
// 		w.Header().Set("HX-Trigger", "contentLoaded")
// 	}
// 	props := models.ContentProps{
// 		Content:   content,
// 		IsAdmin:   utils.IsUserAdmin(r.Context()),
// 		IsEditing: isEditing,
// 	}
// 	views.ContentItem(props).Render(r.Context(), w)
// }
//
// func (h *Handler) GetContentHandler(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil {
// 		http.Error(w, "Invalid content ID", http.StatusBadRequest)
// 		return
// 	}
//
// 	content, err := h.DB.GetContentById(r.Context(), int32(id))
// 	if err != nil {
// 		http.Error(w, "Content not found", http.StatusNotFound)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//
// 	props := models.ContentProps{
// 		Content:   content,
// 		IsAdmin:   utils.IsUserAdmin(r.Context()),
// 		IsEditing: false,
// 	}
// 	views.ContentItem(props).Render(r.Context(), w)
// }
//
// func (h *Handler) EditContentHandler(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil {
// 		http.Error(w, "Invalid content ID", http.StatusBadRequest)
// 		return
// 	}
// 	content, err := h.DB.GetContentById(r.Context(), int32(id))
// 	if err != nil {
// 		http.Error(w, "Content not found", http.StatusNotFound)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	props := models.ContentProps{
// 		Content:   content,
// 		IsAdmin:   utils.IsUserAdmin(r.Context()),
// 		IsEditing: true,
// 	}
// 	views.ContentItem(props).Render(r.Context(), w)
// }
//
// func (h *Handler) UpdateContentHandler(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		http.Error(w, "Failed to parse form", http.StatusBadRequest)
// 		return
// 	}
//
// 	id, _ := strconv.Atoi(r.Form.Get("id"))
// 	content := database.UpdateContentParams{
// 		ID:          int32(id),
// 		ContentType: database.ContentType(r.Form.Get("type")),
// 		Title:       r.Form.Get("title"),
// 		Markdown:    sql.NullString{String: r.Form.Get("content"), Valid: true},
// 	}
//
// 	updatedContent, err := h.DB.UpdateContent(r.Context(), content)
// 	if err != nil {
// 		http.Error(w, "Failed to update content", http.StatusInternalServerError)
// 		return
// 	}
//
// 	props := models.ContentProps{
// 		Content:   updatedContent,
// 		IsAdmin:   utils.IsUserAdmin(r.Context()),
// 		IsEditing: false,
// 	}
// 	views.ContentItem(props).Render(r.Context(), w)
// }
//
// func (h *Handler) DeleteContentHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodDelete {
// 			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 			return
// 		}
// 		id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 		if err != nil {
// 			http.Error(w, "Invalid content ID", http.StatusBadRequest)
// 			return
// 		}
//
// 		err = h.DB.DeleteContent(r.Context(), int32(id))
// 		if err != nil {
// 			http.Error(w, "Content not found", http.StatusNotFound)
// 			return
// 		}
//
// 		if r.Header.Get("HX-Request") == "true" {
// 			w.Header().Set("HX-Trigger", "contentDeleted")
//
// 			return
// 		}
// 	}
// }

func (h *Handler) ListContent(contentTypeStr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if contentTypeStr == "" {
			contentTypeStr = r.URL.Query().Get("type")
		}

		props, err := h.Service.GetContentsByType(r.Context(), contentTypeStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		h.renderContentList(w, r, props)
	}
}

func (h *Handler) GetContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		if err != nil {
			http.Error(w, "", http.StatusNotFound)
			return
		}
		props, err := h.Service.GetContentById(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		h.renderContentList(w, r, props)
	}
}

func (h *Handler) renderContentList(w http.ResponseWriter, r *http.Request, props models.ContentProps) {
	views.Base(views.ContentList(props)).Render(r.Context(), w)
}

//
// func (h *Handler) NewContentFormHandler(w http.ResponseWriter, r *http.Request) {
// 	views.NewContentForm().Render(r.Context(), w)
// }
//
// func (h *Handler) CreateContentHandler(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, "Unable to parse form", http.StatusBadRequest)
// 		return
// 	}
//
// 	contentType := database.ContentType(r.Form.Get("type"))
// 	title := r.FormValue("title")
// 	markdown := r.FormValue("markdown")
// 	imageUrl := r.FormValue("image_url")
// 	link := r.FormValue("link")
//
// 	newContent := database.CreateContentParams{
// 		ContentType: contentType,
// 		Title:       title,
// 		Markdown:    sql.NullString{String: markdown, Valid: true},
// 		ImageUrl:    sql.NullString{String: imageUrl, Valid: imageUrl != ""},
// 		Link:        sql.NullString{String: link, Valid: link != ""},
// 	}
//
// 	_, err = h.DB.CreateContent(r.Context(), newContent)
// 	if err != nil {
// 		http.Error(w, "Unable to create content", http.StatusInternalServerError)
// 		return
// 	}
//
// 	w.Header().Set("HX-Trigger", "contentCreated")
// }
//
// func (h *Handler) GetFullContent() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodGet {
// 			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 			return
// 		}
// 		id := r.URL.Query().Get("id")
// 		idInt, err := strconv.Atoi(id)
// 		if err != nil {
// 			http.Error(w, "Invalid ID", http.StatusBadRequest)
// 			return
// 		}
//
// 		content, err := h.DB.GetContentById(r.Context(), int32(idInt))
// 		if err != nil {
// 			http.Error(w, "Error fetching content", http.StatusInternalServerError)
// 			return
// 		}
//
// 		props := models.ContentProps{
// 			Content:   content,
// 			IsAdmin:   utils.IsUserAdmin(r.Context()),
// 			IsEditing: false,
// 		}
//
// 		w.Header().Set("Content-Type", "text/html; charset=utf-8")
//
// 		if r.Header.Get("HX-Request") == "true" {
// 			// If it's an HTMX request, only render the FullContentView
// 			views.FullContentView(props).Render(r.Context(), w)
// 		} else {
// 			// For a full page load, wrap FullContentView in your Base template
// 			views.Base(views.FullContentView(props)).Render(r.Context(), w)
// 		}
// 	}
// }
//
// func (h *Handler) GetTruncatedContent(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	idInt, _ := strconv.Atoi(id)
// 	content, err := h.DB.GetContentById(r.Context(), int32(idInt))
// 	if err != nil {
// 		http.Error(w, "Error fetching content", http.StatusInternalServerError)
// 		return
// 	}
//
// 	props := models.ContentProps{
// 		Content: content,
// 		IsAdmin: utils.IsUserAdmin(r.Context()),
// 	}
//
// 	views.ContentItem(props).Render(r.Context(), w)
// }
