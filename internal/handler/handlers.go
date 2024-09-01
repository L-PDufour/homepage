package handler

import (
	"homepage/internal/blog"
	"homepage/internal/database"
	"homepage/internal/markdown"
	"net/http"
	"strconv"
)

type Handler struct {
	DB          *database.Queries
	MD          markdown.MarkdownConverter
	BlogService blog.BlogService
}

func NewHandler(db *database.Queries, md markdown.MarkdownConverter) *Handler {
	return &Handler{
		DB: db,
		MD: md,
	}
}

func (h *Handler) HandleBio() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (h *Handler) HandleBlogContent(w http.ResponseWriter, r *http.Request) {
	postID, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 32)
	post, err := h.DB.GetBlogPost(r.Context(), int32(postID))
	if err != nil {
		http.Error(w, "Error fetching blog post", http.StatusInternalServerError)
		return
	}

	htmlContent, err := h.MD.ConvertAndSanitize(post.Content)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(htmlContent))
}

// func (h *Handlers) HandleHomePage() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		dbPosts, err := h.db.GetRecentBlogPosts(r.Context(), 5) // Fetch 5 most recent posts
// 		if err != nil {
// 			http.Error(w, "Error fetching blog posts", http.StatusInternalServerError)
// 			return
// 		}
//
// 		var blogPosts []models.BlogPostData
// 		for _, post := range dbPosts {
// 			blogPosts = append(blogPosts, models.BlogPostData{
// 				ID:    int(post.ID),
// 				Title: post.Title,
// 			})
// 		}
//
// 		data := models.HomePageData{
// 			BlogPosts: blogPosts,
// 		}
//
// 		views.Test(data).Render(r.Context(), w)
// 	}
// }
