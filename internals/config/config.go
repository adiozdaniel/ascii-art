package config

import "html/template"

// WebConfig holds configuration for the web application.
type WebConfig struct {
	TemplateCache map[string]*template.Template
}
