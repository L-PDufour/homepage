package handler

import (
	"encoding/json"
	"homepage/internal/views"
	"net/http"
	"strconv"
)

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

	if r.Header.Get("HX-Request") == "true" {
		component := views.BlogPostList(posts)
		component.Render(r.Context(), w)
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(posts); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
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
