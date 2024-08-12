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

// Test runs tests for routes.
func TestRoutes(t *testing.T) {
	handler, err := runWeb()
	if err != nil {
		t.Fatalf("Failed to run web: %v", err)
	}

	server := httptest.NewServer(handler)
	defer server.Close()

	tests := []struct {
		path               string
		expectedStatusCode int
	}{
		{"/", http.StatusOK},
		{"/ascii-art", http.StatusOK},
		{"/about", http.StatusOK},
		{"/contact", http.StatusOK},
		{"/login", http.StatusOK},
		{"/logout", http.StatusSeeOther},
		{"/error", http.StatusNotFound},
		{"/nonexistent", http.StatusNotFound},
	}

	for _, tt := range tests {
		resp, err := http.Get(server.URL + tt.path)
		if err != nil {
			t.Fatalf("Failed to make request to %v: %v", tt.path, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != tt.expectedStatusCode {
			t.Fatalf("For path %v, expected status %v, got %v", tt.path, tt.expectedStatusCode, resp.Status)
		}
	}
}
