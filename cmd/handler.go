package main

import (
	"net/http"
	"strings"

	ascii "github.com/adiozdaniel/ascii-art/ascii"
	utils "github.com/adiozdaniel/ascii-art/utils"
)

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

	textInput := r.FormValue("textInput")
	utils.Inputs. := r.FormValue("FileName")

	fileContents := ascii.FileContents()
	input := utils.Inputs.Input
	output := ascii.Output(fileContents)
	nonAsciis := utils.NonAsciiOutput(strings.Split(input, "\n"))

	fileContents := ascii.FileContents()
	input := strings.ReplaceAll(textInput, "\\n", "\n")
	output := ascii.Output(strings.Split(input, "\n"), fileContents)
	nonAsciis := utils.NonAsciiOutput(strings.Split(input, "\n"))
	result := output + "\n\n" + nonAsciis

	// fmt.Printf("You typed: %s\nYou selected: %s", textInput, fileName)
	data := struct {
		Body string
	}{
		Body: result,
	}

	tmpl2.Execute(w, data)
}
