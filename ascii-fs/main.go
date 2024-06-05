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
	file := "banner.txt"
	reff := ""
	str := ""
	color := "white"
	banner := "standard"
	out := false
	if len(os.Args) == 1 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}

	if strings.Contains(os.Args[1], "--output=") {
		out = true
		if len(os.Args) == 4 {
			file = os.Args[1][9:]
			str = os.Args[2]
			banner = os.Args[3]
		}
		if len(os.Args) == 3 {
			file = os.Args[1][9:]
			str = os.Args[2]
		}
	} else if strings.Contains(os.Args[1], "--color=") {
		if len(os.Args) == 5 && (os.Args[4] == "standard" || os.Args[4] == "thinkertoy" || os.Args[4] == "shadow") {
			color = os.Args[1][8:]
			reff = os.Args[2]
			str = os.Args[3]
			banner = os.Args[4]
		} else if len(os.Args) == 4 && (os.Args[3] == "standard" || os.Args[3] == "thinkertoy" || os.Args[3] == "shadow") {
			color = os.Args[1][8:]
			reff = os.Args[2]
			str = os.Args[2]
			banner = os.Args[3]
		} else if len(os.Args) == 4 && os.Args[3] != "standard" && os.Args[3] != "thinkertoy" && os.Args[3] != "shadow" {
			color = os.Args[1][8:]
			reff = os.Args[2]
			str = os.Args[3]
			banner = "standard"
		} else if len(os.Args) == 3 {
			color = os.Args[1][8:]
			reff = os.Args[2]
			str = os.Args[2]
			banner = "standard"
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		}
	} else if !strings.Contains(os.Args[1], "--color=") {
		if len(os.Args) == 1 || len(os.Args) > 3 {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		}
		if len(os.Args) == 2 {
			str = os.Args[1]
		} else if len(os.Args) == 3 {
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
	// fmt.Print(output, nonAsciis)
	// os.Remove("./banner.txt")

	if out {
		if strings.HasSuffix(file, ".txt") {
			os.WriteFile(file, []byte(output+nonAsciis+"\n"), 0o644)
		} else {
			os.Remove("./banner.txt")
			fmt.Println("Please input a .txt file")
		}
	} else {
		fmt.Print(output, nonAsciis)
		os.Remove("./banner.txt")
	}
}
