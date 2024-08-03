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
	"restricted": "😣 Oops! this is a restricted path.\nplease use another path.",
}

// ErrorHandler outputs errors and safely exits the program
func (i *Input) ErrorHandler(errType string) {
	if errType == "fatal" {
		path, _ := os.Getwd()

		if strings.HasSuffix(path, "cli") {
			fmt.Println("For cli mode\ngo run .\n\nFor web interface:\nNavigate back to web directory\ngo run .\n\n!")
		}

		if strings.HasSuffix(path, "web") {
			fmt.Println("For web interface:\ngo run .\n\nFor cli mode:\nNavigate back to cli directory\ngo run .\n\n!")
		}
		os.Exit(1)
	}

	for _, err := range errors {
		if strings.Contains(errType, err) {
			fmt.Println(err)
		}
	}
}
