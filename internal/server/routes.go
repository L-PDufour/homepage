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
	mux.HandleFunc("assets/index.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		fileServer.ServeHTTP(w, r)
	})

	// Serve WebAssembly files with the correct MIME type
	mux.HandleFunc("/index.wasm", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/wasm")
		fileServer.ServeHTTP(w, r)
	})
	mux.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/index.html") // Adjust the path as necessary
	})

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

	mux.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/bio", http.StatusFound)
	})
	return mux
}
