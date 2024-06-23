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

//BannerFiles is a map for banner files and their paths
var BannerFiles = map[string]string{
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
	"-color":      true,
	"--justify":   true,
	"--output":    true,
	"-output":     true,
	"-shadow":     true,
	"-thinkertoy": true,
	"-standard":   true,
}

//init initializes the Input
func init() {
	if len(os.Args) < 2 {
		ErrorHandler()
	}

	if strings.Contains(os.Args[0], "test") || os.Args[1] == "-web" {
		return
	}

	flag.StringVar(&Inputs.Color, "color", "", "specify a color")
	flag.StringVar(&Inputs.Justify, "justify", "", "specify text justification")
	flag.StringVar(&Inputs.Output, "output", "", "specify output file")
	flag.StringVar(&Inputs.Input, "thinkertoy", "", "specify thinkertoy as your banner")
	flag.StringVar(&Inputs.Input, "shadow", "", "specify shadow as your banner")
	flag.StringVar(&Inputs.Input, "standard", "", "specify standard as your banner")

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
		Inputs.Args = Inputs.Args[2:]
		return
	}

	if len(Inputs.Args) == 1 {
		Inputs.ColorRef = Inputs.Args[0]
		Inputs.Input = Inputs.Args[0]
	}

	if len(Inputs.Args) > 2 {
		ErrorHandler()
	}
}

// GetFile returns the ascii graphic filepath to use.
func getFile() {
	ourBanner := "../banners/standard.txt"
	args := []string{}

	if len(Inputs.Args) == 0 {
		ErrorHandler()
	}

	if len(Inputs.Args) == 1 {
		Inputs.Banner = ourBanner
		return
	}

	for _, val := range Inputs.Args {
		if !Inputs.isBanner {
			if value, ok := BannerFiles[val]; ok {
				ourBanner = value
				Inputs.isBanner = true
			}
		} else {
			args = append(args, val)
		}
	}

	Inputs.Args = args
	Inputs.Banner = ourBanner
}
