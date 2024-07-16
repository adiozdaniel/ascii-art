package ascii

import (
	"strings"

	"github.com/adiozdaniel/ascii-art/utils"
)

// variables declaration
var (
	reset               = "\033[0m"
	color               = strings.TrimSpace(utils.Inputs.Color)
	reff                = utils.Inputs.ColorRef
	height          int = 8
	fileContents, _     = FileContents()
)

// Output compiles the banner characters to form the desired ascii art work
func Output(inputStr string) string {
	if inputStr == "" {
		return ""
	}

	if inputStr == "\\n" {
		return "\n"
	}

	ascii_map := AsciiMap(fileContents)
	input := strings.Split(strings.ReplaceAll(inputStr, "\\n", "\n"), "\n")
	var art_work strings.Builder

	if utils.Inputs.IsWeb {
		processWebInput(ascii_map, fileContents, &art_work)
	} else if utils.Inputs.Justify == "justify" {
		justifyAlign(ascii_map, fileContents, &art_work)
	} else {
		processTerminalInput(ascii_map, fileContents, input, &art_work)
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
func processTerminalInput(ascii_map map[rune]int, fileContents, input []string, art_work *strings.Builder) {
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

func justifyAlign(ascii_map map[rune]int, fileContents []string, art_work *strings.Builder) {
	width := utils.GetTerminalWidth()

	for _, line := range strings.Split(utils.Inputs.Input, "\n") {
		words := strings.Fields(line)
		wordsLength := 0
		for _, word := range words {
			wordsLength += len(word)
		}
		totalSpaces := width - wordsLength
		spaceSlots := len(words) - 1
		evenSpaces := totalSpaces / spaceSlots
		// extraSpaces := totalSpaces % spaceSlots
		spacesIndex :=  mapSpace(line)
		for i := 0; i < 8; i++ {
			var builder strings.Builder
			for j, char := range line {
				
				if ascii, ok := ascii_map[char]; ok {
					if ok := spacesIndex[j] ; ok {
						builder.WriteString(strings.Repeat(" ", evenSpaces))
					} else {
						builder.WriteString(fileContents[ascii+i])
					}
					
				}
				
				// if j < spaceSlots {
				// 	builder.WriteString(strings.Repeat(" ", evenSpaces))
				// 	if j < extraSpaces {
				// 		builder.WriteString(" ")
				// 		extraSpaces--
				// 	}
				// }
			}
			art_work.WriteString(builder.String())
			art_work.WriteRune('\n')
		}
		art_work.WriteRune('\n')
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

func mapSpace (input string) map[int]bool {
	var spaceMap  = make(map[int]bool)
	for i , char := range input {
		if char == ' ' {
			spaceMap[i] = true
		}
	}
	return spaceMap
}
