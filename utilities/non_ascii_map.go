package utilities

func NonAsciiMap(splitInput []string) map[rune]int {
	nonAsciiMap := make(map[rune]int)

	for index, line := range splitInput {
		for _, char := range line {
			if char < 32 || char > 126 {
				nonAsciiMap[char] = index + 1
			}
		}
	}

	return nonAsciiMap
}