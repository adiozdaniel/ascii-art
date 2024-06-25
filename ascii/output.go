package ascii

import (
	"strings"

	"github.com/adiozdaniel/ascii-art/utils"
)

// Output compiles the banner characters to form the desired ascii art work
func Output(fileContents []string) string {
	if utils.Inputs.Input == "" {
		return ""
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
			for _, char := range word {
				if ascii, ok := ascii_map[char]; ok {
					if color != "" {
						colorCode := GetColorCode(color)
						if strings.Contains(reff, string(char)) {
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
