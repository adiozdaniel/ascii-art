package models

import (
	"net/http"
	"net/url"
)

// errors represents a collection of error messages for various fields
type formErrors map[string][]string

// Add adds an error message to the specified field
func (e *formErrors) Add(field, message string) {
	(*e)[field] = append((*e)[field], message)
}

// Get returns the first error message for a given field, or an empty string if no errors exist
func (e *formErrors) Get(field string) string {
	es := (*e)[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}

// Forms represents a collection of form data and errors
type Forms struct {
	url.Values
	Errors formErrors
}

// NewForms creates a new Forms instance with the given data and initializes an empty error map
func NewForms(data url.Values) *Forms {
	return &Forms{data, make(formErrors)}
}

// Has checks if a field has been submitted and returns true if it has, false otherwise
func (f *Forms) Has(field string, r *http.Request) bool {
	return r.Form.Get(field) != ""
}

// ValidateForm validates the form data and adds errors to the error map if necessary
func (f *Forms) ValidateForm() bool {
	return len(f.Errors) == 0
}
