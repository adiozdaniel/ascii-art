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

	if utils.Inputs.Output != "" {
		fmt.Printf("😋 writing.... '%s'. To check output, kindly use: `cat %s | cat -e`\n", utils.Inputs.Input, utils.Inputs.Output)
		utils.LogOutput(output)
		return
	}

	fmt.Print(output, nonAsciis)
}
