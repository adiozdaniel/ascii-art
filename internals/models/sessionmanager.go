package models

// Session represents a user session
type Session struct {
	CRSFToken string
}

// SessionManager manages user sessions
type SessionManager struct {
	sessions map[string]*Session
}

// NewSessionManager returns a new instance of SessionManager
func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*Session),
	}
}
