package server

import (
	efs "homepage"
	"homepage/internal/views"
	"net/http"
	"os"
	"sync"

	"github.com/a-h/templ"
)

type AuthenticatedUser struct {
	Email   string
	IsAdmin bool
}

var authenticatedUsers = sync.Map{}
var isProduction = os.Getenv("GO_ENV") == "production"

func (s *Server) registerRoutes() http.Handler {
	mux := http.NewServeMux()

	// Static file server
	fileServer := http.FileServer(http.FS(efs.Files))
	mux.Handle("/assets/", fileServer)

	// View routes
	mux.Handle("/", templ.Handler(views.Bio()))
	mux.Handle("GET /projects", templ.Handler(views.Projects()))
	mux.Handle("GET /cv", templ.Handler(views.CV()))
	mux.Handle("GET /kids", templ.Handler(views.Kids()))
	mux.Handle("GET /blog", templ.Handler(views.Blog()))

	// Auth
	mux.HandleFunc("GET /admin", s.Handler.Admin)
	mux.HandleFunc("GET /home", s.Handler.Home)

	// Blog Crud
	mux.HandleFunc("GET /blog/posts", s.Handler.GetBlogPosts)
	mux.HandleFunc("GET /blog/post", s.Handler.GetBlogPost)
	mux.HandleFunc("GET /blog/new", s.Handler.NewBlogPostForm)
	mux.HandleFunc("POST /blog/create", s.Handler.CreateBlogPost)
	// In your server setup
	return mux
}
