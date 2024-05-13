package art_work

import (
	"strings"
)

func Output(input, fileContents []string) string {
	var art_work strings.Builder
	var height int

	ascii_map := AsciiMap(fileContents)

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
					builder.WriteString(fileContents[ascii+i])
				}
			}
			art_work.WriteString(builder.String())
			art_work.WriteRune('\n')
		}
	}
	return art_work.String()
}
