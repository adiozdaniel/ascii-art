package utils

import (
	"flag"
	"os"
	"strings"
)

type Input struct {
	Color    string
	Banner   string
	isBanner bool
	Justify  string
	Output   string
	Input    string
	Args     []string
}

var Inputs Input

//bannerFiles is a map for banner files and their paths
var bannerFiles = map[string]string{
	"-shadow":     "../banners/shadow.txt",
	"shadow":      "../banners/shadow.txt",
	"-thinkertoy": "../banners/thinkertoy.txt",
	"thinkertoy":  "../banners/thinkertoy.txt",
	"-standard":   "../banners/standard.txt",
	"standard":    "../banners/standard.txt",
}

//init initializes the Input
func init() {
	flag.StringVar(&Inputs.Color, "color", "", "specify a color")
	flag.StringVar(&Inputs.Justify, "justify", "", "specify text justification")
	flag.StringVar(&Inputs.Output, "output", "", "specify output file")
	flag.StringVar(&Inputs.Input, "input", "", "specify your text")

	// Override default Error message from flags
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			flagName := strings.SplitN(arg, "=", 2)[0]
			if _, ok := bannerFiles[flagName]; !ok {
				ErrorHandler()
				os.Exit(2)
			}
		}
	}

	flag.Parse()
	Inputs.Args = flag.Args()
	getFile()
}
