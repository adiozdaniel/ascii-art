package main

import (
	"fmt"
	"net/http"

	"github.com/adiozdaniel/ascii-art/internals/ascii"
	"github.com/adiozdaniel/ascii-art/internals/config"
	"github.com/adiozdaniel/ascii-art/internals/renders"
	"github.com/adiozdaniel/ascii-art/internals/routes"
	"github.com/adiozdaniel/ascii-art/pkg/helpers"
)

// get the app state manager
var (
	sm        = config.GetStateManager()
	appData   = sm.GetInput()
	appConfig = sm.GetConfig()
)

func main() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	wrappedMux := routes.RouteChecker(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: wrappedMux,
	}

	banner := appData.BannerFile[appData.Flags["font"]]
	err := helpers.FileContents(banner)
	if err != nil {
		appData.ErrorHandler("fatal")
	}

	serverOutput := ascii.Output(appData.Flags["input"])
	fmt.Println(serverOutput + "=====================================\nserver running @http://localhost:8080")

	appData.Flags["isWeb"] = "true"
	err = server.ListenAndServe()
	if err != nil {
		appData.ErrorHandler("web")
	}
}

// init initializes the web data
func init() {
	appData.Init()

	tc, err := renders.CreateTemplateCache()
	if err != nil {
		appData.ErrorHandler("templates")
	}

	appConfig.TemplateCache = tc

	bc, err := appConfig.CreateBannerCache()
	if err != nil {
		appData.ErrorHandler("banners")
	}

	appData.Flags["font"] = "--standard"
	appData.Flags["input"] = "Ascii~"
	appData.Flags["reff"] = "Ascii"
	appData.Flags["color"] = "#FABB60"
}
