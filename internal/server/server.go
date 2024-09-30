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

	"homepage/internal/auth"
	"homepage/internal/database"
	"homepage/internal/handler"
	"homepage/internal/middleware"
	"homepage/internal/service"

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
		middleware.HTMXMiddleware,
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
	var host, port, user, password, dbname string

	if os.Getenv("GO_ENV") == "production" {
		host = os.Getenv("PROD_DB_HOST")
		port = os.Getenv("PROD_DB_PORT")
		user = os.Getenv("PROD_DB_USER")
		password = os.Getenv("PROD_DB_PASSWORD")
		dbname = os.Getenv("PROD_DB_NAME")
	} else {
		host = os.Getenv("DEV_DB_HOST")
		port = os.Getenv("DEV_DB_PORT")
		user = os.Getenv("DEV_DB_USER")
		password = os.Getenv("DEV_DB_PASSWORD")
		dbname = os.Getenv("DEV_DB_NAME")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return sql.Open("postgres", connStr)
}
