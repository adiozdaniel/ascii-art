package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adiozdaniel/ascii-art/internals/middlewares"
)

// TestRunWeb verifies that the web server starts and responds with a status OK.
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

// TestRoutes checks the status codes for various routes.
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
		{"/logout", http.StatusOK},
		{"/nonexistent", http.StatusOK},
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

// TestMiddleware checks the middleware functionality.
func TestMiddlewares(t *testing.T) {
	middlewares.NewMiddlewares(sessionManager)

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	routeChecker := middlewares.SessionMiddleware(
		sessionManager)(middlewares.RouteChecker(testHandler))

	tests := []struct {
		path               string
		expectedStatusCode int
	}{
		{"/", http.StatusSeeOther},
		{"/ascii-art", http.StatusSeeOther},
		{"/about", http.StatusSeeOther},
		{"/contact", http.StatusOK},
		{"/login", http.StatusOK},
		{"/logout", http.StatusSeeOther},
		{"/nonexistent", http.StatusSeeOther},
	}

	for _, tt := range tests {
		req, err := http.NewRequest("GET", tt.path, nil)
		if err != nil {
			t.Fatalf("Failed to create request for %v: %v", tt.path, err)
		}

		recorder := httptest.NewRecorder()
		routeChecker.ServeHTTP(recorder, req)

		if recorder.Code != tt.expectedStatusCode {
			t.Errorf("For path %v, expected status %v, got %v", tt.path, tt.expectedStatusCode, recorder.Code)
		}
	}
}
