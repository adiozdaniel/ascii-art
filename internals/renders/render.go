package renders

import (
	"html/template"
	"net/http"

	"github.com/adiozdaniel/ascii-art/internals/config"
)

// get the app state manager
var (
	sm        = config.GetStateManager()
	appData   = sm.GetInput()
	appConfig = sm.GetConfig()
)

type FormData struct {
	Body string
}

// global variables
var (
	Data FormData
)

// functions is a map of template functions
var functions = template.FuncMap{}

// RenderTemplate is a helper function to render HTML templates
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t := appConfig.TemplateCache
	ts, ok := t[tmpl]
	if !ok {
		renderServerErrorTemplate(w, tmpl+" is missing, contact the Network Admin.")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := ts.Execute(w, data)
	if err != nil {
		return
	}
}

// renderServerErrorTemplate renders a simple error template directly
func renderServerErrorTemplate(w http.ResponseWriter, errMsg string) {
	tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Server Error</title>
	<style>
		body {color: bisque; background-color: #333; font-family: Arial, sans-serif; }
		.container { text-align: center; margin-top: 50px; }
		.btn { display: inline-block; margin-top: 20px; padding: 10px 20px; background-color: #007bff; color: white; text-decoration: none; border-radius: 5px; }
	</style>
</head>
<body>
	<div class="container">
		<h1>404 Oops We can't find what you are looking for! 🙁</h1>
		<h2>Something went wrong.</h2>
		<h3>{{.Error}}</h3>
		<a href="/" title="Go back to the home page" class="btn">
			<h1>Home</h1>
		</a>
	</div>
</body>
</html>`

	t, err := template.New("error").Parse(tmpl)
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
