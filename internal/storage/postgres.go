package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// PostgresDB holds the database connection pool.
type PostgresDB struct {
	DB *sql.DB
}

// NewPostgresDB initializes a new PostgreSQL connection pool.
func NewPostgresDB(host, port, user, password, dbname string) (*PostgresDB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Failed to connect to PostgreSQL: %v", err)
		return nil, err
	}

	// Test the database connection
	if err := db.Ping(); err != nil {
		log.Printf("Failed to ping PostgreSQL: %v", err)
		return nil, err
	}

	log.Println("Connected to PostgreSQL successfully")
	return &PostgresDB{DB: db}, nil
}

// Close closes the database connection pool.
func (pg *PostgresDB) Close() error {
	return pg.DB.Close()
}

// ExampleQuery is a placeholder for executing a query.
func (pg *PostgresDB) ExampleQuery() error {
	query := "SELECT NOW()" // Example query
	row := pg.DB.QueryRow(query)

	var currentTime string
	if err := row.Scan(&currentTime); err != nil {
		log.Printf("Query failed: %v", err)
		return err
	}

	log.Printf("Current time from database: %s", currentTime)
	return nil
}
