package handlers

import (
	"github.com/adiozdaniel/ascii-art/ascii"
	"github.com/adiozdaniel/ascii-art/internals/renders"
	"github.com/adiozdaniel/ascii-art/utils"
	"net/http"
)

var data renders.FormData

// HomeHandler handles the homepage route '/'
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.html", nil)
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
	data.Body = output

	renders.RenderTemplate(w, "home.page.html", data)
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
