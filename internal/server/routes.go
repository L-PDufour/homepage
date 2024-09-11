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
		mux.Handle(path, handler) // Handle other methods as needed
	}
}

func (s *Server) registerRoutes() http.Handler {
	mux := http.NewServeMux()

	// Static file server
	fileServer := http.FileServer(http.FS(efs.Files))
	mux.Handle("/assets/", fileServer)

	// Define your routes
	routes := []Route{
		{"GET", "/bio", http.HandlerFunc(s.Handler.BioHandler)}, // Only one registration for "/"
		{"GET", "/projects", http.HandlerFunc(s.Handler.ProjectsHandler)},
		{"GET", "/cv", templ.Handler(views.CV())},
		{"GET", "/kids", templ.Handler(views.Kids())},
		{"GET", "/blog", http.HandlerFunc(s.Handler.BlogPage)},
		// {"GET", "/blog", templ.Handler(views.Blog())},
		{"GET", "/admin", http.HandlerFunc(s.Handler.Admin)},
		{"GET", "/admin/auth", http.HandlerFunc(s.Handler.AdminAuth)},

		//Content CRUD routes
		{"GET", "/content/get", http.HandlerFunc(s.Handler.GetContentHandler)},
		{"GET", "/content/edit", http.HandlerFunc(s.Handler.EditContentHandler)},
		{"POST", "/content/update", http.HandlerFunc(s.Handler.UpdateContentHandler)},
		// Blog CRUD routes
		{"GET", "/blog/posts", http.HandlerFunc(s.Handler.GetBlogPosts)},
		{"GET", "/blog/post", http.HandlerFunc(s.Handler.GetBlogPost)},
		{"GET", "/blog/new", http.HandlerFunc(s.Handler.NewBlogPostForm)},
		{"POST", "/blog/create", http.HandlerFunc(s.Handler.CreateBlogPost)},
		// {"GET", "/blog/edit", http.HandlerFunc(s.Handler.EditBlogPostForm)},
		// {"POST", "/blog/update", http.HandlerFunc(s.Handler.UpdateBlogPost)},
		// {"POST", "/blog/delete", http.HandlerFunc(s.Handler.DeleteBlogPost)},
	}

	// Register all routes
	for _, route := range routes {
		s.registerRoute(mux, route.Method, route.Path, route.Handler)
	}

	// Catch-all handler for unhandled routes
	// mux.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "/bio", http.StatusFound)
	// })

	return mux
}
