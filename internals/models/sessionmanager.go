package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
	"time"
)

// contexKey is context type
type contexKey string

const (
	// SessionKey is the key for storing the session in the context
	sessionKey contexKey = "session"
)

// Session represents a user session
type Session struct {
	CRSFToken string
	Expiry    time.Time
	Data      *TemplateData
}

// SessionManager manages user sessions
type SessionManager struct {
	sessions   map[string]*Session
	SessionKey contexKey
	lock       sync.RWMutex
}

// NewSessionManager returns a new instance of SessionManager
func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions:   make(map[string]*Session),
		SessionKey: sessionKey,
	}
}

// CreateSession creates a new session
func (sm *SessionManager) CreateSession() *Session {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	sessionID := sm.GenerateSessionID()
	session := &Session{
		CRSFToken: sessionID,
		Expiry:    time.Now().Add(30 * time.Minute),
		Data:      NewTemplateData(),
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

// GetSessionData retrieves session data by ID
func (sm *SessionManager) GetSessionData(sessionID string) *TemplateData {
	session, exists := sm.GetSession(sessionID)
	if exists {
		return session.Data
	}
	return nil
}

// DeleteSession removes a session
func (sm *SessionManager) DeleteSession(sessionID string) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	delete(sm.sessions, sessionID)
}

// generateSessionID generates a unique session ID
func (sm *SessionManager) GenerateSessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return base64.URLEncoding.EncodeToString(b)
}
