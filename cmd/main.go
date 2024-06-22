package main

import (
	"fmt"
	"strings"

	ascii "github.com/adiozdaniel/ascii-art/ascii"
	utils "github.com/adiozdaniel/ascii-art/utils"
)

func main() {
	fileContents := ascii.FileContents()
	input := utils.Inputs.Input
	output := ascii.Output(utils.Inputs.Color,
		utils.Inputs.ColorRef,
		strings.Split(strings.ReplaceAll(input,
			"\\n", "\n"), "\n"),
		fileContents)
	nonAsciis := utils.NonAsciiOutput(strings.Split(input, "\n"))

	fmt.Print(output, nonAsciis)
}
