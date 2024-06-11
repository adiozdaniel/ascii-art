package main

import (
	"fmt"
	"os"
	"strings"

	ascii "github.com/adiozdaniel/ascii-art/ascii"
	utils "github.com/adiozdaniel/ascii-art/utils"
)

func main() {
	fileContents := ascii.FileContents()
	input := strings.ReplaceAll(strings.Join(os.Args[1:], "\n"), "\\n", "\n")
	fmt.Println(input)
	output := ascii.Output("reff", strings.Split(input, "\n"), fileContents)
	nonAsciis := utils.NonAsciiOutput(strings.Split(input, "\n"))
	fmt.Print(output, nonAsciis)
	
	// Channel()
}

// func Channel() {
	
// 	fileContents := ascii.FileContents(banner)
// 	input := strings.ReplaceAll(str, "\\n", "\n")
// 	output := ascii.Output(color, reff, strings.Split(input, "\n"), fileContents)
// 	nonAsciis := utils.NonAsciiOutput(strings.Split(input, "\n"))
// 	fmt.Print(output, nonAsciis)
// }
