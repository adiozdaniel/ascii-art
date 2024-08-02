package main

import (
	"fmt"
	"net/http"

	"github.com/adiozdaniel/ascii-art/internals/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	wrappedMux := routes.RouteChecker(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: wrappedMux,
	}

	utils.Inputs.BannerPath = "../banners/standard.txt"
	utils.Inputs.Input = "Ascii~"
	serverOutput := ascii.Output(utils.Inputs.Input)
	fmt.Println(serverOutput + "=====================================\nserver running @http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		utils.ErrorHandler("web")
	}
}
