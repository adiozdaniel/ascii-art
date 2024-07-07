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
var tmpl2 = template.Must(template.ParseFiles(utils.CleanPath(filePrefix) + "/index.page.tmpl"))
var tmplNotFound = template.Must(template.ParseFiles(utils.CleanPath(filePrefix) + "/notfound.page.tmpl"))
var tmplBadRequest = template.Must(template.ParseFiles(utils.CleanPath(filePrefix) + "/badrequest.page.tmpl"))
var tmplInternalError = template.Must(template.ParseFiles(utils.CleanPath(filePrefix) + "/serverError.page.tmpl"))

type FormData struct {
	Body string
}

var data FormData

// HomeHandler handles the homepage route '/'
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.Inputs.Input = ""
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
