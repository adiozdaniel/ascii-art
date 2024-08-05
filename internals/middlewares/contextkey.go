package middlewares

type contexKey string

const (
	// SessionKey is the key for storing the session in the context
	SessionKey contexKey = "session"
)
