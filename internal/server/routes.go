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
	mux.Handle("GET /blogy", templ.Handler(views.Blog()))
	// mux.HandleFunc("GET /blog/{id}", s.Handler.GetBlogPost)
	mux.HandleFunc("GET /blog", s.Handler.HandleBlogContent)
	mux.HandleFunc("GET /blog/new", s.Handler.NewBlogPostForm)
	mux.HandleFunc("POST /blog", s.Handler.CreateBlogPost)
	mux.HandleFunc("GET /blog/{id}/edit", s.Handler.EditBlogPostForm)
	mux.HandleFunc("PUT /blog/{id}", s.Handler.UpdateBlogPost)
	mux.HandleFunc("DELETE /blog/{id}", s.Handler.DeleteBlogPost)

	return mux
}
