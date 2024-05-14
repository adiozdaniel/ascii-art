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
	plurals := []string{"This", "was"}
	if count > 1 {
		word = "characters"
		plurals[1] = "were"
		plurals[0] = "These"
	}
	artWork.WriteString(fmt.Sprintf("%s %s %s printed only once!\n", plurals[0], word, plurals[1]))
	return artWork.String()
}
