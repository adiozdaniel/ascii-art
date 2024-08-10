package models

import "net/url"

type Forms struct {
	url.Values
	Errors formErrors
}

func NewForms(data url.Values) *Forms {
	return &Forms{data, make(formErrors)}
}
