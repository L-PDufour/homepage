package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	_ "github.com/lib/pq"
	"homepage/internal/database"
	"homepage/internal/handlers"
	"homepage/internal/middleware"
)

type Server struct {
	port     int
	db       *database.Queries
	handlers *handlers.Handlers
}

func NewServer() (*http.Server, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid port: %v", err)
	}

	db, err := connectDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	dbQueries := database.New(db)

	newServer := &Server{
		port:     port,
		db:       dbQueries,
		handlers: handlers.NewHandlers(dbQueries),
	}

	stack := middleware.CreateStack(
		middleware.AllowCors,
		middleware.CheckPermissions,
		middleware.Logging,
	)

	mux := newServer.registerRoutes()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      stack(mux),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, nil
}

func connectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	return sql.Open("postgres", connStr)
}
