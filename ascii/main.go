package main

import (
	"fmt"
	"os"
	"strings"

	ascii "github.com/adiozdaniel/ascii-art/art_work"
	utils "github.com/adiozdaniel/ascii-art/utilities"
)

func main() {
	fileContents := ascii.FileContents()
	input := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	output := ascii.Output(strings.Split(input, "\n"), fileContents)
	nonAsciis := utils.NonAsciiOutput(strings.Split(input, "\n"))
	fmt.Print(output, nonAsciis)
}
