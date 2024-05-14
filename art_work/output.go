package art_work

import (
	"strings"
)

func Output(input, fileContents []string) string {
	var art_work strings.Builder
	ascii_map := AsciiMap(fileContents)

	for _, word := range input {
		nonAsciis := make(map[rune]string)
		for i := 0; i < 8; i++ {
			var builder strings.Builder
			for _, char := range word {
				if ascii, ok := ascii_map[char]; ok {
					builder.WriteString(fileContents[ascii+i])
				} else {
					if _, found := nonAsciis[char]; !found {
						nonAsciis[char] = string(char)
					}
				}
			}
			if i == 4 {
				for _, char := range nonAsciis {
					builder.WriteString(char)
				}
			}
			art_work.WriteString(builder.String())
			art_work.WriteRune('\n')
		}
	}
	return art_work.String()
}
