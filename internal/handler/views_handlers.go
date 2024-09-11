package handler

import (
	"fmt"
	"homepage/internal/auth"
	"homepage/internal/database"
	"homepage/internal/middleware"
	"homepage/internal/views"
	"net/http"
)

func (h *Handler) AdminAuth(w http.ResponseWriter, r *http.Request) {
	views.AdminAuthPage().Render(r.Context(), w)
}

func RedirectToLogin(w http.ResponseWriter, r *http.Request) {
	loginURL := fmt.Sprintf("https://%s.cloudflareaccess.com/cdn-cgi/access/login?redirect_url=%s", auth.CfTeamDomain, "https://dev.lpdufour.xyz/admin/auth")
	http.Redirect(w, r, loginURL, http.StatusTemporaryRedirect)
}

func (h *Handler) Admin(w http.ResponseWriter, r *http.Request) {
	// Fetch the user from the context
	user, _ := middleware.GetUserFromContext(r.Context())

	// If user is not authenticated, redirect them to the login page
	if user == nil {
		RedirectToLogin(w, r)
		return
	}

	// If the user is authenticated, render the admin page
	views.Adminpage().Render(r.Context(), w)
}

func (h *Handler) BlogPage(w http.ResponseWriter, r *http.Request) {
	views.Blog().Render(r.Context(), w)
}

func (h *Handler) BioHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromContext(r.Context())

	// Fetch content from the database
	contentDB, err := h.DB.GetContentByTitle(r.Context(), database.GetContentByTitleParams{
		Type:  "about",
		Title: "bio",
	})
	if err != nil || !contentDB.Markdown.Valid {
		// Render the Bio template with empty content if not found or an error occurred
		views.Bio("", false).Render(r.Context(), w)
		return
	}

	// Sanitize the Markdown content
	sanitizedHTML, err := h.MD.ConvertAndSanitize(contentDB.Markdown.String)
	if err != nil {
		sanitizedHTML = "Error processing content"
	}

	// Determine if the user is an admin
	isAdmin := user != nil && user.IsAdmin

	// Render the Bio component with sanitized content
	views.Bio(sanitizedHTML, isAdmin).Render(r.Context(), w)
}

func (h *Handler) ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromContext(r.Context())

	contentDB, err := h.DB.GetContentsByType(r.Context(), "project")
	if err != nil {
		// Handle error appropriately, possibly render an error view
		http.Error(w, "Error fetching projects", http.StatusInternalServerError)
		return
	}

	// Sanitize Markdown for each project
	for i := range contentDB {
		if contentDB[i].Markdown.Valid {
			sanitizedHTML, err := h.MD.ConvertAndSanitize(contentDB[i].Markdown.String)
			if err != nil {
				sanitizedHTML = "Error processing content"
			}
			contentDB[i].Markdown.String = sanitizedHTML // Update the field
		} else {
			contentDB[i].Markdown.String = "No content available" // Handle invalid content
		}
	}

	isAdmin := user != nil && user.IsAdmin

	views.Projects(contentDB, isAdmin).Render(r.Context(), w)
}
