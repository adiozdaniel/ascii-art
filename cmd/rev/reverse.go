package reverse

import (
	"fmt"
	"os"
	"strings"

	"github.com/adiozdaniel/ascii-art/internals/models"
)

var (
	sm = models.GetStateManager()
)

func CheckReverse(input string) {
	fileData, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	spaces, lines := RemoveLeadingspace(strings.Split(string(fileData), "\n"))

	asciimap := InitMap()
	result := Reverse(lines, asciimap)
	fmt.Println(strings.Repeat(" ", spaces), result)
}

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

func Reverse(lines []string, asciimap map[string]rune) string {
	final := ""
	res := ""
	for len(lines) != 0 {
		if len(lines[0]) == 0 {
			lines = lines[1:]
			if len(lines) != 0 {
				final += "\\n"
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
						lines = lines[1:]
						final += "\\n"
						break
					}
				}
			}
		} else {
			break
		}
	}
	return final
}

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
