package utils

import (
	"fmt"
	"os"
)

func PrintUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
	os.Exit(0)
}
