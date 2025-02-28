package public

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oisinmulvihill/go-whats-my-ip-notifier/internal/assert"
)

func TestGetAddressOKCase(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestIPAddressOKCase(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			t.Errorf("Expected to request '/', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`101.102.103.104`))
	}))
	defer server.Close()

	address, error := IPAddress(server.URL)
	assert.Equal(t, error, nil)
	assert.Equal(t, address, "101.102.103.104")
}
