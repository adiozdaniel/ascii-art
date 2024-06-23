package handlers

import (
	ascii "github.com/adiozdaniel/ascii-art/ascii"
	utils "github.com/adiozdaniel/ascii-art/utils"
	"html/template"
	"net/http"
)

var tmpl2 = template.Must(template.ParseFiles("../templates/index.page.tmpl"))
var tmplNotFound = template.Must(template.ParseFiles("../templates/notfound.page.tmpl"))
var tmplBadRequest = template.Must(template.ParseFiles("../templates/badrequest.page.tmpl"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl2.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "🧐 Can I treat this as an invalid request?", http.StatusMethodNotAllowed)
		return
	}

	utils.Inputs.Input = r.FormValue("textInput")
	banner := utils.BannerFiles[r.FormValue("FileName")]
	if banner == "" {
		utils.Inputs.Banner = utils.BannerFiles["standard"]
	} else {
		utils.Inputs.Banner = utils.BannerFiles[r.FormValue("FileName")]
	}

	fileContents := ascii.FileContents()
	output := ascii.Output(fileContents)
	nonAsciis := utils.NonAsciiOutput()
	result := output + "\n\n" + nonAsciis

	data := struct {
		Body string
	}{
		Body: result,
	}

	tmpl2.Execute(w, data)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := tmplNotFound.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	err := tmplBadRequest.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
}
