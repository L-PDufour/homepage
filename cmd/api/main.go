package main

import (
	"homepage/internal/server"
	"log"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	log.Printf("Server starting on port %d", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
