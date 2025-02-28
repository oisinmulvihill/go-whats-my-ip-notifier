package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

// Create the database and the ip_log table if they don't exist.
func Init(storageFilePath string) (*Storage, error) {

	db, err := sql.Open("sqlite3", storageFilePath)
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS ip_log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip TEXT NOT NULL UNIQUE,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) AddIPIfNotPresent(ip string) error {
	_, err := s.db.Exec("INSERT OR IGNORE INTO ip_log (ip) VALUES (?)", ip)
	return err
}
