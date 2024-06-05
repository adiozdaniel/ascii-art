package utils

import (
	"fmt"
	"os"
)

// GetFile returns the ascii graphic file to use.
func GetFile(args string) string {
	// fmt.Print("\033c")
	// if len(args) == 1 || len(args) > 3 {
	// 	fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
	// 	os.Exit(0)
	// }

	results := "../banners/standard.txt"

	if (args != "-standard") && (args != "-shadow") && (args != "-thinkertoy") {
		fmt.Println(`Invalid banner. Available banner options are:
			 -thinkertoy
			 -shadow
			 -standard`)
		os.Exit(0)
	}
	switch args {
	case "-thinkertoy":
		results = "../banners/thinkertoy.txt"
	case "-shadow":
		results = "../banners/shadow.txt"
	}

	err := ProtectFilesInDirectory()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// return standard.txt as the default incase thinkertoy or shadow was not selected
	return results
}
