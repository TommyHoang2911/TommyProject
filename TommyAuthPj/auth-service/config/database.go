package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // postgres driver
)

// InitDB reads configuration from the environment and opens a connection
// to the PostgreSQL database. It returns a *sql.DB that must be closed
// by the caller when the application shuts down.
func InitDB() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// default placeholder pointing to a local database. Users should
		// override this with a real connection string in production.
		dsn = "postgres://tommyhoang:Aa@123456@localhost:5432/authdb?sslmode=disable"
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return db, nil
}
