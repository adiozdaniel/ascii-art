package outputs

import (
	"fmt"
	"strings"
)

// AsciiMap maps the indices of the characters in the banner file
func AsciiMap(fileContents []string) map[rune]int {
	ascii_map := make(map[rune]int)
	ascii := 32

	for index, line := range fileContents {
		if len(line) == 0 || len(line) == 1 {
			ascii_map[rune(ascii)] = index + 1
			ascii++
		}
	}

	return ascii_map
}

// NonAsciiMap maps the non-printable ascii characters
func NonAsciiMap(input string) map[rune]rune {
	nonAsciiMap := make(map[rune]rune)

	for _, char := range input {
		if (char < 32 || char > 126) && char != 10 {
			nonAsciiMap[char] = char
		}
	}

	return nonAsciiMap
}

// NonAsciiOutput returns the non-printable ascii characters
func NonAsciiOutput() string {
	var artWork strings.Builder
	nonAsciiMap := NonAsciiMap(Inputs.Input)

	for _, char := range nonAsciiMap {
		if char != 13 {
			artWork.WriteString(string(char) + " ")
		}
	}

	count := len(nonAsciiMap)
	word := "character"
	plurals := []string{"This", "was"}
	if count > 1 {
		word = "characters"
		plurals[0] = "These"
		plurals[1] = "were"
	}
	artWork.WriteString(fmt.Sprintf("%s %s %s skipped!\n", plurals[0], word, plurals[1]))
	if len(artWork.String()) == 28 {
		return ""
	}

	return artWork.String()
}
