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
			utils.ErrorHandler("web")
		}
	}

	fileContents, _ := ascii.FileContents()
	output := ascii.Output(fileContents)
	nonAsciis := utils.NonAsciiOutput()

	if utils.Inputs.Output != "" {
		if utils.Inputs.Justify != "" {
			fmt.Printf("🙄 alignment request: 'align=%s'; was ignored\n=================================================\n\n", utils.Inputs.Justify)
		}
		utils.LogOutput(output)
		fmt.Printf("😋 writing.... '%s'. To check output, kindly use: `cat %s | cat -e`\n====================================\n %s", utils.Inputs.Input, utils.Inputs.Output, nonAsciis)
		return
	}

	if utils.Inputs.Justify != "" {
		utils.Alignment(output)
		return
	}

	fmt.Print(output, nonAsciis)
}
