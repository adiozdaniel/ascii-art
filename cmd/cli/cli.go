package main

import (
	"fmt"
	"os"

	"github.com/adiozdaniel/ascii-art/ascii"
	"github.com/adiozdaniel/ascii-art/utils"
)

// Initialize Global Variables
var output = ascii.Output(utils.Inputs.Input)
var nonAsciis = utils.NonAsciiOutput()

func main() {
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
