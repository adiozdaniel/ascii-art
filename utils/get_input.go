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
func (i *Input) init() {
	if len(os.Args) < 2 {
		i.PrintUsage()
		os.Exit(1)
	}

	i.ParseArgs(os.Args[1:])
	if i.BannerPath != "" {
		i.isBanner = true
	}

	getFile := i.GetBannerPath()
	fmt.Println("Banner file:", getFile)
}

// CheckInput checks if there is invalid input in the command line arguments.
func CheckInput(input []string) {
	for _, arg := range input {
		if arg == "--" {
			Inputs.Args = append([]string{"--"}, Inputs.Args...)
		}
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
	if Inputs.Justify != "" && len(Inputs.Args) == 0 {
		ErrorHandler("justify")
	}
	if Inputs.Color != "" && len(Inputs.Args) == 0 {
		ErrorHandler("colors")
	}
	if Inputs.Output != "" && len(Inputs.Args) == 0 {
		ErrorHandler("output")
	}

	if Inputs.Output != "" && !strings.HasSuffix(Inputs.Output, ".txt") {
		ErrorHandler("txt")
	}

	if len(Inputs.Args) == 0 {
		return
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
	newInput := strings.Fields(input)

	for _, word := range newInput {
		var temp strings.Builder
		var skipNext bool
		var isSpace bool

		for i := 0; i < len(word); i++ {
			if skipNext {
				skipNext = false
				continue
			}

			if word[i] == '=' && i+2 < len(word) && (word[i+1] == '"' || word[i+1] == '\'') {
				temp.WriteByte('=')
				isSpace = true
				skipNext = true
			} else if word[i] == '\\' && i+1 < len(word) && (word[i+1] == '"' || word[i+1] == '\'') {
				if word[i+1] == '"' {
					temp.WriteByte('"')
				}
				if word[i+1] == '\'' {
					temp.WriteByte('\'')
				}
				skipNext = true
			} else if word[i] == '"' || word[i] == '\'' {
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

func (i *Input) GetBannerPath() string {
	if i.BannerPath != "" {
		return i.BannerPath
	}

	defaultBanner := "../banners/standard.txt"
	if i.isBanner {
		if path, ok := BannerFiles[i.Args[len(i.Args)-1]]; ok {
			return path
		}
	}
	return defaultBanner
}

func (i *Input) Validate() error {
	if i.Color != "" && i.ColorRef == "" {
		return fmt.Errorf("color flag requires a color reference")
	}
	if i.Justify != "" && i.Justify == "" {
		return fmt.Errorf("justify flag requires a justification value")
	}
	if i.Output != "" && !strings.HasSuffix(i.Output, ".txt") {
		return fmt.Errorf("output file must have a .txt extension")
	}
	return nil
}

func (i *Input) ParseArgs(args []string) {
	flag.StringVar(&i.Color, "color", "", "specify a color")
	flag.StringVar(&i.Justify, "align", "", "specify text justification")
	flag.StringVar(&i.Output, "output", "", "specify output file")

	flag.Parse()
	i.Args = flag.Args()

	if err := i.Validate(); err != nil {
		fmt.Println("Error:", err)
		i.PrintUsage()
		os.Exit(1)
	}
}

func (i *Input) PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("  -color <color>    Specify a color")
	fmt.Println("  -align <align>    Specify text justification")
	fmt.Println("  -output <file>    Specify output file")
	fmt.Println("  --shadow          Use shadow banner")
	fmt.Println("  --thinkertoy      Use thinkertoy banner")
	fmt.Println("  --standard        Use standard banner")
}
