package handler

import (
	"database/sql"
	"fmt"
	"homepage/internal/auth"
	"homepage/internal/database"
	"homepage/internal/middleware"
	"homepage/internal/models"
	"homepage/internal/service"
	"homepage/internal/utils"
	"homepage/internal/views"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func (h *Handler) DeleteContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if user is admin
		if !utils.IsUserAdmin(r.Context()) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid content ID", http.StatusBadRequest)
			return
		}

		err = h.Service.DeleteContent(r.Context(), id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, "Failed to delete content", http.StatusInternalServerError)
			}
			log.Printf("Error deleting content: %v", err)
			return
		}

		if r.Header.Get("HX-Request") == "true" {
			w.Header().Set("HX-Trigger", "contentDeleted")
			fmt.Fprintf(w, "Content deleted successfully")
		} else {
			http.Redirect(w, r, "/content", http.StatusSeeOther)
		}
	}
}

func (h *Handler) GetForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAdmin := utils.IsUserAdmin(r.Context())
		if isAdmin {
			h.renderNewForm(w, r)
		} else {
			http.Error(w, "Nope", http.StatusUnauthorized)
		}
	}
}

func (h *Handler) GetUpdateForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !utils.IsUserAdmin(r.Context()) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid content ID", http.StatusBadRequest)
			return
		}

		props, err := h.Service.GetContentById(r.Context(), id)
		if err != nil {
			http.Error(w, "Failed to fetch content", http.StatusInternalServerError)
			log.Printf("Error fetching content for update form: %v", err)
			return
		}

		if len(props.Content) == 0 {
			http.Error(w, "Content not found", http.StatusNotFound)
			return
		}

		h.renderUpdateForm(w, r, props.Content[0])
	}
}
func (h *Handler) CreateContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !utils.IsUserAdmin(r.Context()) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		params := database.CreateContentParams{
			ContentType: database.ContentType(r.Form.Get("type")),
			Title:       r.FormValue("title"),
			Markdown:    sql.NullString{String: r.FormValue("markdown"), Valid: true},
			ImageUrl:    sql.NullString{String: r.FormValue("image_url"), Valid: r.FormValue("image_url") != ""},
			Link:        sql.NullString{String: r.FormValue("link"), Valid: r.FormValue("link") != ""},
		}

		_, err = h.DB.CreateContent(r.Context(), params)
		if err != nil {
			http.Error(w, "Unable to create content", http.StatusInternalServerError)
			log.Printf("Error creating content: %v", err)
			return
		}

		http.Redirect(w, r, "/blog", http.StatusSeeOther)
	}
}

func (h *Handler) UpdateContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !utils.IsUserAdmin(r.Context()) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid content ID", http.StatusBadRequest)
			return
		}

		params := database.UpdateContentParams{
			ContentType: database.ContentType(r.Form.Get("type")),
			Title:       r.FormValue("title"),
			Markdown:    sql.NullString{String: r.FormValue("markdown"), Valid: true},
			ImageUrl:    sql.NullString{String: r.FormValue("image_url"), Valid: r.FormValue("image_url") != ""},
			Link:        sql.NullString{String: r.FormValue("link"), Valid: r.FormValue("link") != ""},
			ID:          int32(id),
		}

		content, err := h.DB.UpdateContent(r.Context(), params)
		if err != nil {
			http.Error(w, "Unable to update content", http.StatusInternalServerError)
			log.Printf("Error updating content: %v", err)
			return
		}

		if r.Header.Get("HX-Request") == "true" {
			w.Header().Set("HX-Trigger", "contentUpdated")
			h.renderContentList(w, r, models.ContentProps{Content: []database.Content{content}})
		} else {
			http.Redirect(w, r, "/blog", http.StatusSeeOther)
		}
	}
}
