package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
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

// GetSession retrieves a session by ID
func (sm *SessionManager) GetSession(sessionID string) (*Session, bool) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()

	session, exists := sm.sessions[sessionID]
	if !exists || session.Expiry.Before(time.Now()) {
		return nil, false
	}
	return session, true
}

// DeleteSession removes a session
func (sm *SessionManager) DeleteSession(sessionID string) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	delete(sm.sessions, sessionID)
}

// generateSessionID generates a unique session ID
func generateSessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return base64.URLEncoding.EncodeToString(b)
}
