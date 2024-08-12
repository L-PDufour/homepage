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

	db database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		// db:   database.New(),
	}

	// Declare Server config
	stack := middleware.CreateStack(
		middleware.AllowCors,
		middleware.CheckPermissions,
		middleware.Logging,
	)

	loadRoutes := NewServer.RegisterRoutes()
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      stack(loadRoutes),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
