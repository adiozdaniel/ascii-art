package main

import (
	"fmt"
	"net/http"
	"os"

	ascii "github.com/adiozdaniel/ascii-art/ascii"
	utils "github.com/adiozdaniel/ascii-art/utils"
)

func main() {
	if os.Args[1] == "-web" {
		mux := http.NewServeMux()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				homeHandler(w, r)
			} else if r.URL.Path == "/submit-form" {
				submitHandler(w, r)
			} else {
				notFoundHandler(w, r)
			}
		})

		mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
		mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "../static/favicon.ico")
		})

		server := &http.Server{
			Addr:    ":8080",
			Handler: mux,
		}

		fmt.Println("Server is running on http://localhost:8080")
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}
	fileContents := ascii.FileContents()
	output := ascii.Output(fileContents)
	nonAsciis := utils.NonAsciiOutput()

	fmt.Print(output, nonAsciis)
}

// func main() {
// 	http.HandleFunc("/", homeHandler)
// 	http.HandleFunc("/submit-form", submitHandler)

// 	fmt.Println("Server is running on http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }

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
