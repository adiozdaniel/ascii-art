package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adiozdaniel/ascii-art/ascii"
	"github.com/adiozdaniel/ascii-art/routes"
	"github.com/adiozdaniel/ascii-art/utils"
)

func main() {
	if os.Args[1] == "-web" {
		mux := http.NewServeMux()

		// Use the routes package to set up your routes
		routes.RegisterRoutes(mux)

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
