package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

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
	sessionManager = sm.GetSessionManager()
)

// runWeb initializes the web data
func runWeb() http.Handler {
	handlers.NewRepo(sm)
	middlewares.NewMiddlewares(sessionManager)

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	wrappedMux := middlewares.SessionMiddleware(
		sessionManager)(middlewares.RouteChecker(mux))

	return wrappedMux
}

// shutdownChan channel closes the server gracefully
var shutdownChan = make(chan struct{})

// main starts the web server
func main() {
	appData.Init()
	wrappedMux := runWeb()

	server := &http.Server{
		Addr:    ":8080",
		Handler: wrappedMux,
	}

	banner := appData.BannerFile[appData.Flags["font"]]
	if err := helpers.FileContents(banner); err != nil {
		appData.ErrorHandler("fatal")
	}

	serverOutput := ascii.Output(appData.Flags["input"])
	fmt.Println(serverOutput + "=====================================\nserver running @http://localhost:8080")

	appData.Flags["isWeb"] = "true"

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appData.ErrorHandler("web")
			fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		}
	}()

	<-shutdownChan
	fmt.Println("Shutting down server...")

	if err := server.Shutdown(context.Background()); err != nil {
		appData.ErrorHandler("web")
		fmt.Fprintf(os.Stderr, "Server shutdown error: %v\n", err)
	} else {
		fmt.Println("Server shut down successfully.")
	}
}

// TriggerShutdown gracefully closes the shutdownChan
func TriggerShutdown() {
	close(shutdownChan)
}
