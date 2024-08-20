package tests

import (
	"fmt"
	"os"
	"strings"
)

// FileContentTests checks the content of a bannerfile
func FileContentTests() []string {
	path := app.GetProjectRoot("/views/static/banners", "standard.txt")
	fileName := path
	contents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fileContents := strings.Split(string(contents), "\n")
	return fileContents
}
