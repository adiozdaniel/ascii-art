package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tmpl2 = template.Must(template.ParseFiles("index.html"))

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/submit-form", submitHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// package main

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/adiozdaniel/ascii-art/ascii"
// 	"github.com/adiozdaniel/ascii-art/utils"
// )

// func main() {
// 	fileContents := ascii.FileContents()
// 	input := utils.Inputs.Input
// 	output := ascii.Output(fileContents)
// 	nonAsciis := utils.NonAsciiOutput(strings.Split(input, "\n"))

// 	fmt.Print(output, nonAsciis)
// }
