package models

import (
	"net/url"
	"strings"
)

// TemplateData represents data send to templates
type TemplateData struct {
	StringMap map[string]string `json:"stringMap"`
	Form      *Forms
	EmailData *EmailData
}

// NewTemplateData initializes a new TemplateData
func NewTemplateData() *TemplateData {
	return &TemplateData{
		StringMap: make(map[string]string),
		Form:      NewForms(make(url.Values)),
		EmailData: NewEmailData("", "", "", ""),
	}
}

// CapitalizeFirst capitalizes the first character of a string and lowercases the rest.
func (s *TemplateData) CapitalizeFirst(input string) string {
	if len(input) == 0 {
		return input
	}

	if len(input) > 12 {
		input = strings.ToLower(input[:12])
		return strings.ToUpper(input[:1]) + input[1:]
	}

	return strings.ToUpper(string(input[0])) + strings.ToLower(input[1:])
}
