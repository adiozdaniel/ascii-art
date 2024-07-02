package ascii

import (
	"fmt"
	"os"
	"strings"

	"github.com/adiozdaniel/ascii-art/utils"
)

// Output compiles the banner characters to form the desired ascii art work
func Output(fileContents []string) string {
	if utils.Inputs.Input == "" {
		return ""
	}

	if utils.Inputs.Input == "\\n" && !utils.Inputs.IsWeb {
		fmt.Println()
		os.Exit(0)
	}

	var art_work strings.Builder
	var ascii_map = AsciiMap(fileContents)
	var reset = "\033[0m"
	var height int = 8
	var color = strings.TrimSpace(utils.Inputs.Color)
	var reff = utils.Inputs.ColorRef
	var input = strings.Split(strings.ReplaceAll(utils.Inputs.Input, "\\n", "\n"), "\n")

	for _, word := range input {
		if word == "" {
			height = 1
		} else {
			height = 8
		}

		for i := 0; i < height; i++ {
			var builder strings.Builder
			for j, char := range word {
				if ascii, ok := ascii_map[char]; ok {
					if color != "" {
						colorCode := GetColorCode(color)

						if containsReff(word) && j >= indexes.startIndex && j < indexes.endIndex {
							builder.WriteString(colorCode + fileContents[ascii+i] + reset)
						} else if strings.Contains(reff, string(char)) && indexes.startIndex == 0 {
							builder.WriteString(colorCode + fileContents[ascii+i] + reset)
						} else {
							builder.WriteString(fileContents[ascii+i])
						}
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

// contains hosts the startIndex and endIndex for substrings
type contains struct {
	startIndex, endIndex int
}

// indexes is a placeholder for the struct
var indexes contains

// containsReff checks for color substrings and initialises contains struct
func containsReff(word string) bool {
	var reff = utils.Inputs.ColorRef
	var input = word
	var x, y = len(input), len(reff)

	if y > x {
		return false
	}

	for i := 0; i < x; i++ {
		lastIndex := y + i
		if lastIndex > x {
			lastIndex = x - 1
		}
		if input[i:lastIndex] == reff {
			indexes.startIndex = i
			indexes.endIndex = lastIndex
			return true
		}
	}
	return false
}
