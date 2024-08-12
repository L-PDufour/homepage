package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mutecomm/go-sqlcipher/v4"
	"log"
	"os"
	"strconv"
	"time"
)

type Service interface {
	Health() map[string]string
	Close() error
}

type service struct {
	db *sql.DB
}

var (
	dbPath     = os.Getenv("DB_PATH")
	dbKey      = os.Getenv("DB_ENCRYPTION_KEY")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Enable encryption
	_, err = db.Exec(fmt.Sprintf("PRAGMA key = '%s'", dbKey))
	if err != nil {
		log.Fatalf("Failed to set encryption key: %v", err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Printf("Database health check failed: %v", err)
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)

	// SQLite-specific checks
	var pageCount, pageSize int
	err = s.db.QueryRow("PRAGMA page_count").Scan(&pageCount)
	if err == nil {
		stats["page_count"] = strconv.Itoa(pageCount)
	}
	err = s.db.QueryRow("PRAGMA page_size").Scan(&pageSize)
	if err == nil {
		stats["page_size"] = strconv.Itoa(pageSize)
	}

	// Database size (approximate)
	if pageCount > 0 && pageSize > 0 {
		dbSize := pageCount * pageSize
		stats["db_size_bytes"] = strconv.Itoa(dbSize)
	}

	return stats
}

func (s *service) Close() error {
	log.Printf("Closing SQLite database connection")
	return s.db.Close()
}
