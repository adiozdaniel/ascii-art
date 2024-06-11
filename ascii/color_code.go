package ascii

import (
	"strings"
)

type Color struct {
	color, ansicode string
}

var asciiColors []Color

func init() {
	asciiColors = []Color{
		{color: "Red", ansicode: "\033[31m"},
		{color: "Green", ansicode: "\033[32m"},
		{color: "Yellow", ansicode: "\033[33m"},
		{color: "Blue", ansicode: "\033[34m"},
		{color: "Magenta", ansicode: "\033[35m"},
		{color: "Cyan", ansicode: "\033[36m"},
		{color: "Gray", ansicode: "\033[37m"},
		{color: "White", ansicode: "\033[97m"},
		{color: "Orange", ansicode: "\033[38;5;208m"},
		{color: "Purple", ansicode: "\033[95m"},
		{color: "Lime", ansicode: "\033[38;5;118m"},
		{color: "pink", ansicode: "\033[38;5;205m"},
		{color: "Teal", ansicode: "\033[38;5;37m"},
		{color: "Lavender", ansicode: "\033[38;5;183m"},
		{color: "Brown", ansicode: "\033[38;5;130m"},
		{color: "Beige", ansicode: "\033[38;5;230m"},
		{color: "Maroon", ansicode: "\033[38;5;52m"},
		{color: "Mint", ansicode: "\033[38;5;121m"},
		{color: "Olive", ansicode: "\033[38;5;142m"},
		{color: "Apricot", ansicode: "\033[38;5;215m"},
		{color: "Navy", ansicode: "\033[38;5;18m"},
		{color: "Grey", ansicode: "\033[38;5;245m"},
		{color: "Black", ansicode: "\033[30m"},
	}
}

// GetColorCode gets the ANSI code of the input color after iterating through the structs in asciiColors
func GetColorCode(color string) string {
	for _, c := range asciiColors {
		if strings.ToLower(c.color) == color {
			return c.ansicode
		}
	}
	return "\033[97m" // Default to white
}
