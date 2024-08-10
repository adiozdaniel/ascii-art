package models

// Support represents a contact form submission
type Support struct {
	Name    string
	Email   string
	Message string
}

// NewSupport creates a new Support instance
func NewSupport() *Support {
	return &Support{}
}
