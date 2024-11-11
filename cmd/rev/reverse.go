package reverse

import (
	"fmt"
	"os"
	"strings"

	"github.com/adiozdaniel/ascii-art/internals/models"
)

// sm is an instance of the state manager obtained from models
var (
	sm = models.GetStateManager()
)

// CheckReverse reads a file, processes its contents by removing leading spaces, reverses it using an ASCII map, and formats the result
func CheckReverse(input string) (string, error) {
	fileData, err := os.ReadFile(input)
	if err != nil {
		return "", err
	}
	spaces, lines := RemoveLeadingspace(strings.Split(string(fileData), "\n"))

	asciimap := InitMap()
	result, err := Reverse(lines, asciimap)
	if err != nil {
		return "", err
	}

	spaced := strings.Repeat(" ", spaces)
	formated := fmt.Sprintf("%s%s\n\n", spaced, result)

	return formated, nil
}

// RemoveLeadingspace calculates the number of leading spaces in a set of lines and removes them, returning the adjusted lines
func RemoveLeadingspace(lines []string) (int, []string) {
	if len(lines) < 8 {
		return 0, nil
	}
	for i := 0; i < len(lines[0]); i++ {
		for j := 0; j < 8 && len(lines[j]) == len(lines[0]); j++ {
			if lines[j][i] != ' ' {
				spaces := 0
				if i >= 6 {
					spaces = i / 6
				}
				for k := 0; k < 8; k++ {
					lines[k] = lines[k][i:]
				}
				return spaces, lines
			}
		}
	}
	return 0, lines
}

// Reverse processes a list of strings by mapping ASCII characters and building a reversed output, returning the formatted result
func Reverse(lines []string, asciimap map[string]rune) (string, error) {
	final := ""
	res := ""
	for len(lines) != 0 {
		if len(lines[0]) == 0 {
			lines = lines[1:]
			if len(lines) != 0 {
				final += "\n"
			}
		} else if len(lines) >= 8 {
			start := 0
			for i := 1; i <= len(lines[0]); i++ {
				character := ""
				for j := 0; j < 8 && len(lines[j]) == len(lines[0]); j++ {
					character += lines[j][start:i] + "\n"
				}

				if char, ok := asciimap[character]; ok {
					res += string(char)
					start = i
				}
				if i == len(lines[0]) {
					if res != "" {
						lines = lines[8:]
						final += res
						if len(lines) >= 8 {
							final += "\\n"
						}
						res = ""
						break
					} else {
						final += "\\n"
						return final, fmt.Errorf("unexpected characters found in your file")
					}
				}
			}
		} else {
			break
		}
	}
	return final, nil
}

// InitMap initializes a map that associates ASCII character representations with their corresponding rune values
func InitMap() map[string]rune {
	asciimap := make(map[string]rune)
	files := []string{"standard.txt", "shadow.txt", "thinkertoy.txt"}

	for _, file := range files {
		var char rune = 31
		ascii := ""
		contents := sm.GetConfig().BannerFileCache[file]

		for _, line := range strings.Split(string(contents), "\n") {
			if line == "" {
				asciimap[ascii] = char
				char++
				ascii = ""
			} else {
				ascii += line + "\n"
			}
		}
	}
	return asciimap
}
