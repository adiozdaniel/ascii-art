package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/adiozdaniel/ascii-art/pkg/utils"
)

// global variables declaration
var app = utils.NewInput()

// constants store the hashed sum of the banner files as hexadecimal strings
const (
	standard   = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	thinkertoy = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
	shadow     = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
)

// validateBanner calculates the characters to verify the validity of the banner file
func validateBanner(banner []byte) bool {
	hasher := sha256.New()
	hasher.Write(banner)
	hash := hex.EncodeToString(hasher.Sum(nil))

	if hash == standard || hash == thinkertoy || hash == shadow {
		return true
	}
	return false
}

// FileContents returns a slice of strings containing ascii artwork characters
func FileContents(fileName string) error {
	fileDir := "views/static/banners"
	filePath := app.GetProjectRoot(fileDir, fileName)

	contents, err := os.ReadFile(filePath)
	if err != nil || !validateBanner(contents) {
		if app.Flags["isWeb"] == "true" {
			return fmt.Errorf("not found")
		}
		fmt.Print("Be patient while downloading...\n")
		time.Sleep(1 * time.Second)
		err := app.DownloadFile("https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/"+fileName[11:], filePath)
		if err != nil {
			fmt.Println(err)
			os.Remove(filePath)
			os.Exit(0)
		}
		fmt.Print("\033[1A", "\033[G", "\033[2K")
		fmt.Println("Download Complete, This is awesome right?😎😎😎")
		fmt.Print("\033[2K")
		time.Sleep(2 * time.Second)
		fmt.Print("\033[1A", "\033[2K")

		contents, err = os.ReadFile(fileName)
		if err != nil {
			fmt.Println("not succesfull")
			os.Exit(0)
		}
	}

	var fileContents []string

	if fileName == "thinkertoy.txt" {
		fileContents = strings.Split(string(contents), "\r\n")
	} else {
		fileContents = strings.Split(string(contents), "\n")
	}

	app.FileContents = append(app.FileContents, fileContents...)
	return nil
}
