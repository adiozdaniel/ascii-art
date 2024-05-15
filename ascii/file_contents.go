package main

import (
	"fmt"
	"os"
	"strings"
)

func fileContentTests() []string {
	fileName := "../data/standard.txt"
	contents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fileContents := strings.Split(string(contents), "\n")
	return fileContents
}
