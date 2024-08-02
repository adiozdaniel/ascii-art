package ascii

import (
	"os"
	"path/filepath"
	"regexp"
)

// LogOutput writes ascii art to a given file
func LogOutput(output string) {
	cleanOutput := removeANSICodes(output)

	outputDir := filepath.Dir(Inputs.Output)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			ErrorHandler("restricted")
		}
	}

	err := os.WriteFile(Inputs.Output, []byte(cleanOutput), 0644)
	if err != nil {
		ErrorHandler("restricted")
	}
}

// removeANSICodes removes the ansci escape codes
func removeANSICodes(input string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(input, "")
}
