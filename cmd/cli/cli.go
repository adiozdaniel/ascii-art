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
	loadCli()
}

func init() {
	NewCli(models.GetStateManager())
}

// loadCli runs the alignment mode of the application.
func loadCli() {
	var (
		inputChan                              = make(chan string)
		prevWidth                              int
		prevColor, prevReff, prevFont, tempStr string
	)

	go readInput(inputChan)

	for {
		select {
		case input := <-inputChan:
			if input == "exit" {
				fmt.Println("\n🤩 You were wonderful. Hope you enjoyed.\nExiting the Ascii-Art...")
				return
			} else if input != "" {
				tempStr = input
				helpers.ScanInput(input)
			}
		default:
			newWidth := helpers.GetTerminalWidth()
			if shouldUpdate(newWidth, prevWidth, tempStr, prevColor, prevReff, prevFont) {
				if cli.app.GetInput().Flags["font"] == "" {
					cli.app.GetInput().Flags["font"] = "--standard"
				}

				banner := cli.app.GetInput().BannerFile[cli.app.GetInput().Flags["font"]]
				err := helpers.FileContents(banner)
				if err != nil {
					fmt.Println(err)
				}
				outputs := ascii.Output(cli.app.GetInput().Flags["input"])
				termOutput := helpers.Alignment(outputs, newWidth)
				clearTerminal()
				fmt.Print(termOutput)
				resetCursor()

				prevWidth = newWidth
				tempStr = ""
				prevColor = cli.app.GetInput().Flags["color"]
				prevReff = cli.app.GetInput().Flags["reff"]
				prevFont = cli.app.GetInput().Flags["font"]
			}
			time.Sleep(2 * time.Second)
		}
	}
}

// readInput reads input from the CLI interface and sends it to the input channel.
func readInput(inputChan chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputChan <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
	close(inputChan)
}

// shouldUpdate checks if the terminal output needs to be updated.
func shouldUpdate(newWidth, prevWidth int, tempStr, prevColor, prevReff, prevFont string) bool {
	if cli.app.GetInput().Flags["input"] == "" {
		return false
	}
	return newWidth != prevWidth || tempStr != "" || cli.app.GetInput().Flags["color"] != prevColor || cli.app.GetInput().Flags["reff"] != prevReff || cli.app.GetInput().Flags["font"] != prevFont
}

// clearTerminal clears the terminal screen.
func clearTerminal() {
	fmt.Print("\033[H\033[2J\033[3J\033[?25h")
}

// resetCursor resets the terminal cursor to the start.
func resetCursor() {
	fmt.Print("\033[999;1H")
}
