package utils

import (
	"fmt"
	"os"
)

// errors is a map of error output value in ErrorHandler
var errors = map[string]string{
	"fatal":   "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\" standard",
	"align":   "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\" standard",
	"color":   "🤯 Oops! We couldn't recognise your color\n\nKindly search supported colors here: https://htmlcolorcodes.com/",
	"outputs": "🧐 Oops! We can't find your \"something\"\n\nEX: go run . [OPTION] [STRING] \"something\"",
	"output":  "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard",
	"txt":     "😣Oops! We currently only support text files\n\nSee Documentation in: ../README.md",
}

// ErrorHandler outputs errors and safely exits the program
func ErrorHandler(errType string) {
	fmt.Println(errors[errType])
	os.Exit(0)
}
