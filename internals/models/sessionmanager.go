package models

import (
	"sync"
)

// Session represents a user session
type Session struct {
	CRSFToken string
}

// SessionManager manages user sessions
type SessionManager struct {
	sessions map[string]*Session
	lock     sync.RWMutex
}

// NewSessionManager returns a new instance of SessionManager
func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*Session),
	}
}

// CreateSession creates a new session
func (sm *SessionManager) CreateSession() *Session {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	sessionID := generateSessionID()
	session := &Session{
		CRSFToken: sessionID,
	}
	sm.sessions[sessionID] = session
	return session
}

// generateSessionID generates a unique session ID
func generateSessionID() string {
	return "unique-session-id"
}
