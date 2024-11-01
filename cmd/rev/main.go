package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go --reverse=<file>")
		return
	}
	filename := validate(os.Args[1])
	fileData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(fileData), "\n")

	// fmt.Println(len(lines))
	asciimap := InitMap()
	// fmt.Println(len(asciimap))
	final := ""
	res := ""
	for len(lines) != 0 {
		if  len(lines[0]) == 0 {
			lines = lines[1:]
			if len(lines) != 0 {
				final += "\\n"
			}
		} else {
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
		}
	}
	fmt.Println(final)
	// fmt.Println(len(lines))
}

func validate(input string) string {
	if !strings.HasPrefix(input, "--reverse=") {
		fmt.Println("Usage: go run main.go --reverse=<file>")
		os.Exit(1)
	}
	file := strings.TrimPrefix(input, "--reverse=")
	if !strings.HasSuffix(file, ".txt") {
		fmt.Println("File must have a .txt extension")
		os.Exit(1)
	}
	return file
}

func InitMap() map[string]rune {
	asciimap := make(map[string]rune)
	files := []string{"standard.txt","shadow.txt", "thinkertoy.txt"}
	for _, file := range files {
		var char rune = 31
		ascii := ""
		contents, err := Filecontents(file)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		for _, line := range contents {
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

func Filecontents(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if filename == "thinkertoy.txt" {
		return strings.Split(string(content), "\r\n"), nil
	}
	return strings.Split(string(content), "\n"), nil
}
