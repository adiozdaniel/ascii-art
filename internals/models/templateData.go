package models

// TemplateData holds the data for the HTML templates
type TemplateData struct {
	Ascii     map[string]string
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
