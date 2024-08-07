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

	if len(input) > 12 {
		input = strings.ToLower(input[:12])
		return strings.ToUpper(input[:1]) + input[1:]
	}

	return strings.ToUpper(string(input[0])) + strings.ToLower(input[1:])
}
