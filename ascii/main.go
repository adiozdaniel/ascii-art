package main

import (
	"fmt"
	"os"
	"strings"

	ascii "github.com/adiozdaniel/ascii-art/art_work"
	utils "github.com/adiozdaniel/ascii-art/utilities"
)

func main() {
	arg := os.Args
	if len(os.Args) > 3 {
		fmt.Println("Incorrect number of arguments")
		return
	}
	if len(os.Args) == 1 {
		fmt.Println("Incorrect number of arguments")
		return
	}
	if arg[1] == "" {
		return
	}
	if arg[1] == "\\n" {
		fmt.Println()
		return
	}

	err := utils.ProtectFilesInDirectory()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	new := strings.Split(arg[1], "\\n")
	filename := "../data/standard.txt"

	input, _ := os.ReadFile(filename)

	var input2 []string
	if filename == "../data/thinkertoy.txt" {
		input2 = strings.Split(string(input), "\r\n")
	} else {
		input2 = strings.Split(string(input), "\n")
	}

	output := ascii.Output(new, input2)

	fmt.Print(output)
}
