package renders

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type FormData struct {
	Body string
}

// Data is a global variable to hold the form data
var Data FormData

// functions is a map of template functions
var functions = template.FuncMap{}

// RenderTemplate is a helper function to render HTML templates
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	parsedTemplate, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = parsedTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RenderTemplateString is a helper function to render HTML templates as strings
func RenderTemplateString(w http.ResponseWriter, tmpl string) error {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./views/templates/*.page.html")
	if err != nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return err
		}

		matches, err := filepath.Glob("./views/templates/*.layout.html")
		if err != nil {
			return err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./views/templates/*.layout.html")
			if err != nil {
				return err
			}
		}

		myCache[name] = ts
	}
	return nil
}
