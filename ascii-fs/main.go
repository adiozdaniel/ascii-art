package main

import (
	"fmt"
	// "os"
	"strings"

	ascii "github.com/adiozdaniel/ascii-art/ascii"
	utils "github.com/adiozdaniel/ascii-art/utils"
)

func main() {
	fileContents := ascii.FileContents()
	inputs := utils.GetInputs()
	output := ascii.Output(inputs["color"],
		inputs["reff"],
		strings.Split(strings.ReplaceAll(inputs["inputStr"],
			"\\n", "\n"), "\n"),
		fileContents)
	nonAsciis := utils.NonAsciiOutput(strings.Split(inputs["inputStr"], "\n"))

	fmt.Print(output, nonAsciis)
}
