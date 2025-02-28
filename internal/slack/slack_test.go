package slack

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSlackSendMessageOKCase(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected to request method 'GET', got: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`ok`))
	}))
	defer server.Close()

	err := SendMessage(server.URL, "username", "the message to send")

	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}
}

func TestSlackSendMessageErrorUnauthorised(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			t.Errorf("Expected to request '/', got: %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected to request method 'POST', got: %s", r.Method)
		}
		w.WriteHeader(http.StatusForbidden)
	}))
	defer server.Close()

	error := SendMessage(server.URL, "username", "the message to send")

	expected_error := "failed to call slack webhook '403'"
	if error.Error() != expected_error {
		t.Errorf("Expected error: %+v, got: %+v", expected_error, error)
	}
}
