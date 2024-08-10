package models

import "net/url"

// Forms represents a collection of form data and errors
type Forms struct {
	url.Values
	Errors formErrors
}

// NewForms creates a new Forms instance with the given data and initializes an empty error map
func NewForms(data url.Values) *Forms {
	return &Forms{data, make(formErrors)}
}
