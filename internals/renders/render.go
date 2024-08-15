package renders

import (
	"html/template"
	"net/http"

	"github.com/adiozdaniel/ascii-art/internals/models"
)

// RenderTemplate is a helper function to render HTML templateseWriter, tmpl string, data
func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	t := models.GetStateManager().GetConfig().TemplateCache
	ts, ok := t[tmpl]
	if !ok {
		renderServerErrorTemplate(w, tmpl+" is missing, contact the Network Admin.")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := ts.Execute(w, data)
	if err != nil {
		return
	}
}

// renderServerErrorTemplate renders a simple error template directly
func renderServerErrorTemplate(w http.ResponseWriter, errMsg string) {
	t, err := template.New("error").Parse(Tmpl)
	if err != nil {
		http.Error(w, "Oops, something went wrong", http.StatusInternalServerError)
	}

	data := struct {
		Error string
	}{
		Error: errMsg,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	t.Execute(w, data)
}
