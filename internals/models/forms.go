package models

import (
	"net/http"
	"net/url"
	"strings"
)

// errors represents a collection of error messages for various fields
type formErrors map[string][]string

// Add adds an error message to the specified field
func (e *formErrors) Add(field, message string) {
	(*e)[field] = append((*e)[field], message)
}

// Clear clears all error messages for the specified field
func (e *formErrors) Clear() {
	*e = make(map[string][]string)
}

// Get returns the first error message for a given field, or an empty string if no errors exist
func (e *formErrors) Get(field string) string {
	es := (*e)[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}

// formValues represents a collection of form values submitted by the user
type formValues map[string]string

// Set sets the value of a form field
func (v *Forms) Set(field, value string) {
	v.FormValues[field] = value
}

// Forms represents a collection of form data and errors
type Forms struct {
	url.Values
	Errors     formErrors
	FormValues formValues
}

// NewForms creates a new Forms instance with the given data and initializes an empty error map
func NewForms(data url.Values) *Forms {
	return &Forms{data, make(formErrors), make(formValues)}
}

// Has checks if a field has been submitted and returns true if it has, false otherwise
func (f *Forms) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field is required")
		return false
	}
	return r.Form.Get(field) != ""
}

func (f *Forms) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field is required")
		}
	}
}

// ValidateForm returns true if all fields in the form have been submitted and have no errors, false otherwise
func (f *Forms) IsValidForm() bool {
	return len(f.Errors) == 0
}
