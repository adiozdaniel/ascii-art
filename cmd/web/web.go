package main

import (
	"fmt"
	"net/http"

	"github.com/adiozdaniel/ascii-art/internals/ascii"
	"github.com/adiozdaniel/ascii-art/internals/routes"
	"github.com/adiozdaniel/ascii-art/pkg/utils"
)

// global variables
var (
	app = utils.NewInput()
)

func main() {
	app.Init()
	app.Flags["isWeb"] = "true"
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	wrappedMux := routes.RouteChecker(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: wrappedMux,
	}

	app.Flags["font"] = "--standard"
	app.Flags["input"] = "Ascii~"
	serverOutput := ascii.Output(app.Flags["input"])
	fmt.Println(serverOutput + "=====================================\nserver running @http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		app.ErrorHandler("web")
	}
}
