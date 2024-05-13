package utilities

import (
	"strings"
	// "unicode/utf8"
)

func NonAsciiOutput(input []string) (string) {
	var art_work strings.Builder
	non_ascii_map := NonAsciiMap(input)
	printedChars := make(map[rune]bool)
	// effected := false

	for _, word := range input {

		var builder strings.Builder
		for _, char := range word {
			if _, ok := non_ascii_map[char]; ok && !printedChars[char] {
				builder.WriteString("The non ascii character "+ string(char) +" was not printed\n")
				printedChars[char] = true
				// effected = true
			}
		}
		art_work.WriteString(builder.String())
		// art_work.WriteRune('\n')

	}
	return art_work.String() 
}
