package main

import (
	"fmt"
	"os"
	"strings"

	ascii "github.com/adiozdaniel/ascii-art/art_work"
	"github.com/adiozdaniel/ascii-art/utilities"
)

func main() {
	filename := utilities.CorrectFile(os.Args)

	if os.Args[1] == "" {
		return
	}
	if os.Args[1] == "\\n" {
		fmt.Println()
		return
	}

	new := strings.Split(os.Args[1], "\\n")

	non_ascii , effected := utilities.NonAsciiOutput(new)
	if effected {
		fmt.Print(non_ascii)
		return
	}
	fmt.Print(non_ascii)

	_, err := os.Open(filename)
	if err != nil {
		fmt.Println("file error")
		return
	}

	readFile, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	var splitFile []string

	if filename == "../data/thinkertoy.txt" {
		splitFile = strings.Split(string(readFile), "\r\n")
	} else {
		splitFile = strings.Split(string(readFile), "\n")
	}

	output := ascii.Output(new, splitFile)

	fmt.Print(output)
}
