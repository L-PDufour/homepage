package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/mutecomm/go-sqlcipher/v4"
	"log"
	"net/url"
	"os"
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
	_, err := s.db.Exec("INSERT INTO example_table (name) VALUES (?)", name)
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

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/database.db" // Default value
	}

	dbKey := os.Getenv("DB_ENCRYPTION_KEY")
	if dbKey == "" {
		return nil, fmt.Errorf("DB_ENCRYPTION_KEY must be set")
	}

	// Escape the key and create the connection string
	escapedKey := url.QueryEscape(dbKey)
	connStr := fmt.Sprintf("%s?_pragma_key=%s&_pragma_cipher_page_size=4096", dbPath, escapedKey)

	// Open the database
	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS example_table (id INTEGER PRIMARY KEY, name TEXT)`)
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
