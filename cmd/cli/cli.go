package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/adiozdaniel/ascii-art/internals/ascii"
	"github.com/adiozdaniel/ascii-art/pkg/helpers"
	"github.com/adiozdaniel/ascii-art/pkg/utils"
)

// global variables
var (
	app = &utils.Inputs
)

func main() {
	app.Init()
	loadCli()
}

// justified runs the alignment mode of the application.
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
				os.Exit(0)
			} else if input != "" {
				tempStr = input
				helpers.ScanInput(input)
			}
		default:
			newWidth := helpers.GetTerminalWidth()
			if shouldUpdate(newWidth, prevWidth, tempStr, prevColor, prevReff, prevFont) {
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
	close(inputChan)
}

// shouldUpdate checks if the terminal output needs to be updated.
func shouldUpdate(newWidth, prevWidth int, tempStr, prevColor, prevReff, prevFont string) bool {
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
