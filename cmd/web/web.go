package main

import (
	"fmt"
	"net/http"

	"github.com/adiozdaniel/ascii-art/internals/ascii"
	"github.com/adiozdaniel/ascii-art/internals/handlers"
	"github.com/adiozdaniel/ascii-art/internals/middlewares"
	"github.com/adiozdaniel/ascii-art/internals/models"
	"github.com/adiozdaniel/ascii-art/internals/routes"
	"github.com/adiozdaniel/ascii-art/pkg/helpers"
)

// get the app state manager
var (
	sm             = models.GetStateManager()
	appData        = sm.GetInput()
	appConfig      = sm.GetConfig()
	sessionManager = sm.GetSessionManager()
)

// runWeb initializes the web data
func runWeb() error {
	handlers.NewRepo(sm)
	middlewares.NewMiddlewares(models.GetStateManager().GetSessionManager())

	tc, err := appConfig.CreateTemplateCache()
	if err != nil {
		return err
	}

	appConfig.TemplateCache = tc

	bc, err := appConfig.CreateBannerCache()
	if err != nil {
		return err
	}

	appConfig.BannerFileCache = bc

	appData.Flags["font"] = "--standard"
	appData.Flags["input"] = "Ascii~"
	appData.Flags["reff"] = "Ascii"
	appData.Flags["color"] = "#FABB60"

	return nil
}

// main starts the web server
func main() {
	appData.Init()
	err := runWeb()
	if err != nil {
		appData.ErrorHandler("web")
	}

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	wrappedMux := middlewares.SessionMiddleware(
		sessionManager)(middlewares.RouteChecker(mux))

	server := &http.Server{
		Addr:    ":8080",
		Handler: wrappedMux,
	}

	banner := appData.BannerFile[appData.Flags["font"]]
	err = helpers.FileContents(banner)
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
