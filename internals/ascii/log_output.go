package outputs

import (
	"os"
	"path/filepath"

	"github.com/adiozdaniel/ascii-art/pkg/helpers"
)

// LogOutput writes ascii art to a given file
func LogOutput(output string) {
	cleanOutput := helpers.RemoveANSICodes(output)

	outputDir := filepath.Dir(app.Output)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			app.ErrorHandler("restricted")
		}
	}

	err := os.WriteFile(app.Output, []byte(cleanOutput), 0644)
	if err != nil {
		app.ErrorHandler("restricted")
	}
}
