package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/adiozdaniel/ascii-art/ascii"
	"github.com/adiozdaniel/ascii-art/utils"
)

// html files passed as templates
var filePrefix, _ = (filepath.Abs("templates/"))
var tmpl2 = template.Must(template.ParseFiles(utils.CleanPath(filePrefix) + "/index.page.html"))
var tmplNotFound = template.Must(template.ParseFiles(utils.CleanPath(filePrefix) + "/notfound.page.html"))
var tmplBadRequest = template.Must(template.ParseFiles(utils.CleanPath(filePrefix) + "/badrequest.page.html"))
var tmplInternalError = template.Must(template.ParseFiles(utils.CleanPath(filePrefix) + "/serverError.page.html"))

type FormData struct {
	Body string
}

var data FormData

// renderTemplate is a helper function to render HTML templates
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	parsedTemplate, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HomeHandler handles the homepage route '/'
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, utils.CleanPath(filePrefix)+"/index.page.html", nil)
}

// SubmitHandler handles the output route '/ascii-art'
func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "🧐 Can I treat this as an invalid request?", http.StatusMethodNotAllowed)
		return
	}

	_, err := ascii.FileContents()
	if err != nil {
		ServerError(w, r)
		return
	}

	output := ascii.Output(utils.Inputs.Input)
	data.Body = output

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
