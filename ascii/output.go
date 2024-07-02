package ascii

import (
	"fmt"
	"os"
	"strings"

	"github.com/adiozdaniel/ascii-art/utils"
)

// variables declaration
var art_work strings.Builder
var reset = "\033[0m"
var height int = 8
var color = strings.TrimSpace(utils.Inputs.Color)
var reff = utils.Inputs.ColorRef
var input = strings.Split(strings.ReplaceAll(utils.Inputs.Input, "\\n", "\n"), "\n")

// Output compiles the banner characters to form the desired ascii art work
func Output(fileContents []string) string {
	if utils.Inputs.Input == "" {
		return ""
	}

	if utils.Inputs.Input == "\\n" && !utils.Inputs.IsWeb {
		fmt.Println()
		os.Exit(0)
	}

	var ascii_map = AsciiMap(fileContents)
	for index, word := range input {
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

						if containsReff() {
							if word == indexes.lineIndex[index] && j >= indexes.startIndex[word] && j < indexes.endIndex[word] {
								builder.WriteString(colorCode + fileContents[ascii+i] + reset)
							} else {
								builder.WriteString(fileContents[ascii+i])
							}
						} else if strings.Contains(reff, string(char)) {
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
	startIndex, endIndex map[string]int
	lineIndex            map[int]string
}

// indexes is a placeholder for the struct
var indexes = contains{
	startIndex: make(map[string]int),
	endIndex:   make(map[string]int),
	lineIndex:  make(map[int]string),
}

// containsReff checks for color substrings and initialises contains struct
func containsReff() bool {
	var hasReff bool

	for i, line := range input {
		var x, y = len(line), len(reff)

		for j := 0; j < len(line); j++ {
			lastIndex := y + j

			if lastIndex > x {
				lastIndex = x - 1
			}

			if line[j:lastIndex] == reff {
				indexes.startIndex[line] = j
				indexes.endIndex[line] = lastIndex
				indexes.lineIndex[i] = line
				hasReff = true
			}
		}
	}

	return hasReff
}
