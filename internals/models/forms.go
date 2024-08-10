package models

// errors represents a collection of error messages for various fields
type errors map[string][]string

// Add adds an error message to the specified field
func (e *errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error message for a given field, or an empty string if no errors exist
func (e *errors) Get(field string) string {
	es := (*e)[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
