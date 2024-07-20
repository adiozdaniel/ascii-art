package utils

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Input holds all the data passed around in the application
type Input struct {
	Color      string
	ColorRef   string
	BannerPath string
	isBanner   bool
	Justify    string
	Output     string
	Input      string
	IsWeb      bool
	Args       []string
}

// Inputs is a placeholder for the Input struct
var Inputs Input

// BannerFiles is a map for banner files and their paths
var BannerFiles = map[string]string{
	"-shadow":      "../banners/shadow.txt",
	"--shadow":     "../banners/shadow.txt",
	"shadow":       "../banners/shadow.txt",
	"-thinkertoy":  "../banners/thinkertoy.txt",
	"--thinkertoy": "../banners/thinkertoy.txt",
	"thinkertoy":   "../banners/thinkertoy.txt",
	"-standard":    "../banners/standard.txt",
	"--standard":   "../banners/standard.txt",
	"standard":     "../banners/standard.txt",
}

// validFlags stores allowed flags
var validFlags = map[string]bool{
	"--color":      true,
	"--align":      true,
	"--output":     true,
	"-shadow":      true,
	"--shadow":     true,
	"-thinkertoy":  true,
	"--thinkertoy": true,
	"-standard":    true,
	"--standard":   true,
}

// init initializes the Input
func init() {
	if len(os.Args) < 2 {
		ErrorHandler("fatal")
	}

	if strings.Contains(os.Args[0], "test") || os.Args[1] == "-web" {
		Inputs.IsWeb = true
		return
	}

	flag.StringVar(&Inputs.Color, "color", "", "specify a color")
	flag.StringVar(&Inputs.Justify, "align", "", "specify text justification")
	flag.StringVar(&Inputs.Output, "output", "", "specify output file")

	flag.Usage = func() {
		fmt.Print("\033[1A")
		fmt.Print("\033[2K")
		ErrorHandler("fatal")
	}

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && strings.Contains(arg, "=") {
			flagValue := strings.Split(arg, "=")[1]
			flagName := strings.Split(arg, "=")[0]
			if !validFlags[flagName] {
				ErrorHandler("fatal")
			}
			if flagValue == "" {
				if flagName == "--output" {
					ErrorHandler("output")
				}
				if flagName == "--align" {
					ErrorHandler("justify")
				}
				if flagName == "--color" {
					ErrorHandler("colors")
				}
				ErrorHandler("fatal")
			}
		}
	}

	flag.Parse()
	Inputs.Args = flag.Args()

	CheckInput(os.Args[1:])

	if Inputs.BannerPath != "" {
		Inputs.isBanner = true
	}

	getFile()

	if len(Inputs.Args) == 2 && Inputs.Color != "" {
		Inputs.ColorRef = strings.TrimSpace(Inputs.Args[0])
		Inputs.Input = Inputs.Args[1]
		Inputs.Args = Inputs.Args[2:]
		return
	}

	if strings.Contains(Inputs.Output, "/banners/") {
		ErrorHandler("restricted")
	}

	if Inputs.Output != "" && (len(Inputs.Args) != 1) {
		ErrorHandler("output")
	}

	if Inputs.Output != "" && Inputs.Color != "" {
		ErrorHandler("output")
	}

	if Inputs.Output != "" && Inputs.Justify != "" {
		ErrorHandler("output")
	}

	if len(Inputs.Args) == 1 {
		Inputs.Input = Inputs.Args[0]
		Inputs.ColorRef = strings.TrimSpace(Inputs.Args[0])
	}

	if len(Inputs.Args) > 1 {
		if Inputs.Color != "" {
			ErrorHandler("colors")
		}

		if Inputs.Output != "" {
			ErrorHandler("output")
		}

		if Inputs.Justify != "" {
			ErrorHandler("align")
		}
		ErrorHandler("fatal")
	}
}

func CheckInput(input []string) {
	for _, arg := range input {
		if Inputs.Output != "" && Inputs.Output == arg {
			ErrorHandler("output")
		}
		if Inputs.Color != "" && Inputs.Color == arg {
			ErrorHandler("colors")
		}
		if Inputs.Justify != "" && Inputs.Justify == arg {
			ErrorHandler("justify")
		}
	}
}

// GetFile returns the ascii graphic filepath to use.
func getFile() {
	if len(Inputs.Args) == 0 {
		return
	}
	if Inputs.Color != "" && len(Inputs.Args) == 0 {
		ErrorHandler("fatal")
	}
	if Inputs.Output != "" && len(Inputs.Args) == 0 {
		ErrorHandler("output")
	}

	if Inputs.Output != "" && !strings.HasSuffix(Inputs.Output, ".txt") {
		ErrorHandler("txt")
	}

	ourBanner := "../banners/standard.txt"
	if len(Inputs.Args) == 1 {
		Inputs.BannerPath = ourBanner
		return
	}

	if !Inputs.isBanner {
		if value, ok := BannerFiles[Inputs.Args[len(Inputs.Args)-1]]; ok {
			Inputs.BannerPath = value
			Inputs.isBanner = true
			Inputs.Args = Inputs.Args[:len(Inputs.Args)-1]
			return
		}
	}

	Inputs.BannerPath = ourBanner
}

// RemoveQuotes removes opening or closing quotes in a string
func RemoveQuotes(input string) string {
	var result strings.Builder
	var newInput = strings.Fields(input)

	for _, word := range newInput {
		var temp strings.Builder
		var skipNext bool
		var isSpace bool

		for i := 0; i < len(word); i++ {
			if skipNext {
				skipNext = false
				continue
			}

			if word[i] == '=' && i+2 < len(word) && word[i+1] == '"' {
				temp.WriteByte('=')
				isSpace = true
				skipNext = true
			} else if word[i] == '\\' && i+1 < len(word) && word[i+1] == '"' {
				temp.WriteByte('"')
				skipNext = true
			} else if word[i] == '"' {
				if isSpace {
					temp.WriteByte(word[i])
					isSpace = false
				}
			} else {
				temp.WriteByte(word[i])
			}
		}
		if isSpace {
			result.WriteString(temp.String())
		} else {
			result.WriteString(temp.String() + " ")
		}
	}
	return strings.TrimSpace(result.String())
}
