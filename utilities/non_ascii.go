package utilities

import (
	"fmt"
	"strings"
)

func NonAsciiOutput(input []string) string {
	var artWork strings.Builder
	nonAsciiMap := NonAsciiMap(input)

	for _, char := range nonAsciiMap {
		artWork.WriteString(string(char) + " ")
	}

	count := len(nonAsciiMap)
	word := "character"
	plural := "was"
	if count > 1 {
		word = "characters"
		plural = "were"
	}
	artWork.WriteString(fmt.Sprintf("These %d %s %s printed only once!\n", count, word, plural))
	return artWork.String()
}
