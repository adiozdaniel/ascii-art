package routes

import (
	"net/http"
	"strings"

	"github.com/adiozdaniel/ascii-art/internals/handlers"
	"github.com/adiozdaniel/ascii-art/internals/renders"
)

// Allowed routes
var allowedRoutes = map[string]bool{
	"/":          true,
	"/ascii-art": true,
	"/about":     true,
	"/contact":   true,
}

// RouteChecker is a middleware that checkes allowed routes
func RouteChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/static/") {
			next.ServeHTTP(w, r)
			return
		}

		if _, ok := allowedRoutes[r.URL.Path]; !ok {
			handlers.Repo.NotFoundHandler(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// RegisterRoutes manages the routes
func RegisterRoutes(mux *http.ServeMux) {
	staticDir := renders.GetProjectRoot("views", "static")
	fs := http.FileServer(http.Dir(staticDir))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Repo.HomeHandler(w, r)
	})

	mux.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		handlers.Repo.SubmitHandler(w, r)
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		handlers.Repo.AboutHandler(w, r)
	})

	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		handlers.Repo.ContactHandler(w, r)
	})
}
