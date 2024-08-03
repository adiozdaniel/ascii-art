package ascii

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

	outputString := fmt.Sprintf("  😋 writing.... '%s'. To check output, kindly use: `cat %s | cat -e` %s", app.Input, app.Output, NonAsciiOutput())
	fmt.Printf("%s\n %s\n", outputString, strings.Repeat("=", len(outputString)-3))
}
