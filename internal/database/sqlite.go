package database

import (
	"database/sql"
	"fmt"
	"os"
)

// Connection pool => collection of db connecttions
func InitDB(dataSourceName string) (*sql.DB, error) {
	err := os.MkdirAll("data", 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w: ", err)
	}

	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open databse: %w: ", err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, fmt.Errorf("failed to enable foreign keys: %w: ", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed ping database: %w: ", err)
	}

	return db, nil
}
