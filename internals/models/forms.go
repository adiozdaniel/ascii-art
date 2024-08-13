package models

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
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
func (f *Forms) Set(field, value string) {
	f.FormValues[field] = value
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

// Required checks if a submitted field is empty and prints an error message
func (f *Forms) Required(r *http.Request, fields ...string) {
	for _, field := range fields {
		f.Set(field, r.Form.Get(field))

		value := f.FormValues[field]
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, field+" cannot be empty")
		}
		f.MinLength(field)

		if field == "email" && !IsValidEmail(value) {
			f.Errors.Add(field, "Invalid email address")
		}
	}
}

// fieldLengths represents the minimum length required for each form field
var fieldLengths = map[string]int{
	"name":     3,
	"email":    9,
	"message":  30,
	"username": 3,
}

// MinLength checks if a submitted field has a minimum length and prints an error message if it does not
func (f *Forms) MinLength(field string) {
	value := f.FormValues[field]
	if len(value) < fieldLengths[field] {
		f.Errors.Add(field,
			fmt.Sprintf("%s must be at least %d characters long",
				field, fieldLengths[field]))
	}
}

// IsValidEmail validates an email address using a regular expression.
func IsValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// ValidateForm returns true if all fields in the form have been submitted and have no errors, false otherwise
func (f *Forms) IsValidForm() bool {
	return len(f.Errors) == 0
}
