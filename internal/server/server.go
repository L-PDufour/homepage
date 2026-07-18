package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"homepage"

	_ "github.com/joho/godotenv/autoload"

	"homepage/internal/auth"
	"homepage/internal/database"
	"homepage/internal/handler"
	"homepage/internal/middleware"
	"homepage/internal/service"

	_ "modernc.org/sqlite"
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
	contentService := service.NewContentService(dbQueries)

	handler := handler.NewHandler(dbQueries, contentService)
	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		log.Fatalf("Failed to create authenticator: %v", err)
	}

	newServer := &Server{
		Port:    port,
		DB:      dbQueries,
		Handler: handler,
	}

	stack := middleware.CreateStack(
		middleware.AllowCors,
		middleware.WithAuthenticator(authenticator),
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
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "dev.db"
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		seed, err := homepage.SeedDB.ReadFile("dev.db")
		if err != nil {
			return nil, err
		}
		if err := os.WriteFile(dbPath, seed, 0o644); err != nil {
			return nil, err
		}
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	schema, err := homepage.SQLFiles.ReadFile("sql/schema.sql")
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(string(schema)); err != nil {
		return nil, err
	}

	return db, nil
}
