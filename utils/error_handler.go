package utils

import (
	"fmt"
	"os"
	"strings"
)

// errors is a map of error output value in ErrorHandler
var errors = map[string]string{
	"colors":     "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\" standard",
	"justify":    "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard",
	"color":      "🤯 Oops! We couldn't recognise your color\n\nKindly search supported colors here: https://htmlcolorcodes.com/",
	"output":     "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard",
	"txt":        "😣 Oops! We currently only support text files\n\nSee Documentation in: ../README.md",
	"web":        "😮 Oops! Something went wrong",
	"restricted": "😣 Oops! This is a restricted path: '../banner/'",
}

// ErrorHandler outputs errors and safely exits the program
func ErrorHandler(errType string) {
	if errType == "fatal" {
		fmt.Printf("For color:\n%s\n", strings.Split(errors["colors"], "\n")[2])
		fmt.Printf("For output:\n%s\n", strings.Split(errors["output"], "\n")[2])
		fmt.Printf("For justify:\n%s\n", strings.Split(errors["justify"], "\n")[2])
		fmt.Println("For web:\ngo run . -web")
		os.Exit(0)
	}
	fmt.Println(errors[errType])
	os.Exit(0)
}
