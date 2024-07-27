package handlers

import (
	"net/http"

	"github.com/adiozdaniel/ascii-art/ascii"
	"github.com/adiozdaniel/ascii-art/internals/renders"
	"github.com/adiozdaniel/ascii-art/utils"
)

var data renders.FormData

// HomeHandler handles the homepage route '/'
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.html", nil)
}

// SubmitHandler handles the output route '/ascii-art'
func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("textInput") == "" && r.Method != "POST" {
		renders.RenderTemplate(w, "ascii.page.html", nil)
		return
	}

	if r.Method == "POST" && r.FormValue(("textInput")) == "" {
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

	_, err := ascii.FileContents()
	if err != nil {
		ServerErrorHandler(w, r)
		return
	}

	output := ascii.Output(utils.Inputs.Input)
	noasciis := utils.NonAsciiOutput()
	data.Body = output + "\n" + noasciis

	renders.RenderTemplate(w, "ascii.page.html", data)
}

// NotFoundHandler handles unknown routes; 404 status
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renders.RenderTemplate(w, "notfound.page.html", nil)
}

// BadRequestHandler handles bad requests routes
func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	renders.RenderTemplate(w, "badrequest.page.html", nil)
}

// ServerErrorHandler handles server failures that result in status 500
func ServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	renders.RenderTemplate(w, "serverError.page.html", nil)
}

// AboutHandler handles the about page route '/about'
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.page.html", nil)
}

// ContactHandler handles the contact page route '/contact'
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "contact.page.html", nil)
}
