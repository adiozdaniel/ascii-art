package art_work

import (
	"fmt"
	"os"
	"strings"

	utils "github.com/adiozdaniel/ascii-art/utilities"
)

// FileContents returns a slice of strings containing ascii artwork characters
func FileContents() []string {
	fileName := utils.CorrectFile(os.Args)
	contents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	var fileContents []string

	if fileName == "../data/thinkertoy.txt" {
		fileContents = strings.Split(string(contents), "\r\n")
	} else {
		fileContents = strings.Split(string(contents), "\n")
	}
	return fileContents
}
