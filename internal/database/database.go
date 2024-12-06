package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	SeedSql() error

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type DBInstance struct {
	Db        *pgx.Conn
	DbContext *context.Context
}

var (
	database   = os.Getenv("BLUEPRINT_DB_DATABASE")
	password   = os.Getenv("BLUEPRINT_DB_PASSWORD")
	username   = os.Getenv("BLUEPRINT_DB_USERNAME")
	port       = os.Getenv("BLUEPRINT_DB_PORT")
	host       = os.Getenv("BLUEPRINT_DB_HOST")
	schema     = os.Getenv("BLUEPRINT_DB_SCHEMA")
	dbInstance *DBInstance
)

func New(context context.Context) *DBInstance {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)

	db, err := pgx.Connect(context, connStr)
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &DBInstance{
		Db:        db,
		DbContext: &context,
	}
	return dbInstance
}

func (s *DBInstance) SeedSql() error {
	projectEnv := os.Getenv("APP_ENV")

	if projectEnv != "local" {
		return nil
	}

	path := filepath.Join("./", "seed.sql")

	c, ioErr := os.ReadFile(path)
	if ioErr != nil {
		return ioErr
	}

	sql := string(c)

	_, err := s.Db.Exec(*s.DbContext, sql)

	if err != nil {
		return err
	}

	return nil
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *DBInstance) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.Db.Ping(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *DBInstance) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.Db.Close(*s.DbContext)
}
