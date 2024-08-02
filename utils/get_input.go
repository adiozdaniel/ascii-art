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
	fmt.Println("Banner file:", i.GetBannerPath())
}

// CheckInput checks if there is invalid input in the command line arguments.
func (i *Input) CheckInput() {
	for _, arg := range i.Args {
		if !validFlags[arg] {
			fmt.Printf("Error: Invalid flag %s\n", arg)
			i.PrintUsage()
			os.Exit(1)
		}
	}
}

// GetBannerPath returns the path to the banner file.
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

// Validate checks if the Input contains valid arguments and flags.
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

// ParseArgs parses command-line arguments and sets Input fields.
func (i *Input) ParseArgs(args []string) {
	flag.StringVar(&i.Color, "color", "", "specify a color")
	flag.StringVar(&i.Justify, "align", "", "specify text justification")
	flag.StringVar(&i.Output, "output", "", "specify output file")

	flag.Parse()
	i.Args = flag.Args()

	// Apply RemoveQuotes to relevant fields
	i.Color = RemoveQuotes(i.Color)
	i.Justify = RemoveQuotes(i.Justify)
	i.Output = RemoveQuotes(i.Output)

	for index := range i.Args {
		i.Args[index] = RemoveQuotes(i.Args[index])
	}

	i.CheckInput()

	if err := i.Validate(); err != nil {
		fmt.Println("Error:", err)
		i.PrintUsage()
		os.Exit(1)
	}
}

// PrintUsage prints the usage instructions for the application.
func (i *Input) PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("  -color <color>    Specify a color")
	fmt.Println("  -align <align>    Specify text justification")
	fmt.Println("  -output <file>    Specify output file")
	fmt.Println("  --shadow          Use shadow banner")
	fmt.Println("  --thinkertoy      Use thinkertoy banner")
	fmt.Println("  --standard        Use standard banner")
}
