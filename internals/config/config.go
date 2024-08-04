package config

import "html/template"

// AppConfig holds configuration for the application.
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}

// App is the instance of AppConfig.
func App() *AppConfig {
	return &AppConfig{
		TemplateCache: make(map[string]*template.Template),
	}
}
