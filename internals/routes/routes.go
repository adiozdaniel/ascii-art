package routes

import (
	"net/http"

	"github.com/adiozdaniel/ascii-art/internals/handlers"
)

// Allowed routes
var allowedRoutes = map[string]bool{
	"/":                       true,
	"/ascii-art":              true,
	"/static/styles/main.css": true,
	"/static/styles/home.css": true,
}

// RouteChecker is a middleware that checkes allowed routes
func RouteChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := allowedRoutes[r.URL.Path]; !ok {
			handlers.NotFoundHandler(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// RegisterRoutes manages the routes
func RegisterRoutes(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("views/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r)
	})

	mux.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.SubmitHandler(w, r)
		}
	})
}
