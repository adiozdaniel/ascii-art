package routes

import (
	"net/http"
	"strings"

	appconfig "github.com/adiozdaniel/ascii-art/internals/config"
	"github.com/adiozdaniel/ascii-art/internals/handlers"
)

// get the app state manager
var (
	sm  = appconfig.GetStateManager()
	app = sm.GetInput()
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
			handlers.NotFoundHandler(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// RegisterRoutes manages the routes
func RegisterRoutes(mux *http.ServeMux) {
	staticDir := app.GetProjectRoot("views", "static")
	fs := http.FileServer(http.Dir(staticDir))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r)
	})

	mux.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		handlers.SubmitHandler(w, r)
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		handlers.AboutHandler(w, r)
	})

	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		handlers.ContactHandler(w, r)
	})
}
