package renders

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Define custom template functions
var functions = template.FuncMap{
	"contains": func(s, substr string) bool {
        return strings.Contains(s, substr)
    },
    "toUpper": strings.ToUpper,
}

// Cache to store templates for reuse
var (
	templateCache = make(map[string]*template.Template)
	cacheMutex    sync.RWMutex
)

// RenderTemplate is a helper function to render HTML templates
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Use the cached template if available
	ts, ok := getCachedTemplate(tmpl)
	if !ok {
		renderServerErrorTemplate(w, tmpl+" is missing, contact the Network Admin.")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := ts.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// getCachedTemplate fetches the template from cache or loads and caches it if not present
func getCachedTemplate(name string) (*template.Template, bool) {
	cacheMutex.RLock()
	ts, ok := templateCache[name]
	cacheMutex.RUnlock()

	if ok {
		return ts, true
	}

	// Load and cache the template if not found in cache
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	// Double-check if another goroutine already cached the template
	if ts, ok := templateCache[name]; ok {
		return ts, true
	}

	tmpl, err := loadTemplate(name)
	if err != nil {
		return nil, false
	}
	templateCache[name] = tmpl
	return tmpl, true
}

// loadTemplate loads and parses a template
func loadTemplate(name string) (*template.Template, error) {
	baseDir := GetProjectRoot("views", "templates")

	// Find and parse the template and its layout files
	pagePath := filepath.Join(baseDir, name+".page.html")
	tmpl, err := template.New(filepath.Base(pagePath)).Funcs(functions).ParseFiles(pagePath)
	if err != nil {
		return nil, fmt.Errorf("error parsing page %s: %v", name, err)
	}

	layoutsPath := filepath.Join(baseDir, "*.layout.html")
	if _, err := tmpl.ParseGlob(layoutsPath); err != nil {
		return nil, fmt.Errorf("error parsing layout files: %v", err)
	}

	return tmpl, nil
}

// GetProjectRoot dynamically finds the project root directory
func GetProjectRoot(first, second string) string {
	cwd, _ := os.Getwd()
	baseDir := cwd
	if strings.HasSuffix(baseDir, "cmd") {
		baseDir = filepath.Join(cwd, "../")
	}
	return filepath.Join(baseDir, first, second)
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
		http.Error(w, "Not Found", http.StatusNotFound)
	}

	data := struct {
		Error string
	}{
		Error: errMsg,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	t.Execute(w, data)
}
