package models

import (
	"strings"
)

func AsciiMapper(fileContents []string) map[rune]int {
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

var input = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func CharacterBuilder(ascii_map map[rune]int, fileContents []string) map[string]string {
	var characterMap = make(map[string]string)

	for _, char := range input {
		var builder strings.Builder
		for i := 0; i < 8; i++ {
			if ascii, ok := ascii_map[char]; ok {
				builder.WriteString(fileContents[ascii+i])
			}
			builder.WriteString(builder.String())
			builder.WriteRune('\n')
		}

		characterMap[string(char)] = builder.String()
	}
	return characterMap
}
