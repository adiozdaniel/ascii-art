package utilities

import (
	"fmt"
	"os"
)

func CorrectFile(args []string) string {
	var filename string

	if len(os.Args) > 3 {
		fmt.Println("Incorrect number of arguments. Use: go run . <input_text> <flag>")
		os.Exit(0)
	}

	if len(os.Args) == 1 {
		fmt.Println("Incorrect number of arguments. Use: go run . <input_text> <flag>")
		os.Exit(0)
	}

	if len(os.Args) == 3 {
		if os.Args[2] == "s" || os.Args[2] == "t" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy" || os.Args[2] == "standard" || os.Args[2] == "st" {
			switch os.Args[2] {
			case "s", "shadow":
				filename = "../data/shadow.txt"
			case "t", "thinkertoy":
				filename = "../data/thinkertoy.txt"
			case "st", "standard":
				filename = "../data/standard.txt"
			}
		} else {
			fmt.Println("Available flag options are:\ns & shadow\nt & thinkertoy\nst & standard")
			os.Exit(0)
		}
	} else {
		filename = "../data/standard.txt"
	}

	return filename
}
