package server

import (
	"encoding/json"
	"log"
	"net/http"

	"homepage/cmd/web"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.healthHandler)
	fileServer := http.FileServer(http.Dir("./cmd/web/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets", fileServer))
	mux.Handle("/", templ.Handler(web.Home()))
	mux.Handle("GET /bio", templ.Handler(web.Bio()))
	mux.Handle("GET /about", templ.Handler(web.About()))
	mux.Handle("GET /cv", templ.Handler(web.CV()))
	return mux
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
