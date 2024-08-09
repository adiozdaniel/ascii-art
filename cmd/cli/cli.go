package main

import (
	"github.com/adiozdaniel/ascii-art/internals/ascii"
	"github.com/adiozdaniel/ascii-art/internals/models"
)

// Cli holds the state for the cli interface.
type Cli struct {
	app *models.StateManager
}

// NewCli creates a new Cli instance.
func NewCli(sm *models.StateManager) *Cli {
	return &Cli{app: sm}
}

// app is the input data for the application.
var cli Cli

func main() {
	cli.app.GetInput().Init()
	ascii.LoadCli()
}

// init initializes the cli interface
func init() {
	cli = *NewCli(models.GetStateManager())

	bc, err := cli.app.GetConfig().CreateBannerCache()
	if err != nil {
		cli.app.GetInput().ErrorHandler("banners")
	}

	cli.app.GetConfig().BannerFileCache = bc

	cli.app.GetInput().Flags["font"] = "--standard"
	cli.app.GetInput().Flags["input"] = "Ascii~"
	cli.app.GetInput().Flags["reff"] = "Ascii"
	cli.app.GetInput().Flags["color"] = "#FABB60"
}
