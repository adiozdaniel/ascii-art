package handlers

import (
	// "fmt"

	"html/template"
	"net/http"

	"github.com/adiozdaniel/ascii-art/ascii"
	"github.com/adiozdaniel/ascii-art/internals/renders"
	"github.com/adiozdaniel/ascii-art/utils"
)

// var tmpl2 = template.Must(template.ParseFiles(utils.CleanPath(filePrefix) + "/index.page.html"))
var tmplNotFound = template.Must(template.ParseFiles(utils.GetFilePath("views/templates", "notfound.page.html")))
var tmplBadRequest = template.Must(template.ParseFiles(utils.GetFilePath("views/templates", "badrequest.page.html")))
var tmplInternalError = template.Must(template.ParseFiles(utils.GetFilePath("views/templates", "serverError.page.html")))

// HomeHandler handles the homepage route '/'
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "index.page.html", nil)
}

// SubmitHandler handles the output route '/ascii-art'
func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || r.FormValue("textInput") == "" {
		BadRequestHandler(w, r)
		return
	}

	utils.Inputs.Input = r.FormValue("textInput")
	banner := utils.BannerFiles[r.FormValue("FileName")]

	if banner == "" {
		utils.Inputs.BannerPath = utils.BannerFiles["standard"]
	} else {
		utils.Inputs.BannerPath = banner
	}

	output := ascii.Output(utils.Inputs.Input)
	renders.Data.Body = output

	renders.RenderTemplate(w, "index.page.html", renders.Data.Body)
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
