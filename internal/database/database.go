package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Service interface {
	Close() error
	GetQueries() *Queries
}

type service struct {
	*Queries
	db *sql.DB
}

func NewService() (Service, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	queries := New(db)

	return &service{
		Queries: queries,
		db:      db,
	}, nil
}

func (s *service) Close() error {
	return s.db.Close()
}

func (s *service) GetQueries() *Queries {
	return s.Queries
}
