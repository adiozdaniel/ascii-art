package main

import (
	"fmt"
	"os"
	"strings"

	ascii "github.com/adiozdaniel/ascii-art/ascii"
	utils "github.com/adiozdaniel/ascii-art/utils"
)

func main() {
	// fileContents := ascii.FileContents()
	// // input := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	// // output := ascii.Output(strings.Split(input, "\n"), fileContents)
	// nonAsciis := utils.NonAsciiOutput(strings.Split(input, "\n"))
	// fmt.Print(output, nonAsciis)
	Channel()
}

func Channel() {
	reff := ""
	str := ""
	color := ""
	banner := ""
	if len(os.Args) == 1 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}
	if strings.Contains(os.Args[1], "--color=") {
		// if len(os.Args) == 1 || len(os.Args) > 5 {
		// 	fmt.Println("err")
		// 	os.Exit(1)
		// }
		if len(os.Args) == 8 {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			os.Exit(0)
		}
		if len(os.Args[1]) == 8 {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			os.Exit(0)
		}
		if len(os.Args) == 5 && (os.Args[4] == "-standard" || os.Args[4] == "-thinkertoy" || os.Args[4] == "-shadow") {
			color = os.Args[1][8:]
			reff = os.Args[2]
			str = os.Args[3]
			banner = os.Args[4]
		} else if len(os.Args) == 4 && (os.Args[3] == "-standard" || os.Args[3] == "-thinkertoy" || os.Args[3] == "-shadow") {
			color = os.Args[1][8:]
			reff = os.Args[2]
			str = os.Args[2]
			banner = os.Args[3]
		} else if len(os.Args) == 4 && os.Args[3] != "-standard" && os.Args[3] != "-thinkertoy" && os.Args[3] != "-shadow" {
			color = os.Args[1][8:]
			reff = os.Args[2]
			str = os.Args[3]
			banner = "-standard"
		} else if len(os.Args) == 3 {
			color = os.Args[1][8:]
			reff = os.Args[2]
			str = os.Args[2]
			banner = "-standard"
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			os.Exit(0)
		}
	}
	if !strings.Contains(os.Args[1], "--color=") {
		if len(os.Args) == 1 || len(os.Args) > 3 {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			os.Exit(0)
		}
		if len(os.Args) == 2 {
			color = "white"
			reff = os.Args[1]
			str = os.Args[1]
			banner = "-standard"
		} else if len(os.Args) == 3 {
			color = "white"
			reff = os.Args[1]
			str = os.Args[1]
			banner = os.Args[2]
		}
	}

	if str == "" {
		os.Exit(0)
	}
	if str == "\\n" {
		fmt.Println()
		os.Exit(0)
	}

	fileContents := ascii.FileContents(banner)
	input := strings.ReplaceAll(str, "\\n", "\n")
	output := ascii.Output(color, reff, strings.Split(input, "\n"), fileContents)
	nonAsciis := utils.NonAsciiOutput(strings.Split(input, "\n"))
	fmt.Print(output, nonAsciis)
}
