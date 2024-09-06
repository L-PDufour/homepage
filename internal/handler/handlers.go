package handler

import (
	"context"
	"fmt"
	"homepage/internal/auth"
	"homepage/internal/blog"
	"homepage/internal/database"
	"homepage/internal/markdown"
	"homepage/internal/views"
	"net/http"
	"strconv"
)

type Handler struct {
	DB          *database.Queries
	MD          markdown.MarkdownConverter
	BlogService *blog.BlogService
}

func NewHandler(db *database.Queries, md markdown.MarkdownConverter, blogService *blog.BlogService) *Handler {
	return &Handler{
		DB:          db,
		MD:          md,
		BlogService: blogService,
	}
}

func (h *Handler) BlogPage(w http.ResponseWriter, r *http.Request) {
	component := views.Blog()
	component.Render(r.Context(), w)
}

func (h *Handler) GetBlogPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.BlogService.ListBlogPosts(r.Context())
	if err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	// Check if the request is for the sidebar (recent posts)
	if r.URL.Query().Get("recent") == "true" {
		component := views.RecentPosts(posts)
		component.Render(r.Context(), w)
	} else {
		// For the main content area
		component := views.BlogPostList(posts)
		component.Render(r.Context(), w)
	}
}

func (h *Handler) GetBlogPost(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	post, htmlContent, err := h.BlogService.GetBlogPost(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Error fetching post", http.StatusInternalServerError)
		return
	}

	component := views.BlogContent(post, htmlContent)
	component.Render(r.Context(), w)
}

func (h *Handler) NewBlogPostForm(w http.ResponseWriter, r *http.Request) {
	component := views.BlogPostForm()
	component.Render(r.Context(), w)
}

// This handler needs to change to store Markdown content
func (h *Handler) CreateBlogPost(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content") // This is now Markdown content

	post, err := h.BlogService.CreateBlogPost(r.Context(), title, content)
	if err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	// Convert the Markdown content to HTML for display
	htmlContent, err := h.BlogService.MarkdownService.ConvertAndSanitize(post.Content)
	if err != nil {
		http.Error(w, "Error processing post content", http.StatusInternalServerError)
		return
	}

	component := views.BlogContent(post, htmlContent)
	component.Render(r.Context(), w)
}

func (h *Handler) Admin(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*auth.AuthenticatedUser)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	// Check if the user is an admin
	if user.IsAdmin {
		fmt.Fprintf(w, "Welcome, admin! This is the home page.")
	} else {
		fmt.Fprintf(w, "Welcome, user! This is the home page.")
	}
	views.Adminpage().Render(context.Background(), w)
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*auth.AuthenticatedUser)
	if !ok {
		// Handle the case where user is not in the context
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	// Check if the user is an admin
	if user.IsAdmin {
		fmt.Fprintf(w, "Welcome, admin! This is the home page.")
	} else {
		fmt.Fprintf(w, "Welcome, user! This is the home page.")
	}
	views.Homepage(user.IsAdmin, user.Email).Render(context.Background(), w)
}
