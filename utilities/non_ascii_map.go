package utilities

func NonAsciiMap(splitInput []string) map[rune]bool {
	nonAsciiMap := make(map[rune]bool)

	for _, line := range splitInput {
		for _, char := range line {
			if char < 32 || char > 126 {
				nonAsciiMap[char] = true
			}
		}
	}

	return nonAsciiMap
}