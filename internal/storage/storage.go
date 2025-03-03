package storage

import (
	"database/sql"
	"log"

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

	// ip column is not unique as the same IP address could be set at a later
	// datecan be added multiple times.
	query := `
	CREATE TABLE IF NOT EXISTS ip_log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) AddAddress(ip string) error {
	_, err := s.db.Exec("INSERT INTO ip_log (ip) VALUES (?)", ip)
	return err
}

func (s *Storage) CurrentIP() (string, error) {
	var err error
	var ipAddress string = ""
	var timestamp string

	result := s.db.QueryRow("select ip, timestamp from ip_log order by timestamp desc limit 1")
	if err = result.Scan(&ipAddress, &timestamp); err != nil {
		if err != sql.ErrNoRows {
			log.Println("problem getting the current ip address:", err)
		} else {
			// No rows is fine, ignore this error.
			err = nil
		}
	}

	return ipAddress, err
}
