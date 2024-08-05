package models

import (
	"sync"
	"time"
)

// Session represents a user session
type Session struct {
	CRSFToken string
	Expiry    time.Time
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
		Expiry:    time.Now().Add(30 * time.Minute),
	}
	sm.sessions[sessionID] = session
	return session
}

// generateSessionID generates a unique session ID
func generateSessionID() string {
	return "unique-session-id"
}
