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
		inputChan                                = make(chan string)
		prevWidth                                int
		prevColor, prevReff, prevBanner, tempStr string
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
			if shouldUpdate(newWidth, prevWidth, tempStr, prevColor, prevReff, prevBanner) {
				outputs := ascii.Output(utils.Inputs.Input)
				termOutput := helpers.Alignment(outputs, newWidth)
				clearTerminal()
				fmt.Print(termOutput)
				resetCursor()

				prevWidth = newWidth
				tempStr = ""
				prevColor = utils.Inputs.Color
				prevReff = utils.Inputs.ColorRef
				prevBanner = utils.Inputs.BannerPath
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
func shouldUpdate(newWidth, prevWidth int, tempStr, prevColor, prevReff, prevBanner string) bool {
	return newWidth != prevWidth || tempStr != "" || utils.Inputs.Color != prevColor || utils.Inputs.ColorRef != prevReff || utils.Inputs.BannerPath != prevBanner
}

// clearTerminal clears the terminal screen.
func clearTerminal() {
	fmt.Print("\033[H\033[2J\033[3J\033[?25h")
}

// resetCursor resets the terminal cursor to the start.
func resetCursor() {
	fmt.Print("\033[999;1H")
}

// isValidAlignment checks if the provided alignment is valid.
func isValidAlignment(alignment string) bool {
	return alignment == "left" || alignment == "center" || alignment == "right" || alignment == "justify"
}

// isBannerFile checks if the provided word is a banner file.
func isBannerFile(word string) bool {
	_, exists := utils.BannerFiles[word]
	return exists
}
