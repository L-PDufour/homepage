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

	// Define your routes directly in the mux
	mux.HandleFunc("/bio", s.Handler.ListContent("bio"))
	mux.HandleFunc("/projects", s.Handler.ListContent("project"))
	mux.HandleFunc("/blog", s.Handler.ListContent("blog"))
	mux.HandleFunc("/cv", s.Handler.ServeResume())
	// mux.Handle("/kids", templ.Handler(views.Kids()))
	mux.Handle("/plan", templ.Handler(views.Plan()))
	mux.HandleFunc("/admin", http.HandlerFunc(s.Handler.Admin))
	mux.HandleFunc("/admin/auth", s.Handler.AdminAuth())

	mux.HandleFunc("/content/list", s.Handler.ListContent(""))
	mux.HandleFunc("/content", s.Handler.GetContent())
	mux.HandleFunc("DELETE /content", s.Handler.DeleteContent())
	mux.HandleFunc("/content/new", s.Handler.GetForm())
	mux.HandleFunc("POST /content/new", s.Handler.CreateContent())
	mux.HandleFunc("/content/update", s.Handler.GetUpdateForm())
	mux.HandleFunc("POST /content/update", s.Handler.UpdateContent())

	mux.Handle("/game", templ.Handler(views.Game()))
	mux.HandleFunc("POST /game/input", s.Handler.Test())

	mux.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/bio", http.StatusFound)
	})
	return mux
}
