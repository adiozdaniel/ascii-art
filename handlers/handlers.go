package handlers

import (
	"html/template"
	"net/http"

	ascii "github.com/adiozdaniel/ascii-art/ascii"
)

// html files passed as templates
var tmpl2 = template.Must(template.ParseFiles("../templates/index.page.tmpl"))
var tmplNotFound = template.Must(template.ParseFiles("../templates/notfound.page.tmpl"))
var tmplBadRequest = template.Must(template.ParseFiles("../templates/badrequest.page.tmpl"))
var tmplInternalError = template.Must(template.ParseFiles("../templates/serverError.page.tmpl"))

// HomeHandler handles the homepage route '/'
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl2.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SubmitHandler handles the output route '/ascii-art'
func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "🧐 Can I treat this as an invalid request?", http.StatusMethodNotAllowed)
		return
	}

	fileContents, err := ascii.FileContents()
	if err != nil {
		ServerError(w, r)
		return
	}

	output := ascii.Output(fileContents)
	data := struct {
		Body string
	}{
		Body: output,
	}

	tmpl2.Execute(w, data)
}

// NotFoundHandler handles unknown routes; 404 status
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := tmplNotFound.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

// BadRequestHandler handles the bad requests routes
func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	err := tmplBadRequest.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// ServerError handles server failures that result to status 500
func ServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	err := tmplInternalError.Execute(w, nil)
	if err != nil {
		http.Error(w, "🧐 Internal Server Error", http.StatusInternalServerError)
	}
}
