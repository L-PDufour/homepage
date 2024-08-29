package handlers

import (
	"bytes"
	"homepage/internal/database"
	"homepage/internal/models"
	"homepage/internal/views"
	"net/http"
	"strconv"

	"github.com/yuin/goldmark"
)

type Handlers struct {
	db *database.Queries
}

func NewHandlers(db *database.Queries) *Handlers {
	return &Handlers{db: db}
}

func (h *Handlers) HandleBio() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (h *Handlers) HandleBlogContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 32)
		post, err := h.db.GetBlogPost(r.Context(), int32(postID))
		if err != nil {
			http.Error(w, "Error fetching blog post", http.StatusInternalServerError)
			return
		}

		var buf bytes.Buffer
		if err := goldmark.Convert([]byte(post.Content), &buf); err != nil {
			http.Error(w, "Error converting Markdown to HTML", http.StatusInternalServerError)
			return
		}

		// Convert the buffer content to a string
		htmlContent := buf.String()

		// Add a script to log the content to the console
		script := "<script>console.log(" + strconv.Quote(htmlContent) + ");</script>"
		htmlContent += script

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(htmlContent))
	}
}

func (h *Handlers) HandleHomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbPosts, err := h.db.GetRecentBlogPosts(r.Context(), 5) // Fetch 5 most recent posts
		if err != nil {
			http.Error(w, "Error fetching blog posts", http.StatusInternalServerError)
			return
		}

		var blogPosts []models.BlogPostData
		for _, post := range dbPosts {
			blogPosts = append(blogPosts, models.BlogPostData{
				ID:    int(post.ID),
				Title: post.Title,
			})
		}

		data := models.HomePageData{
			BlogPosts: blogPosts,
		}

		views.Test(data).Render(r.Context(), w)
	}
}
