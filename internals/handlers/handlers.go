package handlers

import (
	"net/http"

	"github.com/adiozdaniel/ascii-art/ascii"
	"github.com/adiozdaniel/ascii-art/internals/config"
	"github.com/adiozdaniel/ascii-art/internals/renders"
	"github.com/adiozdaniel/ascii-art/utils"
)

// data is the variable that holds the form data
var data renders.FormData

// Repository is a struct to hold the application configuration
type Repository struct {
	App *config.AppConfig
}

// Repo is a global variable to hold the Repository instance
var Repo *Repository

// NewRepo initializes a new Repository instance
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers initializes the repository instance and sets the handlers
func NewHandlers(r *Repository) {}

// HomeHandler handles the homepage route '/'
func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.html", nil)
}

// SubmitHandler handles the output route '/ascii-art'
func (m *Repository) SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("textInput") == "" && r.Method != "POST" {
		renders.RenderTemplate(w, "ascii.page.html", nil)
		return
	}

	if r.Method == "POST" && r.FormValue(("textInput")) == "" {
		Repo.BadRequestHandler(w, r)
		return
	}

	utils.Inputs.Input = r.FormValue("textInput")
	banner := utils.BannerFiles[r.FormValue("FileName")]

	if banner == "" {
		utils.Inputs.BannerPath = utils.BannerFiles["standard"]
	} else {
		utils.Inputs.BannerPath = banner
	}

	_, err := ascii.FileContents()
	if err != nil {
		Repo.ServerErrorHandler(w, r)
		return
	}

	output := ascii.Output(utils.Inputs.Input)
	noasciis := utils.NonAsciiOutput()
	data.Body = output + "\n" + noasciis

	renders.RenderTemplate(w, "ascii.page.html", data)
}

// NotFoundHandler handles unknown routes; 404 status
func (m *Repository) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renders.RenderTemplate(w, "notfound.page.html", nil)
}

// BadRequestHandler handles bad requests routes
func (m *Repository) BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	renders.RenderTemplate(w, "badrequest.page.html", nil)
}

// ServerErrorHandler handles server failures that result in status 500
func (m *Repository) ServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	renders.RenderTemplate(w, "serverError.page.html", nil)
}

// AboutHandler handles the about page route '/about'
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.page.html", nil)
}

// ContactHandler handles the contact page route '/contact'
func (m *Repository) ContactHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "contact.page.html", nil)
}
