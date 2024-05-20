package utilities

import (
	"fmt"
	"os"
)

// GetFile returns the ascii graphic file to use.
func GetFile(args []string) string {
	// fmt.Print("\033c")
	if len(args) == 1 || len(args) > 3 {
		fmt.Println("To run the program type go run . <text> [-flag].")
		os.Exit(0)
	}

	if args[1] == "" {
		os.Exit(0)
	}
	if args[1] == "\\n" {
		fmt.Println()
		os.Exit(0)
	}

	results := "../data/standard.txt"

	if len(args) == 3 {
		if (args[2] != "standard") && (args[2] != "shadow") && (args[2] != "-s") && (args[2] != "-thinkertoy") && (args[2] != "thinkertoy") && (args[2] != "-t") {
			fmt.Println(`Invalid flag. Available flag options are:
			-t, -thinkertoy, thinkertoy
			-s, -shadow, shadow`)
			os.Exit(0)
		}
		switch args[2] {
		case "-t", "-thinkertoy", "thinkertoy":
			results = "../data/thinkertoy.txt"
		case "-s", "-shadow", "shadow":
			results = "../data/shadow.txt"
		case "-st", "-standard", "standard":
			results = "../data/standard.txt"
		}
	}
	err := ProtectFilesInDirectory()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// If there was a typing typo, return standard.txt
	return results
}
