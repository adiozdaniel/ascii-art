package routes

import (
	"net/http"
	"strings"
	"time"

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

// Simple in-memory session store
var sessions = map[string]string{} // TODO map[sessionID]userID

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

	mux.HandleFunc("/set-cookie", SetCookieHandler)

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

// SetCookieHandler sets a cookie to the client
func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(15 * time.Minute)
	cookie := http.Cookie{
		Name:     "BaseCookie",
		Value:    "ClientSession",
		Expires:  expiration,
		HttpOnly: true,
		Secure:   false, // TODO Set to true before deploying to production
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.Write([]byte("Cookie has been set"))
}

// GetSessionHandler retrieves session data based on the session cookie
func GetSessionHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}

	sessionID := cookie.Value
	userID, ok := sessions[sessionID]
	if !ok {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Session is valid for user: " + userID))
}
