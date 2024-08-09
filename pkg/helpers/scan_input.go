package helpers

import (
	"fmt"
	"regexp"
	"strings"
)

// scanInput reads input from CLI interface and updates the input struct.
func ScanInput(input string) {
	cleanInput := app.RemoveQuotes(input)
	words := strings.Fields(cleanInput)

	app.Args = append(app.Args, words...)
	app.ParseArgs()
}

// isValidAlignment checks if the provided alignment is valid.
func isValidAlignment(alignment string) bool {
	return alignment == "left" || alignment == "center" || alignment == "right" || alignment == "justify"
}

// removeANSICodes removes the ansci escape codes
func RemoveANSICodes(input string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(input, "")
}

// clearTerminal clears the terminal screen.
func ClearTerminal() {
	fmt.Print("\033[H\033[2J\033[3J\033[?25h")
}

// resetCursor resets the terminal cursor to the start.
func ResetCursor() {
	fmt.Print("\033[999;1H")
}
