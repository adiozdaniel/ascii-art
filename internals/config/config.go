package config

import "html/template"

// AppConfig holds configuration for the application.
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
