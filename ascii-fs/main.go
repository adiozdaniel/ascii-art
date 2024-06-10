package main

import (
	"fmt"
	"os"
	"strings"

	ascii "github.com/adiozdaniel/ascii-art/ascii"
	utils "github.com/adiozdaniel/ascii-art/utils"
)

func main() {
	color, reff, str, banner := parseArgs(os.Args)

	fileContents := ascii.FileContents(banner)
	input := strings.ReplaceAll(str, "\\n", "\n")
	output := ascii.Output(color, reff, strings.Split(input, "\n"), fileContents)
	nonAsciis := utils.NonAsciiOutput(strings.Split(input, "\n"))
	fmt.Print(output, nonAsciis)
}

func parseArgs(args []string) (color, reff, str, banner string) {
	if len(args) == 1 || len(args) > 5 {
		utils.PrintUsage()
	}

	if strings.Contains(args[1], "--color=") {
		color, reff, str, banner = parseColorArgs(args)
	} else {
		color, reff, str, banner = parseStandardArgs(args)
	}

	if str == "" || str == "\\n" {
		os.Exit(0)
	}

	return color, reff, str, banner
}

func parseColorArgs(args []string) (color, reff, str, banner string) {
	color = args[1][8:]
	reff = args[2]
	str = args[2]

	if len(args) == 5 && (args[4] == "-standard" || args[4] == "-thinkertoy" || args[4] == "-shadow") {
		banner = args[4]
	} else if len(args) == 4 && (args[3] == "-standard" || args[3] == "-thinkertoy" || args[3] == "-shadow") {
		banner = args[3]
	} else if len(args) == 4 && args[3] != "-standard" && args[3] != "-thinkertoy" && args[3] != "-shadow" {
		str = args[3]
		banner = "-standard"
	} else if len(args) == 3 {
		banner = "-standard"
	}

	return color, reff, str, banner
}

func parseStandardArgs(args []string) (color, reff, str, banner string) {
	if len(args) == 2 {
		color = "white"
		reff = args[1]
		str = args[1]
		banner = "-standard"
	} else if len(args) == 3 {
		color = "white"
		reff = args[1]
		str = args[1]
		banner = args[2]
	}

	return color, reff, str, banner
}

