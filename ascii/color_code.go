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
		{color: "Reset", ansicode: "\033[0m"},
		{color: "Red", ansicode: "\033[31m"},
		{color: "Green", ansicode: "\033[32m"},
		{color: "Yellow", ansicode: "\033[33m"},
		{color: "Blue", ansicode: "\033[34m"},
		{color: "Magenta", ansicode: "\033[35m"},
		{color: "Cyan", ansicode: "\033[36m"},
		{color: "Gray", ansicode: "\033[37m"},
		{color: "White", ansicode: "\033[97m"},
		{color: "Orange", ansicode: "\033[38;5;208m"},
	}
}

func GetColorCode(name string) string {
	name = strings.ToLower(name)
	for _, c := range asciiColors {
		if strings.ToLower(c.color) == name {
			return c.ansicode
		}
	}
	return "\033[97m" // Default to white
}
