package utils

//NonAsciiMap maps the non-printable ascii characters
func NonAsciiMap(splitInput []string) map[rune]rune {
	nonAsciiMap := make(map[rune]rune)

	for _, line := range splitInput {
		for _, char := range line {
			if char < 32 || char > 126 {
				nonAsciiMap[char] = char
			}
		}
	}

	return nonAsciiMap
}
