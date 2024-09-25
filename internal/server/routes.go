package server

import (
	efs "homepage"
	"net/http"
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
	mux.HandleFunc("/admin", http.HandlerFunc(s.Handler.Admin))
	mux.HandleFunc("/admin/auth", s.Handler.AdminAuth())

	mux.HandleFunc("/content/list", s.Handler.ListContent(""))
	mux.HandleFunc("/content", s.Handler.GetContent())
	// mux.HandleFunc("POST /content", s.Handler...)
	// mux.HandleFunc("PUT /content/", s.Handler...)
	// mux.HandleFunc("DELETE /content/", s.Handler.DeleteContent)
	// mux.HandleFunc("/content/view", s.Handler.ViewContentHandler)
	// Catch-all redirect to /bio
	mux.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/bio", http.StatusFound)
	})
	return mux
}
