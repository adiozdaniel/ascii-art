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

func isBanner(s []string)(string, string){
	for _, flag := range s {
		if _, ok := bannerFiles[flag]; ok {
			return bannerFiles[flag], flag
		}
	}
	return bannerFiles["standard"], ""
}

// GetFile returns the ascii graphic filepath to use.
func GetFile() (string, string) {
	if len(os.Args) == 1 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}

	args := os.Args[2:]

	return isBanner(args)
}
