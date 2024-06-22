package utils

import (
	"flag"
	"os"
	"strings"
)

type Input struct {
	Color    string
	ColorRef string
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

//validFlags stores allowed flags
var validFlags = map[string]bool{
	"--color":     true,
	"--justify":   true,
	"--output":    true,
	"-shadow":     true,
	"-thinkertoy": true,
	"-standard":   true,
}

//init initializes the Input
func init() {
	flag.StringVar(&Inputs.Color, "color", "", "specify a color")
	flag.StringVar(&Inputs.Justify, "justify", "", "specify text justification")
	flag.StringVar(&Inputs.Output, "output", "", "specify output file")
	flag.StringVar(&Inputs.Input, "input", "", "specify your text")

	flag.Usage = func() {
		ErrorHandler()
	}

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			flagName := strings.SplitN(arg, "=", 2)[0]
			if !validFlags[flagName] {
				flag.Usage()
			}
		}
	}

	flag.Parse()
	Inputs.Args = flag.Args()
	getFile()
	if len(Inputs.Args) == 2 {
		Inputs.ColorRef = Inputs.Args[0]
		Inputs.Input = Inputs.Args[1]
		return
	}

	if len(Inputs.Args) == 1 && Inputs.Color != "" {
		Inputs.ColorRef = Inputs.Args[0]
		Inputs.Input = Inputs.Args[0]
		return
	}
}
