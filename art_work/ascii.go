package art_work

import "strings"

func Output(input, outputFile []string) string {
	var art_work strings.Builder

	ascii_map := asciiMap(outputFile)

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
					builder.WriteString(outputFile[ascii+i])
				}
			}
			art_work.WriteString(builder.String())
			art_work.WriteRune('\n')
		}
	}
	return art_work.String()
}
