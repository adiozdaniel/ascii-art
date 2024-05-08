package utilities

import "strings"

func NonAsciiOutput(input []string) string {
	var art_work strings.Builder

	non_ascii_map := NonAsciiMap(input)

	for _, word := range input {

		var builder strings.Builder
		for _, char := range word {
			if _, ok := non_ascii_map[char]; ok {
				builder.WriteString("The non ascii character "+ string(char) +" was not printed\n")
			}
		}
		art_work.WriteString(builder.String())
		art_work.WriteRune('\n')

	}
	return art_work.String()
}
