package utils

import (
	"fmt"
	"os"
)

//constant used as an output value in ErrorHandler
const (
	fatal   = "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\" standard"
	color   = "🤯 Oops!We couldn't recognise your color\n\nKindly search supported colors here: https://htmlcolorcodes.com/"
	outputs = "🧐 Oops!We can't find your \"something\"\n\nEX: go run . [OPTION] [STRING] \"something\""
	output  = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<banner.txt> \"something\" standard"
)

// ErrorHandler outputs errors and safely exits the program
func ErrorHandler(errType string) {
	fmt.Println(errType)
	os.Exit(0)
}
