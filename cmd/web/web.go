package main

import (
	"fmt"
	"net/http"

	appconfig "github.com/adiozdaniel/ascii-art/internals/app_config"
	"github.com/adiozdaniel/ascii-art/internals/ascii"
	"github.com/adiozdaniel/ascii-art/internals/routes"
	"github.com/adiozdaniel/ascii-art/pkg/helpers"
)

// get the app state manager
var (
	sm  = appconfig.GetStateManager()
	app = sm.GetInput()
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

	banner := app.BannerFile[app.Flags["font"]]
	err := helpers.FileContents(banner)
	if err != nil {
		app.ErrorHandler("fatal")
	}

	serverOutput := ascii.Output(app.Flags["input"])
	fmt.Println(serverOutput + "=====================================\nserver running @http://localhost:8080")
	err = server.ListenAndServe()
	if err != nil {
		app.ErrorHandler("web")
	}
}
