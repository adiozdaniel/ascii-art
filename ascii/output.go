package ascii

import (
	"strings"

	"github.com/adiozdaniel/ascii-art/utils"
)

// global variables declaration
var (
	reset               = "\033[0m"
	reff                = utils.Inputs.ColorRef
	height          int = 8
)

// Output compiles the banner characters to form the desired ascii art work
func Output(inputStr string) string {
	if inputStr == "" {
		return ""
	}

	if inputStr == "\\n" {
		return "\n"
	}

	var fileContents, _     = FileContents()
	var ascii_map = AsciiMap(fileContents)
	var art_work strings.Builder

	if inputStr != "" {
		utils.Inputs.Input = strings.ReplaceAll(inputStr, "\\n", "\n")
	}

	if utils.Inputs.IsWeb {
		processWebInput(ascii_map, fileContents, &art_work)
	} else {
		processTerminalInput(ascii_map, fileContents, &art_work)
	}

	return art_work.String()
}

// processWebInput processes input from the web
func processWebInput(ascii_map map[rune]int, fileContents []string, art_work *strings.Builder) {
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
func processTerminalInput(ascii_map map[rune]int, fileContents []string, art_work *strings.Builder) {
	input := strings.Split(utils.Inputs.Input, "\n")
	color := strings.TrimSpace(utils.Inputs.Color)

	for lineIndex, line := range input {
		if line == "" {
			height = 1
		} else {
			height = 8
		}

		for i := 0; i < height; i++ {
			var builder strings.Builder
			for j, char := range line {
				if ascii, ok := ascii_map[char]; ok {
					if char == ' ' && utils.Inputs.Justify == "justify" {
						builder.WriteRune('$')
						continue
					}
					if color != "" {
						colorCode := GetColorCode(color)

						if containsReff(input) {
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
func containsReff(input []string) bool {
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
