package ascii

import (
	"fmt"
	"os"
	"strings"
	"time"

	utils "github.com/adiozdaniel/ascii-art/utils"
)

// FileContents returns a slice of strings containing ascii artwork characters
func FileContents(banner string) []string {
	fileName := utils.GetFile(banner)
	contents, err := os.ReadFile(fileName)
	var ans string
	if err != nil {
		fmt.Printf("File %v has a problem, would you wish to download it (yes/no): ", fileName[11:])
		fmt.Scan(&ans)
		if ans == "yes" {
			fmt.Print("Be patient while downloading...\n")
			time.Sleep(1 * time.Second)
			err := utils.DownloadFile("https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/"+fileName[11:], fileName)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			fmt.Print("\033[2A")
			// Clear the current line
			fmt.Print("\033[2K")
			// Move to the beginning of the line
			fmt.Print("\033[G")
			// Clear the next line
			fmt.Print("\033[2K")
			fmt.Println("Download Complete, This is awesome right?😎😎😎")
			fmt.Print("\033[2K")
			time.Sleep(2 * time.Second)
			fmt.Print("\033[1A")
			// Clear the current line
			fmt.Print("\033[2K")

			// fmt.Print("\033c")
		} else {
			fmt.Println("Program stopped: No banner file")
			os.Exit(0)
		}
		contents, err = os.ReadFile(fileName)
		if err != nil {
			fmt.Println("not succesfull")
			os.Exit(0)
		}
	}

	var fileContents []string

	if fileName == "../banners/thinkertoy.txt" {
		fileContents = strings.Split(string(contents), "\r\n")
	} else {
		fileContents = strings.Split(string(contents), "\n")
	}
	return fileContents
}
