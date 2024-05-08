package utilities

import (
	"fmt"
	"os"
)

func CorrectFile(args []string) string {
	var filename string

	if len(os.Args) > 3 {
		fmt.Println("Incorrect number of arguments. Use: go run . <input_text> <flag>")
		return ""
	}

	if len(os.Args) == 1 {
		fmt.Println("Incorrect number of arguments")
		return ""
	}

	if len(os.Args) == 3 {
		if os.Args[2] == "s" || os.Args[2] == "t" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy" || os.Args[2] == "standard" {
			switch os.Args[2] {
			case "s", "shadow":
				filename = "../data/shadow.txt"
			case "t", "thinkertoy":
				filename = "../data/thinkertoy.txt"
			case "st", "standard":
				filename = "../data/standard.txt"
			default:
				fmt.Println("Available options are shadow.txt and thinkertoy.txt")
			}
		}
	} else {
		filename = "../data/standard.txt"
	}

	return filename
}
