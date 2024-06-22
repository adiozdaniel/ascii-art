package ascii

import (
	"strings"

	"github.com/adiozdaniel/ascii-art/utils"
)

// The function output now Writes our desired Output on the command line
func Output(fileContents []string) string {
	var art_work strings.Builder
	var ascii_map = AsciiMap(fileContents)
	var reset = "\033[0m"
	var height int = 8
	var color = utils.Inputs.Color
	var reff = utils.Inputs.ColorRef
	var word = utils.Inputs.Input

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

	return art_work.String()
}
