package utils

import (
	"flag"
	"os"
	"strings"
)

//Input holds all the data passed around in the application
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

//Inputs is a placeholder for the Input struct
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
		ErrorHandler("fatal")
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
		ErrorHandler("fatal")
	}

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && strings.Contains(arg, "=") {
			flagValue := strings.Split(arg, "=")[1]
			flagName := strings.Split(arg, "=")[0]
			if !validFlags[flagName] {
				flag.Usage()
			}
			if flagValue == "" {
				flag.Usage()
			}
		}
	}

	flag.Parse()
	Inputs.Args = flag.Args()

	if Inputs.Banner != "" {
		Inputs.isBanner = true
	}

	getFile()

	if len(Inputs.Args) == 2 && Inputs.Color != "" {
		Inputs.ColorRef = Inputs.Args[0]
		Inputs.Input = Inputs.Args[1]
		Inputs.Args = Inputs.Args[2:]
		return
	}

	if Inputs.Output != "" && (len(Inputs.Args) != 1) {
		ErrorHandler("output")
	}

	if Inputs.Output != "" && Inputs.Color != "" {
		ErrorHandler("output")
	}

	if len(Inputs.Args) == 1 {
		Inputs.ColorRef = Inputs.Args[0]
		Inputs.Input = Inputs.Args[0]
	}

	if len(Inputs.Args) > 1 {
		ErrorHandler("fatal")
	}
}

// GetFile returns the ascii graphic filepath to use.
func getFile() {
	if len(Inputs.Args) == 0 {
		ErrorHandler("fatal")
	}

	ourBanner := "../banners/standard.txt"
	if len(Inputs.Args) == 1 {
		Inputs.Banner = ourBanner
		return
	}

	if !Inputs.isBanner {
		if value, ok := BannerFiles[Inputs.Args[len(Inputs.Args)-1]]; ok {
			Inputs.Banner = value
			Inputs.isBanner = true
			Inputs.Args = Inputs.Args[:len(Inputs.Args)-1]
			return
		}
	}

	Inputs.Banner = ourBanner
}
