package ascii

import (
	"fmt"
	"os"
	"strings"

	"github.com/adiozdaniel/ascii-art/utils"
)

// variables declaration
var reset = "\033[0m"
var color = strings.TrimSpace(utils.Inputs.Color)
var reff = utils.Inputs.ColorRef
var input = strings.Split(strings.ReplaceAll(utils.Inputs.Input, "\\n", "\n"), "\n")
var art_work strings.Builder
var height int = 8

// Output compiles the banner characters to form the desired ascii art work
func Output(fileContents []string) string {
	if strings.TrimSpace(utils.Inputs.Input) == "" {
		return ""
	}

	var ascii_map = AsciiMap(fileContents)

	if utils.Inputs.IsWeb {
		processWebInput(ascii_map, fileContents)
	} else {
		processTerminalInput(ascii_map, fileContents)
	}

	return art_work.String()
}

// processWebInput processes input from the web
func processWebInput(ascii_map map[rune]int, fileContents []string) {
	for _, line := range strings.Split(utils.Inputs.Input, "\n") {
		for i := 0; i < height; i++ {
			var builder strings.Builder
			for _, char := range line {
				if ascii, ok := ascii_map[char]; ok {
					{
						builder.WriteString(fileContents[ascii+i])
					}
				}
			}
			art_work.WriteString(builder.String())
			art_work.WriteRune('\n')
		}
		art_work.WriteRune('\n')
	}
}

// processTerminalInput processes input from the internal
func processTerminalInput(ascii_map map[rune]int, fileContents []string) {
	if utils.Inputs.Input == "\\n" {
		fmt.Println()
		os.Exit(0)
	}

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

		for j := 0; j <= x-y; j++ {
			if line[j:j+y] == reff {
				indexes.startIndex[line] = j
				indexes.endIndex[line] = j + y
				indexes.lineIndex[i] = line
				hasReff = true
			}
		}
	}

	return hasReff
}
