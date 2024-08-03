package helpers

import (
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

// isBannerFile checks if the provided word is a banner file.
func isBannerFile(word string) bool {
	_, exists := app.BannerFile[word]
	return exists
}
