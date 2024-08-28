package server

import (
	"encoding/json"
	"net/http"

	efs "homepage"
	"homepage/internal/database"
	"homepage/internal/views"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.FS(efs.Files))
	mux.Handle("/assets/", fileServer)
	mux.Handle("/", templ.Handler(views.Bio()))
	mux.Handle("GET /projects", templ.Handler(views.Projects()))
	// mux.Handle("GET /blog", templ.Handler(views.Blog()))
	mux.Handle("GET /cv", templ.Handler(views.CV()))
	mux.Handle("GET /kids", templ.Handler(views.Kids()))
	mux.HandleFunc("/ListAuthor", GetAuthorsHandler(s.db)) // Pass s.queries

	return mux
}

func GetAuthorsHandler(q *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		authors, err := q.ListAuthors(ctx)
		if err != nil {
			http.Error(w, "Error fetching authors", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(authors); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}
