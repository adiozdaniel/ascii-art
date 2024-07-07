package main

import (
	"fmt"
	"net/http"

	"github.com/adiozdaniel/ascii-art/routes"
	"github.com/adiozdaniel/ascii-art/utils"
)

// runOutput writes the output to the specified file.
func runOutput(output, nonAsciis string) {
	if utils.Inputs.Justify != "" {
		fmt.Printf("🙄 alignment request: 'align=%s'; was ignored\n=================================================\n\n", utils.Inputs.Justify)
	}
	utils.LogOutput(output)
	fmt.Printf("😋 writing.... '%s'. To check output, kindly use: `cat %s | cat -e`\n====================================\n %s", utils.Inputs.Input, utils.Inputs.Output, nonAsciis)
}

// runWeb starts the web server to handle HTTP requests.
func runWeb() {
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
