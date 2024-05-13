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
	output := ascii.Output(strings.Split(os.Args[1], "\\n"), fileContents)
	nonAsciis := utils.NonAsciiOutput(strings.Split(os.Args[1], "\\n"))
	fmt.Print(output, nonAsciis)
}
