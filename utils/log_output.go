package utils

import (
	"fmt"
	"os"
	"regexp"
)

// LogOutput writes ascii art to a given file
func LogOutput(output string) {
	file, err := os.OpenFile(Inputs.Output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	cleanOutput := removeANSICodes(output)

	_, err = file.WriteString(cleanOutput)
	if err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}

// removeANSICodes removes the ansci escape codes
func removeANSICodes(input string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(input, "")
}
