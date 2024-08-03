package helpers

import (
	"fmt"
	"os"
	"strings"
)

// scanInput reads input from CLI interface and updates the input struct.
func ScanInput(input string) {
	cleanInput := app.RemoveQuotes(input)
	words := strings.Fields(cleanInput)
	var newInput string

	for i, word := range words {
		switch {
		case strings.Contains(word, "--align") || strings.HasPrefix(word, "-align"):
			alignment := strings.TrimPrefix(strings.TrimPrefix(word, "--align="), "-align=")
			if isValidAlignment(alignment) {
				app.Justify = alignment
				continue
			}
			app.ErrorHandler("justify")
		case strings.Contains(word, "--color") || strings.Contains(word, "-color"):
			color := strings.TrimPrefix(strings.TrimPrefix(word, "--color="), "-color=")
			if color != "" {
				app.Color = color
				continue
			}
			app.ErrorHandler("colors")
		case strings.Contains(word, "--reff") || strings.Contains(word, "-reff"):
			reff := strings.TrimPrefix(strings.TrimPrefix(word, "--reff="), "-reff=")
			if reff != "" {
				app.ColorRef = reff
				continue
			}
			app.ErrorHandler("colors")
		case strings.Contains(word, "--output") || strings.Contains(word, "-output"):
			fmt.Println("🙄 Sorry, FS Mode cannot be set in alignment mode.")
			os.Exit(0)
		case strings.Contains(word, "--standard") || strings.Contains(word, "--thinkertoy") || strings.Contains(word, "--shadow"):
			if value, ok := app.BannerFile[word]; ok {
				app.Font = value
				continue
			}

			newInput += word + " "
		case isBannerFile(word):
			if i == len(words)-1 && len(words) != 1 {
				if value, ok := app.BannerFile[word]; ok {
					app.Font = value
				}
				break
			}
			newInput += word + " "
		default:
			newInput += word + " "
		}
	}

	if newInput != "" {
		app.Input = strings.TrimSpace(newInput)
	}
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
