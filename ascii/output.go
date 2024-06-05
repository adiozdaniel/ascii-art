package ascii

import (
	"strings"
)

var (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
	Orange  = "\033[38;5;208m"
)

// The function output now Writes our desired Output on the command line
func Output(color string, reff string, input []string, fileContents []string) string {
	switch color {
	case "Red" , "red":
		color = Red
	case "Green", "green":
		color = Green
	case "Yellow" , "yellow":
		color = Yellow
	case "Blue" , "blue":
		color = Blue
	case "Magenta", "magneta":
		color = Magenta
	case "Cyan" , "cyan":
		color = Cyan
	case "Gray" , "gray":
		color = Gray
	case "White" , "white":
		color = White
	case "Orange" , "orange":
		color = Orange
	default:
		color = White
	}
	var art_work strings.Builder
	ascii_map := AsciiMap(fileContents)

	var height int
	for _, word := range input {
		if word == "" {
			height = 1
		} else {
			height = 8
		}
		for i := 0; i < height; i++ {
			var builder strings.Builder
			for _, char := range word {
				if ascii, ok := ascii_map[char]; ok {
					if strings.Contains(reff, string(char)) {
						builder.WriteString(color + fileContents[ascii+i] + Reset)
					} else {
						builder.WriteString(fileContents[ascii+i])
					}
				}
			}
			art_work.WriteString(builder.String())
			art_work.WriteRune('\n')
		}
	}
	return art_work.String()
}
