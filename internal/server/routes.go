package server

import (
	efs "homepage"
	"homepage/internal/views"
	"net/http"

	"github.com/a-h/templ"
)

func (s *Server) registerRoutes() http.Handler {
	mux := http.NewServeMux()

	// Static file server
	fileServer := http.FileServer(http.FS(efs.Files))
	mux.Handle("/assets/", fileServer)

	// Main routes
	mux.Handle("/", templ.Handler(views.Bio()))
	mux.Handle("GET /projects", templ.Handler(views.Projects()))
	mux.Handle("GET /cv", templ.Handler(views.CV()))
	mux.Handle("GET /kids", templ.Handler(views.Kids()))

	// Blog routes
	mux.HandleFunc("GET /blogy", s.Handler.BlogPage)
	// mux.HandleFunc("GET /blog/{id}", s.Handler.GetBlogPost)
	http.HandleFunc("/blog", s.Handler.GetBlogPost)
	http.HandleFunc("/blog/posts", s.Handler.GetBlogPosts)
	http.HandleFunc("/blog/new", s.Handler.NewBlogPostForm)
	http.HandleFunc("/blog/create", s.Handler.CreateBlogPost)
	// mux.HandleFunc("GET /blog/{id}/edit", s.Handler.EditBlogPostForm)
	// mux.HandleFunc("PUT /blog/{id}", s.Handler.UpdateBlogPost)
	// mux.HandleFunc("DELETE /blog/{id}", s.Handler.DeleteBlogPost)

	return mux
}
