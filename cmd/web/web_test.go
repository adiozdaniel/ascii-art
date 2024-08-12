package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/adiozdaniel/ascii-art/internals/middlewares"
	"github.com/adiozdaniel/ascii-art/internals/models"
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

// Mocked appData methods
type MockAppData struct {
	flags      map[string]string
	bannerFile map[string]string
}

func (m *MockAppData) Init() {}

func TestMainFunction(t *testing.T) {
	// Setup mock appData
	mockAppData := &MockAppData{
		flags: map[string]string{
			"font":  "--standard",
			"input": "Ascii~",
		},
		bannerFile: map[string]string{
			"--standard": "banner.txt",
		},
	}

	// Replace global appData with mock
	models.NewInputData().Flags = mockAppData.flags
	models.NewInputData().BannerFile = mockAppData.bannerFile

	// Create a test server and client
	handler, err := runWeb()
	if err != nil {
		t.Fatalf("Failed to run web: %v", err)
	}
	server := httptest.NewServer(handler)
	defer server.Close()

	// Start the main function in a separate goroutine
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan struct{})
	go func() {
		defer close(done)
		main()
	}()

	// Simulate HTTP request
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK, got %v", resp.Status)
	}

	// Simulate shutdown signal
	time.Sleep(2 * time.Second) // Allow some time for the server to start
	stop <- syscall.SIGINT

	// Wait for main to finish
	<-done
}
