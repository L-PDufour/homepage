package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"homepage/internal/database"
	"homepage/internal/middleware"
)

type Server struct {
	port int
	db   database.Service
}

func NewServer() (*http.Server, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid port: %v", err)
	}

	db, err := database.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}

	newServer := &Server{
		port: port,
		db:   db,
	}

	// Declare Server config
	stack := middleware.CreateStack(
		middleware.AllowCors,
		middleware.CheckPermissions,
		middleware.Logging,
	)

	loadRoutes := newServer.RegisterRoutes()
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      stack(loadRoutes),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, nil
}
