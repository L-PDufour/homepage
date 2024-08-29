package server

import (
	efs "homepage"
	"homepage/internal/views"
	"net/http"

	"github.com/a-h/templ"
)

func (s *Server) registerRoutes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(efs.Files))
	mux.Handle("/assets/", fileServer)

	mux.Handle("/", templ.Handler(views.Bio()))
	mux.Handle("GET /projects", templ.Handler(views.Projects()))
	mux.Handle("GET /cv", templ.Handler(views.CV()))
	mux.Handle("GET /kids", templ.Handler(views.Kids()))
	mux.Handle("GET /blog", templ.Handler(views.Blog()))
	mux.HandleFunc("/bloggy", s.handlers.HandleBlogContent())
	mux.HandleFunc("/test", s.handlers.HandleHomePage())

	return mux
}
