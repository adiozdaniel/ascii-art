package utils

import (
	"fmt"
	"os"
	"strings"
	// ascii "github.com/adiozdaniel/ascii-art/ascii"
)

func GetInputs() map[string]string {
	var input []string = os.Args[1:]
	inputs := make(map[string]string)

	color, isColor := contains(input, "--color=")
	if isColor && strings.Contains(input[0], "--color=") {
		input = removeItem(input, color)
		color = strings.Split(color, "=")[1]
		if len(color) > 2 {
			inputs["color"] = strings.TrimSpace(color)
		}
	}

	banner, flag := GetFile()
	if flag != "" {
		input = removeItem(input, flag)
		inputs[banner] = flag
	}

	if len(input) == 2 {
		if _, ok := inputs["color"]; !ok {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			os.Exit(0)
		}
		inputs["reff"] = input[0]
		inputs["inputStr"] = input[1]
	}

	if len(input) == 1 {
		inputs["reff"] = input[0]
		inputs["inputStr"] = input[0]
	}

	if len(input) == 0 || len(input) > 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}

	inputs["output"] = "output"
	inputs["justify"] = "justify"
	inputs["reverse"] = "reverse"

	return inputs
}

// removeItem function to remove a specific item from a slice
func removeItem(input []string, item string) []string {
	var index int
	var found bool

	for i, v := range input {
		if v == item {
			index = i
			found = true
			break
		}
	}

	if found {
		input = append(input[:index], input[index+1:]...)
	}

	return input
}

// contains function to check if a slice contains a specific item
func contains(slice []string, item string) (string, bool) {
	for _, v := range slice {
		if strings.HasPrefix(v, item) {
			return v, true
		}
	}
	return "", false
}
