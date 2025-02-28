package storage

import (
	"database/sql"
	"testing"

	"github.com/oisinmulvihill/go-whats-my-ip-notifier/internal/assert"
)

func TestAddIPIfNotPresent(t *testing.T) {
	var row *sql.Row
	var err error
	var count int

	storage, err := Init(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer storage.db.Close()

	// Check that the ip_log table is empty
	row = storage.db.QueryRow("SELECT COUNT(*) FROM ip_log")
	if err := row.Scan(&count); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, count, 0)

	// Add the IP address for the first time
	err = storage.AddIPIfNotPresent("192.168.0.1")
	assert.Equal(t, err, nil)
	row = storage.db.QueryRow("SELECT COUNT(*) FROM ip_log")
	if err = row.Scan(&count); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, count, 1)

	// Adding the same IP address should result is more than one entry
	err = storage.AddIPIfNotPresent("192.168.0.1")
	assert.Equal(t, err, nil)

	row = storage.db.QueryRow("SELECT COUNT(*) FROM ip_log")
	if err = row.Scan(&count); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, count, 1)
}
