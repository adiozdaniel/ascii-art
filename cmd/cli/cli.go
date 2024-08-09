package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/adiozdaniel/ascii-art/internals/ascii"
	"github.com/adiozdaniel/ascii-art/internals/models"
	"github.com/adiozdaniel/ascii-art/pkg/helpers"
)

// Global Cli instance
var cli *Cli

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
	cli = NewCli(models.GetStateManager())

	bc, err := cli.app.GetConfig().CreateBannerCache()
	if err != nil {
		cli.app.GetInput().ErrorHandler("banners")
	}

	go cli.readInput()
	cli.app.GetConfig().BannerFileCache = bc

	cli.app.GetInput().Flags["font"] = "--standard"
	cli.app.GetInput().Flags["input"] = "Ascii~"
}

// shouldUpdate checks if the terminal output needs to be updated.
func shouldUpdate(newWidth, prevWidth int, tempStr, prevColor, prevReff, prevFont string) bool {
	if cli.app.GetInput().Flags["input"] == "" {
		return false
	}
	return newWidth != prevWidth || tempStr != "" || cli.app.GetInput().Flags["color"] != prevColor || cli.app.GetInput().Flags["reff"] != prevReff || cli.app.GetInput().Flags["font"] != prevFont
}

func main() {
	cli.app.GetInput().Init()

	for {
		select {
		case input, ok := <-cli.inputChan:
			if !ok {
				return
			}
			if input == "exit" {
				fmt.Println("\n🤩 You were wonderful. Hope you enjoyed.\nExiting the Ascii-Art...")
				return
			} else if input != "" {
				cli.tempStr = input
				helpers.ScanInput(input)
			}
		default:
			newWidth := helpers.GetTerminalWidth()
			if shouldUpdate(newWidth, cli.prevWidth, cli.tempStr, cli.prevColor, cli.prevReff, cli.prevFont) {
				banner := cli.app.GetInput().BannerFile[cli.app.GetInput().Flags["font"]]
				err := helpers.FileContents(banner)
				if err != nil {
					fmt.Println("Error loading banner file:", err)
				}
				outputs := ascii.Output(cli.app.GetInput().Flags["input"])
				termOutput := helpers.Alignment(outputs, newWidth)
				helpers.ClearTerminal()
				fmt.Print(termOutput)
				helpers.ResetCursor()

				cli.prevWidth = newWidth
				cli.tempStr = ""
				cli.prevColor = cli.app.GetInput().Flags["color"]
				cli.prevReff = cli.app.GetInput().Flags["reff"]
				cli.prevFont = cli.app.GetInput().Flags["font"]
			}
			time.Sleep(2 * time.Second)
		}
	}
}
