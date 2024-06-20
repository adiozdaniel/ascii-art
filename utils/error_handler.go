package utils

import (
	"fmt"
	"os"
)

const (
	usage = "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\""
)

// ErrorHandler outputs errors and safely exits the program
func ErrorHandler() {
	fmt.Println(usage)
	os.Exit(0)
}
