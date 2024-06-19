package ascii

import (
	"strings"
)

// The function output now Writes our desired Output on the command line
func Output(color, reff string, input []string, fileContents []string) string {
	var art_work strings.Builder
	ascii_map := AsciiMap(fileContents)
	reset := "\033[0m"

	var height int
	for _, word := range input {
		if word == "" {
			height = 1
		} else {
			height = 8
		}
		for i := 0; i < height; i++ {
			var builder strings.Builder
			for _, char := range word {
				if ascii, ok := ascii_map[char]; ok {
					colorCode := GetColorCode(color)
					if strings.Contains(reff, string(char)) {
						builder.WriteString(colorCode + fileContents[ascii+i] + reset)
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
