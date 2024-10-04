package server

import (
	"database/sql"
	"grocery-purchase/config" // Ensure that you import the config package
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// InitDB initializes the database connection using the provided configuration.
func InitDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
