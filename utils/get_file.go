package utils

import (
	"fmt"
	"os"
)

//bannerFiles is a map for banner files and their paths
var bannerFiles = map[string]string{
	"-shadow":     "../banners/shadow.txt",
	"shadow":      "../banners/shadow.txt",
	"-thinkertoy": "../banners/thinkertoy.txt",
	"thinkertoy":  "../banners/thinkertoy.txt",
	"-standard":   "../banners/standard.txt",
	"standard":    "../banners/standard.txt",
}

//isBanner returns the appropriate banner file path and the banner name
func isBanner(s []string) (string, string) {
	ourBanner := "../banners/standard.txt"
	flag := ""

	for _, val := range s {
		if value, ok := bannerFiles[val]; ok {
			ourBanner = value
			flag = val
		}
	}

	return ourBanner, flag
}

// GetFile returns the ascii graphic filepath to use.
func GetFile() (string, string) {
	if len(os.Args) == 1 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}

	args := os.Args[1:]

	return isBanner(args)
}
