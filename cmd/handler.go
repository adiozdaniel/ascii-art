package main

import (
	"html/template"
	"net/http"

	ascii "github.com/adiozdaniel/ascii-art/ascii"
	utils "github.com/adiozdaniel/ascii-art/utils"
)

var tmpl2 = template.Must(template.ParseFiles("index.page.tmpl"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl2.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
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

	// fmt.Printf("You typed: %s\nYou selected: %s", textInput, fileName)
	data := struct {
		Body string
	}{
		Body: result,
	}

	tmpl2.Execute(w, data)
}
