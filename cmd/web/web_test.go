package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test runWeb function
func TestRunWeb(t *testing.T) {
	handler, err := runWeb()
	if err != nil {
		t.Fatalf("Failed to run web: %v", err)
	}

	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK, got %v", resp.Status)
	}
}
