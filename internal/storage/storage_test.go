package storage

import (
	"database/sql"
	"testing"

	"github.com/oisinmulvihill/go-whats-my-ip-notifier/internal/assert"
)

func TestCurrentAddressAndAddAddress(t *testing.T) {
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

	var currentAddress string
	currentAddress, err = storage.CurrentIP()
	assert.Equal(t, err, nil)
	assert.Equal(t, currentAddress, "")

	// Add the IP address for the first time
	err = storage.AddAddress("192.168.0.1")
	assert.Equal(t, err, nil)

	currentAddress, err = storage.CurrentIP()
	assert.Equal(t, err, nil)
	assert.Equal(t, currentAddress, "192.168.0.1")
}
