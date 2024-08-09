package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/adiozdaniel/ascii-art/internals/ascii"
	"github.com/adiozdaniel/ascii-art/internals/models"
)

// Cli holds the state for the CLI interface.
type Cli struct {
	app                                    *models.StateManager
	inputChan                              chan string
	prevWidth                              int
	prevColor, prevReff, prevFont, tempStr string
}

// NewCli creates a new Cli instance.
func NewCli(sm *models.StateManager) *Cli {
	return &Cli{
		app:       sm,
		inputChan: make(chan string),
	}
}

// readInput reads input from the CLI interface and sends it to the input channel.
func (cli *Cli) readInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cli.inputChan <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
	close(cli.inputChan)
}

// init initializes the CLI interface.
func init() {
	cli := NewCli(models.GetStateManager())

	bc, err := cli.app.GetConfig().CreateBannerCache()
	if err != nil {
		cli.app.GetInput().ErrorHandler("banners")
	}

	go cli.readInput()
	cli.app.GetConfig().BannerFileCache = bc

	cli.app.GetInput().Flags["font"] = "--standard"
	cli.app.GetInput().Flags["input"] = "Ascii~"
	cli.app.GetInput().Flags["reff"] = "Ascii"
	cli.app.GetInput().Flags["color"] = "#FABB60"
}

func main() {
	cli := NewCli(models.GetStateManager())
	cli.app.GetInput().Init()

}
