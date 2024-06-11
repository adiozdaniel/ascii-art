package utils

import (
	"fmt"
	"os"
)

var bannerFiles = map[string]string{
	"-shadow":     "../banners/shadow.txt",
	"shadow":      "../banners/shadow.txt",
	"-thinkertoy": "../banners/thinkertoy.txt",
	"thinkertoy":  "../banners/thinkertoy.txt",
	"-standard":   "../banners/standard.txt",
	"standard":    "../banners/standard.txt",
}

func isBanner(s []string)string{
	for _, word := range s {
		if _, ok := bannerFiles[word]; ok {
			return bannerFiles[word]
		}
	}
	return bannerFiles["standard"]
}

// GetFile returns the ascii graphic filepath to use.
func GetFile() string {
	if len(os.Args) == 1 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}

	args := os.Args[2:]

	return isBanner(args)
}
