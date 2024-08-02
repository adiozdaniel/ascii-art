package outputs

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