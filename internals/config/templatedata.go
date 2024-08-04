package config

type TemplateData struct {
	StringMap map[string]string
}

func NewTemplateData() *TemplateData {
	return &TemplateData{
		StringMap: make(map[string]string),
	}
}
