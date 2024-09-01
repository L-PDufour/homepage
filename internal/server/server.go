package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"homepage/internal/blog"
	"homepage/internal/database"
	"homepage/internal/handler"
	"homepage/internal/markdown"
	"homepage/internal/middleware"

	_ "github.com/lib/pq"
)

type Server struct {
	Port    int
	DB      *database.Queries
	Handler *handler.Handler
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

	logger := log.New(os.Stdout, "markdown: ", log.LstdFlags)

	mdService := markdown.NewMarkdownService(logger) // Now returns *MarkdownService

	blogService := blog.NewBlogService(dbQueries, mdService) // Pass the pointer

	handler := handler.NewHandler(dbQueries, mdService, blogService)

	newServer := &Server{
		Port:    port,
		DB:      dbQueries,
		Handler: handler,
	}

	stack := middleware.CreateStack(
		middleware.AllowCors,
		middleware.CheckPermissions,
		middleware.Logging,
	)

	mux := newServer.registerRoutes()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.Port),
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
