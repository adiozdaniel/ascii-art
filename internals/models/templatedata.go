package models

import "strings"

type TemplateData struct {
	StringMap map[string]string `json:"stringMap"`
}

func NewTemplateData() *TemplateData {
	return &TemplateData{
		StringMap: make(map[string]string),
	}
}

// CapitalizeFirst capitalizes the first character of a string and lowercases the rest.
func (s *TemplateData) CapitalizeFirst(input string) string {
	if len(input) == 0 {
		return input
	}

	return strings.ToUpper(string(input[0])) + strings.ToLower(input[1:])
}
