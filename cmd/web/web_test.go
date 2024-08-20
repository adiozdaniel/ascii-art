package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestServer runs tests ont the web server
func TestServer(t *testing.T) {
	go main()

	time.Sleep(1 * time.Second)

	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()

	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK, got %v", resp.Status)
	}

	TriggerShutdown()
	time.Sleep(1 * time.Second)

	_, err = http.Get("http://localhost:8080/")
	if err == nil {
		t.Fatalf("Expected server to be shut down, but it's still running")
	}
}
