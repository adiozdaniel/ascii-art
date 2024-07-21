package renders

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
)

type FormData struct {
	Body          string
	TemplateCache map[string]*template.Template
}

// Data is a global variable to hold the form data
var Data FormData

// functions is a map of template functions
var functions = template.FuncMap{}

// RenderTemplate is a helper function to render HTML templates
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tc, err := getTemplateCache()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, ok := tc[tmpl]
	if !ok {
		http.Error(w, "😏Oops, something went wrong", http.StatusNotFound)
		return
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, data)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getTemplateCache is a helper function to cache all HTML templates as a map
func getTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./views/templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./views/templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./views/templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
