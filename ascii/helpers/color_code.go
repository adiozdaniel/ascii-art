package helpers

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/adiozdaniel/ascii-art/utils"
)

// Color struct stores the color name and its ansicode equivalent
type Color struct {
	color, ansicode string
}

// asciiColors global slice stores Color structs
var asciiColors []Color

// init populates asciiColors slice with data
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
	if len(color) < 3 {
		utils.ErrorHandler("color")
	}

	if strings.HasPrefix(color, "#") {
		myColor, err := hexToRGB(color)
		if err == nil {
			return myColor
		}
	}

	if strings.HasPrefix(color, "hsl(") || strings.HasPrefix(color, "HSL(") && strings.HasSuffix(color, ")") {
		myColor, err := getHSLColor(color)
		if err == nil {
			return myColor
		}
	}

	if (strings.HasPrefix(color, "rgb(") || strings.HasPrefix(color, "RGB(")) && strings.HasSuffix(color, ")") {
		ansciColor, err := getAnsiColor(color)
		if err == nil {
			return ansciColor
		}
	}

	for _, c := range asciiColors {
		if strings.EqualFold(c.color, color) {
			return c.ansicode
		}
	}

	utils.ErrorHandler("color")
	return ""
}

// getAnsiColor converts rgb color format to ansicodes
func getAnsiColor(s string) (string, error) {
	temp1 := strings.Split(s, "(")[1]
	temp2 := strings.Split(temp1, ")")[0]
	colorSlice := strings.Split(temp2, ",")

	if len(colorSlice) != 3 {
		return "", fmt.Errorf("invalid rgb format")
	}

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

// hexToRGB converts hexadecimal colors to ansicodes
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

// getHSLColor converts hsl color formats to ansicodes
func getHSLColor(str string) (string, error) {
	temp1 := strings.Split(str, "(")[1]
	temp2 := strings.Split(temp1, ")")[0]
	colorSlice := strings.Split(temp2, ",")

	if len(colorSlice) != 3 {
		return "", fmt.Errorf("invalid hsl format")
	}

	h, err := strconv.ParseFloat(strings.TrimSpace(colorSlice[0]), 64)
	if err != nil {
		return "", err
	}
	s, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(colorSlice[1], "%")), 64)
	if err != nil {
		return "", err
	}
	s /= 100
	l, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(colorSlice[2], "%")), 64)
	if err != nil {
		return "", err
	}
	l /= 100

	r, g, b := hslToRGB(h, s, l)
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b), nil
}

// hslToRGB helper function to calculate hsl values to rgb
func hslToRGB(h, s, l float64) (int, int, int) {
	c := (1 - abs(2*l-1)) * s
	x := c * (1 - abs(math.Mod(h/60.0, 2)-1))
	m := l - c/2

	var r, g, b float64
	if 0 <= h && h < 60 {
		r, g, b = c, x, 0
	} else if 60 <= h && h < 120 {
		r, g, b = x, c, 0
	} else if 120 <= h && h < 180 {
		r, g, b = 0, c, x
	} else if 180 <= h && h < 240 {
		r, g, b = 0, x, c
	} else if 240 <= h && h < 300 {
		r, g, b = x, 0, c
	} else {
		r, g, b = c, 0, x
	}

	return int((r + m) * 255), int((g + m) * 255), int((b + m) * 255)
}

// abs helper function to always return positive numbers
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
