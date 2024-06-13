package ascii

import (
	"fmt"
	"strconv"
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
		{color: "Pink", ansicode: "\033[38;5;205m"},
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
	if len(color) == 0 {
		return "\033[97m" // Default to white
	}

	for _, c := range asciiColors {
		if strings.EqualFold(c.color, color) {
			return c.ansicode
		}
	}

	if color[0] == ('#') {
		myColor, _ := hexToRGB(color)
		return myColor
	}

	ansciColor, err := getAnsiColor(color)
	if err != nil {
		return "\033[97m" // Default to white
	}
	return ansciColor
}

func getAnsiColor(s string) (string, error) {
	if !strings.Contains(s, "(") && !strings.Contains(s, ")"){
		return "", fmt.Errorf("not rgb")
	}

	temp1 := strings.Split(s, "(")[1]
	temp2 := strings.Split(temp1, ")")[0]
	colorSlice := strings.Split(temp2, ",")
	red, err := strconv.Atoi(strings.TrimSpace(colorSlice[0]))
	if err != nil {
		return "", err
	}
	green, err := strconv.Atoi(strings.TrimSpace(colorSlice[1]))
	if err != nil {
		return "", err
	}
	blue, err := strconv.Atoi(strings.TrimSpace(colorSlice[2]))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", red, green, blue), nil
}

func hexToRGB(hex string) (string, error) {
	if len(hex) == 4 {
		hex = fmt.Sprintf("#%c%c%c%c%c%c", hex[1], hex[1], hex[2], hex[2], hex[3], hex[3])
	} else if len(hex) != 7 { // #RRGGBB
		return "", fmt.Errorf("invalid hex color: %s", hex)
	}
	red, err := strconv.ParseInt(hex[1:3], 16, 64)
	if err != nil {
		return "", err
	}
	green, err := strconv.ParseInt(hex[3:5], 16, 64)
	if err != nil {
		return "", err
	}
	blue, err := strconv.ParseInt(hex[5:7], 16, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", red, green, blue), nil
}
