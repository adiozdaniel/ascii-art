package utils

import (
	"fmt"
	"os"
	"regexp"
)

func LogOutput(output string) {
	// Open the log file in append mode, create it if it doesn't exist
	file, err := os.OpenFile(Inputs.Output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	// Remove ANSI escape codes
	cleanOutput := removeANSICodes(output)

	// Write the cleaned output to the log file
	_, err = file.WriteString(cleanOutput + "\n")
	if err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}

func removeANSICodes(input string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(input, "")
}
