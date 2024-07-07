package main

import (
	"fmt"
	"os"

	"github.com/adiozdaniel/ascii-art/ascii"
	"github.com/adiozdaniel/ascii-art/utils"
)

var fileContents, _ = ascii.FileContents()
var output = ascii.Output(fileContents)
var nonAsciis = utils.NonAsciiOutput()

func main() {
	if os.Args[1] == "-web" {
		runWeb()
	}

	if utils.Inputs.Output != "" {
		runOutput()
		return
	}

	if utils.Inputs.Justify != "" {
		justified()
		return
	}

	fmt.Print(output, nonAsciis)
}
