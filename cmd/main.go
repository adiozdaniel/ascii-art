package main

import (
	"fmt"
	"os"

	"github.com/adiozdaniel/ascii-art/ascii"
	"github.com/adiozdaniel/ascii-art/utils"
)

func main() {
	if os.Args[1] == "-web" {
		runWeb()
	}

	fileContents, _ := ascii.FileContents()
	output := ascii.Output(fileContents)
	nonAsciis := utils.NonAsciiOutput()

	if utils.Inputs.Output != "" {
		runOutput(output, nonAsciis)
		return
	}

	if utils.Inputs.Justify != "" {
		utils.Alignment(output)
		return
	}

	fmt.Print(output, nonAsciis)
}
