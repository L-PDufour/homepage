package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/a-h/templ"
	efs "homepage"
	"homepage/internal/views"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.FS(efs.Files))
	mux.Handle("/assets/", fileServer)
	mux.Handle("/", templ.Handler(views.Bio()))
	mux.HandleFunc("/health", s.healthHandler)
	mux.Handle("GET /projects", templ.Handler(views.Projects()))
	// mux.HandleFunc("GET /blog", BlogHandler)
	mux.Handle("GET /cv", templ.Handler(views.CV()))
	mux.Handle("GET /kids", templ.Handler(views.Kids()))
	mux.HandleFunc("POST /add", s.addPerson)
	mux.HandleFunc("GET /list", s.listPeople)
	return mux
}

func (s *Server) addPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.db.AddPerson(person.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Person added successfully"})
}

func (s *Server) listPeople(w http.ResponseWriter, r *http.Request) {
	people, err := s.db.ListPeople()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	err := s.db.HealthCheck()
	status := "healthy"
	if err != nil {
		status = "unhealthy"
	}

	jsonResp, err := json.Marshal(map[string]string{"status": status})
	if err != nil {
		log.Printf("error handling JSON marshal. Err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
