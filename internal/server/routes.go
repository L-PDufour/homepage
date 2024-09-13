package server

import (
	efs "homepage"
	"homepage/internal/views"
	"net/http"

	"github.com/a-h/templ"
)

type Route struct {
	Method  string
	Path    string
	Handler http.Handler
}

func (s *Server) registerRoute(mux *http.ServeMux, method, path string, handler http.Handler) {
	switch method {
	case "GET":
		mux.Handle(path, handler)
	default:
		mux.Handle(path, handler)
	}
}

func (s *Server) registerRoutes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(efs.Files))
	mux.Handle("/assets/", fileServer)

	routes := []Route{
		{"GET", "/bio", s.Handler.UnifiedView("bio")},
		{"GET", "/projects", s.Handler.UnifiedView("project")},
		{"GET", "/blog", s.Handler.UnifiedView("blog")},
		{"GET", "/cv", templ.Handler(views.CV())},
		{"GET", "/kids", templ.Handler(views.Kids())},
		{"GET", "/admin", http.HandlerFunc(s.Handler.Admin)},
		{"GET", "/admin/auth", s.Handler.AdminAuth()},

		{"GET", "/content/list", http.HandlerFunc(s.Handler.ListContentHandler)},
		{"GET", "/content/view", http.HandlerFunc(s.Handler.ViewContentHandler)},
		{"GET", "/content/new", http.HandlerFunc(s.Handler.NewContentFormHandler)},
		{"POST", "/content/create", http.HandlerFunc(s.Handler.CreateContentHandler)},
		{"GET", "/content/edit", http.HandlerFunc(s.Handler.EditContentHandler)},
		{"GET", "/content/get", http.HandlerFunc(s.Handler.GetContentHandler)},
		{"POST", "/content/update", http.HandlerFunc(s.Handler.UpdateContentHandler)},
		{"DELETE", "/content/delete", http.HandlerFunc(s.Handler.DeleteContentHandler)},
	}

	for _, route := range routes {
		s.registerRoute(mux, route.Method, route.Path, route.Handler)
	}

	mux.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/bio", http.StatusFound)
	})

	return mux
}
