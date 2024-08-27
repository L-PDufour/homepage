package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Service interface {
	Close() error
	HealthCheck() error
	AddPerson(name string) error
	ListPeople() ([]Person, error)
}

type service struct {
	db *sql.DB
}

func (s *service) AddPerson(name string) error {
	_, err := s.db.Exec("INSERT INTO example_table (name) VALUES ($1)", name)
	return err
}

func (s *service) ListPeople() ([]Person, error) {
	rows, err := s.db.Query("SELECT id, name FROM example_table")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []Person
	for rows.Next() {
		var p Person
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, err
		}
		people = append(people, p)
	}
	return people, nil
}

func New() (Service, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Construct the connection string from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open the database
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS example_table (id SERIAL PRIMARY KEY, name TEXT)`)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}

	log.Println("Database opened and connection verified")
	return &service{db: db}, nil
}

func (s *service) Close() error {
	return s.db.Close()
}

func (s *service) HealthCheck() error {
	var result int
	err := s.db.QueryRow("SELECT 1").Scan(&result)
	if err != nil || result != 1 {
		return fmt.Errorf("database health check failed: %v", err)
	}
	return nil
}
