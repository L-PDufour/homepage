package handler

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"homepage/internal/database"
)

func (h *Handler) GetBlogPost(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}
	id := parts[1]

	// Convert ID to integer
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Query the database to get the blog post by ID
	ctx := r.Context()
	post, err := h.DB.GetBlogPost(ctx, int32(intID))
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Blog post not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve blog post", http.StatusInternalServerError)
		}
		return
	}

	// Use the MarkdownService to convert and sanitize the content
	sanitizedContent, err := h.MD.ConvertAndSanitize(post.Content)
	if err != nil {
		http.Error(w, "Error processing Markdown content", http.StatusInternalServerError)
		return
	}

	// Set content type and write the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(sanitizedContent))
}

func (h *Handler) ListBlogPosts(ctx context.Context) []database.Post {
	posts, _ := h.DB.ListBlogPosts(ctx)
	return posts
}

// func (h *Handler) ListBlogPost(w http.ResponseWriter, r *http.Request) {
//
// 	ctx := r.Context()
// 	posts, err := h.DB.ListBlogPosts(ctx)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			http.Error(w, "Blog post not found", http.StatusNotFound)
// 		} else {
// 			http.Error(w, "Failed to retrieve blog post", http.StatusInternalServerError)
// 		}
// 		return
// 	}
//
// 	var buf bytes.Buffer
// 	for _, post := range posts {
// 		if err := goldmark.Convert([]byte(post.Content), &buf); err != nil {
// 			http.Error(w, "Error converting Markdown to HTML", http.StatusInternalServerError)
// 			return
// 		}
// 		buf.WriteString("<hr>")
// 	}
//
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(buf.Bytes())
// }

func (h *Handler) NewBlogPostForm(w http.ResponseWriter, r *http.Request) {
	// Render new blog post form
}

func (h *Handler) CreateBlogPost(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	// Insert new blog post into database
	// Redirect to the new blog post or blog list
}

func (h *Handler) EditBlogPostForm(w http.ResponseWriter, r *http.Request) {
	// Fetch blog post from database
	// Render edit form with current data
}

func (h *Handler) UpdateBlogPost(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	// Update blog post in database
	// Redirect to updated blog post
}

func (h *Handler) DeleteBlogPost(w http.ResponseWriter, r *http.Request) {
	// Delete blog post from database
	// Redirect to blog list
}
