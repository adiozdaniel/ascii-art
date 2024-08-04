//go:build cli
// +build cli

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	appconfig "github.com/adiozdaniel/ascii-art/internals/app_config"
	"github.com/adiozdaniel/ascii-art/internals/ascii"
	"github.com/adiozdaniel/ascii-art/pkg/helpers"
)

// get the app state manager
var (
	sm  = appconfig.GetStateManager()
	app = sm.GetInput()
)

func main() {
	app.Init()
	loadCli()
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
				if app.Flags["font"] == "" {
					app.Flags["font"] = "--standard"
				}

				banner := app.BannerFile[app.Flags["font"]]
				err := helpers.FileContents(banner)
				if err != nil {
					fmt.Println(err)
				}
				outputs := ascii.Output(app.Flags["input"])
				termOutput := helpers.Alignment(outputs, newWidth)
				clearTerminal()
				fmt.Print(termOutput)
				resetCursor()

				prevWidth = newWidth
				tempStr = ""
				prevColor = app.Flags["color"]
				prevReff = app.Flags["reff"]
				prevFont = app.Flags["font"]
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
	if app.Flags["input"] == "" {
		return false
	}
	return newWidth != prevWidth || tempStr != "" || app.Flags["color"] != prevColor || app.Flags["reff"] != prevReff || app.Flags["font"] != prevFont
}

// clearTerminal clears the terminal screen.
func clearTerminal() {
	fmt.Print("\033[H\033[2J\033[3J\033[?25h")
}

// resetCursor resets the terminal cursor to the start.
func resetCursor() {
	fmt.Print("\033[999;1H")
}
