package config

type TemplateData struct {
	Body      string            `json:"Body"`
	StringMap map[string]string `json:"stringMap"`
}

func NewTemplateData() *TemplateData {
	return &TemplateData{
		Body:      "",
		StringMap: make(map[string]string),
	}
}
