package config

import "html/template"

// AppConfig holds the configuration for the web application
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
