package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	SSLMode  string
}

// Open opens a new database connection and returns a *sql.DB object
func Open(config Config) (*sql.DB, error) {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.SSLMode,
	)
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
