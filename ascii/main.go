package main

import (
	"fmt"
	"os"
	"strings"

	ascii "github.com/adiozdaniel/ascii-art/art_work"
)

func main() {
	fileContents := ascii.FileContents()
	output := ascii.Output(strings.Split(os.Args[1], "\\n"), fileContents)
	fmt.Print(output)
}
