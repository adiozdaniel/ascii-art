package utils

import (
	"fmt"
	"strings"
)

//NonAsciiOutput returns the non-printable ascii characters
func NonAsciiOutput() string {
	var artWork strings.Builder
	nonAsciiMap := NonAsciiMap(Inputs.Input)

	for _, char := range nonAsciiMap {
		artWork.WriteString(string(char) + " ")
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
