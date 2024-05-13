package art_work

import "strings"

func Output(input, fileContents []string) string {
	var art_work strings.Builder

	ascii_map := AsciiMap(fileContents)

	for _, word := range input {
		for i := 0; i < 8; i++ {
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
