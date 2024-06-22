package utils

import (
	"flag"
	"fmt"
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
	fmt.Println(Inputs.Args)

	if len(Inputs.Args) > 1 {
		Inputs.ColorRef = Inputs.Args[0]
		Inputs.Input = Inputs.Args[1]
		Inputs.Args = Inputs.Args[2:]
	}

	if len(Inputs.Args) == 1 && Inputs.Color != "" {
		Inputs.ColorRef = Inputs.Args[0]
		Inputs.Input = Inputs.Args[0]
	}
	getFile()
}

// GetFile returns the ascii graphic filepath to use.
func getFile() {
	ourBanner := "../banners/standard.txt"
	flag := ""

	for _, val := range Inputs.Args {
		if value, ok := bannerFiles[val]; ok {
			ourBanner = value
			flag = val
		}
	}

	Inputs.Banner = ourBanner
	if flag == "" {
		Inputs.isBanner = false
	} else {
		Inputs.isBanner = true
	}
}
