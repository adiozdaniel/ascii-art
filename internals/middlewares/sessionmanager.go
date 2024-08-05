package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/adiozdaniel/ascii-art/internals/models"
)

const sessionCookieName = "session_id"

// SessionMiddleware sets up sessions for incoming requests
func SessionMiddleware(sm *models.SessionManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(sessionCookieName)
			var sessionID string

			if err == nil {
				sessionID = cookie.Value
			} else {
				session := sm.CreateSession()
				sessionID = session.CRSFToken
				http.SetCookie(w, &http.Cookie{
					Name:     sessionCookieName,
					Value:    sessionID,
					Path:     "/",
					Expires:  time.Now().Add(30 * time.Minute),
					HttpOnly: true,
					Secure:   false, // TODO replace with true for production
				})
			}

			session, _ := sm.GetSession(sessionID)
			if session != nil {
				ctx := context.WithValue(r.Context(), sm.SessionKey, session)
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		})
	}
}
