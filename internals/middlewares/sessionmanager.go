package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/adiozdaniel/ascii-art/internals/models"
)

// SessionMiddleware sets up sessions for incoming requests
func SessionMiddleware(sm *models.SessionManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_id")
			var sessionID string

			if err == nil {
				sessionID = cookie.Value
			} else {
				session := sm.CreateSession()
				sessionID = session.CRSFToken
				http.SetCookie(w, &http.Cookie{
					Name:    "session_id",
					Value:   sessionID,
					Path:    "/",
					Expires: time.Now().Add(30 * time.Minute),
				})
			}

			session, _ := sm.GetSession(sessionID)
			if session != nil {
				ctx := context.WithValue(r.Context(), SessionKey, session)
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		})
	}
}
