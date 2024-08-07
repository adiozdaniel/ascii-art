package middlewares

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/adiozdaniel/ascii-art/internals/models"
)

const sessionCookieName = "session_id"

// SessionMiddleware sets up sessions for incoming requests
func SessionMiddleware(sm *models.SessionManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/static/") {
				next.ServeHTTP(w, r)
				return
			}

			if r.URL.Path == "/contact"{
				next.ServeHTTP(w, r)
				return
			}

			cookie, err := r.Cookie(sessionCookieName)
			var sessionID string

			if err == nil {
				sessionID = cookie.Value
			} else {
				if r.URL.Path != "/login" {
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}
				sessionID = ""
			}

			session, exists := sm.GetSession(sessionID)
			if !exists || session.Expiry.Before(time.Now()) {
				http.SetCookie(w, &http.Cookie{
					Name:     sessionCookieName,
					Value:    "",
					Path:     "/",
					Expires:  time.Now().Add(-time.Hour),
					HttpOnly: true,
					Secure:   false, // TODO replace with true for production
				})

				if r.URL.Path != "/login" {
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}
			}

			if session != nil {
				ctx := context.WithValue(r.Context(), sm.SessionKey, session)
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		})
	}
}
