package middlewares

import (
	"net/http"
	"strings"

	"github.com/adiozdaniel/ascii-art/internals/handlers"
	"github.com/adiozdaniel/ascii-art/internals/models"
)

// get the app state manager
var (
	sm   = models.GetStateManager()
	repo = handlers.NewRepo(sm)
)

// Middlewares struct
type Middlewares struct {
	sm *models.SessionManager
}

// NewMiddlewares creates a new Middlewares instance
func NewMiddlewares(sm *models.SessionManager) *Middlewares {
	return &Middlewares{sm: sm}
}

// Allowed routes
var allowedRoutes = map[string]bool{
	"/":            true,
	"/ascii-art":   true,
	"/about":       true,
	"/contact":     true,
	"/login":       true,
	"/logout":      true,
	"/msg-success": true,
}

// RouteChecker is a middleware that checkes allowed routes
func RouteChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/static/") || strings.HasPrefix(r.URL.Path, "/api/download-ascii"){
			next.ServeHTTP(w, r)
			return
		}

		if _, ok := allowedRoutes[r.URL.Path]; !ok {
			repo.NotFoundHandler(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
