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
					builder.WriteString(fileContents[ascii+i])
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

	var ascii_map = AsciiMap(fileContents)
	for index, word := range input {
		if word == "" {
			height = 1
		} else {
			height = 8
		}

		for i := 0; i < height; i++ {
			var builder strings.Builder
			for j, char := range line {
				if ascii, ok := ascii_map[char]; ok {
					if color != "" {
						colorCode := GetColorCode(color)

						if containsReff() {
							if _, ok := indexes.indexMap[lineIndex][j]; ok {
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

// contains hosts the indexMap for substrings characters
type contains struct {
	indexMap map[int]map[int]int
}

// indexes is a placeholder for the struct
var indexes = contains{
	indexMap: make(map[int]map[int]int),
}

// containsReff checks for color substrings and initialises contains struct
func containsReff() bool {
	var hasReff bool

	for i, line := range input {
		x, y := len(line), len(reff)

		for j := 0; j <= x-y; j++ {
			if line[j:j+y] == reff {
				if _, ok := indexes.indexMap[i]; !ok {
					indexes.indexMap[i] = make(map[int]int)
				}
				for k := j; k < j+y; k++ {
					indexes.indexMap[i][k] = k
				}
				hasReff = true
			}
		}
	}

	return hasReff
}
