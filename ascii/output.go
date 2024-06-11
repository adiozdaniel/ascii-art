package ascii

import (
	"fmt"
	"strings"

	"github.com/adiozdaniel/ascii-art/utils"
)

// The function output now Writes our desired Output on the command line
func Output(reff string, input []string, fileContents []string) string {
	var art_work strings.Builder
	ascii_map := AsciiMap(fileContents)
	reset := "\033[0m"
	str := cleanInput(input)

	var height int
	for _, word := range str {
		if word == "" {
			height = 1
		} else {
			height = 8
		}
		for i := 0; i < height; i++ {
			var builder strings.Builder
			for _, char := range word {
				if ascii, ok := ascii_map[char]; ok {
					colorName := GetColorCode()
					if strings.Contains(reff, string(char)) {
						builder.WriteString(colorName + fileContents[ascii+i] + reset)
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

func cleanInput(input []string) []string {
	_, flag := utils.GetFile()

	for i, word := range input {
		if word == flag {
			input[i] = ""
		}
	}

	for i, word := range input {
		if strings.Contains(word, "--color") || strings.Contains(word, "--output") {
			input = input[i+1:]
		}
	}
	fmt.Println(input)
	return input
}
