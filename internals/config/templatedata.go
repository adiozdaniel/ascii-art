package config

type TemplateData struct {
	StringMap map[string]string `json:"stringMap"`
}

func NewTemplateData() *TemplateData {
	return &TemplateData{
		StringMap: make(map[string]string),
	}
}
