package art_work

func asciiMap(outputFile []string) map[rune]int {
	ascii_map := make(map[rune]int)
	ascii := 32

	for index, line := range outputFile {
		if len(line) == 0 || len(line) == 1 {
			ascii_map[rune(ascii)] = index + 1
			ascii++
		}
	}
	return ascii_map
}
